package di

import (
	"fmt"
	"reflect"
	"runtime"
)

func getFunctionName(i reflect.Value) string {
	return runtime.FuncForPC(i.Pointer()).Name()
}

func getParamTypes(fnType reflect.Type) []reflect.Type {
	var paramTypes []reflect.Type
	for i := 0; i < fnType.NumIn(); i++ {
		paramTypes = append(paramTypes, fnType.In(i))
	}
	return paramTypes
}

func getReturnType(fnType reflect.Type) reflect.Type {
	if fnType.NumOut() == 1 {
		return fnType.Out(0)
	} else {
		message := fmt.Sprintf("Erro, a função %s deve possuir um único tipo de retrono \n", fnType.Name())
		panic(message)
	}
}

func ReduceSliceToSingleElement(sliceElem reflect.Type) reflect.Type {
	if sliceElem.Kind() == reflect.Slice {
		elementType := sliceElem.Elem()
		return elementType
	}
	return sliceElem
}

func isInterface(r reflect.Type) bool {
	return r.Kind() == reflect.Interface
}

// Função para verificar se uma struct implementa uma interface
func implementsInterface(structType reflect.Type, interfaceType reflect.Type) bool {
	return structType.Implements(interfaceType)
}
