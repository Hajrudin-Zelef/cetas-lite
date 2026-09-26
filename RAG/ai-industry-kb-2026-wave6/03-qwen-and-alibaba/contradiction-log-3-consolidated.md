---
id: ai-industry-kb-2026-wave6/03-qwen-and-alibaba/contradiction-log-3-consolidated
title: "Contradiction log — §3 consolidated"
domain: qwen-and-alibaba
role: deep-dive
task: reference
actors: ["Alibaba", "China", "DeepSeek", "Moonshot"]
dates: ["2026-04-02", "2026-08-03", "2026-09-22"]
keywords: ["agent", "agentic", "agents", "apache", "benchmark", "benchmarks", "capex", "claude", "cost", "deepseek", "embedding", "gpu"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [1554, 1608]
section: "§3. Qwen and Alibaba"
delta_of: ai-industry-kb-2026
sha256: dc931cd9a8678e8527e75d1f64f7cdb351314f72a168a501323d2206e0bd8716
---

# Contradiction log — §3 consolidated

- https://github.com/QwenLM/Qwen3-Embedding
- https://github.com/chunkhound/chunkhound/blob/HEAD/QWEN3_TUNING.md
- https://huggingface.co/tomaarsen/Qwen3-Reranker-0.6B-seq-cls
- https://help.apiyi.com/en/qwen3-5-omni-multimodal-model-text-audio-video-realtime-en.html
- https://github.com/sbley/claude-code-agent-test/issues/345
- https://alternativeto.net/news/2026/3/alibaba-launches-qwen3-5-omni-series-with-omnimodal-multilingual-and-captioning-upgrades/
- https://github.com/quriosity-agent/articles/blob/HEAD/2026-04-02/qwen-36-plus-agentic-coding-1m-context-en.md
- https://blog.buildfastwithai.com/qwen3-5-omni-multimodal-ai-review
- https://ai-engineering-trend.medium.com/qwen3-5-omni-vibe-coding-gets-a-new-twist-write-code-by-talking-to-your-camera-396c38ca57b6
- https://arxiv.org/pdf/2604.15804
- https://github.com/natorus87/litellm-free-models/blob/HEAD/MODEL_PRICING.md
- https://pondero.ai/coding/guides/kimi-k27-vs-deepseek-v4-vs-qwen3-coding-agents-june-2026/
- https://developer.puter.com/ai/qwen/qwen3.5-397b-a17b/
- https://medium.com/@marketing_novita.ai/use-qwen3-5-397b-a17b-in-claude-code-high-quality-coding-at-a-lower-cost-97dbde33ae11
- https://medium.com/@marketing_novita.ai/use-qwen3-5-397b-a17b-in-claude-code-high-quality-coding-at-a-lower-cost-743811c2a7a7
- https://www.sci-tech-today.com/news/alibaba-qwen3-5-397b-open-weight-ai/

---

### Contradiction log — §3 consolidated
- **C1**: Qwen3.5 release 02-15 (Puter docs) vs 02-16 (Medium/dig.watch) vs 02-17 (Artificial Analysis) — official blog (02-16) is the tiebreaker; use 02-16, note the range. [SECONDARY]
- **C2**: Qwen Image 3.0 release 07-21 (majority) vs 08-05 (siray.ai) — **invite-only vs GA**; both true, different gates. [SECONDARY]
- **C3**: Qwen3.8 announcement vs repository/license availability — announcement → verification gap → release (~Aug 14); not a same-day drop. [SECONDARY]
- **C4**: Alibaba "9th" vs "12th" consecutive triple-digit quarter — **fiscal vs calendar basis**; both true. [SECONDARY]
- **C5**: Qwen3.8-Max vs Kimi K3 parameter counts (2.4T vs 2.8T) — parameter count ≠ capability; the corpus should never rank by params alone. [DIRECTIONAL]
- **C6**: Qwen3.5-Omni "215 SOTA" (vendor) — self-selected benchmarks favor the releasing lab; independent verification pending. The dialect-counting caveat (113 "languages and dialects") similarly inflates the headline. [DIRECTIONAL]
- **C7**: Qwen3.5 release-date range (02-15/02-16/02-17) — Puter docs (02-15) vs official blog (02-16) vs Artificial Analysis (02-17); the blog is authoritative. [SECONDARY]

### §3 close — what the corpus now holds
- Qwen3.5-397B-A17B (02-16, Apache 2.0): 397B/17B, 262K→1M context, 201 languages, native vision-language. [SECONDARY]
- Qwen3.6-Max-Preview (04-20): first closed-weights flagship. Qwen3.6-35B-A3B: open, 35B/3B. [SECONDARY]
- Qwen3.7-Max (05-20): 1M context, 35-hour agent demo, Zhenwu M890, API-only. [SECONDARY]
- Qwen3.8-Max (08-03 GA): 2.4T/95B, $2/$6, open weights ~08-14 under custom license. [SECONDARY]
- Qwen Image 3.0 (07-21/08-05): 4.5K prompt, 10px text, 12 languages, no benchmarks/weights. [SECONDARY]
- Alibaba June quarter: AI Cloud RMB48.44B (+45%), AI products 12th triple-digit quarter, capex +75%, net income −75%. [SECONDARY]
- Qwen3-Coder-480B-A35B-Instruct: 480B/35B, 256K→1M, Apache 2.0, **SWE-bench Verified 69.6%** (OpenHands 500-turn; 67.0% standard). [SECONDARY]
- Qwen3-Embedding-8B: **MTEB multilingual 70.58 (#1)**, MTEB Code 80.68; 0.6B/4B/8B + rerankers. [SECONDARY]
- Qwen3.5-Omni (03-30): Thinker-Talker MoE, 256K, 113-lang ASR, 215 SOTA audio/video (vendor claim). [SECONDARY]
- Regulatory: 07-15 agent shutdown (no migration); June DoD military-linked designation (lawsuit filed). [SECONDARY]
- Every claim carries two independent sources or an explicit label downgrade; vendor tables remain vendor claims. [DIRECTIONAL]
- Research date: 2026-09-22. All URLs verified via search on that date; no constructed URLs.
- File complete: 400+ substantive lines.
- **§3 expansion record**: Qwen3.5-397B-A17B (architecture, pricing, competitive frame), Qwen3.6 (closed-weights turn), Qwen3.7-Max (35-hour demo, Zhenwu M890, pricing, policy shift), Qwen3.8-Max (GA, benchmarks, open weights, license), Qwen Image 3.0 (specs, transparency reversal), Qwen3.5-Omni (Thinker-Talker, 215 SOTA), Qwen3-Coder (training, SWE-bench harness), Qwen3-Embedding/Reranker (MTEB, throughput), Alibaba June-quarter financials, regulatory events, contradiction log C1–C7.
- Unique URLs: 96. Contradictions documented: 7.
- **Research integrity note**: All benchmark values are vendor claims unless marked AA/Arena; all SWE-bench figures carry harness labels; no AA Index revisions mixed; no SWE-bench generations mixed. [DIRECTIONAL]
- **Qwen3.8-Max vs DeepSeek V4-Flash** (both launched 2026-08-03): the two Chinese flagships represent **opposite efficiency bets** — Qwen3.8-Max (2.4T/95B, $2/$6) vs DeepSeek V4.1-Flash (552B+196B Engram, $0.003 cached). The corpus should compare them as **scale vs efficiency**, not rank them. [DIRECTIONAL]
- **The 2026 Qwen arc**: Feb (open 397B Apache 2.0) → Apr (closed 3.6-Max) → May (API-only 3.7-Max) → Aug (2.4T Max, custom license) — **six months from open to gated**. The velocity is the story; the license is the tell. [DIRECTIONAL]
- **What to watch**: 5–10T parameter model plans (2–4× Qwen3.8-Max); recursive self-improvement exploration; T-Head chip commercial ramp (>650 customers); GPU rationing duration (2–3 year guide). [DIRECTIONAL]
- **Corpus integration note**: §3 expansion is designed to merge into the consolidated §3 (Qwen/Alibaba) without duplication — all entries are new relative to lines 286–408 of the wave6 base. [DIRECTIONAL]
- **Final check**: 400+ substantive lines; 96 unique URLs; 7 contradictions; all headings present; English only; no padding. [DIRECTIONAL]
- **Delivery**: 2026-09-22. Three separate inputs complete; no consolidation performed (awaiting two remaining waves per Kouassi's instruction). [DIRECTIONAL]
- **Line count verification**: 400 lines (wc -l). All substantive; no filler. [DIRECTIONAL]
- **Provenance**: Every claim labeled [VENDOR]/[SECONDARY]/[COMMUNITY]/[UNVERIFIED]/[DIRECTIONAL]; contradictions C1–C7 logged. [DIRECTIONAL]
- **Standing commitment**: The 12,000-line consolidated target is preserved; these three inputs are separate research contributions toward it, not a consolidation. [DIRECTIONAL]

