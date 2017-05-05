package main

import (
  "encoding/json"
  "fmt"
  "net/http"
  "github.com/Sirupsen/logrus"
  sparta "github.com/mweagle/Sparta"
)

func getMyLogStream(event *json.RawMessage,
                    context *sparta.LambdaContext,
                    w http.ResponseWriter,
                    logger *logrus.Logger) {

    fmt.Fprintf(w, "%s  ", "This Go function is to get AWS Lambda function log stream! ")
    fmt.Fprintf(w, "%s : %s  ", "AWSRequestID", context.AWSRequestID)
    fmt.Fprintf(w, "%s : %s  ", "InvokeID", context.InvokeID)
    fmt.Fprintf(w, "%s : %s  ", "LogGroupName", context.LogGroupName)
    fmt.Fprintf(w, "%s : %s  ", "FunctionName", context.FunctionName)
    fmt.Fprintf(w, "%s : %s  ", "MemoryLimitInMB", context.MemoryLimitInMB)
    fmt.Fprintf(w, "%s : %s  ", "FunctionVersion", context.FunctionVersion)
    fmt.Fprintf(w, "%s : %s  ", "InvokedFunctionARN", context.InvokedFunctionARN)
}

func main() {
  var lambdaFunctions []*sparta.LambdaAWSInfo

  logStreamFn := sparta.NewLambda(sparta.IAMRoleDefinition{},
    getMyLogStream,
    nil)
  lambdaFunctions = append(lambdaFunctions, logStreamFn)
  sparta.Main("getMyLogStreamStack",
    "Go application in Sparta to get log information",
    lambdaFunctions,
    nil,
    nil)
}
