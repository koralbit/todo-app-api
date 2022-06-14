package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/koralbit/todo-app-api/api/models"
	"github.com/koralbit/todo-app-api/core/entities"
	"github.com/koralbit/todo-app-api/core/services"
	"github.com/labstack/echo/v4"
)

type TodoController interface {
	Route(e *echo.Echo)
	GetAllTodoList(c echo.Context) error
	GetTodoList(c echo.Context) error
	CreateTodoList(c echo.Context) error
	UpdateTodoList(c echo.Context) error
	DeleteTodoList(c echo.Context) error

	GetAllTodoItems(c echo.Context) error
	GetTodoItem(c echo.Context) error
	CreateTodoItem(c echo.Context) error
	UpdateTodoItem(ctx echo.Context) error
}

type todoController struct {
	service services.TodoService
}

func NewTodoController(service services.TodoService) TodoController {

	return &todoController{
		service: service,
	}
}

func (t todoController) Route(e *echo.Echo) {
	g := e.Group("list")
	g.GET("", t.GetAllTodoList)
	g.GET("/:id", t.GetTodoList)
	g.POST("", t.CreateTodoList)
	g.POST("/:id", t.UpdateTodoList)
	g.DELETE("/:id", t.DeleteTodoList)

	item := g.Group("/:id/items")
	item.GET("", t.GetAllTodoItems)
	item.GET("/:itemId", t.GetTodoItem)
	item.POST("", t.CreateTodoItem)
	item.POST("/:itemId", t.UpdateTodoItem)
	item.DELETE("/:itemId", t.DeleteTodoList)
}

func (t todoController) GetAllTodoList(ctx echo.Context) error {
	res, err := t.service.GetTodoLists()
	if err != nil {
		return errorMessage(ctx, err.Error(), http.StatusInternalServerError)
	}
	return ctx.JSON(http.StatusOK, res)
}

func (t todoController) GetTodoList(ctx echo.Context) error {
	strId := ctx.Param("id")
	id, err := strconv.ParseInt(strId, 10, 32)
	if err != nil {
		return errorMessage(ctx, err.Error(), http.StatusBadRequest)
	}
	res, err := t.service.GetTodoListById(uint(id))
	if err != nil {
		return errorMessage(ctx, err.Error(), http.StatusInternalServerError)
	}
	if res == nil {
		return errorMessage(ctx, fmt.Sprintf("List %d not found", id), http.StatusNotFound)
	}
	return ctx.JSON(http.StatusOK, res)
}

func (t todoController) CreateTodoList(ctx echo.Context) error {
	var request models.TodoListCreateRequest
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	if err := ctx.Validate(request); err != nil {
		return err
	}
	list := entities.TodoList{
		Name:        request.Name,
		Description: request.Description,
	}
	res, err := t.service.InserTodoList(list)
	if err != nil {
		return errorMessage(ctx, err.Error(), http.StatusInternalServerError)
	}
	return ctx.JSON(http.StatusOK, res)
}

func (t todoController) UpdateTodoList(ctx echo.Context) error {
	strId := ctx.Param("id")
	id, err := strconv.ParseInt(strId, 10, 32)
	if err != nil {
		return errorMessage(ctx, err.Error(), http.StatusBadRequest)
	}
	var request models.TodoListCreateRequest
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	if err := ctx.Validate(request); err != nil {
		return err
	}
	list := entities.TodoList{
		Name:        request.Name,
		Description: request.Description,
	}
	res, err := t.service.UpdateTodoList(uint(id), list)
	if err != nil {
		return errorMessage(ctx, err.Error(), http.StatusInternalServerError)
	}
	if res == nil {
		return errorMessage(ctx, fmt.Sprintf("List %d not found", id), http.StatusNotFound)
	}
	return ctx.JSON(http.StatusOK, res)
}

