package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	token := ""
	appURL := ""

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(bot.Self.FirstName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			switch update.Message.Command() {
				case "start":
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет! Нажмите кнопку ниже, чтобы открыть мини-приложение:")
					btn := tgbotapi.NewInlineKeyboardButtonURL("Открыть приложение", appURL)
					kb := tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(btn))
					msg.ReplyMarkup = kb
					bot.Send(msg)
				default:
					bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Команда не распознана. Используйте /start"))
			}
		}
	}
}