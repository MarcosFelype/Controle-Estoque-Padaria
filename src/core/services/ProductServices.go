package services

import (
	"log"
	"padaria/src/core/domain"
	"padaria/src/core/interfaces/primary"
	"padaria/src/core/interfaces/repository"
)

var _ primary.ProductManager = (*ProcuctServices)(nil) //força o adaptador a utilizar a porta primária

type ProcuctServices struct {
	productRepository repository.ProductLoader
}

func (service ProcuctServices) RegisterProduct(product domain.Product) (int, error) { //implementação do método da porta primária
	productID, err := service.productRepository.InsertProduct(product) //chamando a porta secundária
	//a inserção do produto efetivamente será abordada a seguir
	if err != nil {
		log.Print(err)
		return -1, err
	}

	return productID, nil
}
