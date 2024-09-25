package request

//Arquivo de solicitação de produto

import (
	"padaria/src/core/domain"
	"time"
)

type Product struct {
	//uriliza de bandeiras em go (associar os atribuots de um json aos atributos de produto)
	Name  string  `json:"name"`
	Code  string  `json:"code"`
	Price float32 `json:"price"`
}

func (dto Product) ToDomain() *domain.Product {
	return domain.NewProduct(0, dto.Name, dto.Price, dto.Code, time.Time{})
}