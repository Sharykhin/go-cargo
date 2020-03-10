-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE OR REPLACE FUNCTION check_default_address()
  RETURNS trigger AS
$BODY$
DECLARE
	num varchar;
BEGIN
  SELECT "number" INTO num FROM company_addresses ca
  WHERE company_id = new.company_id AND is_default = true;
	IF num IS NOT NULL THEN
	    RAISE unique_violation USING MESSAGE = 'Duplicate default address NUM: ' || num;
	END IF;

	RETURN NEW;
END;
$BODY$
LANGUAGE plpgsql VOLATILE
;

CREATE TRIGGER trigger_check_default_address
  BEFORE insert
  ON company_addresses
  FOR EACH ROW
  EXECUTE PROCEDURE check_default_address();

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP trigger trigger_check_default_address on company_addresses;
DROP FUNCTION check_default_address;