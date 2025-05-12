/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"
	"github.com/spf13/cobra"
	"log"
	telebot "gopkg.in/telebot.v4"
	"os"
	"time"
)

var (
	TeleToken = os.Getenv("TELE_TOKEN")
)

// kbotCmd represents the kbot command
var kbotCmd = &cobra.Command{
	Use:   "kbot",
	Aliases: []string{"start"},
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("kbot %s started", appVersion)
		pref := telebot.Settings{
			URL:    "",
			Token:  TeleToken,
			Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
		}

		kbot, err := telebot.NewBot(pref)
		if err != nil {
			if TeleToken == "" {
				fmt.Println("Please set the TELE_TOKEN environment variable.")
				return
			}
			log.Fatal(err)
			return
		}

		kbot.Handle(telebot.OnText, func(ctx telebot.Context) error {
		    var text = ctx.Text();
			log.Println(ctx.Message().Payload, text)

			switch strings.ToLower(text) {
			    case "hello":
                    ctx.Send("Hello, how can I help you?")
                case "bye":
                    ctx.Send("Bye! Have a great day!")
                default:
                    ctx.Send("Sorry, I don't understand that command.")
            }

			return err
		})

		kbot.Start()
	},
}

func init() {
	rootCmd.AddCommand(kbotCmd)
}
