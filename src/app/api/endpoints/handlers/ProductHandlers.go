// arquivo para disponibilizar a função RegisterProduct, de primary
package handlers

import (
	"net/http"
	"padaria/src/app/api/endpoints/dto/request"
	"padaria/src/app/api/endpoints/dto/response"
	"padaria/src/core/interfaces/primary"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductHandlers struct {
	productService primary.ProductManager
}

// Função para disponibilizar a função RegisterProduct, de primary
func (handler ProductHandlers) PostProduct(c echo.Context) error { //o pacote echo apresenta o resultado de uma requisição

	var product request.Product
	if err := c.Bind(&product); err != nil { //fazendo a atribuição e testando a condição ao mesmo tempo
		return c.JSON(http.StatusBadRequest, response.NewError(
			"Algo está incompatível na sua requisição.",
			http.StatusBadRequest,
		))
	}

	productId, registerErr := handler.productService.RegisterProduct(*product.ToDomain())
	if registerErr != nil {
		return c.JSON(http.StatusBadRequest, response.NewError(
			"Oops! Parece que o serviço de dados está indisponível.",
			http.StatusBadRequest,
		))
	}

	return c.JSON(http.StatusCreated, &response.Created{ID: productId})
}

func (handler ProductHandlers) GetProducts(c echo.Context) error {
	products, err := handler.productService.ListProducts()

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.NewError(
			"Oops! Parece que o serviço de dados está indisponível.",
			http.StatusBadRequest,
		))
	}

	var productDTOList []response.Product
	for _, product := range products {
		productDTOList = append(productDTOList, *response.NewProduct(product))
	}
	return c.JSON(http.StatusOK, productDTOList)
}

func (handler ProductHandlers) PutProduct(c echo.Context) error {
	productId, err := strconv.Atoi(c.Param("productId"))

	var productDTO request.Product //recebe um dto, formato de JSON

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.NewError(
			"O id desse produto não pôde ser processado.",
			http.StatusBadRequest,
		))
	}

	product := productDTO.ToDomainWithId(productId)
	//cria o produto propriamente dito (domain de produto)

	err = handler.productService.EditProduct(*product)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.NewError(
			"Oops! Parece que o serviço de dados está indisponível.",
			http.StatusBadRequest,
		))
	}

	return c.NoContent(http.StatusNoContent)
}

func NewProductHandlers(productService primary.ProductManager) *ProductHandlers {
	return &ProductHandlers{
		productService: productService}
}
