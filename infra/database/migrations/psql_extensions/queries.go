package psqlextensions

const upQuery = `
	CREATE EXTENSION IF NOT EXISTS unaccent;
`

const downQuery = `
	DROP EXTENSION IF EXISTS unaccent;
`
