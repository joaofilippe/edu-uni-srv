package studentrepository

const SaveQuery = `
	INSERT INTO students (
		id, 
		user_id, 
		age, 
		classes_id, 
		grades_id, 
		disabilities, 
		guardian_id, 
		address, 
		phone, 
		created_at, 
		updated_at, 
		active
	)
	VALUES (
		:id, 
		:user_id, 
		:age, 
		:classes_id, 
		:grades_id, 
		:disabilities, 
		:guardian_id, 
		:address, 
		:phone, 
		NOW(), 
		NULL, 
		TRUE)
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

const UpdateQuery = `
	UPDATE students SET
		age = :age,
		classes_id = :classes_id,
		grades_id = :grades_id,
		disabilities = :disabilities,
		guardian_id = :guardian_id,
		address = :address,
		phone = :phone,
		updated_at = NOW()
	WHERE id = :id
`