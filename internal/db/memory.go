package db

import (
	"context"
	"primitive-todo-server/internal/todo"
	"sync"
)

var _ DB = (*memoryDB)(nil)

type memoryDB struct {
	db   map[string]*todo.Todo
	lock sync.RWMutex
}

func (m *memoryDB) PutTodo(ctx context.Context, t *todo.Todo) error {
	m.lock.Lock()
	m.db[t.ID] = t
	m.lock.Unlock()
	return nil
}
