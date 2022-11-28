package db

import (
	"context"
	"primitive-todo-server/internal/todo"
)

type DB interface {
	PutTodo(ctx context.Context, t *todo.Todo) error
}
