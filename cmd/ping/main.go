package main

// This is a little script to test the cloud function to check if we receive the correct response.

import (
	"encoding/json"
	"github.com/jmichiels/cloud-functions-tests/internal"
	"io/ioutil"
	"log"
	"net/http"
)

const cloudFunctionTriggerURL = "https://us-central1-cloud-functions-tests-45711.cloudfunctions.net/ServeHTTP"

func main() {
	response, err := http.Get(cloudFunctionTriggerURL + "/api/v1/ping")
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var decodedResponse internal.PingResponse
	err = json.Unmarshal(body, &decodedResponse)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("response message: " + decodedResponse.Message)
}
