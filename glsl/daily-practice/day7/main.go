package main

import (
	"fmt"

	"github.com/KCkingcollin/go-help-func/ghf"
	"github.com/KCkingcollin/go-help-func/glf"
	"github.com/go-gl/mathgl/mgl32"
)

func main() {
    var n int
	var input [4]float32
	var data []mgl32.Vec4
    var sourceFile string = "shader.comp"
    shaderSource := ghf.LoadFile(sourceFile)
	shaderManager := glf.InitShaderManagerVec3(shaderSource, "shader.comp")
	defer shaderManager.Cleanup()

    fmt.Scan(&n)
    for i := 0; i < n; i++ {
        fmt.Scanf("%f %f %f", &input[0], &input[1], &input[2])
        data = append(data, [4]float32{input[0], input[1], input[2], 0})
	}

    output := shaderManager.Execute(data)
    if output[0][0] == 0 {
        fmt.Println("NO")
    } else {
        fmt.Println("YES")
    }
}
