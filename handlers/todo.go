package handlers

import (
	"fmt"
	"go_fullstack/components"
	"go_fullstack/types"
	"go_fullstack/views"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Todo struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (h *Handler) HandleSearchTodo(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("search")
	todos, err := h.store.SearchTodos(text)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	components.TodoList(todos).Render(r.Context(), w)
}

func (h *Handler) HandleTodo(w http.ResponseWriter, r *http.Request) {
	todos, err := h.store.GetTodos()
	fmt.Println(todos)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	views.Todo(todos).Render(r.Context(), w)
}

func (h *Handler) AddTodo(w http.ResponseWriter, r *http.Request) {
	todo := &types.Todo{
		Title:     r.FormValue("title"),
		Completed: false,
	}
	t, err := h.store.CreateTodo(todo)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	components.TodoTitle(t).Render(r.Context(), w)
}

func (h *Handler) HandleDeleteTodo(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = h.store.DeleteTodo(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}
