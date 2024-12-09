-- +goose Up
SELECT 'up SQL query';
CREATE TABLE categories(
    id char(36) NOT NULL,
    name varchar(256) NOT NULL,
    inventory_id char(36) NOT NULL,
    PRIMARY KEY(id)
);

-- +goose Down
SELECT 'down SQL query';
DROP TABLE categories;
