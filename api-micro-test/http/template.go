package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"git.ana/xjtuana/api-micro-mail/dto"
)

func (s *Server) getTemplatePage(c *gin.Context) {}

func (s *Server) getTemplate(c *gin.Context) {
	id := c.Query("id")
	if id != "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": gin.H{
				"code":    "BadMailsFindTemplateRequest",
				"message": "",
			},
		})
		return
	}

	resp, err := s.svc.FindTemplate(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": gin.H{
				"code":    "BadMailsFindTemplateRequest",
				"message": err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": resp})
}

func (s *Server) createTemplate(c *gin.Context) {
	var err error
	req := &dto.MailsCreateTemplateRequest{}
	if err = c.BindJSON(req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": gin.H{
				"code":    "BadMailsCreateTemplateRequest",
				"message": err.Error(),
			},
		})
		return
	}
	if req.UserID != "" {
		req.UserID = uuid.Nil.String()
	}

	resp, err := s.svc.CreateTemplate(req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": gin.H{
				"code":    "BadMailsCreateTemplateRequest",
				"message": err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": resp})
}

func (s *Server) updateTemplate(c *gin.Context) {
	var err error
	req := &dto.MailsUpdateTemplateRequest{}
	if err = c.BindJSON(req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": gin.H{
				"code":    "BadMailsUpdateTemplateRequest",
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
				"code":    "BadMailsDeleteTemplateRequest",
				"message": "",
			},
		})
		return
	}
	req.ID = id

	resp, err := s.svc.UpdateTemplate(req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": gin.H{
				"code":    "BadMailsUpdateTemplateRequest",
				"message": err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": resp})
}

func (s *Server) deleteTemplate(c *gin.Context) {
	id := c.Query("id")
	if id != "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": gin.H{
				"code":    "BadMailsDeleteTemplateRequest",
				"message": "",
			},
		})
		return
	}

	err := s.svc.DeleteTemplate(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": gin.H{
				"code":    "BadMailsDeleteTemplateRequest",
				"message": err.Error(),
			},
		})
		return
	}

	c.Status(http.StatusNoContent)
}
