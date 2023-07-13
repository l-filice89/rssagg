-- +goose Up

CREATE TABLE feeds
(
    id         UUID PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    name       TEXT                     NOT NULL,
    url        TEXT UNIQUE              NOT NULL,
    user_id    UUID REFERENCES users (id) ON DELETE CASCADE NOT NULL
);

-- +goose Down

DROP TABLE feeds;