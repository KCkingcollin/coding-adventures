package main

import (
	"fmt"

	"github.com/KCkingcollin/go-help-func/ghf"
	"github.com/KCkingcollin/go-help-func/glf"
)

func main() {
    var n int
	var input int
	var data []int32 = make([]int32, 4)
    var sourceFile string = "shader.comp"
    shaderSource := ghf.LoadFile(sourceFile)
	shaderManager := glf.InitShaderManager(shaderSource, "shader.comp")
	defer shaderManager.Cleanup()

    fmt.Scan(&n)
    for i := 0; i < n; i++ {
        fmt.Scan(&input)
        data[input-1]++
	}

    output := shaderManager.Execute(data)
    fmt.Println(output[0]+output[1]+output[2]+output[3])
}
