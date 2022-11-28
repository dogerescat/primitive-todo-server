package db

import (
	"context"
	"primitive-todo-server/internal/todo"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPutTodo(t *testing.T) {
	t.Parallel()
	todo1 := &todo.Todo{
		ID:    "ddd046bc-7035-412b-9009-4f75d62241a4",
		Title: "code reading",
	}
	tests := map[string]struct {
		todo     *todo.Todo
		expected map[string]*todo.Todo
	}{
		"put": {
			todo:     todo1,
			expected: map[string]*todo.Todo{todo1.ID: todo1},
		},
	}
	ctx := context.Background()
	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			d := &memoryDB{db: map[string]*todo.Todo{}}
			if err := d.PutTodo(ctx, test.todo); err != nil {
				t.Fatalf("failed to put a todo: %s", err.Error())
			}
			if diff := cmp.Diff(test.expected, d.db); diff != "" {
				t.Errorf("\n(-expected, + actual)\n%s", diff)
			}
		})
	}
}
