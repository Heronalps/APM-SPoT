package main

import (
  "encoding/json"
  "fmt"
  "net/http"
  logrus "github.com/Sirupsen/logrus"
  sparta "github.com/mweagle/Sparta"
  simplejson "github.com/bitly/go-simplejson"
)


func lambda_handler(event *json.RawMessage,
                    context *sparta.LambdaContext,
                    w http.ResponseWriter,
                    logger *logrus.Logger) {
    data, _ := event.MarshalJSON()
    js, _ := simplejson.NewJson(data)
    msg := js.Get("Records").GetIndex(0).Get("Sns").Get("Message").MustString()
    fmt.Println("From SNS : " + msg)
}

func main() {
  var lambdaFunctions []*sparta.LambdaAWSInfo

  logStreamFn := sparta.NewLambda(sparta.IAMRoleDefinition{},
    lambda_handler,
    nil)
  lambdaFunctions = append(lambdaFunctions, logStreamFn)
  sparta.Main("SNS-X-Account-Go",
    "Go Lambda application to interact with AWS SNS",
    lambdaFunctions,
    nil,
    nil)
}
