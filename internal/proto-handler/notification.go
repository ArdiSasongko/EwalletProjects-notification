package protohandler

import (
	"context"
	"errors"

	"github.com/ArdiSasongko/EwalletProjects-notification/internal/config/logger"
	"github.com/ArdiSasongko/EwalletProjects-notification/internal/mailer"
	"github.com/ArdiSasongko/EwalletProjects-notification/internal/model"
	"github.com/ArdiSasongko/EwalletProjects-notification/internal/proto/notification"
	"github.com/ArdiSasongko/EwalletProjects-notification/internal/service"
	"github.com/ArdiSasongko/EwalletProjects-notification/internal/storage/sqlc"
)

var log = logger.NewLogger()

type NotifEmailInterface interface {
	SendNotification(ctx context.Context, req *notification.SendNotificationRequest) (*notification.SendNotificationResponse, error)
}

type NotifEmail struct {
	notification.UnimplementedNotificationServiceServer
	e service.Service
}

func NewNotifEmail(q *sqlc.Queries, e mailer.MailerSMTP) NotifEmailInterface {
	email := service.NewService(q, e)
	return &NotifEmail{
		e: email,
	}
}

func (n *NotifEmail) SendNotification(ctx context.Context, req *notification.SendNotificationRequest) (*notification.SendNotificationResponse, error) {
	notifReq := model.NotificationRequest{
		TemplateName: req.TemplateName,
		Recipent:     req.Recipient, // Fixed field name
		Placeholder:  req.Placeholder,
	}

	if err := notifReq.Validate(); err != nil {
		errs := errors.New("validate error")
		log.WithError(errs).Error(err.Error())
		return &notification.SendNotificationResponse{
			Message: err.Error(),
		}, nil
	}

	if err := n.e.Email.SendEmail(ctx, notifReq); err != nil {
		log.WithError(err).Error(err.Error())
		return &notification.SendNotificationResponse{
			Message: "failed send email",
		}, nil
	}

	return &notification.SendNotificationResponse{
		Message: "success",
	}, nil
}
