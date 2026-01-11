package commands

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/bufferof/bot/response"
	"github.com/bwmarrin/discordgo"
)

var Slot = &discordgo.ApplicationCommand{
	Name:        "슬롯머신",
	Description: "같은 문양이 3개 연속해서 나오면 상금을 얻습니다",
}

var symbols = []string{"🍒", "🍋", "🔔", "⭐", "💎"}

func make_randomGrid() [][]string {
	rand.Seed(time.Now().UnixNano())

	grid := make([][]string, 3)
	for i := 0; i < 3; i++ {
		grid[i] = make([]string, 3)
		for j := 0; j < 3; j++ {
			grid[i][j] = symbols[rand.Intn(len(symbols))]
		}
	}

	return grid
}

func build_slotMachine(title string, grid [][]string, result string, color int) *discordgo.MessageEmbed {
	view := fmt.Sprintf(
		"   %s | %s | %s\n"+
			"➤ %s | %s | %s ◀\n"+
			"   %s | %s | %s\n\n"+
			"%s",
		grid[0][0], grid[0][1], grid[0][2],
		grid[1][0], grid[1][1], grid[1][2],
		grid[2][0], grid[2][1], grid[2][2],
		result,
	)

	return response.Make_embed(title, "```"+view+"```", color)
}

func judge(grid [][]string) (string, int) {
	a, b, c := grid[1][0], grid[1][1], grid[1][2]

	if a == b && b == c {
		if a == "💎" {
			return "💎 JACKPOT!!!", 0xaa00ff
		}
		return "🎉 WIN!", 0x00ff00
	}

	return "💣 FAILED...", 0xff4444
}

func SlotMachineHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	response.Reply(
		s,
		i,
		&discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				response.Make_embed("슬롯머신", "같은 문양이 3개 연속해서 나오면 상금을 얻습니다", 0x4285f4),
			},
			Components: []discordgo.MessageComponent{
				response.Make_btn("✅", "slot_confirm", discordgo.PrimaryButton),
			},
		},
	)
}

func SlotRunHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredMessageUpdate,
	})
	msg := i.Message
	grid := make_randomGrid()

	response.EditMessageComplex(
		s,
		msg.ChannelID,
		msg.ID,
		[]*discordgo.MessageEmbed{
			build_slotMachine("슬롯머신", grid, "아래 버튼을 눌러 슬롯머신을 돌리세요", 0xffd900),
		},
		[]discordgo.MessageComponent{
			response.Make_btn("🎰", "slot_run", discordgo.PrimaryButton),
		},
	)
}

func Slot_run(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseDeferredMessageUpdate,
	})

	msg := i.Message

	go func() {
		var grid [][]string

		for step := 0; step < 3; step++ {
			time.Sleep(time.Second * 1)
			grid = make_randomGrid()

			response.EditMessageComplex(
				s,
				msg.ChannelID,
				msg.ID,
				[]*discordgo.MessageEmbed{
					build_slotMachine("회전 중...", grid, "", 0xffcc00),
				},
				[]discordgo.MessageComponent{},
			)
		}

		result, color := judge(grid)

		time.Sleep(time.Second * 1)
		response.EditMessageComplex(s,
			msg.ChannelID,
			msg.ID,
			[]*discordgo.MessageEmbed{
				build_slotMachine(result, grid, "", color),
			},
			[]discordgo.MessageComponent{
				response.Make_btn("🎰", "slot_run", discordgo.PrimaryButton),
			},
		)
	}()
}
