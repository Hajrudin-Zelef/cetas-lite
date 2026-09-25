---
id: collect-250926-servers-hardware/servers-hardware/tool-calling-guide-for-local-llms-unsloth-documentation
title: "tool-calling-guide-for-local-llms-unsloth-documentation"
domain: servers-hardware
role: reference
task: reference
actors: ["Apple", "OpenAI", "Unsloth", "Z.ai"]
dates: []
keywords: ["gguf", "glm", "gpu", "inference", "llama", "llama.cpp", "parameters", "tool calling"]
source: docs/RAG/clean4/tool-calling-guide-for-local-llms-unsloth-documentation.md
source_anchor: ""
source_lines: [1, 58]
sha256: 087c27fce6be30f99b1bde3bb0f653daac8bbcb5e660d88bac1812b08d1aa6cd
---

# tool-calling-guide-for-local-llms-unsloth-documentation

Tool calling is when a LLM is allowed to trigger specific functions (like “search my files,” “run a calculator,” or “call an API”) by emitting a structured request instead of guessing the answer in text. You use tool calls because they make outputs **more reliable and up-to-date**, and they let the model **take real actions** (query systems, validate facts, enforce schemas) rather than hallucinating.

In this tutorial, you will learn how to use local LLMs via Tool Calling with Mathematical, story, Python code and terminal function examples. Inference is done locally via llama.cpp, llama-server and OpenAI endpoints.

Tool calling is automatically setup when you use Unsloth Studio. Just select your model and toggle on or off tool-calling.

See right for an example of tool-calling being automatically applied for Gemma 4. Unsloth also has self-healing tool-calling ensuring you always have working tool calls.

Our guide should work for nearly any model including:

Our first step is to Obtain the latest `llama.cpp` on GitHub here. You can follow the build instructions below as well. Change `-DGGML_CUDA=ON` to `-DGGML_CUDA=OFF` if you don't have a GPU or just want CPU inference. **For Apple Mac / Metal devices**, set `-DGGML_CUDA=OFF` then continue as usual - Metal support is on by default.

```
apt-get update
apt-get install pciutils build-essential cmake curl libcurl4-openssl-dev -y
git clone https://github.com/ggml-org/llama.cpp
cmake llama.cpp -B llama.cpp/build \
    -DBUILD_SHARED_LIBS=OFF -DGGML_CUDA=ON
cmake --build llama.cpp/build --config Release -j --clean-first --target llama-cli llama-mtmd-cli llama-server llama-gguf-split
cp llama.cpp/build/bin/llama-* llama.cpp
```
In a new terminal (if using tmux, use CTRL+B+D), we create some tools like adding 2 numbers, executing Python code, executing Linux functions and much more:

We then use the below functions (copy and paste and execute) which will parse the function calls automatically and call the OpenAI endpoint for any model:

Now we'll showcase multiple methods of running tool-calling for many different use-cases below:

In a new terminal, we create some tools like adding 2 numbers, executing Python code, executing Linux functions and much more:

We then use the below functions which will parse the function calls automatically and call OpenAI endpoint for any LLM:

Now we'll showcase multiple methods of running tool-calling for many different use-cases below:

We confirm the file was created and it was!

We first download GLM-4.7 or GLM-4.7-Flash via some Python code, then launch it via llama-server in a separate terminal (like using tmux). In this example we download the large GLM-4.7 model:

If you ran it successfully, you should see:

Now launch it via llama-server in a new terminal. Use tmux if you want:

And you will get:

Now in a new terminal and executing Python code, reminder to run Tool Calling Setup We use GLM 4.7's optimal parameters of temperature = 0.7 and top_p = 1.0

We first download Devstral 2 via some Python code, then launch it via llama-server in a separate terminal (like using tmux):

If you ran it successfully, you should see:

Now launch it via llama-server in a new terminal. Use tmux if you want:

You will see the below if it succeeded:

We then call the model with the following message and with Devstral's suggested parameters of temperature = 0.15 only. Reminder to run Tool Calling Setup

Last updated

Was this helpful?
