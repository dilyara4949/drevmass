package repository

import (
	"context"
	"fmt"
	// "fmt"

	"github.com/dilyara4949/drevmass/internal/domain"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

type productRepository struct {
	db *pgxpool.Pool
}

func NewProductRepository(db *pgxpool.Pool) domain.ProductRepository {
	return &productRepository{db: db}
}

func (p *productRepository) Create(ctx context.Context, product domain.Product) (domain.Product, error) {
	ord := 0
	if err := p.db.QueryRow(ctx, findMaxFromProductOrder).Scan(&ord); err != nil {
		return domain.Product{}, errors.Wrap(err, "Scan")
	}
	ord += 1
	if err := p.db.QueryRow(ctx, createProductQuery,
		&product.Name,
		&product.Title,
		&product.Description,
		&product.ImageSrc,
		&product.VideoSrc,
		&product.Price,
		&product.Weight,
		&product.Length,
		&product.Height,
		&product.Depth,
		&product.Icon,
		&product.Status,
		ord,
	).Scan(&product.ID); err != nil {
		return domain.Product{}, errors.Wrap(err, "Scan9")
	}

	return product, nil
}

func (p *productRepository) GetByID(ctx context.Context, id string) (domain.Product, error) {
	var product domain.Product
	if err := p.db.QueryRow(ctx, getProductQuery, id).Scan(
		&product.ID,
		&product.Name,
		&product.Title,
		&product.Description,
		&product.ImageSrc,
		&product.VideoSrc,
		&product.Price,
		&product.Weight,
		&product.Length,
		&product.Height,
		&product.Depth,
		&product.Icon,
		&product.Status,
	); err != nil {

		return domain.Product{}, errors.Wrap(err, "Scan")
	}

	return product, nil
}

func (p *productRepository) Update(ctx context.Context, product domain.Product) (domain.Product, error) {

	if err := p.db.QueryRow(ctx,
		updateProductQuery,
		&product.Name,
		&product.Title,
		&product.Description,
		&product.ImageSrc,
		&product.VideoSrc,
		&product.Price,
		&product.Weight,
		&product.Length,
		&product.Height,
		&product.Depth,
		&product.Icon,
		&product.Status,
		&product.ID,
	).Scan(&product.ID); err != nil {
		return domain.Product{}, errors.Wrap(err, "(Product with given id does not exist) Scan")
	}

	return product, nil
}

func (p *productRepository) Delete(ctx context.Context, id string) error {
	var ord *int
	if err := p.db.QueryRow(ctx, deleteProductQuery, id).Scan(&ord); ord == nil {
		return errors.Wrap(err, "(Product with given id does not exist) Scan")
	}

	_, err := p.db.Exec(ctx, cleanProductOrder, &ord)
	if err != nil {
		return err
	}

	return nil
}

func (p *productRepository) GetAll(ctx context.Context) ([]domain.Product, error) {

	products := []domain.Product{}
	rows, err := p.db.Query(context.Background(), getAllProductQuery)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		r, err := rows.Values()

		fmt.Println(r)
		if err != nil {
			return nil, err
		}
		product := domain.Product{}
		product.ID = uint(r[0].(int32))
		product.Name = r[1].(string)
		product.Title = r[2].(string)
		product.Description = r[3].(string)
		product.ImageSrc = r[4].(string)
		product.VideoSrc = r[5].(string)
		product.Price = uint(r[6].(float32))
		product.Weight = r[7].(float32)
		product.Length = r[8].(string)
		product.Height = r[9].(string)
		product.Depth = r[10].(string)
		product.Icon = r[11].(string)
		product.Status = uint(r[12].(int32))

		products = append(products, product)
	}

	return products, nil
}

func (p *productRepository) GetAllPriceUp(ctx context.Context) ([]domain.Product, error) {

	products := []domain.Product{}
	rows, err := p.db.Query(context.Background(), getAllProductPriceUpQuery)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		r, err := rows.Values()

		fmt.Println(r)
		if err != nil {
			return nil, err
		}
		product := domain.Product{}
		product.ID = uint(r[0].(int32))
		product.Name = r[1].(string)
		product.Title = r[2].(string)
		product.Description = r[3].(string)
		product.ImageSrc = r[4].(string)
		product.VideoSrc = r[5].(string)
		product.Price = uint(r[6].(float32))
		product.Weight = r[7].(float32)
		product.Length = r[8].(string)
		product.Height = r[9].(string)
		product.Depth = r[10].(string)
		product.Icon = r[11].(string)
		product.Status = uint(r[12].(int32))

		products = append(products, product)
	}

	return products, nil
}

func (p *productRepository) GetAllPriceDown(ctx context.Context) ([]domain.Product, error) {

	products := []domain.Product{}
	rows, err := p.db.Query(context.Background(), getAllProductPriceDownQuery)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		r, err := rows.Values()

		fmt.Println(r)
		if err != nil {
			return nil, err
		}
		product := domain.Product{}
		product.ID = uint(r[0].(int32))
		product.Name = r[1].(string)
		product.Title = r[2].(string)
		product.Description = r[3].(string)
		product.ImageSrc = r[4].(string)
		product.VideoSrc = r[5].(string)
		product.Price = uint(r[6].(float32))
		product.Weight = r[7].(float32)
		product.Length = r[8].(string)
		product.Height = r[9].(string)
		product.Depth = r[10].(string)
		product.Icon = r[11].(string)
		product.Status = uint(r[12].(int32))

		products = append(products, product)
	}

	return products, nil
}

func (p *productRepository) ChangeOrder(ctx context.Context, a uint, b uint) error {
	aa := 0
	bb := 0
	if err := p.db.QueryRow(ctx, ordOfProduct, a).Scan(&aa); err != nil {
		return errors.Wrap(err, "(Product with given id does not exist) Scan")
	}

	if err := p.db.QueryRow(ctx, ordOfProduct, b).Scan(&bb); err != nil {
		return errors.Wrap(err, "(Product with given id does not exist) Scan")
	}

	_, err := p.db.Exec(ctx, changeOrderProduct, a, b, bb, aa)
	if err != nil {
		return err
	}

	return nil
}
