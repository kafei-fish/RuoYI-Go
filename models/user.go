package models

type User struct {
	userId uint `gorm:"primaryKey"`
	Name   string
	Email  string
}
