package handler

import (
	"github.com/xsre/pretulcorect/commands"
	tb "gopkg.in/tucnak/telebot.v2"
)

func LoadHandler(b *tb.Bot) map[string]func(m *tb.Message) {
	commandMap := make(map[string]func(m *tb.Message))

	commandMap["/price"] = func(m *tb.Message) {
		res, _ := commands.GetPrice()
		b.Send(m.Chat, "Pretul la Helium acum este: U$S "+res)
	}

	return commandMap
}
