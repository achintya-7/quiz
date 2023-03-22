-- name: CreateQuiz :one
insert into quiz (
    created_by,
    name,
    description
) values (
    $1, $2, $3
) returning *;

-- name: GetQuizById :one
select * from quiz where id = $1 limit 1;

-- name: ListQuiz :many
select * from quiz;

-- name: UpdateQuiz :one
update quiz 
set name = $2, description = $3
where id = $1
returning *;

-- name: DeleteQuiz :exec
delete from quiz where id = $1;


