---
id: briefing-general-tech-2026/02-gpus-accelerators/12-huawei-ascend-roadmap
title: "Huawei: Ascend roadmap at Huawei Connect 2026"
domain: gpus-accelerators
role: deep-dive
task: hardware
actors: ["AMD", "China", "Huawei", "Intel", "MLCommons", "Nvidia", "Qualcomm"]
dates: ["2026-09"]
keywords: ["ascend", "accelerator", "agentic", "agi", "export controls", "fp4", "fp8", "hbm", "inference", "mlperf", "mlperf v6.1", "sovereignty"]
source: docs/RAG/briefing-general-tech-2026-en.md
source_anchor: "#g03-12"
source_lines: [2203, 2262]
canonical_for: ["huawei-ascend"]
sha256: c4a9fcaccc797ff8a0189194fca897e197f6ca386f815e2b631548de60d245d7
---

# Huawei: Ascend roadmap at Huawei Connect 2026

<a id="g03-12"></a>
### 3.12 Huawei: Ascend roadmap at Huawei Connect 2026

**Huawei Connect 2026** ran **September 17–19 in Shanghai** under the theme **"Advancing the Agentic World,"** with **David Wang's keynote on September 17**. It was, by any measure, the most consequential Huawei event for the global AI-hardware conversation since the US export controls: a full-stack, multi-year accelerator roadmap presented as an answer to the question of whether China can build frontier AI infrastructure without American silicon.

#### The roadmap: "Tao's Law"

The roadmap's headline is cadence. Huawei announced an **annual release rhythm** — internally dubbed **"Tao's Law"** — with the following sequence, all **announced/planned** as of the event:

| Generation | Position | Target | Status |
|---|---|---|---|
| Ascend 950PR | Inference | In production | Shipping (per Huawei) |
| Ascend 950DT | Training | Q4 2026 | Announced |
| **Ascend 960DT** | Training | **Q1 2027** (pulled in 3 quarters) | Announced |
| **Ascend 960PR** | Inference | **Q3 2027** | Announced |
| Ascend 970 | — | 2028 | Announced |
| Ascend 980 | — | 2029 | Announced |

The "pulled in 3 quarters" attached to the 960DT is doing heavy rhetorical work: it says the roadmap is accelerating, not slipping — that Huawei's execution is ahead of its own plan. In an industry where roadmaps chronically slip, claiming a pull-in is the strongest credibility signal available short of shipping. Whether the claim survives Q1 2027 is the test; in September 2026, it is a statement of confidence, recorded as such.

The claimed 960DT specifications: **2 PFLOPS FP8 / 4 PFLOPS FP4, with 288 GB of HBM**. Those are Huawei's claimed figures for unshipped silicon (see §3.15 for what they imply about the per-chip gap, and §3.13 for the SKU split).

#### The DT/PR structure

The training/inference split in the naming — DT vs. PR — mirrors the industry's workload segmentation: distinct SKUs for training (throughput-optimized) and inference (efficiency-optimized). Huawei is not just copying the cadence playbook of Nvidia and AMD; it is copying the product-line structure, down to the suffixes. The 950 generation already established the pattern (950PR in production, 950DT due Q4 2026); the 960 generation extends it on the annual rhythm.

#### The sovereignty subtext

The event's subtext was sovereignty. Every roadmap slide at Huawei Connect was implicitly addressed to two audiences: domestic customers being told they have a non-American path to scale, and foreign observers being shown that the export regime has produced a competitor with its own laws of motion — literally, its own "law." The theme, "Advancing the Agentic World," positioned Huawei's infrastructure as the substrate for the agentic era — the same era Intel's Tan invoked for the CPU (§3.10) and Qualcomm targets for inference (§3.17). Everyone in 2026 agreed on what the era is; they disagreed on who builds its foundation.

Whether "Tao's Law" survives contact with manufacturing reality is unknowable from announcements; what is knowable is that Huawei chose, in September 2026, to be judged on an annual cadence, in public, against the two companies that invented the practice. That choice is itself the news: it means Huawei believes its execution can withstand the comparison.

---

#### Huawei Connect as a venue: why September mattered

Huawei Connect is the company's flagship annual event — its equivalent of GTC or Advancing AI — and the 2026 edition's timing was pointed. It landed in the densest month of the industry's year: eleven days after Huang's "AGI" post, one day after MLPerf v6.1, two days before AMD's trillion-dollar crossing. Whether by design or by calendar collision, Huawei's roadmap shared the news cycle with the Western industry's biggest September — and held its own in it. The coverage treated Connect not as a regional event but as a global roadmap disclosure, which is itself a measure of how seriously the industry now takes Huawei's silicon program.

The three-day structure (September 17–19, Shanghai) front-loaded the substance: David Wang's keynote on day one carried the roadmap, the SuperPoD, and the PyTorch announcement — the entire §3.12–§3.16 content in a single address. The remaining days were ecosystem and partner programming. For reference purposes, September 17 is the date that matters; everything in this chapter's Huawei sections traces to that keynote unless otherwise noted.

#### The roadmap as a timeline

```
2026                2027                2028      2029
 │                   │                   │         │
950PR (shipping)     │                   │         │
950DT (Q4 2026) ───▶│                   │         │
                    960DT (Q1, pulled   │         │
                      in 3 quarters)    │         │
                    960PR (Q3) ────────▶│         │
                                        970 ─────▶980
                                        │         │
                              Atlas 960 SuperPoD │
                              delivery (Q4 2027)  │
```

The visual makes the cadence claim concrete: a new generation every year, training and inference SKUs alternating within the year, and the flagship system delivery (Q4 2027) landing on the 960 generation's hardware. "Tao's Law" is the label; the timeline is the commitment. Annual cadences are easy to announce and hard to keep — Nvidia's is the only one with a multi-year track record — so the 2028 and 2029 entries should be read as intent, the 2027 entries as the testable near term.

---

