package main

import (
	"bytes"
	"fmt"
	"regexp"
)

type JsonPayload struct {
	Buffer bytes.Buffer
}

func (j *JsonPayload) AppendData(dataType string, payload string) {
	if dataType == "span" || dataType == "transaction" || dataType=="error" {
		// Clean timestamp type
		re := regexp.MustCompile(`(?m)("timestamp":"(\d{16})")`)
		payload = re.ReplaceAllString(payload, `"timestamp":$2`)
	}
	j.Buffer.WriteString(fmt.Sprintf("{\"%s\" : %s}\n", dataType, payload))
}
