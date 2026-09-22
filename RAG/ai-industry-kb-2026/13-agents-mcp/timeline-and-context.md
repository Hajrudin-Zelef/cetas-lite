---
id: ai-industry-kb-2026/13-agents-mcp/timeline-and-context
title: "Timeline and context"
domain: agents-mcp
role: deep-dive
task: agents
actors: ["AWS", "Alibaba", "Anthropic", "CISA", "DeepSeek", "EU", "Google", "Meta", "Microsoft", "Mistral", "Nvidia", "OpenAI", "SpaceX", "StepFun", "xAI"]
dates: ["2024-11-05", "2025-03-26", "2025-05", "2025-06-18", "2025-10-22", "2025-11-25", "2026-01-06", "2026-03", "2026-03-25", "2026-04", "2026-04-02", "2026-04-10", "2026-04-14", "2026-04-16", "2026-05", "2026-05-19", "2026-05-21", "2026-05-24", "2026-05-26", "2026-05-29", "2026-06", "2026-06-02", "2026-06-08", "2026-06-10", "2026-06-16", "2026-07", "2026-07-18", "2026-07-24", "2026-07-28", "2026-07-30", "2026-08-01", "2026-08-02", "2026-08-06", "2026-08-10", "2026-08-13", "2026-08-17", "2026-08-28", "2026-09", "2026-09-01", "2026-09-03", "2026-09-08", "2026-09-10", "2026-09-11", "2026-09-19", "2026-09-20", "2026-09-21", "2026-09-22"]
keywords: ["acquisition", "agent", "agentic", "agents", "arr", "astra", "aws", "backlog", "benchmark", "benchmarks", "chatgpt", "claude"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [6793, 6931]
section: "13. Agents & MCP"
sha256: 9c0573b3abe2451aea8ceb005cf26b03132e0973fe2244518503bd94b73756e9
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

1. **MCP won the integration war; governance now decides whether it survives its own success.** The 97M→110M SDK series, the registry growth (9,652 → 12,000+), and the 41% production figure are verified; the OpenAI + Anthropic joint stewardship (AAIF, Dev Summit NYC) institutionalized it. But five breaking revisions in ~20 months with a deprecation policy formalized only in the latest one is the material enterprise risk: most deployed clients still speak 2025-06-18/2025-11-25. The enterprise roadmap (SSO-integrated auth, workload identity federation, gateway standardization) is directionally verified — the stability question is whether the protocol's own churn rate slows now that the lifecycle policy exists.
2. **The RC caveat is dead; replace it with the adoption-lag caveat.** Consolidation must present 2026-07-28 as Final (RC locked 2026-05-21, RC published 2026-05-29, final 2026-07-28; tier-1 SDKs same day). The real, non-obsolete caveat is that deployed clients lag the spec by one or two breaking revisions — that is the fact to surface, not RC status.
3. **Audit trails are guidance, not spec — yet.** The enterprise requirement is confirmed as 2026 guidance converging on gateway-mediated, allow-listed deployment (audit logging, tool allowlisting, identity binding, human-in-the-loop approval for high-risk actions). The EU AI Act's high-risk obligations (enforceable 2026-08-02) add indirect regulatory pressure on logging and oversight for agentic tool-use. Expect audit trails to move from guidance to compliance requirement in the RAG's forward window.
4. **Config portability is aspiration, not a shipped feature.** Documented pain point (Goose YAML vs Codex TOML per-host formats); do not present it as a 2026 deliverable.
5. **The benchmark world requires provenance tags.** Vendor-reported scores run 10–30 points above standardized harnesses; effort/turn budgets and sandbox choice move numbers 5–15 points; harness choice moves 6+ points (Fable 5.1: 91.4% Artificial Analysis vs 85.02% vals.ai, Terminus 2). The July 2026 SWE-bench Pro audit (~30% of public split broken, recommendation retracted) plus Datacurve's verifier audit make cross-harness comparison untrustworthy. 2026-native benchmarks (DeepSWE, Terminal-Bench 4.0, MCP-Atlas) re-differentiate — use them, provenance-tagged.
6. **Coding agents crossed from autocomplete to execution, and the guardrail pattern is now consensus.** Sandboxed shell by default (Codex CLI bubblewrap/Seatbelt; Claude Code permission prompts; OpenHands Docker; Jules ephemeral GCP VMs); plan-review gates; approval policies; MCP allowlists. The residual open problem: unsupervised desktop-GUI autonomy is empirically bounded (~20.6% binary success at 500 steps over multi-app workflows), and the surviving computer-use product is Anthropic Computer Use (OpenAI's line retired Aug 2026).
7. **Test-time compute is now a system-design and pricing axis.** DeepSeek V4.1-Flash (890 bytes/token KV, reasoning dial, 50× cache-hit/miss price gap) is the template: providers price-shape agent behavior, not just token volume. Agent economics must record cache-hit rates and tokens-per-task alongside per-token prices; per-token price ≠ per-task cost (sub-agent fan-out surges total tokens while per-call cost falls). The workload inversion (90–95% of GPU cycles in decode, 16K–128K thinking tokens per request [DIRECTIONAL]) reframes infrastructure planning: the bottleneck is re-reading context, not writing answers.
8. **The protocol layer is the agent-economics lever that matters most for infrastructure.** With reasoning models emitting 16K–128K thinking tokens per request [DIRECTIONAL], protocol-level caching semantics (`ttlMs` + `cacheScope` on list APIs, client-side `defer_loading`) and provider cache-hit pricing (DeepSeek's $0.006/M cache read) determine per-task cost more than raw per-token price. The consolidated document should treat MCP's caching/discovery semantics as cost infrastructure, not plumbing.
9. **Vendor compression claims need denominator discipline.** The V4.1-Flash "437×" episode is the template: always record the denominator (2023 4K-context V1) alongside the headline; prefer the replacement-model comparison (~4× vs V4-Flash) for operational claims.
10. **Agent benchmark numbers keep splitting by harness.** SWE-2's 50% (FrontierCode, human-merge judgment) vs 27.3% (TB4.0) and Grok 4.7's 71.0% (DeepSWE) vs 38% (TB4.0) reinforce the trust-crisis rule: provenance-tag every score and never compare across harnesses. The MCP-Atlas 88%→69% single-to-cross-server gap quantifies multi-server orchestration as the current frontier of tool-use difficulty.
11. **Safety incidents and the agent economy belong to their sections.** Agent-caused incidents (Cline CI compromise, "Comment and Control", Unit 42 campaigns, "Agents of Chaos") are filed in **§17 — Safety & Security** (one-line cross-ref, no duplication here). Agent economy figures (Cursor $4B ARR, Cognition $48B Series E, funding, billing incidents) are filed in **§19 — AI Economics** (attributed forecasts only).
12. **Vendor compression claims need denominator discipline.** The V4.1-Flash "437×" episode is the template: always record the denominator (2023 4K-context V1) alongside the headline; prefer the replacement-model comparison (~4× vs V4-Flash) for operational claims. The general rule extends to benchmarks: harness, effort budget, and sandbox move scores more than most model deltas the RAG will compare.
14. **The 88%→69% MCP-Atlas gap is the next frontier, not a failure.** Single-server tool-use is at 88.1% (Muse Spark 1.1) while cross-server drops to 69.2% (Gemini 3.1 Pro) — multi-server orchestration (discovery, namespace collisions, auth per server) is where 2026 tooling lags. Gateway standardization on the 2026 roadmap is aimed exactly at this gap.
15. **Client-side answers can beat protocol answers.** SEP-1576 / issue #2808 (protocol-level progressive disclosure) were closed with no spec-level answer; Anthropic's tool search + `defer_loading` solved the large-registry token problem in the client instead. When the RAG evaluates "what solved X," the answer is sometimes an implementation, not a spec.

13. **The surviving products define the deployable envelope.** By September 2026 the agent-product landscape had been pruned by market and incident pressure: OpenAI's Operator/ChatGPT-agent/Atlas browser line is gone, Anthropic Computer Use is the surviving computer-use product, SWE-agent is maintenance-only behind mini-SWE-agent, Cascade is sunset behind Rust-rewritten Devin Local, Windsurf the brand is retired. The consolidation should name survivors and the superseded, not just the new.

## Sources and URLs

- MCP specification versioning page (primary — "current protocol version is 2026-07-28"): https://modelcontextprotocol.io/specification/versioning [VERIFIED-PRIMARY]
- Official MCP blog — 2026-07-28 release announcement: https://blog.modelcontextprotocol.io/posts/2026-07-28/ [VERIFIED-PRIMARY]
- Official MCP blog — SDK betas for 2026-07-28: https://blog.modelcontextprotocol.io/posts/sdk-betas-2026-07-28/ [VERIFIED-PRIMARY]
- 2026-07-28 spec reference article (Sept 2026, drawn from the spec + changelog): https://dev.to/jarvisbitztech/mcp-in-2026-what-changed-in-the-2026-07-28-specification-and-how-to-design-production-integrations-16c4 [SECONDARY]
- Migration impact analysis (stateless core, SEPs): https://github.com/crimsonsunset/mcp-mux/blob/HEAD/docs/planning/mcp-2026-07-28-spec-impact.md [SECONDARY]
- Spec evolution research (June 10, 2026 — RC-era tracker, now outdated on status): https://github.com/mayanpathak/gateway/blob/HEAD/docs/research/mcp-spec.md [SECONDARY]
- Outdated tracker (2025-11-25-as-Final advice, RC-era — do not cite as current): https://github.com/azigler/aaif/blob/HEAD/refs/projects/mcp.md [OUTDATED]
- 2026-07-28 final confirmation trackers: https://github.com/oaslananka/kicad-mcp/blob/HEAD/docs/adr/0006-mcp-2026-stateless-compatibility-lane.md ; https://github.com/rul1an/assay/issues/1943 ; https://github.com/nulab/backlog-mcp-server/issues/170 ; https://github.com/lexus2016/hermes-agent-evolution/issues/1414 [COMMUNITY]
- MCP 97M installs / Dev Summit reveal (97M on 2026-03-25, 4,750% growth, roadmap SSO/gateway): https://medium.com/@jahanzaibai/mcp-97-million-installs-what-the-dev-summit-revealed-7c412fc62a01 [SECONDARY]
- 110M/month downloads (June 2026, Dev Summit India): https://levelup.gitconnected.com/mcp-has-110-million-downloads-a-month-most-teams-still-cant-use-it-in-production-b39b709a3760 [SECONDARY]
- Stacklok "State of Model Context Protocol in Software 2026" (PDF, 2026-01): https://stacklok.com/wp-content/uploads/2026/01/State-of-MCP-in-Software-2026_FINAL.pdf [VERIFIED-PRIMARY]
- MCP adoption statistics compilation (Stacklok table, retraction note on the 78% figure): https://www.digitalapplied.com/blog/mcp-adoption-statistics-2026-model-context-protocol [SECONDARY]
- The LLM Book 2026 ed. (Stacklok figures + "Do not use it" retraction note): https://vstorm.co/the-llm-book/ebooks/the-llm-book.pdf [SECONDARY]
- MCP Marketplaces April 2026 (12,000+ servers, 33 registries): https://dev.to/matthias_studiomeyer/mcp-marketplaces-in-april-2026-a-field-report-from-33-platforms-33pn [SECONDARY]
- MCP just went stateless (Aug 2026, SDK readiness, TS/Python 1B cumulative claim): https://medium.com/@vinoth.lingam333/mcp-just-went-stateless-the-protocols-biggest-rewrite-yet-01485e5c75d9 [SECONDARY]
- The Enterprise MCP Guide 2026 (gateway, audit logging, EU AI Act, ROI figures): https://medium.com/@28thjun/the-enterprise-mcp-guide-2026-the-agentics-5f4a68dc6102 [SECONDARY]
- MCP Complete Guide (97M, 13,000+ GitHub servers, Gartner gateway forecast): https://medium.com/design-bootcamp/the-complete-guide-to-mcp-everything-a-developer-needs-to-know-37bf0968f185 [SECONDARY]
- MCP architecture guide 2026 (USB-C analogy, 97M March 2026, 10,000 servers): https://neuralcoretech.com/agentic-ai-model-context-protocol-mcp-architecture-2026/ [SECONDARY]
- AAIF PR — MCP Dev Summit North America 2026 (April 2–3, NYC, 95+ sessions): https://pr.gulfmainmagazine.com/article/Agentic-AI-Foundation-Unveils-MCP-Dev-Summit-North-America-2026-Schedule/699c5e3cf3b9ec0450c8d2ad [VENDOR]
- Agentic-AI briefing (97M, Dev Summit transition framing, X-sourced): https://github.com/yoselabs/insights-trail/blob/HEAD/data/2026-04-10/agentic-ai/briefing.md [COMMUNITY]
- Seven Stages of AI field notes (97M March 2026, Stacklok 41%, NSA MCP guidance May 2026): https://www.amitjadhav.com/field-notes/002-seven-stages-of-ai [SECONDARY]
- MCP enterprise guide 2026 (10k+ servers, 500M downloads claim — do not use as fact): https://aibuzz.blog/model-context-protocol-explained/ [SECONDARY]
- MCP revision churn and 2026-07-28 deprecation policy: https://github.com/kanthipm/medpullkiosk/blob/HEAD/recovery-copilot/docs/backend-design/research/aiarch--model-context-protocol-mcp-what-it-is-in-mid-2026-healt.md [SECONDARY]
- MCP tool search + defer_loading (client-side token fix): https://github.com/albretsen/mcpemails/blob/HEAD/docs/token-cost/research-spec.md [COMMUNITY]
- Enterprise MCP guide (NIST initiative, EU AI Act, ROI): https://medium.com/@28thjun/the-enterprise-mcp-guide-2026-the-agentics-5f4a68dc6102 [SECONDARY]
- DeepSeek V4.1-Flash architecture (552B MoE, CED, CSA2, MXFP4, MIT weights): https://pondero.ai/news/2026-09-11-deepseek-v41-flash/ ; https://www.intelligentliving.co/deepseek-v41-flash-pricing-release/ ; https://docs.agenteum.top/blog/2026-09-10-deepseek-v41-flash-open-weights [SECONDARY]
- DeepSeek V4.1-Flash reasoning dial and agent economics: https://www.theneuron.ai/explainer-articles/deepseek-v41-flash-explained-how-it-cuts-ai-memory-8x/ [SECONDARY]
- DeepSeek V4.1-Flash KV 890 bytes/token, pricing, critique of 437× denominator: https://www.techtimes.com/articles/327755/20260919/deepseek-cuts-ai-agent-memory-cost-4x-new-architecture-fits-more-sessions-per-gpu.htm ; https://www.youtube.com/watch?v=n-_e9bROGHo [SECONDARY]
- DeepSeek V4.1-Flash harder-benchmark context (TB4.0 31.2, HLE 36.8): https://temperaturezero.com/2026/09/10/deepseek-v4-1-flash-inference-kv-cache-benchmark-analysis/ [SECONDARY]
- DeepSeek V4.1-Flash price cuts and agent price analysis (Sept 10): https://amdatalakehouse.substack.com/p/ai-weekly-deepseek-cuts-prices-as [SECONDARY]
- DeepSeek V4.1-Flash API changelog (beta window, model-ID routing): https://github.com/uwuclxdy/ai-pricelog/blob/HEAD/state/announce/deepseek/updates.md [COMMUNITY]
- Test-time compute workload inversion (practitioner series): https://kenhuangus.substack.com/p/chapter-8-test-time-compute-and-reasoning [SECONDARY]
- Qwen3 Thinking Mode Fusion (thinking default): https://github.com/bayesiansapien/cere-bro/blob/HEAD/raw/rss/2026-07-18-ahead-of-ai-controlling-reasoning-effort-in-llms.md [COMMUNITY]
- CoT faithfulness debate (ACL 2026, practitioner summary): https://medium.com/@ai4medical/chain-of-thought-isnt-lying-to-you-it-s-just-leaving-things-out-c47e2fccca87 [SECONDARY]
- Cognition SWE-2 launch, FrontierCode, Poke acquisition (Sept 11): https://www.newsbytesapp.com/news/science/swe-2-challenges-rivals-with-cost-efficiency/story [VENDOR]
- StepFun Step 5 (600B, 1M context, DeepSWE 67.7): https://areeblog.com/stepfuns-600b-step-5-targets-ai-coding-agents-with-1m-token-context/ [SECONDARY]
- Grok 4.7 launch (DeepSWE 71.0%, $2/$6 pricing): https://aiweekly.co/alerts/xai-ships-grok-47-at-26-per-million-tokens-deepswe-71 ; https://en.coinotag.com/spacexai-grok-4-7-launch-deepswe-benchmark [VENDOR]

