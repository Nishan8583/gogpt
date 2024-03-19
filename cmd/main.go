package main

import (
	"fmt"
	openai "gogpt/openAI"
	"os"

	"os/exec"

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
	LocalLLM          bool   `arg:"-l,--local-llm" help:"use local llm to answer the question"`
	VenvPath          string `arg:"--venv" help:"path to the virtual environment path" default:".venv"`
	PythonScript      string `arg:"--python-script" help:"path to python script" default:"code_scan_helper.py"`
}

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

	python_path := fmt.Sprintf("%s/Scripts/python.exe", c.VenvPath)
	if c.LocalLLM {
		log.Info().Msgf("Using local LLM for task")
		cmd := exec.Command(python_path, c.PythonScript, "-f", "bof.c")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Println("during python script execution", err)
			return
		}
		return
	}
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
		sentimens, err := oa.PerformSentimentAnalysis(c.ChatLogFile)
		if err != nil {
			log.Error().Msgf("error during sentiment analysis %v", err)
			return
		}
		fmt.Println(sentimens)
	} else {
		// a one off send message and reply from chatGPT
		fmt.Println(oa.SendSingleMessage(c.Message))
	}

}
