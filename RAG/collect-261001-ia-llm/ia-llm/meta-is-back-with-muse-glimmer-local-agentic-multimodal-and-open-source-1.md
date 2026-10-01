---
id: collect-261001-ia-llm/ia-llm/meta-is-back-with-muse-glimmer-local-agentic-multimodal-and-open-source-1
title: "Load model"
domain: ia-llm
role: reference
task: reference
actors: ["AMD", "Alibaba", "Hugging Face", "Intel", "Meta", "Nvidia", "vLLM"]
dates: []
keywords: ["accelerator", "agentic", "amd", "apache", "attention", "benchmark", "cost", "decode", "embedding", "embeddings", "gpus", "inference"]
source: docs/RAG/collect-261001-ia-llm/meta-is-back-with-muse-glimmer-local-agentic-multimodal-and-open-source.md
source_anchor: ""
source_lines: [1, 118]
sha256: 1306ecacd5c2a3061b6624a1145559b8f0b393d27f24fb3a1166909df6fae4d0
---

# Load model

**30B**parameters, and released under the

**Apache 2.0 license**, it’s ideal deploying locally for privacy, reducing costs, or just hacking around. It’s intended for privacy-aware applications such as coding, document analysis, personal assistants, Claw- or Hermes-like setups.

To celebrate, we are shipping with Meta day-0 support in `transformers`, `llama.cpp`, `vLLM`, Inference Endpoints, and other libraries. We built a few cool things and explain our findings in this blog.

You can find Muse Glimmer on the Hugging Face Hub.

## Benchmark results

Scores are reported as published. **Bold** indicates the best result among the compared models; ↓ indicates lower is better.

| Category | Benchmark | Muse Glimmer-30B High Reasoning | Gemma4-31B Thinking Mode | Qwen3.6-27B Thinking Mode | 
|---|---|---|---|---|
| General Agentic | MCP Atlas | **75.5** | 54.2 | 62.5 | 
| General Agentic | DeepSearch QA | **74.6** | 61.7 | 71.1 | 
| General Agentic | τ³-Banking | **23.5** | 15.1 | 16.7 | 
| General Agentic | WildClawBench | **47.6** | 37.6 | 43.2 | 
| General Agentic | GDPval-AA | 953 | 811 | **1141** | 
| General Agentic | GAIA2 | **43.3** | 36.4 | 40.0 | 
| General Agentic | SkillsBench (With Skills) | 44.3 | 32.4 | **46.6** | 
| General Agentic | OSWorld-Verified | 65.9 | 58.5 | **75.6** | 
| Agentic Coding | SWE-Bench Pro | **51.2** | 36.9 | 50.2 | 
| Agentic Coding | SWE-Bench Verified | 76.0 | 66.6 | **77.2** | 
| Agentic Coding | TerminalBench 2.1 | 51.7 | 43.4 | **60.7** | 
| Agentic Coding | SciCode | **43.6** | 43.4 | 39.8 | 
| Multimodal | Charxiv Reasoning | **78.8** | 77.7 | 78.4 | 
| Multimodal | ScreenSpot Pro | 75.4 | 75.9 | **76.1** | 
| Multimodal | OmniDocBench v1.5 | 75.8 | 72.5 | **77.8** | 
| Multimodal | MMMU Pro | 74 | 73 | **75** | 
| Safety | CI Memories | Violation (↓): 26.4 Coverage: 64.8 | **Violation (↓): 12.1**Coverage: 53.0 | Violation (↓): 53.4 Coverage: 66.9 | 
| Safety | Siren AgentDojo | Attack Success Rate (↓): 28.4 Utility: 94.2 | **Attack Success Rate (↓): 25.6**Utility: 90.8 | Attack Success Rate (↓): 40.3 Utility: 92.7 | 
| General Capabilities and Reasoning | IFBench | **77.0** | 76.0 | 70.8 | 
| General Capabilities and Reasoning | AIME 2026 | **94.7** | 89.2 | 94.1 | 
| General Capabilities and Reasoning | GPQA Diamond | 83.5 | **85.7** | 84.2 | 
| General Capabilities and Reasoning | Humanity’s Last Exam (Text + No Tools) | 22.0 | **23.6** | 23.1 | 
| General Capabilities and Reasoning | AA-LCR | **80.0** | 68.3 | 73.3 | 
| General Capabilities and Reasoning | Beam 128K | **65.1** | 58.2 | 63.0 | 

