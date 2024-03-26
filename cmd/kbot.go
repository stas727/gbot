/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	"gopkg.in/telebot.v3"
)

var (
	//TeleToken bot
	TeleToken = os.Getenv("TELE_TOKEN")
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
		kbot, err := telebot.NewBot(telebot.Settings{
			URL:    "",
			Token:  TeleToken,
			Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
		})
		if err != nil {
			fmt.Println("Telebot error. ", err.Error())
			return
		}

		kbot.Handle(telebot.OnText, func(m telebot.Context) error {
			log.Println(m.Message().Payload, m.Text())
			payload := m.Message().Payload

			switch payload {
			case "hello":
				err := m.Send("Hello, I am Bot.")
				if err != nil {
					fmt.Println("Telebot send error. ", err.Error())

					return err
				}
			case "version":
				err := m.Send(fmt.Sprintf("Version bof Bot is %s!", appVersion))
				if err != nil {
					fmt.Println("Telebot send error. ", err.Error())

					return err
				}
			}

			return err
		})

		kbot.Start()
	},
}

func init() {
	rootCmd.AddCommand(gbotCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// gbotCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gbotCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
