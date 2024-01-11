-- +goose Up
-- +goose StatementBegin
CREATE TABLE bots (
    id int unsigned primary key NOT NULL AUTO_INCREMENT,
    token VARCHAR(50) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE bots
-- +goose StatementEnd
