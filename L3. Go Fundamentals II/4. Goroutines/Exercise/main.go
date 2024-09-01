package main

import (
	"fmt"
	"time"
)

func print(things []string) {
	for _, thing := range things {
		fmt.Println(thing)
	}
}

func main() {
	startTime := time.Now()

	colorNames := []string{"red", "orange", "yellow", "green", "blue", "indigo", "violet"}
	colorCodes := []string{"#FF0000", "#FF7F00", "#FFFF00", "#00FF00", "#0000FF", "#4B0082", "#9400D3"}
	go print(colorNames)
	go print(colorCodes)

	duration := time.Since(startTime)

	fmt.Println("Duration of execution: " + duration.String())

	time.Sleep(time.Second)

}
