package repository


// user queries
const (
	createUserQuery = `INSERT INTO users ( name, email, password, gender, height, weight, birth, activity) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8 ) 
	RETURNING  name, email, password, gender, height, weight, birth, activity`

	getByIDUserQuery = `SELECT id, name, email, gender, height, weight, birth, activity FROM users WHERE id = $1`

	getByEmailUserQuery = `SELECT id, name, email, password, gender, height, weight, birth, activity FROM users WHERE email = $1`

)

