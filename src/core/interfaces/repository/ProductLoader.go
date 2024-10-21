package repository

import "padaria/src/core/domain"

//Porta Secundária
type ProductLoader interface {
	InsertProduct(product domain.Product) (int, error)
	SelectProducts() ([]domain.Product, error)
}
