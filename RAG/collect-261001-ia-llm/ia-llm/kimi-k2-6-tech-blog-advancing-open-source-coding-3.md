---
id: collect-261001-ia-llm/ia-llm/kimi-k2-6-tech-blog-advancing-open-source-coding-3
title: "kimi-k2-6-tech-blog-advancing-open-source-coding"
domain: ia-llm
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "benchmarks", "reasoning", "tool use"]
source: docs/RAG/collect-261001-ia-llm/kimi-k2-6-tech-blog-advancing-open-source-coding.md
source_anchor: ""
source_lines: [117, 125]
sha256: d8f1619191fde81436eb3d0cee79ad3eb0e8475cacb8c9f0409435afe1a4c42e
---

# kimi-k2-6-tech-blog-advancing-open-source-coding

- Terminal-Bench 2.0 scores were obtained with the default agent framework (Terminus-2) and the provided JSON parser, operating in preserve thinking mode.
- For the SWE-Bench series of evaluations (including Verified, Multilingual, and Pro), we used an in-house evaluation framework adapted from SWE-agent. This framework includes a minimal set of tools—bash tool, createfile tool, insert tool, view tool, strreplace tool, and submit tool.
- All reported scores for coding tasks are averaged over 10 independent runs.

**5. Vision Benchmarks**

- Max-tokens = 98,304, averaged over three runs (avg@3).
- Settings with Python tool use max-tokens-per-step = 65,536 and max-steps = 50 for multi-step reasoning.
- MMMU-Pro follows the official protocol, preserving input order and prepending images.
