package product

import (
	"time"

	"github.com/dilyara4949/drevmass/api/controller/product"
	"github.com/dilyara4949/drevmass/internal/repository"
	"github.com/dilyara4949/drevmass/internal/usecase"
	"github.com/dilyara4949/drevmass/pkg"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

func NewProductRouter(env *pkg.Env, timeout time.Duration, db *pgxpool.Pool, group *gin.RouterGroup) {
	r := repository.NewProductRepository(db)
	c := &product.ProductController{
		ProductUsecase: usecase.NewProductUsecase(r, timeout),
	}
	group.POST("/products", c.Create)
	group.GET("/products", c.GetAll)
	group.GET("/products/:id", c.GetOne)
	group.POST("/products/:id", c.Update)
	group.DELETE("/products/:id", c.Delete)
	group.POST("/products/change/:a/:b", c.ChangeOrder)
	group.GET("/products/pricedown", c.GetAllPricedown)
	group.GET("/products/priceup", c.GetAllPriceUp)
}
