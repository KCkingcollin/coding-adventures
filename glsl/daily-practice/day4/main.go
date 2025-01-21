package main

import (
	"fmt"
	"strconv"

	"github.com/KCkingcollin/go-help-func/ghf"
	"github.com/KCkingcollin/go-help-func/glf"
)

func main() {
	var input string
	var data []int32
    var sourceFile string = "shader.comp"
    shaderSource := ghf.LoadFile(sourceFile)
	shaderManager := glf.InitShaderManager(shaderSource, "shader.comp")
	defer shaderManager.Cleanup()

    fmt.Scan(&input)
    for _, char := range input {
        r, _ := strconv.Atoi(string(char))
        data = append(data, int32(r))
	}

    output := shaderManager.Execute(data)
    if output[0] == 0 {
        fmt.Println("NO")
    } else {
        fmt.Println("YES")
    }
}
