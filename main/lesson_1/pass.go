package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int
	fmt.Scan(&i)

	var sum_1 int
	var sum_2 int

	a := strconv.Itoa(i)
	for b := 0; b < len(a); b++ {
		if b <= 2 {
			sum_1 += int(a[b])
		} else {
			sum_2 += int(a[b])
		}
	}

	if sum_1 == sum_2 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")

	}

}
