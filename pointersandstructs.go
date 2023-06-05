package main
import "fmt"

func main() {

  // declare a struct Person
  type Person struct {
    name string
    age int
  }

  person := Person{"John", 25}

  // create a struct type pointer that
  // stores the address of person
  var ptr *Person
  ptr = &person

  // access the name member
  fmt.Println("Name:", ptr.name)

  // access the age member
  fmt.Println("Age:", ptr.age)

}