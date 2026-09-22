---
id: frontier-models-2026/06-cross-cutting-comparison-vague-1-models/overview
title: "4. Cross-cutting comparison (vague 1 models)"
domain: cross-cutting-comparison-vague-1-models
role: deep-dive
task: reference
actors: ["DeepSeek", "Z.ai", "xAI"]
dates: []
keywords: ["agent", "attention", "cyber", "deepseek", "fp4", "glm", "grok", "grok 4", "memory", "multimodal", "training"]
source: docs/RAG/Grands titres IA modèlesEN.md
source_anchor: ""
source_lines: [370, 383]
section: "4. Cross-cutting comparison (vague 1 models)"
sha256: e764f7be05e82beb003145713b9f770dec313a2fe0e3af7314cecc6cf969c02c
---

# 4. Cross-cutting comparison (vague 1 models)

| Dimension | Grok 4.20 | DeepSeek V4-Pro | DeepSeek V4.1 Flash | GLM-5.2 | GLM-5.3 | GLM-5.3-Flash |
|---|---|---|---|---|---|---|
| Release | 17 Feb 2026 (beta) | ~17 Feb 2026 | 10 Sep 2026 | 13–16 Jun 2026 | 14 Aug 2026 | 26 Aug 2026 |
| Weights | Closed | Open (MIT) | Open (MIT) | Open (MIT) | Open (MIT, staged) | Open (MIT) |
| Params (total/active) | n.d. | 1.6T / 49B | 552B / 8B–16B (+196B Engram) | 744B / 40B | ~743B / ~40B | 320B / 18B |
| Context | 256K (→2M) | 1M | 1M (384K out) | 1M | 1M | 1M |
| Multimodal | text+image+video | vision (V4.1) | native vision | text only | text (coding) | native image+video |
| Signature tech | 4-agent council, adversarial consensus | Engram memory, mHC, Muon | Causal Encoder-Decoder, FP4 KV, CSA2 | IndexShare attention | post-training scaling, cyber training | sparse+linear attention, CN-chip serving |
| SWE-bench Verified | n.d. | 80.6% | — | — (Pro: 62.1%) | — | — |
| Terminal-Bench | n.d. | 67.9% (2.0) | 90.6 (2.1) | 81.0 (2.1) | 28.3 (3.0) | 84.3 (2.1) |
| API in/out ($/1M) | 1.25 / 2.50 | 1.74 / 3.48 | 0.15 / 0.60 (off-peak) | 1.40 / 4.40 | n.d. (Coding Plan) | 0.15 / 0.50 |

