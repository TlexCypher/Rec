-- +goose Up
SELECT 'up SQL query';
CREATE TABLE users_inventories(
    id char(36) NOT NULL,
    user_id char(36) NOT NULL,
    inventory_id char(36) NOT NULL,
    represent_product_name varchar(256) NOT NULL,
    represent_product_code varchar(256) NOT NULL,
    PRIMARY KEY(id)
);

-- +goose Down
SELECT 'down SQL query';
DROP TABLE users_inventories;
