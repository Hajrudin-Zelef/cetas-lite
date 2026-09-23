---
id: collect-huggingface/huggingface/minimaxai-minimax-m2-5
title: "MiniMax-M2.5 - Hugging Face model card"
domain: huggingface
role: reference
task: model-card
actors: ["Anthropic", "Hugging Face", "MiniMax", "SGLang", "vLLM"]
dates: ["2026-09-23"]
keywords: ["agent", "agentic", "benchmark", "benchmarks", "claude", "cost", "distribution", "fp8", "license", "mit license", "moe", "open-weight"]
source: docs/RAG/Collect RAG/03_huggingface/MiniMaxAI-MiniMax-M2.5.md
source_anchor: ""
source_lines: [1, 50]
sha256: 40c57738dc4da1a21c8a41c87a745aa3e1b794f6e98916eca26913fe62ffd1ae
---

# MiniMax-M2.5 - Hugging Face model card

## Metadata

- **Source** : https://huggingface.co/MiniMaxAI/MiniMax-M2.5
- **Site** : Hugging Face
- **Type** : Model card
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

MiniMax-M2.5 is a frontier open-weight text-generation model from MiniMax, extensively trained with reinforcement learning in hundreds of thousands of complex real-world environments. It achieves SOTA results in coding, agentic tool use and search, office work, and other economically valuable tasks: 80.2% SWE-Bench Verified, 51.3% Multi-SWE-Bench, 76.3% BrowseComp (with context management). Hub reports 229B params (F32/BF16/F8_E4M3 weights). The model is a MoE in the M2 family (arXiv 2605.26494), served with a custom minimax_m2 architecture in Transformers with trust_remote_code.

Efficiency/cost focus: M2.5 completes SWE-Bench Verified 37% faster than M2.1 (22.8 min vs 31.3 min, 3.52M tokens/task), matching Claude Opus 4.6 speed at only 10% of the cost per task. It costs about $1 to run continuously for an hour at 100 tokens/s ($0.30 at 50 TPS). Two versions with identical capability: M2.5 (50 TPS) and M2.5-Lightning (100 TPS; $0.3/M input, $2.4/M output tokens). On coding harnesses, M2.5 generalizes across out-of-distribution scaffolds (Droid 79.7, OpenCode 76.1, both above Opus 4.6). The model exhibits a spec-writing tendency (architect-like planning before code) and was trained on 10+ languages across 200,000+ real environments. Search: uses ~20% fewer rounds than M2.1 on BrowseComp/Wide Search/RISE (a new realistic interactive search benchmark). Office work: GDPval-MM pairwise win rate 59.0%, Word/PPT/Excel financial modeling, MEWC.

RL scaling: in-house Forge agent-native RL framework (~40x training speedup via tree-structured sample merging), CISPO algorithm for MoE stability, process reward mechanism for long-context credit assignment, and time-based rewards. Appendix benchmarks: AIME25 86.3, GPQA-D 85.2, HLE w/o tools 19.4, SciCode 44.4, IFBench 70.0. License: Modified-MIT. Deployment via SGLang, vLLM, Transformers, KTransformers, ModelScope; recommended parameters temperature=1.0, top_p=0.95, top_k=40 with a default system prompt.

## Key points

- Frontier agentic model: SWE-Bench Verified 80.2%, Multi-SWE-Bench 51.3%, BrowseComp 76.3%.
- RL trained across hundreds of thousands of real-world environments.
- 37% faster than M2.1 on SWE-Bench Verified; matches Claude Opus 4.6 runtime at 10% of cost.
- Two versions (M2.5 50 TPS / M2.5-Lightning 100 TPS); ~$1/hour to run at 100 TPS.
- Spec-writing (architect-like planning) behavior emerged during training.
- Forge agent-native RL framework, CISPO algorithm, process rewards.
- Modified-MIT license; 229B hub params (F32/BF16/FP8).
- Deploy via SGLang, vLLM, Transformers, KTransformers; temp 1.0, top_p 0.95, top_k 40.

## Technical data / figures

| Attribute | Value |
|---|---|
| Organization | MiniMaxAI |
| Model name | MiniMax-M2.5 |
| Architecture | MoE (minimax_m2, custom code) |
| Total params | 229B (hub-reported) |
| Context length | 1M-class (M2 family) |
| License | Modified-MIT |
| Weight formats | F32, BF16, F8_E4M3 |
| Benchmarks | SWE-Bench Verified 80.2, Multi-SWE-Bench 51.3, BrowseComp 76.3, AIME25 86.3, GPQA-D 85.2 |
| Pricing | M2.5-Lightning $0.3/M in, $2.4/M out |
| Sampling | temperature=1.0, top_p=0.95, top_k=40 |
| Deployment | SGLang, vLLM, Transformers, KTransformers, ModelScope |
| Downloads/month | 293,707 |

## Why this source matters for the RAG

This card documents a frontier open-weight agentic model that combines massive RL scaling with extreme cost efficiency, including detailed benchmark tables, RL methodology (Forge, CISPO), and economics. It is essential reference material for retrieval on agentic RL, coding models, and cost-efficient open-weight LLMs.
