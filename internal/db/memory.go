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

func NewMemoryDB() *memoryDB {
	return &memoryDB{db: map[string]*todo.Todo{}}
}

func (m *memoryDB) PutTodo(ctx context.Context, t *todo.Todo) error {
	m.lock.Lock()
	m.db[t.ID] = t
	m.lock.Unlock()
	return nil
}

func (m *memoryDB) GetAllTodos(ctx context.Context) ([]*todo.Todo, error) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	result := make([]*todo.Todo, len(m.db))
	i := 0
	for _, t := range m.db {
		result[i] = t
		i++
	}

	return result, nil
}
