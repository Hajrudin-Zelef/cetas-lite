---
id: etape4-trackb-local-inference/00-local-inference/master-open-verification-log-track-b-carried-forward
title: "Master open-verification log (track B, carried forward)"
domain: step-4-track-b-local-inference-stack-llama-cpp-ollama-lm-stu
role: deep-dive
task: reference
actors: ["Alibaba", "DeepSeek", "Google", "Hugging Face", "Moonshot", "Nvidia", "Z.ai", "vLLM"]
dates: ["2026-09", "2026-09-22"]
keywords: ["benchmark", "benchmarks", "deepseek", "gguf", "glm", "gpus", "inference", "kimi", "llama", "llama.cpp", "memory", "nvidia"]
source: docs/RAG/etape4_trackB_local_inference.md
source_anchor: ""
source_lines: [970, 1010]
section: "Step 4 — Track B: Local Inference Stack (llama.cpp + Ollama + LM Studio)"
sha256: 4bfe04c722c5b81d4cbfcd1f6f258b143f99b29eb72344e2eaf64e235749a661
---

# Master open-verification log (track B, carried forward)

- The in-app catalog is the product's signature flow: search → Hugging Face pull → memory-fit estimate → download → chat; supports GGUF and MLX quants; no terminal required. [secondary] (https://medium.com/@nishilbhave/lm-studio-in-2026-download-models-run-local-llms-vs-ollama-b30567df17f8)
- 2026 catalog landmarks trace the open-model frontier: `openai/gpt-oss-20b` (120B sibling runs on 128 GB unified machines), Kimi K2.7 Code / K3 (Moonshot), GLM 5.2/5.3-Flash (Z.ai), Qwen3.5/3.6/3.8 (Alibaba), Gemma 4 (Google), DeepSeek V4 Pro (cloud), NVIDIA Nemotron-Nano-v2. [official/secondary]
- LM Link free tier: up to 5 devices; paid/enterprise details unpublished at cutoff. [secondary] (https://learning.christiandrapatz.de/lmstudio-en.pdf)
- Community position vs Ollama (2026): LM Studio = GUI-first, Ollama = CLI-first; both llama.cpp-based so raw tok/s nearly identical; Ollama added paid cloud tiers in 2026 (Pro $20/mo, Max $100/mo) while local stays free; LM Studio remains free locally with paid Secure Cloud credits. [secondary] (https://www.kunalganglani.com/blog/lm-studio-vs-ollama)
- vLLM vs llama.cpp at scale (Red Hat 2026 benchmark, cited in independent coverage): at 64 concurrent users vLLM generated ~44× more tokens/s than llama.cpp — i.e., llama.cpp/LM Studio is the single-user king, vLLM the serving king. [secondary] (https://medium.com/@nishilbhave/local-llms-in-2026-which-runtime-to-run-and-the-hardware-you-need-a88450dece2e)
- Menlo Ventures (2025, cited 2026): open-source models hold ~11% of enterprise LLM usage, down from 19% — local inference remains a niche vs hosted APIs, which frames LM Studio's audience. [secondary] (same)

---

*End of report. Research cutoff: September 22, 2026.*


---


## Master open-verification log (track B, carried forward)

1. **Ollama "Turbo" branding** — could not be confirmed; flagged [unverified] (Ollama draft §12).
2. **Ollama Team pricing discrepancy** — $500/mo ($1,000 shared, Sep 9, 2026) vs $25/seat/mo 5-seat minimum (Aug 28, 2026); both kept, flagged; needs official-page verification.
3. **Ollama per-token billing transition date** — not pinned to an official announcement.
4. **Ollama v0.31.x** — versions unaccounted for (conflicting third-party date); minor timeline gap.
5. **Ollama Enterprise/SSO/MDM** — listed "coming soon" (Aug 2026), not shipped; Enterprise custom plan exists with no published price.
6. **LM Studio user/download figures** — no vendor-published number; "millions of downloads" is [secondary]-only.
7. **LM Studio Secure Cloud / Enterprise pricing** — unpublished.
8. **LM Studio DFlash/DSpark drafter details** — [unverified].
9. **IPEX-LLM archival status** — [unverified].
10. **Snapdragon X2 (2026)** — unconfirmed details.
11. **DGX Spark speedup claims** — vendor-adjacent claims not independently verified; thermal issues community-reported.
12. **AI BOX pricing** — [unverified].
13. **llama.cpp Feb–Jul 2026 b-tag history** — only sampled, not exhaustive.
14. **Tom's Hardware DGX Spark chart numbers** — not extractable from accessible sources.
15. **ServeTheHome 2026 llama.cpp benchmark article** — no dedicated article found; coverage is forum-based.
16. **GGUF container version-bump absence (2026)** — inference from unchanged server README cache-type list, not a positive statement.
17. **KleidiAI macOS build DISABLED** — reason unknown.
18. **GGUF quant community benchmark numbers** — Artefact2/ikawrakow and gguf-switchboard PPL tables are [secondary] community measurements; do not mix with vendor-reported figures.
19. **Street prices for GPUs** — dated September 2026 snapshots, move weekly.

## Collection metadata
- Sources used: ggml-org/llama.cpp GitHub releases, ollama/ollama GitHub releases, lmstudio.ai changelog, Ollama docs, ServeTheHome, Tom's Hardware, community benchmarks (Artefact2/ikawrakow, gguf-switchboard), LM Studio SDK docs.
- Method: web research, read-only; no live-browser visits; no external sends.
- Working drafts retained: etape4_draft_llamacpp.md, etape4_draft_ollama.md, etape4_draft_lmstudio_hw.md.
