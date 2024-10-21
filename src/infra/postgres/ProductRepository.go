package postgres

import (
	"log"
	"padaria/src/core/domain"
	"padaria/src/core/interfaces/repository"
	"padaria/src/infra/postgres/dto"
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

func (repo ProductRepository) SelectProducts() ([]domain.Product, error) {
	conn, err := repo.getConnection() //verificar o retorno de getConnection
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer repo.closeConnection(conn)

	query := `SELECT 
				ID AS PRODUCT_ID, 
				NAME AS PRODUCT_NAME,
				CODE AS PRODUCT_CODE,
				PRICE AS PRODUCT_PRICE,
				EXPIRATION_DATE AS PRODUCT_EXPIRATION_DATE
			  FROM PRODUCT`
	var productDTOList []dto.Product
	err = conn.Select(&productDTOList, query)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	var products []domain.Product
	for _, productDTO := range productDTOList {
		products = append(products, *productDTO.ToDomain())
	}

	return products, nil
}

func (repo ProductRepository) UpdateProduct(product domain.Product) error {
	conn, err := repo.getConnection()
	if err != nil {
		log.Print(err)
		return err
	}

	defer repo.closeConnection(conn)
	query := `update product set 
				name            = $1,
				code            = $2,
				price           = $3,
				expiration_date = $4
			  where id = $5`

	_, err = conn.Exec(
		query,
		product.Name(),
		product.Code(),
		product.Price(),
		product.ExpirationDate(),
		product.Id(),
	)

	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

func NewProductRepository(manager connectorManager) *ProductRepository {
	return &ProductRepository{connectorManager: manager}
}