Muse Glimmer is a dense 30B parameter model consisting of:

- 2B ViT-style encoder for vision (Perception Encoder)
- 28B parameter text decoder

In addition to the main VLM, there’s also a speculative decoding drafter implemented on DFlash. Usage of this module is optional, and it can provide much faster generation in exchange for some memory cost. We found this drafter to be particularly well suited to structured content generation such as coding.

The language model uses the following architecture components:

- **Hybrid attention:** Alternating between three sliding window layers (of 2,048 tokens) using rotary position embedding, followed by a fourth layer that uses full attention and NoPE (no positional embedding). The pattern is therefore (SWA, SWA, SWA, Full), repeated 13 times to a total of 52 layers. This allows the model to retain relative order and distance information with RoPE and preserve information globally with NoPE.
- **Gated Grouped-Query Attention:** Each key-value head is shared by 16 query heads, which reduces KV-cache memory by 16x and makes generation faster and cheaper.
- **Q-K normalization with extra query scaling:** Before computing attention, Muse Glimmer applies RMS normalization to every query and key head to keep attention logits stable. After this, queries are multiplied by a scale factor to set the target logit scale after normalization. The extra query scaling behaves like an inverse temperature at the softmax level.

Muse Glimmer uses one image encoder to handle both images and videos. Unlike the relatively small vision encoders used in other VLMs, this is a sizable 2B ViT-like model designed after the Perception Encoder architecture. Perception Encoder was previously introduced by Meta as a backbone for various downstream spatial and multimodal tasks.

The encoder patchifies images to a shape of 2 frames x 3 channels x 14 x 14, and passes them through a linear layer for projection. An interpolated absolute position embedding from a learned position table is then added to these embeddings. These are then sent to the vision tower which consist of 50 layers and GELU MLPs. Similar to the language model, the attention pattern consists of three window attention layers followed by one full attention layer. Inside the attention layers, 2D RoPE is applied to the queries and keys.

After transformer, pixel shuffle concatenates 2x2 groups of neighboring spatial tokens which reduces the number of image tokens 4x without discarding their channels. The merged features are then projected to the shared embedding space of the text decoder.

Videos go through the same encoder frame by frame, where each frame is converted into patches (of shape [batch, temporal groups, grid height, grid width, 2 frames, 3 channels, 14, 14]). The processor targets 2 frames per second and caps the clip at 96 frames sampled evenly across video. The processor creates timestamped video placeholders, interleaving text with frame e.g. “Time: 0.0s <|video|> x N” in which the final video embeddings are replaced before the final projection layer.

Upgrade transformers to the latest version to be able to use Muse Glimmer.

```
pip install --upgrade transformers accelerate
```
Muse Glimmer comes with day-0 support in transformers, both for the main model and the speculative decoding drafter. You can use `AutoModelForMultimodalLM` and `AutoProcessor` classes to load the model and the processor.

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
```
The same snippet runs unchanged on NVIDIA (CUDA), AMD (ROCm) and Intel (XPU) GPUs, `device_map="auto"` places the model on whichever accelerator is available.

After loading the model, you can do text-only inference with it as follows.

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
# Prompt
messages = [
    {"role": "user", "content": "Write a short joke about saving RAM."},
]
# Process input
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
We would need `torchvision` to be able to use images and text.

```
pip install torchvision
```
Muse Glimmer accepts images as input, as demonstrated here:

