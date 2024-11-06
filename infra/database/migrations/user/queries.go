package usermigrations

const upQuery = `
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    userType VARCHAR(50) NOT NULL,
    createdAt TIMESTAMP NOT NULL,
    updatedAt TIMESTAMP NOT NULL,
    active BOOLEAN NOT NULL
);
`
const downQuery = `
DROP TABLE IF EXISTS users;
`
