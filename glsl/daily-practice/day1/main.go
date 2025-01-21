package main

import (
	"fmt"

	"github.com/KCkingcollin/go-help-func/ghf"
	"github.com/KCkingcollin/go-help-func/glf"
)

func main() {
	var input int
	var data []int32
    var sourceFile string = "shader.comp"
    shaderSource := ghf.LoadFile(sourceFile)
    shaderManager := glf.InitShaderManager(shaderSource, "shader.comp")
    defer shaderManager.Cleanup()

    fmt.Scan(&input)
    data = append(data, int32(input))
    output := shaderManager.Execute(data)

    if output[0] == 1 {
        fmt.Println("YES")
    } else {
        fmt.Println("NO")
    }
}
