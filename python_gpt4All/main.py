from gpt4all import GPT4All
print("Loading the model")
model = GPT4All("C:\\D\\projects\\AI\\Models\\orca-mini-3b-gguf2-q4_0.gguf")
print()
c_code = """
#include <stdio.h>
#include <string.h>

#define S 100
#define N 1000

int main(int argc, char *argv[]) {
  char out[S];
  char buf[N];
  char msg[] = "Welcome to the argument echoing program\n";
  int len = 0;
  buf[0] = '\0';
  printf(msg);
  while (argc) {
    sprintf(out, "argument %d is %s\n", argc-1, argv[argc-1]);
    argc--;
    strncat(buf,out,sizeof(buf)-len-1);
    len = strlen(buf);
  }
  printf("%s",buf);
  return 0;
}
"""
print("model loaded")
with model.chat_session("You are a cyber secuirty expert who can find vulnerabilities in source code"):
    print("Chat session started")
    response = model.generate('Please find security vulnerability in provided source code encolsed in triple backticks. The source code here \n ```{}```'.format(c_code), temp=0.1)
    print(response)
