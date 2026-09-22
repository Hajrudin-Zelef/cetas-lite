---
id: frontier-models-2026/01-volet-1-wave-1-research-fiche-en/overview
title: "VOLET 1 — Wave 1 Research Fiche (EN)"
domain: volet-1-wave-1-research-fiche-en
role: deep-dive
task: reference
actors: ["AWS", "Anthropic", "Google", "Microsoft", "OpenAI"]
dates: ["2025-05", "2026-02", "2026-02-03", "2026-02-05", "2026-03", "2026-06-30", "2026-07", "2027-02-04"]
keywords: ["research", "agent", "agentic", "agents", "bedrock", "benchmark", "benchmarks", "chatgpt", "claude", "cloud agent", "context window", "copilot"]
source: docs/RAG/Grands titres IA modèlesEN.md
source_anchor: ""
source_lines: [7, 95]
section: "VOLET 1 — Wave 1 Research Fiche (EN)"
sha256: 9d3e32af642105aa4f2063cf46b513ba429ff363711d1c3511cf57ab0b853a9a
---

# VOLET 1 — Wave 1 Research Fiche (EN)
## Frontier AI Models: February 2026 Release Wave
**Claude Sonnet 5 "Fennec" · GPT-5.3-Codex · Gemini 3.1 Pro**

---

## 1. Claude Sonnet 5 "Fennec" (Anthropic)

### Release & announcement
- **Official launch: February 3, 2026.** Anthropic released Claude Sonnet 5 under the internal codename "Fennec" (the fox known for speed/agility) with model identifier `claude-sonnet-5-20260203`. The launch was timed around Super Bowl week 2026 for maximum visibility.
- The community first spotted the model via a version string `claude-sonnet-5@20260203` leaked in Google Vertex AI error logs before the official announcement.
- **Historical note / discrepancy for the RAG:** a July 2026 retrospective (Medium/SSNTPL) claims the February checkpoint actually became Sonnet 4.6 and that the "genuine" Sonnet 5 shipped **June 30, 2026** with SWE-bench Pro 63.2% and OSWorld-Verified 81.2%. Most Feb–Apr 2026 sources (WaveSpeed, Vertu, AI First Founders) treat Feb 3, 2026 as the Sonnet 5 launch with 82.1% SWE-Bench Verified. Both narratives are documented; the Feb-2026 narrative is the primary one for this wave.

### Architecture & technical specs
- **Architecture:** not officially disclosed by Anthropic (transformer-based frontier LLM). Third-party analyses describe it as using "distilled reasoning" — compressing flagship-class intelligence into an efficient inference engine to solve the latency–intelligence paradox — plus "zero-latency background reasoning" (reasoning without visible thinking blocks).
- **Context window:** 1,000,000 tokens (5× larger than Opus 4.5; via API). Max output ~128,000 tokens.
- **Infrastructure:** first Anthropic model optimized for Google's "Antigravity" layer on TPUv6 architecture, co-developed with Google — speculative decoding (10–20 tokens in parallel), massive throughput (1M context at the latency Sonnet 3.5 needed for 200K), "warm" contextual persistence across days, and ~20–30% faster inference than prior generations.
- **Modalities:** text, images, PDF, code (reported).
- **Security-aware generation:** "Refusal with Explanation" logic — when it detects a security vulnerability being introduced (SQL injection, XSS, etc.), it identifies the vulnerability, explains the risk, and suggests a secure alternative, proceeding only on explicit user confirmation.

