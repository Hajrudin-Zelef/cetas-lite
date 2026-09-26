---
id: ai-industry-kb-2026-wave6/14-google-gemini-and-gemma/gemma-4-full-spec-sheet
title: "Gemma 4 — full spec sheet"
domain: google-gemini-and-gemma
role: deep-dive
task: actor-profile
actors: ["Google", "Hugging Face"]
dates: ["2026-04-02", "2026-04-10", "2026-05-07"]
keywords: ["agentic", "apache", "attention", "compute", "consumer", "embedding", "embeddings", "gemini", "gpus", "inference", "leaderboard", "llama"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [7034, 7066]
section: "§14. Google: Gemini and Gemma"
delta_of: ai-industry-kb-2026
sha256: 9a16a89949f9b609b80f5feed9dd93621683b0c839894955ebd20d799083e74a
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

