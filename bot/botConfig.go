package bot

import (
	"go-discord-bot/util"

	"github.com/bwmarrin/discordgo"
)

var BotId string

func GetDiscordBot() (*discordgo.Session, error) {

	token := util.GetEnv("discord_token", "<token>")

	goBot, err := discordgo.New("Bot " + token)

	return goBot, err
}
