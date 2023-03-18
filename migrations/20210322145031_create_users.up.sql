CREATE TABLE users (
    id bigserial not null primary key,
    login varchar not null unique,
    password varchar not null
);

CREATE TABLE words (
    id bigserial not null primary key,
    english varchar not null unique,
    russian varchar not null,
    theme varchar 
);