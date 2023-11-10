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
	if err := p.db.QueryRow(
		ctx,
		createProductQuery,   
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

	).Scan(&product.ID); err != nil {
		return domain.Product{}, errors.Wrap(err, "Scan")
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
	
	// u, err := p.GetByID(ctx, fmt.Sprint(product.ID))
	// fmt.Println(u)
	// if (u == domain.Product{}) {
	// 	fmt.Printf(err.Error())
	// 	return domain.Product{}, err
	// }
	

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
		).Scan(&product.ID ); err != nil {
			return domain.Product{}, errors.Wrap(err, "(Product with fiven id does not exist) Scan")
	}
	
	return product, nil

}



func (p *productRepository) Delete(ctx context.Context, id string) (error) {

	res, err := p.db.Exec(ctx, deleteProductQuery, id)
	if err != nil {
		return err
	}
	if res.RowsAffected() != 1 {
		return errors.New("No row found to delete")
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
		product.Title  = r[2].(string)  
		product.Description = r[3].(string)     
		product.ImageSrc    = r[4].(string)
		product.VideoSrc   = r[5].(string)
		product.Price = uint(r[6].(float32))
		product.Weight   = r[7].(float32)   
		product.Length  = r[8].(string)
		product.Height  = r[9].(string)
		product.Depth = r[10].(string)
		product.Icon = r[11].(string)
		product.Status = uint(r[12].(int32))

		products = append(products, product)
	}

	return products, nil
}