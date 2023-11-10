package product

import (
	// "fmt"
	// "fmt"
	"net/http"
	"strconv"

	// "time"

	"github.com/dilyara4949/drevmass/internal/domain"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	ProductUsecase domain.ProductUsecase
}


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

func (p *ProductController) GetAll(c *gin.Context) {
	products, err := p.ProductUsecase.GetAll(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}


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


func (p *ProductController) Update(c *gin.Context) {
	ProductID := strToUint( c.Param("id"))
	

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



func (p *ProductController) Delete(c *gin.Context) {
	ProductID := c.Param("id")

	err := p.ProductUsecase.Delete(c, ProductID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Product deleted"})
}

func strToUint(s string) uint {
    i, _ := strconv.Atoi(s)
    return uint(i)
}