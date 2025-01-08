-- +goose Up
-- +goose StatementBegin
-- First create the new table with the desired structure
CREATE TABLE users_new (
    id integer PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(255),
    created_at datetime not null,
    updated_at datetime not null
);

-- Copy the data from the old table to the new table
-- We'll concatenate first_name and last_name for the name field
INSERT INTO users_new (id, name, created_at, updated_at)
SELECT 
    id,
    first_name,
    created_at,
    updated_at
FROM users;

-- Drop the old table
DROP TABLE users;

-- Rename the new table to the original name
ALTER TABLE users_new RENAME TO users;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users ADD COLUMN last_name VARCHAR(255) NOT NULL;
ALTER TABLE users ADD COLUMN email VARCHAR(255) NOT NULL; 
CREATE UNIQUE INDEX IF NOT EXISTS uni_email ON users(email);
ALTER TABLE users RENAME COLUMN name TO first_name;
-- +goose StatementEnd
