package main

import "fmt"

func main() {
	var a, b, n int

	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		fmt.Scanln(&a, &b)

		fmt.Printf("Case #%d: %d + %d = %d\n", i+1, a, b, a+b)
	}
}
