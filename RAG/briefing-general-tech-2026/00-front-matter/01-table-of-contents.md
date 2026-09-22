---
id: briefing-general-tech-2026/00-front-matter/01-table-of-contents
title: "Front matter — table of contents"
domain: front-matter
role: reference
task: reference
actors: ["AMD", "Anthropic", "Apple", "BitNet", "China", "Credo", "FS.com", "Fujitsu", "Google", "Huawei", "IDC", "Intel", "MACOM", "MLCommons", "Meta", "Nokia", "Nvidia", "OIF", "OpenAI", "PrismML", "Qualcomm", "Samsung", "Telxius", "UALink", "UN", "United States"]
dates: ["2026-02", "2026-03", "2026-04", "2026-05", "2026-06", "2026-07", "2026-08", "2026-09", "2026-09-22"]
keywords: ["800v dc", "accelerator", "agentic", "agi", "ai200", "ascend", "awq", "benchmark", "bitnet", "blackwell", "capex", "clearwater forest"]
source: docs/RAG/briefing-general-tech-2026-en.md
source_anchor: "#g01"
source_lines: [10, 171]
sha256: ecc3c001526a1b418d24a7521c5a28f53beafb34d8a806438d296bd86df97a95
---

# Front matter — table of contents

<a id="g01"></a>
## 1. Front matter

This chapter is the dossier's entry point. It contains three things: a complete linked table of contents for all ten chapters and their subsections, an alphabetical index of the actors, products, events and concepts that appear across the dossier, and twenty thematic keywords with full definitions.

Conventions used throughout:

- Every chapter and subsection has a stable anchor of the form `#gNN` and `#gNN-k`; links below jump directly to the right section.
- Figures, dates and prices appear exactly as stated in the chapter bodies — nothing is inferred, rounded into precision, or carried across sections.
- Items dated after 22/09/2026 are marked **announced/planned**; they are not reported as facts.
- Glosses in the index and definitions in the keyword list are descriptive, not verdicts: they say what an item is and where it is covered, without adding new claims.

How to use this dossier:

- **Start from the table of contents** when you want the narrative arc of a topic (e.g. how the memory crisis developed, or the month-by-month chronology).
- **Use the alphabetical index** when you want to find "who/what is X" and jump to its main section.
- **Use the thematic keywords** when you need the dossier's working vocabulary (accelerator, quantization, rack-scale, token economics, …).
- **Do not assert what the dossier does not state.** If a fact is not in the chapter bodies, treat it as unknown even if it looks plausible.

<a id="g01-toc"></a>
### Table of contents

