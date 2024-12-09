-- +goose Up
SELECT 'up SQL query';
ALTER TABLE users_inventories
ADD CONSTRAINT fk_inventories_users_inventories
FOREIGN KEY (inventory_id)
REFERENCES inventories(id)
ON DELETE CASCADE
ON UPDATE CASCADE;

-- +goose Down
SELECT 'down SQL query';
ALTER TABLE users_inventories
DROP FOREIGN KEY fk_inventories_users_inventories;
