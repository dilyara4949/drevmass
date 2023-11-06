package postgres

import (
	"github.com/dilyara4949/drevmass/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type database struct {
	DB *gorm.DB
}

type Database interface {
	InsertUser(object *domain.User) (interface{}, error)
	GetByEmail(email string) (domain.User, error)
	GetUserByID(id uint) (domain.User, error)
	Close() error
}

func NewDatabase(url string) (Database, error){
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&domain.User{})
	
	return &database{
		DB: db,}, 
		err
}


func (db *database) GetUserByID(id uint) (domain.User, error) {
	user := domain.User{}
	result:= db.DB.Where(&domain.User{ID: id}).Find(&user)

	return user, result.Error
}

func (db *database) InsertUser(object *domain.User) (interface{}, error){
	
	result := db.DB.Create(object)

	return object, result.Error
}



func (db *database) GetByEmail(email string) (domain.User, error) {
	// result := db.DB.First(object)

	user := domain.User{}
	result:= db.DB.Where(&domain.User{Email: email}).Find(&user)
	// fmt.Println(user)
	return user, result.Error
}


func (db *database) Close() error{
	// Get generic database object sql.DB to use its functions
	sqlDB, err := db.DB.DB()
	if err != nil {
		return err
	}

	err = sqlDB.Close()
	return err
}


