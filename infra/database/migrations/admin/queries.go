package adminmigrations

const upQuery = `
CREATE TABLE IF NOT EXISTS admins (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    active BOOLEAN NOT NULL,
    CONSTRAINT fk_user
        FOREIGN KEY(user_id)
        REFERENCES users(id)
);
`
const downQuery = `
DROP TABLE IF EXISTS admins;
`
