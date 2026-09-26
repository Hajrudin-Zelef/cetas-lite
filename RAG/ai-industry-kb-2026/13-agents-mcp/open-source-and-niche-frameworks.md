---
id: ai-industry-kb-2026/13-agents-mcp/open-source-and-niche-frameworks
title: "Open-source and niche frameworks"
domain: agents-mcp
role: deep-dive
task: agents
actors: ["AWS", "Anthropic", "Cerebras", "DeepSeek", "EU", "Google", "Hugging Face", "Meta", "Microsoft", "Moonshot", "Nvidia", "OpenAI", "Perplexity", "SpaceX", "StepFun", "United States", "xAI"]
dates: ["2025-10-22", "2026-04", "2026-05-26", "2026-05-31", "2026-06-02", "2026-06-16", "2026-07-01", "2026-07-11", "2026-07-28", "2026-08-02", "2026-09-10", "2026-09-11", "2026-09-20", "2026-09-21"]
keywords: ["acquisition", "agent", "agentic", "agents", "apache", "arr", "aws", "claude", "compute", "copilot", "deepseek", "funding"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [6729, 6763]
section: "13. Agents & MCP"
sha256: d2cf7482f9796c55e2d659cac2ed6df9c25a468237d16440ff56c131c186effd
---

# Open-source and niche frameworks

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

