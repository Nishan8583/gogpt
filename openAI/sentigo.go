package openai

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
)

type Result struct {
	Name      string `json:"name"`
	Sentiment string `json:"sentiment"`
	Summary   string `json:"summary"`
}
type Sentiments struct {
	Players []Result `json:"players"`
}

var prompt = "Perform sentiment analysis on the chat logs and respond in json. The json should contain array, each element in array should contain the following keys: 'name' which has the value the name of the sender, 'sentiment' which has the list of top human expression being displayed like anger, hate,  kindness, toxicity and so on,'summary' key that summarises the content of messages that each individual has sent. They 'summary' key should explains what messages each person has sent, not in terms of sentiment, rather just summarize what they were saying in chat.   Read the entire message the person has sent, and summarize the sentiment in the 'sentiment' key. Do no analyze the sentiment of each message, rather read the all messages of each person, and return a single sentiment for each person. The following chat logs contains messages, the name of the person who sent the message is in the beginning of the line, the senders name and the message is separated by a semicolon. The chat log is given below between triple backticks:```%s```"

const systemRole = `You are a game moderator who is responsbile for reading chat logs in a multiplayer online game, and detect toxic players.`

func (oa OpenAI) PerformSentimentAnalysis(chatLogFile string) (string, error) {
	verdict := ""
	chats, err := os.ReadFile(chatLogFile)
	if err != nil {
		return verdict, errors.Wrap(err, "while trying to read chat logs")
	}

	resp, err := oa.Chat([]Message{
		{Role: "system", Content: systemRole},
		{Role: "user", Content: fmt.Sprintf(prompt, string(chats))},
	})

	if err != nil {
		return verdict, errors.Wrap(err, "while sending chat logs for sentiment analysis")
	}
	if len(resp.Choices) == 0 {
		return verdict, errEmtpyResonse
	}
	return resp.Choices[0].Message.Content, nil

}
