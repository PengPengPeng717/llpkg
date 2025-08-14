package main

import (
	"github.com/goplus/lib/py"
)

func main() {
	// 初始化 Python
	py.Initialize()
	defer py.Finalize()
}
