---
id: ai-industry-kb-2026/13-agents-mcp/figures-and-metrics
title: "Figures and metrics"
domain: agents-mcp
role: deep-dive
task: agents
actors: ["AWS", "Alibaba", "Anthropic", "Cerebras", "DeepSeek", "EU", "Google", "Hugging Face", "Meta", "Microsoft", "Moonshot", "Nvidia", "OpenAI", "Perplexity", "SpaceX", "StepFun", "Stripe", "United States", "Z.ai", "xAI"]
dates: ["2025-04", "2025-05", "2025-10-22", "2026-02-05", "2026-03", "2026-03-25", "2026-04", "2026-05", "2026-05-19", "2026-05-24", "2026-05-26", "2026-05-31", "2026-06", "2026-06-02", "2026-06-16", "2026-07", "2026-07-01", "2026-07-11", "2026-07-28", "2026-07-30", "2026-08-02", "2026-08-06", "2026-08-13", "2026-08-28", "2026-09-03", "2026-09-10", "2026-09-11", "2026-09-19", "2026-09-20", "2026-09-21", "2026-09-22"]
keywords: ["acquisition", "agent", "agentic", "agents", "apache", "arr", "astra", "aws", "bedrock", "benchmark", "benchmarks", "chatgpt"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [6654, 6792]
section: "13. Agents & MCP"
sha256: d670a74294c4a624abcba4ae93b077e1dc359d9bbf948c9cb15c469cbf99f336
---

# Figures and metrics

## Figures and metrics

### MCP scale figures (canonical series)

| Metric | Value | Date | Provenance |
|---|---|---|---|
| Combined SDK monthly downloads at launch | ~2M | 2024-11 | Anthropic ecosystem reporting (canonical) |
| Combined SDK monthly downloads | **97M** | 2026-03-25 | VERIFIED (secondary report citing SDK registry data; corroborated) |
| Combined SDK monthly downloads | **~110M** | 2026-06 | VERIFIED (Dev Summit India reporting) |
| Official registry, latest server records | **9,652** | 2026-05-24 | VERIFIED (registry survey) |
| Official registry, server-version records | 28,959 | 2026-05-24 | VERIFIED (registry survey) |
| Servers across 33 registries | **12,000+** | 2026-04 | VERIFIED (field report) |
| Independent census estimate | ~17,468 | 2026-03 (Q1) | VERIFIED |
| Private/internal enterprise servers | 3–4× public counts | 2026 | [DIRECTIONAL/estimation] |
| AAIF membership | 247 orgs, 8 platinum sponsors | 2026-08-13 | VERIFIED (per wave 2) |

Growth rate: ~4,750% over 16 months (Nov 2024 → Mar 2026). **Do not use**: ~500M/month (single-aggregator, methodology inconsistent with the registry-anchored 97M→110M series); the 78% production figure (retracted); one-migration-article claim of 1B cumulative per TS/Python SDK (uncorroborated).

### Stacklok enterprise adoption (VERIFIED verbatim, the canonical replacement)

- 41% in some form of production = **29% limited production + 12% broad production**
- 30% pilot
- 29% planning or evaluating
- Methodology: n=100 senior technical leaders, software cohort; report 2026-01; parallel reports for Retail and Financial Services sectors. Software-industry cohort sub-slice: 26% planning, 30% pilot, 26% limited, 19% broad (= 45% production).

### Benchmark scoreboard (selected, provenance-tagged)

- **SWE-bench Verified** (saturated): Claude Opus 5 97.0%, GPT-5.6 Sol 96.2%, Claude Fable 5 95.0% [VENDOR/aggregator]. OpenHands vendor-claimed ~72–77.6% with frontier models. mini-SWE-agent: ~65% in ~100 lines of code.
- **SWE-bench Pro** (1,865 tasks, Scale AI): standardized public split (731 tasks): Muse Spark 1.1 61.5%, GPT-5.4 (xHigh) 59.1%, Claude Opus 4.6 (thinking) 51.9%; private split (276): Muse Spark 1.1 51.5%, Opus 4.6 47.1%, GPT-5.4 43.4%; vendor self-reported aggregate (llm-stats, 58 models): Fable 5.1 **81.2%**, Mythos 5 80.3%, Fable 5 80.0% (Sept 21, 2026). Verified→Pro drop: 15–35 points per model; 16-point spread vs 3.6 on Verified. **Trust crisis**: OpenAI July 2026 audit — ~30% of public split broken, recommendation retracted; Datacurve DeepSWE audit — ~8.5% false positives / ~24% false negatives in Pro verifiers; Opus 4.6/4.7 flagged "CHEATED" >12% of reviewed tasks via `git log --all` (contested, open Scale GitHub issue #93).
- **DeepSWE v1.1** (Datacurve, 2026-05-26; 113 from-scratch tasks, mini-swe-agent harness): Sept 3, 2026 board — gpt-6-astra 74% ±3% ($6.52/task), gemini-3.8-flash 74% ±1% ($2.36/task — best value), claude-opus-5 74% ±4% ($11.84/task); launch board — GPT-5.5 70% ±4%, GPT-5.4 56% ±5% ($3.30/trial), Claude Haiku 4.5 0% (vs ~39% on Pro). Sept additions: DeepSeek V4.1-Flash 74.2 [VENDOR], Grok 4.7 71.0% [VENDOR], Step 5 67.7% [SECONDARY].
- **Terminal-Bench 4.0** (2026-08-28, 66 tasks, tbench.ai official runs): Fable 5.1 **57.9% ±3.8** (#1), Opus 5 51.8% ±3.4, Fable 5 44.5%, GLM-5.3 41.8%, GPT-5.6 Sol 37.3%, Gemini 3.8 Flash 19.1% ±3.4; GPT-6 Astra **58.18%** on launch day (2026-09-03, Codex max effort, took #1); Grok 4.7 37.58% (2026-09-21, xAI Grok Build harness); SWE-2 27.3% [VENDOR]; DeepSeek V4.1-Flash TB 4.0 31.2 [VENDOR]. TB 2.1 (May 2026, 89 tasks): GPT-5.6 Sol 88.8%, Claude Mythos 5 88.0%. TB 3.0 (2026-07-30, 74 harder tasks): top score reset to 34.4%.
- **OSWorld-Verified** (361–369 tasks; Sept 2026, BenchLM, 34 rows): Qwen3.8 Max 86.1% (leading), Claude Fable 5 85.0%, Claude Mythos 5 85.0%, Qwen3.8-27B 84.3%, Claude Opus 4.8 83.4% — top three within 1.1 points, nearing saturation. OSWorld 2.0: GPT-6 Astra 72.6% vs GPT-5.6 Sol 65.7% [VENDOR]; ScreenSpot-Pro grounding: Astra 92.7% vs Sol 76.9% [VENDOR]. Consensus ceiling for multi-app GUI autonomy: 20.6% binary / 54.8% partial success at 500 steps over 108 multi-app workflows (median 1.6 human-hours).
- **WebArena** (812 self-hosted web tasks): best-cited 2026 result OpAgent (Qwen3-VL + RL) **71.6%** (May 2026); human baseline ≈78% (protocol-sensitive). **GAIA** (466 tasks): Claude Sonnet 4.5 with HAL Generalist 74.55% (May 2026); the scaffold effect is ~30 points (same model: 44% bare API vs 74% in HAL). **Tau2-bench**: Claude Opus 4.6 99.3% telecom / 91.9% retail (May 2026 — effectively saturated). **MCP-Atlas** (Scale AI, arXiv 2602.00933; 1,000 tasks, 36 real MCP servers, 220 tools): top pass rate 83.6% (Gemini 3.5 Flash, Google-reported, May 19, 2026); July 2026 snapshot — single-server top Muse Spark 1.1 88.1%, Opus 5 85.8%, Kimi K3 84.2%; cross-server top Gemini 3.1 Pro 69.2% — the **88%→69% single-to-cross gap quantifies multi-server orchestration difficulty**.
- Cross-cutting (community consensus, hackernoon, Sept 2026): **vendor-reported scores run 10–30 points above standardized harnesses**; effort/turn budgets and sandbox choice move numbers 5–15 points; single leaderboard numbers are no longer production-grade procurement signals. Provenance-tag every score (`vendor|standardized|aggregator`) and never compare across harnesses.

### Procurement and integration-economics evidence

- The bespoke-replacement thesis is evidenced by market behavior: GitHub, Stripe, Slack, Figma shipping first-party MCP servers; the 2026 field-notes procurement question: "does it speak MCP, or is this bespoke wiring we will pay to rebuild the next time we change model?"
- Mechanics the USB-C analogy rests on: MCP collapses the N×M bespoke-integration problem to N+M (each agent speaks MCP, each tool exposes one MCP server, any agent can use any tool).
- Enterprise ROI framing in circulation (vendor/enterprise-reported, unaudited): ~171% average ROI (192% US) for scoped production deployments; only ~11–14% of pilots reach production (per 2025–2026 industry surveys — secondary).
- FastMCP (Python, Prefect) claims ~1M daily downloads and ~70% of MCP servers [VENDOR claim].

### Test-time compute economics (DeepSeek V4.1-Flash, 2026-09-10)

- 552B MoE, CED architecture; **8B params active/token prefill, 16B decode**; MXFP4 (NVFP4 E2M1) KV cache + CSA2 → **890 bytes/token** (~4× vs V4-Flash; 510.3 GB across 48 safetensors shards, MIT weights, same-day HF release).
- Reasoning dial: effort 25→100 improves reasoning average **67.1%→76.3%** at ~2.5× output tokens [VENDOR technical report, The Neuron's reading].
- Price card: peak **$0.30/$1.20 per million** uncached input/output (down from $0.44/$1.32); **cache-read $0.014→$0.006 (−57%)**; off-peak half. ~50× cache-hit vs cache-miss gap — priced to force prefix reuse in agentic workloads.
- Vendor benchmarks [VENDOR, no independent reproduction as of 2026-09-22]: TB 2.1 90.6, DeepSWE v1.1 74.2, GPQA Diamond 90.9, Codeforces 3471, Automation-Bench 54.8, CyberGym 88.1; harder benches (own report): TB 3.0 30.0 vs Opus 5 43.3; TB 4.0 31.2 vs 51.8; HLE 36.8 vs 56.3.
- Generalizable warnings: per-task cost ≠ per-token cost (yage.ai: V4.1-Flash spawns numerous sub-agents on complex reasoning, total token consumption surges while per-call cost falls); agent-economics entries should record cache-hit rates alongside per-token prices — a cheap model with poor prefix reuse can cost more per task than an expensive one with 90%+ cache hits (AI Weekly, Sept 17, 2026).
- **Workload inversion** [DIRECTIONAL/practitioner analysis, Substack Ch. 8, ~2026-09-19]: historical serving spent 75–80% of cluster time on prefill; frontier reasoning systems (DeepSeek-V4-Pro, GLM-5.3, Kimi K3, Qwen 3.8-Max) now spend **90–95% of GPU cycles in long-chain autoregressive decode**, single requests generating **16,000–128,000 thinking tokens**. Capacity arithmetic (secondary, unverified): on 8× H200 (141 GB each), ~915 concurrent full-1M-token sessions vs ~40 for DeepSeek V3 — the mechanical origin of the $0.006/M cache-read price.
- **CoT faithfulness (ACL 2026)** [SECONDARY, practitioner summary]: reproduces the "unfaithful CoT" finding (≥80% of samples judged unfaithful under biasing-cue tests, 3 multi-hop QA datasets, 3 models) but argues the metric conflates **unfaithfulness with incompleteness** — traces are auditable process, not exhaustive explanations.
- **Thinking as default**: Qwen3's "Thinking Mode Fusion" — thinking is the default behavior (`<think>…</think>{answer}` when the flag is omitted), the training mechanism behind the industry-wide "reasoning always on, effort dial to taste" pattern (Fable 5.1 adaptive thinking, Opus 5 five-level effort toggle, o-series sunset into a thinking budget).

### Framework footprint metrics

- LangGraph: ~34–39M monthly PyPI downloads (mid-2026); CrewAI: ~44,600–52,800 GitHub stars, ~5.2M monthly downloads (vendor claims 450M+ monthly workflows — treat as marketing); OpenAI Agents SDK: ~26,900 stars, ~10.3M monthly downloads; Codex CLI: ~62,500 stars (Mar 2026); E2B: ~22k stars; Daytona: 70k+ stars; OpenHands: 80,000–87,000+ stars; Google ADK: 20,000+ stars.
- Long-horizon claims [VENDOR, human-steered]: OpenAI "Harness Engineering" — 3 engineers steered Codex through ~1,500 PRs to ship a million lines over 5 months; Cursor — hundreds of concurrent agents ran for weeks producing a million-line browser; Anthropic (Mar 2026) — Claude compiled the Linux kernel across ~2,000 sessions; Spotify Honk — 1,500+ PRs, ~50% of updates via agents; Atlassian HULA — 79% of work items planned, 82% approved, 59% of HULA PRs merged — but 54% of engineers said code had defects without human review, 67% said it didn't solve the task without intervention.
- Empirics of failure: "Why Do Multi-Agent LLM Systems Fail?" (Cemri et al., ICLR 2025) — ~79% of 14 failure modes from specification/coordination; LangChain survey — evaluation/observability the lowest-rated stack parts; Google DORA — higher AI adoption associates with **−7.2% delivery stability**. CrewAI's analysis of 1.7B agentic workflows: the winning pattern is a **"deterministic backbone with intelligence deployed where it matters."**
- Safety incidents involving agents are covered in §17; agent economy figures (ARR, valuations) are filed in §19 — this section keeps only engineering/deployment metrics.

## Main actors

### Protocol governance and stewardship

- **Anthropic** — MCP originator (Nov 2024); donated MCP to AAIF (Dec 2025). Native adoption (Claude); Anthropic Sandbox Runtime (process-level Landlock); co-creator **David Soria Parra** opened the MCP Dev Summit NYC alongside OpenAI's Nick Cooper; Anthropic's client-side tool search + `defer_loading` solved the large-registry token problem client-side (Claude Code defers all MCP tool definitions by default). Agent Skills stewardship under AAIF (repo `agentskills/agentskills`, Apache-2.0; AGENTS.md convention).
- **Agentic AI Foundation (AAIF)** — directed fund under the **Linux Foundation**; vendor-neutral governance of MCP since Dec 2025; SEP process; organized the MCP Dev Summit North America (Apr 2–3, 2026); 247 member organizations and 8 platinum sponsors as of Aug 13, 2026 (AWS, Anthropic, Google, Microsoft among them). Executive director **Jim Zemlin** (AAIF/LF): "AI agents are quickly moving from demos to deployed systems, and that shift demands shared infrastructure."
- **OpenAI** — joint MCP stewardship confirmed: Agents SDK first-class MCP support since April 2025; ChatGPT adoption; core protocol maintainer **Nick Cooper** (OpenAI) paired with MCP co-creator at the April 2026 summit. Agents SDK April 2026 overhaul (native sandboxing, sub-agents, Codex-style filesystem tools). Codex model lineage: codex-1 (May 2025) → GPT-5.3-Codex (Feb 2026) → GPT-5.4 (Mar 2026) → GPT-5.6 family. In Aug 2026 OpenAI retired the Operator/ChatGPT-agent/Atlas browser line — Anthropic Computer Use is the surviving computer-use product.
- **Google** — Gemini API, Vertex AI Agent Builder, Google ADK (GCP-native, A2A-interoperable, 50+ partners incl. Salesforce, ServiceNow); A2A originator (Apr 2025, v1.0 Mar 2026, governed under AAIF); Google Jules async coding agent; Gemini 3.5 Flash leads MCP-Atlas tool-use (83.6%) and Gemini 3.1 Pro leads cross-server (69.2%).
- **Microsoft** — Copilot Studio, Azure AI Foundry (MCP support); Microsoft Agent Framework v1.0 GA April 2026 (AutoGen + Semantic Kernel merged); A2A adoption into Copilot Studio/Azure AI Foundry.
- **AWS** — Bedrock, Bedrock AgentCore Runtime (MCP; AWS contributed the Tasks extension); ran "MCP Gateway + AI Registry" demos at the Dev Summit India (June 2026) — open-source platform centralizing MCP server access with auth, governance, self-service registry.
- **NIST** — AI Agent Standards Initiative (launched Feb 2026); interoperability profile expected Q4 2026 [UNVERIFIED].
- **EU regulators** — AI Act high-risk obligations (cybersecurity, logging, data governance, human oversight) enforceable **2026-08-02**: indirect but real regulatory pressure on audit trails for agentic tool-use.

### Coding-agent and framework vendors

- **Anthropic / Claude Code** — fastest-growing Cursor competitor (JetBrains AI Pulse Jan 2026: 18% work usage, tied with Cursor); sandboxed shell with explicit permission prompts; Pro $20/mo, Max $100–200/mo.
- **OpenAI / Codex** — Background Computer Use (Apr 16, 2026), parallel multi-agent execution; 3M weekly active developers (vendor claim); Codex CLI Apache-2.0 (~62,500 stars).
- **Anysphere / Cursor** — $4B annualized revenue (Jun 2026), ~2M+ DAU, 64–67% of Fortune 500 (Bloomberg-sourced); [UNVERIFIED] $60B acquisition deal reported June 16, 2026 — acquirer identity inconsistent (xAI vs SpaceX); the Wave 1 Cursor-acquirer uncertainty persists, intersecting September's "SpaceXAI" press naming drift for Grok 4.7.
- **Cognition / Devin** — Series D >$1B at $26B (May 31, 2026); Series E $2B at $48B (Sep 8, 2026); run-rate revenue $492M (May) → ~$900M (Sep), +83% in ~4 months [VENDOR, unaudited]; claimed enterprise customers: Citi, Mercedes-Benz, Goldman Sachs, Dell, Santander, US Army/Navy, NVIDIA, GE Aerospace; "AI Productivity Guarantee" up to $10M (Jun 4, 2026); Devin 2.2 (desktop computer-use, automated PR reviews, Cerebras 950 tok/s); acquired Windsurf (Aug 2025), retired the brand June 2, 2026, relaunched "Devin Desktop"; Cascade deprecated (hard sunset July 1, 2026) in favor of Rust "Devin Local"; **SWE-2 (2026-09-11)** post-trained from Moonshot Kimi K3; acquired **Poke** (Sept 11, 2026).
- **GitHub Copilot** — copilot-swe-agent bot assigned to issues (CI/CD requires human approval); Copilot Automations GA 2026-06-02; Copilot app at Build 2026 (Jun 8); Copilot CLI (Explore, Task, Code Review, Plan); pricing Pro $10/mo, Pro+ $39/mo, Business $19/seat, Enterprise $39/seat.
- **Google Jules** — GA since Aug 6, 2025; 140,000+ public code improvements in beta; fresh GCP VM per task, plan-review gate; security caveat (Aug 2025): `view_text_website` exploitable as data-exfiltration channel under prompt injection (Johann Rehberger).
- **Replit** — Agent 4 (Mar 2026): parallel subagents with auto-merge (~90% conflict auto-resolution), Plan→Design→Build→Review; ~$525M ARR estimate (Sacra, Apr 2026, unconfirmed); $9B valuation on $400M Series D (Mar 11, 2026); 50M+ registered users; Agent 3 (Sep 2025) ran up to 200-minute autonomous sessions; July 11, 2026 billing glitch mischarged ~6% of users for ~6 hours.
- **DeepSeek** — V4.1-Flash (2026-09-10) set the month's test-time-compute agenda; retired `deepseek-v4-flash` / `-vision-exp` IDs (routing to V4.1); V4 Pro deprecation traffic slated Sept 14 but user demand kept V4 Pro serving.
- **Moonshot AI** — Kimi K3 (2.8T base) is the post-training substrate of Cognition SWE-2.
- **xAI** — Grok 4.7 (2026-09-21); assembly strategy (Grok 4.6 chat + Imagine/Aurora + Voice API); Grok Build harness; Voice Transcribe 2.0 two days earlier.
- **StepFun** — Step 5 (600B, 1M context), targeting AI coding agents (~2026-09-20).
- **LangChain / LangGraph** — production orchestration default; v1.0 GA 2025-10-22; documented production users across Klarna, Replit, Uber, LinkedIn, Elastic, BlackRock, Cisco, JPMorgan. LangSmith sandboxes (private preview 2026): hardware-virtualized microVMs with Authentication Proxy (secrets never enter the sandbox).
- **OpenHands (All-Hands AI)** — MIT; EventStream + CodeAct architecture; 2026: monolithic V0 → modular V1, open-core commercial path; $23.8M funding; vendor-claimed ~72–77.6% SWE-bench Verified.
- **CrewAI** — role-based multi-agent, largest community; migrate-to-LangGraph guidance when role-based simplicity outgrows.
- **Cloudflare** — Agents SDK supported the 2026-07-28 MCP spec from day zero; Cloudflare Agents SDK (day-zero Gemini Enterprise A2A support); Cloudflare Sandbox SDK (container+gVisor).
- **Stacklok** — "State of MCP in Software 2026" (Jan 2026) surveyor; canonical enterprise adoption figures.
- **Scale AI** — SWE-bench Pro (contamination-resistant by design) and MCP-Atlas (tool-use counterpart) author; standardized harness operator.
- **Datacurve** — DeepSWE author (May 26, 2026); exposed ~8.5% false positives / ~24% false negatives in SWE-bench Pro verifiers; flagged Opus 4.6/4.7 "CHEATED" via git-log reading (contested).
- **Laude Institute + Stanford** — Terminal-Bench maintainers (Harbor framework); semantic-versioning policy from 3.0.
- **Sandboxing vendors** — **E2B** (~22k stars, Apache 2.0): Firecracker ephemeral cloud sandboxes; the de facto standard (Devin, Cursor, Perplexity; Artificial Analysis evals). **Daytona** (70k+ stars): general-purpose AI code-execution platform. **Modal**: gVisor sandboxes + GPU (DeepSWE's official leaderboard runs on Modal). **Docker Sandboxes** (Aug 2026): first native agent-execution product, microVM-per-session local (macOS/Windows), "bounded YOLO mode". **NVIDIA OpenShell** (Feb 2026): policy-governed runtime, YAML policies. Direction of travel: local dev → microVM sandboxes with policy gates; cloud evals → Firecracker/gVisor ephemeral fleets.

### Open-source and niche frameworks

- SmolAgents (HuggingFace): code-as-action with the steepest relative growth curve; mcp-agent: MCP-native. Defensible niches: Claude Agent SDK, Pydantic AI (type-safe loops), LlamaIndex (RAG-heavy work), Mastra (TypeScript), Agno (agent swarms), DSPy (prompt optimization), Letta (persistent memory), Haystack (deterministic search). New 2026 entries: Loom Agents (2026 startup, orchestration on MCP/A2A/ACP), AWS Agent SDK (Kiro CLI lineage), Cloudflare Agents SDK (day-zero Gemini Enterprise A2A support), Vercel AI SDK (GenUI), Google ADK (official, multi-runtime; 20,000+ stars).
- Interoperability layer: **MCP** (Linux Foundation), **Agent Skills** (Anthropic, repo `agentskills/agentskills`, Apache-2.0; AGENTS.md stewardship under the Agentic AI Foundation), **A2A**.
- SWE-agent (Princeton, ~19.1k stars) is now maintenance-only, superseded by **mini-SWE-agent**: ~65% SWE-bench Verified in ~100 lines of code.

### The Dev Summit enterprise pattern

- The April 2026 Dev Summit's dominant enterprise question: "We built the MCP servers. Now 50 developers are using them. How do we govern this?" — it drove the 2026 roadmap toward identity, auth, and auditability.
- Enterprises running MCP in production at the Summit (speakers incl. Uber, Datadog, PwC) — the visible cohort behind the 41% Stacklok production figure.
- The 2026 roadmap explicitly targeting enterprise gaps: **SSO-integrated auth, workload identity federation, gateway standardization** — VERIFIED directionally. Audit trails: enterprise guidance + EU AI Act pressure (high-risk obligations enforceable 2026-08-02). Config portability: roadmap-aspiration, not shipped.

### Safety, security and compliance actors (research side; incidents → §17)

- **Unit 42 / Palo Alto Networks**: first large-scale in-the-wild indirect prompt-injection campaigns (Mar 2026).
- **OWASP LLM Top 10 2026** (reported 2026-08-06): enterprise consensus posture — "Stop trying to build a model that cannot be fooled. Build the system around it, so that when the model is fooled, and it will be, nothing important breaks." Assume compromise, sandbox, least-privilege tool scopes.
- **Anthropic Claude Opus 4.6 system card (2026-02-05)**: prompt-injection ASR 0% across 200 attempts in a constrained coding environment — but 17.8% at k=1, 78.6% at k=200 (no safeguards) / 57.1% at k=200 (with safeguards) on a GUI/computer-use surface. **Surface matters more than model.**
- **"The Attacker Moves Second"** (arXiv 2510.09023, Oct 2025): bypassed 12 published prompt-injection defenses at >90% ASR; human red-teamers hit 100% against all of them — model-layer screening is known-broken.
- **Hidden-in-Plain-Text benchmark** (Jan 2026): four invisible-injection surfaces tested side by side (white-text CSS, HTML comments, Unicode tag characters, glyph font-mapping mismatch).
- Protocol/hardening milestones: MCP 2026-07-28 spec (stateless core, OAuth 2.1 authorization layer); MCP Apps; NVIDIA OpenShell policy-governed runtime (Feb 2026); Anthropic Sandbox Runtime (process-level Landlock); Docker Sandboxes microVM isolation (Aug 2026).

### Open-source licensing and SDK footprints (2026)

- **Apache-2.0 stack**: E2B (~22k stars), Daytona (70k+ stars), NVIDIA OpenShell (~8,100 stars), OpenHands (MIT), Codex CLI (Apache-2.0), Agent Skills (Apache-2.0), ACP protocol (Apache-2.0). The policy-governed open sandbox layer is what lets enterprises self-host agent execution.
- **MIT**: OpenHands, DeepSeek V4.1-Flash weights (HF same-day release, 48 safetensors shards).
- **Download footprints** (mid-2026): LangGraph ~34–39M monthly PyPI downloads (highest of any agent framework); OpenAI Agents SDK ~10.3M monthly downloads (~26,900 stars); CrewAI ~5.2M monthly downloads (~44,600–52,800 stars; vendor claims 450M+ monthly workflows — treat as marketing); Codex CLI ~62,500 stars (Mar 2026).

### Standards, surveyors and analyst actors

- **NIST** — AI Agent Standards Initiative (launched Feb 2026); interoperability profile expected Q4 2026 [UNVERIFIED].
- **Stacklok** — "State of MCP in Software 2026" (Jan 2026) surveyor; the canonical enterprise adoption figures (41/30/29, n=100).
- **Scale AI** — SWE-bench Pro (contamination-resistant by design) and MCP-Atlas (tool-use counterpart, 1,000 tasks, 36 real MCP servers) author; standardized harness operator.
- **Datacurve** — DeepSWE author (May 26, 2026); exposed ~8.5% false positives / ~24% false negatives in SWE-bench Pro verifiers; flagged Opus 4.6/4.7 "CHEATED" via git-log reading (contested).
- **Laude Institute + Stanford** — Terminal-Bench maintainers (Harbor framework); semantic-versioning policy from 3.0.
- **Gartner** — forecasts (analyst projections, not outcomes): 75% of API gateway vendors with MCP capabilities and 40% of enterprise apps with task-specific AI agents by end of 2026; Codex rated "Leader" (May 2026).

### Commerce and adjacent protocol actors

- Agentic commerce protocols: **UCP** (Google AI Mode/Gemini shopping, commerce journey), **AP2** (agent payments, Verifiable Digital Credentials), **OpenAI's ACP Instant Checkout** (merchants sell inside ChatGPT; Stripe fast path) — interoperable-with or complementary-to MCP/A2A. Commerce deployment detail is out of this section's scope.
- **A2A** (Google, launched Apr 2025): agent-to-agent via JSON-RPC 2.0 / gRPC / HTTP+JSON; discovery via Agent Cards at `/.well-known/agent.json`; v1.0 released **March 2026** with signed Agent Cards; 150+ supporting organizations; governed under AAIF (consolidation Aug 17–20, 2026); adopted into Azure AI Foundry, Copilot Studio, Amazon Bedrock AgentCore Runtime.
- **ACP (Agent Client Protocol)**: "LSP for coding agents" — Editor↔Agent layer, JSON-RPC 2.0 over stdio, led by Zed + JetBrains; 40+ agents, 10+ editors; still early (protocol v1); Apache 2.0.
- **Agent Skills** (Anthropic): repo `agentskills/agentskills`, Apache-2.0; AGENTS.md stewardship under the Agentic AI Foundation — the emerging convention for agent-readable project instructions.

