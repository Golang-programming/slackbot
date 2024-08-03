package main

import (
	"context"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/shomali11/slacker"
)

func main() {
	godotenv.Load()
	response := handleGeminiMessage("hi")

	slackAppToken := goDotEnvVariable("SLACK_APP_TOKEN")
	slackBotToken := goDotEnvVariable("SLACK_BOT_TOKEN")

	bot := slacker.NewClient(slackBotToken, slackAppToken)

	fmt.Println(fmt.Println(response))

	go printCommandEvents(bot.CommandEvents())
	bot.Command("ping", &slacker.CommandDefinition{
		Handler: func(bc slacker.BotContext, r slacker.Request, w slacker.ResponseWriter) {
			w.Reply(response)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatalf("Error listening: %v", err)
	}
}
