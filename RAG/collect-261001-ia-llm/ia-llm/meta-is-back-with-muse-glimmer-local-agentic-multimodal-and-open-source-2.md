---
id: collect-261001-ia-llm/ia-llm/meta-is-back-with-muse-glimmer-local-agentic-multimodal-and-open-source-2
title: "Load model"
domain: ia-llm
role: reference
task: reference
actors: ["Meta", "Unsloth"]
dates: []
keywords: ["agents", "benchmark", "decode", "diffusion", "gguf", "inference", "llama", "llama.cpp", "multimodal", "muse", "parameters", "reasoning"]
source: docs/RAG/collect-261001-ia-llm/meta-is-back-with-muse-glimmer-local-agentic-multimodal-and-open-source.md
source_anchor: ""
source_lines: [119, 322]
sha256: a201ff744c317a3e36fce53d908673691f972f250029417fedabcd9c8bb5fbe1
---

# Load model

```
from transformers import AutoProcessor, AutoModelForMultimodalLM
MODEL_ID = "meta-models/Muse-Glimmer-30B"
# Load model
processor = AutoProcessor.from_pretrained(MODEL_ID)
model = AutoModelForMultimodalLM.from_pretrained(
    MODEL_ID,
    dtype="auto",
    device_map="auto"
)
# Images + Text
messages = [
    {
        "role": "user", "content": [
            {"type": "image", "image": "https://huggingface.co/datasets/merve/vl-test-suite/resolve/main/SF.png"},
            {"type": "text", "text": "What is shown in this image?"}
        ]
    }
]
inputs = processor.apply_chat_template(
    messages,
    tokenize=True,
    return_dict=True,
    return_tensors="pt",
    add_generation_prompt=True,
    reasoning_strength="low"
).to(model.device)
input_len = inputs["input_ids"].shape[-1]
# Generate output
outputs = model.generate(**inputs)
response = processor.decode(outputs[0][input_len:], skip_special_tokens=False)
print(response)
```
To work with videos we recommend installing `torchcodec` into the environment.

```
pip install torchcodec
```
Muse Glimmer can answer complex questions about videos *without audio*. You can do video inference as follows, here’s an example from VideoMME2, which is the most popular video question answering benchmark.

```
from transformers import AutoProcessor, AutoModelForMultimodalLM
MODEL_ID = "meta-models/Muse-Glimmer-30B"
# Load model
processor = AutoProcessor.from_pretrained(MODEL_ID)
model = AutoModelForMultimodalLM.from_pretrained(
    MODEL_ID,
    dtype="auto",
    device_map="auto"
)
# Videos + Text
messages = [
    {
        "role": "user",
        "content": [
            {"type": "video", "video": "https://huggingface.co/datasets/merve/vl-test-suite/resolve/main/IMG_8137.mp4"},
            {"type": "text", "text": "Describe what happens in this video."},
        ],
    },
]
inputs = processor.apply_chat_template(
    messages,
    tokenize=True,
    return_dict=True,
    return_tensors="pt",
    add_generation_prompt=True,
    reasoning_strength="low",
    processor_kwargs={"num_frames": 96},
).to(model.device)
input_len = inputs["input_ids"].shape[-1]
outputs = model.generate(**inputs)
response = processor.decode(
    outputs[0, input_len:],
    skip_special_tokens=False,
)
print(response)
```
Muse Glimmer can do multimodal tool calling, here’s how you can do it. In the example below, we ask the model to call the weather tool based on the city in the image.

```
import json
import re
from transformers import AutoProcessor, AutoModelForMultimodalLM
MODEL_ID = "meta-models/Muse-Glimmer-30B"
# Load model
processor = AutoProcessor.from_pretrained(MODEL_ID)
model = AutoModelForMultimodalLM.from_pretrained(
    MODEL_ID,
    dtype="auto",
    device_map="auto"
)
tools = [
    {
        "type": "function",
        "function": {
            "name": "weather.get",
            "description": "Get the current weather for a city.",
            "parameters": {
                "type": "object",
                "properties": {
                    "city": {"type": "string"},
                },
                "required": ["city"],
            },
        },
    }
]
messages = [
    {
        "role": "user",
        "content": [
            {"type": "image", "image": "https://huggingface.co/datasets/merve/vl-test-suite/resolve/main/SF.png"},
            {"type": "text", "text": "I'm going to the city in this picture. What clothes should I wear?"},
        ],
    },
]
inputs = processor.apply_chat_template(
    messages,
    tools=tools,
    tokenize=True,
    return_dict=True,
    return_tensors="pt",
    add_generation_prompt=True,
    reasoning_strength="low"
).to(model.device)
input_len = inputs["input_ids"].shape[-1]
outputs = model.generate(**inputs)
response = processor.decode(outputs[0][input_len:], skip_special_tokens=False)
print(response)
```
You can use Muse Glimmer to do open ended object detection in images as follows.

```
from transformers import AutoProcessor, AutoModelForMultimodalLM
MODEL_ID = "meta-models/Muse-Glimmer-30B"
# Load model
processor = AutoProcessor.from_pretrained(MODEL_ID)
model = AutoModelForMultimodalLM.from_pretrained(
    MODEL_ID,
    dtype="auto",
    device_map="auto"
)
messages = [{
    "role": "user",
    "content": [
        {"type": "image", "image": "https://huggingface.co/datasets/merve/vl-test-suite/resolve/main/SF.png"},
        {
            "type": "text",
            "text": (
                "Detect the bridge. Return only the detection in the model's "
                "native object-detection format, with no explanation."
            ),
        },
    ],
}]
inputs = processor.apply_chat_template(
    messages,
    tokenize=True,
    return_dict=True,
    return_tensors="pt",
    add_generation_prompt=True,
    reasoning_strength="low",
).to(model.device)
input_len = inputs["input_ids"].shape[-1]
outputs = model.generate(**inputs, max_new_tokens=128)
response = processor.decode(outputs[0][input_len:], skip_special_tokens=False)
detections = json.loads(response.removesuffix("<|eot|>"))
print(detections)
```
Here is an end to end script to perform object detection GitHub Gist


Muse Glimmer comes with day-0 llama.cpp support. Meta has distributed calibrated quants in this repo, and Unsloth is releasing optimized quants as well. DFlash speculative decoding is supported as well. You can use a pre-built llama binary to start a llama server or a CLI. To install llama.cpp, run:

```
curl -LsSf https://llama.app/install.sh | sh
```
Then you can start the server as follows.

```
llama serve -hf meta-models/Muse-Glimmer-30B-GGUF
```
Once the server has started, you can head to localhost:8080 to chat with the built-in WebUI.

You can also query the server as follows.

```
curl http://localhost:8080/v1/chat/completions \
    -H "Content-Type: application/json" \
    -d '{
        "messages": [
            {"role": "system", "content": "You are a helpful assistant."},
            {"role": "user", "content": "Write a limerick about python exceptions"}
        ]
    }'
```
You can also use llama server with coding agents like Pi.

DFlash uses a lightweight block-diffusion drafter model to provide same output with extra speed-ups in decoding phase. Transformers and llama.cpp ship support for DFlash drafter of Muse Glimmer day-0.

Below you can see how speculative decoding can speed-up generation in realistic setups. The video shows llama.cpp webui with DFlash on the left and regular generation on the right.

You can load the drafter and model as follows, and infer like how you would with base model with an additional parameter (shown in the upcoming snippets).

