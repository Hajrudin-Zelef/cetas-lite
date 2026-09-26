---
id: ai-industry-kb-2026-wave6/16-cohere/sources-and-urls
title: "Sources and URLs"
domain: cohere
role: deep-dive
task: actor-profile
actors: ["AWS", "Anthropic", "Cohere", "Microsoft", "Mistral"]
dates: ["2025-08", "2026-05"]
keywords: ["agent", "agentic", "apache", "arr", "attribution", "aws", "bedrock", "benchmark", "claude", "cohere", "copilot", "cost"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [8158, 8190]
section: "§16. Cohere"
delta_of: ai-industry-kb-2026
sha256: 030ee475846cfe057f8a6962b1e30368a2a4929d6093fc1ba5b34c6a01f2927c
---

# Sources and URLs

1. Command A+ under Apache 2.0 removes the licensing tax on self-hosted frontier-class enterprise inference: enterprises with existing H100 capacity can run a 218B-class model with zero marginal license cost, turning the hosted-API decision into a pure operational tradeoff rather than a licensing one [DIRECTIONAL]. (S2, S6)
2. Native citation emission during inference collapses a layer of RAG scaffolding (re-ranker, post-hoc attribution) into the model itself — relevant for regulated buyers who need audit trails on generated answers [DIRECTIONAL]. (S5, S1)
3. The 24B-versus-25B active-parameter reporting and the disputed hosted pricing ($2.50/$10 mirrored vs. no published Cohere row) mean procurement comparisons should cite the Apache 2.0 self-host path as the stable fact and treat hosted Command A+ pricing as subject to change [SECONDARY]. (S1, S5, S2, S7)
4. The Cohere–Aleph Alpha combination at an ~83x ARR multiple prices the sovereignty thesis, not current revenue; whether state-adjacent capital (Schwarz Digits) counts as durable enterprise demand is the central open question [DIRECTIONAL]. (S9, S10)
5. Distribution partnerships (Carahsoft for U.S. public sector, OpenText for regulated industries) monetize North's private-deployability rather than raw model capability — a channel strategy that matches the 85%-private-deployment revenue mix claimed in secondary reporting [DIRECTIONAL]. (S13, S14, S9)
6. AWS Bedrock narrowing on-demand Cohere availability to Embed 4 and Rerank 3.5, with Command moved to Provisioned Throughput, changes the cost shape for Bedrock buyers from pay-per-token to capacity-based — a 2026 deployment fact worth surfacing to customers choosing clouds [SECONDARY]. (S24)
7. Embed v4's $0.47/M image-token price versus $0.12/M text-token price is a cost trap for image-heavy RAG pipelines: indexing screenshots of documents at scale can outpace per-quarter budgets, per Computerworld's analyst interviews [SECONDARY]. (S23)
8. The absence of Cohere prompt-caching discounts (vs. Anthropic's cached tiers) keeps Cohere cheaper on raw unit price for some workloads but removes a levers for high-cache-hit agentic loops [DIRECTIONAL]. (S18)
9. Contradictions in this file (active params 24B vs 25B, Rerank 4 context 4K vs 32K, Command A context 128K vs 256K in mirrors) reflect documentation drift across Cohere's 2025–2026 model generations — buyers should verify against the official docs catalog, not secondary mirrors [DIRECTIONAL]. (S1, S5, S8, S21, S17, S27)

10. The Command-family lineage (A → A Reasoning → A Vision → A+) shows Cohere iterating on the same 111B/112B base before the 218B MoE jump — Command A+ is an architectural break, not a scaled-up checkpoint [DIRECTIONAL]. (S32)
11. North's 2025 early-access → August 2025 GA cadence means the 2026 Carahsoft/OpenText deals are distribution plays for a mature platform, not launch news — frame them as channel expansion [DIRECTIONAL]. (S33, S35, S13, S14)
12. MCP support in North plus the Command A+ tool-use design gives Cohere a coherent enterprise agent stack (model + search/Compass + orchestration/North), which is the bundle the OpenText deal actually sells [DIRECTIONAL]. (S35, S32, S14)
13. FedRAMP High (May 2026) is the credential that makes the Carahsoft public-sector play credible; buyers should verify its scope (Second Front Systems path) rather than assuming full agency authorization [SECONDARY]. (S13b)
14. Vendor benchmark claims in this file (Command A 73 tok/s, ADI2 24.7, Command A Reasoning vs R1 0528, North vs Copilot) are all single-source vendor figures — do not cite them as independent measurements [VENDOR]. (S29, S31, S33)
15. The North internal evaluation used a Llama Index benchmark that accepts answers "at least relevant and correct" — a lenient bar disclosed in the reporting itself, limiting the weight of the claimed wins [SECONDARY]. (S33)

16. The three-model cadence (Transcribe → Command A+ → North Mini Code) is a composable open-weight enterprise agent stack under one permissive license: voice ingestion, planning/reasoning and coding can each be self-hosted inside the security perimeter without licensing negotiation — the practical expression of the sovereignty thesis [DIRECTIONAL]. (S39)
17. North Automations' governance surface (human checkpoints, audit paths, per-user token dashboards) directly answers the spend-control and oversight questions that have blocked enterprise agent rollouts — this is the feature set to verify in bake-offs, not the agent demo itself [DIRECTIONAL]. (S40)
18. Cohere's compliance language is precisely scoped ("the Cohere API" is SOC 2 Type II): regulated buyers must confirm attestation scope for VPC/on-prem North instances rather than assuming the badge transfers — a diligence line for procurement checklists [SECONDARY]. (S40)
19. Command A+'s benchmark profile (strong agentic/structured-reasoning, weak HLE broad knowledge, index 37 below Mistral Medium 3.5's 39) means it should be evaluated as an agentic-workhorse model, not a general-knowledge one — the $550B-Gartner framing is marketing, the τ²-Bench/Terminal-Bench deltas are the working figures [VENDOR]. (S43, S41)
20. QAD post-training (quantized student matching the full-precision teacher's distribution) is the technical mechanism behind the W4A4-parity claim — the claim is vendor-measured, but the mechanism is a concrete, checkable engineering choice rather than a black-box assertion [DIRECTIONAL]. (S42)
21. The late-2026 product velocity (Automations July, Parse August, Waterloo August, Carahsoft July, OpenText September, Aleph Alpha September) shows Cohere converting its 2025 platform build into channel and ecosystem revenue in 2026 — the ARR-doubling story has a product-side counterpart worth tracking [DIRECTIONAL]. (S40, S13, S14, S9)

## Sources and URLs
- https://cohere.com/blog/command-a-plus
- https://www.businesswire.com/news/home/20260520121796/en/Cohere-Releases-Command-A-An-Open-Source-Enterprise-AI-Model-Built-for-Sovereign-Critical-Infrastructure
- https://github.com/mnfst/modeldeprecations.dev/blob/HEAD/snapshots/cohere/deprecations.md
- https://github.com/sbley/claude-code-agent-test/issues/403


### New sources — expansion

