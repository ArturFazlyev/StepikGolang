package main

import "fmt"

func main() {
	var (
		name string
		age  int
	)

	fmt.Println("Введите имя: ")
	fmt.Scan(&name)
	fmt.Println("Введите возраст: ")
	fmt.Scan(&age)

	fmt.Println(name, age)

}
