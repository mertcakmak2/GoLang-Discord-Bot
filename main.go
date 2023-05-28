package main

import (
	"fmt"
	"go-discord-bot/bot"
)

func main() {

	discordBot, err := bot.GetDiscordBot()
	if err != nil {
		fmt.Println(err.Error())
	}

	botUser, err := discordBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
	}

	bot.BotId = botUser.ID

	discordBot.AddHandler(bot.MessageHandler)
	discordBot.AddHandler(bot.ReactionHandler)

	err = discordBot.Open()

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Discord Bot is running...")

	forever := make(chan bool)
	<-forever
}

// 👏 🔥
