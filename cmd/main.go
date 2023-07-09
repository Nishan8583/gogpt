package main

import (
	"flag"
	"fmt"
	openai "gogpt/openAI"
	"log"
	"os"
)

// -i option to make it interactive
// (username) >>>
// (gpt) >>>
func main() {

	// setting up flags for interactive
	interactiveCmd := flag.NewFlagSet("interactive", flag.ExitOnError)
	username := interactiveCmd.String("username", "you", "--username <your username>")

	// setting up one shot message sending
	msg := ""
	configFile := ""
	flag.StringVar(&msg, "message", "", "--message <message to send chatGPT>")
	flag.StringVar(&configFile, "config", "../settings.yaml", "--message <path to config>")

	if len(os.Args) < 2 {
		log.Fatal("ERROR please provide flags for me to work with.")
	}

	switch os.Args[1] {
	case "interactive", "--interactive", "-i":
		interactiveCmd.Parse(os.Args[2:])
		gpt, err := openai.New(configFile)
		if err != nil {
			log.Fatal("could not craft gpt instance", err)
		}
		gpt.InteractiveChat(*username)
	default:
		flag.Parse()
		gpt, err := openai.New(configFile)
		if err != nil {
			log.Fatal("could not craft gpt instance", err)
		}
		fmt.Println(gpt.SendSingleMessage(msg))
	}

}
