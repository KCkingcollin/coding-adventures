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
    scanner.Scan()
    maxLines, _ := strconv.Atoi(scanner.Text())
    var X int
    for loopN := 0; scanner.Scan() && loopN < maxLines; loopN++ {
        line := strings.Split(scanner.Text(), "")
        for _, char := range line {
            if char == "+" {
                X++
                break
            } else if char == "-" {
                X--
                break
            }
        }
    }
    fmt.Println(X)
}
