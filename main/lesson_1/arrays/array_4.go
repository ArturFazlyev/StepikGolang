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

	vector_slice := vector[0:a]

	for i := 0; i < len(vector_slice); i++ {
		if i%2 == 0 {
			fmt.Print(vector_slice[i], " ")
		}
	}

}
