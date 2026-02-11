package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(in, &n)
	var a, b, c int
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a, &b, &c)
		if a == b {
			fmt.Println(c)
		} else if a == c {
			fmt.Println(b)
		} else {
			fmt.Println(a)
		}
	}
}