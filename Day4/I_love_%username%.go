package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	count := 0
	min := 0
	max := 0

	for i := 0; i < n; i++ {
		var score int
		fmt.Scan(&score)
		if i == 0 {
			min = score
			max = score
		} else {
			if score > max {
				max = score
				count++
			} else if score < min {
				min = score
				count++
			}
		}
	}
	fmt.Println(count)
}