package discord

import (
	"bytes"
	"crypto/ed25519"
	"encoding/hex"
	"log"

	"github.com/aws/aws-lambda-go/events"
)

func Validate(publicKey string, req events.APIGatewayV2HTTPRequest) (*events.APIGatewayV2HTTPResponse, error) {

	log.Printf("Public Key in validate %s", publicKey)

	var resp events.APIGatewayV2HTTPResponse
	typedKey, err := hex.DecodeString(publicKey)
	if err != nil {
		resp.StatusCode = 401
		resp.Body = "Could not decode public key"
		return &resp, err
	}

	log.Printf("Public Key after decoding %v", typedKey)

	signature := req.Headers["x-signature-ed25519"]
	sig, err := hex.DecodeString(signature)
	if err != nil || len(sig) != ed25519.SignatureSize {
		resp.StatusCode = 401
		resp.Body = "Failed manual len check"
		return &resp, err
	}

	timestamp := req.Headers["x-signature-timestamp"]
	if timestamp == "" {
		resp.StatusCode = 401
		resp.Body = "Failed on find timestamp"
		return &resp, nil
	}

	var msg bytes.Buffer
	msg.WriteString(timestamp)
	msg.WriteString(req.Body)
	if !ed25519.Verify(typedKey, msg.Bytes(), sig) {
		resp.StatusCode = 401
		resp.Headers = req.Headers
		return &resp, nil
	}

	return nil, nil
}
