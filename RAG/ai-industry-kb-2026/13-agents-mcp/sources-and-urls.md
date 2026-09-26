---
id: ai-industry-kb-2026/13-agents-mcp/sources-and-urls
title: "Sources and URLs"
domain: agents-mcp
role: deep-dive
task: agents
actors: ["Anthropic", "DeepSeek", "EU", "Google", "Meta", "OpenAI", "xAI"]
dates: ["2025-06-18", "2025-11-25", "2026-05-21", "2026-05-29", "2026-07", "2026-07-28", "2026-08-02", "2026-09"]
keywords: ["agent", "agentic", "agents", "arr", "benchmark", "benchmarks", "chatgpt", "claude", "compute", "cost", "decode", "deepseek"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [6875, 6893]
section: "13. Agents & MCP"
sha256: 19f0e0d4c0561b826d631a56f5bd36f17d3b1db6b361e69388cf3c6d8424d057
---

# Sources and URLs

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

