package main

import "fmt"

func main() {
	var a [6]int

	fmt.Scanln(&a[0], &a[1], &a[2], &a[3], &a[4], &a[5])
	fmt.Printf("%d %d %d %d %d %d", 1-a[0], 1-a[1], 2-a[2], 2-a[3], 2-a[4], 8-a[5])
}
