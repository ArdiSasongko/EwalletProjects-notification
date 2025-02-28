// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: emailhistory.sql

package sqlc

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const insertEmailHistory = `-- name: InsertEmailHistory :exec
INSERT INTO notification_history (recipient, template_id, status, error_message) 
VALUES ($1, $2, $3, $4)
`

type InsertEmailHistoryParams struct {
	Recipient    string
	TemplateID   int32
	Status       string
	ErrorMessage pgtype.Text
}

func (q *Queries) InsertEmailHistory(ctx context.Context, arg InsertEmailHistoryParams) error {
	_, err := q.db.Exec(ctx, insertEmailHistory,
		arg.Recipient,
		arg.TemplateID,
		arg.Status,
		arg.ErrorMessage,
	)
	return err
}
