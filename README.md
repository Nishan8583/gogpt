# GOGPT
 - Its a simple CLI program that uses openAI to do some interesting stuffs.
 - Has some basic CLI chat functionality, code scanning for vulnerabilities and sentiment analysis.

### Build 
 - `go build -o main.exe cmd/main.go` or use make.
 - Make sure to set openAI api key in settings.yaml.

### Usage
 - After your OpenAI API key is set in settings.yaml.
 - `main.exe --audit-code -c settings.yaml --code-dir sample_vulnerable_code`.

### TODO
- [x] A simple chat interactive interface to OpenAI.
- [x] Send source code to openAI to find vulnerbility in soruce code and output results in directory.
- [ ] Create a front end browser UI to upload zip files containing source codes.
- [x] Sentiment analysis.