---
id: ai-industry-kb-2026-wave6/15-nvidia-and-nemotron/nemotron-3-super-the-march-2026-mid-tier-flagship
title: "Nemotron 3 Super — the March 2026 mid-tier flagship"
domain: nvidia-and-nemotron
role: deep-dive
task: actor-profile
actors: ["Alibaba", "Hugging Face", "Nvidia", "OpenRouter"]
dates: ["2025-06", "2026-02", "2026-03", "2026-03-11", "2026-03-16"]
keywords: ["agentic", "attention", "benchmarks", "context window", "fp8", "license", "nvfp4", "nvidia", "parameters", "pricing", "quantization", "reasoning"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [7520, 7532]
section: "§15. NVIDIA and Nemotron"
delta_of: ai-industry-kb-2026
sha256: 6798267e7a587e1f1a8d8887080025f0d58eccd63069b89f289e4b0df6c94bab
---

# Nemotron 3 Super — the March 2026 mid-tier flagship

### Nemotron 3 Super — the March 2026 mid-tier flagship
- NVIDIA released Nemotron 3 Super (`nemotron-3-super-120b-a12b`) at GTC in March 2026 — reported as March 11, 2026 in secondary catalogs, while the base corpus dates GTC 2026 announcements to March 16, 2026; preserve both dates. [SECONDARY] [S31][S32]
- Architecture: hybrid Mamba-2 + Transformer sparse-attention + LatentMoE with a Multi-Token Prediction (MTP) head; 120B total parameters, ~12B active per token (reported as 12.7B in some catalogs — preserve both). 1M-token context window. [SECONDARY] [S31][S32][S33]
- CONTRADICTION: active parameters reported as 12B (TokenCost/awesomeagents/dev.to) versus 12.7B (other catalogs). Preserve both; they describe the same checkpoint. [SECONDARY] [S31][S32]
- Speed: 449 output tokens/sec per Artificial Analysis — ranked #1 of 51 models in its intelligence tier, 2.2x faster than GPT-OSS-120B. (single secondary coverage) [SECONDARY] [S31]
- Long-context: 91.75% on RULER at full 1M length; 96.30% at 256K, 95.67% at 512K. The 1M window defaults to 256K in the shipped configuration due to VRAM constraints; extending to the full million requires `VLLM_ALLOW_LONG_MAX_MODEL_LEN=1`. [SECONDARY] [S32][S36]
- Benchmarks (vendor-reported, via community): SWE-Bench Verified 60.47% (vs GPT-OSS-120B 41.90%), MMLU-Pro 83.73%, AIME 2025 90.21% (no tools), HMMT Feb 2025 93.67% (no tools), LiveCodeBench 81.19%, GPQA 79.23 (82.70% with tools), SciCode subtask 42.05%. Artificial Analysis: Intelligence Index 35.97, Coding 31.19, Agentic 40.18; GPQA Diamond 80.0%, HLE 19.2%, SciCode 36.0%, TAU-2 67.8%, IFBench 71.5%, LCR 60.0%. (single vendor coverage) [VENDOR] [S36]
- Competitive position is mixed: it beats GPT-OSS-120B on most tasks but trades wins with Qwen3.5-122B-A10B — leading on long-context retrieval (RULER@1M 91.75%), HMMT math and LiveCodeBench, trailing on general knowledge (MMLU-Pro 83.73 vs 86.70) and science reasoning (GPQA 79.23 vs 86.60), and on agentic coding (SWE-Bench 60.47 vs 66.40). (single secondary coverage) [SECONDARY] [S32]
- Training: 15.6T tokens, 153 datasets, 20 languages, 43 programming languages, 10+ RL environments; pre-training cutoff June 2025, post-training cutoff February 2026. NVFP4 quantization-aware training from day one; BF16/FP8/NVFP4 weights published. (single secondary coverage) [SECONDARY] [S32]
- Full open release: weights on Hugging Face, 153 datasets and 15 RL environments; technical report at research.nvidia.com and arXiv 2604.12374. [COMMUNITY] [S35][S32]
- Deployment: minimum 8x H100-80GB self-host; NIM containers; 15+ cloud providers; free on OpenRouter and build.nvidia.com. Median hosted pricing $0.30/$0.80 per 1M input/output tokens; as low as $0.10/$0.50 on DeepInfra. [SECONDARY] [S31][S32]
- License: reported as the NVIDIA Nemotron Open Model License for Super [S32] — while Lightning/Cosmos 3/Alpamayo 2 Super are reported under OpenMDW-1.1 [S1][S13][S21]. Preserve both attributions; they may reflect NVIDIA's license renaming over 2026. [SECONDARY]

