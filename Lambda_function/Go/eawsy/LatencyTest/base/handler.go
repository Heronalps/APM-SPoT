package main

import (
  "encoding/json"
  "log"
  "time"

  "github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
)

func Handle(evt json.RawMessage, ctx *runtime.Context) (interface{}, error) {
    log.Printf("Log stream name: %s\n", ctx.LogStreamName)
	log.Printf("Log group name: %s\n", ctx.LogGroupName)
	log.Printf("Request ID: %s\n", ctx.AWSRequestID)
	log.Printf("Mem. limits(MB): %d\n", ctx.MemoryLimitInMB)

	select {
	case <-time.After(1 * time.Second):
		log.Printf("Time remaining (MS): %d\n", ctx.RemainingTimeInMillis())
	}

	return nil, nil
}
