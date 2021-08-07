package main

import "fmt"

func main() {
	fmt.Println(minimumFromFour())
}

func minimumFromFour() int {

	vector := []int{}
	var a, b, c, d int

	fmt.Scan(&a, &b, &c, &d)
	vector = append(vector, a, b, c, d)
	var min = vector[0]
	for i := 1; i < len(vector); i++ {
		if vector[i] < min {
			min = vector[i]
		}

	}

	return min
}
