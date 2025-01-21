package main

import (
	"fmt"
	"math/big"
)

func strToBig(value string) *big.Int {
    uint128 := new(big.Int)
    _, ok := uint128.SetString(value, 2)
    if !ok {
        panic("could not convert the string")
    }
	return uint128
}

func main() {
    var w string
    fmt.Scan(&w)
    input := strToBig(w)
    prevBit := input.Bit(0)
    adjBits := 1
    for i := 1; i < len(w); i++ {
        curBit := input.Bit(i)
        if curBit == prevBit {
            adjBits++
        } else {
            adjBits = 1
        }
        if adjBits >= 7 {
            fmt.Println("YES")
            return
        }
        prevBit = curBit
    }
    fmt.Println("NO")
}
