package dto

import (
	"encoding/json"
	"strings"

	"git.ana/xjtuana/api-micro-mail/model"
)

type MailsGetMailRequest struct{}

type MailsGetMailResponse struct {
	ID string `json:"id"`
	To string `json:"to"`

	Fields map[string]string `json:"fields"`
	Status string            `json:"status"`

	TmplID string `json:"tmpl_id"`
	UserID string `json:"user_id"`
}

func (resp *MailsGetMailResponse) Bind(data *model.Mail) error {
	resp.ID = data.ID
	resp.To = data.To
	if err := json.Unmarshal(data.Fields, &resp.Fields); err != nil {
		return err
	}
	resp.TmplID = data.TmplID
	resp.UserID = data.UserID
	return nil
}

type MailsCreateMailRequest struct {
	To string `json:"to"`

	Fields map[string]string `json:"fields"`
	Status string            `json:"status"`

	TmplID string `json:"tmpl_id"`
	UserID string `json:"user_id"`
}

func (req MailsCreateMailRequest) Bind(data *model.Mail) error {
	data.To = req.To
	data.Fields, _ = json.Marshal(req.Fields)
	switch strings.ToUpper(req.Status) {
	case "DRAFT":
		data.Status = "DRAFT"
	case "WAIT":
		data.Status = "WAIT"
	default:
		data.Status = "UNKNOWN"
	}
	data.TmplID = req.TmplID
	data.UserID = req.UserID
	return nil
}

type MailsCreateMailResponse struct {
	ID string `json:"id"`
}

type MailsUpdateMailRequest struct {
	ID string `json:"id"`
	To string `json:"to"`

	Fields map[string]string `json:"fields"`
	Status string            `json:"status"`

	TmplID string `json:"tmpl_id"`
	UserID string `json:"user_id"`
}

func (req MailsUpdateMailRequest) Bind(data *model.Mail) error {
	data.ID = req.ID
	data.To = req.To
	data.Fields, _ = json.Marshal(req.Fields)
	switch strings.ToUpper(req.Status) {
	case "DRAFT":
		data.Status = "DRAFT"
	case "WAIT":
		data.Status = "WAIT"
	default:
		data.Status = "UNKNOWN"
	}
	data.TmplID = req.TmplID
	data.UserID = req.UserID
	return nil
}

type MailsUpdateMailResponse struct {
	ID string `json:"id"`
}

type MailsSendMailRequest struct {
	ID string `json:"id"`
}

type MailsSendMailResponse struct{}
