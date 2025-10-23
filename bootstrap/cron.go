package bootstrap

import (
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"venturan/app/services"
	"venturan/global"
)

func InitializeCron() {
	c := cron.New()
	_, err := c.AddFunc("0 0 2 * * ?", func() {
		err := services.UserService.DeleteResgin()
		if err != nil {
			global.App.Log.Error("clear resgin employee account failed, err:", zap.Any("err", err))
		}
	})
	if err != nil {
		global.App.Log.Error("Cron task add failed, err:", zap.Any("err", err))
	}
	c.Start()
}
