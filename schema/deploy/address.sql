-- Deploy demo:address to pg

BEGIN;

ALTER TABLE users ADD COLUMN adress VARCHAR DEFAULT '' NOT NULL;

COMMIT;
