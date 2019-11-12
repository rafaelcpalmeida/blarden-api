package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type Response struct {
	Status string `json:"status"`
}

func RequestOpenDoor(door string) (Response, error) {
	internalAPIUrl := os.Getenv("INTERNAL_API_URL")
	api2apiToken := os.Getenv("API2API_TOKEN")

	key := GetAESToken()

	encrypted, _ := Encrypt([]byte(fmt.Sprintf("{\"key\": \"%s\", \"timestamp\": %d}", api2apiToken, time.Now().Unix())), &key)

	requestBody, err := json.Marshal(map[string] string {
		"message": fmt.Sprintf("%x", encrypted),
	})

	if err != nil {
		return Response{}, err
	}

	resp, err := http.Post(fmt.Sprintf("%s/%s", internalAPIUrl, door), "application/json", bytes.NewBuffer(requestBody))

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
