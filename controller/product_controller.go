package controller

import (
	"fmt"
	"go-api/model"
	"go-api/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productController struct {
	productUsecase usecase.ProductUsecase
}

func NewProductController(usecase usecase.ProductUsecase) productController {
	return productController{productUsecase: usecase}
}

func (p *productController) GetProducts(ctx *gin.Context) {
	products, err := p.productUsecase.GetProducts()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, products)
}

func (controller *productController) CreateProduct(ctx *gin.Context) {
	var product model.Product

	err := ctx.BindJSON(&product)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		fmt.Println(err)
		return
	}

	insertedProduct, err := controller.productUsecase.CreateProduct(product)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertedProduct)
}

func (controller *productController) GetProductById(ctx *gin.Context) {
	id := ctx.Param("id")

	productId, validationError := validateProductId(id)
	if validationError != nil {
		ctx.JSON(validationError.Status, validationError.Message)
		return
	}

	product, err := controller.productUsecase.GetProductById(productId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if product == nil {
		response := model.ValidationError{
			Message: "Product not found.",
		}

		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func validateProductId(id string) (int, *model.ValidationError) {
	if id == "" {
		return 0, &model.ValidationError{
			Status:  http.StatusBadRequest,
			Message: "Id cannot be empty",
		}
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		return 0, &model.ValidationError{
			Status:  http.StatusBadRequest,
			Message: "ID must be a number.",
		}
	}

	return productId, nil
}
