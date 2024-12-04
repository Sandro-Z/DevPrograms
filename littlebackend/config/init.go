package config

import (
	"exam_go/service/validator"
)

func init() {
	initConfig()
	initLogger()
	validator.InitValidator()
}
