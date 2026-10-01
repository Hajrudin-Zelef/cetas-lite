---
id: collect-261001-ia-llm/ia-llm/meta-is-back-with-muse-glimmer-local-agentic-multimodal-and-open-source-3
title: "Load model"
domain: ia-llm
role: reference
task: reference
actors: ["Hugging Face", "Meta", "OpenAI", "vLLM"]
dates: []
keywords: ["agent", "agents", "decode", "gguf", "gpu", "gpus", "inference", "llama", "lora", "mcp", "muse", "quantization"]
source: docs/RAG/collect-261001-ia-llm/meta-is-back-with-muse-glimmer-local-agentic-multimodal-and-open-source.md
source_anchor: ""
source_lines: [323, 489]
sha256: 2ec9ac99e06d40cc192e0832f4693557886ed583b8110e9fdc888f5d134f52bb
---

# Load model

```
import torch
from transformers import AutoProcessor, MuseGlimmerAssistantModel, MuseGlimmerForConditionalGeneration
model_id = "meta-models/Muse-Glimmer-30B"
assistant_model_id = "meta-models/Muse-Glimmer-30B-assistant"
target = MuseGlimmerForConditionalGeneration.from_pretrained(model_id, dtype=torch.bfloat16, device_map="auto")
assistant = MuseGlimmerAssistantModel.from_pretrained(assistant_model_id, dtype=torch.bfloat16, device_map="auto")
processor = AutoProcessor.from_pretrained(model_id)
messages = [
    {
        "role": "user", "content": [
            {"type": "image", "url": "https://huggingface.co/datasets/merve/vl-test-suite/resolve/main/SF.png"},
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
).to(target.device)
input_len = inputs["input_ids"].shape[-1]
outputs = target.generate(
    **inputs,
    assistant_model=assistant,
    speculation_type="dflash",
    do_sample=True
)
response = processor.decode(outputs[0][input_len:], skip_special_tokens=False)
print(response)
```
You can start llama server using following command. `--spec-draft-n-max` argument controls how many future tokens DFlash proposes during each speculative-decoding step. Muse Glimmer’s DFlash model was trained with a block size of 16, one anchor token plus 15 proposed tokens, so any value above 15 will be clamped to 15.

```
llama serve -hf meta-models/Muse-Glimmer-30B-GGUF --spec-type draft-dflash --spec-draft-n-max 15
```
You can also use llama cli with speculative decoding drafter as follows.

```
llama cli -hf meta-models/Muse-Glimmer-30B-GGUF --spec-type draft-dflash
```
For a managed, autoscaling deployment, open the Muse Glimmer 30B Inference Endpoints preset. The model is already selected: choose the organization, cloud provider, region, compatible GPU instance, authentication, and autoscaling settings, then review the hourly price and click **Create Endpoint**. Once its status is **Running**, you can test it in the Playground and copy the endpoint URL and model name from the Overview.

The deployed model exposes an OpenAI-compatible Chat Completions API. Keep your Hugging Face token in an environment variable, set `HF_ENDPOINT_URL` to the URL shown in the Overview without `/v1`, and set `HF_ENDPOINT_MODEL` to the endpoint's model name.

```
export HF_TOKEN="hf_..."
export HF_ENDPOINT_URL="https://<endpoint-id>.<region>.<cloud>.endpoints.huggingface.cloud"
export HF_ENDPOINT_MODEL="<endpoint-model-name>"
pip install --upgrade openai
```
```
import os
from openai import OpenAI
client = OpenAI(
    base_url=f"{os.environ['HF_ENDPOINT_URL'].rstrip('/')}/v1/",
    api_key=os.environ["HF_TOKEN"],
)
response = client.chat.completions.create(
    model=os.environ["HF_ENDPOINT_MODEL"],
    messages=[
        {"role": "system", "content": "You are a helpful assistant."},
        {"role": "user", "content": "Write a limerick about Python exceptions."},
    ],
    max_tokens=256,
)
print(response.choices[0].message.content)
```
See the Inference Endpoints documentation for configuration, autoscaling, security, logs, and monitoring.

