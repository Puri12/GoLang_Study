package main

import "fmt"

func main() {
	N := 0
	fmt.Scanln(&N)

	for i := 0; i < N; i++ {
		var s string = ""
		fmt.Scanln(&s)
		fmt.Println(s[:1] + s[len(s)-1:])
	}
}
