package repository

import (
	"github.com/koralbit/todo-app-api/core/entities"
	"gorm.io/gorm"
)

type TodoRepository interface {
	GetAllTodoList() ([]entities.TodoList, error)
	GetTodoList(id uint) (*entities.TodoList, error)
	CreateTodoList(list entities.TodoList) (*entities.TodoList, error)
	UpdateTodoList(id uint, list entities.TodoList) (*entities.TodoList, error)
	DeleteTodoList(id uint) error

	GetAllTodoItem(todoListId uint) ([]entities.TodoItem, error)
	GetTodoItem(todoListId uint, id uint) (*entities.TodoItem, error)
	CreateTodoItem(todoListId uint, item entities.TodoItem) (*entities.TodoItem, error)
	UpdateTodoItem(todoListId uint, id uint, list entities.TodoItem) (*entities.TodoItem, error)
	DeleteTodoItem(todoListId uint, id uint) error
}

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	err := db.AutoMigrate(&entities.TodoItem{})
	if err != nil {
		panic(err.Error())
	}
	err = db.AutoMigrate(&entities.TodoList{})
	if err != nil {
		panic(err.Error())
	}
	return &todoRepository{
		db: db,
	}
}
func (r todoRepository) GetAllTodoList() ([]entities.TodoList, error) {
	var todoLists []entities.TodoList
	err := r.db.Order("id asc").Find(&todoLists).Error
	if err != nil {
		return nil, err
	}
	return todoLists, nil
}

func (r todoRepository) GetTodoList(id uint) (*entities.TodoList, error) {
	todoList := entities.TodoList{Id: id}
	res := r.db.Limit(1).Find(&todoList)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}
	return &todoList, nil
}

func (r todoRepository) CreateTodoList(list entities.TodoList) (*entities.TodoList, error) {
	err := r.db.Create(&list).Error
	if err != nil {
		return nil, err
	}
	return &list, nil
}

func (r todoRepository) UpdateTodoList(id uint, list entities.TodoList) (*entities.TodoList, error) {
	err := r.db.Model(&list).Omit("created_at").Updates(entities.TodoList{Name: list.Name, Description: list.Description}).Error
	if err != nil {
		return nil, err
	}
	return &list, nil
}

func (r todoRepository) DeleteTodoList(id uint) error {
	todoList := entities.TodoList{Id: id}
	err := r.db.Delete(&todoList).Error
	if err != nil {
		return err
	}
	return nil
}

func (r todoRepository) GetAllTodoItem(todoListId uint) ([]entities.TodoItem, error) {
	var todoItems []entities.TodoItem
	err := r.db.Order("id asc").Model(entities.TodoItem{TodoListId: todoListId}).Find(&todoItems).Error
	if err != nil {
		return nil, err
	}
	return todoItems, nil
}

func (r todoRepository) GetTodoItem(todoListId uint, id uint) (*entities.TodoItem, error) {
	todoItem := entities.TodoItem{Id: id, TodoListId: todoListId}
	res := r.db.Limit(1).Find(&todoItem)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, nil
	}
	return &todoItem, nil
}

func (r todoRepository) CreateTodoItem(todoListId uint, item entities.TodoItem) (*entities.TodoItem, error) {
	item.TodoListId = todoListId
	err := r.db.Create(&item).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r todoRepository) UpdateTodoItem(todoListId uint, id uint, list entities.TodoItem) (*entities.TodoItem, error) {
	err := r.db.Model(&list).Omit("created_at").Update("complete", list.Done).Error
	if err != nil {
		return nil, err
	}
	return &list, nil
}

func (r todoRepository) DeleteTodoItem(todoListId uint, id uint) error {
	todoItem := entities.TodoItem{Id: id, TodoListId: todoListId}
	err := r.db.Delete(&todoItem).Error
	if err != nil {
		return err
	}
	return nil
}
