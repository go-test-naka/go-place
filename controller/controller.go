package controller

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/go-test-naka/go-place/usecase/person"

	"github.com/go-test-naka/go-place/log"
)

type Controller struct {
	pc person.Create
	ps person.Search
}

func CreateNewController(createPerson person.Create, searcPerson person.Search) *Controller {
	return &Controller{
		createPerson,
		searcPerson,
	}
}

func (c *Controller) Execute() {
	http.HandleFunc("/ping", c.ping)
	http.HandleFunc("/person", c.handler)
	http.HandleFunc("/persons", c.handler)

	err := http.ListenAndServe(":8081", nil)
	if errors.Is(err, http.ErrServerClosed) {
		log.Info("Server closed")
	} else if err != nil {
		log.Error("Error starting server", err)
		os.Exit(1)
	}

}

func (c *Controller) handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		queryParams := r.URL.Query()
		name := queryParams.Get("name")
		if len(name) > 0 {
			c.findPersonByName(w, r)
		} else {
			c.listPersons(w, r)
		}
	case http.MethodPost:
		c.createPerson(w, r)
	}
}

func (c *Controller) ping(w http.ResponseWriter, r *http.Request) {
	log.Debug("Received ping request")
	io.WriteString(w, "UP!\n")
}
func (c *Controller) createPerson(w http.ResponseWriter, r *http.Request) {
	log.Debug("Received /person request")
	io.WriteString(w, "Create person\n")
	c.pc.Create()
}
func (c *Controller) listPersons(w http.ResponseWriter, r *http.Request) {
	log.Debug("Received /persons request")
	io.WriteString(w, "List persons\n")
	c.ps.SearchAll()
}

func (c *Controller) findPersonById(w http.ResponseWriter, r *http.Request) {
	log.Debug("Received /person/{id} request")
	path := r.URL.Path
	parts := strings.Split(path, "/")
	idStr := parts[len(parts)-1]

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Erro ao converter 'id' para uint", http.StatusBadRequest)
		return
	}

	io.WriteString(w, fmt.Sprintf("Find person with id %v", id))
	c.ps.SearchById(uint(id))
}

func (c *Controller) findPersonByName(w http.ResponseWriter, r *http.Request) {
	log.Debug("Received /person/name/{name} request")
	queryParams := r.URL.Query()
	name := queryParams.Get("name")
	log.Info(fmt.Sprintf("Find person with name %v", name))
	c.ps.SearchByName(name)
}
