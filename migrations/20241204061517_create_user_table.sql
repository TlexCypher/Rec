-- +goose Up
SELECT 'up SQL query';
CREATE TABLE users (
    id char(36) NOT NULL,
    role varchar(20),
    username varchar(50),
    PRIMARY KEY(id)
);

-- +goose Down
SELECT 'down SQL query';
DROP TABLE users;
