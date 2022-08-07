package main

import "fmt"

func main() {
	var n int

	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		var a, b int
		fmt.Scanln(&a, &b)
		fmt.Println(a + b)
	}
}
