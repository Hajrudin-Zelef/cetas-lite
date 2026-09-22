---
id: ai-industry-kb-2026-wave6/16-cohere/key-dated-facts
title: "Key dated facts"
domain: cohere
role: deep-dive
task: actor-profile
actors: ["AMD", "AWS", "Anthropic", "Cohere", "CoreWeave", "DeepSeek", "Google", "Hugging Face", "Meta", "Microsoft", "Mistral", "Moonshot", "Nvidia", "OpenAI", "OpenRouter", "Oracle", "vLLM"]
dates: ["2025-01-09", "2025-01-10", "2025-03", "2025-04-15", "2025-07", "2025-08", "2025-08-06", "2025-08-21", "2025-09", "2026-02", "2026-03", "2026-04", "2026-04-04", "2026-04-24", "2026-05", "2026-05-12", "2026-05-20", "2026-05-21", "2026-06", "2026-06-09", "2026-06-25", "2026-07-27", "2026-07-30", "2026-08-06", "2026-08-27", "2026-09", "2026-09-16"]
keywords: ["agent", "agentic", "agents", "alignment", "amd", "apache", "attention", "aws", "backlog", "bedrock", "benchmark", "benchmarks"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [7814, 7971]
section: "§16. Cohere"
delta_of: ai-industry-kb-2026
sha256: 73df7370e750c789d8a19cdf3b6ddd2161e004cb43399ce26f8c327300145c64
---

# Key dated facts

## Key dated facts
### The Command A baseline — CC-BY-NC research-only
- The pre-A+ Command line — **Command A, Command R, Command R+** — shipped under **CC-BY-NC research-only** terms [SECONDARY, wave6/02 context].
- Command A+ is therefore not just a bigger model but a **license break**: the first Cohere flagship-class release under a permissive OSS license, aimed at the commercial and sovereign deployments CC-BY-NC excluded [VENDOR/SECONDARY].

### Command A+ — May 20, 2026
- **2026-05-20** — **Command A+** released [VENDOR].
- **218B total / 25B active** MoE [VENDOR] — the largest open-weight Cohere checkpoint and a sparse-activation design.
- **Apache 2.0** license [VENDOR] — vs CC-BY-NC research-only terms on the earlier Command A / R / R+ line; the first time Cohere shipped a flagship-class model under a permissive OSS license.
- The announcement's positioning: an open-source enterprise AI model **built for sovereign and critical infrastructure** [VENDOR] — language aimed at government and regulated-industry procurement.
- Framing rule: "Cohere's first Apache-2.0 model" and "first open-weight Cohere model" are vendor-supported; unqualified "first fully open-source" is not — vendor framing must be labeled [VENDOR].
- Command A+ is one of the 2026 Apache-2.0 firsts (Meta Glimmer 30B, Google Gemma 4, OpenAI gpt-oss) tracked in §21's license map.

### The quantization claim — "near lossless"
- Cohere describes the release's quantization as **"near lossless"** [VENDOR] — not "lossless." The distinction is deliberate vendor language; upgrading the adjective invents a claim.
- The claim matters because Command A+ ships quantized for deployment: the "near lossless" framing is the vendor's quality guarantee on the shippable checkpoint, and it is explicitly qualified — benchmark it before assuming parity with the full-precision training artifact [DIRECTIONAL].

### The sovereign-framing announcement
- The BusinessWire announcement's headline frames Command A+ as **"An Open-Source Enterprise AI Model Built for Sovereign, Critical Infrastructure"** [VENDOR] — the vendor's own words, and the positioning that ties the release to the FedRAMP High authorization eight days earlier.
- This is the same sovereign-infrastructure pitch NVIDIA makes with the July 24 open-weights letter ("enable sovereignty") and that Cohere pairs with permissive licensing — the 2026 pattern where open weights are sold as a sovereignty feature, not just a cost feature [DIRECTIONAL].

### Specs and deployment floor
- **48 languages** supported [VENDOR].
- **128K input** context window; **64K generation** [VENDOR].
- Minimum deployment hardware: **1× B200 or 2× H100** [VENDOR] — the stated floor for running the 218B checkpoint.

### The Aya 8B API retirement
- **2026-04-04** — **Aya Expanse 8B and Aya Vision 8B retired from the API** [SECONDARY] (modeldeprecations.dev snapshot).
- The Aya 8B API line ended ~6 weeks before Command A+ launched — the API catalog consolidated around the Command line.

### No 2026 successors — R7B, Embed 4, Rerank 4
- **No 2026 Command R7B successor** found in the corpus — the small-Command line has no 2026 refresh.
- **Embed 4 and Rerank 4 are 2025 releases** with **no 2026 successors** found.
- Unresolved in the master contradictions log: Rerank 4's context window is **32K [SECONDARY] vs 4K [COMMUNITY]** — logged, not resolved; do not assert either without the tag.

### FedRAMP High — May 12, 2026
- **2026-05-12** — Cohere received **FedRAMP High authorization** [VENDOR].
- This is the compliance substrate for the sovereign/critical-infrastructure positioning of Command A+ — one week apart, the two events read as a single enterprise push [DIRECTIONAL].

### Pricing context
- Command R+ API pricing in Sept 2026 snapshots: **$3.00/$15.00** per million input/output [SECONDARY] — the closed-API Command line's standing tier; Command A+'s open-weight release sits alongside it, not above it.
- No Command A+ API price row was found in the corpus's September price table — the wave tracks A+ as an open-weight release, with the closed Command R+ line carrying the published API rate [SECONDARY].
- In the 2026 price-war context, $3/$15 puts Command R+ at the Sonnet-5/Kimi-K3 tier — mid-market closed API, while the A+ weights compete with the Apache-2.0 open tier (Gemma 4, Glimmer 30B, gpt-oss) [DIRECTIONAL].

### Where Cohere appears elsewhere in the corpus
- **§21 (licensing map)** owns the full license-spectrum treatment: Command A+ is listed among the 2026 Apache-2.0 firsts and the permissive tier of the license map; the CC-BY-NC history of Command A/R/R+ is the "before" picture.
- **§22 (pricing tables)** owns the September 2026 price table; this section carries only the Command R+ row as context.
- The Aya 8B API retirement is tracked by **community deprecation infrastructure** (modeldeprecations.dev snapshots; the sbley/claude-code-agent-test issue #403 is cited in the sources as a community tracking reference) — not by a vendor changelog [SECONDARY].


### New verified facts — expansion

1. The Command A+ announcement is dated May 20, 2026 with secondary coverage dated May 21, 2026; the license break is explicit against named predecessors — Command A, Aya and Command R+ shipped under CC-BY-NC — making Apache 2.0 Cohere's first fully permissive license. [SECONDARY] (S1, S2)
2. Command A+ is a decoder-only sparse MoE Transformer routing across 128 experts with 8+1 active per generation step — the sparse-activation design behind the 218B-total / 25B-active footprint. [SECONDARY] (S1, S2)
3. Command A+ accepts text and image inputs and emits text, with a 128K-token input context and 64K-token output cap [SECONDARY]. (S1, S3)
4. Command A+ ships in three precision formats — BF16, FP8 and W4A4 — with the vendor claiming benchmark parity for the W4A4 build against the uncompressed model [VENDOR]. (S4, S1)
5. Cohere says the W4A4 quantization lets Command A+ run on two NVIDIA H100 GPUs (or one B200), where the previous Command family required larger footprints [VENDOR]. (S1, S5)
6. Command A+ emits source citations natively during inference rather than as a downstream RAG post-processing layer, a design pitched at enterprise audit-trail requirements in legal, compliance and healthcare workloads [SECONDARY]. (S5, S1)
7. Command A+ consolidates four Command A-series models — text, reasoning, vision and translation — into one weights package, per third-party analysis of the release [SECONDARY]. (S5, S6)
8. Open weights for Command A+ are downloadable free on Hugging Face under Apache 2.0, with no per-token rate on self-hosted inference [SECONDARY]. (S2, S5)
9. Command A+ pricing on hosted surfaces is disputed: OpenRouter, Azure AI Foundry and third-party trackers list $2.50 input / $10.00 output per million tokens — the existing Command A rate — but Cohere has not published a dedicated Command A+ row on its own pricing page [UNVERIFIED]; a later source reports Command A+ is free via Cohere's API up to rate limits with no published metered per-token price. (S2, S7)
10. In Cohere's official docs catalog the model is tracked as `command-a-plus-05-2026` (128K context, 64K max output, Chat endpoint), corroborating the 128K/64K specification [SECONDARY]. (S8, S1)
11. Cohere and Aleph Alpha signed the definitive agreement for their combination on September 16, 2026, converting an April 2026 term sheet into a real transaction — structured publicly as a "merger" in which Cohere shareholders hold approximately 90% of the combined entity and Aleph Alpha shareholders 10% [SECONDARY]. (S9, S10)
12. The combined Cohere–Aleph Alpha entity is targeting a headline valuation of about $20 billion; Cohere's standalone valuation was approximately $6.8–7 billion after extending its Series D by $100 million in September 2025 [SECONDARY]. (S10, S11)
13. Schwarz Digits, the digital-infrastructure arm of the Schwarz Group retail conglomerate (owner of STACKIT sovereign cloud), committed structured financing to the deal — reported as €500 million by one source and approximately $600 million by another — with the combined entity committed to running Schwarz Group's AI workloads on STACKIT [SECONDARY]. (S10, S11)
14. Cohere disclosed roughly $240 million in annual recurring revenue in February 2026, more than doubling from about $100 million annualized in mid-2025; the company told Upstarts in April 2026 that revenue had grown sixfold over the prior year [SECONDARY]. (S9, S11)
15. Analysts estimate gross margins near 70% with about 85% of revenue from private deployments rather than API metering, per secondary reporting that should not be read as audited company disclosure [UNVERIFIED]. (S9)
16. Aidan Gomez remains CEO of the combined company and Aleph Alpha CEO Jonas Andrulis continues with the organization; the deal still required Aleph Alpha shareholder approval at announcement [SECONDARY]. (S10, S9)
17. PitchBook lists Cohere's Series E at $2.5 billion in progress (dated April 24, 2026), while separate secondary reporting describes advanced talks for up to $3 billion — the round size is not confirmed [SECONDARY]. (S12, S11)
18. On July 30, 2026, Cohere and Carahsoft Technology Corp. announced a partnership making Carahsoft a Public Sector distributor for Cohere, exposing Cohere's models and the North agentic platform to U.S. government agencies through Carahsoft's reseller partners [SECONDARY]. (S13, S13b)
18b. Carahsoft's own announcement page (vendor-side copy) corroborates the July 30, 2026 date, the distributor role, and the five contract vehicles [VENDOR]. (S13c)
19. The Carahsoft distribution runs through NASA SEWP V, ITES-SW2, NASPO ValuePoint, TIPS and OMNIA Partners contracts [SECONDARY]. (S13 — single secondary coverage)
20. Cohere's VP of Global Public Sector Dave Ferris is the named executive on the Carahsoft partnership; the partnership positions North as enabling agencies to adopt agentic AI while aligning with government security requirements [SECONDARY]. (S13, S13b)
20b. ExecutiveBiz's independent coverage adds that Cohere's software is built to run in air-gapped and on-premises settings with no data leaving the customer environment, and that it secured FedRAMP High authorization in May 2026 through Second Front Systems [SECONDARY]. (S13b)
20c. Purchases under the Carahsoft deal run through Carahsoft's reseller network and five contract vehicles — two federal (NASA SEWP V, ITES-SW2) and three cooperative purchasing agreements (NASPO ValuePoint, TIPS, OMNIA Partners) [SECONDARY]. (S13, S13b)
21. On September 16–17, 2026, OpenText Corporation and Cohere announced a strategic partnership at the ALL IN AI conference in Montreal, pairing OpenText's data-and-context layer (unstructured, operational, transactional data) with Cohere's North platform and models as the application-and-orchestration layer for governments and regulated industries [SECONDARY]. (S14, S15)
22. OpenText CEO Ayman Antoun's statement frames the deal as giving clients what they need "to build the agentic enterprise"; Cohere's own framing is moving agentic AI from pilot to production, with customers able to run on-premises or in private, public or sovereign clouds [SECONDARY]. (S14, S15)
23. No contract value, revenue sharing, minimum purchases or customer commitments were disclosed for the OpenText partnership — the evidence supports distribution, not attributable revenue [SECONDARY]. (S16 — single secondary coverage)
24. Command A is Cohere's March 2025 flagship-generation successor to Command R+: a 111B-parameter open-weights model with 256K-token context, running on as few as two A100 or H100 GPUs, delivering roughly 150% higher throughput than Command R+ 08-2024 (about 32.8 output tokens/second, time-to-first-token near 706 ms) [SECONDARY]. (S17, S17b)
25. Command A and Command R+ share the same hosted rate card as of mid-2026: $2.50 per million input tokens / $10.00 per million output tokens; Command R+ reached that level in its 08-2024 refresh, down from $3.00/$15.00 at launch [SECONDARY]. (S17, S18)
26. Command A Reasoning, released August 21, 2025 per Cohere's changelog, is a 2025 model — it should not be presented as a 2026 launch; Hugging Face hosts it as `CohereLabs/command-a-reasoning-08-2025` [VENDOR]. (S19, S20)
27. Rerank 4 is a 2025-era release in two variants: `rerank-v4.0-pro` and `rerank-v4.0-fast`; one 2026 catalog audit lists 4K-token context per model while launch-era reporting described 32K-token context — the two figures conflict and are preserved as a contradiction [SECONDARY]. (S21, S8)
28. Rerank v4 handles English and non-English documents and semi-structured JSON reranking [COMMUNITY]. (S8 — single community source)
29. Cohere's dedicated audio transcription model appears in official docs as `cohere-transcribe-03-2026` (25MB max audio input shown in the model table), indicating a March 2026 model version [COMMUNITY]. (S8 — single community source)
30. Embed v4.0 was released April 15, 2025: Cohere's fourth-generation multimodal embedding model reaching a 65.2 MTEB score versus OpenAI's text-embedding-3-large at 64.6 [SECONDARY]. (S22, S23)
31. Embed v4 embeds interleaved text and images in the same vector space, so screenshots of PDFs, slides, figures and tables can be indexed alongside text without OCR preprocessing [SECONDARY]. (S22, S23)
32. Embed v4 supports four output dimensions — 256, 512, 1024 and 1536 — with Matryoshka-style nested representations; truncation to 1024/512/256 dims costs approximately -0.5%/-1.2%/-2.8% quality for 33%/67%/83% storage reduction [SECONDARY]. (S22, S23)
33. Embed v4 input is billed at $0.12 per million text tokens and $0.47 per million image tokens at listed rates; the text rate matches across Vercel AI Gateway, AWS Bedrock and third-party pricing surveys [SECONDARY]. (S22, S23)
34. Embed v4 is available on the Cohere API, Amazon Bedrock, Amazon SageMaker JumpStart, Azure AI Foundry and Google Cloud Vertex AI [SECONDARY]. (S23, S24)
35. Command R7B remains Cohere's budget line at $0.0375 per million input tokens / $0.15 per million output tokens, presented as 27x cheaper than competitors at comparable task quality [SECONDARY]. (S18, S26)
36. Cohere offers no prompt-caching discounts on the order of Anthropic's cached-token tiers as of mid-2026 [SECONDARY]. (S18 — single secondary coverage)
37. Cohere's Trial API key allows 1,000 free monthly calls but is explicitly not permitted for production or commercial use [SECONDARY]. (S18, S26)
38. On AWS Bedrock the Cohere catalog was pared back: Embed 4 and Rerank 3.5 remain on-demand while the Command family moved to Provisioned Throughput only — $49.50/hour (no commit), $39.60 (1-month) and $23.77 (6-month) for Command [SECONDARY]. (S24 — single secondary coverage)
39. Legacy pricing for existing customers on older Command versions persists (e.g., Command R 03-2024 at $0.50/$1.50, Command R+ 04-2024 at $3.00/$15.00, Rerank 2 at $1.00 per 1K searches) [SECONDARY]. (S18 — single secondary coverage)
40. Cohere uses an OpenAI-compatible chat format in its v2 API; the v1 custom format is deprecated, and RAG is a first-class citizen via the `documents` parameter for grounded generation [SECONDARY]. (S26, S17)
41. One community pricing mirror places Command A context at 256K tokens while another places it at 128K — preserved as a contradiction; the official docs track command-a-plus-05-2026 at 128K, but Command A (non-plus) and Command A+ are distinct models [COMMUNITY]. (S17b, S27)
42. A research memo labels Cohere "the world's leading sovereign AI company" in its own announcement copy — the phrasing comes from Cohere's press materials, not independent measurement [VENDOR]. (S10, S13)


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

### North Automations (GA July 27, 2026)
76. North Automations launched generally available on July 27, 2026: an orchestration layer for enterprise agents with granular roles and permissions, human-approval checkpoints on agent next-steps, an auditable execution path, and dashboards for active users and token consumption; vendor blog confirms "available today to all North customers". [SECONDARY] (S40, S45)
77. North Automations deploys in the customer's VPC, fully on-premises, or through Model Vault — Cohere's dedicated single-tenant hosting option; Cohere states it has no access to customer infrastructure or data in private deployments. [SECONDARY] (S40 — single secondary coverage)
78. On compliance scope, Cohere's security page states "the Cohere API" is SOC 2 Type II compliant — language scoped to the API platform, not independently confirmed to extend to every on-premises or VPC instance of North Automations. Regulated buyers should verify scope per deployment model. [SECONDARY] (S40 — single secondary coverage)
79. Cohere's North Automations launch cites Gartner analysts' framing of a $550 billion market opportunity for enterprise agentic transformation — a figure attributed to Gartner via Cohere, not an independently verified market sizing. [SECONDARY] (S40, S45)
80. A Wiz security-agent case study published June 25, 2026 demonstrates North's MCP-native architecture in a security workflow. [SECONDARY] (S40 — single secondary coverage)
81. Cohere announced a University of Waterloo certificate partnership on August 6, 2026. [SECONDARY] (S40 — single secondary coverage)
82. Cohere Parse 5 (parse-v5.0) launched August 27, 2026: a 2.3B-parameter vision-language document parser on the North-Micro-Vision-Instruct architecture, converting PDF pages, PowerPoint slides and JPEG images into Markdown with HTML tables and bounding boxes in a single model run — no separate OCR stage. API price $1.50 per 1,000 pages; Model Vault single-tenant from $2,500/month; ParseBench self-reported 79.2 (vendor-reported, averaged over 3 of 5 dimensions, trailing GPT-5.5's 84.4). [SECONDARY] (S46, S48)

### Command A+ vendor benchmarks and engineering (second-source pass)
83. Vendor benchmark package (via Cohere's announcement blog and secondary synthesis): τ²-Bench Telecom 85% (vs 37% on Command A Reasoning), Terminal-Bench Hard 25% (vs 3%), North agentic Q&A +20% accuracy, North spreadsheet analysis +32% quality, Memory Usage Quality 54% vs 39%. All vendor-reported, scored with LLM-as-a-judge on internal evals. [VENDOR] (S41, S42)
84. Multimodal reasoning (vendor-reported): MMMU 75.1% (vs Command A Vision 65.3%), MMMU Pro 63%, MathVista 80.6% (vs 73.5%), CharXiv reasoning 52.7% (vs 46.9%). [VENDOR] (S41, S42)
85. Structured reasoning (vendor-reported): GPQA Diamond 76%, AIME 2025 90%, HLE around 11% — disproportionate strength on structured tasks versus the composite index score. [VENDOR] (S41, S43)
86. Artificial Analysis Intelligence Index: Command A+ scores 37, below Mistral Medium 3.5 (39) and the closed leaders (GPT-5.5 at 60) — strongest in agentic benchmarks, weakest on broad-knowledge HLE. [SECONDARY] (S41, S43)
87. Speed engineering (vendor): at equal quantization and concurrency, Command A+ delivers up to 63% higher output tokens/sec and 17% lower TTFT than Command A Reasoning; the W4A4 build adds a further 47% speed / 13% latency cut; MoE-optimized speculative decoding adds 1.5–1.6x speedup on text and multimodal inputs. [VENDOR] (S41, S42)
88. Cohere closes the quantization gap with Quantization-Aware Distillation (QAD) in post-training: the quantized student is trained to match the full-precision teacher's output distribution using fake quantization operators forward and straight-through estimators backward — the basis of the W4A4-with-parity claim. [VENDOR] (S41, S42)
89. Command A+ expands multilingual coverage from 23 to 48 languages, with vendor-reported gains in machine translation and multilingual reasoning. [VENDOR] (S41, S42)

### Sovereign positioning (vendor release copy)
90. Cohere's Command A+ launch release frames the model as sovereign critical infrastructure: "zero hidden backdoors" (full visibility into architecture and behavior), total data sovereignty (on-premises/private cloud, no external data transmission), regulatory alignment by design, and no vendor lock-in with predictable costs; a separate Carahsoft partnership release repeats the "world's leading sovereign AI company" framing. [VENDOR] (S44, S49)
91. The same release positions the company as a security-first enterprise AI leader founded in 2019, headquartered in Toronto and San Francisco with offices in London, New York, Montreal, Paris and Seoul (single vendor release; boilerplate not independently corroborated); cumulative fundraising reported at ~$1.6B from Nvidia, AMD Ventures, Salesforce Ventures, Oracle, Cisco and institutional investors including Radical Ventures, Inovia, PSP Investments, HOOPP and BDC, plus individual AI pioneers (Hinton, Li, Abbeel, Urtasun). [VENDOR] (S44)

