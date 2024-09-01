package main

// Uncomment the line below after implementing your reduce() function
import "fmt"

// TODO: create a reduce() function here
func reduce(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func main() {
	// Uncomment the line below after implementing your reduce() function
	fmt.Println(reduce([]int{0, 1, 1, 2, 3, 5, 8, 13, 21}))
}
