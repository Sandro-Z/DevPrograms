package http

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"git.ana/xjtuana/api-micro-mail/config"
	"git.ana/xjtuana/api-micro-mail/service"
)

type Server struct {
	cfg *config.Config
	svc *service.Service

	e   *gin.Engine
	srv *http.Server
}

func NewServer(cfg *config.Config) *Server {
	s := &Server{
		cfg: cfg,
		svc: service.New(cfg),
	}
	s.e = gin.Default()
	s.srv = &http.Server{
		Addr:    cfg.Server.Addr,
		Handler: s.e,
	}
	return s
}

func (s *Server) ListenAndServe() error {
	return s.srv.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}

func (s *Server) SetRoutes() {
	s.HandlePing()
	s.HandlePrivate()
	s.HandlePublic()
}

func (s *Server) HandlePing() {
	s.e.GET("/ping", func(c *gin.Context) {
		if err := s.svc.Ping(c); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		c.Status(http.StatusOK)
	})
}

func (s *Server) HandlePrivate() {
	mails := s.e.Group("/x/api/micro/mails") // verify permission
	{
		mails.GET("", s.getMailPage)        // select mails
		mails.GET("/:id", s.getMail)        // select mail
		mails.POST("", s.createMail)        // create mail
		mails.PATCH("/:id", s.updateMail)   // update mail
		mails.DELETE("/:id", s.deleteMail)  // delete mail
		mails.POST("/sendMail", s.sendMail) // send mail

		templates := mails.Group("/templates") // verify permission
		{
			templates.GET("", s.getTemplatePage)       // select templates
			templates.GET("/:id", s.getTemplate)       // select template
			templates.POST("", s.createTemplate)       // create template
			templates.PATCH("/:id", s.updateTemplate)  // update template
			templates.DELETE("/:id", s.deleteTemplate) // delete template
		}
	}
}

func (s *Server) HandlePublic() {
	mails := s.e.Group("/mails") // verify scope
	{
		mails.GET("", hasScope("MicroMail.Read.All"), s.getMailPage)            // select mails
		mails.GET("/:id", hasScope("MicroMail.Read.All"), s.getMail)            // select mail
		mails.POST("", hasScope("MicroMail.ReadWrite.All"), s.createMail)       // create mail
		mails.PATCH("/:id", hasScope("MicroMail.ReadWrite.All"), s.updateMail)  // update mail
		mails.DELETE("/:id", hasScope("MicroMail.ReadWrite.All"), s.deleteMail) // delete mail
		mails.POST("/sendMail", hasScope("MicroMail.Send.All"), s.sendMail)     // send mail

		templates := mails.Group("/templates") // verify scope
		{
			templates.GET("", hasScope("MicroMailTemplate.*.All"), s.getTemplatePage)               // select templates
			templates.GET("/:id", hasScope("MicroMailTemplate.*.All"), s.getTemplate)               // select template
			templates.POST("", hasScope("MicroMailTemplate.ReadWrite.All"), s.createTemplate)       // create template
			templates.PATCH("/:id", hasScope("MicroMailTemplate.ReadWrite.All"), s.updateTemplate)  // update template
			templates.DELETE("/:id", hasScope("MicroMailTemplate.ReadWrite.All"), s.deleteTemplate) // delete template
		}
	}
}

func hasScope(scope ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
