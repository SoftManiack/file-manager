CREATE TABLE users
(
    uid UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    email VARCHAR(255) NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    password_hash VARCHAR(255) NOT NULL
);

CREATE TABLE directories
(
    uid UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    users_uid REFERENCES users (uid) ON DELETE CASCADE NOT NULL,
    root_uid UUID REFERENCES directories (uid) ON DELETE CASCADE NOT NULL,
    date_update TIMESTAMP NOT NULL DEFAULT current_timestamp,
    is_favorites BOOLEAN DEFAULT FALSE,
    size VARCHAR(255) NOT NULL DEFAULT '0',
    count_element INTEGER NOT NULL DEFAULT 0
);

CREATE TABLE files
(
    uid UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    root_uid UUID REFERENCES directories (uid) ON DELETE CASCADE NOT NULL,
    date_update TIMESTAMP NOT NULL DEFAULT current_timestamp,
    name VARCHAR(255) NOT NULL,
    is_favorites BOOLEAN DEFAULT FALSE,
    extension VARCHAR(255) NOT NULL,
    size VARCHAR(255) NOT NULL DEFAULT '0'
);

CREATE TABLE link
(
    uid UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    users_uid UUID REFERENCES users (uid) ON DELETE CASCADE NOT NULL
);

CREATE TABLE link_files
(
    link_uid UUID REFERENCES link (uid) ON DELETE CASCADE NOT NULL,
    files_uid UUID REFERENCES files (uid) ON DELETE CASCADE NOT NULL
);

CREATE TABLE link_directories
(
    link_uid UUID REFERENCES link (uid) ON DELETE CASCADE NOT NULL,
    directories_uid UUID REFERENCES directories (uid) ON DELETE CASCADE NOT NULL
);


CREATE TABLE reсent
(
    uid UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    files_uid UUID REFERENCES files (uid) ON DELETE CASCADE NOT NULL
);

CREATE TABLE trash
(
    uid UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    users_uid UUID REFERENCES users (uid) ON DELETE CASCADE NOT NULL
);

CREATE TABLE trash_files
(
    trash_uid UUID REFERENCES trash (uid) ON DELETE CASCADE NOT NULL,
    files_uid UUID REFERENCES files (uid) ON DELETE CASCADE NOT NULL
);

CREATE TABLE trash_directories
(
    trash_uid UUID REFERENCES trash (uid) ON DELETE CASCADE NOT NULL,
    directories_uid UUID REFERENCES directories (uid) ON DELETE CASCADE NOT NULL
);
