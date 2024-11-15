package classrepository

const SaveQuery = `
	INSERT INTO classes (
		id, 
		name, 
		teacher_id, 
		students_ids, 
		contents_ids, 
		created_at,
		updated_at
	) VALUES (
		:id,
		:name,
		:teacher_id,
		:students_ids,
		:contents_ids,
		:created_at,
		:updated_at
	)
`