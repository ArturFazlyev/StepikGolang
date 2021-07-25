package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int
	fmt.Scan(&i)

	a := strconv.Itoa(i)
	fmt.Print(string(a[0]))
}
