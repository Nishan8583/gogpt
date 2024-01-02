from gpt4all import GPT4All
model = GPT4All("nous-hermes-llama2-13b.Q4_0.gguf")
with model.chat_session("You are a cyber secuirty expert who can find vulnerabilities in source code"):
    response = model.generate('who are you?', temp=0)
    print(response)