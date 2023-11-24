package domain

import "context"

type Product struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	ImageSrc    string  `json:"image_src"`
	VideoSrc    string  `json:"video_src"`
	Price       uint    `json:"price"`
	Weight      float32 `json:"weight"`
	Length      string  `json:"length"`
	Height      string  `json:"height"`
	Depth       string  `json:"depth"`
	Icon        string  `json:"icon"`
	Status      uint    `json:"status"`
}

type ProductRepository interface {
	Create(ctx context.Context, product Product) (Product, error)
	GetByID(ctx context.Context, id string) (Product, error)
	GetAll(ctx context.Context) ([]Product, error)
	Update(ctx context.Context, product Product) (Product, error)
	Delete(ctx context.Context, id string) error
	ChangeOrder(ctx context.Context, a uint, b uint) error
	GetAllPriceUp(ctx context.Context) ([]Product, error)
	GetAllPriceDown(ctx context.Context) ([]Product, error)
}

type ProductUsecase interface {
	Create(c context.Context, product Product) (Product, error)
	GetByID(c context.Context, id string) (Product, error)
	Update(c context.Context, product Product) (Product, error)
	Delete(c context.Context, id string) error
	GetAll(c context.Context) ([]Product, error)
	ChangeOrder(ctx context.Context, a uint, b uint) error
	GetAllPriceUp(ctx context.Context) ([]Product, error)
	GetAllPriceDown(ctx context.Context) ([]Product, error)
}
