package openai

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewConfig(t *testing.T) {
	if err := os.WriteFile("./test_config.yaml", []byte(`API_KEY: sampleAPIKey`), 0555); err != nil {
		t.Fatal("during test config file creation", err)
	}
	defer os.Remove("./test_config.yaml")
	is := require.New(t)
	c, err := New("./test_config.yaml")
	is.Nil(err, "error when calling NewConfig %v", err)

	is.Equal("sampleAPIKey", c.APIKey, "Expected API KEY sampleAPIKey found %s", c.APIKey)
}
