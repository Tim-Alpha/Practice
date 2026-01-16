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
		var s string
		fmt.Fscan(in, &s)
		if s[0]+s[1]+s[2] == s[3]+s[4]+s[5] {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}