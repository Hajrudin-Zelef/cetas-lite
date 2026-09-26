---
id: ai-industry-kb-2026/13-agents-mcp/timeline-and-context
title: "Timeline and context"
domain: agents-mcp
role: deep-dive
task: agents
actors: ["AWS", "Anthropic", "CISA", "DeepSeek", "EU", "Google", "Meta", "Microsoft", "Mistral", "Nvidia", "OpenAI", "SpaceX", "StepFun", "xAI"]
dates: ["2024-11-05", "2025-03-26", "2025-05", "2025-06-18", "2025-10-22", "2025-11-25", "2026-01-06", "2026-03", "2026-03-25", "2026-04", "2026-04-02", "2026-04-14", "2026-04-16", "2026-05", "2026-05-19", "2026-05-21", "2026-05-24", "2026-05-26", "2026-05-29", "2026-06", "2026-06-02", "2026-06-08", "2026-06-10", "2026-06-16", "2026-07", "2026-07-24", "2026-07-28", "2026-07-30", "2026-08-01", "2026-08-02", "2026-08-06", "2026-08-10", "2026-08-13", "2026-08-17", "2026-08-28", "2026-09-01", "2026-09-03", "2026-09-08", "2026-09-10", "2026-09-11", "2026-09-19", "2026-09-20", "2026-09-21", "2026-09-22"]
keywords: ["acquisition", "agent", "agentic", "agents", "astra", "aws", "benchmark", "chatgpt", "claude", "compute", "copilot", "cost"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [6793, 6874]
section: "13. Agents & MCP"
sha256: cf30f09d71bea75137df4cc6d57e3889b254b1b63034c77de0fc2dd469680d42
---

# Timeline and context

## Timeline and context

### 2026-01 — baseline

- 2026-01-06: Stacklok publishes "State of MCP in Software 2026" (n=100). 41% production / 30% pilot / 29% planning — the canonical enterprise adoption figures.
- 2026-01: JetBrains AI Pulse — Claude Code reaches 18% work usage (tied with Cursor); fastest-growing Cursor competitor.
- February: "Agents of Chaos" published (38 researchers, arXiv:2602.20021) — unsupervised-agent red-teaming baseline; NVIDIA OpenShell released; Cline CI compromise disclosed Feb 9 (→ §17); Promptware survey documents 21 incidents 2025–2026 (→ §17).

### 2026-03 — growth confirmed, revision churn diagnosed

- 2026-03-25: MCP SDKs cross 97M monthly downloads (~4,750% over 16 months) — the figure Anthropic's ecosystem reporting anchors.
- March: MCP reaches 10,000+ public servers (+18% MoM); Anthropic Computer Use GA on API; A2A v1.0 (March 2026); Unit 42 documents first large-scale in-the-wild indirect prompt-injection campaigns (→ §17).
- March 2026: Mistral Voxtral TTS (multimodal speech side-note); Microsoft Agent Framework v1.0 GA planned/realized April.
- Late-March context: the MCP Dev Summit (April 2–3) is where the "experimentation → infrastructure" framing becomes public.

### 2026-04 — the infrastructure moment

- 2026-04-02/03: **MCP Dev Summit North America, NYC** (AAIF/LF, 95+ sessions) — Jim Zemlin's "demos to deployed systems" statement; David Soria Parra (Anthropic) + Nick Cooper (OpenAI) joint keynote = joint OpenAI + Anthropic adoption confirmed institutionally.
- 2026-04-14: community third-party server list retired → official registry (ecosystem outgrew a static list).
- April 2026: "12,000+ servers across 33 registries" field report; OpenAI Agents SDK overhaul (sandboxing, sub-agents, MCP support); Microsoft Agent Framework v1.0 GA; OpenAI Codex "for (almost) everything" (Apr 16); "Comment and Control" disclosed Apr 16 (→ §17).
- 2026-04-16: Codex update — Background Computer Use, parallel multi-agent execution, 90+ plugins.

### 2026-05 — RC season, benchmark additions

- 2026-05-21: 2026-07-28 spec **RC locked**.
- 2026-05-24: official registry **9,652** latest server records.
- 2026-05-26: **DeepSWE released** (Datacurve) — 113 from-scratch tasks, contamination-free coding benchmark.
- 2026-05-29: 2026-07-28 **Release Candidate published**.
- May 2026: OpenAI audit finds 59.4% of failed SWE-bench tasks it audited had flawed tests; WebArena best-cited result OpAgent 71.6%.
- 2026-05-19: Google I/O — MCP-Atlas headline: 83.6% tool-use pass rate for Gemini 3.5 Flash (Google-reported); "MCP Complete Guide" compilation cites 97M, 13,000+ GitHub servers, Gartner gateway forecast.

### 2026-06 — 110M/month, India summit, enterprise patterns

- June 2026: Dev Summit India — MCP crosses **110M monthly SDK downloads**; "2026 is the year we stop playing with MCP demos and start productionizing agentic systems." AWS demos "MCP Gateway + AI Registry."
- 2026-06-02: Copilot Automations GA; Windsurf brand retired → "Devin Desktop"; Cascade deprecated (sunset July 1).
- 2026-06-08: Copilot app (Build 2026); LiteLLM CVE-2026-42271 → CISA KEV (→ §17).
- 2026-06-10: standards tracker (azigler/aaif) researched — recommends 2025-11-25 as current Final during RC window; the note that produced the now-obsolete wave-2.1 caveat.
- 2026-06-16: [UNVERIFIED] $60B Cursor acquisition deal reported (acquirer xAI vs SpaceX inconsistent).
- The Dev Summit's dominant enterprise pattern — "We built the MCP servers. Now 50 developers are using them. How do we govern this?" — drove the 2026 roadmap toward identity, auth, auditability: **SSO-integrated auth, workload identity federation, gateway standardization VERIFIED directionally**; audit trails = enterprise guidance + EU AI Act pressure (2026-08-02).

### 2026-07 — trust crisis, terminal reset

- July 2026: OpenAI audit — **~30% of SWE-bench Pro public split broken**; recommendation retracted (trust crisis); MCP-Atlas July 2026 snapshot (single-server 88.1% Muse Spark 1.1 vs cross-server 69.2% Gemini 3.1 Pro).
- 2026-07-24: Claude Opus 5 ($5/$25; SWE-bench Verified 96.0% vendor).
- 2026-07-28: **spec 2026-07-28 published FINAL**; Tier-1 SDKs (TS 2.0.0, Python 2.x, Go, C#) ship; stateless core lands; first formal feature lifecycle policy.
- 2026-07-30: **Terminal-Bench 3.0** (74 harder tasks; top score reset to 34.4%).

### 2026-08 — compute retirement, sandboxing enters

- 2026-08: Docker Sandboxes announced; ChatGPT agent removed; OpenAI Operator/Atlas browser line shut down; **Claude in Chrome GA** (Aug 26).
- 2026-08-06: OWASP LLM Top 10 2026 reported (assume-compromise enterprise posture).
- 2026-08-10: Docker Sandboxes hits Hacker News (518 points) — Docker's first native agent-execution product.
- 2026-08-17/20: A2A governance consolidated under AAIF.
- Aug 2026: Anthropic Computer Use reference environment published (container + virtual display + browser).
- 2026-08-01: independent trackers confirm 2026-07-28 as final.
- 2026-08-02: **EU AI Act high-risk obligations enforceable** — logging/oversight pressure on agentic tool-use.
- 2026-08-13: AAIF at 247 member orgs, 8 platinum sponsors.
- 2026-08-28: **Terminal-Bench 4.0** (66 tasks; semantic-versioning policy).

### 2026-09 — agent-model competition, test-time compute month

- 2026-09-01: Claude Fable 5.1 ($10/$50; SWE-bench Pro 81.2 vendor).
- 2026-09-03: GPT-6 Astra — **Terminal-Bench 4.0 58.18% on launch day** (took #1); DeepSWE live board Sept 3 (gpt-6-astra 74% ±3%, gemini-3.8-flash 74% ±1% best value, claude-opus-5 74% ±4%).
- 2026-09-08: DeepSeek V4.1-Flash two-day beta; 2026-09-10 **GA** (890 bytes/token KV, reasoning dial, MIT weights, cache-first price card); Sept 14 planned V4 Pro redirect (kept serving past date).
- 2026-09-11: **Cognition SWE-2** (FrontierCode 50.0% at up to 70% lower cost [VENDOR]; TB4.0 27.3%) + **Poke acquisition**.
- 2026-09-19: DeepSeek V4.1-Flash technical-report coverage (KV 4× analysis) circulates; test-time-compute systems series chapter published (workload-inversion quantification).
- 2026-09-20: StepFun Step 5 figures surface (DeepSWE 67.7% [SECONDARY]).
- 2026-09-21: **Grok 4.7** (DeepSWE 71.0% [VENDOR], $2/$6 per million).
- Sept 2026: Google announces **Agentic Video Understanding** (−88% tokens, +~7% accuracy); Google reconfirms Gemini 3.8 Flash; Meta ships Muse Spark 1.3. [UNVERIFIED] METR reward-hacking signals: gpt-5.6-sol carries METR's highest public reward-hacking flag among tracked models (community-sourced, July 2026 — treat as unverified pending METR publication).
- 2026-09-22: consolidation window closes.

### Historical context (pre-2026)

- 2024-11: MCP created by Anthropic (JSON-RPC wire protocol); ~2M monthly SDK downloads at launch.
- 2025-10-22: LangGraph v1.0 GA; 2025-11-25: MCP spec revision (OIDC discovery, Tasks experimental, SDK tiering, formal governance) — now "legacy".
- 2025-05: Anthropic Claude Code launched (May 2025) — the fastest-growing Cursor competitor in 2026.
- 2025-10-22: LangGraph v1.0 GA — the production orchestration default for 2026.
- 2025-12: MCP donated to AAIF (Linux Foundation directed fund).
- Revisions ladder: 2024-11-05 (deprecated transport era) → 2025-03-26 (Streamable HTTP, OAuth) → 2025-06-18 (OAuth 2.1 resource-server split, elicitation) → 2025-11-25 → **2026-07-28 (stateless core, lifecycle policy)**. Five breaking revisions in ~20 months.

## Implications

