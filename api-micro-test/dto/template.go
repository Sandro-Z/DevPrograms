package dto

import (
	"git.ana/xjtuana/api-micro-mail/model"
)

type MailsGetTemplateRequest struct{}

type MailsGetTemplateResponse struct {
	ID string `json:"id"`

	Subject string `json:"subject"`
	Body    string `json:"body"`
	Type    string `json:"type"`

	UserID string `json:"user_id"`
}

func (resp *MailsGetTemplateResponse) Bind(data *model.Template) error {
	resp.ID = data.ID
	resp.Subject = data.Subject
	resp.Body = data.ContentBody
	resp.Type = data.ContentType
	resp.UserID = data.UserID
	return nil
}

type MailsCreateTemplateRequest struct {
	Subject string `json:"subject"`
	Body    string `json:"body"`
	Type    string `json:"type"`

	UserID string `json:"user_id"`
}

func (req MailsCreateTemplateRequest) Bind(data *model.Template) error {
	data.Subject = req.Subject
	data.ContentBody = req.Body
	data.ContentType = req.Type
	data.UserID = req.UserID
	return nil
}

type MailsCreateTemplateResponse struct {
	ID string `json:"id"`
}

type MailsUpdateTemplateRequest struct {
	ID string `json:"id"`

	Subject string `json:"subject"`
	Body    string `json:"body"`
	Type    string `json:"type"`

	UserID string `json:"user_id"`
}

func (req MailsUpdateTemplateRequest) Bind(data *model.Template) error {
	data.ID = req.ID
	data.Subject = req.Subject
	data.ContentBody = req.Body
	data.ContentType = req.Type
	data.UserID = req.UserID
	return nil
}

type MailsUpdateTemplateResponse struct {
	ID string `json:"id"`
}
