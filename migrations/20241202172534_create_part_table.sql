-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE parts (
    id char(36) NOT NULL,
    car_id char(36) NOT NULL,
    name varchar(256),
    description varchar(256),
    price float,
    quantity_in_stock int,
    PRIMARY KEY(id)
);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE parts;
-- +goose StatementEnd