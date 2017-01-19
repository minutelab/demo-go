-- Deploy demo:initial to pg

BEGIN;

CREATE TABLE users (
    id       VARCHAR(20)  PRIMARY KEY,
    name     VARCHAR(256) NOT NULL,
    email    VARCHAR(256)
);

CREATE UNIQUE INDEX users_email_idx ON users (email)  WHERE email <> '';

COMMIT;
