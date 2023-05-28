package bot

import (
	"bytes"
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func MessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotId {
		return
	}

	var buf bytes.Buffer

	buf.WriteString(m.Author.Username)
	buf.WriteString(" ")
	buf.WriteString(m.Content)

	fmt.Println(buf.String())

	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "selam "+m.Author.Username)

		err := s.MessageReactionAdd(m.ChannelID, m.ID, "🔥")
		if err != nil {
			fmt.Println(err.Error)
		}
	}
}

func ReactionHandler(s *discordgo.Session, m *discordgo.MessageReactionAdd) {
	fmt.Println("Reacted to message ", m.Emoji.Name, m.Emoji.ID)
	// msg, err := s.ChannelMessage(m.ChannelID, m.MessageID)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(msg.Content)
}
