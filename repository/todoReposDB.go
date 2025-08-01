package repository

import (
	"context"
	"time"

	"github.com/hsnylmz1283/goApi/models"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type TodoRepositoryDB struct {
	TodoCollection *mongo.Collection
}

type TodoRepository interface {
	Insert(todo models.Todo) (bool, error)
}

func (t TodoRepositoryDB) Insert(todo models.Todo) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := t.TodoCollection.InsertOne(ctx, todo)
	if result.InsertedID == nil || err != nil {
		// errors.New("failed add")
		return false, err
	}
	return true, nil
}

func NewTodoRepositoryDb(dbClient *mongo.Collection) TodoRepositoryDB {
	return TodoRepositoryDB{TodoCollection: dbClient}
}
