package main

import (
  "encoding/json"
  "time"
  "fmt"
  "net/http"
  "io/ioutil"
  "log"

  "github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
  simplejson "github.com/bitly/go-simplejson"
)

type TimeStamp struct {
        Source  string
        Timestamp int64
}

func http_handler(w http.ResponseWriter, r *http.Request){

    dest_ts := time.Now().UnixNano()
    data, _ := ioutil.ReadAll(r.Body)
    js, _ := simplejson.NewJson(data)
    base_ts := js.Get("Base").MustInt64()
    log.Println("Base timestamp: %d\n", base_ts)
    time_diff := base_ts - dest_ts
    log.Println("Time Diff: %d\n", time_diff)
    fmt.Fprintf(w, string(time_diff))

}

func Handle(evt json.RawMessage, ctx *runtime.Context) (interface{}, error) {
    http.HandleFunc("/", http_handler)
    http.ListenAndServe(":8080", nil)

    return nil, nil
}
