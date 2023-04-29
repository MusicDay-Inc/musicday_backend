package core

type Song struct {
	Id     int    `json:"-"`
	Name   string `json:"name" binding:"required"`
	Author string `json:"author" binding:"required"`
}

func NewSong(id int, name string, author string) Song {
	return Song{Id: id, Name: name, Author: author}
}
