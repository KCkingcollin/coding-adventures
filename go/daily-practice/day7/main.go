package main

import "fmt"

var (
    n int 
    Vec3 [3]int
    num int
)

func main() {
    fmt.Scan(&n)
    for i := 0; i < n; i++ {
        for j := 0; j < 3; j++ {
            fmt.Scan(&num)
            Vec3[j] += num
        }
    }
    for _, Vec := range Vec3 {
        if Vec != 0 {
            fmt.Println("NO")
            return
        }
    }
    fmt.Println("YES")
}
