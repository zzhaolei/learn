CREATE TABLE "user" (
    id serial primary key,
    username varchar(255) unique not null,
    password varchar(255) not null,
    gender smallint not null
);

CREATE INDEX idx_user_username_password ON "user" ("username", "password");
