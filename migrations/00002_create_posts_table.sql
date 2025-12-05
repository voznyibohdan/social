-- +goose Up
CREATE TABLE IF NOT EXISTS posts
(
    id         BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    title      TEXT                        NOT NULL,
    user_id    BIGINT                      NOT NULL,
    created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW(),
    content    TEXT                        NOT NULL,
    tags       TEXT[]                      NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS posts;
