package pkg


import (
	"fmt"
	"log"

	p "github.com/dilyara4949/drevmass/internal/postgres"
	// "gorm.io/driver/postgres"
	// "gorm.io/gorm"
)


func NewDatabase(env *Env) p.Database {

	username := env.DBUser
	password := env.DBPass
	host := env.DBHost
	port := env.DBPort
	dbname := env.DBName

	url := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",host, username, password, dbname, port)

	db, err := p.NewDatabase(url)
	if err != nil {
		log.Fatal(err)
	}

	return db
}



func CloseDatabase(db p.Database) {
	if db == nil {
		return
	}
	err := db.Close()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connection to postgres closed")
}