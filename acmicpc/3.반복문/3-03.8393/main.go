package main

import "fmt"

func main() {
	var n, a int

	fmt.Scanln(&n)

	for i := 1; i <= n; i++ {
		a += i
	}
	fmt.Println(a)
}
