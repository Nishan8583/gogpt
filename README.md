# GOGPT
 - Its a simple CLI program to send your code to OpenAI to find for any vulnerabilities in your code.
 - Also has some basic cli chat functionality.

### Build 
 - In `cmd` directory `go build main.go -o main.exe`

### Usage
 - Add your OpenAI API key in settings.yaml.
 - `main.exe --message="your message here"` 
 - `main.exe -i --username="imauser"`
 - `main.exe --audit-code -c settings.yaml`

### TODO
 - [] Create a front end browser UI to upload zip files containing source codes.