package repository


// user queries
const (
	createUserQuery = `INSERT INTO users ( name, email, password, gender, height, weight, birth, activity) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8 ) 
	RETURNING  name, email, password, gender, height, weight, birth, activity`

	getByIDUserQuery = `SELECT id, name, email, gender, height, weight, birth, activity FROM users WHERE id = $1`

	getByEmailUserQuery = `SELECT id, name, email, password, gender, height, weight, birth, activity FROM users WHERE email = $1`

	deleteByIDUserQuery = `delete from users where id=$1`

	updateUserQuery = `UPDATE users
						SET name = $1, email = $2, gender = $3, height = $4, weight = $5, birth = $6, activity = $7
						WHERE id = $8 returning id`

)

// day queries
const (
	createDayQuery = `INSERT INTO days ( userid, mon, tue, wed, thu, fri, sat, sun, t) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) 
	RETURNING  id, userid, mon, tue, wed, thu, fri, sat, sun, t`

	getDayQuery = `SELECT id, userid, mon, tue, wed, thu, fri, sat, sun, t FROM days WHERE userid = $1`

	updateDayQuery = `UPDATE days
						SET mon = $1, tue = $2, wed = $3, thu = $4, fri = $5, sat = $6, sun = $7, t = $8
						WHERE userid = $9 returning id`
)

// product queries
const (
	createProductQuery = `INSERT INTO products (Name, Title , Description, ImageSrc, VideoSrc, Price  ,Weight ,Length, Height, Depth, Icon, Status) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) 
	RETURNING  id`

	getProductQuery = `SELECT id, Name, Title , Description, ImageSrc, VideoSrc, Price  ,Weight ,Length, Height, Depth, Icon, Status FROM products WHERE id = $1`

	getAllProductQuery = `SELECT id, Name, Title , Description, ImageSrc, VideoSrc, Price  ,Weight ,Length, Height, Depth, Icon, Status FROM products`

	updateProductQuery = `UPDATE products
						SET name = $1, title = $2, Description = $3, ImageSrc = $4, VideoSrc = $5, Price = $6, Weight = $7, Length = $8, height = $9, depth = $10, icon=$11, status = $12
						WHERE id = $13 returning id`

	deleteProductQuery = `delete from products where id=$1`
)


// lesson queries
const (
	createLessonQuery = `INSERT INTO lessons (Name, Title , Description, ImageSrc, VideoSrc, Duration  ,Created_at, Updated_at) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8) 
	RETURNING  id`

	getLessonQuery = `SELECT id, Name, Title , Description, ImageSrc, VideoSrc, Duration  ,Created_at, Updated_at FROM lessons WHERE id = $1`

	getAllLessonQuery = `SELECT id, Name, Title , Description, ImageSrc, VideoSrc, Duration  ,Created_at, Updated_at FROM lessons`

	updateLessonQuery = `UPDATE lessons
						SET name = $1, title = $2, Description = $3, ImageSrc = $4, VideoSrc = $5, Duration = $6,  Updated_at = $7
						WHERE id = $8 returning id`

	deleteLessonQuery = `delete from lessons where id=$1`
)

// favorites queries
const (
	createfavoritesQuery = `INSERT INTO favorites (userid, lessonid) 
	VALUES ($1, $2) 
	RETURNING  id`

	getfavoritesQuery = `SELECT id, Name, Title , Description, ImageSrc, VideoSrc, Duration  ,Created_at, Updated_at FROM lessons WHERE id in (select lessonid from favorites where userid = $1);`

	// getAllfavoritesQuery = `SELECT id, userid, lessonid FROM favorites where userid = $1`

	deletefavoritesQuery = `delete from favorites where lessonid=$1`
)

// support queries
const (
	createsupportQuery = `INSERT INTO supports (userid, ProblemDescription, AnswerDescription ) 
	VALUES ($1, $2, $3) 
	RETURNING  id`

	getsupportQuery = `SELECT id, userid, ProblemDescription, AnswerDescription FROM supports WHERE userid = $1 `

	getAllsupportQuery = `SELECT id, userid, ProblemDescription, AnswerDescription from supports`

	deletesupportQuery = `delete from supports where userid=$1`

	updateSupportQuery = `UPDATE supports
						SET AnswerDescription = $1
						WHERE userid = $2 returning id, userid, ProblemDescription, AnswerDescription`
)