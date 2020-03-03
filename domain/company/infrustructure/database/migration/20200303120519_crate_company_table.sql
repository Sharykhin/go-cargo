-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE companies (
	id bigint NOT NULL,
	uuid char(36) NOT NULL,
	name varchar(80) NOT NULL,
	created_at timestamp NOT NULL
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE companies