package main

import (
	"fmt"
	"strconv"
)

func fizzbuzz(n int) []string {
	outputs := []string{}

	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			outputs = append(outputs, "FizzBuzz")
		} else if i%3 == 0 {
			outputs = append(outputs, "Fizz")
		} else if i%5 == 0 {
			outputs = append(outputs, "Buzz")
		} else {
			outputs = append(outputs, strconv.Itoa(i))
		}
	}

	return outputs
}

func main() {
	n := 15
	fmt.Println(fizzbuzz(n))
}
