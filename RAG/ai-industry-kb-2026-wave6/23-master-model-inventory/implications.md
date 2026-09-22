---
id: ai-industry-kb-2026-wave6/23-master-model-inventory/implications
title: "Implications"
domain: master-model-inventory
role: deep-dive
task: reference
actors: ["Alibaba", "Anthropic", "DeepSeek", "Falcon", "Glasswing", "Google", "Meta", "MiniMax", "Mistral", "Moonshot", "OpenAI", "Poolside", "United States", "Z.ai", "vLLM", "xAI"]
dates: ["2024-09-25", "2026-05-19", "2026-08-26", "2026-09-01", "2026-09-03", "2026-09-21"]
keywords: ["acquisition", "agent", "agentic", "agents", "apache", "astra", "bedrock", "benchmark", "benchmarks", "claude", "cost", "deepseek"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [11429, 11546]
section: "§23. Master Model Inventory"
sha256: 267dff06ab4ea52461a066df278fe1c2be6bb188ee00dcc497fe6efa33fabd7d
---

# Implications

## Implications

1. The inventory is a snapshot, not a leaderboard — params alone never identify a model (MiMo V2.5-Pro and V2.6-Pro share 1.02T/42B as distinct generations).
2. License and status columns are load-bearing: geo-exclusions (MiniMax Community) and non-commercial clauses (M2.7) change deployability more than param counts.
3. Contradicted entries (V4-Lite, R2, Grok 5, Gen-5, v4) persist in SEO noise — the superseded list above is the canonical filter.
4. Dated pricing and version-pinned benchmarks live in §22; licensing analysis in §21; per-model detail in §§2–17. This section condenses only.
5. Version-pinning rule from §22 applies to every row above: a model name without a date and checkpoint is not a retrievable fact.


### New verified implications — expansion

- The inventory's UNVERIFIED third is the work plan: licenses for Large 3, Kimi K2.6/K2.7, GLM-4.7, Falcon-H1, and prices for GPT-6 Astra, Grok 4.5/4.6, Mythos 5/5.1 are the highest-value next verifications [DIRECTIONAL].
- Dated snapshots beat merged scores: GPT-5.5's DeepSWE 70% (May) vs 67% (Sept) and Large 3's $2→$0.50 are the same lesson — every number needs a date [SECONDARY].
- The "open" label now spans MIT to geo-fenced: procurement must read the LICENSE file, not the launch post [SECONDARY].
- Retired-model entries (Medium 3, Small 3.1, Pixtral Large, V4 alias) are load-bearing: without them, stale prices and scores get attributed to current models [DIRECTIONAL].

## Sources and URLs

- https://github.com/meta-llama/llama-models/blob/main/models/llama4/USE_POLICY.md
- https://explainx.ai/blog/longcat-2-0-open-source-moe-coding-agent-2026
- https://groundtruth.day/news/longcat-2-trillion-param-open-model-domestic-chips.html
- https://dev.to/pueding/glm-52-becomes-the-top-open-weights-model-active-vs-total-parameters-2284
- https://aideveloper44.com/blog/zai-glm-5-2-open-source-moe-1m-context-window
- https://www.implicator.ai/alibaba-ships-qwen3-6-27b-an-open-weight-coding-model-that-beats-its-397b-moe/
- https://medium.com/@OpenCSG/trillion-parameter-model-goes-open-source-13361c16d134
- https://leanware.co/insights/kimi-k2
- https://dev.to/onirestart/kimi-k2-the-1-trillion-parameter-open-source-ai-thats-redefining-agentic-intelligence-jmk
- https://awesomeagents.ai/news/nvidia-nemotron-cascade-2-open-moe-30b/
- https://www.pymnts.com/news/artificial-intelligence/2026/nvidia-pays-6-billion-to-license-poolside-ai-model-development-software/
- https://forkast.news/nvidias-7-billion-poolside-deal-reveals-a-licensing-playbook-that-sidesteps-acquisition-scrutiny/
- https://the-decoder.com/nvidias-nemotron-4-aims-for-one-trillion-parameters-a-scale-chinese-labs-already-surpassed/
- https://kie.ai/blog/what-is-qwen-image-3-0
- https://gigazine.net/gsc_news/en/20250415-openai-gpt-4-1-released/
- https://mobilesyrup.com/2025/04/15/openai-gpt-41-one-million-context-window/


### New sources — expansion

- https://deepswe.datacurve.ai/
- https://github.com/jetbrains/ytdb-slate/blob/HEAD/research/gaps.md
- https://github.com/antigmalabs/ante/blob/HEAD/docs-site/docs/benchmarks/eval.mdx
- https://github.com/yyykf/cpa-usage-lens/blob/HEAD/.trellis/tasks/archive/2026-07/07-12-fix-codex-usage-pricing/research/openai-usage-pricing-contract.md
- https://www.morphllm.com/claude-code-pricing
- https://github.com/alt-f4-llc/dotfiles.vorpal/blob/HEAD/docs/facts/1785618019_claude_v5_models_cost_and_benchmarks.md
- https://medium.com/@automation.labs/opus-5-sonnet-5-haiku-4-5-which-claude-model-for-which-job-bce5e8346233
- https://github.com/aaronmarchant96-max/rei-ai/blob/HEAD/docs/CACHE_PRICING_LANDSCAPE.md
- https://venturebeat.com/technology/googles-gemini-3-7-flash-targets-coding-and-agents-with-a-50-introductory-price-cut
- https://devtk.ai/en/blog/gemini-api-pricing-guide-2026/
- https://github.com/caudena/beam_weaver/blob/HEAD/docs/partners/google.md
- https://www.aipricing.guru/xai-grok-pricing/
- https://github.com/xp-py/llm-prep-2026/blob/HEAD/docs/Model_Zoo/Vision_Language_Models/Gemma_3.md
- https://themunicheye.com/google-launches-gemma-3-open-source-ai-model-13085
- https://github.com/ufal/atrium-project/issues/9
- https://github.com/gunnarguy/openintelligence/blob/HEAD/Docs/Research/EMBEDDING_AND_INGESTION_UPGRADE_2026-08.md
- https://www.codesota.com/benchmarks/mteb
- https://kimi-k2.org/blog/56-kimi-k3-vs-deepseek-v4-flash-0731
- http://dev.to/best_codes/qwen-3-benchmarks-comparisons-model-specifications-and-more-4hoa
- https://huggingface.co/unsloth/Qwen3-Coder-30B-A3B-Instruct-1M-GGUF
- https://huggingface.co/unsloth/Qwen3-30B-A3B-GGUF
- https://huggingface.co/Qwen/Qwen3-30B-A3B-MLX-4bit
- https://github.com/QwenLM/Qwen3-Embedding
- https://ayinedjimi-consultants.fr/static/pdf/qwen3-7-max-thinking-alibaba-benchmark-2026.pdf
- https://github.com/kzinmr/ai-topics/blob/HEAD/wiki/comparisons/llm-api-pricing.md
- https://github.com/vLLM-X/ktransformers/blob/HEAD/doc/en/api/server.rst
- https://emergent.sh/learn/what-is-glm-5-3
- https://www.explainx.ai/blog/glm-5-3-flash-ox-alpha-official-launch-august-2026?ref=hackernoon.com
- https://github.com/prism-shadow/agenthub/blob/HEAD/changelog/0.4.8/2026-08-26-glm-5.3-flash-vision.md
- https://www.orcarouter.ai/blog/deepseek-v4-1-flash-vs-kimi-k3
- https://orphentisai.com/kimi-k3-vs-deepseek-v4-pro-vs-glm-5-2-open-trillion-scale-moe-models-compared-on-benchmarks-license-and-serving-cost/
- https://github.com/utensils/mold/blob/HEAD/docs/qualification/minimax-h3.md
- https://huggingface.co/OpenVDN/vdn-minimax-h3/blob/main/README.md
- https://github.com/makhmudovmurod/autodirector/blob/HEAD/Docs/license-notes.md
- https://digitalmatters.me/artificial-intelligence-ai/minimax-music-3/
- https://github.com/XUEXUE-XING/minimax-h3-webui/pull/12
- https://techcommunity.microsoft.com/blog/educatordeveloperblog/phi-4-small-language-models-that-pack-a-punch/4464167
- https://www.turing.com/blog/exploring-phi-4
- https://github.com/labscommunity/cascadia/blob/HEAD/docs/architectures/phi.md
- https://github.com/ualberta-rcg/aleph/blob/HEAD/models/phi-4-reasoning/CLAUDE.md
- https://medium.com/@mealermed/microsofts-phi-4-reasoning-vision-15b-the-ai-that-knows-when-to-think-and-when-to-just-answer-de6f2a83bb9b
- https://github.com/huggingface/blog/blob/HEAD/smollm3.md
- https://atomic.chat/models/smollm3-3b
- https://github.com/huggingface/smollm/
- https://www.aipricing.guru/mistral-ai-pricing/
- https://aiworldtoday.com/guides/mistral-ai-pricing
- https://devtk.ai/en/blog/mistral-api-pricing-guide-2026/
- https://www.secondtalent.com/resources/every-mistral-ai-model-explained-compared/
- https://curlscape.com/blog/mistral-api-pricing-2026
- https://www.smashingapps.com/mistral-ai-review/
- https://github.com/mikeschirtzinger/silent-notetaker/blob/HEAD/docs/research/model-licenses.md
- https://github.com/ather-techie/rag-interview-system/blob/HEAD/01_concepts/embeddings.md
- https://github.com/tardellirs/colibri-embed
- https://huggingface.co/jinaai/jina-reranker-v3
- https://github.com/2389-research/dippin-lang/commit/437a8a2889b92a4842582195f2a525836ad93645

- **Llama 3.2 1B** (2024-09-25): 1.23B actual params; 128K context; up to 9T tokens training; Dec 2023 cutoff; GQA; text-only multilingual; Llama 3.2 Community License; ~0.75GB VRAM at 4-bit [SECONDARY]. Sources: https://huggingface.co/meta-llama/Llama-3.2-3B and https://github.com/budecosystem/bud-connect/blob/HEAD/scripts/model_docs_workflow/raw_docs/bedrock/meta_llama_models.md
- **Llama 3.2 3B** (2024-09-25): 3.21B actual params; 128K context; up to 9T tokens; Dec 2023 cutoff; instruction-tuned for dialogue, agentic retrieval, summarization; Llama 3.2 Community License; ~1.75GB VRAM at 4-bit [SECONDARY]. Sources: https://huggingface.co/meta-llama/Llama-3.2-3B and https://hf.global-rail.com/meta-llama/Llama-3.2-1B/resolve/main/README.md?download=true
- **Llama 3.2 90B Vision**: multimodal 90B; 128K context; Llama 3.2 Community License [SECONDARY]. Source: https://github.com/budecosystem/bud-connect/blob/HEAD/scripts/model_docs_workflow/raw_docs/bedrock/meta_llama_models.md
- **Mistral Large 3** (see updated entry above): the 675B/41B MoE replaces the earlier "flagship" placeholder — architecture now specified [SECONDARY].
- **Ministral 3** (3B/8B/14B): dense with built-in vision encoders; 256K context per Red Hat (supersedes 128K for 3B in older pricing pages); Apache 2.0 [SECONDARY]. Source: https://docs.redhat.com/en/documentation/red_hat_ai_inference_server/3.3/pdf/inference_serving_mistral_3_models/Red_Hat_AI_Inference_Server-3.3-Inference_serving_Mistral_3_models-en-US.pdf

- **Claude Fable 5.1**: newer point release, 81.2% SWE-bench Pro (September 1 BenchLM update), $10/$50, 1.0M context — above Fable 5's 80.0% [SECONDARY]. Source: https://futuretweets.com/swe-bench-pro-leaderboard-2026/
- **Claude Mythos 5 / Mythos Preview**: 77.8% Pro (vendor); Glasswing partners only; pricing not published [SECONDARY]. Source: https://www.morphllm.com/swe-bench-pro
- **Qwen3.8 Max**: 2.4T; 67.7% Pro; Closed API (Alibaba) — NOT open-weight unlike Qwen3.7 Max; output $4.95/M (September) [SECONDARY]. Sources: https://www.morphllm.com/swe-bench-pro and https://futuretweets.com/swe-bench-pro-leaderboard-2026/
- **Qwen3.8 Flash / Qwen3.8-Flash-Next**: 62.5% Pro; output $0.47/M [SECONDARY]. Source: https://www.morphllm.com/swe-bench-pro
- **Qwen3.8-27B**: 61.7% Pro; output $3.00/M [SECONDARY]. Source: https://www.morphllm.com/swe-bench-pro
- **Tencent Hy4 preview**: 65.7% Pro; pricing n/a [SECONDARY]. Source: https://www.morphllm.com/swe-bench-pro
- **Tencent Hy3**: 57.9% Pro; output $0.58/M [SECONDARY]. Source: https://www.morphllm.com/swe-bench-pro
- **Meta Muse Spark 1.1**: 61.5% Pro (vendor) AND leads Scale SEAL standardized at 61.5%; output $4.25/M [SECONDARY]. Sources: https://www.morphllm.com/swe-bench-pro and https://localaimaster.com/models/swe-bench-explained-ai-benchmarks
- **Poolside Laguna S 2.1**: 59.4% Pro; output $0.20/M — the Model Factory license line surfacing as product [SECONDARY]. Source: https://www.morphllm.com/swe-bench-pro
- **DeepSeek V4 Pro 0813**: 1.6T/49B MoE; 1M context, 384K max output; 500-request concurrency cap; vendor-reported 80.6% SWE-bench Verified [VENDOR]. Source: https://github.com/full-stack-assets/wireandlogic/blob/HEAD/content/posts/deepseek-v4-pro-0813-release.mdx

- **GPT-6 Astra** (2026-09-03): OpenAI's new flagship; entered TB 4.0 on launch day at 58.18% (Codex, max effort) — top score; $10/$50 short-context, $20/$75 long-context; two lower effort settings tied Fable 5.1 at 57.88% [SECONDARY]. Sources: https://capitalandcompute.net/ai-benchmarks/terminal-bench/ and https://github.com/parhumm/product-excellence/blob/HEAD/docs/researches/api-pricing-claude-openai.md
- **Grok 4.7** (2026-09-21): xAI; TB 4.0 rank 17 at 37.58% on launch day (Grok Build harness, flagged partial 324/330 trials) [SECONDARY]. Source: https://capitalandcompute.net/ai-benchmarks/terminal-bench/
- **Gemini 3.8 Flash**: TB 4.0 19.1% ±3.4 (mini-SWE-agent, Sep 2) [SECONDARY]. Source: https://codingfleet.com/blog/terminal-bench-leaderboard-2026/
- **Claude Fable 5.1** (2026-09-01): $10/$50; 1M context / 128K out; thinking always on; TB 4.0 57.9% ±3.8 official (#1), Anthropic's own TB 4.0 figure 55.8%; SWE-bench Pro 81.2% (BenchLM Sep 1) [SECONDARY]. Sources: https://codingfleet.com/blog/terminal-bench-leaderboard-2026/ and https://terranettechnologies.com/blog/claude-fable-opus-sonnet-haiku-explained and https://futuretweets.com/swe-bench-pro-leaderboard-2026/
- **Claude Mythos 5.1**: Anthropic's own TB 4.0 figure 60.9% — above Fable 5.1's 55.8% on the vendor scaffold [VENDOR]. Source: https://codingfleet.com/blog/terminal-bench-leaderboard-2026/
- **Gemini 3.5 Flash** (2026-05-19, Google I/O): TB 2.1 76.2%, MCP Atlas 83.6%, GDPval-AA 1656 Elo, CharXiv Reasoning 84.2% — beats Gemini 3.1 Pro on coding/agentic at Flash latency [VENDOR]. Source: https://blog.imseankim.com/gemini-3-5-flash-vs-3-1-pro-terminal-bench-mcp-atlas-agentic-coding-may-2026/

