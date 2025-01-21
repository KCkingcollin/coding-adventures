package main

import (
	"fmt"
)

func main() {
    var n int
    var k int
    var result int
    fmt.Scan(&n)
    fmt.Scan(&k)
    if k <= (n+2-1)/2 {
        // ods
        result = k*2-1
    } else {
        // evens
        result = (k-(n+2-1)/2)*2
    }
    fmt.Println(result)
}
