// 열거형 2
package main

import "fmt"

func main() {
	const (
		// iota : 0부터 +1 씩 증가
		A = iota
		B
		C
	)

	fmt.Println(A, B, C)

	const (
		Jan = iota + 1
		Feb
		Mar
		Apr
		May
		Jun
	)

	fmt.Println(Jan)
	fmt.Println(Feb)
	fmt.Println(Mar)
	fmt.Println(Apr)
	fmt.Println(May)
	fmt.Println(Jun)
}
