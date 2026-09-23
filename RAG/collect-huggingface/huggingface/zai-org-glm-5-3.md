---
id: collect-huggingface/huggingface/zai-org-glm-5-3
title: "GLM-5.3 - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Alibaba", "Anthropic", "DeepSeek", "ExploitGym", "Huawei", "Hugging Face", "Moonshot", "Nvidia", "OpenAI", "SGLang", "Unsloth", "Z.ai", "vLLM"]
dates: ["2026-09-23"]
keywords: ["glm", "agent", "agentic", "agents", "ascend", "attention", "benchmark", "benchmarks", "claude", "cyber", "deepseek", "fable 5"]
source: docs/RAG/Collect RAG/03_huggingface/zai-org-GLM-5.3.md
source_anchor: ""
source_lines: [1, 55]
sha256: 47be5b2643dc70c0d17a68fc7879e5a28b2fb3450e773495b3991164ddeedb69
---

# GLM-5.3 - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/zai-org/GLM-5.3
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

GLM-5.3 is Z.ai's flagship open-weight language model in the GLM-5 series, sharing the same base model as GLM-5.2 — every gain comes from post-training. It is released under a custom GLM-5.3 license, reported at 753B params (BF16/F8_E4M3/F32), architecture `glm_moe_dsa` (MoE with dynamic sparse attention). It is far better than GLM-5.2 at complex coding and long-horizon tasks, and is the most capable open-weights model for coding, with a 50% improvement over GLM-5.2 on Z.ai's in-house Code Bench, plus open-source SOTA on Terminal Bench 3.0 and Agents' Last Exam.

It also shows an "emergent cyber capability": state of the art on CyberGym for vulnerability discovery, more than doubling GLM-5.2 on exploitation benchmarks (ExploitGym 2h/6h: 105/130 vs 29/39; ExploitBench 54.4 vs 24.4).

Benchmarks vs competitors (GLM-5.2, Kimi K3, DeepSeek-V4 Pro-0813, Qwen3.8-Max, Opus 4.8, Fable 5, GPT-5.6 Sol): Terminal Bench 2.1 88.2, Terminal Bench 3.0 28.3, DeepSWE (v1.1) 66.9, NL2Repo 58.0, FrontierSWE 78.1, SWE-Marathon 42.5, PostTrainBench 39.8, CyberGym 84.5, Toolathlon Verified 73.0, AutomationBench 48.2, Agents' Last Exam 28.5, HLE w/ Tools 62.5, GDPval-AA v2 1769. Detailed footnotes specify harnesses (Claude Code 2.1.207, mini-swe-agent, etc.) and context lengths up to 1M.

Deployment: SGLang, vLLM, TokenSpeed, Transformers, KTransformers, Unsloth, and Ascend NPU platforms (vLLM-Ascend, xLLM, SGLang). `reasoning_effort` (low/high/max, default max) controls thinking budget; `clear_thinking=true` recommended for chat. ~1.08M monthly downloads.

## Key points

- GLM-5.3 = GLM-5.2 base + post-training; 753B params MoE (glm_moe_dsa).
- Most capable open-weights coding model; +50% over GLM-5.2 on Z.ai Code Bench.
- Open-source SOTA on Terminal Bench 3.0 and Agents' Last Exam.
- Emergent cyber capability: SOTA on CyberGym (84.5); doubles GLM-5.2 on ExploitBench.
- Terminal Bench 2.1 88.2; DeepSWE 66.9; HLE w/ tools 62.5.
- Custom GLM-5.3 license; deployable on NVIDIA + Ascend NPU stacks.

## Technical data / figures

| Attribute | Value |
|---|---|
| Total parameters | 753B (reported) |
| Architecture | MoE with dynamic sparse attention (glm_moe_dsa) |
| Base model | Same as GLM-5.2 (post-training only) |
| License | GLM-5.3 (custom) |
| Precision | BF16 / F8_E4M3 / F32 |
| reasoning_effort | low / high / max (default max) |
| Terminal Bench 2.1 | 88.2 |
| Terminal Bench 3.0 | 28.3 |
| DeepSWE (v1.1) | 66.9 |
| NL2Repo | 58.0 |
| FrontierSWE | 78.1 |
| CyberGym | 84.5 |
| ExploitBench | 54.4 |
| Toolathlon Verified | 73.0 |
| Agents' Last Exam | 28.5 |
| HLE w/ Tools | 62.5 |
| Monthly downloads | ~1.08M |

## Why this source matters for the RAG

GLM-5.3 represents the frontier of open-weight coding/agentic models, with exhaustive benchmark data against closed and open competitors, which is essential for capability comparisons. Its detailed evaluation footnotes clarify methodology (harness, context, anti-cheat) and prevent hallucination about scores. It anchors the RAG's coverage of top-tier open-weight models in 2026.
