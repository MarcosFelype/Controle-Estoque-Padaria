package handlers

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"padaria/src/core/domain"
	"strings"
	"time"

	//adicionar: serviceMock "padaria/src/core/interfaces/primary/mocks"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"padaria/src/app/api/endpoints/dto/request"
	"testing"
)

const (
	postProductURL = "/api/product/new"
)

func TestPostProduct(t *testing.T) {
	caseManager := postProductCases{}

	t.Run("Test PostProduct when it has a happy path", caseManager.testWhenItHasHappyPath)
	t.Run("Test PostProduct when it cannot process the product", caseManager.testWhenItCannotProcessTheProduct)
	t.Run("Test PostProduct when it receives a connection error", caseManager.testWhenItReceivesAConnectionError)
}

type postProductCases struct{}

func (postProductCases) testWhenItHasHappyPath(t *testing.T) {
	controller := gomock.NewController(t)
	productService := serviceMock.NewMockProductManager(controller)
	productHandlers := NewProductHandlers(productService)
	product := domain.NewProduct(0, "Produto Teste", 0.1, "0000", time.Now().Add(30*time.Second))
	requestBody := fmt.Sprintf(
		`{	
		"name": "%s",
		"code": "%s",
		"price": %f,
	}`,
		product.Name(),
		product.Code(),
		product.Price(),
	)
	expectedID := 1
	productService.EXPETED().RegisterProduct(*product).Return(expectedID, nil).MaxTimes(1)
	body := strings.NewReader(requestBody)
	request := httptest.NewRequest(http.MethodPost, postProductURL, body)
	request.Header.Add(echo.HeaderContentType, "application/json")
	recorder := httptest.NewRecorder()
	server := echo.New()
	context := server.NewContext(request, recorder)
	expectedJSON := fmt.Sprintf(`{"id": %d}`, expectedID)
	expectedStatusCode := http.StatusCreated

	_ = productHandlers.PostProduct(context)

	assert.JSONEq(t, expectedJSON, recorder.Body.String())
	assert.Equal(t, expectedStatusCode, recorder.Code)
}

func (postProductCases) testWhenItCannotProcessTheProduct(t *testing.T) {
	controller := gomock.NewController(t)
	productServices := serviceMock.NewMockProductManager(controller)
	productHandlers := NewProductHandlers(productServices)
	requestBody := `{}`
	body := strings.NewReader(requestBody)
	request := httptest.NewRequest(http.MethodPost, postProductURL, body)
	recorder := httptest.NewRecorder()
	server := echo.New()
	context := server.NewContext(request, recorder)
	expectedMsg := "Algo está incompatível na sua requisição."
	expectedStatusCode := http.StatusBadRequest

	expectedJSON := fmt.Sprintf(`{
		"msg": "%s",
		"statusCode": %i,
	}`,
		expectedMsg,
		expectedStatusCode,
	)

	_ = productHandlers.PostProduct(context)
	assert.JSONEq(t, expectedJSON, recorder.Body.String())
	assert.Equal(t, expectedStatusCode, recorder.Code)
}

func (postProductCases) testWhenItReceivesAConnectionError(t *testing.T) {
	controller := gomock.NewController(t)
	productServices := serviceMock.NewMockProductManager(controller)
	err := errors.New("connection error")
	productService.EXPETED().RegisterProduct(gomock.Any()).Return(-1, err).MaxTimes(1)
	productHandlers := NewProductHandlers(productServices)
	requestBody := `{}`
	body := strings.NewReader(requestBody)
	request := httptest.NewRequest(http.MethodPost, postProductURL, body)
	request.Header.Add(echo.HeaderContentType, "application/json")
	recorder := httptest.NewRecorder()
	server := echo.New()
	context := server.NewContext(request, recorder)
	expectedMsg := "Oops! Parece que o serviço de dados está indisponível."
	expectedStatusCode := http.StatusInternalServerError

	expectedJSON := fmt.Sprintf(`{
		"msg": "%s",
		"statusCode": %i,
	}`,
		expectedMsg,
		expectedStatusCode,
	)

	_ = productHandlers.PostProduct(context)
	assert.JSONEq(t, expectedJSON, recorder.Body.String())
	assert.Equal(t, expectedStatusCode, recorder.Code)
}
