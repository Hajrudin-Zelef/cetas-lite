---
id: ai-industry-kb-2026-wave6/16-cohere/main-actors
title: "Main actors"
domain: cohere
role: deep-dive
task: actor-profile
actors: ["Cohere", "CoreWeave", "Meta", "Mistral", "OpenAI"]
dates: ["2025-01-09", "2025-04-15", "2025-07", "2025-08-06", "2025-08-21", "2025-09", "2026-02", "2026-03", "2026-04", "2026-04-04", "2026-05", "2026-05-12", "2026-05-20", "2026-05-24", "2026-06-09", "2026-06-25", "2026-07-27", "2026-07-30", "2026-08-06", "2026-08-10", "2026-08-27", "2026-09", "2026-09-16"]
keywords: ["agent", "agentic", "apache", "arr", "cohere", "distillation", "distribution", "embedding", "fp8", "governance", "gpu", "latency"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [8065, 8144]
section: "§16. Cohere"
delta_of: ai-industry-kb-2026
sha256: e46041ab80926521b24395b659ad1800ccf72bb77063cb80d38b0e3e182e8033
---

# Main actors

| North Mini Code | Released 2026-06-09; 30B total / 3B active MoE, 256K context (64K max gen), Apache 2.0, 1xH100 FP8 min | S39, S47 | [SECONDARY] |
| Three-model Apache 2.0 cadence | Transcribe (Mar), A+ (May), Mini Code (Jun) | S39 (single source) | [SECONDARY] |
| North Automations | GA 2026-07-27, orchestration + governance layer | S40 (single source) | [SECONDARY] |
| North Automations governance | roles/permissions, human checkpoints, audit path, usage dashboards | S40 (single source) | [SECONDARY] |
| Cohere Parse 5 (parse-v5.0) | Launched 2026-08-27; 2.3B VLM, PDF/PPT/JPEG->Markdown, $1.50/1k pages, Model Vault from $2,500/mo, ParseBench 79.2 (vendor-reported) | S46, S48 | [SECONDARY] |
| Waterloo certificate partnership | Announced 2026-08-06 | S40 (single source) | [SECONDARY] |
| Wiz security-agent case study | Published 2026-06-25 | S40 (single source) | [SECONDARY] |
| Gartner market-opportunity figure (via Cohere) | $550B | S40 (single source) | [SECONDARY] |
| Command A+ τ²-Bench Telecom | 85% (vs 37% Command A Reasoning) | S41, S42 | [VENDOR] |
| Command A+ Terminal-Bench Hard | 25% (vs 3% Command A Reasoning) | S41, S42 | [VENDOR] |
| Command A+ North agentic Q&A | +20% accuracy vs A Reasoning | S41, S42 | [VENDOR] |
| Command A+ North spreadsheet analysis | +32% quality vs A Reasoning | S41, S42 | [VENDOR] |
| Command A+ Memory Usage Quality | 54% vs 39% (A Reasoning) | S41, S42 | [VENDOR] |
| Command A+ MMMU / MMMU Pro | 75.1% / 63% | S41, S42 | [VENDOR] |
| Command A+ MathVista | 80.6% (vs 73.5%) | S41, S42 | [VENDOR] |
| Command A+ CharXiv reasoning | 52.7% (vs 46.9%) | S41, S42 | [VENDOR] |
| Command A+ GPQA Diamond / AIME 2025 / HLE | 76% / 90% / ~11% | S43 (single source) | [SECONDARY] |
| Command A+ AA Intelligence Index | 37 (Mistral Medium 3.5: 39; GPT-5.5: 60) | S43 (single source) | [SECONDARY] |
| Command A+ TOPS vs A Reasoning | +63% higher output tok/s | S41, S42 | [VENDOR] |
| Command A+ TTFT vs A Reasoning | −17% | S41, S42 | [VENDOR] |
| Command A+ W4A4 incremental | +47% speed, −13% latency | S41, S42 | [VENDOR] |
| Command A+ speculative decoding | +1.5–1.6× speedup (text + multimodal) | S41, S42 | [VENDOR] |
| Command A+ quantization method | Quantization-Aware Distillation (QAD) | S42 (single source) | [SECONDARY] |
| Command A+ languages | 48 (vs 23 on Command A) | S41 (single source) | [VENDOR] |
| Cohere cumulative fundraising (vendor-reported) | ~$1.6B | S44 (single source) | [VENDOR] |

## Main actors
- **Cohere** — enterprise/sovereign AI lab; Command A+ (first Apache-2.0 model); FedRAMP High (May 12) [VENDOR].
- **Command A+** — 218B/25B MoE, Apache 2.0, "near lossless" quantization, 48 languages [VENDOR].
- **The Command closed line (A / R / R+)** — CC-BY-NC research-only predecessors; Command R+ carries the $3/$15 API tier [SECONDARY].
- **The Aya line** — Aya Expanse 8B / Aya Vision 8B retired from API 2026-04-04; no 2026 Aya refresh in the corpus [SECONDARY].
- **Sovereign and critical-infrastructure buyers** — the named audience of the Command A+ announcement and the FedRAMP High push [VENDOR].

## Timeline and context
- **2025** — Embed 4 and Rerank 4 released; Command A/R/R+ on CC-BY-NC research-only terms [SECONDARY].
- **2026-04-04** — Aya Expanse 8B and Aya Vision 8B retired from the API [SECONDARY].
- **2026-05-12** — FedRAMP High authorization [VENDOR].
- **2026-05-20** — Command A+ launches (218B/25B MoE, Apache 2.0) [VENDOR].
- **2026-08-10** — Meta Muse Glimmer 30B (Apache 2.0) — the other Western Apache-2.0 first of the summer [SECONDARY].
- **Sept 2026** — Command R+ API snapshot at $3/$15; no Command A+ API price row in the corpus's price table [SECONDARY].


### New verified timeline entries — expansion

- April 15, 2025 — Embed v4.0 released: first production multimodal embedding model, 65.2 MTEB, 128-language coverage per community mirror [SECONDARY]. (S22, S26)
- August 21, 2025 — Command A Reasoning released per Cohere changelog (2025, not 2026) [VENDOR]. (S19)
- September 2025 — Cohere extends Series D with $100M, reaching approximately $6.8–7B valuation [SECONDARY]. (S10, S11)
- Mid-2025 — Cohere ARR annualized at roughly $100M [SECONDARY]. (S11)
- February 2026 — Cohere discloses roughly $240M ARR [SECONDARY]. (S9)
- March 2026 — `cohere-transcribe-03-2026` audio transcription model appears in official docs catalog [COMMUNITY]. (S8)
- April 2026 — Cohere and Aleph Alpha sign term sheet; Series E of $2.5B listed in progress by PitchBook [SECONDARY]. (S9, S12)
- May 2026 — Cohere secures FedRAMP High authorization through Second Front Systems [SECONDARY]. (S13b)
- May 20–21, 2026 — Command A+ release: announced May 20 (vendor), secondary coverage May 21; ships in three precision formats BF16/FP8/W4A4 — the format detail not in the base launch entry [SECONDARY]. (S1, S5)
- May 24, 2026 — third-party deployment memo documents the 48-hour market impact window of the Command A+ release [SECONDARY]. (S6)
- July 30, 2026 — Cohere × Carahsoft partnership announced; North and Cohere models become available to U.S. public sector via Carahsoft distribution contracts [SECONDARY]. (S13)
- September 16, 2026 — Cohere × Aleph Alpha definitive agreement signed; ~$20B combined headline valuation; Schwarz Digits financing [SECONDARY]. (S9, S10)
- September 16–17, 2026 — OpenText × Cohere strategic partnership announced at ALL IN AI, Montreal, for trusted agentic AI in governments and regulated industries [SECONDARY]. (S14, S15)
- September 2026 — series-E round reported in advanced talks at up to $3B, per secondary reporting [SECONDARY]. (S11)

- January 9, 2025 — North introduced in early access as Cohere's "secure AI workplace platform" [SECONDARY]. (S33, S34, S37)
- March 13–16, 2025 — Command A released: 111B parameters, 256K context, 23 languages, two-GPU deployment, CC-BY-NC open weights [SECONDARY]. (S28, S29, S30)
- April 15, 2025 — Embed v4.0 released [SECONDARY]. (S22)
- July 2025 — $500M round at $5.5B valuation; North development accelerated on CoreWeave GB200 NVL72 [VENDOR]. (S33, S36)
- August 6, 2025 — North general availability (Dell, RBC as launch customers) [SECONDARY]. (S35)
- August 21, 2025 — Command A Reasoning released (111B, token budget, reasoning toggle) [VENDOR]. (S19, S31)
- September 2025 — Series D extended by $100M; valuation ~$6.8–7B [SECONDARY]. (S10, S11)
- February 2026 — $240M ARR disclosed [SECONDARY]. (S9)
- March 2026 — cohere-transcribe-03-2026 in docs catalog [COMMUNITY]. (S8)
- May 2026 — FedRAMP High authorization via Second Front Systems [SECONDARY]. (S13b)
- July 30, 2026 — Cohere × Carahsoft U.S. public-sector distribution deal [SECONDARY]. (S13, S13b, S13c)
- September 16, 2026 — Cohere × Aleph Alpha definitive agreement (~$20B combined valuation) [SECONDARY]. (S9, S10)
- September 16–17, 2026 — OpenText × Cohere strategic partnership at ALL IN AI [SECONDARY]. (S14, S15)

- March 2026 — Cohere Transcribe (speech-to-structured-text) released under Apache 2.0, opening the three-model cadence [SECONDARY]. (S39)
- June 9, 2026 — North Mini Code released: agentic coding sub-agent for large codebases [SECONDARY]. (S39, S40)
- June 25, 2026 — Wiz security-agent case study published, demonstrating North's MCP-native architecture [SECONDARY]. (S40)
- July 27, 2026 — North Automations generally available: agent orchestration with roles/permissions, human checkpoints, audit paths and usage dashboards [SECONDARY]. (S40)
- August 6, 2026 — Cohere × University of Waterloo certificate partnership announced [SECONDARY]. (S40)
- August 27, 2026 — Cohere Parse 5 (parse-v5.0) launches: 2.3B vision-language document parser, $1.50/1,000 pages [SECONDARY]. (S46, S48)

