package main

import "fmt"

func main() {
	var t, m, a int

	fmt.Scanln(&t, &m)
	fmt.Scanln(&a)

	m += a
	t += m / 60
	m %= 60
	t %= 24

	fmt.Println(t, m)
}
