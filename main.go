package main

import (
	"fmt"
	"os"

	"github.com/robertkrimen/otto"
)

func main() {
	filePath := "encrypt.js"
	//先读入文件内容
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	vm := otto.New()

	_, err = vm.Run(bytes)
	if err != nil {
		panic(err)
	}

	data := "你需要传给JS函数的参数"
	//encodeInp是JS函数的函数名
	value, err := vm.Call("encodeInp", nil, data)
	if err != nil {
		panic(err)
	}
	fmt.Println(value.String())
}
