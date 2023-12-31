package openai

import (
	"errors"
	"fmt"

	"github.com/rs/zerolog/log"
)

var errEmtpyResonse = errors.New("empty response from OpenAI")

// ScanCode adds prompt indicating that its a cyber security researcher
func (oa OpenAI) ScanCode(code string) (string, error) {
	msg := fmt.Sprintf("Find a vulnerability in the provided code, where the code is included inside three backticks. Please give the output in json format with the following keys:'vulnerability' that holds the name of the vulnerability, 'severity' that classifies the severity of the vulnerability with high, medium or low, 'description' which gives detailed description about the vulnerability found, 'fix' which describes possible solution to the vulnerability and 'exploit' which gives a sample code to exploit the vulnerability. The code is given inside the following 3 backticks ```%s```", code)
	resp, err := oa.Chat([]Message{
		{Role: "system", Content: "you are a cyber security researcher who has experience doing source code audit, code review and in finding vulnerabilities in source code in multiple different programming languages."},
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
