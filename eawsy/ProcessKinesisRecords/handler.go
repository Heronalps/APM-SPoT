package main

import (
  "encoding/json"
  "log"

  "github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
  simplejson "github.com/bitly/go-simplejson"
)

func Handle(evt json.RawMessage, ctx *runtime.Context) (interface{}, error) {
    data, _ := evt.MarshalJSON()
    js, _ := simplejson.NewJson(data)
    length := len(js.Get("Records").MustMap())
    for i := 0; i < length; i++{
        decoded := js.Get("Records").GetIndex(i).Get("kinesis").Get("data").MustString()
        payload, _ := base64.StdEncoding.DecodeString(decoded)
        log.Printf("Decoded payload : %q", payload)
    }
    return nil, nil
}
