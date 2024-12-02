-- +goose Up
CREATE TABLE cars (
    id char(36) NOT NULL,
    make varchar(256) NOT NULL,
    model varchar(256) NOT NULL,
    year int,
    color varchar(256),
    price float,
    mileage float,
    transmission varchar(256),
    engine varchar(256),
    status varchar(256),
    created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
);

-- +goose Down
DROP TABLE cars;
