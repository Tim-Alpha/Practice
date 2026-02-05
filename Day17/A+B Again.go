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
		var num int
		fmt.Fscan(in, &num)

		sum := 0
		n := num

		for n > 0 {
			sum += n % 10
			n = n / 10
		}

		fmt.Println(sum)
	}
}
