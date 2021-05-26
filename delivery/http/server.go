package http 

import (
	"github.com/nimitt14/gameZop-go/domain/todo"
	"github.com/gorilla/mux"
	// h "net/http"
	"github.com/nimitt14/gameZop-go/repository/eventQueue"
)

type httpServer struct {
	Router *mux.Router
	todoRepo todo.Repository
	eqRepo eventQueue.EventQueueRepo
}

func NewHttpServer(todoRepo todo.Repository, eqRepo eventQueue.EventQueueRepo) *httpServer {
	httpServerObj := &httpServer{
		Router : mux.NewRouter().StrictSlash(true),
		todoRepo : todoRepo,
		eqRepo : eqRepo,
	}

        // var handler h.Handler
        // handler = route.HandlerFunc
        // handler = utils.Logger(handler, route.Name)
        httpServerObj.Router.
            Methods("POST").
            Path("/todo").
            Name("Create").
            HandlerFunc(httpServerObj.Create)

	return httpServerObj
}