func (t todoController) DeleteTodoList(ctx echo.Context) error {
	strId := ctx.Param("id")
	id, err := strconv.ParseInt(strId, 10, 32)
	if err != nil {
		return errorMessage(ctx, err.Error(), http.StatusBadRequest)
	}
	res, err := t.service.DeleteTodoList(uint(id))
	if err != nil {
		return errorMessage(ctx, err.Error(), http.StatusInternalServerError)
	}
	if res == nil {
		return errorMessage(ctx, fmt.Sprintf("List %d not found", id), http.StatusNotFound)
	}
	return ctx.JSON(http.StatusOK, echo.Map{
		"Message": fmt.Sprintf("List %d deleted", *res),
	})
}

func errorMessage(ctx echo.Context, message string, code int) error {
	return ctx.JSON(code, echo.Map{
		"error": message,
	})
}

func (t todoController) GetAllTodoItems(ctx echo.Context) error {
	strId := ctx.Param("id")
	id, err := strconv.ParseInt(strId, 10, 32)
	if err != nil {
		return errorMessage(ctx, err.Error(), http.StatusBadRequest)
	}
	res, err := t.service.GetAllTodoItem(uint(id))
	if err != nil {
		return errorMessage(ctx, err.Error(), http.StatusInternalServerError)
	}
	return ctx.JSON(http.StatusOK, res)
}

func (t todoController) GetTodoItem(ctx echo.Context) error {
	strId := ctx.Param("id")
	id, err := strconv.ParseInt(strId, 10, 32)
	if err != nil {
		return errorMessage(ctx, err.Error(), http.StatusBadRequest)
	}
	strItemId := ctx.Param("itemId")
	itemId, err := strconv.ParseInt(strItemId, 10, 32)
	if err != nil {
		return errorMessage(ctx, err.Error(), http.StatusBadRequest)
	}
	res, err := t.service.GetTodoItem(uint(id), uint(itemId))
	if err != nil {
		return errorMessage(ctx, err.Error(), http.StatusInternalServerError)
	}
	if res == nil {
		return errorMessage(ctx, fmt.Sprintf("Item %d for %d list not found", itemId, id), http.StatusNotFound)
	}
	return ctx.JSON(http.StatusOK, res)
}

func (t todoController) CreateTodoItem(ctx echo.Context) error {
	strId := ctx.Param("id")
	id, err := strconv.ParseInt(strId, 10, 32)
	if err != nil {
		return errorMessage(ctx, err.Error(), http.StatusBadRequest)
	}
	var request models.TodoItemCreateRequest
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	if err := ctx.Validate(request); err != nil {
		return err
	}
	list := entities.TodoItem{
		Description: request.Description,
		TodoListId:  uint(id),
	}
	res, err := t.service.CreateTodoItem(uint(id), list)
	if err != nil {
		return errorMessage(ctx, err.Error(), http.StatusInternalServerError)
	}
	return ctx.JSON(http.StatusOK, res)
}

func (t todoController) UpdateTodoItem(ctx echo.Context) error {
	strId := ctx.Param("id")
	id, err := strconv.ParseInt(strId, 10, 32)
	if err != nil {
		return errorMessage(ctx, err.Error(), http.StatusBadRequest)
	}
	strItemId := ctx.Param("itemId")
	itemId, err := strconv.ParseInt(strItemId, 10, 32)
	if err != nil {
		return errorMessage(ctx, err.Error(), http.StatusBadRequest)
	}
	var request models.TodoIteUpdateRequest
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	if err := ctx.Validate(request); err != nil {
		return err
	}
	item := entities.TodoItem{
		Done: request.Done,
	}
	res, err := t.service.UpdateTodoItem(uint(id), uint(itemId), item)
	if err != nil {
		return errorMessage(ctx, err.Error(), http.StatusInternalServerError)
	}
	if res == nil {
		return errorMessage(ctx, fmt.Sprintf("Item %d for %d list not found", itemId, id), http.StatusNotFound)
	}
	return ctx.JSON(http.StatusOK, res)
}
