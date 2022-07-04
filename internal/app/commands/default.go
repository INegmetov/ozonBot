package commands

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Default(inputMessage *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "You wrote: "+inputMessage.Text)

	c.bot.Send(msg)
}

func (c *Commander) HandleUpdate(update tgbotapi.Update) {

	defer func() {
		if panicValue := recover(); panicValue != nil {
			log.Printf("recover from panic: %v", panicValue)
		}
	}()

	if update.CallbackQuery != nil {
		msg := tgbotapi.NewMessage(
			update.CallbackQuery.Message.Chat.ID,
			"Data: "+update.CallbackQuery.Data,
		)
		c.bot.Send(msg)
		return
	}

	if update.Message == nil { // ignore any non-Message Updates
		return
	}

	switch update.Message.Command() {
	case "help":
		c.Help(update.Message)
	case "list":
		c.List(update.Message)
	case "get":
		c.Get(update.Message)
	default:
		c.Default(update.Message)
	}
}
