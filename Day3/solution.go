package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)

    for i := 0; i < n; i++ {
        j := "codeforces"

        var str string
        fmt.Scan(&str)

        cnt := 0
        for a := 0; a < len(j); a++ {
            if j[a] != str[a] {
                cnt++
            }
        }

        fmt.Println(cnt)
    }
}
