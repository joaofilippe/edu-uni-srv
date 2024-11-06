package teachermigrations

const upQuery = `
CREATE TABLE IF NOT EXISTS teachers (
    id UUID PRIMARY KEY,
    userID UUID NOT NULL,
    name VARCHAR(255) NOT NULL,
    classesIDs UUID[] NOT NULL,
    createdAt TIMESTAMP NOT NULL,
    updatedAt TIMESTAMP NOT NULL,
    active BOOLEAN NOT NULL,
    CONSTRAINT fk_user
        FOREIGN KEY(userID) 
        REFERENCES users(id)
);
`

const downQuery = `
DROP TABLE IF EXISTS teachers;
`
