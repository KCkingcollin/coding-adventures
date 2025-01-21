package main

import (
	"fmt"
	"unicode"
)

func isNotVowel(p rune) bool {
    set := []rune{'a', 'e', 'i', 'o', 'u', 'y'}
    for _, val := range set {
        if unicode.ToLower(p) == val {
            return false
        }
    }
    return true
}

func main() {
    var p string
    fmt.Scan(&p)
    var result []rune
    for _, value := range []rune(p) {
        if isNotVowel(value) {
            result = append(result, '.')
            result = append(result, unicode.ToLower(value))
        }
    }
    fmt.Println(string(result))
}
