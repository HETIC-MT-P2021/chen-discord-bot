package command

import (
	"github.com/HETIC-MT-P2021/chen-discord-bot/discord"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/HETIC-MT-P2021/chen-discord-bot/pokeapi"
	"github.com/bwmarrin/discordgo"
)

func formatStats(p pokeapi.Pokemon) (titleStats string, scaleStats string) {
	var title []string
	var stats []string
	for _, s := range p.Stats() {
		scaleMax := 15
		scale := strings.Repeat("⚪", scaleMax)
		calc := int(math.Floor(float64(s.Value) / float64(225) * float64(scaleMax)))
		scale = strings.Replace(scale, "⚪", "🔵", calc)
		stats = append(stats, scale)
		title = append(title, s.Name)
	}
	return strings.Join(title, "\n"), strings.Join(stats, "\n")
}

func cardEmbed(p pokeapi.Pokemon) *discordgo.MessageEmbed {
	titleStats, scaleStats := formatStats(p)
	return &discordgo.MessageEmbed{
		Type:        "rich",
		Title:       p.Title(),
		Description: p.Description(),
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Height",
				Value:  strconv.Itoa(p.Pokemon.Height),
				Inline: true,
			},
			{
				Name:   "Weight",
				Value:  strconv.Itoa(p.Pokemon.Weight),
				Inline: true,
			},
			{
				Name:   "Category",
				Value:  p.Category(),
				Inline: true,
			},
			{
				Name:   "Type",
				Value:  p.Types(),
				Inline: false,
			},
			{
				Name:   "Stats",
				Value:  titleStats,
				Inline: true,
			},
			{
				Name:   "\u200b",
				Value:  scaleStats,
				Inline: true,
			},
		},
		Color:     discord.C_RED,
		Thumbnail: &discordgo.MessageEmbedThumbnail{URL: p.Image()},
		URL:       p.Link(),
		Author:    getAuthorPokedex(),
		Timestamp: time.Now().Format(time.RFC3339),
	}
}
