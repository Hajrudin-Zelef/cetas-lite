---
id: ai-industry-kb-2026-wave6/23-master-model-inventory/part-11
title: "§23. Master Model Inventory (part 11)"
domain: master-model-inventory
role: deep-dive
task: reference
actors: ["Alibaba", "Anthropic", "DeepSeek", "Glasswing", "Google", "Meta", "Mistral", "OpenAI", "Poolside", "United States", "xAI"]
dates: ["2024-09-25", "2026-05-19", "2026-09-01", "2026-09-03", "2026-09-21"]
keywords: ["agent", "agentic", "apache", "astra", "bedrock", "benchmarks", "claude", "deepseek", "fable 5", "gemini", "gemini 3.8", "gpt-6"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [11523, 11546]
section: "§23. Master Model Inventory"
delta_of: ai-industry-kb-2026
sha256: 2d82ed490724a8c5a2be2f3d8726870bada610f8f1a18fe7b203766311e8dc24
---

# §23. Master Model Inventory (part 11)

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

