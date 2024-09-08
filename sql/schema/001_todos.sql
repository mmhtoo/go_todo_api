-- +goose Up

CREATE TABLE todos (
    id SERIAL NOT NULL,
    title VARCHAR(256) NOT NULL,
    status VARCHAR(256) NOT NULL DEFAULT 'In-Progress',
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE todos;
