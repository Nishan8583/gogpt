package main

import (
	"embed"
	"fmt"
	openai "gogpt/openAI"

	args "github.com/alexflint/go-arg"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type config struct {
	Interactive       bool   `arg:"-i,--interactive" help:"make the program run interactively"`
	AuditCode         bool   `arg:"-a,--audit-code" help:"audits code"`
	ConfigFile        string `arg:"-c,--config" help:"path to config file that has the openAI API key" default:"./settings.yaml"`
	SentimentAnalysis bool   `arg:"-s,--sentiment-analysis" help:"perform sentiment analysis"`
	ChatLogFile       string `arg:"-f,--chat-log" help:"chat log file path" default:"./sample_chat_logs/chat_log.txt"`
	Debug             bool   `arg:"-d,--debug" help:"set it to enable debug mode" default:"false"`
	Username          string `arg:"-u,--username" help:"username to display" default:"you"`
	Message           string `arg:"-m,--message" help:"message to send to GPT"`
	CodeDirectory     string `arg:"--code-dir" help:"give directory to all code files" default:"./sample_vulnerable_code"`
	OutputDirectory   string `arg:"-o,--ouput-dir" help:"output directory for the report" default:"./outputdir"`
}

//go:embed templates/chat.html
var content embed.FS

// -i option to make it interactive
// (username) >>>
// (gpt) >>>
func main() {

	// parse command line flags and set debug log level if requested
	c := config{}
	args.MustParse(&c)
	log.Debug().Msgf("%+v", c)
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if c.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	//server.StartServer(content)

	// Create a new openAI instance
	oa, err := openai.New(c.ConfigFile)
	if err != nil {
		log.Error().Msgf("while creating openai instance %v", err)
		return
	}

	if c.Interactive {
		// start a interactive chat session
		oa.InteractiveChat(c.Username)
	} else if c.AuditCode {
		// start code auditing
		log.Debug().Msg("starting to audit code")
		oa.AuditCode(c.CodeDirectory, c.OutputDirectory)
	} else if c.SentimentAnalysis {
		log.Debug().Msg("sentiment analysis")
		fmt.Println(oa.PerformSentimentAnalysis(c.ChatLogFile))
	} else {
		// a one off send message and reply from chatGPT
		fmt.Println(oa.SendSingleMessage(c.Message))
	}

}
