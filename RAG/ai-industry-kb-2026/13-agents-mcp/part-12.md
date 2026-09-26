---
id: ai-industry-kb-2026/13-agents-mcp/part-12
title: "13. Agents & MCP (part 12)"
domain: agents-mcp
role: deep-dive
task: agents
actors: ["Alibaba", "CISA", "DeepSeek", "EU", "StepFun", "xAI"]
dates: ["2025-11-25", "2026-03", "2026-03-25", "2026-04", "2026-04-10", "2026-05", "2026-06", "2026-06-10", "2026-07-18", "2026-07-28", "2026-09-10", "2026-09-11"]
keywords: ["agent", "agents", "mcp", "acquisition", "agentic", "backlog", "benchmark", "compute", "copilot", "cost", "deepseek", "gpu"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [6894, 6931]
section: "13. Agents & MCP"
sha256: ae1da3f6fa6a89b6ec438f2d38c8a7ca914b4753597960d7b48cc0bebcfd40ab
---

# 13. Agents & MCP (part 12)

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

