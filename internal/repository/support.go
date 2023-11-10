package repository

import (
	"context"
	"fmt"

	"github.com/dilyara4949/drevmass/internal/domain"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

type supportRepository struct {
	db *pgxpool.Pool
}

func NewsupportRepository(db *pgxpool.Pool) domain.SupportRepository {
	return &supportRepository{db: db}
}

func (l *supportRepository) Create(ctx context.Context, support domain.Support) (domain.Support, error) {



	if err := l.db.QueryRow(
		ctx,
		createsupportQuery,   
		&support.UserID, 
		&support.ProblemDescription,
		&support.AnswerDescription,        

	).Scan(&support.ID); err != nil {
		return domain.Support{}, errors.Wrap(err, "Scan")
	}

	return support, nil
}




func (l *supportRepository) GetAll(ctx context.Context) ([]domain.Support, error) {

	supports := []domain.Support{}
	rows, err := l.db.Query(context.Background(), getAllsupportQuery)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		r, err := rows.Values()
		fmt.Println(r)
		if err != nil {
			return nil, err
		}
		support := domain.Support{}
		support.ID = uint(r[0].(int32))
		support.UserID = uint(r[1].(int64))
		support.ProblemDescription = r[2].(string)
		support.AnswerDescription  = r[3].(string)  

		supports = append(supports, support)
	}

	return supports, nil
}

func (l *supportRepository) Get(ctx context.Context, userid uint) (domain.Support, error) {

	var support domain.Support
	if err := l.db.QueryRow(ctx, getsupportQuery, userid).Scan(
		&support.ID,
		&support.UserID,
		&support.ProblemDescription,
		&support.AnswerDescription,
	); err != nil {
		
		return domain.Support{}, errors.Wrap(err, "Scan")
	}
	return support, nil
}

func (p *supportRepository) Update(ctx context.Context, ans string, userid string) (domain.Support, error) {
	support := domain.Support{}
	if err := p.db.QueryRow(ctx, 
		updateSupportQuery,
		&ans,
		&userid,
		).Scan(&support.ID,
			&support.UserID,
			&support.ProblemDescription,
			&support.AnswerDescription,	
			 ); err != nil {
			return domain.Support{}, errors.Wrap(err, "(Support with given id does not exist) Scan")
	}
	
	return support, nil

}