-- +migrate Up
-- +migrate StatementBegin
ALTER TABLE shortenerurls ADD COLUMN user_id CHAR(128) DEFAULT NULL;
-- +migrate StatementEnd