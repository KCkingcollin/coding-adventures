package main

import (
	"fmt"

	"github.com/KCkingcollin/go-help-func/ghf"
	"github.com/KCkingcollin/go-help-func/glf"
)

func main() {
    var n int
    var input string
    var data []int32
    var answer []rune
    var sourceFile string = "shader.comp"
    shaderSource := ghf.LoadFile(sourceFile)
	shaderManager := glf.InitShaderManager(shaderSource, "shader.comp")
	defer shaderManager.Cleanup()

    fmt.Scanln(&n)
    for i := 0; i < n; i++ {
        fmt.Scanln(&input)
        data = []int32{}
        for _, char := range input {
            data = append(data, int32(char))
        }
        output := shaderManager.Execute(data)
        for _, char := range output {
            if char == '\x00' {
                break
            }
            answer = append(answer, rune(char))
        }
        answer = append(answer, '\n')
    }

    fmt.Println(string(answer))
}
