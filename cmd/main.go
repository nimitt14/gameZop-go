package main 

import (
	"log"
	"os"
	h "net/http"
	"github.com/nimitt14/go-node/database"
	TodoRepo "github.com/nimitt14/go-node/repository/todo"
	"github.com/nimitt14/go-node/delivery/http"
	EqRepo "github.com/nimitt14/go-node/repository/eventQueue"
)

func main(){
	rdb := database.NewRedisClient();
	amqp := database.NewAmqpClient();
	todoRepo := TodoRepo.NewRedisTodoRepo(rdb)
	eqRepo := EqRepo.NewAmqpRepo(amqp)
	server := http.NewHttpServer(todoRepo, eqRepo) 
	
	log.Fatal(h.ListenAndServe(fmt.Sprintf(":%s",os.Getenv("PORT")), server.Router))
}