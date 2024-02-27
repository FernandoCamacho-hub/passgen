package main

import (
	"flag"
	"fmt"
)

var specialCharacters string = "~`!@#$%^&*()-_+={}[]|\\;:\"<>,./?"

func main() {
    fmt.Printf("hello, world\n")
    var numDigits = flag.Int("d", 1, "number of digits the password contains") 
    var numSymbols = flag.Int("s", 1, "number of special characters the password contains") 
    var upperCase = flag.Bool("u", true, "have uppercase")
    var lowerCase = flag.Bool("l", true, "have lowercase")
    flag.Parse()
    fmt.Printf("numdigits: %v\n", *numDigits)
    fmt.Printf("numSymbols: %v\n", *numSymbols)
    fmt.Printf("uppercase: %v\n", *upperCase)
    fmt.Printf("lowercase: %v\n", *lowerCase)
}
