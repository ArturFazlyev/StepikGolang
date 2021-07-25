package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int
	fmt.Scan(&i)

	a := strconv.Itoa(i)
	if a[0] == a[1] || a[0] == a[2] || a[1] == a[2] {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}

}
