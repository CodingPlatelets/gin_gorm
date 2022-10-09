package main

import (
	"os"

	"github.com/WenkanHuang/gin_gorm/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		println("start fail: ", err.Error())
		os.Exit(-1)
	}
}
