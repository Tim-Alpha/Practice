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

		if n >= 1900 {
			fmt.Println("Division 1")
		} else if n >= 1600 && n <= 1899 {
			fmt.Println("Division 2")
		} else if n >= 1400 && n <= 1599 {
			fmt.Println("Division 3")
		} else {
			fmt.Println("Division 4")
		}
	}
}