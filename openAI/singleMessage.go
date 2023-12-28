package openai

import (
	"flag"
	"os"

	"github.com/rs/zerolog/log"
)

// SendSingleMessage takes the message and the path to config file to read config from
func (oa OpenAI) SendSingleMessage(msg string) string {

	if len(msg) == 0 {
		flag.Usage()
		os.Exit(-1)
	}

	resp, err := oa.Chat([]Message{
		{Role: "system", Content: "you are a friend"},
		{Role: "user", Content: msg},
	})

	if err != nil {
		log.Fatal().Msgf("while chatting with gpt %+v", err)
	}
	log.Debug().Msgf("reply from server %+v", resp)
	if len(resp.Choices) <= 0 {
		log.Fatal().Msgf("WARN chat gpt empty reply %+v", err)
	}

	return resp.Choices[0].Message.Content
}
