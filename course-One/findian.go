package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	lowerText := strings.ToLower(text)
	finalText := strings.TrimSpace(lowerText)

	if strings.HasPrefix(finalText, "i") && string(finalText[len(finalText) - 1]) == "n" &&
		strings.ContainsAny(finalText , "a"){

		fmt.Println("Found!")

	}else {
		fmt.Println("Not Found!")
	}
}
