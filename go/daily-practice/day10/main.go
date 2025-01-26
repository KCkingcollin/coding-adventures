package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var n int
var minX int
var maxX int
var y []int
var x []int

func p(i int) int {
    if i > maxX {
        y[len(y)-1] = p(maxX-1) + x[maxX]
        return y[len(y)-1]
    }
    if y[i] == 0 && i > minX {
        y[i] = p(i-1) + x[i]
    }
    return y[i]
}

func main() {
	reader := bufio.NewReader(os.Stdin)
    fmt.Fscan(reader, &n)
    x = make([]int, 0, 100001)
    for i := 0; i < n; i++ {
        var temp int
		fmt.Fscan(reader, &temp)
        if temp < minX {minX = temp}
        if temp > maxX {maxX = temp}
        x = x[:maxX+1]
        x[temp]++
    }
    y = make([]int, len(x)+1)

	var q int
    fmt.Fscan(reader, &q)
    var m int
    var answer []string = make([]string, q)
    for i := 0; i < q; i++ {
        fmt.Fscan(reader, &m)
        answer[i] = fmt.Sprintf("%d", p(m))
    }
    fmt.Println(strings.Join(answer, "\n"))
}
