package repository

import (
	"context"
	// "fmt"

	"github.com/dilyara4949/drevmass/internal/domain"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

type dayRepository struct {
	db *pgxpool.Pool
}

func NewDayRepository(db *pgxpool.Pool) domain.DayRepository {
	return &dayRepository{db: db}
}

func (u *dayRepository) Create(ctx context.Context, day domain.Day) (domain.Day, error) {
	
	if err := u.db.QueryRow(
		ctx,
		createDayQuery,
		// &day.ID,   
		&day.UserID, 
		&day.Mon,    
		&day.Tue,     
		&day.Wed,   
		&day.Thu,  
		&day.Fri,
		&day.Sat,     
		&day.Sun, 
		&day.Time,  
	).Scan( &day.ID,
		&day.UserID, 
		&day.Mon,    
		&day.Tue,     
		&day.Wed,   
		&day.Thu,  
		&day.Fri,
		&day.Sat,     
		&day.Sun, 
		&day.Time,); err != nil {
		return domain.Day{}, errors.Wrap(err, "Scan")
	}

	return day, nil
}



func (d *dayRepository) Get(ctx context.Context, userID uint) (domain.Day, error) {
	var day domain.Day
	if err := d.db.QueryRow(ctx, getDayQuery, userID).Scan(
		&day.ID,
		&day.UserID, 
		&day.Mon,    
		&day.Tue,     
		&day.Wed,   
		&day.Thu,  
		&day.Fri,
		&day.Sat,     
		&day.Sun, 
		&day.Time,
	); err != nil {
		return domain.Day{}, errors.Wrap(err, "Scan")
	}

	return day, nil
}

func (u *dayRepository) Update(ctx context.Context, day domain.Day) (domain.Day, error) {
	
	if err := u.db.QueryRow(
		ctx,
		updateDayQuery,  
		&day.Mon,    
		&day.Tue,     
		&day.Wed,   
		&day.Thu,  
		&day.Fri,
		&day.Sat,     
		&day.Sun, 
		&day.Time,  
		&day.UserID, 
	).Scan(&day.ID); err != nil {
		return domain.Day{}, errors.Wrap(err, "Scan10")
	}

	return day, nil
}