---
id: ai-industry-kb-2026/13-agents-mcp/open-source-licensing-and-sdk-footprints-2026
title: "Open-source licensing and SDK footprints (2026)"
domain: agents-mcp
role: deep-dive
task: agents
actors: ["AWS", "Anthropic", "DeepSeek", "Google", "Meta", "Microsoft", "Nvidia", "OpenAI", "Stripe"]
dates: ["2026-02-05", "2026-03", "2026-05", "2026-05-26", "2026-07-28", "2026-08-06"]
keywords: ["agent", "agentic", "agents", "apache", "bedrock", "benchmark", "chatgpt", "claude", "copilot", "deepseek", "foundry", "gemini"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [6764, 6792]
section: "13. Agents & MCP"
sha256: f48bc3adb304bb1cda94a86e919b7aeb00fe17c046850133f774aff1c4ad8d90
---

# Open-source licensing and SDK footprints (2026)

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

