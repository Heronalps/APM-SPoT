package main

import (
    "fmt"
)

func main(){
    a := 11.1
    fmt.Printf("Tyep: %T, Value: %v \n",fmt.Sprint(a), fmt.Sprint(a))
    str := fmt.Sprint(a)
    fmt.Println(str)
}
