package cloud_functions_tests

import (
	"github.com/jmichiels/cloud-functions-tests/internal"
	"net/http"
)

var httpHandler http.Handler

func init() {
	// Initializes the HTTP handler once when the instance starts.
	httpHandler = internal.NewHttpHandler()
}

// Cloud function that acts as a single entry point to handle all HTTP requests.
func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	httpHandler.ServeHTTP(w, r)
}
