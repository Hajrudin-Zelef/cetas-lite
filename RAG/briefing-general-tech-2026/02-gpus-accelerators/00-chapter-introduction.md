---
id: briefing-general-tech-2026/02-gpus-accelerators/00-chapter-introduction
title: "GPUs and AI accelerators — introduction"
domain: gpus-accelerators
role: deep-dive
task: hardware
actors: ["AMD", "Broadcom", "China", "Fujitsu", "Huawei", "Intel", "MLCommons", "Meta", "Micron", "Nvidia", "OCP", "OpenAI", "Qualcomm"]
dates: ["2025-10", "2026-09", "2026-09-22", "2026-11"]
keywords: ["accelerator", "gpu", "gpus", "18a", "agi", "ai200", "ascend", "asic", "benchmark", "clearwater forest", "crescent island", "decode"]
source: docs/RAG/briefing-general-tech-2026-en.md
source_anchor: "#g03"
source_lines: [1332, 1441]
sha256: f1fb25f8b22dcefbe3ad23bfd895696acf3cf52f6d5229b0afca064e446203a1
---

# GPUs and AI accelerators — introduction

<a id="g03"></a>
## 3. GPUs and AI accelerators

> **Coverage:** February → September 2026 · **Chapter status:** all figures below are sourced from vendor announcements, press coverage, or benchmark publications in the coverage window; 2025 dates appear only where they frame a 2026 event (e.g. partnership agreements whose 2026 milestones are the story). Vendor performance claims are labeled as vendor claims; anything dated after September 22, 2026 is "announced/planned", not shipped.

If 2025 was the year the AI accelerator market learned to think in racks, 2026 is the year the rack became the product. Every major vendor in this chapter — Nvidia, AMD, Intel, Huawei, Qualcomm, Fujitsu — now sells, or claims to be building, a rack-scale system rather than a chip. The unit of competition moved up one level: from per-GPU FLOPS to per-rack throughput, memory capacity, interconnect bandwidth, and power envelope. This chapter walks through that shift vendor by vendor, then places them on a single competitive map.

The cast divides into three groups. First, the incumbents fighting over the frontier-training and large-scale inference market: Nvidia with Vera Rubin, AMD with Helios and the MI455X, Intel trying to re-enter through Xeon 6+ and a dedicated inference GPU called Crescent Island. Second, Huawei, which staged at Huawei Connect 2026 the most ambitious system-level answer to the US export regime: an entire annual accelerator roadmap, a 15,488-chip SuperPoD, and a first-of-its-kind optical fabric — all aimed at a China-first market. Third, the specialists: Qualcomm's AI200 rack-scale inference push and Fujitsu's MONAKA, a sovereign Japanese inference CPU.

#### Why hardware gets a chapter

A word on the dossier's architecture: AI news in 2026 is dominated by models, labs, and capital — the subjects of the dossier's other chapters. Hardware gets its own chapter because it is the *substrate* those stories run on: every model launch is a demand event for accelerators, every funding round is partly a compute purchase, and every "AGI" claim (§3.4) is implicitly a claim about hardware scale. The chapter's job is to keep the substrate visible — to record what the machines are, what they cost, when they ship, and whose claims about them are measured versus announced. Models are the weather; hardware is the climate. This chapter is the climate report.

### How this chapter is organized

Sections 3.1–3.4 cover Nvidia: the production ramp of Vera Rubin, the NVL72's architecture, its first MLPerf outing, and the "AGI has arrived" media moment. Sections 3.5–3.8 cover AMD: the Advancing AI event and Helios ramp, the MI455X's memory-first positioning, the three lab deals, and the $1 trillion crossing. Sections 3.9–3.11 cover Intel across client (Panther Lake/X9), datacenter CPU (Xeon 6+ Clearwater Forest), and the Crescent Island inference GPU. Sections 3.12–3.16 cover Huawei: the Connect 2026 roadmap, the 960DT/960PR split, the Atlas 960 SuperPoD, the per-chip gap, and the China-first strategy plus PyTorch support. Sections 3.17–3.18 cover Qualcomm and Fujitsu. Section 3.19 synthesizes the competitive landscape.

### A note on claims and evidence

This chapter applies a strict evidentiary discipline, because 2026 was a year of dueling announcements and the press frequently blurred three different things: **claimed specifications** (numbers in a keynote for unshipped silicon), **measured results** (benchmarks on real hardware), and **announced calendars** (delivery dates). Throughout the chapter:

