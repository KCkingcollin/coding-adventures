package main

import "fmt"

var (
    hello = []rune{'h', 'e', 'l', 'l', 'o'}
    word string
    pos int
)

func main() {
    fmt.Scan(&word)
    for _, char := range []rune(word) {
        if char == hello[pos] {
            pos++
            if pos > 4 {
                fmt.Println("YES")
                return
            }
        }
    }
    fmt.Println("NO")
}
