CREATE TABLE author (
    id      bigint  not null auto_increment primary key,
    name    text    not null,
    bio     text,
    config  json    not null
);
