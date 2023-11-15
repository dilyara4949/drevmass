package usecase

import (
	"context"
	"time"

	"github.com/dilyara4949/drevmass/internal/domain"
)



type productUsecase struct {
	productRepository domain.ProductRepository
	contextTimeout time.Duration
}

func NewProductUsecase(productRepository domain.ProductRepository, timeout time.Duration) domain.ProductUsecase {
	return &productUsecase{
		productRepository: productRepository,
		contextTimeout: timeout,
	}
}

func (p *productUsecase) Create(c context.Context, product domain.Product) (domain.Product, error) {
	_, cancel := context.WithTimeout(c, p.contextTimeout)
	defer cancel()
	return p.productRepository.Create(c, product)
}

func (p *productUsecase) GetByID(c context.Context, id string) (domain.Product, error) {
	_, cancel := context.WithTimeout(c, p.contextTimeout)
	defer cancel()
	return p.productRepository.GetByID(c, id)
}

func (p *productUsecase) GetAll(c context.Context) ([]domain.Product, error) {
	_, cancel := context.WithTimeout(c, p.contextTimeout)
	defer cancel()
	return p.productRepository.GetAll(c)
}

func (p *productUsecase) Update(c context.Context, product domain.Product) (domain.Product, error) {
	_, cancel := context.WithTimeout(c, p.contextTimeout)
	defer cancel()
	return p.productRepository.Update(c, product)
}

func (p *productUsecase) Delete(c context.Context, id string) ( error) {
	_, cancel := context.WithTimeout(c, p.contextTimeout)
	defer cancel()
	return p.productRepository.Delete(c, id)
}

func (p *productUsecase) ChangeOrder(ctx context.Context, a uint, b uint) error {
	_, cancel := context.WithTimeout(ctx, p.contextTimeout)
	defer cancel()
	return p.productRepository.ChangeOrder(ctx, a, b)
}


