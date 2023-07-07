package openai

import (
	"net/http"
	"os"
	"time"

	"github.com/pkg/errors"
	yaml "gopkg.in/yaml.v3"
)

type OpenAI struct {
	APIKey string `yaml:"API_KEY"` // will hold the OPENAI Api key
	client *http.Client
}

// New will take path to a config file, parse yaml from that file, and return the config structure
func New(filepath string) (OpenAI, error) {
	oai := OpenAI{}
	content, err := os.ReadFile(filepath)
	if err != nil {
		return oai, errors.Wrap(err, "while reading the config file "+filepath)
	}

	if err := yaml.Unmarshal(content, &oai); err != nil {
		return oai, errors.Wrap(err, "while unmarshalling")
	}

	oai.client = &http.Client{Timeout: 10 * time.Second}

	return oai, err
}
