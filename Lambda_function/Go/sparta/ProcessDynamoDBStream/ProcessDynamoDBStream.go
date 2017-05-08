package main

import (
  "encoding/json"
  "fmt"
  "net/http"
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
    for i := 0; i < 3; i++{
        eventID := js.Get("Records").GetIndex(i).Get("eventName").MustString()
        eventName := js.Get("Records").GetIndex(i).Get("eventID").MustString()
        fmt.Println("Event ID : " + eventID)
        fmt.Println("Event Name : " + eventName)
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
