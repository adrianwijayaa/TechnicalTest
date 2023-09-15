package repositories

import (
	"TechnicalTest/models"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

type ToDoListRepository interface {
	CreateToDoList(list models.ToDoList) (models.ToDoList, error)
	FindToDoList() ([]models.ToDoList, error)
	GetToDoListById(Id int) (models.ToDoList, error)
	UpdateToDoListById(list models.ToDoList, Id int) (models.ToDoList, error)
	DeleteToDoList(Id int) (models.ToDoList, error)
}

func RepositoryToDoLists(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateToDoList(list models.ToDoList) (models.ToDoList, error) {
	err := r.db.Create(&list).Error

	return list, err
}

func (r *repository) FindToDoList() ([]models.ToDoList, error) {
	var lists []models.ToDoList
	err := r.db.Find(&lists).Error

	return lists, err
}

func (r *repository) GetToDoListById(Id int) (models.ToDoList, error) {
	var list models.ToDoList
	err := r.db.First(&list, Id).Error

	return list, err
}

func (r *repository) UpdateToDoListById(list models.ToDoList, Id int) (models.ToDoList, error) {
	var listUpdate models.ToDoList
	err := r.db.First(&listUpdate, Id).Error
	if err != nil {
		return listUpdate, err
	}

	errV := r.db.Model(&listUpdate).Updates(list).Error
	if errV != nil {
		return listUpdate, errV
	}

	return listUpdate, nil
}

func (r *repository) DeleteToDoList(Id int) (models.ToDoList, error) {
	var list models.ToDoList
	err := r.db.Where("id = ?", Id).Delete(&list).Error

	return list, err
}
