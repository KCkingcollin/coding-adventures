// A. Way Too Long Words
// time limit per test
// 1 second
// memory limit per test
// 256 megabytes
//
// Sometimes some words like "localization" or "internationalization" are so long that writing them many times in one text is quite tiresome.
//
// Let's consider a word too long, if its length is strictly more than 10 characters. All too long words should be replaced with a special abbreviation.
//
// This abbreviation is made like this: we write down the first and the last letter of a word and between them we write the number of letters between the first and the last letters. That number is in decimal system and doesn't contain any leading zeroes.
//
// Thus, "localization" will be spelt as "l10n", and "internationalization» will be spelt as "i18n".
//
// You are suggested to automatize the process of changing the words with abbreviations. At that all too long words should be replaced by the abbreviation and the words that are not too long should not undergo any changes.
// Input
//
// The first line contains an integer n (1 ≤ n ≤ 100). Each of the following n lines contains one word. All the words consist of lowercase Latin letters and possess the lengths of from 1 to 100 characters.
// Output
//
// Print n lines. The i-th line should contain the result of replacing of the i-th word from the input data.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        input := scanner.Text()
        word := []rune(input)
        if _, err := strconv.Atoi(input); err != nil {
            if len(word) > 10 {
                char1 := string(word[0])
                char2 := string(word[len(word)-1])
                num := strconv.Itoa(len(word)-2)
                fmt.Println(char1+num+char2)
            } else {
                fmt.Println(input)
            }
        }
    }
}
