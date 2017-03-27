-- Revert demo:address from pg

BEGIN;

ALTER TABLE users DROP COLUMN address;

COMMIT;
