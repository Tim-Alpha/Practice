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

	for ; t > 0; t-- {
		var n int
		fmt.Fscan(in, &n)

		a := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &a[i])
		}

		common := a[0]
		if a[0] != a[1] {
			if a[0] == a[2] {
				common = a[0]
			} else {
				common = a[1]
			}
		}

		for i := 0; i < n; i++ {
			if a[i] != common {
				fmt.Println(i + 1)
				break
			}
		}
	}
}
