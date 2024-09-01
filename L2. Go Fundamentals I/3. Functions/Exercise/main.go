package main

import (
	"fmt"
	"strconv"
)

func getRectangleArea(width, length int) string {
	area := width * length
	if area < 50 {
		return "The area is " + strconv.Itoa(area) + ", which is less than 50"
	}
	return "The area is " + strconv.Itoa(area) + ", which is greater than or equal to 50"
}

func main() {
	result := getRectangleArea(58, 30)
	fmt.Println(result)
}
