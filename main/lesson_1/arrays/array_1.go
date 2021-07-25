package main

import "fmt"

func main() {
	var a uint8
	var b, c int
	var workArray [10]uint8
	for i := 0; i < len(workArray); i++ {
		fmt.Scan(&a)
		workArray[i] = a
	}

	for i := 0; i < 3; i++ {
		fmt.Scan(&b, &c)
		workArray[b], workArray[c] = workArray[c], workArray[b]
	}

	for i := 0; i < len(workArray); i++ {
		fmt.Print(workArray[i], ",")
	}

}
