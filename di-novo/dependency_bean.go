package dinovo

import "reflect"

type DependencyBean struct {
	ContainsVariadicParam bool
	IsFunction            bool
	Name                  string
	constructorType       reflect.Type
	fnValue               reflect.Value
	constructorReturn     reflect.Type
	ParamTypes            []reflect.Type
}

func generateDependencyBean(fn interface{}) DependencyBean {
	fnType := reflect.TypeOf(fn)
	fnValue := reflect.ValueOf(fn)
	nameFunction := getFunctionName(fnValue)
	paramTypes := getParamTypes(fnType)
	returnType := getReturnType(fnType)
	isVariadic := fnType.IsVariadic()
	return DependencyBean{
		constructorType:       fnType,
		fnValue:               fnValue,
		Name:                  nameFunction,
		IsFunction:            true,
		ContainsVariadicParam: isVariadic,
		constructorReturn:     returnType,
		ParamTypes:            paramTypes,
	}
}

func generateDependencyBeansMap(funcs []interface{}) map[string]DependencyBean {
	ReflectTypeArray := make(map[string]DependencyBean)
	for _, fn := range funcs {
		dep := generateDependencyBean(fn)
		ReflectTypeArray[dep.Name] = dep
	}
	return ReflectTypeArray
}

func GenerateDependenciesList(funcs ...interface{}) map[string]DependencyBean {
	return generateDependencyBeansMap(funcs)
}
