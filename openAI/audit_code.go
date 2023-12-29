package openai

import (
	"os"
	"path/filepath"

	"github.com/rs/zerolog/log"
)

// AuditCode takes in path to directory containing source code and path to output directories
// It reads the content of each source code, and sends openai to find any vulnerability it can find.
func (oa OpenAI) AuditCode(code_dir, output_dir string) {

	if _, err := os.Stat(output_dir); err != nil {
		if os.IsNotExist(err) {
			log.Info().Msgf("%s directory does not exist for report output, so creating new", output_dir)
			if err := os.Mkdir(output_dir, 0555); err != nil {
				log.Fatal().Msgf("creation of directory %s failed due to error %v", output_dir, err)
			}
		}
	}

	filepath.Walk(code_dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			content, err := os.ReadFile(path)
			if err != nil {
				log.Warn().Msgf("error while trying to read file=%s error=%v", path, err)
				return nil
			}
			log.Debug().Msgf("sending file content %s to openai", path)
			resp, err := oa.ScanCode(string(content))
			if err != nil {
				log.Warn().Msgf("while getting report for %s", path)
				return nil
			}

			_, fileName := filepath.Split(path)
			fileName = fileName + "_output.txt"
			if err := os.WriteFile(filepath.Join(output_dir, fileName), []byte(resp), 0555); err != nil {
				log.Warn().Msgf("while creating output file for %s error=%s", path, err)
			}
			log.Debug().Msgf("report generated and saved in file %s", fileName)

		}
		return nil
	})

}
