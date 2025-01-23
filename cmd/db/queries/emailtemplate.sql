-- name: GetTemplateByName :one
SELECT id, template_name, subject, body FROM notification_template WHERE template_name = $1;


