package main

import "github.com/bwmarrin/discordgo"

var integerOptionMinValueHour = 1.0
var integerOptionMinValueMinute = 0.0

var commands = []*discordgo.ApplicationCommand{
	{
		Name:        "reminder",
		Description: "set a reminder",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "hour",
				Description: "Hour of the reminder",
				Type:        discordgo.ApplicationCommandOptionInteger,
				MinValue:    &integerOptionMinValueHour,
				MaxValue:    12,
				Required:    true,
			},
			{
				Name:        "minute",
				Description: "Minute of the reminder",
				Type:        discordgo.ApplicationCommandOptionInteger,
				MinValue:    &integerOptionMinValueMinute,
				MaxValue:    59,
				Required:    true,
			},
			// {
			// 	Name:        "test",
			// 	Description: "test value",
			// 	Type:        discordgo.ApplicationCommandOptionInteger,
			// 	MinValue:    &integerOptionMinValueMinute,
			// 	MaxValue:    59,
			// 	Required:    true,
			// },
			{
				Name:        "am-pm",
				Description: "AM or PM for reminder",
				Type:        discordgo.ApplicationCommandOptionString,
				Required:    true,
				Choices: []*discordgo.ApplicationCommandOptionChoice{
					{
						Name:  "AM",
						Value: "AM",
					},
					{
						Name:  "PM",
						Value: "PM",
					},
				},
			},
		},
	},
}
