package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var t int
	fmt.Fscan(in, &t)
	for ; t > 0; t-- {
		var n int
		fmt.Fscan(in, &n)
		var arr = make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &arr[i])
		}

		sort.Ints(arr)

		for i := 0; i < len(arr)-1; {
			if arr[i] == arr[i+1] {
				arr = append(arr[:i], arr[i+1:]...)
			} else if arr[i]+1 == arr[i+1] {
				arr = append(arr[:i], arr[i+1:]...)
			} else {
				i++
			}
		}

		if len(arr) == 1 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
