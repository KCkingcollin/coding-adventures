package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/KCkingcollin/go-help-func/ghf"
	"github.com/KCkingcollin/go-help-func/glf"
)

func main() {
	var data []int32
    var sourceFile string = "shader.comp"
    shaderSource := ghf.LoadFile(sourceFile)
	shaderManager := glf.InitShaderManager(shaderSource, "shader.comp")
	defer shaderManager.Cleanup()

    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    input := scanner.Text()
    for _, r := range input {
        data = append(data, int32(r))
	}

    output := shaderManager.Execute(data)
    if output[0] == 0 {
        fmt.Println("NO")
    } else {
        fmt.Println("YES")
    }
}
