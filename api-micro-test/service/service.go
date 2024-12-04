package service

import (
	"context"
	"net/smtp"
	"sync"

	"git.ana/xjtuana/api-micro-mail/config"
	"git.ana/xjtuana/api-micro-mail/dao"
	"git.ana/xjtuana/api-micro-mail/util"
)

type Service struct {
	cfg *config.Config
	dao *dao.DAO

	mutex    sync.Mutex // protects following
	smtpAuth smtp.Auth
	closing  bool // user has called Close
	shutdown bool // server has told us to stop
}

func New(cfg *config.Config) *Service {
	svc := &Service{cfg: cfg, dao: dao.New(cfg)}
	svc.init()
	return svc
}

func (s *Service) init() {
	s.mutex.Lock()
	s.smtpAuth = util.NewLoginAuth(s.cfg.SMTP.Username, s.cfg.SMTP.Password)
	s.mutex.Unlock()
}

func (s *Service) Close() error {
	s.mutex.Lock()
	if s.closing {
		s.mutex.Unlock()
		return ErrShutdown
	}
	s.closing = true
	s.mutex.Unlock()
	return s.dao.Close()
}

func (s *Service) Ping(ctx context.Context) error {
	return s.dao.Ping(ctx)
}
