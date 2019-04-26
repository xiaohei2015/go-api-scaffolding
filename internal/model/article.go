package model

type Article struct {
	Id int64
	Title string
	Content string
	Status int32
	CreateTime int64
	UpdateTime int64
	IsDeleted int32
}