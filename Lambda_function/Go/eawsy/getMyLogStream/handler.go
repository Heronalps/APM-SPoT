package main

import (
  "encoding/json"
  "log"
  "time"
  "strconv"

  "github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
)

func dosomething(){
    var m map[string]int = make(map[string]int)
    for i := 0; ; i++{
        s := strconv.Itoa(i)
        m["name" + s] = i
        log.Printf(string(m["name" + s]))
    }
}

func Handle(evt json.RawMessage, ctx *runtime.Context) (interface{}, error) {

    // Insert ill-behaved instructions to test AWS Lambda limits
    // Infinite loop
    /*
    for {
        log.Printf(string(evt))
    }
    */
    //Memory Leak
    /*
    for {
        go dosomething()
    }
    */

    //Exit with error
    /*
    panic("Stop the lambda function") //Panic exit
    os.Exit(1) // OS Exit
    fmt.Fprintf(os.Stderr, "error: %v \n ", MyError("Intentional Error")) // Fatal error : Stack overflow
    */


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
