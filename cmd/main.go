package main 

import (
	"log"
	"fmt"
	"os"
	h "net/http"
	"github.com/nimitt14/gameZop-go/database"
	TodoRepo "github.com/nimitt14/gameZop-go/repository/todo"
	"github.com/nimitt14/gameZop-go/delivery/http"
	EqRepo "github.com/nimitt14/gameZop-go/repository/eventQueue"
)

func main(){
	rdb := database.NewRedisClient();
	amqp := database.NewAmqpClient();
	todoRepo := TodoRepo.NewRedisTodoRepo(rdb)
	eqRepo := EqRepo.NewAmqpRepo(amqp)
	server := http.NewHttpServer(todoRepo, eqRepo) 
	
	log.Fatal(h.ListenAndServe(fmt.Sprintf(":%s",os.Getenv("PORT")), server.Router))
}