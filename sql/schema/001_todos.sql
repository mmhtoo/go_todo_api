-- +goose Up

CREATE TABLE todos (
    id SERIAL NOT NULL,
    title VARCHAR(256) NOT NULL,
    status VARCHAR(256) NOT NULL DEFAULT 'In-Progress',
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP
);

-- +goose Down
DROP TABLE todos;
