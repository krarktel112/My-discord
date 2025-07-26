package main

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

func main() {
	// Initialize Discord session with your bot token
	dg, err := discordgo.New("Bot " + os.Getenv("DISCORD_BOT_TOKEN"))
	if err != nil {
		fmt.Println("Error creating Discord session:", err)
		return
	}

	// Register a handler for MessageCreate events
	dg.AddHandler(messageCreate)

	// Open a websocket connection to Discord
	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening connection:", err)
		return
	}

	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	<-make(chan struct{}) // Keep the program running until interrupted
	return
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore messages from the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Example: Forward a DM to a specific user (replace with actual logic)
	if m.ChannelID == "YOUR_SOURCE_DM_CHANNEL_ID" { // Replace with the ID of the DM you want to forward from
		targetUserID := "YOUR_TARGET_USER_ID" // Replace with the ID of the user to forward to

		// Create or get the DM channel for the target user
		dmChannel, err := s.UserChannelCreate(targetUserID)
		if err != nil {
			fmt.Println("Error creating DM channel:", err)
			return
		}

		// Send the message content to the target DM channel
		_, err = s.ChannelMessageSend(dmChannel.ID, "Forwarded message: " + m.Content)
		if err != nil {
			fmt.Println("Error sending forwarded message:", err)
		}
	}
}
