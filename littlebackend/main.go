package main

import (
	"encoding/gob"
	"exam_go/config"
	"exam_go/middleware"
	"exam_go/model"
	"exam_go/router"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type usersession struct {
	UserId int    `json:"userId"`
	Number string `json:"number"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

func main() {
	gob.Register(usersession{})
	gin.SetMode(config.Config.AppMode)
	r := gin.Default()

	model.InitModel()
	config.InitSession(r)
	store := cookie.NewStore([]byte("snaosnca"))
	store.Options(sessions.Options{MaxAge: 1000, Path: "/"})
	r.Use(sessions.Sessions("SESSIONID", store))

	r.Use(middleware.Error)
	router.InitRouter(r)

	r.Run(":8088")
}
