package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    var advancing int
    scanner.Scan()
    line := strings.Split(scanner.Text(), " ")
    n, _ := strconv.Atoi(line[0])
    k, _ := strconv.Atoi(line[1])
    scanner.Scan()
    line = strings.Split(scanner.Text(), " ")
    kthScore, _ := strconv.Atoi(line[k-1])
    for i := 0; i < n; i++ {
        score, _ := strconv.Atoi(line[i])
        if score >= kthScore && score > 0 {
            advancing++
        }
    }
    fmt.Println(advancing)
}
