package main

import (
	"fmt"

	"github.com/KCkingcollin/go-help-func/ghf"
	"github.com/KCkingcollin/go-help-func/glf"
)

func main() {
    var n int
	var input int
	var data []int32
    var sourceFile string = "shader.comp"
    shaderSource := ghf.LoadFile(sourceFile)
	shaderManager := glf.InitShaderManager(shaderSource, "shader.comp")
	defer shaderManager.Cleanup()
    fmt.Scan(&n)
    for i := 0; i < n+1; i++ {
        fmt.Scan(&input)
        data = append(data, int32(input))
	}

    output := shaderManager.Execute(data)
    fmt.Println(output[0])
}
