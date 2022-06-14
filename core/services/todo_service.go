package services

import (
	"github.com/koralbit/todo-app-api/core/entities"
	"github.com/koralbit/todo-app-api/core/repository"
)

type TodoService interface {
	GetTodoLists() ([]entities.TodoList, error)
	GetTodoListById(id uint) (*entities.TodoList, error)
	InserTodoList(list entities.TodoList) (*entities.TodoList, error)
	UpdateTodoList(id uint, list entities.TodoList) (*entities.TodoList, error)
	DeleteTodoList(id uint) (*uint, error)

	GetAllTodoItem(todoListId uint) ([]entities.TodoItem, error)
	GetTodoItem(todoListId uint, id uint) (*entities.TodoItem, error)
	CreateTodoItem(todoListId uint, item entities.TodoItem) (*entities.TodoItem, error)
	UpdateTodoItem(todoListId uint, id uint, list entities.TodoItem) (*entities.TodoItem, error)
	DeleteTodoItem(todoListId uint, id uint) (*uint, error)
}
type todoService struct {
	repo repository.TodoRepository
}

func NewTodoService(todoRepository repository.TodoRepository) TodoService {
	return &todoService{
		repo: todoRepository,
	}
}
func (s todoService) GetTodoLists() ([]entities.TodoList, error) {
	return s.repo.GetAllTodoList()
}

func (s todoService) GetTodoListById(id uint) (*entities.TodoList, error) {
	return s.repo.GetTodoList(id)
}

func (s todoService) InserTodoList(list entities.TodoList) (*entities.TodoList, error) {
	return s.repo.CreateTodoList(list)
}

func (s todoService) UpdateTodoList(id uint, list entities.TodoList) (*entities.TodoList, error) {
	l, err := s.repo.GetTodoList(id)
	if err != nil {
		return nil, err
	}
	if l == nil {
		return nil, nil
	}
	l.Description = list.Description
	l.Name = list.Name
	return s.repo.UpdateTodoList(id, *l)
}

func (s todoService) DeleteTodoList(id uint) (*uint, error) {
	l, err := s.repo.GetTodoList(id)
	if err != nil {
		return nil, err
	}
	if l == nil {
		return nil, nil
	}
	err = s.repo.DeleteTodoList(id)
	if err != nil {
		return nil, err
	}
	return &id, nil
}

func (s todoService) GetAllTodoItem(todoListId uint) ([]entities.TodoItem, error) {
	return s.repo.GetAllTodoItem(todoListId)
}

func (s todoService) GetTodoItem(todoListId uint, id uint) (*entities.TodoItem, error) {
	return s.repo.GetTodoItem(todoListId, id)
}

func (s todoService) CreateTodoItem(todoListId uint, item entities.TodoItem) (*entities.TodoItem, error) {
	return s.repo.CreateTodoItem(todoListId, item)
}

func (s todoService) UpdateTodoItem(todoListId uint, id uint, item entities.TodoItem) (*entities.TodoItem, error) {
	l, err := s.repo.GetTodoItem(todoListId, id)
	if err != nil {
		return nil, err
	}
	if l == nil {
		return nil, nil
	}
	l.Done = item.Done
	return s.repo.UpdateTodoItem(todoListId, id, *l)
}

func (s todoService) DeleteTodoItem(todoListId uint, id uint) (*uint, error) {
	l, err := s.repo.GetTodoItem(todoListId, id)
	if err != nil {
		return nil, err
	}
	if l == nil {
		return nil, nil
	}
	err = s.repo.DeleteTodoItem(todoListId, id)
	if err != nil {
		return nil, err
	}
	return &id, nil
}
