package product

import (
	"net/http"
	"strconv"

	_ "github.com/dilyara4949/drevmass/docs"
	"github.com/dilyara4949/drevmass/internal/domain"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	ProductUsecase domain.ProductUsecase
}

// @Summary GetProduct
// @Security ApiKeyAuth
// @Tags product
// @Description get product by id
// @ID get-product
// @Produce  json
// @Param        id   path      int  true  "Product ID"
// @Success 200 {object} domain.Product
// @Failure 500 {object} domain.ErrorResponse
// @Failure default {object} domain.ErrorResponse
// @Router /products/{id} [get]
func (p *ProductController) GetOne(c *gin.Context) {
	productID := c.Param("id")
	// fmt.Println(productID)
	product, err := p.ProductUsecase.GetByID(c, productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}

// @Summary GetProducts
// @Security ApiKeyAuth
// @Tags product
// @Description get all products
// @ID get-products
// @Produce  json
// @Success 200 {object} []domain.Product
// @Failure 500 {object} domain.ErrorResponse
// @Failure default {object} domain.ErrorResponse
// @Router /products [get]
func (p *ProductController) GetAll(c *gin.Context) {
	products, err := p.ProductUsecase.GetAll(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

// @Summary GetProducts price up
// @Security ApiKeyAuth
// @Tags product
// @Description get all products price up
// @ID get-products-price-up
// @Produce  json
// @Success 200 {object} []domain.Product
// @Failure 500 {object} domain.ErrorResponse
// @Failure default {object} domain.ErrorResponse
// @Router /products/priceup [get]
func (p *ProductController) GetAllPriceUp(c *gin.Context) {
	products, err := p.ProductUsecase.GetAllPriceUp(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

// @Summary GetProducts price down
// @Security ApiKeyAuth
// @Tags product
// @Description get all products price down
// @ID get-products-price-down
// @Produce  json
// @Success 200 {object} []domain.Product
// @Failure 500 {object} domain.ErrorResponse
// @Failure default {object} domain.ErrorResponse
// @Router /products/pricedown [get]
func (p *ProductController) GetAllPricedown(c *gin.Context) {
	products, err := p.ProductUsecase.GetAllPriceDown(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

// @Summary CreateProduct
// @Security ApiKeyAuth
// @Tags product
// @Description create product
// @ID create-product
// @Produce  json
// @Accept   json
// @Param	product	body  domain.Product	true	"Create product"
// @Success 200 {object} domain.Product
// @Failure 500 {object} domain.ErrorResponse
// @Failure default {object} domain.ErrorResponse
// @Router /products [post]
func (p *ProductController) Create(c *gin.Context) {

	var request domain.Product

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	product, err := p.ProductUsecase.Create(c, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

// @Summary UpdateProduct
// @Security ApiKeyAuth
// @Tags product
// @Description update product
// @ID update-product
// @Produce  json
// @Accept       json
// @Param        id   path      int  true  "Product ID"
//@Param	product	body	domain.Product	true	"update product"
// @Success 200 {object} domain.Product
// @Failure 500 {object} domain.ErrorResponse
// @Failure default {object} domain.ErrorResponse
// @Router /products/{id} [post]
func (p *ProductController) Update(c *gin.Context) {
	ProductID := strToUint(c.Param("id"))

	var request domain.Product

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		// c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "3"})
		return
	}
	request.ID = ProductID

	_, err = p.ProductUsecase.Update(c, request)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, request)
}

// @Summary DeleteProduct
// @Security ApiKeyAuth
// @Tags product
// @Description Delete product
// @ID Delete-product
// @Produce  json
// @Accept       json
// @Param    id   path      int  true  "Product ID"
// @Success 200 {object} domain.SuccessResponse
// @Failure 500 {object} domain.ErrorResponse
// @Failure default {object} domain.ErrorResponse
// @Router /products/{id} [delete]
func (p *ProductController) Delete(c *gin.Context) {
	ProductID := c.Param("id")

	err := p.ProductUsecase.Delete(c, ProductID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Product deleted"})
}


// @Summary ChangeOrderOfProducts
// @Security ApiKeyAuth
// @Tags product
// @Description change order of products
// @ID change-order-product
// @Produce  json
// @Accept       json
// @Param    a   path      int  true  "First Product"
// @Param    b   path      int  true  "Second Product"
// @Success 200 {object} domain.SuccessResponse
// @Failure 500 {object} domain.ErrorResponse
// @Failure default {object} domain.ErrorResponse
// @Router /products/change/{a}/{b} [post]
func (p *ProductController) ChangeOrder(c *gin.Context) {
	a := strToUint(c.Param("a"))
	b := strToUint(c.Param("b"))
	
	err := p.ProductUsecase.ChangeOrder(c, a, b)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Order changed"})
}



func strToUint(s string) uint {
	i, _ := strconv.Atoi(s)
	return uint(i)
}


