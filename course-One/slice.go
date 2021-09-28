package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var slice []int = make([]int, 3)
	var input string
	fmt.Println("Enter a number (X or x to exit):")
	for true {
		fmt.Scanln(&input)
		if input == "X" || input == "x" {
			break
		}
		aInput, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Wrong Input! make sure you enter an integer.")
			continue
		}
		slice = append(slice, aInput)
		sort.Ints(slice[:])
		/*Print from index 3 so skip show 3 zero at first evey time.*/
		fmt.Println(slice[3:])
		fmt.Println("Enter a number (X or x to exit):")
	}
}
