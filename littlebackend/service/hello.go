package service

import (
	"fmt"
	"time"
)

type HelloService struct {
}

func (h HelloService) Hello() string {
	return fmt.Sprintln("你好谢谢小笼包再见")
}

func (h HelloService) HelloTime(date time.Time) string {
	return fmt.Sprintf("tomorrow is %v", date.Format("2006-01-02"))
}
