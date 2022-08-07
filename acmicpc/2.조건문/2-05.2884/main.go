package main

import "fmt"

func main() {
	var a, b int

	fmt.Scanln(&a, &b)

	b -= 45

	if b < 0 {
		a--
		b += 60
	}
	if a < 0 {
		a = 23
	}

	fmt.Println(a, b)
}
