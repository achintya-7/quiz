// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: quiz.sql

package db

import (
	"context"
)

const createQuiz = `-- name: CreateQuiz :one
insert into quiz (
    created_by,
    name,
    description
) values (
    $1, $2, $3
) returning id, created_by, name, description, created_at
`

type CreateQuizParams struct {
	CreatedBy   string `json:"created_by"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (q *Queries) CreateQuiz(ctx context.Context, arg CreateQuizParams) (Quiz, error) {
	row := q.db.QueryRowContext(ctx, createQuiz, arg.CreatedBy, arg.Name, arg.Description)
	var i Quiz
	err := row.Scan(
		&i.ID,
		&i.CreatedBy,
		&i.Name,
		&i.Description,
		&i.CreatedAt,
	)
	return i, err
}

const deleteQuiz = `-- name: DeleteQuiz :exec
delete from quiz where id = $1
`

func (q *Queries) DeleteQuiz(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteQuiz, id)
	return err
}

const getQuizById = `-- name: GetQuizById :one
select id, created_by, name, description, created_at from quiz where id = $1 limit 1
`

func (q *Queries) GetQuizById(ctx context.Context, id int64) (Quiz, error) {
	row := q.db.QueryRowContext(ctx, getQuizById, id)
	var i Quiz
	err := row.Scan(
		&i.ID,
		&i.CreatedBy,
		&i.Name,
		&i.Description,
		&i.CreatedAt,
	)
	return i, err
}

const listQuiz = `-- name: ListQuiz :many
select id, created_by, name, description, created_at from quiz
`

func (q *Queries) ListQuiz(ctx context.Context) ([]Quiz, error) {
	rows, err := q.db.QueryContext(ctx, listQuiz)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Quiz{}
	for rows.Next() {
		var i Quiz
		if err := rows.Scan(
			&i.ID,
			&i.CreatedBy,
			&i.Name,
			&i.Description,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateQuiz = `-- name: UpdateQuiz :one
update quiz 
set name = $2, description = $3
where id = $1
returning id, created_by, name, description, created_at
`

type UpdateQuizParams struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (q *Queries) UpdateQuiz(ctx context.Context, arg UpdateQuizParams) (Quiz, error) {
	row := q.db.QueryRowContext(ctx, updateQuiz, arg.ID, arg.Name, arg.Description)
	var i Quiz
	err := row.Scan(
		&i.ID,
		&i.CreatedBy,
		&i.Name,
		&i.Description,
		&i.CreatedAt,
	)
	return i, err
}