- Figures from vendor keynotes and press releases are labeled **vendor claims**.
- Benchmark numbers from MLCommons/MLPerf are **measured results** — with the submitter and division noted.
- Delivery dates after September 22, 2026 are **announced/planned**, and the text distinguishes "in production" (manufacturing) from "shipping to customers" (delivery) wherever the sources allow it.
- Analyst characterizations are attributed as analyst framing, not vendor statements.

### The three structural shifts of 2026

Three patterns recur across every vendor in this chapter, and they are worth naming up front because they explain decisions that otherwise look idiosyncratic.

**First, the annual cadence became the industry standard.** Nvidia set the tempo with Rubin → Rubin Ultra → Feynman; AMD answered with MI450 → MI500; Huawei announced its own annual "Tao's Law." Whether or not every vendor hits every date, the *announcement* of an annual rhythm is now table stakes — a signal to buyers that committing to a platform today will not strand them tomorrow.

**Second, memory displaced raw FLOPS as the contested metric.** AMD's entire Helios pitch is memory capacity and bandwidth; Intel's Crescent Island bets on LPDDR5X capacity-per-dollar; Huawei's SuperPoD is a memory-and-interconnect story; even Nvidia's MLPerf wins were attributed to software (NVFP4, disaggregated serving) as much as silicon. The binding constraint on serving large models in 2026 is widely understood to be memory — capacity for KV caches at long context, bandwidth for token generation — and every vendor positioned accordingly.

**Third, power and cooling became deployment constraints, not footnotes.** Rubin NVL72 at 190–230 kW per rack, Qualcomm's 160 kW liquid-cooled racks, Huawei's 220-cabinet SuperPoD with >550 kW of optical power savings — the datacenter, not the chip, is now the binding constraint on how fast new platforms can actually deploy. Several of 2026's "delay" narratives are better read as power-delivery narratives.

With that frame in place: Nvidia first, because everyone else's 2026 was defined in reaction to it.

#### What this chapter does not cover

