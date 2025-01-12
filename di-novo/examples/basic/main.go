package main

import (
	"fmt"
	"net/http"

	dinovo "github.com/sebastiaofortes/novo-de/di-novo"
)

type Repository struct {
}

func (r Repository) GetData() {
	fmt.Println("Chamando GetData")
}

type RepositoryInterface interface {
	GetData()
}

type Service struct {
	R RepositoryInterface
}

func (s Service) Apply() {
	s.R.GetData()
	fmt.Println("Chamando Apply")
}

type ServiceInterface interface {
	Apply()
}

type Controller struct {
	S ServiceInterface
}

func main() {

	app := dinovo.NewContainer()

	app.AddDependencies(newController, newService, newRepository)

	app.InjectDependenciesInFunction(InitializeAPP)
}

func newRepository() Repository {
	fmt.Println("Criando Repository")
	return Repository{}
}

func newService(r RepositoryInterface) Service {
	fmt.Println("Criando Service")
	return Service{
		R: r,
	}
}

func newController(s ServiceInterface) Controller {
	fmt.Println("Criando controller")
	return Controller{
		S: s,
	}
}

func (c Controller) handler(w http.ResponseWriter, r *http.Request) {
	c.S.Apply()
	fmt.Fprintf(w, "Olá, Mundo!")
}

func InitializeAPP(c Controller) string {
	http.HandleFunc("/", c.handler)
	http.ListenAndServe(":8080", nil)
	return ""
}
