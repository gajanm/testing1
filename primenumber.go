package main

import "fmt"

func factors(inputNumber int) {
	divisor := int(inputNumber)
	for i := 1; i <= divisor; i++ {
		if inputNumber%i == 0 {
			fmt.Println(i, "is a factor of ", inputNumber)
		}
	}
}

func main() {
	factors(50)
}
