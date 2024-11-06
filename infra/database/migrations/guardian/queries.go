package guardianmigrations

const upQuery = `
CREATE TABLE IF NOT EXISTS guardians (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    student_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    active BOOLEAN NOT NULL,
    CONSTRAINT fk_user
        FOREIGN KEY(userID)
        REFERENCES users(id)
);
`

const downQuery = `
DROP TABLE IF EXISTS guardians;
`
