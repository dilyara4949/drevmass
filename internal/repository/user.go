package repository

import (
	"context"
	// "fmt"

	"github.com/dilyara4949/drevmass/internal/domain"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

type userRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) domain.UserRepository {
	return &userRepository{db: db}
}

func (u *userRepository) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
	
	if err := u.db.QueryRow(
		ctx,
		createUserQuery,
		// &user.ID,   
		&user.Name, 
		&user.Email,    
		&user.Password,     
		&user.Gender,   
		&user.Height,  
		&user.Weight,
		&user.Birth,     
		&user.Activity,  
	).Scan(&user.Name, &user.Email, &user.Password, &user.Gender,&user.Height, &user.Weight, &user.Birth, &user.Activity, ); err != nil {
		return nil, errors.Wrap(err, "Scan")
	}

	return user, nil
}

func (u *userRepository) GetByID(ctx context.Context, id uint) (*domain.User, error) {
	var user domain.User
	if err := u.db.QueryRow(ctx, getByIDUserQuery, id).Scan(
		&user.ID, &user.Name, &user.Email, &user.Gender, &user.Height, &user.Weight, &user.Birth, &user.Activity,
	); err != nil {
		return nil, errors.Wrap(err, "Scan")
	}

	return &user, nil
}



func (u *userRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	var user domain.User

	if err := u.db.QueryRow(ctx, getByEmailUserQuery, email).Scan(
		&user.ID, &user.Name, &user.Email, &user.Password, &user.Gender, &user.Height, &user.Weight, &user.Birth, &user.Activity,
	); err != nil {
		return nil, errors.Wrap(err, "Scan0")
	}
	// fmt.Println(user)
	return &user, nil
}



