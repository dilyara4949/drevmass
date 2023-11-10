package domain

import "context"

type Product struct {
	ID          uint
	Name        string
	Title       string
	Description string
	ImageSrc    string
	VideoSrc    string
	Price       uint
	Weight      float32
	Length      string
	Height      string
	Depth       string
	Icon        string
	Status      uint
}


type ProductRepository interface {
	Create(ctx context.Context, product Product) (Product, error)
	GetByID(ctx context.Context, id string) (Product, error) 
	GetAll(ctx context.Context, ) ([]Product, error) 
	Update(ctx context.Context,  product Product) (Product, error) 
	Delete(ctx context.Context, id string) ( error)
}

type ProductUsecase interface {
	Create(c context.Context, product Product) (Product, error)
	GetByID(c context.Context, id string) (Product, error)
	Update(c context.Context, product Product) (Product, error)
	Delete(c context.Context, id string) ( error)
	GetAll(c context.Context) ([]Product, error)
}