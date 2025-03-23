package response

import (
	"padaria/src/core/domain"
	"time"
)

type Product struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Code           string    `json:"code"`
	Price          float32   `json:"price"`
	ExpirationDate time.Time `json:"expiration_date"`
}

func NewProduct(product domain.Product) *Product {
	return &Product{
		ID:             product.Id(),
		Name:           product.Name(),
		Code:           product.Code(),
		Price:          product.Price(),
		ExpirationDate: product.ExpirationDate(),
	}
}