- [2. Introduction, methodology and narrative chronology](#g02)
  *Month-by-month narrative of the general-tech landscape from February to September 2026,*
  *plus the methodology and verification approach behind this dossier.*
  *Use it to locate any event in time before diving into the topical chapters.*
  - [2.1 Purpose and how to use this dossier](#g02-1)
  - [2.2 Methodology and verification approach](#g02-2)
  - [2.3 February 2026](#g02-3)
  - [2.4 March 2026](#g02-4)
  - [2.5 April 2026](#g02-5)
  - [2.6 May 2026](#g02-6)
  - [2.7 June 2026](#g02-7)
  - [2.8 July 2026](#g02-8)
  - [2.9 August 2026](#g02-9)
  - [2.10 September 2026](#g02-10)
- [3. GPUs and AI accelerators](#g03)
  *The accelerator race in 2026: Nvidia's Vera Rubin generation in full production and its first MLPerf outing,*
  *AMD's Advancing AI event and the Helios ramp with MI455X, Intel's Panther Lake and Clearwater Forest,*
  *Huawei's Ascend roadmap and the SuperPoD claims, plus Qualcomm's AI200 and Fujitsu's MONAKA.*
  *It closes with a competitive-landscape synthesis: who stands where at the end of September 2026.*
  - [3.1 Nvidia: Vera Rubin in full production](#g03-1)
  - [3.2 Rubin NVL72: architecture and specifications](#g03-2)
  - [3.3 MLPerf Inference v6.1: Rubin's first benchmark outing](#g03-3)
  - [3.4 Jensen Huang and the "AGI has arrived" moment](#g03-4)
  - [3.5 AMD: Advancing AI and the Helios ramp](#g03-5)
  - [3.6 MI455X and the memory play](#g03-6)
  - [3.7 AMD's deals: Anthropic, OpenAI, Meta](#g03-7)
  - [3.8 AMD crosses $1 trillion](#g03-8)
  - [3.9 Intel: Core Ultra X9 and the Panther Lake generation](#g03-9)
  - [3.10 Intel: Xeon 6+ Clearwater Forest and the datacenter fightback](#g03-10)
  - [3.11 Intel Crescent Island: the LPDDR5X inference bet](#g03-11)
  - [3.12 Huawei: Ascend roadmap at Huawei Connect 2026](#g03-12)
  - [3.13 Ascend 960DT vs 960PR: training and inference SKUs](#g03-13)
  - [3.14 Atlas 960 SuperPoD: scale, NPO optics and the 2.3x/2.5x claims](#g03-14)
  - [3.15 The per-chip gap: ~2 years behind Blackwell, ~10x under Rubin](#g03-15)
  - [3.16 Huawei's China-first strategy and PyTorch support](#g03-16)
  - [3.17 Qualcomm AI200: rack-scale inference](#g03-17)
  - [3.18 Fujitsu MONAKA: sovereign inference from Japan](#g03-18)
  - [3.19 Competitive landscape: who stands where](#g03-19)
- [4. Servers and datacenters](#g04)
  *The datacenter economy in 2026: a server market up 52% in Q2 (IDC), GPU-accelerated systems*
  *dominating value, the "RAMageddon" memory crisis, ~$725–730 billion of hyperscaler capex,*
  *custom ASICs gaining ground, Apple's Private Cloud Compute moving to M5, the 800V DC transition and UALink.*
  *It also covers Apple's datacenter moves and the two big infrastructure bets of the year: 800V DC power and open UALink interconnects.*
  - [4.1 The server market at +52%: IDC Q2 2026](#g04-1)
  - [4.2 GPU-accelerated systems dominate value](#g04-2)
  - [4.3 The memory crisis: "RAMageddon"](#g04-3)
  - [4.4 DRAM and NAND: prices, quotes, horizons](#g04-4)
  - [4.5 Hyperscaler capex: ~$725-730 billion in 2026](#g04-5)
  - [4.6 Custom ASICs gain ground](#g04-6)
  - [4.7 Apple: Private Cloud Compute moves to M5](#g04-7)
  - [4.8 The M8 Ultra enterprise server rumor](#g04-8)
  - [4.9 Power delivery: the 800V DC transition](#g04-9)
  - [4.10 UALink: the open coalition against NVLink](#g04-10)
- [5. Training vs inference architectures and quantization](#g05)
  *The great specialization of 2026: dense HBM training racks versus disaggregated, LPDDR-based inference,*
  *cost per token as the decisive metric, and the quantization ladder from FP8 to BitNet's 1.58-bit frontier.*
  *It ends with the runtimes that carry these techniques into production, and what quantization changes for the industry.*
  - [5.1 The great specialization: training vs inference](#g05-1)
  - [5.2 MLPerf v6.1 as an illustration of the shift](#g05-2)
  - [5.3 Nvidia Dynamo and disaggregated serving](#g05-3)
  - [5.4 Training racks: dense, HBM, high power](#g05-4)
  - [5.5 Inference racks: disaggregated, LPDDR, edge](#g05-5)
  - [5.6 Cost per token as the decisive metric](#g05-6)
  - [5.7 Quantization: the economics (McKinsey, June 2026)](#g05-7)
  - [5.8 FP8: the safe production standard](#g05-8)
  - [5.9 INT4, NVFP4, MXFP4: the aggressive standard](#g05-9)
  - [5.10 AWQ vs GPTQ, Marlin kernels](#g05-10)
  - [5.11 BitNet and the 1.58-bit frontier](#g05-11)
  - [5.12 PrismML Bonsai: 27B models in a few gigabytes](#g05-12)
  - [5.13 TurboQuant: 3-bit KV cache](#g05-13)
  - [5.14 Runtimes: vLLM, TensorRT-LLM, SGLang, llama.cpp](#g05-14)
  - [5.15 What quantization changes for the industry](#g05-15)
- [6. Networking and optics](#g06)
  *The plumbing of the AI era: FS.com's AI networking portfolio, 1.6T Ethernet shipping now,*
  *OIF 1600ZR coherent optics over a single wavelength, ECOC 2026 in Malaga,*
  *Linear Pluggable Optics, and the race toward 3.2T.*
  *It also tracks FS.com's own portfolio and H1 2026 financials alongside the industry-wide standards.*
  - [6.1 FS.com: the AI networking portfolio](#g06-1)
  - [6.2 1.6T OSFP and the "verified on NVIDIA" claim](#g06-2)
  - [6.3 D7070 800G muxponder](#g06-3)
  - [6.4 Wi-Fi 7 campus and AmpCon](#g06-4)
  - [6.5 FS.com H1 2026 financials](#g06-5)
  - [6.6 1.6T Ethernet is shipping now](#g06-6)
  - [6.7 OIF 1600ZR: coherent 1.6T over a single wavelength](#g06-7)
  - [6.8 ECOC 2026 in Malaga](#g06-8)
  - [6.9 Linear Pluggable Optics (LPO)](#g06-9)
  - [6.10 MACOM: 448G/lane toward 3.2T](#g06-10)
  - [6.11 Credo ZeroFlap 1.6T](#g06-11)
  - [6.12 Telxius deploys Nokia 800G coherent pluggables](#g06-12)
- [7. Consumer devices](#g07)
  *Foldables, AI PCs and AI laptops in 2026: Apple's first foldable iPhone Duo and the "Surprise and shine" event,*
  *Samsung's "Switchers" offensive and the Galaxy Z Fold 8, Google's return to AI laptops with the Googlebook,*
  *and the consumer price pressure driven by "RAMageddon".*
  *It closes with the market outlook: foldable share projections and the memory-driven pressure on consumer prices.*
  - [7.1 iPhone Duo: Apple's first foldable](#g07-1)
  - [7.2 The "Surprise and shine" event](#g07-2)
  - [7.3 Duo specifications in detail](#g07-3)
  - [7.4 Nano-texture and the crease question](#g07-4)
  - [7.5 Pricing, availability and market forecasts](#g07-5)
  - [7.6 Samsung's offensive: the "Switchers" campaign](#g07-6)
  - [7.7 Galaxy Z Fold 8: lighter, thinner, IP48](#g07-7)
  - [7.8 Foldable market shares: projections](#g07-8)
  - [7.9 Googlebook: Google returns to AI laptops](#g07-9)
  - [7.10 "Aluminium": codename, not a commercial name](#g07-10)
  - [7.11 Googlebook lineup, pricing and availability](#g07-11)
  - [7.12 AI PCs: 55% of shipments, but units fall](#g07-12)
  - [7.13 Consumer price pressure and "RAMageddon"](#g07-13)
- [8. AGI debate and governance](#g08)
  *The autumn 2026 AGI debate and its governance calendar: Huang's "AGI has arrived", Brockman's "AGI era",*
  *the DeepMind Institute, Amodei's "We Must Pace the Frontier", recursive self-improvement and the Hugging Face precedent,*
  *the 22-country declaration, the US–China incident line, and the planned UN Security Council and Trump–Xi summit.*
  *It ends with the definitions-and-timelines debate and the objections of the caution camp.*
  - [8.1 Huang: "AGI has arrived"](#g08-1)
  - [8.2 Brockman: "the AGI era"](#g08-2)
  - [8.3 The DeepMind Institute](#g08-3)
  - [8.4 Amodei: "We Must Pace the Frontier"](#g08-4)
  - [8.5 OpenAI: "Building standards for the next phase of AI"](#g08-5)
  - [8.6 Recursive self-improvement and the Hugging Face precedent](#g08-6)
  - [8.7 The 22-country declaration](#g08-7)
  - [8.8 US-China: Bessent-He Lifeng and the incident line](#g08-8)
  - [8.9 Altman at the UN Security Council (planned 23/09)](#g08-9)
  - [8.10 Trump-Xi summit (planned 24/09)](#g08-10)
  - [8.11 The UN Independent International Scientific Panel on AI](#g08-11)
  - [8.12 Definitions, timelines and the caution camp](#g08-12)
- [9. Market, deals and cross-cutting analysis](#g09)
  *What ties it together: a recap table of deals and funding, memory as the binding constraint,*
  *the US–China stack competition, the agentic infrastructure shift, and what to watch next.*
  *Memory as the binding constraint is the thread that ties the whole market story together.*
  - [9.1 Deals and funding recap table](#g09-1)
  - [9.2 Memory as the binding constraint](#g09-2)
  - [9.3 US-China stack competition](#g09-3)
  - [9.4 The agentic infrastructure shift](#g09-4)
  - [9.5 What to watch next](#g09-5)
- [10. Appendices](#g10)
  *Reference material: glossary, actors index, and methodological notes including sensitive points.*
  *Start here for terminology, then use the actors index to jump back into the narrative.*
  - [10.1 Glossary](#g10-1)
  - [10.2 Actors index](#g10-2)
  - [10.3 Methodological notes and sensitive points](#g10-3)

