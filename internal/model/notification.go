package model

import "github.com/go-playground/validator/v10"

var Validate *validator.Validate

func init() {
	Validate = validator.New(validator.WithRequiredStructEnabled())
}

type NotificationRequest struct {
	TemplateName string            `json:"template_name" validate:"required"`
	Recipent     string            `json:"recipent" validate:"required"`
	Placeholder  map[string]string `json:"placeholder"`
}

func (u NotificationRequest) Validate() error {
	return Validate.Struct(u)
}

type NotificationHistory struct {
	Recipent   string `json:"recipent"`
	TemplateID int32  `json:"template_id"`
	Status     string `json:"status"`
}
