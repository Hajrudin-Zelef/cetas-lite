---
id: collect-250926-servers-hardware/servers-hardware/nvidia-nemotron-3-nano-omni-how-to-run-locally-2
title: "NVIDIA Nemotron 3 Nano Omni - How To Run Locally"
domain: servers-hardware
role: reference
task: reference
actors: ["Alibaba", "Hugging Face", "Nvidia", "OpenAI", "Unsloth"]
dates: []
keywords: ["nvidia", "omni", "agent", "benchmark", "benchmarks", "decode", "fine-tuning", "gguf", "llama", "llama.cpp", "memory", "multimodal"]
source: docs/RAG/clean4/NVIDIA Nemotron 3 Nano Omni - How To Run Locally.md
source_anchor: ""
source_lines: [41, 180]
sha256: 0888953cacf996e7e9884134c3a0b72b95e644b7deda111d1c37af4d9803e72d
---

# NVIDIA Nemotron 3 Nano Omni - How To Run Locally

, which is just our mini logo showing how finetunes are made with Unsloth:
{% code overflow="wrap" %}
```bash
wget https://raw.githubusercontent.com/unslothai/unsloth/refs/heads/main/images/unsloth%20made%20with%20love.png -O unsloth.png
```
{% endcode %}
Let's get the 2nd image at
{% code overflow="wrap" %}
```bash
wget https://files.worldwildlife.org/wwfcmsprod/images/Sloth_Sitting_iStock_3_12_2014/story_full_width/8l7pbjmj29_iStock_000011145477Large_mini__1_.jpg -O picture.png
```
{% endcode %}
{% endstep %}
{% step %}
Now let's download the model manually. We can do this via the code below (after installing pip install huggingface\_hub). If downloads get stuck, see: [Hugging Face Hub, XET debugging](/docs/basics/troubleshooting-and-faqs/hugging-face-hub-xet-debugging.md)
{% code overflow="wrap" %}
```bash
pip install huggingface_hub
hf download unsloth/NVIDIA-Nemotron-3-Nano-Omni-30B-A3B-Reasoning-GGUF \
--local-dir unsloth/NVIDIA-Nemotron-3-Nano-Omni-30B-A3B-Reasoning-GGUF \
--include "*mmproj-BF16*" \
--include "*UD-Q4_K_XL*" # Use "*UD-Q2_K_XL*" for Dynamic 2bit
```
{% endcode %}
{% endstep %}
{% step %}
Then run the model in conversation mode:
{% code overflow="wrap" %}
```bash
./llama.cpp/llama-cli \
--model unsloth/NVIDIA-Nemotron-3-Nano-Omni-30B-A3B-Reasoning-GGUF/NVIDIA-Nemotron-3-Nano-Omni-30B-A3B-Reasoning-UD-Q4_K_XL.gguf \
--mmproj unsloth/NVIDIA-Nemotron-3-Nano-Omni-30B-A3B-Reasoning-GGUF/mmproj-BF16.gguf \
--temp 0.6 \
--top-p 0.95 \
--min-p 0.01
```
{% endcode %}
{% endstep %}
{% step %}
You will then see the below:
{% endstep %}
{% step %}
Then use `/image` to load both images in and ask "What is this image":
{% endstep %}
{% step %}
And for the sloth image:
{% endstep %}
{% endstepper %}
#### Llama-server serving & deployment
To deploy Nemotron 3 Nano Omni locally, use `llama-server`. In a new terminal, for example via `tmux`, deploy the model:
```bash
./llama.cpp/llama-server \
-hf unsloth/NVIDIA-Nemotron-3-Nano-Omni-30B-A3B-Reasoning-GGUF:UD-Q4_K_XL \
--alias "unsloth/NVIDIA-Nemotron-3-Nano-Omni-30B-A3B-Reasoning" \
--prio 3 \
--temp 0.6 \
--top-p 0.95 \
--port 8001
```
If you downloaded the model manually, use:
{% code overflow="wrap" %}
```bash
./llama.cpp/llama-server \
--model unsloth/NVIDIA-Nemotron-3-Nano-Omni-30B-A3B-Reasoning-GGUF/NVIDIA-Nemotron-3-Nano-Omni-30B-A3B-Reasoning-UD-Q4_K_XL.gguf \
--mmproj unsloth/NVIDIA-Nemotron-3-Nano-Omni-30B-A3B-Reasoning-GGUF/mmproj-BF16.gguf \
--alias "unsloth/NVIDIA-Nemotron-3-Nano-Omni-30B-A3B-Reasoning" \
--prio 3 \
--temp 0.6 \
--top-p 0.95 \
--port 8001
```
{% endcode %}
Then in a new terminal, after installing the OpenAI client with `pip install openai`:
```python
from openai import OpenAI
openai_client = OpenAI(
base_url = "http://127.0.0.1:8001/v1",
api_key = "sk-no-key-required",
)
completion = openai_client.chat.completions.create(
model = "unsloth/NVIDIA-Nemotron-3-Nano-Omni-30B-A3B-Reasoning",
messages = [
{"role": "user", "content": "What is 2+2?"},
],
)
print(completion.choices[0].message.reasoning_content)
print(completion.choices[0].message.content)
```
Which will show something like the below:
#### Image input through the OpenAI-compatible server
Let's use `picture.png` which was the sloth image like in [#llama.cpp-tutorial](#llama.cpp-tutorial "mention")
{% code expandable="true" %}
```python
from openai import OpenAI
import base64
import mimetypes
image_link = "picture.png"
def file_to_data_url(path: str) -> str:
mime = mimetypes.guess_type(path)[0] or "application/octet-stream"
with open(path, "rb") as f:
data = base64.b64encode(f.read()).decode("utf-8")
return f"data:{mime};base64,{data}"
openai_client = OpenAI(
base_url = "http://127.0.0.1:8001/v1",
api_key = "sk-no-key-required",
)
completion = openai_client.chat.completions.create(
model = "unsloth/NVIDIA-Nemotron-3-Nano-Omni-30B-A3B-Reasoning",
messages = [ {
"role": "user",
"content": [
{ "type": "text", "text": "What is this image?", },
{
"type": "image_url",
"image_url": { "url": file_to_data_url(image_link), },
},
],
} ],
)
print(completion.choices[0].message.reasoning_content)
print(completion.choices[0].message.content)
```
{% endcode %}
Which will show something like below:
### 🦥 Fine-tuning Nemotron 3 Nano Omni
Unsloth supports the entire [Nemotron](/docs/models/nemotron-3.md) model family. Nemotron 3 Nano Omni is useful for multimodal agent datasets. You can train on audio, vision or text via Unsloth. **Video input** fine-tuning is currently not supported.
For text-only and notebooks, you can start from the existing [Nemotron 3 Nano fine-tuning flow](/docs/models/nemotron-3.md#fine-tuning-nemotron-3-and-rl). For multimodal adapters, make sure your dataset includes the modality your agent actually needs:
* **Computer use:** screenshots, UI state, cursor/context, expected next action
* **Document intelligence:** PDFs, screenshots, charts, tables, structured extraction targets
* **Audio understanding:** audio clips, sampled frames, summaries, timestamps, events and follow-up questions
* **Agent loops:** observation → reasoning → action → validation examples
For Omni, do not blindly reuse text-only VRAM numbers. Multimodal encoders, projector weights, image tokens, audio chunks and long context all increase memory use. Start with shorter contexts and smaller batch sizes, then scale up.
### Benchmarks
Nemotron 3 Nano Omni is the strongest omni model for its size. It is also the highest-efficiency open multimodal model with leading accuracy. The model surpasses Qwen3-Omni-30B-A3B on every benchmark.
&goal=
```
`ask` is the immediate question: it should be specific, self-contained, and written in natural language.
`goal` is optional and describes the broader end goal you are ultimately trying to accomplish on behalf of the user. GitBook uses it to tailor the answer towards what is most useful for that goal.
The response will contain a direct answer to the question and relevant excerpts and sources from the documentation.
Use this mechanism when the answer is not explicitly present in the current page, you need clarification or additional context, or you want to retrieve related documentation sections.
