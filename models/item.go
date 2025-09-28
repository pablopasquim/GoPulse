package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model         // add ID, CreatedAt, UpdatedAt, DeletedAt
	Name        string `json:"name"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}
