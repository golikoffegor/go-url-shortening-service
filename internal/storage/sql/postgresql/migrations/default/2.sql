-- +migrate Up
-- +migrate StatementBegin
CREATE UNIQUE INDEX url_idx ON shortenerurls (url);
-- +migrate StatementEnd