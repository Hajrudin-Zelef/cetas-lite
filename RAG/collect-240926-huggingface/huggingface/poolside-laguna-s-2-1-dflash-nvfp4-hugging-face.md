---
id: collect-240926-huggingface/huggingface/poolside-laguna-s-2-1-dflash-nvfp4-hugging-face
title: "poolside-laguna-s-2-1-dflash-nvfp4-hugging-face"
domain: huggingface
role: reference
task: reference
actors: ["OpenRouter", "SGLang", "vLLM"]
dates: []
keywords: ["nvfp4", "latency", "sglang", "speculative decoding", "vllm"]
source: docs/RAG/clean_en/huggingface/poolside-laguna-s-2-1-dflash-nvfp4-hugging-face.md
source_anchor: ""
source_lines: [1, 31]
sha256: 750edb43f2c6969b4eb42ed7f212d1d469615ba022d40f04907481d5df081dec
---

# poolside-laguna-s-2-1-dflash-nvfp4-hugging-face

<!-- source: https://huggingface.co/poolside/Laguna-S-2.1-DFlash-NVFP4 -->

**Use on OpenRouter** ·
  **Use on Vercel AI Gateway** ·
  **Release blog post**

DFlash speculator for the NVFP4 target poolside/Laguna-S-2.1-NVFP4. The speculator is a 6-layer Laguna-style draft model (BF16); pair it with the NVFP4 base for lower-latency serving via speculative decoding.

Trained: `e0630_rhiemann_baseline` SFT, DFlash Stage-2, 15k steps. Recommended
serving setting: `num_speculative_tokens=7`.
DFlash upstream support is in progress (vLLM #46853, SGLang #29446, TRT-LLM #15666). Use
`poolside/Laguna-S-2.1-NVFP4` as the target model.

Measured with TP=2, `temperature=0`, and `num_speculative_tokens=15`.

| Concurrency | GSM8K | MATH-500 | HumanEval | MBPP | MT-Bench | 
|---|---|---|---|---|---|
| 1 | 3.324x | 2.893x | 3.692x | 2.426x | 2.338x | 
| 4 | 2.498x | 2.174x | 2.742x | 1.848x | 1.831x | 
| 8 | 2.279x | 1.948x | 2.634x | 1.719x | 1.772x | 
| 16 | 2.302x | 1.965x | 2.626x | 1.731x | 1.672x | 

| Concurrency | GSM8K | MATH-500 | HumanEval | MBPP | MT-Bench | 
|---|---|---|---|---|---|
| 1 | 5.775 | 4.942 | 6.438 | 4.171 | 4.017 | 
| 4 | 5.804 | 4.935 | 6.298 | 4.161 | 4.091 | 
| 8 | 5.719 | 4.885 | 6.537 | 4.193 | 4.373 | 
| 16 | 5.758 | 4.896 | 6.412 | 4.123 | 3.981 | 

- Downloads last month
- 42,033
