package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"net/mail"
	"net/smtp"
	"strings"
)

func (s *Service) SendMail(id string) error {
	data, err := s.dao.FindMailByID(id)
	if err != nil {
		return err
	}

	switch data.Status {
	case "WAIT", "FAILED":
	default:
		return fmt.Errorf("mail cannot be sent, status: " + data.Status)
	}

	tmpl, err := template.New(data.Template.ID).Parse(data.Template.ContentBody)
	if err != nil {
		return err
	}
	var body bytes.Buffer
	var fields map[string]interface{}
	if err = json.Unmarshal(data.Fields, &fields); err != nil {
		return err
	}
	if err = tmpl.Execute(&body, fields); err != nil {
		return err
	}

	from := mail.Address{Name: s.cfg.SMTP.Nicename, Address: s.cfg.SMTP.Username}
	msg := "From: " + from.String()
	msg += "\nTo: " + data.To
	msg += "\nSubject: " + data.Template.Subject
	switch strings.ToUpper(data.Template.ContentType) {
	case "HTML":
		msg += "\nMIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"
	default:
		msg += "\nMIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";"
	}
	msg += "\n\n" + body.String()

	return smtp.SendMail(s.cfg.SMTP.Hostname, s.smtpAuth, s.cfg.SMTP.Username, []string{data.To}, []byte(msg))
}
