---
id: collect-huggingface/huggingface/minimaxai-minimax-m2-7
title: "MiniMax-M2.7 - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Anthropic", "Hugging Face", "MiniMax", "Nvidia", "OpenAI", "SGLang", "vLLM"]
dates: ["2026-09-23"]
keywords: ["agent", "agentic", "benchmarks", "fp8", "incident", "license", "memory", "moe", "nvidia", "open-weight", "opus 4", "reasoning"]
source: docs/RAG/Collect RAG/03_huggingface/MiniMaxAI-MiniMax-M2.7.md
source_anchor: ""
source_lines: [1, 48]
sha256: 90e99240f4f4eb091138e16ddce164204502ecd67646e2b5fd9f086324feb685
---

# MiniMax-M2.7 - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/MiniMaxAI/MiniMax-M2.7
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

MiniMax-M2.7 is MiniMax's first model deeply participating in its own evolution: during development the model updated its own memory, built dozens of complex skills for RL experiments, and improved its own learning process from experiment results. An internal version autonomously optimized a programming scaffold over 100+ rounds (analyzing failures, modifying code, running evals, deciding keep/revert), achieving a 30% performance improvement. On MLE Bench Lite (22 ML competitions) it reached a 66.6% medal rate, second only to Opus-4.6 and GPT-5.4. It is a text-generation MoE (minimax_m2 architecture, custom code) with 229B hub params in F32/BF16/F8_E4M3.

Professional software engineering: SWE-Pro 56.22% (matching GPT-5.3-Codex), SWE Multilingual 76.5, Multi SWE Bench 52.7, VIBE-Pro 55.6% (near Opus 4.6), Terminal Bench 2 57.0%, NL2Repo 39.8%. It supports native Agent Teams for multi-agent collaboration with stable role identity and autonomous decision-making, and shows SRE-level system reasoning (correlating monitoring metrics, trace analysis, root-cause verification in databases); production incident recovery reduced to under three minutes. Professional work: ELO 1495 on GDPval-AA (highest among open-weight models, surpassing GPT-5.3), Toolathon 46.3% accuracy (global top tier), 97% skill compliance across 40+ complex skills on MM Claw, MM Claw end-to-end 62.7% (close to Sonnet 4.6); Word/Excel/PPT high-fidelity multi-round editing with editable deliverables. Entertainment: strengthened character consistency and emotional intelligence; OpenRoom open-sourced interactive demo.

License: "other" (MiniMax M2.7 LICENSE, per the card's LICENSE file). Deployment: SGLang, vLLM, Transformers, ModelScope, NVIDIA NIM endpoint; recommended temperature=1.0, top_p=0.95, top_k=40, default system prompt "You are a helpful assistant. Your name is MiniMax-M2.7 and is built by MiniMax." Hub reports 1.38M monthly downloads, 115 community quantizations, 27 finetunes.

## Key points

- First MiniMax model engaging in self-evolution during training; 30% scaffold improvement autonomously.
- MLE Bench Lite 66.6% medal rate (second only to Opus-4.6 and GPT-5.4).
- SWE-Pro 56.22% (matches GPT-5.3-Codex); SWE Multilingual 76.5; Multi SWE Bench 52.7.
- Native Agent Teams for multi-agent collaboration with stable role identity.
- GDPval-AA ELO 1495, highest among open-weight models.
- 97% skill compliance across 40+ complex skills on MM Claw.
- "Other" license (MiniMax M2.7 LICENSE); F32/BF16/FP8 weights.
- Deploy via SGLang, vLLM, Transformers, ModelScope, NVIDIA NIM.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | MiniMaxAI |
| Model name | MiniMax-M2.7 |
| Architecture | MoE (minimax_m2, custom code) |
| Total params | 229B (hub-reported) |
| License | other (MiniMax M2.7 LICENSE) |
| Weight formats | F32, BF16, F8_E4M3 |
| Benchmarks | SWE-Pro 56.22, SWE Multilingual 76.5, Multi SWE Bench 52.7, MLE Bench Lite 66.6% medal, GDPval-AA ELO 1495 |
| Sampling | temperature=1.0, top_p=0.95, top_k=40 |
| Deployment | SGLang, vLLM, Transformers, ModelScope, NVIDIA NIM |
| Downloads/month | 1,381,521 |

## Why this source matters for the RAG

This card documents a model trained with self-evolution loops, providing a unique data point on autonomous scaffold optimization and strong multi-agent/office-work capabilities. It is essential reference for retrieval on self-improving models, multi-agent collaboration, agentic RL, and software-engineering benchmarks.
