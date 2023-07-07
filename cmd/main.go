package main

import (
	"flag"
	"fmt"
	openai "gogpt/openAI"
	"log"
	"os"
)

func main() {

	msg := "" // holds user message
	flag.StringVar(&msg, "message", "", "--message <message to send chatGPT>")
	flag.Parse()

	if len(msg) == 0 {
		flag.Usage()
		os.Exit(-1)
	}

	gpt, err := openai.New("../settings.yaml")
	if err != nil {
		log.Fatal("could not create type", err)
	}
	resp, err := gpt.Chat([]openai.Message{
		{"user", msg},
	})

	if err != nil {
		log.Fatal("while chatting with gpt", err)
	}

	if len(resp.Choices) <= 0 {
		log.Fatal("WARN chat gpt empty reply")
	}

	fmt.Println(resp.Choices[0].Message.Content)

}
