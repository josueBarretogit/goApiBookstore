package controllers

type IController interface {
	Create()
	Update()
	Delete()
	GetAll()
}
