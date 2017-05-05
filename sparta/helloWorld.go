package main

import (
  "encoding/json"
  "fmt"
  "net/http"
  "github.com/Sirupsen/logrus"
  sparta "github.com/mweagle/Sparta"
)

func helloWorld(event *json.RawMessage, context *sparta.LambdaContext, w http.ResponseWriter, logger *logrus.Logger) {
  fmt.Fprintf(w, "Hello World!")
}
