package studentmigrations

const upQuery = `
CREATE TABLE IF NOT EXISTS students (
    id UUID PRIMARY KEY,
    userID UUID NOT NULL,
    name VARCHAR(255) NOT NULL,
    age INT NOT NULL,
    classes UUID[] NOT NULL,
    grades UUID[] NOT NULL,
    disabilities UUID[] NOT NULL,
    guardianID UUID NOT NULL,
    address VARCHAR(255) NOT NULL,
    phone VARCHAR(20) NOT NULL,
    createdAt TIMESTAMP NOT NULL,
    updatedAt TIMESTAMP NOT NULL,
    active BOOLEAN NOT NULL,
    CONSTRAINT fk_user
        FOREIGN KEY(userID) 
        REFERENCES users(id)
);`

const downQuery = `
DROP TABLE IF EXISTS students;
`
