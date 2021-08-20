package main

import "fmt"

// 运行：go run *.go
// 参考：https://www.pixelstech.net/article/1621090998-Run-code-with-multiple-files-in-the-same-main-package-in-GoLang
func main() {
	fmt.Println("main")

	// 可以调用处于同一个package的函数
	hey()
	bye()
}
