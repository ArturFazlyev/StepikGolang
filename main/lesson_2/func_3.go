package main

import "fmt"

func main() {
	fmt.Print(sumInt(1, 2, 3, 4))
}

func sumInt(a ...int) (int, int) {

	var num, sum int

	for elem := 0; elem < len(a); elem++ {
		num += 1
		sum += a[elem]
	}

	return num, sum

}
