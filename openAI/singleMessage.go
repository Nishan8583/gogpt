package openai

import (
	"flag"
	"log"
	"os"
)

// SendSingleMessage takes the message and the path to config file to read config from
func (oa OpenAI) SendSingleMessage(msg string) string {

	if len(msg) == 0 {
		flag.Usage()
		os.Exit(-1)
	}

	resp, err := oa.Chat([]Message{
		{"user", msg},
	})

	if err != nil {
		log.Fatal("while chatting with gpt", err)
	}

	if len(resp.Choices) <= 0 {
		log.Fatal("WARN chat gpt empty reply")
	}

	return resp.Choices[0].Message.Content
}
