package main

import (
	"fmt"

	"github.com/KCkingcollin/go-help-func/ghf"
	"github.com/KCkingcollin/go-help-func/glf"
)

func main() {
	var data []float64 = make([]float64, 2)
    var sourceFile string = "shader.comp"
    shaderSource := ghf.LoadFile(sourceFile)
	shaderManager := glf.InitShaderManagerFloat(shaderSource, "shader.comp")
	defer shaderManager.Cleanup()

    fmt.Scanf("%f %f", &data[0], &data[1])
    output := shaderManager.Execute(data)
    fmt.Println(int(output[0]))
}
