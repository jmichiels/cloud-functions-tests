package cloud_functions_tests

import (
	"context"
	firebase "firebase.google.com/go"
	"github.com/jmichiels/cloud-functions-tests/internal"
	"log"
	"net/http"
)

var httpHandler http.Handler

// Runs once when the instance starts.
func init() {
	conf := &firebase.Config{
		ProjectID: "cloud-functions-tests-45711",
	}
	// Initialize the Firebase client.
	app, err := firebase.NewApp(context.Background(), conf)
	if err != nil {
		log.Fatalln(err)
	}
	// Initialize the HTTP handler.
	httpHandler, err = internal.NewHttpHandler(app)
	if err != nil {
		log.Fatalln(err)
	}
}

// Cloud function that acts as a single entry point to handle all HTTP requests.
func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	httpHandler.ServeHTTP(w, r)
}
