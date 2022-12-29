// 열거형 3
package main

import "fmt"

func main() {
	const (
		_ = iota // _로 특정부분 생략
		A
		B
		C
	)

	const (
		_ = iota + 0.75*2
		DEFAULT
		SILVER
		GOLD
		PLATINUM
	)

	fmt.Println("DEFAULT :", DEFAULT)
	fmt.Println("SILVER :", SILVER)
	fmt.Println("GOLD :", GOLD)
	fmt.Println("PLATINUM :", PLATINUM)
}
