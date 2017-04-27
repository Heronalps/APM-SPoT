package main

import (
	"strings"
    "fmt"
)

func WordCount(s string) map[string]int {
	Map := make(map[string]int)
	split := strings.Fields(s)
	for _, value := range split {
		if _, ok := Map[value]; ok {
			Map[value] = Map[value] + 1
		} else {
			Map[value] = 1
		}
	}
	return Map
}

func main() {
	fmt.Println(WordCount("Let's Go!"))
}
