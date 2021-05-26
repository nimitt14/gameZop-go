package http

import (
    "encoding/json"
	h "net/http"
	"io"
	"io/ioutil"
	"github.com/nimitt14/gameZop-go/domain/todo"

    // "github.com/gorilla/mux"
)

func (s *httpServer) Create(w h.ResponseWriter, r *h.Request) {
    var todo todo.Todo
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
        panic(err)
    }
    if err := r.Body.Close(); err != nil {
        panic(err)
    }
    if err := json.Unmarshal(body, &todo); err != nil {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(422) // unprocessable entity
        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
        }
    }

	t := s.todoRepo.Create(todo)
	err = s.eqRepo.Create(todo)
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(h.StatusCreated)
    if err := json.NewEncoder(w).Encode(t); err != nil {
        panic(err)
    }

}
