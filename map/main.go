package main

import "fmt"

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Color is:", color, "Value is:", hex)
	}
}

func main() {
	// var colors map[string]string

	// colors := make(map[int]string)

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#556t6",
	}
	fmt.Println(colors)

	printMap(colors)
}
