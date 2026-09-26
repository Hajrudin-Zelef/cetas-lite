---
id: ai-industry-kb-2026/15-hardware-chips/overview
title: "15. Hardware & Chips"
domain: hardware-chips
role: deep-dive
task: hardware
actors: ["AMD", "AWS", "Anthropic", "Broadcom", "Cerebras", "CoreWeave", "DeepSeek", "Google", "Groq", "Meta", "Microsoft", "Nebius", "Nvidia", "OpenAI", "Samsung", "TSMC"]
dates: ["2025-03", "2025-04", "2025-06", "2025-12-24", "2026-01-05", "2026-03", "2026-03-16", "2026-05-14", "2026-06-30", "2026-07", "2026-07-21", "2026-08-18", "2026-08-24", "2026-09", "2026-09-22", "2026-10"]
keywords: ["accelerator", "acquisition", "amd", "asic", "benchmarks", "blackwell", "compute", "custom silicon", "decode", "deepseek", "ethernet", "fp4"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [7389, 7440]
section: "15. Hardware & Chips"
sha256: b59d7dc7edc71e335415c8a8ca4eb8a5937f466ae10738647426825173d0f4ac
---

# 15. Hardware & Chips
Keywords: Vera Rubin NVL72, NVL144 rename, Rubin full production CES 2026, Rubin GPU, Vera CPU, NVLink 6, HBM4, Rubin shipments Q3 2026, CoreWeave Rubin bring-up, Dell PowerEdge XE9812, AMD Helios, MI455X, EPYC Venice, Pensando Vulcano, MI300X rental, Groq 3 LPU, Groq 3 LPX, Cerebras IPO, WSE-3, Etched Sohu, Broadcom XPU, TPU v7 Ironwood, Trainium 3, Maia 200, GB200 NVL72, Rubin Ultra Kyber, tokens per dollar per watt, rack-scale AI

## Summary

### The headline: Vera Rubin NVL72 is in full production, on schedule

- **Vera Rubin entered "full production" at CES 2026 (January 5, 2026), per Jensen Huang's keynote — VERIFIED.** All six platform chips had returned from manufacturing partners; customer deployments were guided for H2 2026.
- **CES 2026 was the production confirmation, NOT the announcement.** The Rubin platform was first announced at **GTC 2025 (March 2025)**. Any RAG entry describing CES 2026 as the Rubin reveal is wrong; it is the "full production" milestone of a March-2025-announced platform.
- **Canonical status as of September 22, 2026: ON SCHEDULE within the original H2 2026 window — not an acceleration.** Production shipments began in NVIDIA's fiscal Q3 (August–October 2026), exactly inside the guided window. The framing "accelerated vs H2 2026" found in some briefing material is corrected: Q3 sits inside H2, so no pulled-in date is supported by evidence.
- **Production was re-confirmed on July 21, 2026** by NVIDIA hyperscale/HPC VP Ian Buck: "We are absolutely in full production right now. All the major customers are deploying these systems."

### Naming correction: NVL144 (2025) = NVL72 (2026), same rack

- The GTC 2025 name **NVL144** counted GPU *dies* (72 packages × 2 dies = 144). In 2026 materials the rack was renamed **NVL72** (counting packages). It is the same hardware.
- Every RAG entry saying "Vera Rubin NVL144" is using the 2025 name — map it to NVL72.
- **Do not conflate with Rubin Ultra NVL576** ("Kyber" rack, 2027): 576 Rubin GPU dies, 15 EF FP4 / 5 EF FP8, HBM4e, NVLink 7. SemiAnalysis (July 2026) reported a possible Ultra/Kyber slip to 2028; NVIDIA said the roadmap was intact [SECONDARY — reported, not confirmed].

### Vera Rubin specs — VERIFIED as NVIDIA-stated [VENDOR]

- **3.6 exaflops FP4 inference / 1.2 exaflops FP8 training** per rack, ~3.3× a GB300 NVL72 rack; 75 TB HBM4 total (288 GB per GPU, 8 stacks, up to ~15 TB/s per GPU); NVLink 6 at 3.6 TB/s per GPU bidirectional, 260 TB/s all-to-all per rack.
- Vera CPU: 88 custom "Olympus" Arm v9.2-A cores, 176 threads, 1.8 TB/s NVLink-C2C. Process: TSMC N3P (3 nm). HBM4 suppliers: Samsung and SK Hynix.
- Six co-designed platform chips: Rubin GPU, Vera CPU, NVLink 6 switch, ConnectX-9 SuperNIC, BlueField-4 DPU, Spectrum-6 Ethernet switch.
- Cable-free modular compute-tray design cut rack assembly from ~2 hours to ~5 minutes (~95%); truck-to-power-on in 47 minutes claimed — engineered as the answer to Blackwell's 2024–2025 ramp issues (reported overheating/design problems when racked at scale).
- **Measured datapoint (CoreWeave, [VENDOR] via company publication):** DeepSeek R1 on Vera Rubin NVL72 delivers **10× tokens/sec/MW** vs the previous Blackwell generation.
- **Pricing (third-party only, never NVIDIA-stated):** one buyer reported **$7–8M per Vera Rubin NVL72 rack** vs ~$5M for a GB300 rack [UNVERIFIED — single secondary]. Earlier March 2026 press put it at $5–7M/rack.
- **Sell-side estimate [UNVERIFIED]:** ~5.7M Rubin GPUs shipping in 2026 (cryptobriefing). The viral "1,000 racks/day" figure traces to no NVIDIA statement or named analyst — **do not cite**.

### AMD answered with Helios

- **Lisa Su's CES 2026 keynote (January 5/6, 2026): detailed reveal of the Helios rack-scale platform** — "world's best AI rack." Correction: Helios was first teased at Advancing AI (June 2025); CES 2026 was the detailed reveal, not the first mention.
- 72 Instinct MI455X per rack, 31 TB HBM4, 1.4 PB/s memory bandwidth, **2.9 FP4 exaFLOPS inference / 1.4 FP8 exaFLOPS training** [VENDOR]; Zen-6 EPYC "Venice" (256 cores/512 threads, 2 nm); Pensando Vulcano NICs; liquid-cooled, ~7,000 lbs; launches **later in 2026**. MI440X (enterprise 8-GPU) and MI500 (2027, CDNA 6, HBM4E, up to "1,000× vs MI300X" claim [VENDOR]) also previewed.
- MI300X remains a thinly listed rental product ($2.00–$3.59/hr at neoclouds); MI355X was "contact sales" only — **no public rental listing found** [UNVERIFIED for rental availability].

### The inference-silicon wave is the other half of 2026

- **NVIDIA–Groq, December 24, 2025: ~$20B non-exclusive technology licensing agreement** (NOT an acquisition — corrects earlier "acquisition" phrasing). Groq founder Jonathan Ross and president Sunny Madra joined NVIDIA; Groq continues independently under CEO Simon Edwards; GroqCloud unaffected.
- **Groq 3 LPU** unveiled at GTC 2026 (March 16, 2026) as the 7th Vera Rubin chip; **in full production announced at Hot Chips, August 24, 2026**; Nebius is the first cloud adopter (Token Factory). NVIDIA's design split: Rubin NVL72 handles prefill/context, Groq 3 LPX racks handle decode.
- **Rubin CPX NVL rack put on hold/cancelled at GTC 2026**, replaced by the Groq 3 LPX in the Rubin platform family (after the $20B Groq licensing deal).
- **Cerebras is now public: IPO completed May 14, 2026, Nasdaq CBRS, priced at $185/share, $5.55B raised** — the largest IPO of 2026. WSE-3 at scale; Artificial Analysis measured Gemma 4 31B at 1,990.8 tok/s (Jul 2026). Caveat: independent academic modeling gives NVIDIA's B200 1.5–3× better performance-per-watt-per-dollar at scale.
- **Etched Sohu transformer ASIC ships June–September 2026** (stealth exit June 30, 2026; TSMC N4P first-pass silicon success; >$1B in signed customer contracts claimed; first rack to Jane Street in July 2026; $700M raised at $21B valuation August 18, 2026). **Still no independent third-party benchmarks as of September 22, 2026** — all throughput claims are company-reported.
- **Broadcom is the quiet inference giant:** custom AI XPU business did **$20B+ in FY2025**; $8.4B AI semiconductor revenue in Q1 FY2026 (+106% YoY); serves Google, Meta, and OpenAI simultaneously; a 10-GW AI accelerator + networking deal with OpenAI projected above $100B. 2026 market shorthand: "NVIDIA owns training, Broadcom builds inference."

### Hyperscaler custom silicon ramps in 2026

- **Google TPU v7 "Ironwood"** — announced at Google Cloud Next **2025** (April 2025), 2026 is the production-deployment ramp year: 4.6 PFLOPS dense FP8 per chip, 192 GB HBM3E, 9,216-chip Superpod scaling to 42.5 EF aggregate inference [SECONDARY, consistent across sources].
- **Amazon Trainium 3** — volume production in 2026 (3 nm class, 2.52 PFLOPS MXFP8, 128–144 GB HBM3E); **validated by both Anthropic and OpenAI** workloads — third-party validation the earlier generations lacked.
- **Microsoft Maia 200 ("Braga")** — 3 nm, 216 GB HBM3e at 7 TB/s (most memory of the 2026 hyperscaler cohort); Microsoft claims 3× the FP4 performance of Trainium 3 [VENDOR]; Azure inference deployment.
- Meta MTIA 500 (3 nm roadmap) is internal-only with no external availability.

### The bifurcation thesis frames the whole market

