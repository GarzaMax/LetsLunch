-- +goose Up
-- +goose StatementBegin
ALTER TABLE users
    ADD COLUMN phone_number VARCHAR(255);
ALTER TABLE users
    ADD COLUMN birthday date;
ALTER TABLE users
    DROP column age;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
ALTER TABLE users
    ADD COLUMN age int;

ALTER TABLE users
    DROP COLUMN phone_number,
    DROP COLUMN birthday;
-- +goose StatementEnd
