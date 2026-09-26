---
id: ai-industry-kb-2026-wave6/16-cohere/overview
title: "§16. Cohere"
domain: cohere
role: deep-dive
task: actor-profile
actors: ["Cohere", "Google", "Meta", "Moonshot", "Nvidia", "OpenAI"]
dates: ["2026-04-04", "2026-05-12", "2026-05-20", "2026-09"]
keywords: ["cohere", "agent", "apache", "benchmark", "claude", "context window", "cost", "kimi", "license", "moe", "nvidia", "open weights"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [7799, 7865]
section: "§16. Cohere"
delta_of: ai-industry-kb-2026
sha256: 682034020e563217184e80bd441e06f87513e58ac13b0de4ebf91afa6cc5a87a
---

# §16. Cohere

Keywords: Cohere, Command A+, Command A, Command R, Command R+, Aya Expanse 8B, Aya Vision 8B, Embed 4, Rerank 4, Apache 2.0, sovereign AI, FedRAMP High, model retirement

## Summary
- **Command A+** launched **2026-05-20** [VENDOR]: **218B total / 25B active MoE**, **Apache 2.0** — Cohere's **first Apache-2.0 / open-weight model** [VENDOR].
- Vendor framing discipline matters: say "first Apache-2.0 / open-weight Cohere model," not unqualified "first fully open-source" — the vendor's own announcement language is the ceiling of the claim [VENDOR].
- Quantization is described as **"near lossless," not "lossless"** [VENDOR] — the adjective is load-bearing; do not upgrade it.
- Specs: **48 languages**, **128K input** context, **64K generation**, minimum deployment **1× B200 or 2× H100** [VENDOR].
- The announcement positioned Command A+ for **sovereign and critical-infrastructure** AI — the enterprise/sovereignty pitch that also frames the FedRAMP High authorization (2026-05-12) [SECONDARY].
- **Aya Expanse 8B and Aya Vision 8B were retired from the API on 2026-04-04** [SECONDARY] — the Aya 8B API line ended a month before Command A+ shipped.
- No 2026 successors found for **Command R7B**, **Embed 4**, or **Rerank 4** — all three are 2025-era or earlier lines with no 2026 refresh in the corpus.
- **FedRAMP High authorization: 2026-05-12** [VENDOR] — the compliance milestone underpinning the sovereign-infrastructure positioning.
- The section is thin on 2026 releases because Cohere's 2026 is a one-release year in this corpus: Command A+ is the only new model; everything else is retirements, compliance, and standing 2025 lines [DIRECTIONAL].

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

