package main

import "fmt"

func main() {
	for {
		var a, b int
		fmt.Scanln(&a, &b)
		if a == 0 && b == 0 {
			break
		}
		fmt.Println(a + b)
	}
}
