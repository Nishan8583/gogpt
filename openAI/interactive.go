package openai

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// InteractiveChat takes a username, loops through users input, its gonna cost you.
func (oa OpenAI) InteractiveChat(username string) {
	prompt := fmt.Sprintf("(%s) >>> ", username)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("%s", prompt)
		userInput, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("ERROR while reading user input", err)
		}

		userInput = strings.ReplaceAll(userInput, "\n", "")
		userInput = strings.ReplaceAll(userInput, "\r", "")

		if userInput == "exit" || userInput == "quit" {
			fmt.Println("(system) >>> Thanks for the love !!! ❤️. See you tomorrow bye. 👋 👍")
			break
		} else {
			msg := oa.SendSingleMessage(userInput)
			fmt.Printf("(gpt) >>> %s\n", msg)
		}

	}

}
