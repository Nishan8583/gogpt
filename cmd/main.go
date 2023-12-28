package main

import (
	"embed"
	"fmt"
	openai "gogpt/openAI"
	"path/filepath"

	"os"

	args "github.com/alexflint/go-arg"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type config struct {
	Interactive   bool   `arg:"-i,--interactive" help:"make the program run interactively"`
	AuditCode     bool   `arg:"-a,--audit-code" help:"audits code"`
	ConfigFile    string `arg:"-c,--config" help:"path to config file that has the openAI API key" default:"./settings.json"`
	Debug         bool   `arg:"-d,--debug" help:"set it to enable debug mode" default:"false"`
	Username      string `arg:"-u,--username" help:"username to display" default:"you"`
	Message       string `arg:"-m,--message" help:"message to send to GPT"`
	CodeDirectory string `arg:"--code-dir" help:"give directory to all code files" default:"./sample_code"`
}

//go:embed templates/chat.html
var content embed.FS

// -i option to make it interactive
// (username) >>>
// (gpt) >>>
func main() {

	c := config{}
	args.MustParse(&c)
	log.Debug().Msgf("%+v", c)
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if c.Debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	//server.StartServer(content)

	if len(os.Args) < 2 {
		log.Fatal().Msg("ERROR please provide flags for me to work with.")
	}
	fmt.Println("Well this is interesting")

	if c.Interactive {
		gpt, err := openai.New(c.ConfigFile)
		if err != nil {
			log.Fatal().Msgf("could not craft gpt instance %+v", err)
		}
		gpt.InteractiveChat(c.Username)
	} else if c.AuditCode {
		auditor, err := openai.New(c.ConfigFile)
		if err != nil {
			log.Fatal().Msgf("error while creating code auditor %+v", err)
		}

		filepath.Walk(c.CodeDirectory, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				content, err := os.ReadFile(path)
				if err != nil {
					log.Warn().Msgf("error while trying to read file=%s error=%v", path, err)
				}
				auditor.ScanCode(string(content))
			}
			return nil
		})

	} else {

		gpt, err := openai.New(c.ConfigFile)
		if err != nil {
			log.Fatal().Msgf("could not craft gpt instance %+v", err)
		}
		fmt.Println(gpt.SendSingleMessage(c.Message))
	}

}
