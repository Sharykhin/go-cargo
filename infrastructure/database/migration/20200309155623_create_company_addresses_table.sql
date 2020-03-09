-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE company_addresses (
	company_id int NOT NULL, 
	address_id cint NOT NULL, 
	number varchar(50) NOT NULL,
    is_default boolean NOT NULL,
    PRIMARY KEY (company_id, address_id)
);
ALTER TABLE company_addresses 
ADD CONSTRAINT fk_company_addresses_company_id_companies_id FOREIGN KEY (company_id) REFERENCES companies (id);

ALTER TABLE company_addresses 
ADD CONSTRAINT fk_company_addresses_address_id_addresses_id FOREIGN KEY (address_id) REFERENCES addresses (id);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
ALTER TABLE addresses DROP CONSTRAINT fk_company_addresses_company_id_companies_id;
ALTER TABLE addresses DROP CONSTRAINT fk_company_addresses_address_id_addresses_id;
DROP TABLE company_addresses;
