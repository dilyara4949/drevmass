package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/dilyara4949/drevmass/internal/domain"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

type favoritesRepository struct {
	db *pgxpool.Pool
}

func NewfavoritesRepository(db *pgxpool.Pool) domain.FavoritesRepository {
	return &favoritesRepository{db: db}
}

func (l *favoritesRepository) Create(ctx context.Context, favorites domain.Favorites) (domain.Favorites, error) {

	// format := "2006-01-02"

	// s1, _ := time.Parse(format, "2018-12-12")

	if err := l.db.QueryRow(
		ctx,
		createfavoritesQuery,   
		&favorites.UserID, 
		&favorites.LessonID,        

	).Scan(&favorites.ID); err != nil {
		return domain.Favorites{}, errors.Wrap(err, "Scan")
	}

	return favorites, nil
}







func (l *favoritesRepository) Delete(ctx context.Context, lessonid string) (error) {

	res, err := l.db.Exec(ctx, deletefavoritesQuery, lessonid)
	if err != nil {
		return err
	}
	if res.RowsAffected() != 1 {
		return errors.New("No row found to delete")
	}
	return nil
}



func (l *favoritesRepository) Get(ctx context.Context, userid uint) ([]domain.Lesson, error) {

	lessons := []domain.Lesson{}
	rows, err := l.db.Query(context.Background(), getfavoritesQuery, userid)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		r, err := rows.Values()
		fmt.Println(r)
		if err != nil {
			return nil, err
		}

		lesson := domain.Lesson{}
		lesson.ID = uint(r[0].(int32))
		lesson.Name = r[1].(string)
		lesson.Title  = r[2].(string)  
		lesson.Description = r[3].(string)     
		lesson.ImageSrc    = r[4].(string)
		lesson.VideoSrc   = r[5].(string)
		lesson.Duration = float64(r[6].(float32))
		lesson.Created_at  = r[7].(time.Time).String()
		lesson.Updated_at  = r[8].(time.Time).String()

		lessons = append(lessons, lesson)
	}

	return lessons, nil
}