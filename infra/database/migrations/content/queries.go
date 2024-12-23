package contentmigrations

const upQuery = `
CREATE TABLE IF NOT EXISTS content (
    id UUID PRIMARY KEY,
    class_id UUID NOT NULL,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    thumbnail_link VARCHAR(2555),
    content_link VARCHAR(255),
    content_type VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    viewed BOOLEAN NOT NULL,
    CONSTRAINT fk_class
        FOREIGN KEY (class_id)
        REFERENCES classes(id)
);`

const downQuery = `
DROP TABLE IF EXISTS content;`
