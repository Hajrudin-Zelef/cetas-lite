---
id: briefing-general-tech-2026/02-gpus-accelerators/18-fujitsu-monaka
title: "Fujitsu MONAKA: sovereign inference from Japan"
domain: gpus-accelerators
role: deep-dive
task: hardware
actors: ["AMD", "Broadcom", "China", "Fujitsu", "Intel", "Nvidia", "Qualcomm"]
dates: ["2026-09-14", "2026-11", "2027-03", "2027-04"]
keywords: ["inference", "monaka", "accelerator", "benchmark", "funding", "packaging", "sovereignty"]
source: docs/RAG/briefing-general-tech-2026-en.md
source_anchor: "#g03-18"
source_lines: [2587, 2642]
canonical_for: ["fujitsu-monaka"]
sha256: 599eb5f76c22a93896f475d7dc7f2c9dc6174569ca26bb6f877e634818304c59
---

# Fujitsu MONAKA: sovereign inference from Japan

<a id="g03-18"></a>
### 3.18 Fujitsu MONAKA: sovereign inference from Japan

The most explicitly *sovereign* product in this chapter comes from Japan. Fujitsu's **MONAKA** — announced via **official press release on September 14, 2026, from Kawasaki** — is a **3D-stacked CPU** designed as sovereign inference infrastructure: Japanese-designed, Japanese-manufactured, aimed at Japanese and European buyers who want AI compute without American or Chinese silicon in the loop.

#### Technical package

| Specification | MONAKA |
|---|---|
| Stacking | 3D-stacked: **2 nm cores / 5 nm cache + I/O** |
| Clock | **3.8 GHz max** |
| Cores | **144** (4 dies × 36 + 4 SRAM chiplets; Broadcom 3.5D XDSiP; Armv9) |
| Acceleration | Hardware matrix acceleration + SVE2; Arm CCA |
| Servers | **1U (1–2 CPUs)** and **2U (2 CPUs)**, air-cooled, inference-optimized |
| Manufacturing | Made in Japan — **Kasashima** fab |
| Positioning | "Sovereign AI infrastructure" |

The packaging is the technical headline: **2 nm compute cores stacked with 5 nm cache and I/O**, assembled via **Broadcom's 3.5D XDSiP** into a 144-core device (four 36-core dies plus four SRAM chiplets). It is the same heterogeneous-integration direction the whole industry is moving (see Chan's "advanced packaging" item, §3.15), applied to a CPU rather than an accelerator. Hardware matrix acceleration plus SVE2 gives it the vector/matrix throughput inference needs; Arm CCA (Confidential Compute Architecture) addresses the security requirements of sovereign and enterprise buyers.

#### Armv9 in one paragraph

MONAKA's Armv9 architecture places it in the mainstream of the datacenter-CPU architecture shift: Arm's server instruction set, with SVE2 (Scalable Vector Extension 2) for vector throughput and dedicated hardware matrix acceleration for AI math. Arm CCA (Confidential Compute Architecture) adds hardware-enforced isolation for sensitive workloads — the security primitive sovereign buyers require. The architectural choice is also a supply-chain choice: an Arm-based design avoids x86 licensing entanglements and aligns Fujitsu with the broader Arm server ecosystem (which includes Nvidia's Grace/Vera lineage and the hyperscalers' in-house CPUs). For a sovereignty product, Arm is the pragmatic ISA: licensable, ecosystem-rich, and free of single-vendor American control at the architecture level.

#### Schedule: the shortest fuse in the chapter

| Milestone | Date | Status |
|---|---|---|
| Official press release (Kawasaki) | September 14, 2026 | Published |
| Sales start (Japan, Europe) | November 2026 | **Announced start** |
| Production ramp | Into April 2027 | Announced |
| Globalization (US, APAC) | Q4 FY2026 (Jan–Mar 2027) | Announced |

The schedule is precise and near-term: **sales start November 2026** — an **announced start**, Japan and Europe first — ramping into **April 2027**, with globalization to the **US and APAC in Q4 of Fujitsu's fiscal 2026 (January–March 2027)**. MONAKA is therefore the closest to delivery of the "challenger" products in this chapter's second half — weeks away from first sales at the time of writing, not quarters. That proximity is also what makes it the most falsifiable: by early 2027, either MONAKA servers are selling or they aren't, and the answer will be public.

#### The performance claim, labeled

The performance claim requires the standard label: **"2x throughput versus other CPUs"** is a **Fujitsu internal estimate** — no independent benchmark exists, and "other CPUs" is undefined in the verified sources. Treat it as a vendor estimate pending measurement. The comparison class matters enormously (2x vs. what? which workload? which precision?), and none of it is specified. Until Fujitsu or a third party publishes a methodology, the number is a placeholder.

#### Sovereignty as a product category

MONAKA's significance is not really about beating Xeon or EPYC on throughput; it is about the sovereign-compute thesis reaching silicon. A 144-core 3D-stacked Arm CPU, built in Japan, sold as inference infrastructure to governments and enterprises with sovereignty requirements, is a product category that did not exist five years ago. The buyers are not choosing on tokens-per-dollar alone — they are choosing on supply-chain provenance, jurisdictional control, and the political durability of their AI infrastructure. Whether the category has a market beyond procurement mandates is the open question — but Fujitsu, with a November 2026 sales start, is about to find out first, and the answer will inform every other sovereign-silicon program watching from Europe and Asia.

---

#### 3D stacking in one paragraph: what MONAKA's package buys

Conventional chips are flat: compute, cache, and I/O share one plane of silicon, and the distances between them are set by the die's footprint. 3D stacking builds upward — MONAKA places 2 nm compute dies atop 5 nm cache-and-I/O, bonded through Broadcom's 3.5D XDSiP integration, with dedicated SRAM chiplets in the stack. The payoff is bandwidth density: cache sits micrometers, not millimeters, from the cores, which collapses the latency and power cost of the cache hierarchy — exactly the tax that dominates inference, where the same weights are streamed repeatedly. The four-dies-plus-four-SRAM-chiplets construction (144 cores total) is the heterogeneous-integration mainstream of 2026 applied to a CPU: no single monolithic die could economically reach 144 cores at 2 nm, so the industry builds big chips out of small ones. MONAKA's package is unremarkable in method and notable in provenance — the same techniques, executed in Japan, for Japanese buyers.

#### The sovereign-silicon wave: MONAKA's company

MONAKA is not alone. The sovereign-compute thesis — that nations and blocs want AI infrastructure free of foreign supply-chain leverage — produced parallel programs across 2025–2026: European processor initiatives, Middle Eastern sovereign-AI buildouts (the HUMAIN-class buyers of §3.17), and Asian national programs. What distinguishes MONAKA is *delivery proximity*: a November 2026 sales start puts Fujitsu months, not years, from revenue, while most sovereign programs remain in the funding-and-roadmap phase. The 1U/2U air-cooled server packaging reinforces the point — these are products designed to slot into existing government and enterprise datacenters, not greenfield AI factories. Sovereignty, in Fujitsu's execution, looks like ordinary servers with extraordinary provenance.

The open question is whether sovereignty commands a price premium or imposes a performance discount the market will bear. Fujitsu's "2x throughput" internal estimate (unlabeled, unverified) is the company's opening bid in that negotiation. The November sales start — Japan and Europe first, US and APAC in Q4 FY2026 — will show whether the bid clears. If it does, expect the sovereign category to attract entrants; if it doesn't, MONAKA becomes a case study in the gap between procurement mandates and market demand.

---

