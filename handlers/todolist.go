package handlers

import (
	dto "TechnicalTest/dto/result"
	todolistdto "TechnicalTest/dto/todolist"
	"TechnicalTest/models"
	"TechnicalTest/repositories"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
)

type handlerToDoLists struct {
	ToDoListRepository repositories.ToDoListRepository
}

func HandlerToDoLists(ToDoListRepository repositories.ToDoListRepository) *handlerToDoLists {
	return &handlerToDoLists{ToDoListRepository}
}

func (h *handlerToDoLists) CreateToDoList(c echo.Context) error {
	request := todolistdto.CreateToDoListRequest{
		Title:       c.FormValue("title"),
		Description: c.FormValue("description"),
	}

	var err error

	validation := validator.New()
	err = validation.Struct(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	toDoLists := models.ToDoList{
		Title:       request.Title,
		Description: request.Description,
	}

	dataToDoLists, err := h.ToDoListRepository.CreateToDoList(toDoLists)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: dataToDoLists})
}

func (h *handlerToDoLists) FindToDoList(c echo.Context) error {
	lists, err := h.ToDoListRepository.FindToDoList()
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: lists})
}

func (h *handlerToDoLists) GetToDoListById(c echo.Context) error {
	Id, _ := strconv.Atoi(c.Param("id"))

	list, err := h.ToDoListRepository.GetToDoListById(Id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: list})
}

func (h *handlerToDoLists) UpdateToDoListById(c echo.Context) error {
	Id, _ := strconv.Atoi(c.Param("id"))
	list, err := h.ToDoListRepository.GetToDoListById(Id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	request := todolistdto.UpdateToDoListRequest{
		Title:       c.FormValue("title"),
		Description: c.FormValue("description"),
	}

	if request.Title != "" {
		list.Title = request.Title
	}

	if request.Description != "" {
		list.Description = request.Description
	}

	newList, err := h.ToDoListRepository.UpdateToDoListById(list, Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: newList})

}

func (h *handlerToDoLists) DeleteToDoList(c echo.Context) error {
	Id, _ := strconv.Atoi(c.Param("id"))

	list, err := h.ToDoListRepository.DeleteToDoList(Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: list})
}
