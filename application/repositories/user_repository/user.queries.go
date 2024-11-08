package userrepository

const SaveUserQuery = `
	INSERT INTO users (id, username, email, password, user_type, created_at, updated_at, active)
	VALUES (:id, :username,:email, :password, :user_type, NOW(), NULL, TRUE)`

const FindByEmailQuery = `
	SELECT id, email, password, username, user_type, created_at, updated_at, active 
	FROM users
	WHERE email = $1`
