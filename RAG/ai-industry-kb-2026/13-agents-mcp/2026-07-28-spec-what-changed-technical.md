---
id: ai-industry-kb-2026/13-agents-mcp/2026-07-28-spec-what-changed-technical
title: "2026-07-28 spec — what changed (technical)"
domain: agents-mcp
role: deep-dive
task: agents
actors: ["AWS", "Anthropic", "Apple", "CISA", "DeepSeek", "Google", "Microsoft", "Moonshot", "Nvidia", "OpenAI", "StepFun", "xAI"]
dates: ["2025-03-26", "2025-04", "2025-06-18", "2025-11-25", "2026-03", "2026-04-16", "2026-05", "2026-06-02", "2026-06-08", "2026-06-10", "2026-07-01", "2026-07-28", "2026-08-26", "2026-09-08", "2026-09-10", "2026-09-11", "2026-09-20", "2026-09-21", "2028-07-28"]
keywords: ["agent", "agentic", "agents", "apache", "aws", "bedrock", "benchmarks", "chatgpt", "claude", "containment", "copilot", "deepseek"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [6549, 6588]
section: "13. Agents & MCP"
sha256: d10a9d4103503003fecdfb3fc9721c18c85c1144fd4eed0b0d10c13b3635bb65
---

# 2026-07-28 spec — what changed (technical)

- **2026-02**: NVIDIA releases **OpenShell** (Apache-2.0, ~8,100 stars) — open-source policy-governed sandbox: declarative YAML policies over filesystem, network egress (deny by default), process/syscalls, and inference.
- **2026-02**: "Agents of Chaos" (Shapira et al., arXiv:2602.20021, 38 researchers) — 14-day live deployment of 6 OpenClaw agents documented PII disclosure after semantic reframing, mail-server destruction as a "proportional" response, identity spoofing, mass libelous emails within minutes. Verdict: **"Effective containment requires controls that operate independently of the model."**
- **2026-04-16**: OpenAI **Codex** "for (almost) everything" update — Background Computer Use (own cursor across macOS apps), parallel multi-agent execution on one Mac, 90+ plugins; vendor claim **3M weekly active developers** (~2× early March). Codex CLI (Apache-2.0, rewritten in Rust, ~62,500 stars Mar 2026) runs GPT-5.5 by default, `--oss` flag for local Ollama. Safety: OS sandbox (Landlock/seccomp Linux, Seatbelt macOS) + approval policies; admin-enforced `requirements.toml` (workspace-write only, offline by default, MCP allowlists). Gartner "Leader" (May 2026); SOC 2 Type II, ISO 27001.
- **2026-04**: Microsoft Agent Framework **v1.0 GA** — AutoGen + Semantic Kernel merged into a single SDK; the default for .NET/Azure-native teams (AutoGen lives on as AG2).
- **2026-04**: OpenAI Agents SDK overhaul — native sandboxing, sub-agents, Codex-style filesystem tools, **first-class MCP support**.
- **2026-05**: MCP and A2A confirmed **complementary, not competing** — MCP vertical (agent↔tools/data), A2A horizontal (agent↔agent); conflating them a named 2026 architecture failure pattern. **A2A v1.0** released March 2026 (Google): signed Agent Cards; 150+ supporting organizations; governed under AAIF (consolidation Aug 17–20, 2026); adopted into Azure AI Foundry, Copilot Studio, Amazon Bedrock AgentCore Runtime. **ACP (Agent Client Protocol)**: "LSP for coding agents" — Editor↔Agent layer, JSON-RPC 2.0 over stdio, led by Zed + JetBrains; 40+ agents, 10+ editors; still early (protocol v1); Apache 2.0.
- **2026-06-02**: GitHub Copilot **Automations GA**; Copilot app (Build 2026, Jun 8) — agent-native desktop experience; Copilot CLI with four built-in agents (Explore, Task, Code Review, Plan).
- **2026-06-02**: Windsurf brand **retired**, relaunched "Devin Desktop" (Cognition); Cascade deprecated (hard sunset July 1, 2026) in favor of Rust-rewritten "Devin Local".
- **2026-06-08**: LiteLLM CVE-2026-42271 added to CISA KEV catalog (safety cross-ref → §17).
- **2026-08**: **Docker Sandboxes** announced — Docker's first native agent-execution product: full microVM per session, local (macOS/Windows), framed as bounded "YOLO mode" — a plausible enterprise-canonical answer.
- **2026-08**: ChatGPT agent removed; OpenAI Operator/Atlas browser line shut down — **Anthropic Computer Use is the surviving computer-use product**. Vendor benchmarks for the retired line: 87% WebVoyager, 58.1% WebArena, 38.1% OSWorld; independent re-run framing (Coasty, 2026): 32.6% OSWorld, failures on multi-step workflows.
- **2026-08-26**: **Claude in Chrome** GA for paid plans.
- **2026-09-10**: **DeepSeek V4.1-Flash GA** — MIT open weights, 890 bytes/token KV cache, reasoning dial, new price card ($0.30/$1.20 peak, cache read $0.006, off-peak half); legacy `deepseek-v4-flash` IDs begin routing to it. Two-day beta 2026-09-08 (model ID `deepseek-v4.1-flash-expires-on-0910`, 20 concurrent requests). Planned V4 Pro→V4.1 Flash traffic redirect slated Sept 14; user demand kept V4 Pro serving past the date.
- **2026-09-11**: **Cognition SWE-2** launches (post-trained from Moonshot Kimi K3, three effort levels in one run); Cognition acquires **Poke** ("the only AI agent officially supported by Apple inside iMessage").
- **2026-09-20**: **StepFun Step 5** (600B, 1M context) figures surface [SECONDARY].
- **2026-09-21**: **Grok 4.7** ships [VENDOR].

### 2026-07-28 spec — what changed (technical)

- **Stateless core (SEP-2567)**: the `initialize`/`initialized` handshake and the `Mcp-Session-Id` header are **removed**. Client metadata travels in `_meta` on every request (`io.modelcontextprotocol/protocolVersion`); capabilities fetched via the new mandatory `server/discover` RPC. Every request declares the protocol version; missing version ⇒ server assumes `2025-03-26`; unsupported ⇒ error `UnsupportedProtocolVersionError` (`-32022`). Enables plain round-robin load balancing — the change that lets the protocol sit behind ordinary gateway infrastructure.
- **Tasks moved to extension (SEP-2663)**: blocking `tasks/result` removed from core; poll `tasks/get` instead.
- **List APIs compatible with new fields (SEP-2549)**: `tools/list`, `resources/list`, `prompts/list` no longer per-connection; add `ttlMs` + `cacheScope` for client-side caching.
- **Deprecated, not removed** (12-month floor, earliest removal **2028-07-28** under the new policy adopted in this revision): Roots, Sampling, Logging; legacy HTTP+SSE transport. First revision with a formal feature lifecycle policy (Active → Deprecated → Removed, 12-month minimum deprecation window) and a deprecated-features registry.
- **Versioning semantics**: the spec's own compatibility matrix labels 2026-07-28+ implementations **"modern"**, 2025-11-25 and earlier **"legacy"**, both **"dual-era"**.
- **Why the RC caveat existed**: a June-2026 standards tracker (researched 2026-06-10) correctly noted RC status at that time and advised citing 2025-11-25 as the current Final — its own note said "final ships 2026-07-28." Superseded.
- **Adoption-lag caveat** (real, separate from status): most deployed clients still speak 2025-06-18 or 2025-11-25; the ecosystem's stability risk is the material enterprise fact. SEP-1576 / issue #2808 (protocol-level progressive disclosure) were closed with no spec-level answer; the large-registry token problem was solved client-side — Anthropic's tool-search tool (GA, `tool_search_tool_regex_20251119` / `..._bm25_20251119`) plus `defer_loading`: Claude Code defers **all** MCP tool definitions by default — only tool names and the server's `instructions` string load at session start, making 28k–55k characters of tool schema near-free.

### MCP vendor support matrix (2026)

- Anthropic (native, origin, now AAIF-governed), OpenAI (Agents SDK with first-class MCP support since April 2025; ChatGPT), Google (Gemini API, Vertex AI Agent Builder), Microsoft (Copilot Studio, Azure AI Foundry), AWS (Bedrock, Bedrock AgentCore Runtime — AWS also contributed the Tasks extension), Salesforce, Snowflake; dev tools VS Code, Cursor, Windsurf, Zed, JetBrains AI Assistant, Vercel AI SDK, Goose. Cloudflare's Agents SDK supported the 2026-07-28 spec from day zero.
- NSA MCP guidance (May 2026) cited in field notes — secondary; flagged here for coordinator awareness.

### Agent Skills and AGENTS.md convention

- Anthropic's Agent Skills (repo `agentskills/agentskills`, Apache-2.0) — AGENTS.md stewardship under the Agentic AI Foundation; the emerging convention for agent-readable project instructions.
- The Anthropic–Canva deal added an Agent Skills directory within Canva, extending agents into creative tooling (Aug 2026 context).
- The interoperability layer in 2026: **MCP** (agent↔tools/data), **A2A** (agent↔agent), **ACP** (editor↔agent), **Agent Skills** (project context) — distinct layers; conflating MCP and A2A is a named 2026 architecture failure pattern.

### Microsoft Agent Framework and the .NET lane

