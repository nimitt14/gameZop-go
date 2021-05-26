package todo

import (
    "encoding/json"
    "fmt"
    "context"
	"github.com/go-redis/redis/v8"
	"github.com/nimitt14/gameZop-go/domain/todo"
)

type redisTodoRepo struct {
	db *redis.Client
}

func (repo *redisTodoRepo) Create(t todo.Todo) todo.Todo {
    str, _ := json.Marshal(t)
	err := repo.db.Set(context.Background(), fmt.Sprintf("TODO:%d",t.Id), str, 0).Err()
    if err != nil {
        panic(err)
    }
	return t
}

func NewRedisTodoRepo(db *redis.Client) todo.Repository {
	return &redisTodoRepo{db:db}
}