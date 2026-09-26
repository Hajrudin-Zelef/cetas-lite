---
id: ai-industry-kb-2026/18-governance-regulation/eu-ai-act-enforcement-august-2026
title: "EU AI Act enforcement (August 2026)"
domain: governance-regulation
role: deep-dive
task: regulation
actors: ["DeepSeek", "EU", "Google", "Z.ai"]
dates: ["2026-05-25", "2026-07-17", "2026-08", "2026-08-14"]
keywords: ["agents", "alignment", "apache", "benchmarks", "cyber", "deepseek", "glm", "governance", "guardrails", "llama", "open-weight", "refusals"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [9101, 9112]
section: "18. Governance & Regulation"
sha256: a493bdf6d4b8346f6eba99e283a442fbdaab081bb8929e52157495772b8fc99e
---

# EU AI Act enforcement (August 2026)

- **Publication:** joint investigation by the **Financial Times** and AI safety research group **Alice** (CEO **Noam Schwartz**), published **May 25, 2026** (FT behind a paywall; summarized by Lexology, JDSupra, Cointelegraph, cxotoday).
- **The finding:** the free tool **Heretic**, hosted on GitHub, can strip **all** safety protections from open-weight AI models in **under ten minutes** using only a **standard laptop**, requiring little technical expertise. The mechanism is automated **abliteration** — machine-optimized removal of the refusal/alignment directions from the weight space.
- **Demonstrations:** an FT journalist removed Llama 3.3's safety alignment in under ten minutes; the modified model then answered prompts the original refused, including calculating lethal dosages of biological agents and generating functional malware. A modified Gemma 3 gave instructions for dispersing chemical agents in enclosed spaces, generated credit-card theft code, and produced child exploitation content. Heretic's creator, **Philipp Emanuel Weidmann**, told the FT he removed the safety guardrails from Google's newest model **Gemma 4 within 90 minutes of its public release**.
- **Scale:** the creator reports **over 3,500 modified model variants** and **13 million cumulative downloads**.
- **Policy implication (reported):** unlike proprietary models, open-weight systems can be downloaded, altered, and redistributed outside the original developer's control — post-release enforcement of safety constraints is structurally harder; regulators question whether development-focused regulation suffices.
- **AISI confirmation (July 17, 2026):** the AI Safety Institute's open-weight gap report found leading open-weight models trail the closed frontier by **4–7 months on cyber benchmarks** — and crucially, once weights are distributed they **cannot be recalled** by any regulatory action. DeepSeek V4-Pro's refusals on cyber tasks were overcome by simply retrying failed requests.
- **The regulatory paradox it creates:** the models easiest to govern at development (open, documented, auditable) are the hardest to govern after release; the models hardest to govern at development (closed, opaque) remain governable after release through API controls and licensing. 2026's enforcement record reflects this exactly — every successful action (BIS order, NDRC unwinding, deemed-export compliance) operated on a **closed or physical** artifact; no action touched distributed weights.
- **Z.ai's weight hold as the market answer:** GLM-5.3's two-week pre-release weight hold (announced August 14, 2026 — a first for the GLM line, because cyber capabilities grew faster than Z.ai expected) shows labs internalizing the non-recallability constraint voluntarily. Pre-release gating is becoming industry practice ahead of any mandate — and it is the only governance mechanism that works on weights at all.
- **Why this sits in §18:** the non-recallability finding is the structural foundation of both the EU's open-source exemption logic (regulate at development, exempt genuine openness) and the compliance advantage of MIT/Apache-2.0 releases documented above.

### EU AI Act enforcement (August 2026)

