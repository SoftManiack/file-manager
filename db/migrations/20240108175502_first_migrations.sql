-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
    uid UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    email VARCHAR(255) NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    size BIGINT NOT NULL DEFAULT 0
);

CREATE TABLE directories
(
    uid UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    users_uid UUID REFERENCES users(uid) ON DELETE CASCADE NOT NULL,
    root_uid UUID REFERENCES directories(uid) ON DELETE CASCADE NOT NULL,
    date_create TIMESTAMP NOT NULL DEFAULT current_timestamp,
    date_update TIMESTAMP NOT NULL DEFAULT current_timestamp,
    name VARCHAR(255) NOT NULL,
    path VARCHAR(255) NOT NULL DEFAULT '/',
    is_delete BOOLEAN NOT NULL DEFAULT FALSE,
    is_link  BOOLEAN NOT NULL DEFAULT FALSE,
    is_favorites BOOLEAN NOT NULL DEFAULT FALSE,
    type VARCHAR(10) NOT NULL DEFAULT 'directory',
    size BIGINT NOT NULL NOT NULL DEFAULT 0,
    count_element INTEGER NOT NULL DEFAULT 0
);

CREATE TABLE files
(
    uid UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    root_uid UUID REFERENCES directories (uid) ON DELETE CASCADE NOT NULL,
    date_create TIMESTAMP NOT NULL DEFAULT current_timestamp,
    date_update TIMESTAMP NOT NULL DEFAULT current_timestamp,
    name VARCHAR(255) NOT NULL,
    path VARCHAR(255) NOT NULL DEFAULT '/',
    is_delete BOOLEAN NOT NULL DEFAULT FALSE,
    is_link BOOLEAN NOT NULL DEFAULT FALSE,
    is_favorites BOOLEAN NOT NULL DEFAULT FALSE,
    type VARCHAR(8) NOT NULL DEFAULT 'default',
    size BIGINT NOT NULL DEFAULT 0,
    data BYTEA DEFAULT '\000'
);

CREATE TABLE trash
(
    uid UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    users_uid UUID REFERENCES users(uid) ON DELETE CASCADE NOT NULL
);

CREATE TABLE trash_direcories
(
    uid UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    trash_uid UUID REFERENCES trash(uid) ON DELETE CASCADE NOT NULL,
    directories_uid UUID REFERENCES directories(uid) ON DELETE CASCADE NOT NULL
);

CREATE TABLE trash_files
(
    uid UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    trash_uid UUID REFERENCES trash(uid) ON DELETE CASCADE NOT NULL,
    files_uid UUID REFERENCES files(uid) ON DELETE CASCADE NOT NULL
);

CREATE TABLE recents
(
    uid UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    users_uid UUID REFERENCES users(uid) ON DELETE CASCADE NOT NULL
);

CREATE TABLE recents_file
(
    uid UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    recents_uid UUID REFERENCES recents(uid) ON DELETE CASCADE NOT NULL,
    files_uid UUID REFERENCES files(uid) ON DELETE CASCADE NOT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
    DROP TABLE IF EXISTS users CASCADE;

    DROP TABLE IF EXISTS files CASCADE;

    DROP TABLE IF EXISTS directories CASCADE;

    DROP TABLE IF EXISTS trash CASCADE;

    DROP TABLE IF EXISTS trash_direcories CASCADE;

    DROP TABLE IF EXISTS trash_files CASCADE;

    DROP TABLE IF EXISTS recents CASCADE;

    DROP TABLE IF EXISTS recents_file CASCADE;

-- +goose StatementEnd

-- Таблица кбелей мощностей в ней скрипт

-- Подбор кабеля по парамтерам
-- Програмка в питоне с формам и тоже самое как в экселе
