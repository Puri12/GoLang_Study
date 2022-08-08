package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var n, a, b int
	fmt.Fscanln(reader, &n)
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &a, &b)
		fmt.Fprintln(writer, a+b)
	}
	defer writer.Flush()
}
