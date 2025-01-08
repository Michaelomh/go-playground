-- +goose Up
create table if not exists users(
	id integer PRIMARY KEY AUTOINCREMENT,
	email VARCHAR(255) UNIQUE NOT NULL,
	first_name VARCHAR(255),
	last_name VARCHAR(255),
	created_at datetime not null,
	updated_at datetime not null
);

-- +goose Down
drop table if exists users;
