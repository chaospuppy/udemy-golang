package main

import "fmt"

func main() {
	var is []int
	even := true
	for i := 0; i < 10; i++ {
		is = append(is, i)
	}
	for i := range is {
		if even {
			fmt.Printf("%v is even\n", i)
			even = false
		} else {
			fmt.Printf("%v is odd\n", i)
			even = true
		}
	}
}
