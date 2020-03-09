-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE addresses (
	id serial NOT NULL, 
	uuid char(36) NOT NULL,
	country varchar(80) NOT NULL,
	city varchar(50) NOT NULL,
    state varchar(10) NOT NULL,
    street varchar(100) NOT NULL,
    zip varchar(10) NOT NULL,
    PRIMARY KEY (uuid)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE addresses;