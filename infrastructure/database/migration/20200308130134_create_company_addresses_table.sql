-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE addresses (
	id serial NOT NULL,
    company_id int NOT NULL,
	uuid char(36) NOT NULL,
	country varchar(80) NOT NULL,
	city varchar(80) NOT NULL,
    state varchar(80) NOT NULL,
    street varchar(80) NOT NULL,
    number varchar(80) NOT NULL
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE addresses;