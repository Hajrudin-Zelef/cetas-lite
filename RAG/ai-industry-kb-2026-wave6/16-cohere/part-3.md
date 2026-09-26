---
id: ai-industry-kb-2026-wave6/16-cohere/part-3
title: "§16. Cohere (part 3)"
domain: cohere
role: deep-dive
task: actor-profile
actors: ["AWS", "Anthropic", "Cohere", "Google", "OpenAI"]
dates: ["2025-04-15"]
keywords: ["cohere", "aws", "bedrock", "embedding", "foundry", "multimodal", "pricing", "research", "throughput"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [7898, 7912]
section: "§16. Cohere"
delta_of: ai-industry-kb-2026
sha256: 8d9cce954b922d031f79f2d48fe66f26de83988dd9d14135cce153ba65926015
---

# §16. Cohere (part 3)

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


