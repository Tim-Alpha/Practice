package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var t int
	fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {
		var a, b, c int
		fmt.Fscan(in, &a, &b, &c)
		if (a+b == c) || (a+c == b) || (b+c == a) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}