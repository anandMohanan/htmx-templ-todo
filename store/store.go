package store

import (
	"database/sql"
	"fmt"
	"go_fullstack/types"
)

type Storage struct {
	db *sql.DB
}

type Store interface {
	CreateTodo(todo *types.Todo) (*types.Todo, error)
	DeleteTodo(id int) error
	GetTodo(id int) (*types.Todo, error)
	GetTodos() ([]*types.Todo, error)
	SearchTodos(search string) ([]*types.Todo, error)
}

func NewStore(db *sql.DB) *Storage {
	return &Storage{
		db: db,
	}
}

func (s *Storage) CreateTodo(todo *types.Todo) (*types.Todo, error) {
	var id int
	err := s.db.QueryRow("INSERT INTO todo (title, completed) VALUES ($1, $2) RETURNING id", todo.Title, todo.Completed).Scan(&id)
	if err != nil {
		return nil, err
	}
	todo.Id = id
	return todo, nil
}

func (s *Storage) DeleteTodo(id int) error {
	_, err := s.db.Exec("DELETE FROM todo WHERE id = $1", id)
	return err
}

func (s *Storage) GetTodo(id int) (*types.Todo, error) {
	row := s.db.QueryRow("SELECT id, title, completed FROM todo WHERE id = $1", id)
	var todo types.Todo
	err := row.Scan(&todo.Id, &todo.Title, &todo.Completed)
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (s *Storage) GetTodos() ([]types.Todo, error) {
	rows, err := s.db.Query("SELECT id, title, completed FROM todo")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	todos := make([]types.Todo, 0)
	for rows.Next() {
		todo, err := scanTodo(rows)
		if err != nil {
			return nil, err
		}
		fmt.Println(todo)
		todos = append(todos, todo)
	}
	fmt.Println(todos, "todos")
	return todos, nil
}

func (s *Storage) SearchTodos(search string) ([]types.Todo, error) {
	fmt.Println(search, "search")
	rows, err := s.db.Query("SELECT id, title, completed FROM todo WHERE title LIKE $1", "%"+search+"%")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()
	todos := make([]types.Todo, 0)
	for rows.Next() {
		fmt.Println("rows")
		todo, err := scanTodo(rows)
		if err != nil {
			return nil, err
		}
		fmt.Println(todo)
		todos = append(todos, todo)
	}
	return todos, nil
}

func scanTodo(row *sql.Rows) (types.Todo, error) {
	var todo types.Todo
	err := row.Scan(&todo.Id, &todo.Title, &todo.Completed)
	return todo, err
}
