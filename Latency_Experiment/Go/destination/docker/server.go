package main

import (
  "time"
  "fmt"
  "net/http"
  "io/ioutil"
  "log"

  simplejson "github.com/bitly/go-simplejson"
)

func http_handler(w http.ResponseWriter, r *http.Request) {
    dest_ts := time.Now().UnixNano()
    log.Printf("Dest timestamp: %d\n", dest_ts)
    if r.Body != nil {
        data, _ := ioutil.ReadAll(r.Body)
        js, _ := simplejson.NewJson(data)
        base_ts := js.Get("Timestamp").MustInt64()
        log.Printf("Base timestamp: %d\n", base_ts)
        time_diff := dest_ts - base_ts
        log.Printf("Time Diff: %d\n", time_diff)
        fmt.Fprint(w, time_diff)
    }
}

func main(){
    http.HandleFunc("/", http_handler)
    http.ListenAndServe(":8080", nil)
}
