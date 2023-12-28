package openai

import (
	"errors"
	"fmt"

	"github.com/rs/zerolog/log"
)

var errEmtpyResonse = errors.New("empty response from OpenAI")

func (oa OpenAI) ScanCode(code string) (string, error) {
	msg := fmt.Sprintf(`Could you find a vulnerability in the following code %s`, code)
	resp, err := oa.Chat([]Message{
		{Role: "system", Content: "you are a cyber security researcher who has experience in finding vulnerabilities in code."},
		{Role: "user", Content: msg},
	})

	log.Debug().Msgf("sending code with msg %+v", msg)
	if err != nil {
		return "", err
	}
	if len(resp.Choices) == 0 {
		return "", errEmtpyResonse
	}

	return resp.Choices[0].Message.Content, nil
}
