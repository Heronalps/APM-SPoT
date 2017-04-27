package main

import (
    "fmt"
    "time"
)

func readword(ch chan string) {
    fmt.Println("Type a word, then hit Enter.")
    var word string
    fmt.Scanf("%s", &word)
    ch <- word
}

func timeout(t chan bool) {
    time.Sleep(5 * time.Second)
    t <- true
}

func main() {
    t := make(chan bool)
    ch := make(chan string)
    go timeout(t)

    for counter := 0; counter < 5;{
        go readword(ch)
        select {
            case word := <-ch:
                fmt.Println("Received", word)
            case <-t:
                fmt.Println("Timeout.")
                go timeout(t)
                counter++
        }
    }
}
