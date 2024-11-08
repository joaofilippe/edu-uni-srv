package studentrepository

const SaveQuery = `
	INSERT INTO students (id, user_id, age, classes, grades, disabilities, guardian_id, address, phone, created_at, updated_at, active)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
`

const FindByIDQuery = `
	SELECT * FROM students WHERE id = $1
`

const FindByUserIDQuery = `
	SELECT * FROM students WHERE user_id = $1
`

const FindAllQuery = `
	SELECT * FROM students
`

const FindByGuardianIDQuery = `
	SELECT * FROM students WHERE guardian_id = $1
	`
