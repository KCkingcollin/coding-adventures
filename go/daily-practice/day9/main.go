package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
    var n int
    fmt.Fscan(reader, &n)
	var x []int = make([]int, n)
    for i := range x {
		fmt.Fscan(reader, &x[i])
    }
    slices.Sort(x)
    // fmt.Println(x)

	var q int
    fmt.Fscan(reader, &q)
    var m int
    var answer []string = make([]string, q)
    for i := 0; i < q; i++ {
        fmt.Fscan(reader, &m)
        // fmt.Printf("%d ", m)
        if m < x[n-1] {
            var L int
            var R int = n
            for L < R {
                a := (L+R)/2
                if x[a] > m {
                    R = a
                } else {
                    L = a + 1
                }
            }
            answer[i] = fmt.Sprintf("%d", R)
        } else {
            answer[i] = fmt.Sprintf("%d", n)
        } 
        // fmt.Println()
    }
    fmt.Println(strings.Join(answer, "\n"))
}
