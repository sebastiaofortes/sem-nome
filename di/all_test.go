package dinovo

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestWithInterfaces(t *testing.T) {

	app := NewContainer()

	assert.NotPanics(t, func() {
		app.AddDependencies(newController, newService, newRepository)

		app.InjectDependenciesInFunction(InitializeAPP)
	})
}

// Duplicate constructor registration
func TestPanicDplicateConstructorRegistration(t *testing.T) {

	app := NewContainer()

	assert.Panics(t, func() {
		app.AddDependencies(newController, newService, newRepository)
		
		app.AddDependencies(newService)

		app.InjectDependenciesInFunction(InitializeAPP)
	})
}

// More than one constructor found for the same type
func TestMoreOneConstructorForSameType(t *testing.T) {

	app := NewContainer()

	assert.Panics(t, func() {
		app.AddDependencies(newController, newService, newService2, newRepository)

		app.InjectDependenciesInFunction(InitializeAPP)
	})
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

func newService2(r RepositoryInterface) Service {
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
	return ""
}
