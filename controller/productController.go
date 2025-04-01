package controller

import (
	"github.com/TiagoNora/GoCRUDV2/schemas"
	"github.com/TiagoNora/GoCRUDV2/service"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
)

type ProductController interface {
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	FindById(c *gin.Context)
	FindAll(c *gin.Context)
}

type productController struct {
	productService service.ProductService
}

// @Summary      Create a new product
// @Description  Create a new product with name, description, price, and stock
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param        product  body      schemas.ProductRequest  true  "Product data"
// @Router       /product [post]
func (p *productController) Create(c *gin.Context) {
	log.Info().Msg("Called create from Product Controller")
	productRequest := schemas.ProductRequest{}
	if err := c.ShouldBindJSON(&productRequest); err != nil {
		sendError(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	product := schemas.Product{
		Name:        productRequest.Name,
		Description: productRequest.Description,
		Price:       productRequest.Price,
		Stock:       productRequest.Stock,
	}

	product, err := p.productService.Create(product)
	if err != nil {
		sendError(c, http.StatusInternalServerError, "Error creating product")
		return
	}
	log.Info().Msg("Product created")
	sendSuccess(c, "Create Product", product)

}

func (p *productController) Update(c *gin.Context) {
	log.Info().Msgf("Called update from Product Controller id: %s", c.Param("id"))
	id := c.Param("id")

	productRequest := schemas.ProductRequest{}
	if err := c.ShouldBindJSON(&productRequest); err != nil {
		sendError(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	productNew := schemas.Product{
		Name:        productRequest.Name,
		Description: productRequest.Description,
		Price:       productRequest.Price,
		Stock:       productRequest.Stock,
	}

	productFound, err := p.productService.FindById(id)
	if err != nil {
		sendError(c, http.StatusInternalServerError, "Error getting product")
		return
	}

	productUpdated, err := p.productService.Update(productFound, productNew)
	if err != nil {
		sendError(c, http.StatusInternalServerError, "Error updating product")
	}

	sendSuccess(c, "Update Product", productUpdated)

}

func (p *productController) Delete(c *gin.Context) {
	log.Info().Msgf("Called delete from Product Controller id: %s", c.Param("id"))
	idParam := c.Param("id")
	product, err := p.productService.Delete(idParam)
	if err != nil {
		sendError(c, http.StatusInternalServerError, "Error deleting product")
		return
	}
	sendSuccess(c, "delete-product", product)
}

func (p *productController) FindById(c *gin.Context) {
	log.Info().Msgf("Called find by id from Product Controller id: %s", c.Param("id"))
	idParam := c.Param("id")
	product, err := p.productService.FindById(idParam)
	if err != nil {
		sendError(c, http.StatusInternalServerError, "Error getting product")
		return
	}
	sendSuccess(c, "get-product", product)
}

func (p *productController) FindAll(c *gin.Context) {
	log.Info().Msg("Called findAll from Product Controller")
	all, err := p.productService.FindAll()
	if err != nil {
		sendError(c, http.StatusInternalServerError, "Error creating product")
		return
	}
	sendSuccess(c, "list-products", all)
}

func NewProductController() ProductController {
	return &productController{
		productService: service.NewProductService(),
	}

}
