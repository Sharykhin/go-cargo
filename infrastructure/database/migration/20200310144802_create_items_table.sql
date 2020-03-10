-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE items (
	id serial NOT NULL UNIQUE,
	uuid char(36) NOT NULL,
	company_id int NOT NULL,
	name varchar(80) NOT NULL,
	attributes jsonb NOT NULL,
	created_at timestamp NOT NULL,
	PRIMARY KEY (uuid)
);

ALTER TABLE items
ADD CONSTRAINT fk_items_company_id_companies_id FOREIGN KEY (company_id) REFERENCES companies (id);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
ALTER TABLE items DROP CONSTRAINT fk_items_company_id_companies_id;
DROP TABLE items;