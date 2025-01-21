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
    var sourceFile string = "shader.comp"
    shaderSource := ghf.LoadFile(sourceFile)
	shaderManager := glf.InitShaderManager(shaderSource, "shader.comp")
	defer shaderManager.Cleanup()

    fmt.Scanf("%d", &n)
    for i := 0; i < n; i++ {
        fmt.Scanf("%s", &input)
		data = append(data, int32(input[0]))
		data = append(data, int32(input[1]))
		data = append(data, int32(input[2]))
	}

    output := shaderManager.Execute(data)

    fmt.Println(output[0])
}
