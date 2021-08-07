package main

import "fmt"

func vote(x int, y int, z int) int {
	var zero, one int
	fmt.Scan(&x, &y, &z)
	arr := [3]int{x, y, z}

	zero = 0
	one = 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			zero += 1
		} else if arr[i] == 1 {
			one += 1
		}

	}

	if zero > one {
		return 0
	} else {
		return 1
	}
}
