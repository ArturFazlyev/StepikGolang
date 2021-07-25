package main

import "fmt"

func main() {

	var (
		max   = 0
		num   int
		count = 0
	)

	for true {
		fmt.Scan(&num)
		if num > max {
			max = num
			count = 1
		} else if num == max {
			count += 1
		} else if num == 0 {
			break
		}
	}

	fmt.Println(count)

}
