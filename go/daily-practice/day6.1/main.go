package main

import "fmt"

func main() {
    var p string
    fmt.Scan(&p)
    for _, char := range []rune(p) {
        if char == 'H' || char == 'Q' || char == '9' {
            fmt.Println("YES")
            return
        }
    }
    fmt.Println("NO")
}
