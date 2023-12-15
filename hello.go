package main

import (
	"fmt"
	"os"
	"time"
)

var javascript = "bad"
var integer int = 24

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func main() {
	Goo := "good"
	floating := float64(integer)
	fmt.Println("Hello, World!", time.Now())
	newString1, newString2 := swap("string 1", "string2")
	val1, mandatoryVal2 := split(10)
	fmt.Println(newString1, newString2)
	fmt.Println(val1, mandatoryVal2)
	fmt.Println(javascript)
	fmt.Println(Goo)
	fmt.Println(floating)
	setEnv()
	fmt.Print(os.Getenv("FOO"))
	fmt.Print(os.Getenv("maxint"))
}
