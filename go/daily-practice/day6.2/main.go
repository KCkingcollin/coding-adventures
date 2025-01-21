package main

import (
	"fmt"
)

func main() {
    var n, m, a int
    fmt.Scan(&n, &m, &a)
    result := ((n+a-1)/a)*((m+a-1)/a)
    fmt.Println(result)
}
