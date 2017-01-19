-- Revert demo:initial from pg

BEGIN;

DROP TABLE users;

COMMIT;
