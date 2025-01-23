package service

import (
	"context"

	"github.com/ArdiSasongko/EwalletProjects-notification/internal/mailer"
	"github.com/ArdiSasongko/EwalletProjects-notification/internal/model"
	"github.com/ArdiSasongko/EwalletProjects-notification/internal/storage/sqlc"
)

type Service struct {
	Email interface {
		SendEmail(context.Context, model.NotificationRequest) error
	}
}

func NewService(q *sqlc.Queries, e mailer.MailerSMTP) Service {
	return Service{
		Email: &EmailService{
			q: q,
			e: e,
		},
	}
}
