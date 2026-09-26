---
id: ai-industry-kb-2026-wave6/23-master-model-inventory/xiaomi-meituan
title: "Xiaomi / Meituan"
domain: master-model-inventory
role: deep-dive
task: reference
actors: ["DeepSeek", "EU", "Hugging Face", "LongCat", "Meituan", "Meta", "Microsoft", "MiniMax", "Mistral", "OpenAI", "United States", "Xiaomi", "vLLM"]
dates: ["2025-02-26", "2025-04-30", "2025-07-08", "2025-09", "2026-01", "2026-06-01", "2026-07-05", "2026-08-02", "2026-08-06", "2026-08-31", "2026-09-08", "2026-09-21", "2026-11-30"]
keywords: ["agentic", "apache", "attribution", "benchmarks", "claude", "deepseek", "inference", "license", "llama", "mistral", "moe", "multimodal"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [11055, 11095]
section: "§23. Master Model Inventory"
delta_of: ai-industry-kb-2026
sha256: 60c3a566228b5be4fd2a953667ee67b97a03c2e2406ac4c23a2268bcee0234ef
---

# Xiaomi / Meituan

- **MiniMax M3** (2026-06-01): Community License, geo-restricted local deployment (US/EU/UK/SK excluded); Terminal-Bench 2.1 (Ante): 62.1%±2.33 [SECONDARY]. Sources: https://github.com/utensils/mold/blob/HEAD/docs/qualification/minimax-h3.md and https://github.com/antigmalabs/ante/blob/HEAD/docs-site/docs/benchmarks/eval.mdx
- **MiniMax H3** (license effective 2026-08-02): H3 Community License — EU/UK/SK/US excluded, $20M/yr revenue trigger, "MiniMax H3" UI attribution, Exhibit A 20 restricted uses, no training other models on outputs [SECONDARY]. Sources: https://huggingface.co/OpenVDN/vdn-minimax-h3/blob/main/README.md and https://github.com/makhmudovmurod/autodirector/blob/HEAD/Docs/license-notes.md
- **MiniMax Music 3** (license effective 2026-08-06): community license with NO territory exclusion; $20M revenue trigger kept [SECONDARY]. Source: https://digitalmatters.me/artificial-intelligence-ai/minimax-music-3/
- **MiniMax M2.7**: named in corpus; license UNVERIFIED [UNVERIFIED].

### Xiaomi / Meituan

- **MiMo V2.5** (2026): MIT; Terminal-Bench 2.1 (Ante): 65.8%±2.30 [SECONDARY]. Sources: wave6/03-model-weights-wave3.md and https://github.com/antigmalabs/ante/blob/HEAD/docs-site/docs/benchmarks/eval.mdx
- **MiMo V2.5-Pro**: 1.02T/42B; MIT line [SECONDARY]. Source: wave6/03-model-weights-wave3.md
- **MiMo V2.6-Pro** (2026-09-21/22): distinct release, same 1.02T/42B profile; MIT line [SECONDARY]. Source: wave6/03-model-weights-wave3.md
- **LongCat-2.0** (2026-07-05): MIT [SECONDARY]. Source: https://github.com/XUEXUE-XING/minimax-h3-webui/pull/12
- **LongCat Flash Thinking**: September 2025 release; January 2026 was the 2601 refresh — do not date the model to 2026 [SECONDARY]. Source: wave6/03-model-weights-wave3.md

### Meta / Microsoft / HuggingFace

- **Llama 3.2 11B**: 128K context, Llama 3 Community License (EU/multimodal acceptance clause per §21); MMLU 80.1/HumanEval 72.3/MATH 68.9 (2025 compilation) [SECONDARY]. Source: https://ayinedjimi-consultants.fr/static/pdf/small-language-models-2026.pdf
- **Meta Muse Glimmer** (2026): Apache 2.0, no revenue trigger, no attribution — released ~3 days before MiniMax Music 3 [SECONDARY]. Source: https://digitalmatters.me/artificial-intelligence-ai/minimax-music-3/
- **Phi-4** (14B, 2024-12): MIT; 16K context; trained 3 weeks on 9.8T tokens; MMLU 84.8/HumanEval 82.6/MATH 80.4; outperformed GPT-4 on some STEM tasks [SECONDARY]. Sources: https://techcommunity.microsoft.com/blog/educatordeveloperblog/phi-4-small-language-models-that-pack-a-punch/4464167 and https://www.turing.com/blog/exploring-phi-4
- **Phi-4-mini** (3.8B, 2025-02-26): MIT; 128K context; code/math/visual reasoning; vocab 200064 [SECONDARY]. Sources: https://www.turing.com/blog/exploring-phi-4 and https://github.com/labscommunity/cascadia/blob/HEAD/docs/architectures/phi.md
- **Phi-4-multimodal** (5.6B, 2025-02-26): MIT; speech/vision/text in one model [SECONDARY]. Source: https://www.turing.com/blog/exploring-phi-4
- **Phi-4-reasoning / reasoning-plus** (14B, 2025-04-30): MIT; reasoning traces distilled from o3-mini; production on L40S via vLLM [SECONDARY]. Sources: https://www.turing.com/blog/exploring-phi-4 and https://github.com/ualberta-rcg/aleph/blob/HEAD/models/phi-4-reasoning/CLAUDE.md
- **Phi-4-reasoning-vision-15B** (2026-03): MIT; adaptive thinking ("knows when to think") [SECONDARY]. Source: https://medium.com/@mealermed/microsofts-phi-4-reasoning-vision-15b-the-ai-that-knows-when-to-think-and-when-to-just-answer-de6f2a83bb9b
- **SmolLM3-3B** (2025-07-08): Apache 2.0; 16.5B params? No — 3B; 11.2T tokens, 6 languages native; 128K (YaRN), 65,536 shipped; think/no_think; AIME 2025 36.7/9.3 [SECONDARY]. Sources: https://github.com/huggingface/blog/blob/HEAD/smollm3.md and https://atomic.chat/models/smollm3-3b
- **SmolVLM**: compact multimodal companion (visual QA, multi-image); on-device [SECONDARY]. Source: https://github.com/huggingface/smollm/

### Mistral

- **Mistral Large 3** (2025-12): 675B total / 41B active sparse MoE (DeepSeek-V3-style, top-4 experts, softmax routing); Apache 2.0 (verified — UNVERIFIED flag lifted); context contradiction: 128K (older secondary) vs 256k (2026-09-08 catalog check) UNRESOLVED; $0.50/$1.50 current (was $2/$6); LMArena open-source non-reasoning #2 at debut [SECONDARY]. Sources: https://airevolution.poltextlab.com/mistral-launches-its-new-10-model-release/ and https://github.com/samouraiworld/awesome-mistral/blob/HEAD/README.md and https://docs.redhat.com/en/documentation/red_hat_ai_inference/3.5/pdf/inference_serving_mistral_3_models/Red_Hat_AI_Inference-3.5-Inference_serving_Mistral_3_models-en-US.pdf?utm_source=openai
- **Mistral Medium 3.5**: 128B dense; Modified MIT (not Apache); context contradiction 131K vs 256k UNRESOLVED; $1.50/$7.50, $0.15 cached; API-only [SECONDARY]. Source: https://github.com/samouraiworld/awesome-mistral/blob/HEAD/README.md
- **Mistral Medium 3**: legacy $0.40/$2.00; retired 2026-08-31 [SECONDARY]. Source: https://devtk.ai/en/blog/mistral-api-pricing-guide-2026/
- **Mistral Small 4**: 119B MoE / 6.5B active (not dense); Apache 2.0; context contradiction 131K vs 256k UNRESOLVED; $0.15/$0.60, $0.015 cached; multimodal [SECONDARY]. Sources: https://github.com/samouraiworld/awesome-mistral/blob/HEAD/README.md and https://devtk.ai/en/blog/mistral-api-pricing-guide-2026/
- **Mistral Small 3.1**: retired 2026-11-30 — use Small 4 [SECONDARY]. Source: https://devtk.ai/en/blog/mistral-api-pricing-guide-2026/
- **Mistral Nemo**: Apache 2.0 open lightweight; price contradiction $0.02/$0.10 vs $0.15/$0.15 UNRESOLVED [SECONDARY — contradiction]. Sources: https://www.secondtalent.com/resources/every-mistral-ai-model-explained-compared/ versus https://curlscape.com/blog/mistral-api-pricing-2026
- **Ministral 3** (3B/8B/14B): edge/on-device; 3B $0.10/$0.10 (128K), 8B $0.15/$0.15 (256K), 14B $0.20/$0.20 (256K) [SECONDARY]. Source: https://curlscape.com/blog/mistral-api-pricing-2026
- **Codestral**: 256K, $0.30/$0.90, code completion specialist; API-only [SECONDARY]. Source: https://curlscape.com/blog/mistral-api-pricing-2026
- **Devstral 2** ($0.40/$2.00, 256K) / **Devstral Small 2** ($0.10/$0.30, 256K): agentic coding; API-only [SECONDARY]. Source: https://curlscape.com/blog/mistral-api-pricing-2026
- **Magistral Medium** ($2.00/$5.00, 128K) / **Magistral Small** (~$0.10/$0.30, 128K): reasoning; API-only [SECONDARY]. Source: https://curlscape.com/blog/mistral-api-pricing-2026
- **Pixtral Large**: retired as standalone — vision folded into Large 3/Medium 3.5/Small 4 [SECONDARY]. Source: https://curlscape.com/blog/mistral-api-pricing-2026
- **Voxtral** (Mini Transcribe 2 $0.003/min; Small $0.004/min + $0.10/$0.40 text; TTS $16/M chars): speech line [SECONDARY]. Source: https://curlscape.com/blog/mistral-api-pricing-2026
- **Mistral Embed** ($0.10/M) / **Codestral Embed** ($0.15/M) / **Mistral Moderation** ($0.10/M): utility APIs [SECONDARY]. Source: https://curlscape.com/blog/mistral-api-pricing-2026

