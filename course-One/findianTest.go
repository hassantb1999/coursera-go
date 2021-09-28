package main

import (
	"fmt"
	"strings"
)

func main() {
	var name string
	
	fmt.Println("Type something: ")
	fmt.Scanln(&name)
	strings.ToLower(name)
	if strings.ContainsAny(name,"ian && i-a-n") == true {
		fmt.Println("Found")
	}else {
		fmt.Println("Not Found")
	}
}