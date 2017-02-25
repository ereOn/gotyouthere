package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/go-kit/kit/log"

	"github.com/ereOn/gotyouthere/pkg/server"
)

func main() {
	endpoint := flag.String("endpoint", ":9999", "The endpoint to listen on.")

	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))
	logger = log.NewContext(logger).With("timestamp", log.DefaultTimestampUTC)

	server := server.NewHTTPHandler(logger)

	// Start the application and block.
	logger.Log("event", "application started", "endpoint", *endpoint)
	err := http.ListenAndServe(*endpoint, server)

	if err != nil {
		panic(err)
	}
}
