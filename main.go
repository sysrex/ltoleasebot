package main

import (
	"log"
	"time"

	"github.com/sysrex/ltoleasebot/config"
	"github.com/sysrex/ltoleasebot/handler"

	tb "gopkg.in/tucnak/telebot.v2"
	"go.uber.org/zap"
)



func main() {

	loggerConfig := zap.NewDevelopmentConfig()
	loggerConfig.DisableStacktrace = true
	loggerConfig.DisableCaller = true
	zapLogger, _ := loggerConfig.Build()
	defer zapLogger.Sync()
	logger := zapLogger.Sugar()

	b, err := tb.NewBot(tb.Settings{
		Token: config.LoadConfig().Token
		PollerL &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	for k, v := range handler.LoadHandler(b) {
		b.Handle(k, v)
		log.Println(k + "✅ Loaded!")
	}

	b.Start()
}