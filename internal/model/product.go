package domain

import "time"

type Product struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float32   `json:"price"`
	Quantity    int64     `json:"quantity"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
}
