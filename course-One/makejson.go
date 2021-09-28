package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var myMap map[string]string
	myMap = make(map[string]string)

	fmt.Println("name: ")

	buffer := bufio.NewReader(os.Stdin)
	secondBuffer := bufio.NewReader(os.Stdin)

	bufferName, _, _ := buffer.ReadLine()
	name := string(bufferName)
	myMap["name"]=name

	fmt.Println("address: ")
	bufferAddress, _, _ := secondBuffer.ReadLine()
	address := string(bufferAddress)
	myMap["address"] = address

	data, _ := json.Marshal(myMap)

	fmt.Println("Decoded data: ")
	fmt.Println(string(data))


}
