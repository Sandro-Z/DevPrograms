package service

import (
	"git.ana/xjtuana/api-micro-mail/dto"
	"git.ana/xjtuana/api-micro-mail/model"
)

func (s *Service) FindTemplate(id string) (*dto.MailsGetTemplateResponse, error) {
	data, err := s.dao.FindTemplateByID(id)
	if err != nil {
		return nil, err
	}

	resp := &dto.MailsGetTemplateResponse{}
	if err := resp.Bind(data); err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *Service) CreateTemplate(req *dto.MailsCreateTemplateRequest) (*dto.MailsCreateTemplateResponse, error) {
	data := &model.Template{}
	if err := req.Bind(data); err != nil {
		return nil, err
	}

	if err := s.dao.InsertTemplate(data); err != nil {
		return nil, err
	}

	return &dto.MailsCreateTemplateResponse{ID: data.ID}, nil
}

func (s *Service) UpdateTemplate(req *dto.MailsUpdateTemplateRequest) (*dto.MailsUpdateTemplateResponse, error) {
	data := &model.Template{}
	if err := req.Bind(data); err != nil {
		return nil, err
	}

	if err := s.dao.InsertTemplate(data); err != nil {
		return nil, err
	}

	return &dto.MailsUpdateTemplateResponse{ID: data.ID}, nil
}

func (s *Service) DeleteTemplate(id string) error {
	return s.dao.DeleteTemplate(id)
}
