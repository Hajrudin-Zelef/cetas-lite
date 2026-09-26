---
id: ai-industry-kb-2026/13-agents-mcp/framework-footprint-metrics
title: "Framework footprint metrics"
domain: agents-mcp
role: deep-dive
task: agents
actors: ["AWS", "Alibaba", "Anthropic", "DeepSeek", "EU", "Google", "Microsoft", "Moonshot", "OpenAI", "Z.ai"]
dates: ["2025-04", "2025-05", "2026-04", "2026-06", "2026-08-02", "2026-09-19", "2026-09-22"]
keywords: ["agent", "agentic", "agents", "apache", "arr", "aws", "bedrock", "benchmarks", "chatgpt", "claude", "copilot", "cost"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [6698, 6728]
section: "13. Agents & MCP"
sha256: 2748213e70b06fe1b8c66c954711ac4de94ada14c612b212f3a8a7a60f2b138f
---

# Framework footprint metrics

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

