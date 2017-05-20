package main

import (
  "encoding/json"
  "log"
  "time"
  "net/http"
  "bytes"
  "io/ioutil"
)

type TimeStamp struct {
        Source  string
        Timestamp int64
}

func main() {

    base_ts := time.Now().UnixNano()
    log.Printf("Base timestamp: %d\n", base_ts)
    var ts = TimeStamp{Source: "Base", Timestamp: base_ts}
    js, err := json.Marshal(ts)

    url := "http://ec2-52-43-80-154.us-west-2.compute.amazonaws.com:8080"
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(js))
    req.Header.Set("X-Custom-Header", "latency_experiment")
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    log.Println("response Status:", resp.Status)
    log.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    log.Println("response Body:", string(body))
}
