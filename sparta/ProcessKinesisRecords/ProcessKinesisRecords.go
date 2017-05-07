package main

import (
  "encoding/json"
  "fmt"
  "net/http"
  "encoding/base64"
  logrus "github.com/Sirupsen/logrus"
  sparta "github.com/mweagle/Sparta"
  simplejson "github.com/bitly/go-simplejson"
)


func processDynamoDBStream(event *json.RawMessage,
                    context *sparta.LambdaContext,
                    w http.ResponseWriter,
                    logger *logrus.Logger) {

    data, _ := event.MarshalJSON()
    js, _ := simplejson.NewJson(data)
    length := len(js.Get("Records").MustMap())
    for i := 0; i < length; i++{
        decoded := js.Get("Records").GetIndex(i).Get("kinesis").Get("data").MustString()
        payload, _ := base64.StdEncoding.DecodeString(decoded)
        fmt.Println("Decoded payload : %q", payload)
}


}

func main() {
  var lambdaFunctions []*sparta.LambdaAWSInfo

  logStreamFn := sparta.NewLambda(sparta.IAMRoleDefinition{},
    processDynamoDBStream,
    nil)
  lambdaFunctions = append(lambdaFunctions, logStreamFn)
  sparta.Main("Process-DynamoDB-Stream-Go",
    "Go Lambda application to interact with Dynamo DB",
    lambdaFunctions,
    nil,
    nil)
}
