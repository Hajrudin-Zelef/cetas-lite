---
id: collect-261001-ia-llm/ia-llm/etape4-trackb-local-inference-18
title: "Step 4 — Track B: Local Inference Stack (llama.cpp + Ollama + LM Studio)"
domain: ia-llm
role: reference
task: reference
actors: []
dates: ["2026-09"]
keywords: ["inference", "llama", "llama.cpp", "benchmark", "benchmarks", "gguf", "gpus", "pricing", "research"]
source: docs/RAG/collect-261001-ia-llm/etape4_trackB_local_inference.md
source_anchor: ""
source_lines: [987, 1010]
sha256: d8feffb1740a8700f17daa970b941037da875ca7a8ce925d41663d2a1ea2342a
---

# Step 4 — Track B: Local Inference Stack (llama.cpp + Ollama + LM Studio)

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
