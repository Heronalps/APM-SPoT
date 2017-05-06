package main

import (
  "encoding/json"
  "fmt"
  "net/http"
  "os"
  "strconv"
  logrus "github.com/Sirupsen/logrus"
  sparta "github.com/mweagle/Sparta"
)

func init() {
  // Log as JSON instead of the default ASCII formatter.
  logrus.SetFormatter(&logrus.JSONFormatter{})

  // Output to stdout instead of the default stderr
  // Can be any io.Writer, see below for File example
  logrus.SetOutput(os.Stdout)

  // Only log the warning severity or above.
  logrus.SetLevel(logrus.InfoLevel)
}

func dosomething(){
    var m map[string]int = make(map[string]int)
    for i := 0; ; i++{
        s := strconv.Itoa(i)
        m["name" + s] = i
        fmt.Println(m)
    }
}

func getMyLogStream(event *json.RawMessage,
                    context *sparta.LambdaContext,
                    w http.ResponseWriter,
                    logger *logrus.Logger) {

    // Insert ill-behaved instructions to test AWS Lambda limits
    // Infinite loop
    /*
    for {
        fmt.Println(event)
    }
    */
    //Memory Leak
    /*
    for {
        go dosomething()
    }
    */

    //Exit with error






    fmt.Fprintf(w, "%s  ", "This Go function is to get AWS Lambda function log stream! ")
    fmt.Fprintf(w, "%s : %s  ", "AWSRequestID", context.AWSRequestID)
    fmt.Fprintf(w, "%s : %s  ", "InvokeID", context.InvokeID)
    fmt.Fprintf(w, "%s : %s  ", "LogGroupName", context.LogGroupName)
    fmt.Fprintf(w, "%s : %s  ", "FunctionName", context.FunctionName)
    fmt.Fprintf(w, "%s : %s  ", "MemoryLimitInMB", context.MemoryLimitInMB)
    fmt.Fprintf(w, "%s : %s  ", "FunctionVersion", context.FunctionVersion)
    fmt.Fprintf(w, "%s : %s  ", "InvokedFunctionARN", context.InvokedFunctionARN)

    logrus.WithFields(logrus.Fields{
      "AWSRequestID" : context.AWSRequestID,
      "InvokeID" : context.InvokeID,
      "LogGroupName" : context.LogGroupName,
      "FunctionName" : context.FunctionName,
      "MemoryLimitInMB" : context.MemoryLimitInMB,
      "FunctionVersion" : context.FunctionVersion,
      "InvokedFunctionARN" : context.InvokedFunctionARN,
    }).Info("All runtime information from AWS Lambda context.")
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
