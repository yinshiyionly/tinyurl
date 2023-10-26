package main

import (
    "fmt"
    "tinyurl/pkg/generator"
)

func main() {
    fmt.Println(generator.Hash("https://www.baidu.com/"))
}
