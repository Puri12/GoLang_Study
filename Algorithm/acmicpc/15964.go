package main

import "fmt"

func main() {
	a, b := 0, 0
	p := func(a int, b int) int {
		return (a + b) * (a - b)
	}

	fmt.Scanln(&a, &b)
	fmt.Println(p(a, b))

}
