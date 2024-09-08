package entities

import "time"

type Todo struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Status string `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}