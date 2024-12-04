package service

import (
	"git.ana/xjtuana/api-micro-mail/dto"
	"git.ana/xjtuana/api-micro-mail/model"
)

func (s *Service) FindMail(id string) (*dto.MailsGetMailResponse, error) {
	data, err := s.dao.FindMailByID(id)
	if err != nil {
		return nil, err
	}

	resp := &dto.MailsGetMailResponse{}
	if err := resp.Bind(data); err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *Service) CreateMail(req *dto.MailsCreateMailRequest) (*dto.MailsCreateMailResponse, error) {
	data := &model.Mail{}
	if err := req.Bind(data); err != nil {
		return nil, err
	}

	if err := s.dao.InsertMail(data); err != nil {
		return nil, err
	}

	return &dto.MailsCreateMailResponse{ID: data.ID}, nil
}

func (s *Service) UpdateMail(req *dto.MailsUpdateMailRequest) (*dto.MailsUpdateMailResponse, error) {
	data := &model.Mail{}
	if err := req.Bind(data); err != nil {
		return nil, err
	}

	if err := s.dao.InsertMail(data); err != nil {
		return nil, err
	}

	return &dto.MailsUpdateMailResponse{ID: data.ID}, nil
}

func (s *Service) DeleteMail(id string) error {
	return s.dao.DeleteMail(id)
}
