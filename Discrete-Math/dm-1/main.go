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
    fmt.Println("Truth table for pxvq (p!=q)")
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
    fmt.Println("Truth table for p->q (!p || q)")
    fmt.Println("p    |  q  |")
	for i, j := 1, 1; i <= int(math.Pow(2, float64(unsafe.Sizeof(q)) + float64(unsafe.Sizeof(p)))); i, j = i+1, j+1 {
        fmt.Println("-------------------------")
        fmt.Printf("%t | %t | %t\n", p, q, !p || q)
        q = !q
        if j == int(math.Pow(2, float64(unsafe.Sizeof(p)))) {
            p = !p
            j = 0
        }
    }

    fmt.Println()
	p = true
    q = true
    fmt.Println("Truth table for q->p (!q || p)")
    fmt.Println("q    |  p  |")
	for i, j := 1, 1; i <= int(math.Pow(2, float64(unsafe.Sizeof(q)) + float64(unsafe.Sizeof(p)))); i, j = i+1, j+1 {
        fmt.Println("-------------------------")
        fmt.Printf("%t | %t | %t\n", q, p, !q || p)
        p = !p
        if j == int(math.Pow(2, float64(unsafe.Sizeof(q)))) {
            q = !q
            j = 0
        }
    }

    fmt.Println()
	p = true
    q = true
    fmt.Println("Truth table for ~p->~q (p || !q)")
    fmt.Println("p    |  q  |")
	for i, j := 1, 1; i <= int(math.Pow(2, float64(unsafe.Sizeof(q)) + float64(unsafe.Sizeof(p)))); i, j = i+1, j+1 {
        fmt.Println("-------------------------")
        fmt.Printf("%t | %t | %t\n", p, q, p || !q)
        q = !q
        if j == int(math.Pow(2, float64(unsafe.Sizeof(p)))) {
            p = !p
            j = 0
        }
    }

    fmt.Println()
	p = true
    q = true
    fmt.Println("Truth table for p<->q (p == q)")
    fmt.Println("p    |  q  |")
	for i, j := 1, 1; i <= int(math.Pow(2, float64(unsafe.Sizeof(q)) + float64(unsafe.Sizeof(p)))); i, j = i+1, j+1 {
        fmt.Println("-------------------------")
        fmt.Printf("%t | %t | %t\n", p, q, p == q)
        q = !q
        if j == int(math.Pow(2, float64(unsafe.Sizeof(p)))) {
            p = !p
            j = 0
        }
    }

    fmt.Println()
	p = true
    q = true
    var r bool = true
    fmt.Println("Truth table for pvq->~r (!(p || q) || !r)")
    fmt.Println("  p  |  q  |  r  |  pvq  |  ~r  |  pvq->~r")
	for i, j, l := 1, 1, 1; i <= int(math.Pow(2, float64(unsafe.Sizeof(q)) + float64(unsafe.Sizeof(p)) + float64(unsafe.Sizeof(r)))); i, j, l = i+1, j+1, l+1 {
        fmt.Println("--------------------------------------------------------")
        fmt.Printf("%t | %t | %t | %t | %t | %t\n", q, p, r, p || q, !r, !(p || q) || !r)
        r = !r
        if j == int(math.Pow(2, float64(unsafe.Sizeof(p)))) {
            p = !p
            j = 0
        }
        if l == int(math.Pow(2, float64(unsafe.Sizeof(q)) + float64(unsafe.Sizeof(p)))) {
            q = !q
            l = 0
        }
    }
}
