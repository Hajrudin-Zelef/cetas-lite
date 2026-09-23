---
id: vague2-vision-ia/vision-ia/nvidia-e-gale-le-mode-le-ouvert-d-openai-avec-4-fois-moins-de-parame-tres
title: "Nvidia égale le modèle ouvert d'OpenAI avec 4 fois moins de paramètres"
domain: vision-ia
role: reference
task: article
actors: ["Alibaba", "Anthropic", "CoreWeave", "Fireworks AI", "Google", "Hugging Face", "Meta", "Microsoft", "Nebius", "Nvidia", "OpenAI"]
dates: ["2026-09-23"]
keywords: ["nvidia", "agentic", "agents", "apache", "attention", "claude", "context window", "cost", "gemini", "ipo", "latency", "license"]
source: docs/RAG/Collect RAG Vague 2/04_vision_ia/nvidia-e-gale-le-mode-le-ouvert-d-openai-avec-4-fois-moins-de-parame-tres.md
source_anchor: ""
source_lines: [1, 43]
sha256: 9756ef37a1069f911ba0bfd9b4f28529ee7feb606d000bf18998d9e3f52007a1
---

# Nvidia égale le modèle ouvert d'OpenAI avec 4 fois moins de paramètres

## Metadata

- **Source** : https://vision-ia.beehiiv.com/p/nvidia-e-gale-le-mode-le-ouvert-d-openai-avec-4-fois-moins-de-parame-tres
- **Site** : Vision-IA (beehiiv)
- **Type** : Newsletter article
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This Vision-IA newsletter (Aug 12, 2026) leads with Nvidia's release of Nemotron 3.5 Lightning, an open-weight model with 30 billion total parameters of which only 3.6 billion are active per request. It matches OpenAI's gpt-oss-120b on the Artificial Analysis Intelligence Index despite being four times smaller, and generates 669 tokens per second — the fastest in its tested category. It is a hybrid architecture combining Mamba-2 layers, a Mixture-of-Experts system and some standard attention layers, text-only, with a one-million-token context window. It scores an Intelligence Index of 24 (vs 15 for Nemotron 3 Nano, nine points gained in one generation and parity with gpt-oss-120b), 86% accuracy on PinchBench with 30% more speed than Qwen3.6 35B at comparable accuracy, and 24.3% on Terminal-Bench v2.1 vs 26.2% for OpenAI's model. Weights are available in BF16 and NVFP4 on Hugging Face under the permissive OpenMDW-1.1 license with commercial use, hosted via NVIDIA NIM, DeepInfra, Fireworks, CoreWeave or Nebius, plus NeMo Switchyard, an in-house router. Nvidia frames it for autonomous agents running for hours and making hundreds of tool calls, where cost and latency matter more than the last IQ point. Other headlines: Google DeepMind leadership change — Koray Kavukcuoglu becomes SVP of Google DeepMind reporting to the Alphabet CEO, while Demis Hassabis becomes DeepMind president and Alphabet Chief Scientist; Jeff Dean leaves after 27 years, joined by Sanjay Ghemawat, Oriol Vinyals and Quoc Le; John Jumper and two AlphaFold colleagues left for Anthropic; Google hasn't shipped a frontier model since Gemini 3.1 Pro in February. Alibaba released Qwen-MM-Plugins (Apache-2.0), adding multimodal capabilities to Claude Code, Codex, Qoder, OpenClaw, Qwen Code and Gemini CLI. English courts now confiscate Meta Glasses at the entrance (since Aug 11). Also: lab-grown brain organoids, Google's AMIE medical AI doing video consultations, extraction of hidden model reasoning, Microsoft's CARE-X chest X-ray model, Anthropic's potential IPO at a $965B valuation, and more.

## Key points

- Nemotron 3.5 Lightning: 30B total / 3.6B active params, matches gpt-oss-120b on the Intelligence Index while being 4× smaller.
- Generates 669 tokens/second; hybrid Mamba-2 + MoE + attention architecture; 1M-token context; text-only.
- Open weights (BF16, NVFP4) on Hugging Face under OpenMDW-1.1 with commercial use; NeMo Switchyard router included.
- Google DeepMind leadership reshuffle: Kavukcuoglu leads, Hassabis becomes president/Chief Scientist, Jeff Dean departs.
- Qwen-MM-Plugins adds vision/audio tools to six coding agents under Apache-2.0.
- Meta Glasses banned/confiscated in all English and Welsh courts since Aug 11, 2026.

## Technical data / figures

| Item | Value |
|---|---|
| Nemotron 3.5 Lightning params | 30B total / 3.6B active |
| Intelligence Index | 24 (gpt-oss-120b parity; Nemotron 3 Nano 15) |
| Generation speed | 669 tokens/s |
| Context window | 1 million tokens |
| PinchBench accuracy | 86% |
| Terminal-Bench v2.1 | 24.3% (OpenAI 26.2%) |
| License | OpenMDW-1.1 (commercial use) |
| Meta Glasses sales 2025 | 7 million pairs |
| Meta Glasses price | from $299 |
| Anthropic IPO valuation (reported) | $965B |
| Gemini monthly users | 1 billion |

## Why this source matters for the RAG

It demonstrates the efficiency frontier in open models — matching a much larger OpenAI model with 4× fewer parameters and far lower serving cost — a key data point for cost-sensitive agentic deployments. It also documents a major leadership shake-up at Google DeepMind and the talent exodus to competitors, central to tracking the competitive balance among frontier labs.
