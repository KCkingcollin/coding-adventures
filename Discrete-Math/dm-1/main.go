package main

import (
	"fmt"
	"math"
	"unsafe"
)

func main() {
	// truth tables

	var p bool = true
    fmt.Println("Truth table for ~p")
    fmt.Println("p    |    !p")
	for i := 1; i <= int(math.Pow(2, float64(unsafe.Sizeof(p)))); i++ {
        fmt.Println("--------------")
        fmt.Printf("%t | %t\n", p, !p)
        p = !p
    }

    fmt.Println()
	p = true
    var q bool = true
    fmt.Println("Truth table for p q")
    fmt.Println("p    |    q")
	for i, j := 1, 1; i <= int(math.Pow(2, float64(unsafe.Sizeof(q)) + float64(unsafe.Sizeof(p)))); i, j = i+1, j+1 {
        fmt.Println("--------------")
        fmt.Printf("%t | %t\n", p, q)
        q = !q
        if j == int(math.Pow(2, float64(unsafe.Sizeof(p)))) {
            p = !p
            j = 0
        }
    }
	
    fmt.Println()
	p = true
    q = true
    fmt.Println("Truth table for p^q")
    fmt.Println("p    |  q  |")
	for i, j := 1, 1; i <= int(math.Pow(2, float64(unsafe.Sizeof(q)) + float64(unsafe.Sizeof(p)))); i, j = i+1, j+1 {
        fmt.Println("-------------------------")
        fmt.Printf("%t | %t | %t\n", p, q, p && q)
        q = !q
        if j == int(math.Pow(2, float64(unsafe.Sizeof(p)))) {
            p = !p
            j = 0
        }
    }

    fmt.Println()
	p = true
    q = true
    fmt.Println("Truth table for pvq")
    fmt.Println("p    |  q  |")
	for i, j := 1, 1; i <= int(math.Pow(2, float64(unsafe.Sizeof(q)) + float64(unsafe.Sizeof(p)))); i, j = i+1, j+1 {
        fmt.Println("-------------------------")
        fmt.Printf("%t | %t | %t\n", p, q, p || q)
        q = !q
        if j == int(math.Pow(2, float64(unsafe.Sizeof(p)))) {
            p = !p
            j = 0
        }
    }

    fmt.Println()
	p = true
    q = true
    fmt.Println("Truth table for px^q (p!=q)")
    fmt.Println("p    |  q  |")
	for i, j := 1, 1; i <= int(math.Pow(2, float64(unsafe.Sizeof(q)) + float64(unsafe.Sizeof(p)))); i, j = i+1, j+1 {
        fmt.Println("-------------------------")
        fmt.Printf("%t | %t | %t\n", p, q, p != q)
        q = !q
        if j == int(math.Pow(2, float64(unsafe.Sizeof(p)))) {
            p = !p
            j = 0
        }
    }

    fmt.Println()
	p = true
    q = true
    fmt.Println("Truth table for p->q (if p {q} el {true})")
    fmt.Println("p    |  q  |")
	for i, j := 1, 1; i <= int(math.Pow(2, float64(unsafe.Sizeof(q)) + float64(unsafe.Sizeof(p)))); i, j = i+1, j+1 {
        fmt.Println("-------------------------")
        var a bool = true
        if p {a = q}
        fmt.Printf("%t | %t | %t\n", p, q, a)
        q = !q
        if j == int(math.Pow(2, float64(unsafe.Sizeof(p)))) {
            p = !p
            j = 0
        }
    }

    fmt.Println()
	p = true
    q = true
    fmt.Println("Truth table for q->p (if q {p} el {true})")
    fmt.Println("q    |  p  |")
	for i, j := 1, 1; i <= int(math.Pow(2, float64(unsafe.Sizeof(q)) + float64(unsafe.Sizeof(p)))); i, j = i+1, j+1 {
        fmt.Println("-------------------------")
        var a bool = true
        if q {a = p}
        fmt.Printf("%t | %t | %t\n", q, p, a)
        p = !p
        if j == int(math.Pow(2, float64(unsafe.Sizeof(q)))) {
            q = !q
            j = 0
        }
    }

    fmt.Println()
	p = true
    q = true
    fmt.Println("Truth table for ~p->~q (if !p {!q} el {true})")
    fmt.Println("p    |  q  |")
	for i, j := 1, 1; i <= int(math.Pow(2, float64(unsafe.Sizeof(q)) + float64(unsafe.Sizeof(p)))); i, j = i+1, j+1 {
        fmt.Println("-------------------------")
        var a bool = true
        if !p {a = !q}
        fmt.Printf("%t | %t | %t\n", p, q, a)
        q = !q
        if j == int(math.Pow(2, float64(unsafe.Sizeof(p)))) {
            p = !p
            j = 0
        }
    }
}
