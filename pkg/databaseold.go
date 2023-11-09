package pkg

// import (
// 	"fmt"
// 	"log"

// 	p "github.com/dilyara4949/drevmass/internal/postgres"
// )


// func NewDatabase(env *Env) p.Database {

// 	username := env.DBUser
// 	password := env.DBPass
// 	host := env.DBHost
// 	port := env.DBPort
// 	dbname := env.DBName

// 	// dbHost := os.Getenv("DB_HOST")
// 	// userName := os.Getenv("DB_USERNAME")
// 	// pass := os.Getenv("DB_PASSWORD")
// 	// dbName := os.Getenv("DB_NAME")
// 	// dbPort := os.Getenv("DB_PORT")


// 	url := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",host, username, password, dbname, port)
	
// 	db, err := p.NewDatabase(url)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return db
// }



// func CloseDatabase(db p.Database) {
// 	if db == nil {
// 		return
// 	}
// 	err := db.Close()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Println("Connection to postgres closed")
// }