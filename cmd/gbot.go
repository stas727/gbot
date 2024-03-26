/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	service "github.com/stas727/gbot/cmd/services"
	storage "github.com/stas727/gbot/cmd/storages"
	"os"
	"time"

	"github.com/spf13/cobra"
)

const MessageHi = "Цей бот на базі штучного інтелекту володіє великою базою даних і здатен аналізувати великі обсяги інформації у реальному часі. Він може надати користувачам різноманітну інформацію і рекомендації залежно від їх потреб і запитань.\n\nКорисність: AI Vision може бути корисним для студентів, дослідників, бізнесменів, маркетологів та будь-яких інших людей, які потребують швидкого і точного аналізу даних. Він може допомогти у прийнятті рішень, проведенні досліджень, аналітиці даних, а також в автоматичному оновленні інформації."
const MessageVersion = "Hello, I am AI %s!"

var (
	TeleToken = os.Getenv("TELE_TOKEN")
	AiToken   = os.Getenv("AI_TOKEN")
)

var (
	storages = storage.NewStorages()
	services = service.NewServices(storages)
)

// gbotCmd represents the gbot command
var gbotCmd = &cobra.Command{
	Use:     "gbot",
	Aliases: []string{"start"},
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Gbot %s started", appVersion)
		ctx := context.Background()

		//new AI Client
		client := services.AIService.NewClient(ctx, AiToken)

		//new Bot
		newBot, err := services.BotService.NewBot(ctx, TeleToken, 10*time.Second, "")
		if err != nil {
			fmt.Println("Telebot error. ", err.Error())
			return
		}

		//conversation
		e := services.BotService.Handle(ctx, func(message string) string {
			switch message {
			case "/start", "hello", "Hello", "hi", "hey", "Hi", "Hey", "привіт", "Привіт":
				return MessageHi
			case "version":
				return fmt.Sprintf(MessageVersion, appVersion)
			default:
				response, err := services.AIService.Response(ctx, message, client)

				if err != nil {
					fmt.Println("Error AI response. Err : " + err.Error())
					return ""
				}
				return *response
			}
		}, newBot)

		if e != nil {
			return
		}

		newBot.Client.Start()
	},
}

func init() {
	rootCmd.AddCommand(gbotCmd)
}