For this release, we ship support for vLLM with transformers backend.

```
# tensor parallel serving across 4 GPUs
vllm serve meta-models/Muse-Glimmer-30B --model-impl transformers --tensor-parallel-size 4
# infer
curl -s http://127.0.0.1:8000/v1/chat/completions \
    -H 'Content-Type: application/json' \
    -d '{
      "model": "meta-models/Muse-Glimmer-30B",
      "messages": [
        {"role": "user", "content": "Explain tensor parallelism briefly."}
      ],
      "temperature": 0.0,
      "max_tokens": 256
    }'
```
You can use TRL to fine-tune Muse Glimmer using various methods from SFT to Async GRPO. We have run two experiments on bf16 with Hopper-class GPUs with 80GB VRAM each.

| Workload | Practical minimum | 
|---|---|
| Inference / eval, BF16 | 1×80 GB H100 | 
| LoRA SFT, BF16 | 1×80 GB H100, microbatch 1 + checkpointing | 
| Full SFT, BF16 | 8×80 GB H100 with FSDP/ZeRO-3 | 
| LoRA GRPO, Transformers rollouts | 1×80 GB H100, but slow/tight | 
| LoRA GRPO, separate vLLM rollout server | 8×H100: 4 rollout + 4 training | 
| Full-finetune GRPO | 8 GPUs is usually insufficient | 

As part of this release, we ship an example to fine-tune Muse Glimmer on small split of MolmoWeb dataset. This shows how to make model generate structured outputs and how to fine-tune on images.

We also experimented with running the model on OpenCode with AsyncGRPO example. Model shows strong coding capabilities, so we encourage you to try training with coding environments in OpenEnv and TRL.

Here are some fun ways to try out Muse Glimmer. In our opinion, the coolest thing about this model is that it is a local scale personal assistant that can code. That means you can make it do things like, quantize itself, find quantized weights on the Hub, deploy itself to inference endpoints, and even optimize itself for specific hardware! Let’s go team local 🚀

Assume the Inference Endpoint exposes an OpenAI-compatible `/v1` API.

Set `HF_TOKEN` in the OpenClaw gateway environment, then add this to `~/.openclaw/openclaw.json`:

## OpenClaw configuration

```
{
  models: {
    mode: "merge",
    providers: {
      muse: {
        baseUrl: "https://YOUR-ENDPOINT.endpoints.huggingface.cloud/v1",
        apiKey: {
          source: "env",
          provider: "default",
          id: "HF_TOKEN"
        },
        api: "openai-completions",
        authHeader: true,
        models: [{
          id: "meta-models/Muse-Glimmer-30B",
          name: "Muse Glimmer",
          reasoning: false,
          input: ["text", "image"],
          contextWindow: 32768,
          maxTokens: 8192
        }]
      }
    }
  },
  agents: {
    defaults: {
      model: { primary: "muse/meta-models/Muse-Glimmer-30B" }
    }
  }
}
```
Restart OpenClaw:

```
openclaw gateway restart
```
Validate from a fresh session:

```
openclaw agent --message "Reply with: muse-ready"
```
Use the exact model ID returned by the endpoint’s `/v1/models` response if it differs.

If we hook up Muse Glimmer to the Hugging Face MCP and update its `AGENTS.md` we give it the capability to find a quantized version of itself on the hub and run locally. This is handy if you want to work on something private, or just cut costs.

If you do this a second time, Muse Glimmer will find the cached weights and switch to them, so feel free to add a convenient command like `/spawn`. 

Muse Glimmer inspects the machine and Hub, selects or creates a Q4_K_M GGUF, launches llama-server, and validates model discovery and chat completion. The result is a smaller local build behind an OpenAI-compatible API. Here’s the prompt we added to `AGENTS.md`.

By adding this to `AGENTS.md` openclaw or hermes will be able to solve the rest.

## Local quantization prompt

