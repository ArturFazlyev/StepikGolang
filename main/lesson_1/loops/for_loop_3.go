package main

import (
	"fmt"
)

func main() {
	var n, a, sum int
	fmt.Scan(&n)

	vector := []int{}

	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		vector = append(vector, a)
	}

	for i := 0; i < len(vector); i++ {
		if vector[i]%8 == 0 && vector[i] < 100 && vector[i] > 10 {
			sum += vector[i]
		}
	}

	fmt.Println(sum)

}
