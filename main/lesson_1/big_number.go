package main

import "fmt"

func main() {
	var i int
	fmt.Scan(&i)
	if i > 10000 {
		fmt.Println("Число больше 10000, введите новое число")
		fmt.Scan(&i)
	} else if i < 0 {
		fmt.Println("Число меньше 0, введите новое число")
		fmt.Scan(&i)
	} else {
		fmt.Println(i % 10)
	}

}
