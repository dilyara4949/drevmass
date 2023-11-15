package repository

import (
	"context"
	// "fmt"
	"time"

	// "fmt"

	"github.com/dilyara4949/drevmass/internal/domain"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

type lessonRepository struct {
	db *pgxpool.Pool
}

func NewLessonRepository(db *pgxpool.Pool) domain.LessonRepository {
	return &lessonRepository{db: db}
}

func (l *lessonRepository) Create(ctx context.Context, lesson domain.Lesson) (domain.Lesson, error) {
	ord := 0
	if err := l.db.QueryRow(ctx, findMaxFromLessonOrder,).Scan(&ord); err != nil {
		return domain.Lesson{}, errors.Wrap(err, "Scan")
	}
	ord += 1

	if err := l.db.QueryRow(
		ctx,
		createLessonQuery,   
		&lesson.Name, 
		&lesson.Title,    
		&lesson.Description,     
		&lesson.ImageSrc,   
		&lesson.VideoSrc,  
		&lesson.Duration,
		&lesson.Created_at,     
		&lesson.Updated_at,
		ord,

	).Scan(&lesson.ID); err != nil {
		return domain.Lesson{}, errors.Wrap(err, "Scan")
	}

	return lesson, nil
}

func (l *lessonRepository) GetByID(ctx context.Context, id string) (domain.Lesson, error) {

	created := time.Time{}
	updated := time.Time{}

	var lesson domain.Lesson
	if err := l.db.QueryRow(ctx, getLessonQuery, id).Scan(
		&lesson.ID,
		&lesson.Name, 
		&lesson.Title,    
		&lesson.Description,     
		&lesson.ImageSrc,   
		&lesson.VideoSrc,  
		&lesson.Duration,
		&created,     
		&updated, 
	); err != nil {
		
		return domain.Lesson{}, errors.Wrap(err, "Scan")
	}
	lesson.Created_at = created.String()
	lesson.Updated_at = updated.String()
	return lesson, nil
}


func (l *lessonRepository) Update(ctx context.Context, lesson domain.Lesson) (domain.Lesson, error) {


	if err := l.db.QueryRow(ctx, 
		updateLessonQuery, 
		&lesson.Name, 
		&lesson.Title,    
		&lesson.Description,     
		&lesson.ImageSrc,   
		&lesson.VideoSrc,  
		&lesson.Duration,    
		&lesson.Updated_at,
		&lesson.ID,
		).Scan(&lesson.ID ); err != nil {
			return domain.Lesson{}, errors.Wrap(err, "(Lesson with given id does not exist) Scan")
	}
	
	return lesson, nil

}



func (l *lessonRepository) Delete(ctx context.Context, id string) (error) {

	var ord *int
	if err := l.db.QueryRow(ctx, deleteLessonQuery, id).Scan(&ord); ord == nil {
			return errors.Wrap(err, "(Lesson with given id does not exist) Scan")
	}

	_, err := l.db.Exec(ctx, cleanLessonOrder, &ord)
	if err != nil {
		return err
	}
	
	return nil
}



func (l *lessonRepository) GetAll(ctx context.Context) ([]domain.Lesson, error) {

	lessons := []domain.Lesson{}
	rows, err := l.db.Query(context.Background(), getAllLessonQuery)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		r, err := rows.Values()
		// fmt.Println(r)
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


func (p *lessonRepository) ChangeOrder(ctx context.Context, a uint, b uint)  (error) {
	aa := 0
	bb := 0
	if err := p.db.QueryRow(ctx, ordOfLesson, a).Scan(&aa); err != nil {
		return errors.Wrap(err, "(Lesson with given id does not exist) Scan")
	}

	if err := p.db.QueryRow(ctx, ordOfLesson, b).Scan(&bb); err != nil {
		return errors.Wrap(err, "(lesson with given id does not exist) Scan")
	}

	_, err := p.db.Exec(ctx, changeOrderLesson, a, b, bb, aa)
	if err != nil {
		return err
	}

	return nil
}