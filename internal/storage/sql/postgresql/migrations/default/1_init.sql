-- +migrate Up
-- +migrate StatementBegin
CREATE TABLE shortenerurls (
    id SERIAL PRIMARY KEY,
    url text NOT NULL,
    url_key text NOT NULL
);
-- +migrate StatementEnd