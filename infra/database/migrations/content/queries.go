package contentmigrations

const upQuery = `
CREATE TABLE IF NOT EXISTS content (
    id UUID PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    thumbnail_link VARCHAR(255),
    content_link VARCHAR(255),
    content_type VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    viewed BOOLEAN NOT NULL
);`

const downQuery = `
DROP TABLE IF EXISTS content;`
