package main

import "fmt"

func main() {
	var a, b int

	fmt.Scan(&a)
	vector := []int{}

	for i := 0; i <= a; i++ {
		fmt.Scan(&b)
		vector = append(vector, b)
	}

	fmt.Println(vector[3])

}
