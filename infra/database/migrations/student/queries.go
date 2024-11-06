package studentmigrations

const upQuery = `
CREATE TABLE IF NOT EXISTS students (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    age INT NOT NULL,
    classes UUID[] NOT NULL,
    grades UUID[] NOT NULL,
    disabilities UUID[] NOT NULL,
    guardian_id UUID NOT NULL,
    address VARCHAR(255) NOT NULL,
    phone VARCHAR(20) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    active BOOLEAN NOT NULL,
    CONSTRAINT fk_user
        FOREIGN KEY(userID) 
        REFERENCES users(id)
);`

const downQuery = `
DROP TABLE IF EXISTS students;
`
