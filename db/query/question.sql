-- name: CreateQuestion :one
insert into question (
    quiz_id,
    question,
    answer
) values (
    $1, $2, $3
) returning *;

-- name: GetQuestionById :one
select * from question where id = $1 limit 1;

-- name: ListQuestion :many
select * from question;

-- name: UpdateQuestion :one
update question set
    question = $2,
    answer = $3
where id = $1
returning *;

-- name: GetAllQuestionByQuizId :many
select * from question where quiz_id = $1;

-- name: DeleteQuestion :exec
delete from question where id = $1;