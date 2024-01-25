package db

const CREATE_TAGS string = `
CREATE TABLE IF NOT EXISTS tags (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	color TEXT,
	icon TEXT
);
`

const CREATE_CARDS string = `
CREATE TABLE IF NOT EXISTS cards (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	title TEXT NOT NULL,
	description TEXT,
	has_date BOOLEAN NOT NULL,
	date TEXT,
	frequency INTEGER,
	frequency_unit TEXT,
	tag_id INTEGER FOREIGN KEY REFERENCES tags(id),
	parent_id INTEGER FOREIGN KEY REFERENCES cards(id)
);
`
