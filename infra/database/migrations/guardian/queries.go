package guardianmigrations

const upQuery = `
CREATE TABLE IF NOT EXISTS guardians (
    id UUID PRIMARY KEY,
    userID UUID NOT NULL,
    studentID UUID NOT NULL,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    createdAt TIMESTAMP NOT NULL,
    updatedAt TIMESTAMP NOT NULL,
    active BOOLEAN NOT NULL,
    CONSTRAINT fk_user
        FOREIGN KEY(userID)
        REFERENCES users(id)
);
`

const downQuery = `
DROP TABLE IF EXISTS guardians;
`
