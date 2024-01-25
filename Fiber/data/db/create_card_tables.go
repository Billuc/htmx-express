package db

const CREATE_TAGS string = `
CREATE TABLE IF NOT EXISTS tags (
	id STRING PRIMARY KEY,
	name TEXT NOT NULL,
	color TEXT DEFAULT '',
	icon TEXT DEFAULT ''
);
`

const CREATE_CARDS string = `
CREATE TABLE IF NOT EXISTS cards (
	id STRING PRIMARY KEY,
	title TEXT NOT NULL,
	description TEXT DEFAULT '',
	has_date BOOLEAN NOT NULL,
	date TEXT DEFAULT '',
	frequency INTEGER DEFAULT 0,
	frequency_unit TEXT DEFAULT '',
	tag_id STRING,
	parent_id STRING,
	FOREIGN KEY (tag_id) REFERENCES tags(id),
	FOREIGN KEY (parent_id) REFERENCES cards(id)
);
`
