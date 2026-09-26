---
id: ai-industry-kb-2026-wave6/06-longcat-and-meituan/new-verified-facts-expansion
title: "New verified facts — expansion"
domain: longcat-and-meituan
role: deep-dive
task: reference
actors: ["LongCat"]
dates: ["2025-09"]
keywords: ["compute", "context window", "moe", "parameters", "reasoning", "throughput"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [2787, 2796]
section: "§6. LongCat and Meituan"
delta_of: ai-industry-kb-2026
sha256: 37fbc768d20606a789a6d241714e6badac137ed5dcbbfebc8177ca5348706fbe
---

# New verified facts — expansion

### New verified facts — expansion

### LongCat-Flash (original) — technical spec from the technical report
- **560B total parameters**; dynamic active compute **18.6B–31.3B, average ~27B**; **128K context window** [SECONDARY] (huggingface.co/docs/transformers model_doc/longcat_flash; akihikowatanabe/paper_notes #4995).
- Architecture: **shortcut-connected MoE (ScMoE)**; **zero-computation experts** implemented as **identity/skip experts** — experts that can be skipped entirely, so the dynamic active range is load-dependent rather than fixed top-k [SECONDARY] (transformers model doc; paper notes).
- Reported throughput: **>100 tokens/second** [SECONDARY] (transformers model doc).
- Vendor MMLU: **89.71** [VENDOR] (transformers model doc reporting vendor figures).
- Technical report: arXiv **2509.01322** (September 2025) [SECONDARY] (arxiv.org/pdf/2509.01322).
- The report's evaluation harness covers: MMLU, MMLU-Pro, ArenaHard, CEval, CMMLU (general); IFEval, COLLIE, Meeseeks (instruction following); MATH500, AIME24, AIME25, BeyondAIME (math); GPQA-diamond, DROP, ZebraLogic, GraphWalks (reasoning); HumanEval+, MBPP+, LiveCodeBench (2024.08–2025.05), SWE-Bench-Verified, TerminalBench (coding) [SECONDARY] (arXiv report §4.4.1).

