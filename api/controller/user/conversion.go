package user

import (
	// "time"

	"github.com/dilyara4949/drevmass/internal/domain"
)


func userInfoToClient(u domain.User) domain.UserInfo{
	
	user := domain.UserInfo{
		ID: u.ID,
        Name: u.Name,
        Email: u.Email,
		Gender: u.Gender,
		Height: u.Height,
		Weight: u.Weight,
		Birth: u.Birth,
		Activity: u.Activity,
    }
	return user
}

func userToBasic(u domain.User) UserBasic {
	
	user := UserBasic {
		ID : u.ID,
        Name: u.Name,
        Email: u.Email,
    }
	return user
}

type UserBasic struct {
	ID        uint
	Name      string 
	Email     string 
}