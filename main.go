package main

import (
	"net/http"
	"os"

	"log"
	
	logkit "github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	svc := stringService{}

	svcWithLogger := &loggingMiddleware{
		next: svc,
		logger: logkit.NewLogfmtLogger(os.Stderr),
	}

	uppercaseHandler := httptransport.NewServer(
		makeUppercaseEndpoint(svcWithLogger),
		decodeUppercaseRequest,
		encodeResponse,
	)

	countHandler := httptransport.NewServer(
		makeCountEndpoint(svcWithLogger),
		decodeCountRequest,
		encodeResponse,
	)

	http.Handle("/uppercase", uppercaseHandler)
	http.Handle("/count", countHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
