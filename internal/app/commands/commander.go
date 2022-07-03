package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/inegmetov/ozonBot/internal/service/product"
)

var registeredCommander = map[string]func(c *Commander, msg *tgbotapi.Message){}

type Commander struct {
	bot            *tgbotapi.BotAPI
	productService *product.Service
}

func NewCommander(
	bot *tgbotapi.BotAPI,
	productService *product.Service,
) *Commander {
	return &Commander{
		bot:            bot,
		productService: productService,
	}
}

func (c *Commander) HendleUpdate(update tgbotapi.Update) {
	if update.Message != nil { // If we got a message

		command, ok := registeredCommander[update.Message.Command()]

		if ok {
			command(c, update.Message)
		} else {
			c.Default(update.Message)
		}
	}
}
