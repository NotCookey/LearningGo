package main

import (
	"fmt"
	"time"
	"net/http"
)

type Person struct {
	age     int
	sdfsdff string
}

func calculateAge(age chan int, year int) {
	currentAge := time.Now().Year() - year
	age <- currentAge
}

func main() {
	var year int
	var age = make(chan int)
	var person *Person

	fmt.Print("Enter your birth year: ")
	fmt.Scan(&year)

	go calculateAge(age, year)

	person = &Person{age: <-age}

	defer fmt.Println("Thank you for using my product!")

	fmt.Printf("Your current age is, %d\n", person.age)
}
