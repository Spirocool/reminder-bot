package main

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"gorm.io/gorm"
)

var handlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate, db *gorm.DB){
	//TODO - add message to reminder
	"reminder": func(s *discordgo.Session, i *discordgo.InteractionCreate, db *gorm.DB) {
		options := i.ApplicationCommandData().Options
		optionsMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
		for _, opt := range options {
			optionsMap[opt.Name] = opt
		}

		res := db.Create(&TimeReminder{
			Author:    i.Member.User.ID,
			ChannelID: i.ChannelID,
			Hour:      int(optionsMap["hour"].IntValue()),
			Minute:    int(optionsMap["minute"].IntValue()),
		})

		if res.Error != nil {
			log.Fatalf("Error creating record in db: %s", res.Error)
		}

		var test TimeReminder
		res = db.Last(&test)

		if res.Error != nil {
			log.Fatalf("Error querying db: %s", res.Error)
		}

		fmt.Println("From DB")
		fmt.Printf("Author: %s\nChannelID: %s\nHour: %d\nMinute: %d", test.Author, test.ChannelID, test.Hour, test.Minute)

		margs := make([]interface{}, 0, len(options))
		msg := "## Reminder Data\n"

		if option, ok := optionsMap["hour"]; ok {
			margs = append(margs, option.IntValue())
			msg += "Hour: %d\n"
		}

		if option, ok := optionsMap["minute"]; ok {
			margs = append(margs, option.IntValue())
			msg += "Minute: %d\n"
		}

		if option, ok := optionsMap["am-pm"]; ok {
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
