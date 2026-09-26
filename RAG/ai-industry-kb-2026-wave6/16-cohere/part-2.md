---
id: ai-industry-kb-2026-wave6/16-cohere/part-2
title: "§16. Cohere (part 2)"
domain: cohere
role: deep-dive
task: actor-profile
actors: ["Cohere", "Hugging Face", "Nvidia", "OpenRouter"]
dates: ["2025-03", "2025-08-21", "2025-09", "2026-02", "2026-03", "2026-04", "2026-04-24", "2026-05", "2026-05-20", "2026-05-21", "2026-07-30", "2026-09-16"]
keywords: ["cohere", "agentic", "apache", "benchmark", "disclosure", "distribution", "foundry", "fp8", "gpus", "inference", "license", "merger"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [7866, 7897]
section: "§16. Cohere"
delta_of: ai-industry-kb-2026
sha256: 1709b83e657fe7c1a145937eb91230b677776029f7a990b4e9992eb007cfceb7
---

# §16. Cohere (part 2)

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
