package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

//Name - contains fname as first name and lname as last name
type Name struct {
	fname string
	lname string
}

func main() {
	path := ""
	fmt.Println("Enter Path:(sth like ~/home/projects/go/test.txt")
	_, err := fmt.Scanln(&path)
	if err != nil {
		fmt.Println("Cant scan!!")
		return
	}
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Something Wrong with reading file")
		return
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(file)

	var name []Name
	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()

		if err != nil || io.EOF == err {
			break
		}
		split := strings.Split(string(line), " ")
		typename := Name{
			split[0],
			split[1],
		}
		name = append(name, typename)
	}
	for i := 0; i < len(name); i++ {
		fmt.Println(i + 1)
		fmt.Println("First Name:" + name[i].fname)
		fmt.Println("Last Name:" + name[i].lname)
	}
}
