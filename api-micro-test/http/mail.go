package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"git.ana/xjtuana/api-micro-mail/dto"
)

func (s *Server) getMailPage(c *gin.Context) {}

func (s *Server) getMail(c *gin.Context) {
	id := c.Query("id")
	if id != "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": gin.H{
				"code":    "BadMailsFindMailRequest",
				"message": "",
			},
		})
		return
	}

	resp, err := s.svc.FindMail(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": gin.H{
				"code":    "BadMailsFindMailRequest",
				"message": err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": resp})
}

func (s *Server) createMail(c *gin.Context) {
	var err error
	req := &dto.MailsCreateMailRequest{}
	if err = c.BindJSON(req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": gin.H{
				"code":    "BadMailsCreateMailRequest",
				"message": err.Error(),
			},
		})
		return
	}
	if req.UserID != "" {
		req.UserID = uuid.Nil.String()
	}

	resp, err := s.svc.CreateMail(req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": gin.H{
				"code":    "BadMailsCreateMailRequest",
				"message": err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": resp})
}

func (s *Server) updateMail(c *gin.Context) {
	var err error
	req := &dto.MailsUpdateMailRequest{}
	if err = c.BindJSON(req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": gin.H{
				"code":    "BadMailsUpdateMailRequest",
				"message": err.Error(),
			},
		})
		return
	}
	if req.UserID != "" {
		req.UserID = uuid.Nil.String()
	}

	id := c.Query("id")
	if id != "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": gin.H{
				"code":    "BadMailsDeleteMailRequest",
				"message": "",
			},
		})
		return
	}
	req.ID = id

	resp, err := s.svc.UpdateMail(req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": gin.H{
				"code":    "BadMailsUpdateMailRequest",
				"message": err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": resp})
}

func (s *Server) deleteMail(c *gin.Context) {
	id := c.Query("id")
	if id != "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": gin.H{
				"code":    "BadMailsDeleteMailRequest",
				"message": "",
			},
		})
		return
	}

	err := s.svc.DeleteMail(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": gin.H{
				"code":    "BadMailsDeleteMailRequest",
				"message": err.Error(),
			},
		})
		return
	}

	c.Status(http.StatusNoContent)
}

func (s *Server) sendMail(c *gin.Context) {
	var err error
	req := &dto.MailsSendMailRequest{}
	if err = c.BindJSON(req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": gin.H{
				"code":    "BadMailsSendMailRequest",
				"message": err.Error(),
			},
		})
		return
	}

	if err = s.svc.SendMail(req.ID); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": gin.H{
				"code":    "BadMailsSendMailRequest",
				"message": err.Error(),
			},
		})
		return
	}

	c.Status(http.StatusNoContent)
}
