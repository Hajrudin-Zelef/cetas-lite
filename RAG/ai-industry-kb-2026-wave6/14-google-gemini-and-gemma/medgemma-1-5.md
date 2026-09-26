---
id: ai-industry-kb-2026-wave6/14-google-gemini-and-gemma/medgemma-1-5
title: "MedGemma 1.5"
domain: google-gemini-and-gemma
role: deep-dive
task: actor-profile
actors: ["Google", "OpenAI", "TSMC", "United States"]
dates: ["2025-04", "2025-05-20", "2026-01-13", "2026-05"]
keywords: ["3nm", "attention", "benchmarks", "chiplet", "cost", "fp8", "gemini", "hbm", "inference", "liquid cooling", "multimodal", "nvidia"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [7067, 7119]
section: "§14. Google: Gemini and Gemma"
delta_of: ai-industry-kb-2026
sha256: 9a055cfe1ca88cb1097a636637755c72b4bb0110b8a5bf0bd6b66e526aa3e1c5
---

# MedGemma 1.5

### MedGemma 1.5
- Released 2026-01-13 as part of the Health AI Developer Foundations program [SECONDARY](https://www.webpronews.com/medgemma-1-5-googles-open-ai-unlocks-3d-scans-and-clinical-speech-for-healthcare-builders/) [VENDOR](https://developers.google.com/health-ai-developer-foundations/medgemma/model-card)
- 4B-parameter multimodal medical model on a Gemma 3 backbone with a MedSigLIP image encoder [SECONDARY](https://www.webpronews.com/medgemma-1-5-googles-open-ai-unlocks-3d-scans-and-clinical-speech-for-healthcare-builders/)
- Open weights [SECONDARY](https://www.webpronews.com/medgemma-1-5-googles-open-ai-unlocks-3d-scans-and-clinical-speech-for-healthcare-builders/)
- Official model card: decoder-only Transformer with grouped-query attention [VENDOR](https://developers.google.com/health-ai-developer-foundations/medgemma/model-card)
- Text+vision input, text-only output [VENDOR](https://developers.google.com/health-ai-developer-foundations/medgemma/model-card)
- ≥128K input context [VENDOR](https://developers.google.com/health-ai-developer-foundations/medgemma/model-card)
- 8192-token max output [VENDOR](https://developers.google.com/health-ai-developer-foundations/medgemma/model-card)
- Images normalized to 896×896, encoded to 256 tokens each [VENDOR](https://developers.google.com/health-ai-developer-foundations/medgemma/model-card)
- Reported metrics: MedQA 69.1% (+5% vs MedGemma 1) [SECONDARY](https://www.webpronews.com/medgemma-1-5-googles-open-ai-unlocks-3d-scans-and-clinical-speech-for-healthcare-builders/)
- EHRQA 90% (+22%) [SECONDARY](https://www.webpronews.com/medgemma-1-5-googles-open-ai-unlocks-3d-scans-and-clinical-speech-for-healthcare-builders/)
- Chest ImaGenome IoU 38% (+35%) [SECONDARY](https://www.webpronews.com/medgemma-1-5-googles-open-ai-unlocks-3d-scans-and-clinical-speech-for-healthcare-builders/)
- CT classification 61% (+3%) [SECONDARY](https://www.webpronews.com/medgemma-1-5-googles-open-ai-unlocks-3d-scans-and-clinical-speech-for-healthcare-builders/)
- MRI 65% (+14%) [SECONDARY](https://www.webpronews.com/medgemma-1-5-googles-open-ai-unlocks-3d-scans-and-clinical-speech-for-healthcare-builders/)
- Histopathology ROUGE-L 0.49 [SECONDARY](https://www.webpronews.com/medgemma-1-5-googles-open-ai-unlocks-3d-scans-and-clinical-speech-for-healthcare-builders/)
- MIMIC-CXR RadGraph F1 30.3 (fine-tuned) [SECONDARY](https://www.webpronews.com/medgemma-1-5-googles-open-ai-unlocks-3d-scans-and-clinical-speech-for-healthcare-builders/)
- Technical report arXiv:2604.05081 [SECONDARY](https://www.webpronews.com/medgemma-1-5-googles-open-ai-unlocks-3d-scans-and-clinical-speech-for-healthcare-builders/)
- Paired with MedASR, a clinical speech-to-text model reported to cut transcription errors by up to 82% vs OpenAI Whisper large-v3 [SECONDARY](https://www.webpronews.com/medgemma-1-5-googles-open-ai-unlocks-3d-scans-and-clinical-speech-for-healthcare-builders/)

### AI Ultra subscription — pricing contradiction
- Launched at Google I/O 2025 (2025-05-20) at $249.99/mo [SECONDARY](https://www.hindustantimes.com/technology/google-ai-ultra-what-is-it-how-much-it-cost-and-everything-else-you-need-to-know-101747802454090.html) [SECONDARY](https://www.techspot.com/news/108006-google-launches-ai-ultra-250-reshaping-expectations-top.html)
- 50% off for the first 3 months (= $124.99 initially) [SECONDARY](https://www.hindustantimes.com/technology/google-ai-ultra-what-is-it-how-much-it-cost-and-everything-else-you-need-to-know-101747802454090.html)
- US-only at launch [SECONDARY](https://www.hindustantimes.com/technology/google-ai-ultra-what-is-it-how-much-it-cost-and-everything-else-you-need-to-know-101747802454090.html)
- Included 30TB storage [SECONDARY](https://www.hindustantimes.com/technology/google-ai-ultra-what-is-it-how-much-it-cost-and-everything-else-you-need-to-know-101747802454090.html)
- Included highest Gemini limits [SECONDARY](https://www.techspot.com/news/108006-google-launches-ai-ultra-250-reshaping-expectations-top.html)
- Included Deep Think [SECONDARY](https://www.hindustantimes.com/technology/google-ai-ultra-what-is-it-how-much-it-cost-and-everything-else-you-need-to-know-101747802454090.html)
- Included Veo 3 early access [SECONDARY](https://www.hindustantimes.com/technology/google-ai-ultra-what-is-it-how-much-it-cost-and-everything-else-you-need-to-know-101747802454090.html)
- Included Flow, Whisk, NotebookLM [SECONDARY](https://www.hindustantimes.com/technology/google-ai-ultra-what-is-it-how-much-it-cost-and-everything-else-you-need-to-know-101747802454090.html)
- CURRENT PRICE CONTRADICTION: as of May 2026 one source reports ~$100/mo in supported markets with 20TB storage (not 30TB) [UNVERIFIED](https://memeburn.com/google-ai-ultra-turns-gemini-into-a-premium-ai-subscription/)
- Another reports $124.99/mo — the two figures conflict [SECONDARY](https://www.nxcode.io/resources/news/gemini-3-1-pro-complete-guide-benchmarks-pricing-api-2026)
- $124.99 was originally the intro-offer price, not the list price [SECONDARY](https://www.hindustantimes.com/technology/google-ai-ultra-what-is-it-how-much-it-cost-and-everything-else-you-need-to-know-101747802454090.html)

### TPU Ironwood (v7)
- Seventh-generation TPU (TPU v7 / TPU7x), codename Ironwood [SECONDARY](https://www.techradar.com/pro/google-cloud-unveils-ironwood-its-7th-gen-tpu-to-help-boost-ai-performance-and-inference) [SECONDARY](https://www.sdxcentral.com/news/google-unveils-seventh-generation-tpu-ironwood/)
- Unveiled at Google Cloud Next '25 (April 2025) [SECONDARY](https://www.techradar.com/pro/google-cloud-unveils-ironwood-its-7th-gen-tpu-to-help-boost-ai-performance-and-inference)
- First TPU designed specifically for inference [SECONDARY](https://www.techradar.com/pro/google-cloud-unveils-ironwood-its-7th-gen-tpu-to-help-boost-ai-performance-and-inference)
- Per-chip: 4,614 TFLOPS peak (FP8) [SECONDARY](https://aiwiki.ai/wiki/tpu_ironwood)
- Dual-chiplet multi-die package [SECONDARY](https://aiwiki.ai/wiki/tpu_ironwood)
- 2 TensorCores + 4 SparseCores per chip [SECONDARY](https://aiwiki.ai/wiki/tpu_ironwood)
- MXU array 256×256 (logical 512×512 at FP8) [SECONDARY](https://aiwiki.ai/wiki/tpu_ironwood)
- 192GB HBM3E (8×24GB stacks) at ~7.37 TB/s [SECONDARY](https://aiwiki.ai/wiki/tpu_ironwood)
- ICI 1.2 TB/s bidirectional [SECONDARY](https://aiwiki.ai/wiki/tpu_ironwood)
- TSMC N3P (3nm class) per AI-wiki specs [SECONDARY](https://aiwiki.ai/wiki/tpu_ironwood)
- Third-gen liquid cooling [SECONDARY](https://aiwiki.ai/wiki/tpu_ironwood)
- Google Axion host CPU [SECONDARY](https://aiwiki.ai/wiki/tpu_ironwood)
- Pod scale: 9,216 liquid-cooled chips → 42.5 exaflops [SECONDARY](https://www.sdxcentral.com/news/google-unveils-seventh-generation-tpu-ironwood/) [SECONDARY](https://mlq.ai/news/google-prepares-next-gen-tpu-reveal-at-cloud-next-to-challenge-nvidia-in-ai-inference/)
- 6× HBM capacity vs Trillium (v6) [SECONDARY](https://www.sdxcentral.com/news/google-unveils-seventh-generation-tpu-ironwood/)
- 4.5× bandwidth vs Trillium (v6) [SECONDARY](https://www.sdxcentral.com/news/google-unveils-seventh-generation-tpu-ironwood/)
- ~2× performance per watt vs Trillium [SECONDARY](https://www.sdxcentral.com/news/google-unveils-seventh-generation-tpu-ironwood/)
- ~30× more power-efficient than the first Cloud TPU (2018) [SECONDARY](https://www.sdxcentral.com/news/google-unveils-seventh-generation-tpu-ironwood/)
- PRECISION CAVEAT: Ironwood's 42.5 exaflops are FP8; El Capitan's 1.7 exaflops are FP64 — the "24× a supercomputer" comparisons are not like-for-like precision [SECONDARY](https://www.sdxcentral.com/news/google-unveils-seventh-generation-tpu-ironwood/)
- AlphaChip (DeepMind RL floorplanning) used for the physical die layout, as on every generation since TPU v4 [SECONDARY](https://aiwiki.ai/wiki/tpu_ironwood)

