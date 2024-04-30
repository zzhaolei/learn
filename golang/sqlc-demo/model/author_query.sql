-- name: GetAuthor :one
select * from author where id = ?;

-- name: ListAuthors :many
select * from author
order by name;

-- name: CreateAuthor :execresult
insert into author (
    name, bio, config
) values (
    ?, ?, ?
);

-- name: DeleteAuthor :exec
delete from author where id = ?;
