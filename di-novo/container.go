package dinovo

import (
	"fmt"
	"log"
	"maps"
	"reflect"
)

type Container struct {
	dependencies map[string]DependencyBean
}

func NewContainer() Container {
	return Container{
		dependencies: make(map[string]DependencyBean),
	}
}

func (c *Container) AddDependencies(deps ...interface{}) {
	// Gera o array com as dependencias
	ReflectTypeArray := generateDependencyBeansMap(deps)
	c.checkingNameUnit(ReflectTypeArray)
	maps.Copy(c.dependencies, ReflectTypeArray)
}

func (c *Container) checkingNameUnit(reflectTypeArray map[string]DependencyBean) {
	for _, v := range reflectTypeArray {
		if _, exists := c.dependencies[v.Name]; exists {
			panic("Duplicate constructor registration")
		}
	}
}

func (c *Container) InjectDependenciesInFunction(startFunc interface{}) {
	fmt.Println("Iniciando Injeção...")
	quantDep := len(c.dependencies)
	fmt.Println(quantDep, " dependiencias registradas")

	dep := generateDependencyBean(startFunc)

	args := c.getDependenciesForConstructor(dep)

	fmt.Println("Chamando função ", dep.Name, " ...")
	fmt.Println()

	// Chamando o construtor e enviando os parametros encontrados
	dep.fnValue.Call(args)
}

func (c *Container) getDependenciesForConstructor(constructor DependencyBean) []reflect.Value {
	args := []reflect.Value{}
	fmt.Printf("constructor: %s, number of parameters: %d\n", constructor.Name, len(constructor.ParamTypes))
	for position, paramType := range constructor.ParamTypes {

		// Check if trhe variadic param
		if constructor.ContainsVariadicParam {
			if position == (len(constructor.ParamTypes) - 1) {
				// Redice slice elements to single element
				paramType = ReduceSliceToSingleElement(paramType)
			}
		}

		// Procura na lista de um contrutuores um tipo igual ao do parametro
		injectableDependencies := c.searchInjectableDependencies(paramType, constructor.constructorReturn, constructor.ContainsVariadicParam)

		for _, injectableDependency := range injectableDependencies {
			if injectableDependency.IsFunction {
				argumants := c.getDependenciesForConstructor(injectableDependency)
				resp := injectableDependency.fnValue.Call(argumants)
				args = append(args, resp...)
				log.Println("Injecting: ", injectableDependency.Name, " in ", constructor.Name)
			} else {
				args = append(args, injectableDependency.fnValue)
			}
		}
	}
	return args
}

func (c *Container) searchInjectableDependencies(paramType reflect.Type, returnType reflect.Type, isVariadic bool) []DependencyBean {
	var dependenciesFound []DependencyBean
	var depsFound []DependencyBean
	if isInterface(paramType) {
		dependenciesFound = c.searchImplementations(paramType)
	} else {
		dependenciesFound = c.searchTypes(paramType)
	}
	if len(dependenciesFound) > 1 {
		if isVariadic {
			depsFound = dependenciesFound
		} else {
			panic("Mais de um construtor encontrado para um mesmo tipo")
		}
	} else if len(dependenciesFound) == 0 {
		panic("nemhum construtor para o parametro foi encontrado")
	} else {
		depsFound = append(depsFound, dependenciesFound[0])
	}
	return depsFound
}

func (f *Container) searchTypes(paramType reflect.Type) []DependencyBean {
	dependenciesFound := []DependencyBean{}
	for fnName, dependency := range f.dependencies {
		for i := 0; i < dependency.constructorType.NumOut(); i++ {
			returnType := dependency.constructorType.Out(i)
			if returnType == paramType {
				fmt.Println("parameter: ", paramType, " compatible => ", fnName, " type ", returnType)
				dependenciesFound = append(dependenciesFound, dependency)
			}
		}
	}
	return dependenciesFound
}

func (f *Container) searchImplementations(paramType reflect.Type) []DependencyBean {
	dependenciesFound := []DependencyBean{}
	for fnName, dependency := range f.dependencies {
		for i := 0; i < dependency.constructorType.NumOut(); i++ {
			returnType := dependency.constructorType.Out(i)
			implements := implementsInterface(returnType, paramType)
			if implements {
				fmt.Println("parameter: ", paramType, " implementation => ", fnName, " type ", returnType)
				dependenciesFound = append(dependenciesFound, dependency)
			}
		}
	}
	return dependenciesFound
}
