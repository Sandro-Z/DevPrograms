package main

import (
	"context"
	"log"
	"os"

	"git.ana/xjtuana/qqguildgo/bot"
	"git.ana/xjtuana/qqguildgo/plugin/antispam"
	"git.ana/xjtuana/qqguildgo/plugin/auditlog"
	_ "git.ana/xjtuana/qqguildgo/plugin/auth"
	"git.ana/xjtuana/qqguildgo/plugin/captcha"
	_ "git.ana/xjtuana/qqguildgo/plugin/ping"
)

func main() {
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatalln("No $BOT_TOKEN given.")
	}
	bot.NewClient("Bot " + token)
	bot.Use(auditlog.Middleware)
	bot.Use(antispam.Middleware)
	bot.Use(captcha.Middleware)
	if err := bot.Run(context.Background()); err != nil {
		log.Fatalln("Bot exited with error: ", err)
	}
	select {}
}
