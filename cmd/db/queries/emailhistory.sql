-- name: InsertEmailHistory :exec
INSERT INTO notification_history (recipient, template_id, status, error_message) 
VALUES ($1, $2, $3, $4);

