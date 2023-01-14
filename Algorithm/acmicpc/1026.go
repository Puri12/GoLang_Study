package main

import (
	"fmt"
	"sort"
)

func main() {
	var N int
	fmt.Scan(&N)
	var a, b = make([]int, N), make([]int, N)

	for i := range a {
		fmt.Scan(&a[i])
	}
	for i := range b {
		fmt.Scan(&b[i])
	}

	sort.Sort(sort.IntSlice(a))
	sort.Sort(sort.Reverse(sort.IntSlice(b)))

	result := 0

	for i := range a {
		result += a[i] * b[i]
	}

	fmt.Println(result)
}
