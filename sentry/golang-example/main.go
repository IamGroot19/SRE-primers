package main

import (
	"log"
	"time"

	"github.com/getsentry/sentry-go"
)

func main() {
	err := sentry.Init(sentry.ClientOptions{

		Dsn: "",  // NOTE: Add the relevant DSN
		SampleRate: 0.25,  // for eerror events

		// Set TracesSampleRate to 1.0 to capture 100%
		// of transactions for performance monitoring.
		// We recommend adjusting this value in production,
		TracesSampleRate: 1.0,    // for traces
		Debug:            true,
		Environment: "demo-environment",
		Release: "my-project-name@1.0.0",
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	// Flush buffered events before the program terminates.
	defer sentry.Flush(2 * time.Second)

	
	sentry.CaptureMessage("Firing event at 1 pm for golang project=2")
}
