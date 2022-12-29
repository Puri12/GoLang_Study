package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, x int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	fmt.Fscanln(reader, &n, &x)
	defer writer.Flush()

	var num = make([]int, n)
	for i := range num {
		fmt.Fscanf(reader, "%d ", &num[i])
		if num[i] < x {
			fmt.Fprintf(writer, "%d ", num[i])
		}
	}
	fmt.Fprint(writer, "\n")
}
