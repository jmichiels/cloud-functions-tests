package main

// This is a little script to test the cloud function to check if we receive the correct response.

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/jmichiels/cloud-functions-tests/internal"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const cloudFunctionTriggerURL = "https://us-central1-cloud-functions-tests-45711.cloudfunctions.net/ServeHTTP"

func main() {
	if len(os.Args) < 2 {
		log.Fatal(fmt.Sprintf("usage: %s <name>", os.Args[0]))
	}
	encodedRequest, err := json.Marshal(internal.HelloRequest{
		Name: os.Args[1],
	})
	response, err := http.Post(cloudFunctionTriggerURL+"/api/v1/hello", "application/json", bytes.NewBuffer(encodedRequest))
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var decodedResponse internal.HelloResponse
	err = json.Unmarshal(body, &decodedResponse)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("response message: " + decodedResponse.Message)
}
