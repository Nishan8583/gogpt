package openai

import (
	"fmt"

	"github.com/rs/zerolog/log"
)

func (oa OpenAI) ScanCode(code string) {
	msg := fmt.Sprintf(`Could you find a vulnerability in the following code %s`, code)
	resp, err := oa.Chat([]Message{
		{Role: "system", Content: "you are a cyber security researcher who has experience in finding vulnerabilities in code."},
		{Role: "user", Content: msg},
	})

	log.Debug().Msgf("sending code with msg %+v", msg)
	if err != nil {
		log.Fatal().Msgf("could not get response from OpenAI %+v", err)
	}

	fmt.Println(resp.Choices[0].Message.Content)
}
