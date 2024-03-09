-- +goose Up
-- +goose StatementBegin
CREATE TABLE bot_actions (
     id INT UNSIGNED PRIMARY KEY NOT NULL AUTO_INCREMENT,
     bot_id INT UNSIGNED NOT NULL,
     parent_id INT UNSIGNED,
     name VARCHAR(255),
     parent_condition VARCHAR(255),
     action_name VARCHAR(255),
     action_data TEXT,

     FOREIGN KEY (bot_id) REFERENCES bots(id) ON UPDATE RESTRICT ON DELETE CASCADE,
     INDEX idx_id (id),
     INDEX idx_bot_id (bot_id),
     INDEX idx_parent_id (parent_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE bot_actions
-- +goose StatementEnd
