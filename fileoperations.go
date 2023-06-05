package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("test.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	if success, err := file.WriteString("Meow :D"); err == nil {
		fmt.Println(success)
	}

	data, _ := os.ReadFile("test.txt")
	fmt.Print(string(data))
}
