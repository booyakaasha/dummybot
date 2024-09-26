package main

import (
	"log"

	"github.com/booyakaasha/bot/internal/app/commands"

	"github.com/booyakaasha/bot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("7739049195:AAExe5lJVPhXWgwN8HtXMaStWBn9Tthxjbg")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authoreized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	productService := product.NewService()

	commander := commands.NewCommander(bot, productService)

	for update := range updates {
		commander.HandleUpdate(update)
	}
}
