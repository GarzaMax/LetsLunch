-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
ALTER TABLE gathering_places
ADD column description text,
ADD column photo_link text,
ADD column title text;
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
ALTER TABLE gathering_places
DROP column description,
DROP column photo_link,
DROP column title;
-- +goose StatementEnd
