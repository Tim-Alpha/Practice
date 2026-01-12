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

	p := make([]int, n+1)
	ans := make([]int, n+1)

    for i:= 1; i<=n;i++ {
        fmt.Fscan(in, &p[i])
        ans[p[i]] = i
    }

	for i := 1; i <= n; i++ {
		fmt.Print(ans[i], " ")
	}
}