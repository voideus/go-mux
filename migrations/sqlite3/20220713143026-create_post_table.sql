
-- +migrate Up
CREATE TABLE IF NOT EXISTS posts (
    id      INTEGER    AUTO_INCREMENT NOT NULL,
    title    VARCHAR(255)                NOT NULL,
    text    VARCHAR(255)                NOT NULL,
    PRIMARY KEY (id)
);
-- +migrate Down
DROP TABLE posts;