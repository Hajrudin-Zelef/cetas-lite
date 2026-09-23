---
id: collect-huggingface/huggingface/deepseek-ai-deepseek-v4-flash-0731
title: "DeepSeek-V4-Flash-0731 - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["DeepSeek", "Hugging Face", "OpenAI", "SGLang", "Z.ai", "vLLM"]
dates: ["2026-09-23"]
keywords: ["deepseek", "agent", "agentic", "agents", "benchmark", "benchmarks", "fp4", "fp8", "glm", "inference", "kv cache", "license"]
source: docs/RAG/Collect RAG/03_huggingface/deepseek-ai-DeepSeek-V4-Flash-0731.md
source_anchor: ""
source_lines: [1, 52]
sha256: 2654ccba89faf10b83b53a20248fc4e713c007c919378f7124465956f973daed
---

# DeepSeek-V4-Flash-0731 - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/deepseek-ai/DeepSeek-V4-Flash-0731
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

DeepSeek-V4-Flash-0731 is the official release of DeepSeek-V4-Flash, superseding the earlier preview version and delivering substantially enhanced agentic capabilities. It shares the same model structure as DeepSeek-V4-Flash-DSpark, meaning it comes with an attached speculative decoding module (DSpark) for faster generation. The model is a text-generation MoE with 304B parameters reported on the hub (BF16 weights, plus I64/F32/F8_E4M3/I8 tensor types). It outperforms DeepSeek-V4-Pro (Preview) on the listed benchmarks despite a far smaller activated parameter count, and is described as broadly competitive with the strongest proprietary models available.

Benchmark results are the focus of the card. Compared with the preview and with proprietary competitors (GLM-5.2, Opus-4.8), DeepSeek-V4-Flash-0731 scores: Terminal Bench 2.1 = 82.7 (preview 61.8; V4-Pro 72.1), NL2Repo = 54.2, Cybergym = 76.7, DeepSWE = 54.4 (preview 7.3), Toolathlon-Verified = 70.3, Agents' Last Exam = 25.2, AutomationBench Public = 25.1, and internal DSBench-FullStack = 68.7 / DSBench-Hard = 59.6. Code-agent tasks were evaluated with the minimal mode of DeepSeek Harness at `max` reasoning effort with temperature = 1.0 and top_p = 0.95.

The release does not ship a Jinja chat template; instead an `encoding` folder provides Python scripts and test cases to encode OpenAI-compatible messages and parse output. The `reasoning_effort` parameter supports three levels — `low`, `high`, and `max`. DSpark speculative decoding is enabled in vLLM with a `--speculative-config` flag (`{"method":"dspark","num_speculative_tokens":7,"draft_sample_method":"greedy"}`), and in SGLang with `--speculative-algorithm DSPARK` (no separate draft model path needed since target and draft weights come from the same checkpoint). Example vLLM (4xGB300, DP=4, expert parallel, deep_gemm_mega_moe, fp8 KV cache, fp4 indexer cache) and SGLang (tp 4, flashinfer_mxfp4, mem-fraction-static 0.90) launch commands are given. Recommended sampling is temperature 1.0, top_p 0.95 for agentic scenarios and 1.0 otherwise; for high/max reasoning effort, max output length 384K tokens. License is MIT.

## Key points

- Official release of DeepSeek-V4-Flash, superseding the preview; strongly improved agentic abilities.
- Same structure as DeepSeek-V4-Flash-DSpark: includes an attached DSpark speculative decoding module.
- Hub size 304B params (BF16 + FP8/I8 tensor types); text-generation MoE.
- Outperforms DeepSeek-V4-Pro (Preview) on listed benchmarks despite fewer activated params.
- reasoning_effort supports low / high / max levels.
- No Jinja template; dedicated `encoding` Python folder instead.
- DSpark enabled via vLLM `--speculative-config` or SGLang `--speculative-algorithm DSPARK`.
- MIT license; 3.95M downloads/month; arXiv 2606.19348.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | deepseek-ai |
| Model name | DeepSeek-V4-Flash-0731 |
| Architecture | MoE with DSpark speculative decoding module |
| Reported hub size | 304B params |
| Weight formats | BF16, I64, F32, F8_E4M3, I8 |
| Context / output | up to 384K output at high/max effort |
| Reasoning effort | low / high / max |
| License | MIT |
| Sampling | temp 1.0; top_p 0.95 (agentic) or 1.0 |
| Serving | vLLM, SGLang (DSpark flags) |
| Key benchmarks | Terminal Bench 2.1 82.7; NL2Repo 54.2; Cybergym 76.7; DeepSWE 54.4; Toolathlon-V 70.3; ALE 25.2 |
| Internal benchmarks | DSBench-FullStack 68.7; DSBench-Hard 59.6 |
| arXiv | 2606.19348 |
| Downloads/month | 3,954,272 |

## Why this source matters for the RAG

This card documents an official, heavily-downloaded agentic MoE release with detailed benchmark comparisons and concrete speculative-decoding serving recipes, making it a key source for retrieval on agent frameworks, efficient inference, and coding-agent performance. The vLLM/SGLang commands and reasoning-effort controls are directly actionable.
