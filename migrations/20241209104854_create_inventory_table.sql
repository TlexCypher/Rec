-- +goose Up
SELECT 'up SQL query';
CREATE TABLE inventories(
    id char(36) NOT NULL,
    product_code varchar(256) NOT NULL,
    product_name varchar(256) NOT NULL,
    remaining_quantity bigint NOT NULL,
    remarks JSON,
    PRIMARY KEY(id)
);

-- +goose Down
SELECT 'down SQL query';
DROP TABLE inventories;
