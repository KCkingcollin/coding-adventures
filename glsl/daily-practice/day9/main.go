package main

import (
	"fmt"

	"github.com/KCkingcollin/go-help-func/ghf"
	"github.com/KCkingcollin/go-help-func/glf"
)

func main() {
    var n int
	var input int32
    var sourceFile string = "shader.comp"
    shaderSource := ghf.LoadFile(sourceFile)
	shaderManager := glf.InitShaderManager(shaderSource, "shader.comp")
	defer shaderManager.Cleanup()

    fmt.Scan(&n)
    data := make([]int32, n+1)
    data[0] = int32(n)
    for i := 0; i < n; i++ {
        fmt.Scan(&input)
        data[i+1] = input
	}
    var q int
    fmt.Scan(&q)
    data = append(data, int32(q))
    for i := 0; i < q; i++ {
        fmt.Scan(&input)
        data = append(data, input)
	}

    output := shaderManager.Execute(data)

    for _, result := range output {
        if result == 0 {
            break
        }
        fmt.Printf("%d\n", result)
    }
    fmt.Println()
}
