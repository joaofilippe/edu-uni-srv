package adminmigrations

const upQuery = `
CREATE TABLE IF NOT EXISTS admins (
    	id SERIAL PRIMARY KEY,
    	name VARCHAR(255) NOT NULL,
    	email VARCHAR(255) NOT NULL,
    	password VARCHAR(255) NOT NULL,
    	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    	updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
`

const downQuery = `
DROP TABLE IF EXISTS admins;
`
