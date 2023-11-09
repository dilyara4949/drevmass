package repository

// import (
// 	"github.com/dilyara4949/drevmass/internal/domain"
// 	"github.com/dilyara4949/drevmass/internal/postgres"
// )


// type userRepository struct {
// 	database postgres.Database
// }


// func NewUserRepository(db postgres.Database) domain.UserRepository {
// 	return &userRepository{ database:   db,}
// }


// func (ur *userRepository) Create(user *domain.User) error {
	
// 	_, err := ur.database.InsertUser(user)

// 	return err
// }


// func (ur *userRepository) GetByEmail(email string) (domain.User, error) {
	
// 	user := domain.User{}
	
// 	user, err := ur.database.GetByEmail(email) 

// 	return user, err
// }


// func (ur *userRepository) GetByID(id uint) (domain.User, error) {
	
// 	// user := &domain.User{ID: id}
	
// 	user, err := ur.database.GetUserByID(id) 
	

// 	return user, err

// }