---
id: ai-industry-kb-2026-wave6/14-google-gemini-and-gemma/gemma-4-full-spec-sheet
title: "Gemma 4 — full spec sheet"
domain: google-gemini-and-gemma
role: deep-dive
task: actor-profile
actors: ["Google", "Hugging Face", "OpenAI", "TSMC", "United States"]
dates: ["2025-04", "2025-05-20", "2026-01-13", "2026-04-02", "2026-04-10", "2026-05", "2026-05-07"]
keywords: ["3nm", "agentic", "apache", "attention", "benchmarks", "chiplet", "compute", "consumer", "cost", "embedding", "embeddings", "fp8"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [7034, 7119]
section: "§14. Google: Gemini and Gemma"
sha256: a8d2c777d6810c9b35315d2f23176f91b5886b23d1a8134462456384a034353c
---

# Gemma 4 — full spec sheet

### Gemma 4 — full spec sheet
- Released 2026-04-02 [SECONDARY](https://github.com/ajay-sainy/gemofgemma/blob/HEAD/Gemma4Research/00-overview.md) [SECONDARY](https://lilting.ch/en/articles/google-gemma-4-open-model-family)
- Open weights under Apache 2.0 (commercially permissive) [SECONDARY](https://github.com/ajay-sainy/gemofgemma/blob/HEAD/Gemma4Research/00-overview.md)
- Built with Gemini 3 technology [SECONDARY](https://lilting.ch/en/articles/google-gemma-4-open-model-family)
- Decoder-only Transformers [SECONDARY](https://github.com/ajay-sainy/gemofgemma/blob/HEAD/Gemma4Research/00-overview.md)
- E2B variant: 5.1B total / 2.3B effective parameters, 35 layers, 128K context, text/image/audio — edge/phones [SECONDARY](https://github.com/xp-py/llm-prep-2026/blob/HEAD/docs/Model_Zoo/Vision_Language_Models/Gemma_4.md) [SECONDARY](https://medium.com/@danushidk507/gemma-4-google-deepminds-most-capable-open-multimodal-models-3a7f3e47e764)
- E4B variant: 8.0B total / 4.5B effective, 42 layers, 128K, text/image/audio + ~300M audio encoder — laptops/edge [SECONDARY](https://github.com/xp-py/llm-prep-2026/blob/HEAD/docs/Model_Zoo/Vision_Language_Models/Gemma_4.md) [SECONDARY](https://medium.com/@danushidk507/gemma-4-google-deepminds-most-capable-open-multimodal-models-3a7f3e47e764)
- 26B A4B MoE: 25.2B total / 3.8B active, ~30 layers, 256K, text/image/video — consumer GPUs [SECONDARY](https://github.com/xp-py/llm-prep-2026/blob/HEAD/docs/Model_Zoo/Vision_Language_Models/Gemma_4.md) [SECONDARY](https://medium.com/@danushidk507/gemma-4-google-deepminds-most-capable-open-multimodal-models-3a7f3e47e764)
- 31B dense: 30.7B, 60 layers, 256K — workstations/servers [SECONDARY](https://github.com/xp-py/llm-prep-2026/blob/HEAD/docs/Model_Zoo/Vision_Language_Models/Gemma_4.md) [SECONDARY](https://medium.com/@danushidk507/gemma-4-google-deepminds-most-capable-open-multimodal-models-3a7f3e47e764)
- Naming: "E" = effective parameters — edge models with Per-Layer Embeddings (PLE), each transformer layer getting its own embedding input, shrinking compute footprint [SECONDARY](https://github.com/xp-py/llm-prep-2026/blob/HEAD/docs/Model_Zoo/Vision_Language_Models/Gemma_4.md)
- "A" = active parameters (MoE inference) [SECONDARY](https://github.com/xp-py/llm-prep-2026/blob/HEAD/docs/Model_Zoo/Vision_Language_Models/Gemma_4.md)
- Architecture: hybrid local/global attention [SECONDARY](https://github.com/xp-py/llm-prep-2026/blob/HEAD/docs/Model_Zoo/Vision_Language_Models/Gemma_4.md)
- MoE option [SECONDARY](https://github.com/xp-py/llm-prep-2026/blob/HEAD/docs/Model_Zoo/Vision_Language_Models/Gemma_4.md)
- Per-layer embeddings for small models [SECONDARY](https://github.com/xp-py/llm-prep-2026/blob/HEAD/docs/Model_Zoo/Vision_Language_Models/Gemma_4.md)
- Native `system` role [SECONDARY](https://medium.com/@danushidk507/gemma-4-google-deepminds-most-capable-open-multimodal-models-3a7f3e47e764)
- Native function calling [SECONDARY](https://medium.com/@danushidk507/gemma-4-google-deepminds-most-capable-open-multimodal-models-3a7f3e47e764)
- Structured JSON output [SECONDARY](https://medium.com/@danushidk507/gemma-4-google-deepminds-most-capable-open-multimodal-models-3a7f3e47e764)
- Configurable thinking mode [SECONDARY](https://medium.com/@danushidk507/gemma-4-google-deepminds-most-capable-open-multimodal-models-3a7f3e47e764)
- 262K vocabulary [SECONDARY](https://github.com/xp-py/llm-prep-2026/blob/HEAD/docs/Model_Zoo/Vision_Language_Models/Gemma_4.md)
- 140+ languages [SECONDARY](https://github.com/xp-py/llm-prep-2026/blob/HEAD/docs/Model_Zoo/Vision_Language_Models/Gemma_4.md)
- Video as frame sequences (up to 60s at 1 fps) [SECONDARY](https://medium.com/@danushidk507/gemma-4-google-deepminds-most-capable-open-multimodal-models-3a7f3e47e764)
- Hugging Face repos: google/gemma-4-31B-it (32.7B incl. embeddings) [SECONDARY](https://github.com/powerfulmoves/pmoves.ai/commit/8eb6480ee4b6148da8f005acb77abf379cba8bb0)
- google/gemma-4-E4B-it [SECONDARY](https://github.com/powerfulmoves/pmoves.ai/commit/8eb6480ee4b6148da8f005acb77abf379cba8bb0)
- google/gemma-4-E2B-it [SECONDARY](https://github.com/powerfulmoves/pmoves.ai/commit/8eb6480ee4b6148da8f005acb77abf379cba8bb0)
- google/gemma-4-26B-A4B-it (all repos dated 2026-04-10) [SECONDARY](https://github.com/powerfulmoves/pmoves.ai/commit/8eb6480ee4b6148da8f005acb77abf379cba8bb0)
- Q4 memory footprint: ~3GB (E2B) [SECONDARY](https://github.com/powerfulmoves/pmoves.ai/commit/8eb6480ee4b6148da8f005acb77abf379cba8bb0)
- ~5GB (E4B) [SECONDARY](https://github.com/powerfulmoves/pmoves.ai/commit/8eb6480ee4b6148da8f005acb77abf379cba8bb0)
- ~15GB (26B-A4B) [SECONDARY](https://github.com/powerfulmoves/pmoves.ai/commit/8eb6480ee4b6148da8f005acb77abf379cba8bb0)
- ~18GB (31B) [SECONDARY](https://github.com/powerfulmoves/pmoves.ai/commit/8eb6480ee4b6148da8f005acb77abf379cba8bb0)
- Reported LMArena: 26B-A4B 1441, 31B 1452 [SECONDARY](https://github.com/powerfulmoves/pmoves.ai/commit/8eb6480ee4b6148da8f005acb77abf379cba8bb0)
- Secondary claims 31B beats Llama 4 on AIME 2026 Math, LiveCodeBench v6, GPQA Diamond, τ2-bench Agentic — leaderboard/secondary, keep attributed [SECONDARY](https://github.com/powerfulmoves/pmoves.ai/commit/8eb6480ee4b6148da8f005acb77abf379cba8bb0)
- Gemma 4 MTP drafter (Multi-Token Prediction): one community test on M1 Max 64GB (2026-05-07) found only 26B-A4B gained (+13%) while 31B Dense and E4B got slower — inverted the official prediction; single-community, [COMMUNITY] only [COMMUNITY](https://lilting.ch/en/articles/google-gemma-4-open-model-family)

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

