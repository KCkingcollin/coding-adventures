package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
    // make a map of coin values and the amount of that value coin as the value while also counting up a total
    scanner := bufio.NewScanner(os.Stdin)
    for i := 0; i < 2; i++ {
        scanner.Scan()
    }
    line := strings.Split(scanner.Text(), " ")
    coins := make(map[int]int)
    var totalValue int
    var maxNum int
    var minNum int
    for i, value := range line {
        coinValue, _ := strconv.Atoi(value)
        coins[coinValue]++
        totalValue += coinValue
        if coinValue > maxNum {
            maxNum = coinValue
            if i == 0 {minNum = coinValue}
        } else if coinValue < minNum {
            minNum = coinValue
        }
    }

    // count how many coins are needed to get to greater than half the total
    halfValueUp := totalValue / 2
    var coinsToTake int
    var curentValue int
    firstLoop:
    for i := maxNum; i >= minNum; i-- {
        if curentValue + i * coins[i] > halfValueUp {
            for j := 1; j <= coins[i]; j++{
                if curentValue + i * j > halfValueUp {
                    coinsToTake += j
                    break firstLoop
                }
            }
        } else if coins[i] > 0 {
            curentValue += i * coins[i]
            coinsToTake += coins[i]
        }
    }

    // print the answer baby
    fmt.Println(coinsToTake)
}
