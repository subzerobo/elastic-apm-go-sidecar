# Welcome to Groot Application
# About
Hi, as you may know Elastic APM  does not provide any UDP interface for scripting languages like PHP to send data in non-blocking way.
this project act as a middleware (**sidecar**) which gets the data from `subzerobo/elastic-apm-php-agent` library which implemented elastic-apm intake api V2 with both ***TCP*** ( **NDJSON + gzip**) and ***UDP*** (**protobuf + gzip**) connectors. 

**How to Launch :** 

Just fork the repository, i've included the protobuf generated .pb.go files already but you can create them yourself using

    protoc --go_out=. ./protos/*.proto 
if you need the ***protoc*** cli code-generator you may consult this page to get your version  [https://github.com/protocolbuffers/protobuf/releases](https://github.com/protocolbuffers/protobuf/releases) 



# Config 

This application uses environment variables for configuration as it use the elastic-apm go agent as dependency to post data to your apm server so you may consult it`s documentation for environment variables settings [https://www.elastic.co/guide/en/apm/agent/go/current/configuration.html](https://www.elastic.co/guide/en/apm/agent/go/current/configuration.html)

### `ELASTIC_APM_SERVER_URL`
The URL for your Elastic APM server. The server supports both HTTP and HTTPS. If you use HTTPS, then you may need to configure your client machines so that the server certificate can be verified. You can also disable certificate verification with [`ELASTIC_APM_VERIFY_SERVER_CERT`](https://github.com/elastic/apm-agent-go/blob/1.x/docs/configuration.asciidoc#config-verify-server-cert).

### `ELASTIC_APM_SERVER_TIME`
| Environment | Default | Example |
|--|--|--|
| `ELASTIC_APM_SERVER_TIMEOUT` | `30s` | `30s` |
The timeout for requests made to your Elastic APM server. When set to zero or a negative value, timeouts will be disabled

### `ELASTIC_APM_SECRET_TOKEN`
| Environment | Default | Example |
|--|--|--|
| `ELASTIC_APM_SECRET_TOKEN` |  | `A Random String` |

This string is used to ensure that only your agents can send data to your APM server. Both the agents and the APM server have to be configured with the same secret token.

> the secret token is sent as plain-text in every request to the server, so you should also secure your communications using HTTPS. Unless you do so, your secret token could be observed by an attacker.

### `ELASTIC_APM_SERVER_CERT`
| Environment | Default |
|--|--|
| `ELASTIC_APM_SERVER_CERT` |  |
If you have configured your APM Server with a self signed TLS certificate, or you just wish to pin the server certificate, you can specify the path to the PEM-encoded certificate via the `ELASTIC_APM_SERVER_CERT` configuration.

### `ELASTIC_APM_VERIFY_SERVER_CERT`
| Environment | Default |
|--|--|
| `ELASTIC_APM_VERIFY_SERVER_CERT` | `true` |
By default, the agent verifies the server’s certificate if you use an HTTPS connection to the APM server. Verification can be disabled by changing this setting to `false`. This setting is ignored when `ELASTIC_APM_SERVER_CERT` is set.

## Command line options

#### Help 
Dispalys the complete help for the application
`$ go run main.go help `

#### Listen Port
Set the listen port of application : `-l | --listen`
`$ go run main.go -l :port `  
`$ go run main.go -l ":1113"`

#### Version
Shows the version of build 
`$ groot -v | --version`
