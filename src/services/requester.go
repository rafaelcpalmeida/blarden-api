package services

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type Response struct {
	Status string `json:"status"`
}

func RequestOpenDoor() (Response, error) {
	internalAPIUrl := os.Getenv("INTERNAL_API_URL")
	api2apiToken := os.Getenv("API2API_TOKEN")
	aesToken := os.Getenv("AES_TOKEN")

	keyStr, _ := hex.DecodeString(aesToken)
	var key [32]byte
	copy(key[:], keyStr)

	encrypted, _ := Encrypt([]byte(fmt.Sprintf("{\"key\": \"%s\", \"timestamp\": %d}", api2apiToken, time.Now().Unix())), &key)

	requestBody, err := json.Marshal(map[string] string {
		"message": fmt.Sprintf("%x", encrypted),
	})

	if err != nil {
		return Response{}, err
	}

	resp, err := http.Post(fmt.Sprintf("%s/open-door", internalAPIUrl), "application/json", bytes.NewBuffer(requestBody))

	if err != nil {
		return Response{}, err
	}

	defer resp.Body.Close()

	response := &Response{}

	err = json.NewDecoder(resp.Body).Decode(&response)

	if err != nil {
		return Response{}, err
	}

	return *response, nil
}
