package service

import (
	"bytes"
	"context"
	"html/template"
	"log"
	"strings"

	"github.com/ArdiSasongko/EwalletProjects-notification/internal/mailer"
	"github.com/ArdiSasongko/EwalletProjects-notification/internal/model"
	"github.com/ArdiSasongko/EwalletProjects-notification/internal/storage/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type EmailService struct {
	q *sqlc.Queries
	e mailer.MailerSMTP
}

func (s *EmailService) SendEmail(ctx context.Context, payload model.NotificationRequest) error {
	emailTemplate, err := s.q.GetTemplateByName(ctx, payload.TemplateName)
	if err != nil {
		return err
	}

	tpl, err := template.New("emailtemplate").Parse(emailTemplate.Body)
	if err != nil {
		return err
	}

	var tmpl bytes.Buffer
	if err := tpl.Execute(&tmpl, payload.Placeholder); err != nil {
		return err
	}

	username := strings.Split(payload.Recipent, "@")
	log.Println(username[0])
	if err := s.e.Send(username[0], payload.Recipent, emailTemplate.Subject, tmpl.String()); err != nil {
		if err := s.q.InsertEmailHistory(ctx, sqlc.InsertEmailHistoryParams{
			Recipient:  payload.Recipent,
			TemplateID: emailTemplate.ID,
			Status:     "failed",
			ErrorMessage: pgtype.Text{
				String: err.Error(),
				Valid:  true,
			},
		}); err != nil {
			return err
		}
		return err
	}

	if err := s.q.InsertEmailHistory(ctx, sqlc.InsertEmailHistoryParams{
		Recipient:  payload.Recipent,
		TemplateID: emailTemplate.ID,
		Status:     "success",
	}); err != nil {
		return err
	}

	return nil
}
