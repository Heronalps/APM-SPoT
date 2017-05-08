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
    msg := js.Get("Records").GetIndex(0).Get("Sns").Get("Message").MustString()
    log.Printf("From SNS : " + msg)
    return nil, nil
}
