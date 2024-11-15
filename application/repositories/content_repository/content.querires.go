package contentrepository

const SaveQuery = `
	INSERT INTO content (
		id, 
		title, 
		description, 
		thumbnail_link, 
		content_link, 
		content_type, 
		created_at, 
		updated_at, 
		viewed
	) VALUES (
		:id,
		:title,
		:description,
		:thumbnail_link,
		:content_link,
		:content_type,
		NOW(),
		NULL
		FALSE
	)		
`

const FindAllQuery = `
	SELECT * FROM content
    WHERE unaccent(description) ILIKE '%' || unaccent(:description) || '%'
`

const FindByIDQuery = `
	SELECT * FROM content WHERE id = :id
`

const FindByDescriptionQuery = `
	SELECT * FROM content
    WHERE unaccent(description) ILIKE '%' || unaccent(:description) || '%'
`
