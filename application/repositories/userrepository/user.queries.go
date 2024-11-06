package userrepository

const SaveUserQuery = `
	INSERT INTO users (id, email, password, username, user_type, created_at, updated_at, active)
	VALUES ($1, $2, $3, $4, $5, NOW(), NULL, TRUE)`
