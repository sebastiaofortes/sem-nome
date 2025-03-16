# Dinovo - Framework de Injeção de Dependência para Golang

Dinovo é um framework leve e dinâmico para injeção de dependências em Golang, suportando injeção em construtores com parâmetros variádicos.

## Recursos

- Injeção automática de dependências
- Suporte a construtores com parâmetros variádicos
- Registro dinâmico de funções e instâncias
- Resolução de dependências baseada em reflexão

## Instalação

Adicione o `dinovo` ao seu projeto:

```sh
 go get github.com/sebastiaofortes/dinovo
```

## Uso Básico

### Criando um Container de Dependências

```go
package main

import (
	"fmt"
	"github.com/sebastiaofortes/dinovo"
)

func main() {
	container := dinovo.NewContainer()

	// Registrando dependências
	container.AddDependencies(NewService, NewRepository)

	// Injetando dependências automaticamente
	container.InjectDependenciesInFunction(StartApplication)
}

func NewService(repo Repository) Service {
	return Service{repo: repo}
}

func NewRepository() Repository {
	return Repository{}
}

func StartApplication(service Service) {
	fmt.Println("Aplicativo iniciado com serviço:", service)
}

// Definição de tipos

type Service struct {
	repo Repository
}

type Repository struct {}
```

## Funcionalidades Avançadas

### Suporte a Variadic Parameters

```go
func NewHandler(services ...Service) Handler {
	return Handler{services: services}
}
```

O `dinovo` gerencia automaticamente a redução de slices para suportar parâmetros variádicos.

### Tratamento de Interfaces

Se um construtor retorna uma interface, `dinovo` automaticamente resolve implementações compatíveis.

```go
func NewRepository() RepositoryInterface {
	return &Repository{}
}
```

## Erros e Depuração

- **Duplicate constructor registration**: ocorre quando duas funções com o mesmo nome são registradas.
- **Mais de um construtor encontrado para um mesmo tipo**: ocorre quando mais de um construtor pode ser usado para um mesmo tipo.
- **Nenhum construtor encontrado**: ocorre quando não há um construtor registrado para uma dependência requerida.

## Contribuição

Contribuições são bem-vindas! Fique à vontade para abrir issues e pull requests.

## Licença

Este projeto é licenciado sob a Licença MIT - veja o arquivo [LICENSE](LICENSE) para detalhes.

