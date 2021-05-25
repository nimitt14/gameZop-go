package eventQueue

import (
    "fmt"
	"github.com/nimitt14/go-node/domain/todo"
	"github.com/streadway/amqp"
)

type amqpRepo struct {
	db *amqp.Channel
}

type EventQueueRepo interface{
	Create(t todo.Todo) error
}

func (repo *amqpRepo) Create(t todo.Todo) error {	
	q, err := repo.db.QueueDeclare(
	"hello", // name
	false,   // durable
	false,   // delete when unused
	false,   // exclusive
	false,   // no-wait
	nil,     // arguments
  )
  
  if err!=nil{
	  return err
  }

  body := fmt.Sprintf("TODO:%d", t.Id)
  err = repo.db.Publish(
	"",     // exchange
	q.Name, // routing key
	false,  // mandatory
	false,  // immediate
	amqp.Publishing {
	  ContentType: "text/plain",
	  Body:        []byte(body),
	})
	
	return err
}

func NewAmqpRepo(db *amqp.Channel) EventQueueRepo {
	return &amqpRepo{db:db}
}
