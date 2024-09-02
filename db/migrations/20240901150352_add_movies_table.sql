-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE IF NOT EXISTS movies (
  id SERIAL PRIMARY KEY,
  original_title VARCHAR(255) NOT NULL,
  original_language VARCHAR(255) NOT NULL,
  overview TEXT NOT NULL UNIQUE,
  release_date DATE NOT NULL,
  adult BOOLEAN NOT NULL,
  poster_path TEXT NOT NULL UNIQUE, -- Store the URL of the poster if using external storage
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down

-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
DROP TABLE IF EXISTS movies CASCADE;
