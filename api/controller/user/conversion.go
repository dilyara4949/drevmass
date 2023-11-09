package user

import (
	"time"

	"github.com/dilyara4949/drevmass/internal/domain"
)


func userInfoToClient(u domain.User) struct {
	ID        uint
	Name      string 
	Email     string 
	Gender    uint 
	Height    uint 
	Weight    uint 
	Birth     time.Time
	Activity  uint 
} {
	
	user := struct {
		ID        uint
		Name      string 
		Email     string 
		Gender    uint 
		Height    uint 
		Weight    uint 
		Birth     time.Time
		Activity  uint 
    }{
		u.ID,
        u.Name,
        u.Email,
		u.Gender,
		u.Height,
		u.Weight,
		u.Birth,
		u.Activity,
    }
	return user
}

func userToBasic(u domain.User) struct {
	ID        uint
	Name      string 
	Email     string 
} {
	
	user := struct {
		ID        uint
		Name      string 
		Email     string 
    }{
		u.ID,
        u.Name,
        u.Email,
    }
	return user
}