package main

import (
	"bytes"
	"compress/zlib"
	"context"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	log "github.com/sirupsen/logrus"
	"github.com/subzerobo/groot/protos"
	"math/rand"
	"net"
	"os"
	"time"
	
	"github.com/urfave/cli"
	
	"github.com/gemnasium/logrus-graylog-hook"
	"go.elastic.co/apm/transport"
	_ "gopkg.in/Graylog2/go-gelf.v2/gelf"
)

var (
	// Version is injected by buildflags
	VERSION = "SELFBUILD"
)

func main() {
	// Randomize Timer
	rand.Seed(int64(time.Now().Nanosecond()))
	
	appFlags := []cli.Flag{
		cli.StringFlag{
			Name:   "gelf, i",
			Value:  "127.0.0.1:2515",
			Usage:  "GELF Logger server full url",
			EnvVar: "APP_GROOT_GELF_SERVER_ADDR",
		},
		cli.StringFlag{
			Name:   "listen,l",
			Value:  ":1113",
			Usage:  "UDP server listen `PORT` for application",
			EnvVar: "APP_GROOT_PORT",
		},
		cli.BoolFlag{
			Name:   "quiet",
			Usage:  "to suppress the 'stream open/close' messages",
			EnvVar: "APP_GROOT_QUITE",
		},
		cli.IntFlag{
			Name:   "size",
			Value:  4194304, // socket buffer size in bytes
			Usage:  "per-socket buffer size in bytes",
			EnvVar: "APP_GROOT_BUFFER_SIZE",
		},
		cli.StringFlag{
			Name:   "config, c",
			Value:  "", // when the value is not empty, the config path must exists
			Usage:  "config from json file, which will override the command from shell",
			EnvVar: "APP_GROOT_CONFIG_FILE",
		},
	}
	
	app := cli.NewApp()
	app.Name = "Groot"
	app.Usage = "PHP Elastic APM Agent [ UDP Worker ]"
	app.Version = VERSION
	app.Flags = appFlags
	
	app.Action = func(c *cli.Context) error {
		err := cli.ShowAppHelp(c)
		checkFatalError(err)
		return nil
	}
	
	app.Commands = []cli.Command{
		{
			Name:  "start",
			Usage: "Start UDP Server",
			Flags: appFlags,
			Action: func(c *cli.Context) error {
				config := Config{}
				config.FillFromContext(c)
				
				// fmt.Println(config)
				// os.Exit(1)
				
				setupGrayLogHook(&config)
				
				udpAddr, errResolve := net.ResolveUDPAddr("udp", config.Listen)
				checkFatalError(errResolve)
				
				pc, errListen := net.ListenUDP("udp", udpAddr)
				checkFatalError(errListen)
				
				defer pc.Close()
				
				log.Println("Started Listening on:", config.Listen)
				log.Println("Quiet Mode :", config.Quiet)
				log.Println("No Data !? sudo sysctl -w net.inet.udp.maxdgram=65535")
				
				for {
					buf := make([]byte, config.BufferSize)
					n, addr, err := pc.ReadFrom(buf)
					if err != nil {
						fmt.Println(err)
						continue
					}
					go processUDP(pc, addr, buf[:n], config.Quiet)
				}
				
				return nil
			},
		},
	}
	
	err := app.Run(os.Args)
	checkFatalError(err)
}

func processUDP(pc net.PacketConn, addr net.Addr, buf []byte, quiet bool) {
	
	if !quiet {
		log.Println("Data Received")
	}
	
	pbp := &protos.Payload{}
	err := proto.Unmarshal(buf, pbp)
	
	checkNonFatalError(err)
	
	// IF passed create NDJson string
	ndJsonPayload := JsonPayload{}
	
	metadataPB := pbp.Metadata
	spansPB := pbp.Spans
	transactionsPB := pbp.Transactions
	errorsPB := pbp.Errors
	
	// Emit Default Value in JSON file
	m := jsonpb.Marshaler{
		EmitDefaults: true,
	}
	
	// Append MetaData
	metadata, metadataError := m.MarshalToString(metadataPB)
	checkNonFatalError(metadataError)
	ndJsonPayload.AppendData("metadata", metadata)
	
	// Append Spans
	for n := 0; n < len(spansPB); n++ {
		spanData, spanErr := m.MarshalToString(spansPB[n])
		checkNonFatalError(spanErr)
		ndJsonPayload.AppendData("span", spanData)
	}
	
	// Append Transactions
	for n := 0; n < len(transactionsPB); n++ {
		txData, txErr := m.MarshalToString(transactionsPB[n])
		checkNonFatalError(txErr)
		ndJsonPayload.AppendData("transaction", txData)
	}
	
	// Append Errors
	for n := 0; n < len(errorsPB); n++ {
		errData, txErr := m.MarshalToString(errorsPB[n])
		checkNonFatalError(txErr)
		ndJsonPayload.AppendData("error", errData)
	}
	
	ndjsonBytes := ndJsonPayload.Buffer.Bytes()
	tr, err := transport.InitDefault()
	checkNonFatalError(err)
	var requestBuf bytes.Buffer
	zlibWriter, _ := zlib.NewWriterLevel(&requestBuf, zlib.BestSpeed)
	_,err = zlibWriter.Write(ndjsonBytes)
	checkNonFatalError(err)
	err = zlibWriter.Close()
	checkNonFatalError(err)
	err =tr.SendStream(context.Background(), &requestBuf)
	checkNonFatalError(err)
	pc.WriteTo(buf, addr)
}

/** Setup GELF Hook for Logrus Logger **/
func setupGrayLogHook(c *Config) {
	graylogAddr := c.Gelf
	log.Infof("Setting-up GELF Logger on %s", graylogAddr)
	hook := graylog.NewGraylogHook(graylogAddr, map[string]interface{}{"facility": "Groot"})
	// Set GELF Logger for Levels above warning
	hook.Level = log.WarnLevel
	// Add Hook Lto Logrus
	log.AddHook(hook)
	// Main Level is Debug
	log.SetLevel(log.DebugLevel)
}

/** Check Error and Throw an Fatal Error **/
func checkFatalError(err error) {
	if err != nil {
		log.Fatal(fmt.Sprintf("%+v\n", err))
		os.Exit(-1)
	}
}

/** Check Error and Send Error Log to STDOUT and Graylog Server **/
func checkNonFatalError(err error) {
	if err != nil {
		log.Error(fmt.Sprintf("%+v\n", err))
	}
}
