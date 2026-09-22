---
id: briefing-general-tech-2026/03-servers-datacenters/00-chapter-introduction
title: "Servers and datacenters — introduction"
domain: servers-datacenters
role: deep-dive
task: infrastructure
actors: ["Amazon", "Apple", "Broadcom", "Google", "IDC", "Meta", "Microsoft", "Nvidia", "OCP", "UALink"]
dates: ["2026-08-12", "2026-09", "2026-09-16"]
keywords: ["800v dc", "accelerator", "capex", "custom silicon", "dram", "gpu", "gpus", "hbm", "hbm4", "hyperscaler", "maia", "mtia"]
source: docs/RAG/briefing-general-tech-2026-en.md
source_anchor: "#g04"
source_lines: [2882, 2964]
sha256: 5a55c6d6ea92c8c60395e4fef83234152e2cb0e5273683a15f614009adfdc62c
---

# Servers and datacenters — introduction

<a id="g04"></a>
## 4. Servers and datacenters

> **Chapter scope:** the physical substrate of the AI boom — server markets, memory, hyperscaler capital spending, custom silicon, Apple, power delivery and interconnects. **Coverage period:** February → September 2026. **Anti-fabrication rule:** no figure, date or quote beyond the verified facts below; vendor claims, rumors and projections are labeled as such throughout.

If 2025 was the year the AI datacenter became the most valuable real estate on
Earth, 2026 is the year the invoice came due. Every layer of the hardware stack
told the same story from a different angle: the server market hit an all-time
record in the second quarter, driven almost entirely by GPU-accelerated systems
whose unit shipments were actually *falling* while their prices soared. Memory —
DRAM and NAND — became the binding constraint of the whole industry, a shortage
so severe that analysts coined portmanteaus like "RAMageddon" and chief
executives spoke of 2028, even 2030, before balance returned. The four American
hyperscalers committed roughly three-quarters of a trillion dollars in capital
expenditure for a single year. Everyone with the scale to justify it — Meta,
Google, Amazon, Microsoft — doubled down on custom silicon to escape the
merchant accelerator market's pricing. Apple moved its Private Cloud Compute
infrastructure onto the M5, skipping two silicon generations. And beneath the
racks, the industry quietly started rebuilding the fundamentals: higher-voltage
power delivery and an open interconnect standard meant to break a single
vendor's lock on scale-up fabrics.

There is a unifying thread running through all ten sections of this chapter, and
it is worth naming up front: **2026 was the year the AI buildout hit physical
limits**. Not limits of demand — demand, as the capex figures show, was
essentially unbounded. Limits of supply: of memory fab capacity, of wafer
starts, of power that can be pushed through copper, of the number of
accelerators that can be wired into a single coherent domain. Every section
below is, in one way or another, about what happens when exponential demand
meets the stubborn physics and lead times of the material world. The industry's
answers — pay more, build your own chips, redesign the power chain, standardize
the interconnect — are the real subject of this chapter.

A note on method before the sections begin. This chapter covers everything
*around and beneath* the AI chips themselves: the boxes they ship in, the memory
crisis that repriced those boxes, the money that funded them, the power that
feeds them, and the wires that connect them. The chips themselves — GPUs,
accelerators, process nodes — belong to Chapter 3 ("Chips"). Where the two
chapters touch, cross-references are given rather than repeated.

**Chapter table of contents**

- [4.1. The server market at +52%: IDC Q2 2026](#g04-1)
- [4.2. GPU-accelerated systems dominate value](#g04-2)
- [4.3. The memory crisis: "RAMageddon"](#g04-3)
- [4.4. DRAM and NAND: prices, quotes, horizons](#g04-4)
- [4.5. Hyperscaler capex: ~$725–730 billion in 2026](#g04-5)
- [4.6. Custom ASICs gain ground](#g04-6)
- [4.7. Apple: Private Cloud Compute moves to M5](#g04-7)
- [4.8. The M8 Ultra enterprise server rumor](#g04-8)
- [4.9. Power delivery: the 800V DC transition](#g04-9)
- [4.10. UALink: the open coalition against NVLink](#g04-10)

**Chapter at a glance — the ten headline facts**

| # | Section | Headline figure |
|---|---|---|
| 4.1 | Server market, IDC Q2 2026 | +52% YoY to $166.3B — all-time record; units +15.4% |
| 4.2 | GPU-accelerated systems | 52.6% of value ($87.4B); units −10.8%; ASP +43.6% to ~$170,200 |
| 4.3 | Memory crisis ("RAMageddon") | Server DRAM ×3.5 by September; NAND ×3; IDC: worst in 15 years |
| 4.4 | Prices, quotes, horizons | Suppliers guide relief in 2028 at earliest; wafers short into 2030 |
| 4.5 | Hyperscaler capex | Big 4 ~$725–730B in 2026 vs ~$410B in 2025 |
| 4.6 | Custom ASICs | Meta "Iris" in mass production (Sep 2026); TPU v7, Trainium3, Maia 200 |
| 4.7 | Apple PCC | M5 transition (9to5Mac Feb 2026; corroborated Aug 2026); M3 Ultra/M4 skipped |
| 4.8 | M8 Ultra rumor | The Information 16/09/2026: enterprise servers, not before 2029, may be cancelled |
| 4.9 | 800V DC power | Nvidia + Google + Microsoft + 80 partners via OCP (12/08/2026); racks H2 2026 |
| 4.10 | UALink | Open consortium vs NVLink; 2.0 spec Apr 2026; no shipping silicon in 2026 |

**Key terms used in this chapter.**

| Term | Meaning in this chapter |
|---|---|
| ASP | Average selling price — here, of GPU-accelerated servers (~$170,200, IDC Q2 2026) |
| HBM | High-bandwidth memory — DRAM stacks bonded to AI accelerators (HBM3E, HBM4) |
| DRAM / NAND | Main memory (volatile) / flash storage (non-volatile) — the two media of the memory crisis |
| OCP | Open Compute Project — the open-hardware consortium behind the 800V DC program |
| MGX | Nvidia's open rack architecture — the 800V DC racks are MGX-compatible |
| NVLink / NVLink Fusion | Nvidia's proprietary scale-up interconnect; Fusion extends it to third-party chips |
| UALink | Ultra Accelerator Link — the open scale-up interconnect consortium |
| MTIA / TPU / Trainium / Maia | Custom AI ASICs of Meta, Google, Amazon, Microsoft respectively |
| PCC | Private Cloud Compute — Apple's privacy-sealed, Apple-silicon cloud for Apple Intelligence |
| Capex | Capital expenditure — here, overwhelmingly datacenter and AI infrastructure build |

