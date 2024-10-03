package dto

import (
	"padaria/src/core/domain"
	"time"
)

type Product struct {
	ID             int       `db:"product-id"` //obs: a bandeira db é para receber um dado do banco de dados
	Name           string    `db:"product-name"`
	Code           string    `db:"product-code"`
	Price          float32   `db:"product-price"`
	ExpirationDate time.Time `db:"product-expiration_date"`
}

func (dto Product) ToDomain() *domain.Product {
	return domain.NewProduct(dto.ID, dto.Name, dto.Price, dto.Code, dto.ExpirationDate)
}
