package openai

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/pkg/errors"
)

type Result struct {
	Name      string `json:"name"`
	Sentiment string `json:"sentiment"`
}

var prompt = "Perform sentiment analysis on the chat logs and respond in json. The json should contain array, each element in array should contain the following keys: 'name' which has the value the name of the sender, 'sentiment' human expression being displayed like anger, hate,  kindness, toxicity and so on. Replace hateful emotions towards others such as hate and anger as 'toxic'. Read the entire message the person has sent, and summarize the sentiment in the 'sentiment' key. The following chat logs contains messages, the name of the person who sent the message is in the beginning of the line, the senders name and the message is separated by a semicolon. The chat log is given below between triple backticks:```%s```"

const systemRole = `You are a game moderator who is responsbile for reading chat logs in a multiplayer online game, and detect toxic players.`

func (oa OpenAI) PerformSentimentAnalysis(chatLogFile string) (Result, error) {
	verdict := Result{}
	chats, err := os.ReadFile(chatLogFile)
	if err != nil {
		return verdict, errors.Wrap(err, "while trying to read chat logs")
	}

	resp, err := oa.Chat([]Message{
		{Role: "system", Content: systemRole},
		{Role: "user", Content: fmt.Sprintf(prompt, string(chats))},
	})

	if len(resp.Choices) == 0 {
		return verdict, errEmtpyResonse
	}

	err = json.Unmarshal([]byte(resp.Choices[0].Message.Content), &verdict)
	return verdict, err
}
