package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	str, _ := in.ReadString('\n')

	letters := make(map[rune]bool)
	for _, ch := range str {
		if ch >= 'a' && ch <= 'z' {
			letters[ch] = true
		}
	}

	fmt.Println(len(letters))
}