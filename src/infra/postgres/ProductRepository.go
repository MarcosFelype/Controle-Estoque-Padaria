package postgres

import (
	"log"
	"padaria/src/core/domain"
	"padaria/src/core/interfaces/repository"
)

var _ repository.ProductLoader = (*ProductRepository)(nil)

type ProductRepository struct {
	connectorManager //"herança" da interface connection
}

func (repo ProductRepository) InsertProduct(product domain.Product) (int, error) {
	conn, err := repo.getConnection() //verificar o retorno de getConnection
	if err != nil {
		log.Println(err)
		return -1, err
	}
	defer repo.closeConnection(conn) //o defer é executado depois do código antecedente

	query := `
		INSERT INTO PRODUCT(NAME, CODE, PRICE, EXPIRATION_DATE) 
		VALUES($1, $2, $3, $4) RETURNING ID;
	`

	var productID int
	err = conn.Get(&productID, query, product.Name(), product.Code(), product.Price(), product.ExpirationDate())

	if err != nil {
		log.Println(err)
		return -1, err
	}

	return productID, nil
}

func NewProductRepository(manager connectorManager) *ProductRepository {
	return &ProductRepository{connectorManager: manager}
}
