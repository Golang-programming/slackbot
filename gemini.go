package main

import (
	"context"
	"fmt"
	"log"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func handleGeminiMessage(_ string) string {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(goDotEnvVariable("GEMINI_API_KEY")))
	model := client.GenerativeModel("gemini-1.0-pro")

	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	resp, err := model.GenerateContent(
		ctx,
		genai.Text("Write a paragraph on cow in 5 parts and mention part with heading"),
	)

	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	fmt.Println(resp)

	return resp.Candidates[0].Content.Parts[0]
}
