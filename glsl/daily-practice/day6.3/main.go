package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/KCkingcollin/go-help-func/ghf"
	"github.com/KCkingcollin/go-help-func/glf"
)

func main() {
	var data []uint32
    var sourceFile string = "shader.comp"
    shaderSource := ghf.LoadFile(sourceFile)
	shaderManager := glf.InitShaderManagerUint(shaderSource, "shader.comp")
	defer shaderManager.Cleanup()

	scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    input := scanner.Text()
    for _, r := range input {
        data = append(data, uint32(r))
    }
    for range input {
        data = append(data, 0)
    }

    output := shaderManager.Execute(data)
    var result []rune
    for _, char := range output {
        if char == '\x00' {
            break
        }
        result = append(result, rune(char))
    }
    fmt.Println(strings.ToLower(string(result)))
}
