package http

import (
	"encoding/json"
	"net/http"
	"primitive-todo-server/internal/db"
	"primitive-todo-server/internal/todo"

	"github.com/google/uuid"
)

var _ http.Handler = (*createHandler)(nil)

type createHandler struct {
	db db.DB
}

func (h *createHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var t todo.Todo
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	t.ID = uuid.New().String()
	if err := h.db.PutTodo(r.Context(), &t); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(&t); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
