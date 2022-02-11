package discord

import (
	"bytes"
	"crypto/ed25519"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// func Validate(publicKey string, req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
func Validate(publicKey string, req *http.Request) bool {

	log.Println(json.Marshal(req))

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

	var msg bytes.Buffer
	msg.WriteString(timestamp)

	b, _ := io.ReadAll(req.Body)

	msg.WriteString(string(b))
	if !ed25519.Verify(typedKey, msg.Bytes(), sig) {
		log.Println("%w", "failed verify")
		return false
	}

	return true
}