### Key benchmarks (reported values, Feb–Apr 2026 sources)
| Benchmark | Score |
|---|---|
| **SWE-Bench Verified** | **82.1%** — first model ever above 80% |
| SWE-Bench vs Opus 4.5 | 82.1% vs 80.9% (Opus 4.5) vs ~78% (GPT-5 est.) |
| OSWorld (computer use) | >85% (leaked); 81.2% (June-build reports) |
| Terminal-Bench 2.1 | 80.4% (June-build reports) |
| SWE-bench Pro | 63.2% (June-build reports) |
| GDPval-AA v2 | 1,618 (edges Opus 4.8's 1,615) |

### API pricing & availability
- **$3.00 / 1M input tokens, $15.00 / 1M output tokens** (~50–80% cheaper than Opus 4.5 depending on the report).
- Intro pricing on the June build: $2/$10 per 1M through Aug 31, 2026, then $3/$15.
- Included in Claude Pro ($20/month); default model on Free and Pro plans (June build).
- Channels: Anthropic API, Amazon Bedrock, Google Vertex AI, Microsoft Foundry, Claude Code.
- Model ID: `claude-sonnet-5-20260203` / `claude-sonnet-5@20260203`.

### Agentic / coding use cases
- **"Dev Team" mode** in the Claude Code CLI: Sonnet 5 acts as a multi-agent orchestrator — a Manager agent decomposes a high-level goal, spawns specialized sub-agents (Backend, QA, Infrastructure/Researcher), they work in parallel on different files (e.g., Backend writes an API route while QA generates unit tests), and the Manager reconciles conflicts before presenting the final PR.
- Self-correcting code execution with built-in terminal; full-feature builds from briefs autonomously; bug-report → write → test → verify patch loops on first try ("human parity" milestone for junior-to-mid-level dev work).
- Solo-founder/agency use: entire feature pipelines (architecture, implementation, debugging) built in AI-first loops.

### What distinguished it in 2026
- The **"Opus-killer"**: flagship-class coding (beating Opus 4.5) at mid-tier price — the best price/performance ratio in frontier models at launch, and the catalyst for the February 2026 agentic-coding arms race (OpenAI launched GPT-5.3-Codex minutes later).

---

## 2. GPT-5.3-Codex (OpenAI)

### Release & announcement
- **Released February 5, 2026** — GA of the code-specialized GPT-5.3 family. OpenAI described it as "the most capable agentic coding model to date," combining GPT-5.2-Codex's coding frontier with GPT-5.2's reasoning/professional knowledge.
- Launched **minutes after Anthropic's own agentic coding announcement** (both were planned for 10 a.m. PST; Anthropic moved theirs up 15 minutes) — the defining Feb 2026 "arms race" moment.
- **March 2026:** designated **LTS (Long-Term Support)** — the first time OpenAI/GitHub guaranteed a model version remains available for a fixed period, **through February 4, 2027**, giving enterprises deployment stability.
- Self-referential milestone: early versions of GPT-5.3-Codex were used during its own training, deployment, and test diagnosis — the first OpenAI model instrumental in creating itself (recursive self-improvement loop).

### Architecture & technical specs
- **Architecture:** not officially disclosed; analyses point to advanced Mixture-of-Experts / sparse activation optimized for programming logic.
- **Context window:** 400,000 tokens; max output 128,000 tokens.
- **Speed:** **25% faster** than GPT-5.2-Codex on agentic tasks, with fewer tokens used, thanks to inference-stack improvements.
- Knowledge work hybrid: coding specialization + reasoning/professional knowledge of GPT-5.2.

### Key benchmarks (reported values)
| Benchmark | GPT-5.3-Codex | Predecessors |
|---|---|---|
| **SWE-Bench Pro (Public)** | **56.8%** (SOTA) | 56.4% (5.2-Codex), 55.6% (5.2) |
| **Terminal-Bench 2.0** | **77.3%** | 64.0% (5.2-Codex), 62.2% (5.2) |
| **OSWorld-Verified** (agentic computer use) | **64.7%** | 38.2% (5.2-Codex) |
| SWE-Lancer (IC-Diamond) | 81.4% | — |
| Cybersecurity CTFs | 77.6% | — |
| LiveBench | 72.8% | — |
- Contamination-resistant design: SWE-Bench Pro spans 4 programming languages.

### API pricing & availability
- **$1.75 / 1M input tokens, $14.00 / 1M output tokens** (confirmed via llm-stats.com / official pricing pages, Aug 2026 recheck).
- Note: early Feb 2026 coverage said "pricing not separately announced from the GPT-5.3 family"; $1.75/$14 figures come from later official pricing pages.
- **Access (at launch):** ChatGPT paid tiers (Plus, Pro, Business, Enterprise, Edu) via Codex app, Codex CLI (`npm i -g @openai/codex`), IDE extensions, web; **API access delayed** because the model was flagged **"High Capability" in cybersecurity** under OpenAI's Preparedness Framework (dual-use exploit-discovery concern) — a precautionary access control.
- GitHub Copilot: GA February 2026, 25% faster on agentic tasks.
- Three products share the "Codex" name: v1 = the 2021 API model (deprecated 2023), v2 = the CLI agent (Apr 2025–), v3 = the cloud agent (May 2025–).

### Agentic / coding use cases
- Interactive agentic coding paradigm: **real-time steering** — developers intervene mid-execution (e.g., switching protocol requirements) without restarting; frequent progress updates; plan → tool-use → iterate loops.
- Beyond code: debugging, test creation, documentation, refactoring, deployment support, monitoring, PRD writing, user research — near every professional computing task.
- Long-horizon projects: builds functional complex games/apps from scratch over multi-day runs.

### What distinguished it in 2026
- **Terminal-first agentic specialist**: it dominated terminal-based workflows (Terminal-Bench 2.0 77.3% vs rivals) at 2.9× cheaper input pricing than Opus-class models — the price-efficiency leader for terminal-heavy engineering, with enterprise-grade LTS stability.

---

