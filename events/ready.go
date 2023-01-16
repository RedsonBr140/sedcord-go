package events

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func Ready(s *discordgo.Session, r *discordgo.Ready) {
	log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	status := fmt.Sprintf("sed at %d servers", len(r.Guilds))

	s.UpdateGameStatus(0, status)
}
