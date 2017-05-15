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
    for i := 0; i < 3; i++{
        eventID := js.Get("Records").GetIndex(i).Get("eventName").MustString()
        eventName := js.Get("Records").GetIndex(i).Get("eventID").MustString()
        log.Printf("Event ID : " + eventID)
        log.Printf("Event Name : " + eventName)
    }
    return nil, nil
}
