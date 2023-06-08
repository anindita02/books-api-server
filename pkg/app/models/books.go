package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Id       int    `json:"id" gorm:"primary_key"`
	Name     string `json:"name"`
	Author   string `json:"author"`
	Year     int    `json:"year"`
	Language string `json:"language"`
}
