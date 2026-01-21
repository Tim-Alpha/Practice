package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var s1, s2, s3 string
	fmt.Fscan(in, &s1)
	fmt.Fscan(in, &s2)
	fmt.Fscan(in, &s3)
	
	charCount := make(map[rune]int)
	for _, ch := range s1 + s2 + s3 {
		charCount[ch]++
	}
	
	for _, count := range charCount {
		if count%2 != 0 {
			fmt.Println("NO")
			return
		}
	}
	fmt.Println("YES")
}
