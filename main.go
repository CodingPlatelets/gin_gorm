package main

import (
	"github.com/WenkanHuang/gin_gorm/Cmd"
	"os"
)

func main() {
	if err := Cmd.Execute(); err != nil {
		println("start fail: ", err.Error())
		os.Exit(-1)
	}
}
