# GOGPT
 - Its a simple CLI program to send your code to OpenAI to find for any vulnerabilities in your code.
 - Also has some basic cli chat functionality.

### Build 
 - In `cmd` directory `go build -o main.exe main.go`.

### Usage
 - Add your OpenAI API key in settings.yaml.
 - `main.exe --message="your message here"` 
 - `main.exe -i --username="imauser"`
 - `main.exe --audit-code -c settings.yaml`

### TODO
- [x] A simple chat interactive interface to OpenAI.
- [x] Send source code to openAI to find vulnerbility in soruce code and output results in directory.
- [ ] Create a front end browser UI to upload zip files containing source codes.
