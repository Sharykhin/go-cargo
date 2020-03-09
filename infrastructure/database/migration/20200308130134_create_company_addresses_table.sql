-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE addresses (
	id serial NOT NULL, 
    company_id int NOT NULL,
	uuid char(36) NOT NULL,
	country varchar(80) NOT NULL,
	city varchar(50) NOT NULL,
    state varchar(10) NOT NULL,
    street varchar(100) NOT NULL,
    number varchar(80) NOT NULL,
    PRIMARY KEY (uuid)
);

ALTER TABLE addresses 
ADD CONSTRAINT fk_addresses_company_id_companies_id FOREIGN KEY (company_id) REFERENCES companies (id);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
ALTER TABLE addresses DROP CONSTRAINT fk_addresses_company_id_companies_id;
DROP TABLE addresses;