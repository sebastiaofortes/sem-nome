# sem-nome - Framework de Injeção de Dependência para Golang

sem-nome é um framework leve e dinâmico para injeção de dependências em Golang, suportando injeção em construtores com parâmetros variádicos.

## Funcionamento

O contêiner de injeção de dependências funciona por meio de um processo de empilhamento e desempilhamento (injeção). Na fase de empilhamento, as dependências de um objeto são identificadas, assim como as dependências dessas dependências, em um ciclo recursivo que termina ao alcançar objetos que não necessitam de injeção. Na fase de injeção, os objetos mapeados na pilha são criados de forma que aqueles no topo da pilha sejam utilizados como parâmetros para a construção dos objetos nas camadas inferiores.

![flow](flow-1.png)

## Recursos

- Injeção automática de dependências
- Suporte a construtores com parâmetros variádicos
- Registro dinâmico de funções e instâncias
- Resolução de dependências baseada em reflexão

## Instalação

Adicione o `sem-nome` ao seu projeto:

```sh
 go get github.com/sebastiaofortes/sem-nome
```

## Uso Básico

### Criando um Container de Dependências

```go
package main

import (
	"fmt"
	"github.com/sebastiaofortes/sem-nome/di"
)

func main() {
	container := di.NewContainer()

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

O `sem-nome` gerencia automaticamente a redução de slices para suportar parâmetros variádicos.

### Tratamento de Interfaces

Se um construtor retorna uma interface, `sem-nome` automaticamente resolve implementações compatíveis.

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

