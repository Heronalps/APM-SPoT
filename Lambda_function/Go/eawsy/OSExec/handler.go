package main

import (
    "os/exec"
    "log"

    "encoding/json"
    "github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
)

func Handle(evt json.RawMessage, ctx *runtime.Context) (interface{}, error) {
    cmd := exec.Command("ntpstat")
    stdout, err := cmd.Output()
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Output: %q\n", string(stdout))
    return nil, nil
}
