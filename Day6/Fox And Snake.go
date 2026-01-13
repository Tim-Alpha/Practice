package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, m int
	fmt.Fscan(in, &n, &m)

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			for j := 0; j < m; j++ {
				fmt.Print("#")
			}
		} else {
			if (i/2)%2 == 0 {
				for j := 0; j < m-1; j++ {
					fmt.Print(".")
				}
				fmt.Print("#")
			} else {
				fmt.Print("#")
				for j := 0; j < m-1; j++ {
					fmt.Print(".")
				}
			}
		}
		fmt.Println()
	}
}