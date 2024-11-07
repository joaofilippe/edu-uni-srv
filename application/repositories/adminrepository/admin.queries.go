package adminrepository

const SaveAdminQuery = `
	INSERT INTO admins (id, user_id, created_at, updated_at, active)
	VALUES ($1, $2, NOW(), NULL, TRUE)`

const FindAllAdminsQuery = `
	SELECT id, email, password, username, created_at, updated_at, active
	FROM admins`

const FindByIDQuery = `
	SELECT id, created_at, updated_at, active
	FROM admins
	WHERE id = $1`

const FindByUserIDQuery = `
	SELECT id, user_id, created_at, updated_at, active
	FROM admins
	WHERE user_id = $1`

const UpdateAdminQuery = `
	UPDATE admins
	SET email = $1, password = $2, username = $3, updated_at = NOW()
	WHERE id = $4`

const DeleteAdminQuery = `
	DELETE FROM admins
	WHERE id = $1`
