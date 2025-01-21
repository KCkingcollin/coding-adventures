package main

import (
	"fmt"

	"github.com/KCkingcollin/go-help-func/ghf"
	"github.com/KCkingcollin/go-help-func/glf"
)

func main() {
	var data []uint32 = make([]uint32, 3)
    var sourceFile string = "shader.comp"
    shaderSource := ghf.LoadFile(sourceFile)
	shaderManager := glf.InitShaderManagerUint(shaderSource, "shader.comp")
	defer shaderManager.Cleanup()

    fmt.Scanf("%d %d %d", &data[0], &data[1], &data[2])

    output := shaderManager.Execute(data)
    fmt.Println(output[0])
}
