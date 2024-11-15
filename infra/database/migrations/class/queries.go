package classmigrations

const upQuery = `
	CREATE TABLE IF NOT EXISTS classes (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    teacher_id UUID NOT NULL,
    students_ids UUID[],
    contents_ids UUID[],
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_teacher
        FOREIGN KEY (teacher_id) 
        REFERENCES teachers(id)
	);
`

const downQuery = `
	DROP TABLE IF EXISTS classes;
`