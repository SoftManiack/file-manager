CREATE TABLE users
(
    uid UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    email VARCHAR(255) NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    password_hash VARCHAR(255) NOT NULL
);

CREATE TABLE items
(
    uid UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    root_uid UUID REFERENCES items (uid) ON DELETE CASCADE NOT NULL,
    name VARCHAR(255) NOT NULL,
    size VARCHAR(255) NOT NULL DEFAULT '0',
    type VARCHAR(255) NOT NULL,
    is_favorites BOOLEAN DEFAULT FALSE,
    date TIMESTAMP NOT NULL DEFAULT current_timestamp,
    link BOOLEAN NOT NULL DEFAULT FALSE
);

INSERT INTO items ( root_uid, name, type, is_favorites, date, link) VALUES
    ( 'c5661a60-2bb2-4b8e-8e51-890efcfdec5e','Папка','Directory', false, '2022-01-21 04:05:06', false ),
    ( 'c5661a60-2bb2-4b8e-8e51-890efcfdec5e','Фильмы','Directory', false, '2022-01-21 04:14:06', false ),
    ( 'c5661a60-2bb2-4b8e-8e51-890efcfdec5e','Text1.txt','File', true, '2022-01-21 04:15:06', false ),
    ( 'c5661a60-2bb2-4b8e-8e51-890efcfdec5e','Text3.txt','File', false, '2022-01-21 04:18:06', false );

INSERT INTO items (  uid, root_uid, name, type, is_favorites, date, link) VALUES
    ( 'c5661a60-2bb2-4b8e-8e51-890efcfdec5e','c5661a60-2bb2-4b8e-8e51-890efcfdec5e','Root','Directory', false, '2022-01-21 04:02:06', false );