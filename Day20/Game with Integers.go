package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var t int
	fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Fscan(in, &n)

		if n % 3 == 0 {
			fmt.Println("Second")
		} else {
			fmt.Println("First")
		}
	}
}