-- +goose Up
SELECT 'up SQL query';
ALTER TABLE categories
ADD CONSTRAINT fk_inventories_categories
FOREIGN KEY (inventory_id)
REFERENCES inventories(id)
ON DELETE CASCADE
ON UPDATE CASCADE;

-- +goose Down
SELECT 'down SQL query';
ALTER TABLE categories
DROP FOREIGN KEY fk_inventories_categories;
