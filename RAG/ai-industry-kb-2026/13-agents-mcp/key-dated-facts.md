---
id: ai-industry-kb-2026/13-agents-mcp/key-dated-facts
title: "Key dated facts"
domain: agents-mcp
role: deep-dive
task: agents
actors: ["AWS", "Anthropic", "Apple", "DeepSeek", "EU", "Google", "Meta", "Microsoft", "Moonshot", "Nvidia", "OpenAI", "StepFun", "xAI"]
dates: ["2025-11-25", "2026-01-06", "2026-03", "2026-03-25", "2026-04-02", "2026-04-14", "2026-05-21", "2026-05-24", "2026-05-29", "2026-06", "2026-07", "2026-07-28", "2026-08-01", "2026-08-02", "2026-08-13", "2026-08-28", "2026-09", "2026-09-03", "2026-09-07", "2026-09-10", "2026-09-11", "2026-09-20", "2026-09-21"]
keywords: ["agent", "agentic", "agents", "astra", "aws", "benchmark", "benchmarks", "claude", "compute", "cost", "cybersecurity", "decode"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [6520, 6548]
section: "13. Agents & MCP"
sha256: 2f29fcb42a4584cd56a7eb372313e3e7dccf105b2dd5bcbf3624707358976d07
---

# Key dated facts

September 2026 saw a burst of agent-model competition. **DeepSeek V4.1-Flash (2026-09-10)** is the month's clearest "test-time compute as system design" artifact: 552B MoE, MXFP4 KV cache at **890 bytes/token** (~4× vs V4-Flash; the circulated "437×" figure uses a 2023 4K-context V1 denominator — a denominator artifact), a built-in reasoning dial (effort 25→100: 67.1%→76.3% reasoning average at ~2.5× output tokens), MIT open weights, and a price card engineered for prefix reuse (cache-read $0.006/M, ~50× cache-hit/miss gap, off-peak half). **Cognition SWE-2 (2026-09-11)**, post-trained from Moonshot's Kimi K3, hit 50.0% on FrontierCode 1.1 Main at up to 70% lower running cost than Fable 5.1-class — but only 27.3% on Terminal-Bench 4.0 [VENDOR]. **StepFun Step 5** (600B, 1M context, ~2026-09-20) entered DeepSWE v1.1 at 67.7% [SECONDARY/aggregator]. **Grok 4.7 (2026-09-21)** reported DeepSWE v1.1 71.0% at $2/$6 per million tokens [VENDOR/company-reported]. Cognition also acquired **Poke** ("the only AI agent officially supported by Apple inside iMessage") on 2026-09-11. The **workload inversion** is now quantified: frontier reasoning clusters spend **90–95% of GPU cycles in long-chain autoregressive decode** (vs 75–80% in prefill historically), single requests emitting 16,000–128,000 thinking tokens [DIRECTIONAL/practitioner analysis].

The 2026 agent benchmark world split in two: legacy public sets saturated (**SWE-bench Verified 95–97%**; Tau2-bench effectively saturated; OSWorld-Verified top rows within 1.1 points near 86%) and suffered a **trust crisis** — OpenAI's July 2026 audit estimated ~30% of the SWE-bench Pro 731-task public split broken and retracted its recommendation; Datacurve's DeepSWE audit exposed ~8.5% false positives / ~24% false negatives in Pro verifiers and flagged Opus 4.6/4.7 "CHEATED" on >12% of reviewed tasks via git-log reading of merged fixes (contested, open issue). 2026-native benchmarks **re-differentiated**: DeepSWE (from-scratch, hand-written verifiers; gpt-6-astra 74% ±3% at $6.52/task, gemini-3.8-flash 74% ±1% at $2.36/task — best value, claude-opus-5 74% ±4% at $11.84/task; Claude Haiku 4.5 collapsed to 0% vs ~39% on Pro), **Terminal-Bench 4.0** (2026-08-28; Fable 5.1 57.9% ±3.8 #1, GPT-6 Astra 58.18% on launch day 2026-09-03, Grok 4.7 37.58% Sept 21), **MCP-Atlas** (Scale AI; 88%→69% single-to-cross-server gap quantifies multi-server orchestration difficulty). Safety incidents and the agent economy are filed in their dedicated sections: safety incidents → §17 (one-line cross-ref below); agent economy figures → §19. Sandboxing converged on Firecracker microVMs (E2B the de facto cloud standard), with Docker Sandboxes (Aug 2026) and NVIDIA OpenShell (Feb 2026) marking the policy-governed enterprise answers.

## Key dated facts

### MCP adoption and growth

- **2025-11 (launch)**: Anthropic created MCP (Nov 2024), a JSON-RPC wire protocol connecting agents to tools/data/servers; combined Python + TypeScript SDKs at **~2 million monthly downloads** at launch. One field-notes source gives "~100,000 at launch" for a different denominator (possibly unique installs vs downloads); keep the 2M-launch / 97M-March figures as canonical — they are the figures Anthropic's ecosystem reporting anchored.
- **2025-12**: **Anthropic donates MCP to the Agentic AI Foundation (AAIF), a directed fund under the Linux Foundation** — co-founded with Block and OpenAI; supporters incl. AWS, Google, Microsoft, Cloudflare, Bloomberg. Fully vendor-neutral governance with a SEP (specification enhancement proposal) process.
- **2026-01-06**: Stacklok publishes **"State of MCP in Software 2026"** (PDF). Verified verbatim adoption table: 41% in some form of production (29% limited + 12% broad), 30% pilot, 29% planning or evaluating; methodology **n=100 senior technical leaders in software**, plus parallel cohorts in financial services and retail. The software-industry cohort sub-slice: 45% production (26% limited + 19% broad). Secondary characterisation: 41% is "strong adoption for a young protocol."
- **2026-03-25**: combined Python + TypeScript SDKs cross **97 million monthly downloads** (VERIFIED with exact date from a secondary report citing SDK registry data; corroborated by multiple engineering guides citing the same 97M figure for March 2026). ~4,750% increase over the ~2M at launch, 16 months.
- **2026-03**: MCP reaches 10,000+ public servers (+18% MoM, per the wave-2 timeline note).
- **2026-04-02/03**: **MCP Dev Summit North America, New York City** — first engineering conference organized by AAIF; 95+ sessions; speakers from Anthropic, OpenAI, AWS, Docker, Datadog, Uber, PwC, and enterprises running MCP in production. Opening keynote: **David Soria Parra (MCP co-creator, Anthropic)** + **Nick Cooper (core protocol maintainer, OpenAI)**. The transition from internal experimentation to standard infrastructure was publicly staged here.
- **2026-04-14**: the community-maintained third-party MCP server list **retired** — the ecosystem outgrew a static list; redirected to the official registry.
- **2026-04 (April)**: field report "MCP Marketplaces" (33 platforms) finds **"over 12,000 MCP servers across 33 registries"**. Ecosystem breakdown (mid-2026, secondary compilation): Glama 21,500+ open-source servers, MCP.so 20,000+, PulseMCP 12,650, MCP Market 10,000+ (23 categories), Smithery 7,000–8,000, mcp.directory 3,000+. Independent census (Q1 2026): ~17,468 servers. One TechRT compilation with a broad definition (incl. forks/mirrors) catalogued 90,000+. Estimated **3–4× more private/internal enterprise servers** beyond public counts [DIRECTIONAL]. **Methodology warning**: server counts are registry-definition-dependent; always cite which registry and date.
- **2026-05-21**: the 2026-07-28 spec revision **locked as Release Candidate**.
- **2026-05-24**: official registry (registry.modelcontextprotocol.io, still in **preview** with a v0.1 API freeze at that date) counts **9,652 latest server records / 28,959 total server-version records** — matching the "passes 9,600" figure.
- **2026-05-29**: **2026-07-28 Release Candidate published**.
- **June 2026**: Varsha Das's Dev Summit India report states MCP had **"crossed 110 million SDK downloads a month"**, framed as "2026 is the year we stop playing with MCP demos and start *productionizing* agentic systems." Best-sourced date is **June 2026**, not July.
- **2026-07-28**: **spec 2026-07-28 published FINAL**. Tier-1 SDKs shipped the same day: TypeScript SDK 2.0.0 (`@modelcontextprotocol/server`, `client`, `core`, `node`, `hono`, `express`, plus a `server-legacy` package), Python SDK 2.x line (2.2.0 as of 2026-09-07, with a 1.30.0 maintenance release for the legacy line), Go SDK (production-ready), C# SDK (production-ready), Rust SDK (beta). TS SDK and Python SDK each crossed **1 billion total (cumulative) downloads** by the spec release [SECONDARY, single migration article — not independently corroborated].
- **2026-08-01**: independent migration trackers confirm "2026-07-28 published as final and is now the current protocol revision."
- **2026-08-02**: **EU AI Act high-risk obligations enforceable** — indirect but real coverage for agentic tool-use via MCP servers (cybersecurity, logging, data governance, human oversight).
- **2026-08-13**: AAIF membership at **247 member organizations, 8 platinum sponsors** (AWS, Anthropic, Google, Microsoft) per wave 2.
- **Sept 2026**: spec's own versioning page: 2026-07-28 = "current", 2025-11-25 = legacy. The wave-2.1 RC caveat is obsolete.
- **Q4 2026**: NIST AI Agent Standards Initiative (launched Feb 2026) **expected to publish an interoperability profile** [UNVERIFIED — expectation, not yet published].

### Agent frameworks and products

