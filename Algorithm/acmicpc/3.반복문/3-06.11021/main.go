package main

import "fmt"

func main() {
	var a, b, n int

	fmt.Scanln(&n)

	for i := 1; i <= n; i++ {
		fmt.Scanln(&a, &b)
		fmt.Printf("Case #%d: %d\n", i, a+b)
	}
}
