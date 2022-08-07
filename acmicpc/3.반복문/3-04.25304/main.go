package main

import "fmt"

func main() {
	var money, n int

	fmt.Scanln(&money)
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		var m, nn int
		fmt.Scanln(&m, &nn)
		money -= m * nn
	}
	if money == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
