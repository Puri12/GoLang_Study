package main

import "fmt"

func main() {
	var a, b, c int

	fmt.Scanln(&a, &b, &c)

	if a == b && b == c {
		fmt.Println(10000 + a*1000)
	} else if a == b && b != c {
		fmt.Println(1000 + a*100)
	} else if a == c && b != c {
		fmt.Println(1000 + a*100)
	} else if b == c && a != c {
		fmt.Println(1000 + b*100)
	} else {
		var max int
		if a > b && a > c {
			max = a
		} else if b > a && b > c {
			max = b
		} else {
			max = c
		}
		fmt.Println(max * 100)
	}
}
