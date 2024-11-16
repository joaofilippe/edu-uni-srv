package teacherrepository

const Save = `
	INSERT INTO teachers (
		id, 
		user_id, 
		classes_ids, 
		created_at, 
		updated_at, 
		active
	) VALUES (
		:id,
		:user_id,
		:classes_ids,
		NOW(),
		NULL,
		TRUE
	)`

const FindByID = `
SELECT id, user_id, classes_ids, created_at, updated_at, active
FROM teachers
WHERE id = $1`

const FindByUserID = `
SELECT *
FROM teachers
WHERE user_id = $1`

const FindAll = `
SELECT id, user_id, classes_ids, created_at, updated_at, active
FROM teachers`

const Update = `
UPDATE teachers
SET user_id = $2, classes_ids = $3, updated_at = $4, active = $5
WHERE id = $1`

const Delete = `
DELETE FROM teachers
WHERE id = $1`
