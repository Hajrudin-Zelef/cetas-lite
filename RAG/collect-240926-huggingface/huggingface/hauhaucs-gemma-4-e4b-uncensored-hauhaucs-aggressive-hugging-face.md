---
id: collect-240926-huggingface/huggingface/hauhaucs-gemma-4-e4b-uncensored-hauhaucs-aggressive-hugging-face
title: "Text only"
domain: huggingface
role: reference
task: reference
actors: ["Google", "Hugging Face", "Nvidia"]
dates: []
keywords: ["attention", "gguf", "guardrails", "llama", "llama.cpp", "memory", "multimodal", "nvidia", "parameters", "quantization", "refusals", "training"]
source: docs/RAG/clean_en/huggingface/hauhaucs-gemma-4-e4b-uncensored-hauhaucs-aggressive-hugging-face.md
source_anchor: ""
source_lines: [1, 74]
sha256: 75f570015184e5015950c272fadc211f6ad89d8aec9132f2fed07622d12a5428
---

# Text only

<!-- source: https://huggingface.co/HauhauCS/Gemma-4-E4B-Uncensored-HauhauCS-Aggressive -->

**Join the Discord** for updates, roadmaps, projects, or just to chat.


Gemma 4 E4B-IT uncensored by HauhauCS. **0/465 Refusals***

**HuggingFace's "Hardware Compatibility" widget doesn't recognize K_P quants** — it may show fewer files than actually exist. Click **"View +X variants"** or go to **Files and versions** to see all available downloads.


No changes to datasets or capabilities. Fully functional, 100% of what the original authors intended - just without the refusals.

These are meant to be the best lossless uncensored models out there.

Stronger uncensoring — model is fully unlocked and won't refuse prompts. May occasionally append short disclaimers (baked into base model training, not refusals) but full content is always generated.

For a more conservative uncensor that keeps some safety guardrails, check the Balanced variant when it's available.

| File | Quant | BPW | Size | 
|---|---|---|---|
| Gemma-4-E4B-Uncensored-HauhauCS-Aggressive-Q8_K_P.gguf | Q8_K_P | 9.4 | 7.6 GB | 
| — | Q8_0 | 8.5 | — | 
| Gemma-4-E4B-Uncensored-HauhauCS-Aggressive-Q6_K_P.gguf | Q6_K_P | 7.0 | 5.9 GB | 
| — | Q6_K | 6.6 | — | 
| Gemma-4-E4B-Uncensored-HauhauCS-Aggressive-Q5_K_P.gguf | Q5_K_P | 6.1 | 5.5 GB | 
| Gemma-4-E4B-Uncensored-HauhauCS-Aggressive-Q5_K_M.gguf | Q5_K_M | 5.7 | 5.4 GB | 
| Gemma-4-E4B-Uncensored-HauhauCS-Aggressive-Q4_K_P.gguf | Q4_K_P | 5.2 | 5.1 GB | 
| Gemma-4-E4B-Uncensored-HauhauCS-Aggressive-Q4_K_M.gguf | Q4_K_M | 4.8 | 5.0 GB | 
| Gemma-4-E4B-Uncensored-HauhauCS-Aggressive-IQ4_XS.gguf | IQ4_XS | 4.3 | 4.8 GB | 
| Gemma-4-E4B-Uncensored-HauhauCS-Aggressive-Q3_K_P.gguf | Q3_K_P | 4.1 | 4.6 GB | 
| Gemma-4-E4B-Uncensored-HauhauCS-Aggressive-Q3_K_M.gguf | Q3_K_M | 3.9 | 4.6 GB | 
| Gemma-4-E4B-Uncensored-HauhauCS-Aggressive-IQ3_M.gguf | IQ3_M | 3.7 | 4.4 GB | 
| Gemma-4-E4B-Uncensored-HauhauCS-Aggressive-Q2_K_P.gguf | Q2_K_P | 3.5 | 4.2 GB | 
| mmproj-Gemma-4-E4B-Uncensored-HauhauCS-Aggressive-f16.gguf | mmproj (f16) | — | 945 MB | 

All quants generated with importance matrix (imatrix) for optimal quality preservation on abliterated weights.

K_P ("Perfect") quants are HauhauCS custom quantizations that use model-specific analysis to selectively preserve quality where it matters most. Each model gets its own optimized quantization profile.

A K_P quant effectively bumps quality up by 1-2 quant levels at only ~5-15% larger file size than the base quant. Fully compatible with llama.cpp, LM Studio, and any GGUF-compatible runtime — no special builds needed.

**Note:** K_P quants may show as "?" in LM Studio's quant column. This is a display issue only — the model loads and runs fine.

- 4B parameters
- 42 layers, mixed sliding window (512) + full attention
- 131K context
- Natively multimodal (text, image, video, audio)
- 18 KV shared layers for memory efficiency
- Based on google/gemma-4-e4b-it

From the official Google Gemma 4 authors:

- `temperature=1.0, top_p=0.95, top_k=64`

**Important:**

- Use `--jinja` flag with llama.cpp for proper chat template handling
- Vision/audio support requires the `mmproj` file alongside the main GGUF

Works with llama.cpp, LM Studio, Jan, koboldcpp, and other GGUF-compatible runtimes.

```
# Text only
llama-cli -m Gemma-4-E4B-Uncensored-HauhauCS-Aggressive-Q4_K_M.gguf \
  --jinja -c 8192 -ngl 99
# With vision/audio
llama-cli -m Gemma-4-E4B-Uncensored-HauhauCS-Aggressive-Q4_K_M.gguf \
  --mmproj mmproj-Gemma-4-E4B-Uncensored-HauhauCS-Aggressive-f16.gguf \
  --jinja -c 8192 -ngl 99
```
***** Gemma 4 didn't get as much manual testing time at longer context as my other releases. Google is now using techniques similar to NVIDIA's GenRM — generative reward models that act as internal critics — making (true) uncensoring an increasingly challenging field. I expect 99.999% of users won't hit edge cases, but the asterisk is there for honesty.

- Downloads last month
- 1,605,760
