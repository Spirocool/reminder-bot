package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var handlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"reminder": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		options := i.ApplicationCommandData().Options
		optionsMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
		for _, opt := range options {
			optionsMap[opt.Name] = opt
		}

		margs := make([]interface{}, 0, len(options))
		msg := "Reminder Data: \n"

		if option, ok := optionsMap["hour"]; ok {
			margs = append(margs, option.IntValue())
			msg += "Hour: %d\n"
		}

		if option, ok := optionsMap["minute"]; ok {
			margs = append(margs, option.IntValue())
			msg += "Minute: %d\n"
		}

		if option, ok := optionsMap["AM/PM"]; ok {
			margs = append(margs, option.StringValue())
			msg += "AM/PM: %s"
		}

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf(
					msg,
					margs...,
				),
			},
		})
	},
}
