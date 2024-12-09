-- +goose Up
SELECT 'up SQL query';
ALTER TABLE users_inventories
ADD CONSTRAINT fk_users_users_inventories
FOREIGN KEY (user_id)
REFERENCES users(id)
ON DELETE CASCADE
ON UPDATE CASCADE;

-- +goose Down
SELECT 'down SQL query';
ALTER TABLE users_inventories
DROP FOREIGN KEY fk_users_users_inventories;
