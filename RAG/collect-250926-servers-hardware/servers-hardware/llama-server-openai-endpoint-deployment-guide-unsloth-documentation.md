---
id: collect-250926-servers-hardware/servers-hardware/llama-server-openai-endpoint-deployment-guide-unsloth-documentation
title: "!pip install huggingface_hub hf_transfer"
domain: servers-hardware
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["gguf", "gpu", "inference", "llama", "llama.cpp", "speculative decoding", "tool calling"]
source: docs/RAG/clean4/llama-server-openai-endpoint-deployment-guide-unsloth-documentation.md
source_anchor: ""
source_lines: [1, 45]
sha256: 819c6953a306759c4f167162bad787f7834c084e8c70145a719ec51cb7d3758e
---

# !pip install huggingface_hub hf_transfer

We are going to deploy Devstral-2 - see Devstral 2 for more details on the model.

Obtain the latest `llama.cpp` on GitHub here. You can follow the build instructions below as well. Change `-DGGML_CUDA=ON` to `-DGGML_CUDA=OFF` if you don't have a GPU or just want CPU inference. **For Apple Mac / Metal devices**, set `-DGGML_CUDA=OFF` then continue as usual - Metal support is on by default.

```
apt-get update
apt-get install pciutils build-essential cmake curl libcurl4-openssl-dev -y
git clone https://github.com/ggml-org/llama.cpp
cmake llama.cpp -B llama.cpp/build \
    -DBUILD_SHARED_LIBS=OFF -DGGML_CUDA=ON -DLLAMA_CURL=ON
cmake --build llama.cpp/build --config Release -j --clean-first --target llama-cli llama-mtmd-cli llama-server llama-gguf-split
cp llama.cpp/build/bin/llama-* llama.cpp
```
First download Devstral 2:

```
# !pip install huggingface_hub hf_transfer
import os
os.environ["HF_HUB_ENABLE_HF_TRANSFER"] = "1"
from huggingface_hub import snapshot_download
snapshot_download(
    repo_id = "unsloth/Devstral-2-123B-Instruct-2512-GGUF",
    local_dir = "Devstral-2-123B-Instruct-2512-GGUF",
    allow_patterns = ["*UD-Q2_K_XL*", "*mmproj-F16*"],
)
```
To deploy Devstral 2 for production, we use `llama-server` In a new terminal say via tmux, deploy the model via:

When you run the above, you will get:

Then in a new terminal, after doing `pip install openai`, do:

Which will simply print 4. You can go back to the llama-server screen and you might see some statistics which might be interesting:

For arguments like using speculative decoding, see https://github.com/ggml-org/llama.cpp/blob/master/tools/server/README.md

- When using `--jinja` llama-server appends the following system message if tools are supported:`Respond in JSON format, either with tool_call (a request to call tools) or with response reply to the user's request` . This sometimes causes issues with fine-tunes! See the llama.cpp repo for more details.
You can stop this by using`--no-jinja` but then`tools` becomes unsupported.
For example FunctionGemma by default uses:But because of llama-server appending an extra message, we get: We reported the issue to https://github.com/ggml-org/llama.cpp/issues/18323 and llama.cpp developers are working on a fix! In the meantime, for all fine-tunes, please add the prompt specifically for tool calling!

See Tool Calling Guide on how to do tool calling!

Last updated

Was this helpful?
