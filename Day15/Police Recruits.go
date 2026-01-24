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
	var police, crime int
	for ; t > 0; t-- {
		var event int
		fmt.Fscan(in, &event)
		if event > 0 {
			police += event
		} else {
			if police > 0 {
				police--
			} else {
				crime++
			}
		}
	}
	fmt.Println(crime)
}