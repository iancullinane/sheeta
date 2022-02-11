package discord

import (
	"bytes"
	"crypto/ed25519"
	"encoding/hex"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// func Validate(publicKey string, req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
func Validate(publicKey string, req *http.Request) bool {

	typedKey, err := hex.DecodeString(publicKey)
	if err != nil {
		log.Println("%w", err)
		return false
	}

	signature := req.Header.Get("X-Signature-Ed25519")
	sig, err := hex.DecodeString(signature)
	if err != nil || len(sig) != ed25519.SignatureSize {
		log.Println("%w", "failed manual length check")
		return false
	}

	timestamp := req.Header.Get("X-Signature-Timestamp")
	if timestamp == "" {
		log.Println("%w", "failed timestamp")
		return false
	}

	//
	//

	var msg bytes.Buffer
	msg.WriteString(timestamp)

	defer req.Body.Close()
	var body bytes.Buffer

	// at the end of the function, copy the original body back into the request
	defer func() {
		req.Body = ioutil.NopCloser(&body)
	}()

	// copy body into buffers
	_, err = io.Copy(&msg, io.TeeReader(req.Body, &body))
	if err != nil {
		return false
	}

	//
	//

	// msg.WriteString(body)
	if !ed25519.Verify(typedKey, msg.Bytes(), sig) {
		log.Println("%w", "failed verify")
		return false
	}

	return true
}
