---
id: ai-industry-kb-2026-wave6/16-cohere/three-model-apache-2-0-cadence-marchjune-2026
title: "Three-model Apache 2.0 cadence (March–June 2026)"
domain: cohere
role: deep-dive
task: actor-profile
actors: ["Cohere", "CoreWeave", "DeepSeek", "Google", "Hugging Face", "Meta", "Microsoft", "Mistral", "Nvidia", "OpenAI", "OpenRouter", "vLLM"]
dates: ["2025-01-09", "2025-01-10", "2025-07", "2025-08", "2025-08-06", "2026-03", "2026-05", "2026-06", "2026-06-09"]
keywords: ["apache", "agent", "agentic", "agents", "attention", "backlog", "benchmark", "benchmarks", "cohere", "consumer", "copilot", "deepseek"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [7913, 7949]
section: "§16. Cohere"
delta_of: ai-industry-kb-2026
sha256: 201fa6d65c14c4a125d449392be03b324261f17d38c25433b982d7b30e258d63
---

# Three-model Apache 2.0 cadence (March–June 2026)

43. Command A was released March 13–16, 2025 as a 111B-parameter model with 256K-token context and support for 23 languages, targeting enterprise deployments on just two GPUs [SECONDARY]. (S28, S29)
44. Command A employs an optimized transformer design with three layers of sliding-window attention (4096-token window) plus a global-attention layer, balancing local context modeling with full-sequence interactions [SECONDARY]. (S28 — single secondary coverage)
45. At launch Cohere claimed 73 tokens/sec streaming at 100K-token context for Command A versus 38/sec for GPT-4o and 32/sec for DeepSeek-V3 — vendor figures, not independent measurements [VENDOR]. (S29 — single secondary coverage)
46. Cohere reported an ADI2 dialect-consistency score of 24.7 for Command A versus 15.9 (GPT-4o) and 15.7 (DeepSeek-V3) — vendor benchmark, single-source claim [VENDOR]. (S29 — single secondary coverage)
47. The Hugging Face research release `CohereLabs/c4ai-command-a-03-2025` is licensed CC-BY-NC with Cohere Labs Acceptable Use Policy; the card notes the model supports 256K but is configured in HF for 128K, adjustable in configuration [VENDOR]. (S30, S53)
48. Command A launch pricing was $2.50 per million input / $10.00 per million output, with private and on-prem deployments available on request [SECONDARY]. (S29, S18)
49. Command A Reasoning introduces a token-budget feature letting users specify how much reasoning to allocate per input; the HF release exposes reasoning on/off via a simple parameter so the same model runs in "reasoning mode" or low-latency mode [VENDOR]. (S31, S30b)
50. Cohere's vendor benchmarks claim Command A Reasoning outpaces DeepSeek-R1 0528, gpt-oss-120b and Mistral Magistral Medium on enterprise reasoning tasks — vendor figures, not independent [VENDOR]. (S31 — single vendor source)
51. A community consolidation table distinguishes four Command-family generations: Command A (2025) 111B dense, no reasoning, 23 languages; Command A Reasoning 111B; Command A Vision 112B multimodal, 6 languages; Command A+ 218B/25B, 48 languages [COMMUNITY]. (S32 — single community source)
52. Command A+ Hugging Face checkpoints use model IDs of the form `CohereLabs/command-a-plus-05-2026-w4a4` (and `-bf16`); vLLM serving requires vllm>=0.21.0 plus `cohere_melody` tool-call/reasoning parsers [COMMUNITY]. (S32 — single community source)
53. The Cohere dashboard API model name for Command A+ is `command-a-plus-05-2026`, matching the docs catalog entry [COMMUNITY]. (S32, S8)
54. North was introduced January 9, 2025 as a "secure AI workplace platform" in early access — a 2025 platform, not a 2026 launch [SECONDARY]. (S33, S34)
55. North reached general availability on August 6, 2025, about seven months after early access, with an initial installed base including Dell Technologies and the Royal Bank of Canada [SECONDARY]. (S35 — single secondary coverage)
56. North is a low-code platform for building and deploying agents across business functions (HR, finance, customer support, IT), positioned against Microsoft autonomous agents, Google Vertex AI agents and Salesforce Agentforce [SECONDARY]. (S34, S33)
57. North is powered by the Command models plus Cohere's Compass multimodal search-and-discovery framework, which combines Embed/Rerank retrieval models, document parsing (PDF, PPT, DOCX, XLSX) and a managed index [SECONDARY]. (S33, S34)
58. In an internal vendor evaluation, Cohere compared North against Microsoft Copilot and Google Vertex AI Agent Builder across finance, HR, customer support and IT tasks, claiming North outperformed both on all four — vendor evaluation, not independent [VENDOR]. (S33 — single vendor source)
59. North ships with a security system that can absorb an enterprise's identity-and-access-management rules and deploys in private cloud or on-premises [SECONDARY]. (S34 — single secondary coverage)
60. North connects to cloud applications through prepackaged integrations and links to custom software via MCP (Model Context Protocol) [SECONDARY]. (S35 — single secondary coverage)
61. Cohere partnered with CoreWeave to deploy one of the industry's first NVIDIA GB200 NVL72 clusters in production for training North-era models; Cohere's VP of Engineering credits the deployment with accelerating North development [VENDOR]. (S36, S50)
62. The CoreWeave case study cites 3x faster AI training for Cohere, plus ARM64 validation through GH200 nodes and CoreWeave AI Object Storage throughput [VENDOR]. (S36, S50)
63. Cohere was valued at $5.5 billion following a $500 million round in July 2025 — consistent with the $6.8B figure cited at Command A Reasoning's August 2025 launch after the Series D extension [SECONDARY]. (S33, S31)
64. A TechTarget report (January 10, 2025) confirms the North early-access launch as Cohere's move to compete with Microsoft and Google in the enterprise workspace [SECONDARY]. (S37 — single secondary coverage)
65. A community-built manufacturing digital-twin showcase for North demonstrates five factory zones with specialized agent types (SOP Search, Work Order, OEE Dashboard, Demand Forecast, Root Cause Analyzer) and claimed before/after KPIs (onboarding weeks→days, downtime −50–65%) — community demo claims, not verified outcomes [COMMUNITY]. (S38 — single community source)
66. The VentureBeat Command A+ coverage confirms the OpenRouter/Azure-mirrored $2.50/$10 rate structure as third-party-listed while noting Cohere has not pushed a dedicated A+ row to its public pricing page [SECONDARY]. (S4 — single secondary coverage)
67. Command A+ is described as designed for sovereign/air-gapped enterprise and government deployments — the same sovereign positioning underpinning the Aleph Alpha combination and the Carahsoft/OpenText distribution deals [SECONDARY]. (S1, S10)
68. The wealthengine analysis cautions that Cohere "does not publish audited revenue, cash-burn or backlog figures" that would let public investors verify the $20B figure [SECONDARY]. (S16 — single secondary coverage)
69. The OpenText partnership is the second major 2026 distribution deal after Carahsoft, both aimed at regulated/government buyers rather than consumer AI [SECONDARY]. (S13, S14)
70. Command A Vision (112B, 6 languages, limited tool use) is the vision sibling in the Command A generation per the community consolidation table — not to be confused with Command A+ multimodal input [COMMUNITY]. (S32, S54)
71. Cohere's own docs describe Embed v4 as helping organizations "securely retrieve their multimodal data to build agentic AI applications" — vendor positioning [VENDOR]. (S22 — single vendor source)
72. Analyst commentary quoted by Computerworld warns Embed 4 image-token pricing ($0.47/M) could "outpace quarter-by-quarter budgets if usage scales" for image-heavy workloads, and notes Cohere lacks the developer ecosystem of OpenAI/Meta/Google [SECONDARY]. (S23 — single secondary coverage)


### Three-model Apache 2.0 cadence (March–June 2026)
73. Cohere shipped three models under Apache 2.0 in six months: Transcribe (March 2026), Command A+ (May 2026) and North Mini Code (June 9, 2026) — a deliberately staggered stack where Transcribe ingests voice into structured text, Command A+ serves as the reasoning/planning core and North Mini Code acts as the coding sub-agent. All three run inside the customer's security perimeter. [SECONDARY] (S39, S47)
74. North Mini Code, released June 9, 2026, is an agentic coding model specialized for navigating and modifying large codebases — the coding leg of the three-model stack; MarkTechPost confirms 30B total / 3B active, 256K context (64K max generation), Apache 2.0 on Hugging Face, 1xH100 FP8 minimum. [SECONDARY] (S39, S47)
75. The three-model Apache 2.0 cadence gives agent developers a vertically integrated open-weight stack (speech → reasoning/planning → code) with no vendor lock-in, per secondary analysis of Cohere's announcements. [SECONDARY] (S39, S47)

