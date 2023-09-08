package scriptsystem

import (
	"os"
	"sync"

	"github.com/robertkrimen/otto"

	"example.com/otto-test/internal/logger"
)

var (
	vmLogger       *logger.Logger = logger.NewLogger("VMFACTORY")
	apiFuncExports sync.Map
	vmList         sync.Map
)

// NewVM
// 创建otto虚拟机
//
//	@param filePath js脚本文件路径
func NewVM(filePath string) {
	vm := otto.New()
	apiFuncExports.Range(func(key, value any) bool {
		vm.Set(key.(string), value)
		return true
	})

	bytes, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	_, err = vm.Run(bytes)
	if err != nil {
		panic(err)
	}

	vmList.Store(filePath, vm)
}

// PutApiFuncExport
// 添加对外暴露的api接口
//
//	@param funcName 函数名
//	@param apiFunc otto函数 func(call otto.FunctionCall) otto.Value
func PutApiFuncExport(funcName string, apiFunc func(call otto.FunctionCall) otto.Value) {
	apiFuncExports.Store(funcName, apiFunc)
}

// ActionCall
// 用于执行生命周期函数
//
//	@param action
func ActionCall(action string) {
	vmList.Range(func(key, value any) bool {
		vm := value.(*otto.Otto)
		if _, err := vm.Call(action, nil, nil); err != nil {
			// TODO:输出错误
			vmLogger.Log(logger.Error, "action call fail %s", err)
		}
		return true
	})
}
