---
id: collect-240926-huggingface/huggingface/deepseek-ai-deepseek-v4-flash-vision-exp-hugging-face
title: "deepseek-ai-deepseek-v4-flash-vision-exp-hugging-face"
domain: huggingface
role: reference
task: reference
actors: ["DeepSeek", "Hugging Face", "OpenAI", "SGLang", "vLLM"]
dates: []
keywords: ["deepseek", "agent", "agents", "attention", "benchmark", "benchmarks", "fp8", "gpus", "inference", "leaderboard", "license", "mit license"]
source: docs/RAG/clean_en/huggingface/deepseek-ai-deepseek-v4-flash-vision-exp-hugging-face.md
source_anchor: ""
source_lines: [1, 93]
sha256: 3113b873ccc99ffd0cd148275db4ad4f4ba07e29dd1a6535cf9e93512b9f04fc
---

# deepseek-ai-deepseek-v4-flash-vision-exp-hugging-face

<!-- source: https://huggingface.co/deepseek-ai/DeepSeek-V4-Flash-Vision-Exp -->

We are excited to introduce **DeepSeek-V4-Flash-Vision-Exp**, our first experimental multimodal model in the DeepSeek-V4 family. It builds on the DeepSeek-V4-Flash architecture by incorporating visual modules and undergoing continued training to unlock visual understanding capabilities.

Compared to DeepSeek-V4-Flash-0731, DeepSeek-V4-Flash-Vision-Exp achieves substantial improvements on its multimodal agent capabilities, while maintaining comparable performance on text-only agent tasks.

| Benchmark | DeepSeek-V4-Flash-Vision-Exp | DeepSeek-V4-Flash-0731 | Opus-4.8 | 
|---|---|---|---|
| **Text Agent Capabilities** |  |  |  | 
| Terminal Bench 2.1 | 83.9 | 82.7 | 85.0 | 
| NL2Repo | 57.7 | 54.2 | 69.7 | 
| Cybergym | 75.3 | 76.7 | 78.3 | 
| DeepSWE | 59.3 | 54.4 | 58.0 | 
| Toolathlon-Verified | 75.9 | 70.3 | 76.2 | 
| DSBench-Hard | 63.6 | 59.6 | 71.7 | 
| AutomationBench (Public) | 25.7 | 25.1 | 27.2 | 
| **Multimodal Agent Capabilities** |  |  |  | 
| ApexBench (Pass@1) | 36.5 | 26.2† | 39.4 | 
| Agents' Last Exam | 27.3 | 25.2† | 25.7 | 
| Chartography | 64.3 | - | 65.0 | 
| ZeroBench (Pass@5) | 35.0 | - | 34.0 | 

Notes:

1. For the text agent benchmarks above, DeepSeek models are evaluated with the minimal mode of DeepSeek Harness as the agent framework, using the `max` reasoning effort level with`temperature = 1.0, top_p = 0.95` .
2. † For ApexBench and Agents' Last Exam, DeepSeek-V4-Flash-0731 ignores the multimodal elements in the input.

This repository contains the tokenizer, prompt encoding reference, and a minimal PyTorch inference implementation for DeepSeek-V4 Flash Vision. The reference inference covers the vision encoder and aligner, DFlash attention, MoE, Hyper-Connections, and the DSpark forward path.

```
.
├── encoding/                  # OpenAI-style messages -> model prompt
├── inference/                 # weight conversion and minimal inference
│   └── examples/              # equivalent TXT and JSON vision prompts
├── config.json                # Hugging Face model metadata
├── generation_config.json
├── model.safetensors.index.json
├── tokenizer.json
└── tokenizer_config.json
```
`encoding/` and `inference/` deliberately remain separate: prompt formatting
does not depend on PyTorch, while inference imports the sibling encoding module
with an explicit Python path. No symlinks are required.

The tokenizer files are regular files so that the repository can be uploaded
to Hugging Face without relying on local filesystem symlinks. The large model
shards are described by `model.safetensors.index.json` and are not duplicated
inside the source checkout used to assemble this repository.

See `encoding/README.md`. Both OpenAI-style JSON content
blocks and the compact `<image>path</image>` TXT notation are supported. The two
examples under `inference/examples/` encode to identical prompts and token IDs.

See `inference/README.md` for dependency installation,
checkpoint conversion, and TXT/JSON inference commands.

For example, the command below serves the model with vLLM on a single 4×GB300 node. See the vLLM recipe for detailed instructions and other hardware configurations.

```
docker run --gpus all \
  vllm/vllm-openai:deepseekv4-flash-vision deepseek-ai/DeepSeek-V4-Flash-Vision-Exp \
  --kv-cache-dtype fp8 \
  --block-size 256 \
  --tensor-parallel-size 4 \
  --tool-call-parser deepseek_v4 \
  --enable-auto-tool-choice \
  --reasoning-parser deepseek_v4 \
  --reasoning-config '{"reasoning_parser":"deepseek_v4","reasoning_start_str":"","reasoning_end_str":""}' \
  --speculative-config '{"method":"dspark","model":"deepseek-ai/DeepSeek-V4-Flash-Vision-Exp","num_speculative_tokens":3,"draft_sample_method":"probabilistic","enable_adaptive_verification":true}'
```
Enable DSpark with --speculative-algorithm DSPARK and do not set a separate --speculative-draft-model-path as the target and draft weights therefore come from the same checkpoint. See the SGLang cookbook for detailed instructions, benchmarks and other hardwares configurations.

```
sglang serve \
  --model-path deepseek-ai/DeepSeek-V4-Flash-Vision-Exp \
  --tp 4 \
  --speculative-algorithm DSPARK \
  --mem-fraction-static 0.85 \
  --host 0.0.0.0 \
  --port 30000
```
This repository is licensed under the MIT License.

- Downloads last month
- 889,035

## Spaces using deepseek-ai/DeepSeek-V4-Flash-Vision-Exp 5

## Collection including deepseek-ai/DeepSeek-V4-Flash-Vision-Exp

- datacurve/deep-swe · Deep Swe View evaluation results leaderboard
- harborframework/terminal-bench-2.1 · Terminalbench 2 1 View evaluation results    leaderboard  83.9
- hkust-nlp/Toolathlon · Toolathlon Verified View evaluation results