Scope discipline is part of the method. This chapter covers **AI accelerators and the datacenter CPUs paired with them** — GPUs, NPUs, and inference CPUs sold as AI infrastructure. It does not cover: general-purpose consumer chips beyond Panther Lake/X9 (which is included as Intel's 18A proof point); networking vendors except where their products are part of an accelerator story (Broadcom appears only in the trillion-dollar club, §3.8); memory vendors except Micron's club membership; or the cloud providers' in-house silicon programs except as customers. Model developments are covered only where they intersect hardware narratives (Astra in §3.4, GPT-class pilots in §3.7). The chapter's unit of analysis is the *system you can buy or that a vendor claims you will be able to buy* — everything else belongs to other chapters of the dossier.

#### How the anchors work

Like the rest of this dossier, the chapter uses stable HTML anchors (`#g03` for the chapter, `#g03-1` through `#g03-19` for the sections) placed on their own line immediately before each heading. They exist so that models, search tools, and cross-references can cite a section unambiguously — `#g03-14` always means the Atlas 960 SuperPoD section, regardless of how the surrounding text is edited. When attaching this chapter to a model's context for Q&A, the anchors let the model ground its answers in specific sections rather than reciting from weights — the same mechanism the dossier's other chapters use.

#### The coverage window: why February to September

The dossier's February–September 2026 window captures the accelerator industry's decisive eight months: from the Meta–AMD deal (February 24) that opened the year's demand story, through the Rubin production evidence (June 1) and the Advancing AI launch (July 22–23), to the September pile-up — Astra (3rd), the Huang post (6th), MLPerf v6.1 (16th), Huawei Connect (17th–19th), and AMD's trillion-dollar crossing (21st). October 2025 dates appear only as prologue (the OpenAI–AMD deal, Qualcomm's AI200 announcement, Crescent Island's OCP unveiling); post-September-22 dates appear only as announced calendars. The window is, in effect, the industry's 2026 as it looked one week before the end of Q3 — claims fresh, deliveries pending, verdicts unwritten.

#### Dramatis personae: the six vendors in one line each

| Vendor | 2026 hardware story | One-line verdict (Sept 2026) |
|---|---|---|
| Nvidia | Vera Rubin: in production, benchmarking, deploying | The incumbent; everyone else's roadmap is calibrated against it |
| AMD | Helios/MI455X: announced, 14 GW of lab deals, unshipped | The credible challenger; execution is the entire question |
| Intel | Xeon 6+ shipping; Crescent Island inference GPU announced | The CPU incumbent fighting on control-plane + value |
| Huawei | Ascend 960 roadmap; Atlas 960 SuperPoD; NPO optics | The system-level compensator; China-first by choice and constraint |
| Qualcomm | AI200 rack-scale inference; HUMAIN 200 MW | The efficiency player scaling up to racks |
| Fujitsu | MONAKA 3D-stacked CPU; sales November 2026 | The sovereign bet, closest to revenue |

---

#### A reader's guide to the evidentiary labels

The chapter uses a small controlled vocabulary for evidence status, applied consistently:

| Label | Meaning | Example |
|---|---|---|
| Vendor claim | Stated by the vendor; no independent verification | AMD's 30% tokens/dollar |
| Measured | Third-party benchmark under standard rules | MLPerf 3.7x on Qwen3-VL |
| Announced / announced calendar | Dated future commitment | Helios first deliveries September |
| Announced/planned | Roadmap item dated after the cutoff | Atlas 960 SuperPoD Q4 2027 |
| Analyst framing | Market interpretation, not a vendor statement | "Ideal for memory-bound workloads" |
| Per sources / unofficial | Press-cited but undisclosed | $60–100B Meta deal valuation |
| Internal estimate | Vendor's own measurement, methodology undisclosed | Fujitsu's "2x throughput" |

When a figure appears without a label, it is either a dated event (keynotes, launches — inherently verifiable) or arithmetic on labeled figures. The discipline is the point: in a year of dueling announcements, the label is often more informative than the number.

#### A note on sources

The chapter's facts come in four kinds, and the text labels them throughout: **vendor disclosures** (keynotes, press releases, Hot Chips talks — the majority of spec figures); **press reporting** (Reuters, SCMP, Nikkei Asia, Bloomberg, The Register, Motley Fool — dates, quotes, market data); **benchmark publications** (MLCommons/MLPerf v6.1 — the only third-party measurements); and **official standards bodies** (the Linux Foundation for PyTorch backend recognition). No single kind suffices: vendor disclosures give specs but not verification, press gives dates and quotes but not measurements, benchmarks give measurements but only where hardware exists to test. The dossier's rule — stated in the chapter header and applied throughout — is that the *kind* of a fact is part of the fact, and omitting it is a form of fabrication.

#### Glossary: the chapter's working vocabulary

Readers meeting rack-scale terminology for the first time will encounter the same dozen terms in every vendor's materials. They are defined once, here, in vendor-neutral language:

| Term | Meaning in this chapter |
|---|---|
| NVL72 / NVL domain | A rack-scale compute domain: 72 GPUs (plus host CPUs) joined by a scale-up fabric into a single addressable system |
| Scale-up | Interconnect *within* the rack/domain (e.g. NVLink 6) — high bandwidth, short distance |
| Scale-out | Interconnect *between* racks (e.g. Ethernet) — longer distance, lower bandwidth per link |
| SuperPoD | Huawei's term for a datacenter-scale system: hundreds of cabinets presented as one product |
| NPO (near-packaged optics) | Optical interconnect engines placed adjacent to the switch ASIC, replacing pluggable optical modules |
| HBM4 / HBM | High-bandwidth memory stacked on/near the accelerator package; HBM4 is the 2026 generation |
| LPDDR5X | Low-power DDR memory — cheaper and more capacious per dollar than HBM, at lower peak bandwidth |
| FP4 / FP8 | 4-bit and 8-bit floating-point precisions; fewer bits = smaller models, higher effective throughput |
| NVFP4 / MXFP4 | Vendors' 4-bit floating-point formats (Nvidia's / the Microscaling shared format) |
| KV cache | Stored attention state that grows with context length and batch size — the capacity hog of long-context inference |
| Tokens per dollar | Cost-efficiency metric for inference: how many generated tokens a dollar buys |
| Disaggregated serving | Splitting inference's prefill (prompt processing) and decode (token generation) phases onto separate resources |
| MoE (mixture of experts) | Model architecture routing each token to a subset of "expert" sub-networks |
| TOFU / warrant structures | *(Not hardware — financial terms appearing in §3.7's deal analysis)* |

#### A note on units

Compute figures in this chapter span six orders of magnitude, so the units are worth fixing: **PFLOPS** = 10¹⁵ floating-point operations per second (per-chip figures), **EFLOPS** = 10¹⁸ (rack- and system-level figures). Bandwidth is in **TB/s** (terabytes per second) for fabrics and **Tbit/s** (terabits) for optical modules — a factor-of-8 difference that matters when comparing Huawei's 7.2 Tbit/s Hi-ONE modules against TB/s-scale fabric numbers. Power is **kW** per rack, **MW** per deployment, **GW** per deal — and §3.7's gigawatt figures refer to *datacenter power capacity committed*, not to chip counts. Keeping the three ladders (compute, bandwidth, power) separate is the simplest defense against the chapter's most common misreadings.

---

