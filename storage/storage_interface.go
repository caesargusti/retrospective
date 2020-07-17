package storage

import "github.com/caesargusti/todo2/model"

type Storage interface{
	//func create menerima object Todo dari model
	Create(obj model.Customers) error
	Detail(id int) (model.Customers, error)
	List()([]model.Customers, error)
	Update(obj model.Customers) error
	Delete(id int) error
}