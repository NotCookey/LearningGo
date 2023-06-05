package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
)

func main() {
    scanner: = bufio.NewReader(os.Stdin)
    fmt.Print("Enter a string : ")
    input,
    _: = scanner.ReadString('\n')
    input = strings.TrimSpace(input)

    if !(len(input) > 0) {
        panic("No input provided!")
    }

    fmt.Print(reverseCheck(input))

}

func reverseCheck(str string)(result string) {
    var check string = "HelloWorld"
    if str[0] >= 'A' && str[0] <= 'Z' {
        if str == check {
            return "!!!"
        } else if str == check[0: len(str)] {
            return check[len(str): len(check)]
        }
    }
    return "???"
}