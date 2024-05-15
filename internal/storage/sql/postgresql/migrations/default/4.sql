-- +migrate Up
-- +migrate StatementBegin
ALTER TABLE shortenerurls ADD COLUMN is_deleted BOOLEAN NOT NULL DEFAULT FALSE;
-- +migrate StatementEnd