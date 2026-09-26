---
id: ai-industry-kb-2026-wave6/03-qwen-and-alibaba/new-verified-timeline-entries-expansion
title: "New verified timeline entries — expansion"
domain: qwen-and-alibaba
role: deep-dive
task: actor-profile
actors: ["Alibaba", "California", "China", "DeepSeek", "Google", "SGLang", "United States"]
dates: ["2026-02-03", "2026-02-15", "2026-02-16", "2026-02-24", "2026-03-30", "2026-03-31", "2026-04-02", "2026-04-16", "2026-04-20", "2026-04-22", "2026-05-18", "2026-05-20", "2026-06-01", "2026-07", "2026-07-15", "2026-07-21", "2026-08-03", "2026-08-05", "2026-08-14", "2026-08-20", "2026-08-26", "2026-09-02", "2026-09-17"]
keywords: ["agent", "agentic", "agents", "apache", "benchmark", "benchmarks", "capex", "cost", "deepseek", "distribution", "gpu", "license"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [1367, 1421]
section: "§3. Qwen and Alibaba"
delta_of: ai-industry-kb-2026
sha256: 28a97d9518afcca28004da7178df4e049955b350be873c47cc2089cc8841d1c8
---

# New verified timeline entries — expansion

- **2026-02-03/04** — Qwen3-Coder-Next (80B/~3B, Apache 2.0, Gated DeltaNet, 262K, 800K+ coding tasks). [VENDOR]
- **2026-02-16** — Qwen3.5-397B-A17B. [SECONDARY]
- **2026-02-24/25** — Medium open Qwen3.5 models + hosted Qwen3.5-Flash. [SECONDARY]
- **2026-03-30** — Qwen3.5-Omni. [SECONDARY]
- **2026-04-02** — Qwen3.6-Plus (closed). [SECONDARY]
- **2026-04-16** — Qwen3.6-35B-A3B (open). [SECONDARY]
- **2026-04-20** — Qwen3.6-Max-Preview (closed). [SECONDARY]
- **2026-04-22** — Qwen3.6-27B (open). [SECONDARY]
- **2026-05-20** — Qwen3.7-Max (closed). [SECONDARY]
- **2026-06-01** — Qwen3.7-Plus (closed). [SECONDARY]
- **2026-07** — Qwen3.7-Flash snapshot/release. [SECONDARY]
- **2026-08-26** — Qwen3.8-Flash-Next (125B/6B). [SECONDARY]
- **2026-08 (SGLang day-0)** — `qwen4_exp` experimental checkpoint (not a Qwen 4 release). [SECONDARY]
- **2026-09-02** — Qwen3.8-Max-0902 snapshot. [SECONDARY]
- **2026-09-17** — 0902 checkpoint reported at AA v4.3 = 45 (China lead). [SECONDARY]


### New verified timeline entries — expansion

- **2026-02-15/16** — Qwen3.5-397B-A17B released (Apache 2.0): 397B/17B, 512 experts, 262K context (1M extensible), 201 languages, early-fusion multimodal. [SECONDARY]
- **2026-03-30** — Qwen3.5-Omni released: text+image+audio+video, 113 languages. [SECONDARY]
- **2026-03-31** — Qwen3.6-Plus Preview: agentic coding, 1M context. [SECONDARY]
- **2026-04-20** — Qwen3.6-Max-Preview released: **first closed-weights Qwen flagship**. [SECONDARY]
- **2026-05-18** — Qwen3.7-Max-Preview / Qwen3.7-Plus-Preview hit Arena (Alibaba #6 lab text, #5 vision). [SECONDARY]
- **2026-05-20** — Qwen3.7-Max GA at Apsara Conference: 1M context, 35-hour agent-task claims. [SECONDARY]
- **July 2026** — Qwen3.8-Max previewed at World AI Conference; preview opens 07-19. [SECONDARY]
- **2026-07-21** — Qwen Image 3.0 / 3.0 Pro released (no benchmarks, no weights, no model card). [SECONDARY]
- **2026-08-03** — Qwen3.8-Max GA: 2.4T/95B, $2/$6, 131K output. [SECONDARY]
- **2026-08-14 (approx)** — Qwen3.8-2.4T-A95B open weights released under the Qwen3.8-Max License (custom, not Apache 2.0); Qwen3.8-27B Apache 2.0. [SECONDARY]
- **2026-08-20** — Alibaba June-quarter earnings: AI Cloud +45% to RMB48.44B; AI products RMB12.38B (12th consecutive triple-digit quarter); capex +75% to RMB67.68B. [SECONDARY]
- **2026-09-02** — Qwen3.8-Max-0902 snapshot: coding/agent post-training, AA Index 45, cost per task $5.41. [SECONDARY]
- **2026-09** — Alibaba's new T-Head AI chip unveiled; 5–10T parameter model plans disclosed; cloud unit reports rationing GPU access. [SECONDARY]
- **2026-07-15** — Qwen app + Doubao shut down user-created AI agents (China humanlike-AI rules); Qwen offered no migration path. [SECONDARY]
- **Early July 2026** — Alibaba cut Qwen3.7-Max ~80% / Qwen3.7-Plus ~60% for international Qoder users off-peak (US developer play). [SECONDARY]
- **2026-08-05** — Qwen Image 3.0 opened to all Qwen AI platform users (was invite-only since 07-21). [SECONDARY]
- **2026-06** — US DoD designated Alibaba a military-linked firm; Alibaba sued in California federal court. [SECONDARY]
- **2026-08-03** — DeepSeek V4-Flash launched same day as Qwen3.8-Max GA (dual Chinese flagship day). [SECONDARY]
- **2026-08-14 (approx)** — Qwen3.8-2.4T-A95B open weights (custom license); Qwen3.8-27B (Apache 2.0). [SECONDARY]
- **2026-09-02** — Qwen3.8-Max-0902 snapshot (AA Index 45, cost/task $5.41). [SECONDARY]

---

## Implications

1. Alibaba runs a two-speed release policy: open Apache-2.0 mid-tiers (3.5/3.6/3.8 ≤27B-class) plus closed/custom-licensed Max flagships — procurement must read the LICENSE file per checkpoint, not per brand. [DIRECTIONAL]
2. The 0902 snapshot is checkpoint news, not generation news; the "Qwen3.8-Max-0902" name invites misreading as a new model. [DIRECTIONAL]
3. The unfulfilled "Plus will be open source" plan is a reminder to treat vendor open-sourcing promises as intent, not inventory. [DIRECTIONAL]
4. `qwen4_exp` must never appear in a Qwen 4 release timeline; conflating an experimental checkpoint with the announced-in-training generation repeats the V4-Lite class of error. [DIRECTIONAL]
5. Qwen's download scale (>3B cumulative, 4.9× Google on HF in 7 months) is the strongest distribution fact for any open-weight family in the corpus — a vendor press claim, but Reuters-cited and directionally consistent with HF's 41%-China figure. [VENDOR]
6. The DeepSWE vendor/secondary split (56.6 vs 69.3) is the canonical "same model, different harness" case in this corpus — any benchmark claim without a named harness is incomplete. [DIRECTIONAL]
7. Qwen3.7-Plus at $0.276/M input (Sept 2026) shows the price war's Chinese front: closed-flagship capability at open-weight prices, repriced quarterly. [SECONDARY]


### New verified implications — expansion

