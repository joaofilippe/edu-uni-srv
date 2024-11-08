package guardianrepository

const SaveQuery = `
	INSERT INTO guardians (id, user_id, student_id, created_at, updated_at, active)
	VALUES ($1, $2, $3, $4, $5, $6)
`

const FindByUserIDQuery = `
	SELECT * FROM guardians WHERE user_id = $1
`

const FindByIDQuery = `
	SELECT * FROM guardians WHERE id = $1
`

const FindByStudentIDQuery = `
	SELECT * FROM guardians WHERE student_id = $1
`

const FindAllQuery = `
	SELECT * FROM guardians
`