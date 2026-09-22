---
id: briefing-general-tech-2026/02-gpus-accelerators/15-per-chip-gap
title: "The per-chip gap: ~2 years behind Blackwell, ~10x under Rubin"
domain: gpus-accelerators
role: deep-dive
task: hardware
actors: ["AMD", "China", "Huawei", "Nvidia"]
dates: ["2026-09-17"]
keywords: ["blackwell", "ascend", "fp4", "fp8", "gpu", "helios", "inference", "npo", "optics", "packaging", "rack-scale", "superpod"]
source: docs/RAG/briefing-general-tech-2026-en.md
source_anchor: "#g03-15"
source_lines: [2395, 2457]
canonical_for: ["huawei-ascend"]
sha256: ba7cd2e76a264b3e5b5dfbd612dca180640156ca68ef19e4b76c870e860ba8fc
---

# The per-chip gap: ~2 years behind Blackwell, ~10x under Rubin

<a id="g03-15"></a>
### 3.15 The per-chip gap: ~2 years behind Blackwell, ~10x under Rubin

System-level ambition does not erase the per-chip gap, and the most clear-eyed assessment of it came from **The Register on September 17, 2026**: the Ascend **960DT** delivers roughly **one-half the FP8 and one-third the FP4** of Nvidia's **B300** — placing it, per chip, in the **Blackwell class**, approximately **two years behind** the leading edge. Against Rubin, the arithmetic is starker: with Rubin claimed at **35–50 PFLOPS FP4** per GPU against the 960DT's claimed 4 PFLOPS FP4, **one Rubin is roughly 10x one 960DT** on paper.

| Comparison (claimed specs, no independent benchmarks) | Ratio |
|---|---|
| Ascend 960DT vs Nvidia B300 (FP8) | ~1/2 |
| Ascend 960DT vs Nvidia B300 (FP4) | ~1/3 |
| 1× Rubin GPU (35–50 PFLOPS FP4 claimed) vs 1× 960DT (4 PFLOPS FP4 claimed) | ~10x |
| Implied per-chip generation gap | ~2 years (Blackwell-class) |

Every number in that table is a claimed spec, not a measurement: the 960DT does not ship until Q1 2027 at the earliest, and **no independent 960 benchmarks exist**. The gap is therefore a paper gap — but paper gaps are what roadmaps are for, and Huawei's paper shows a chip two years behind per unit of compute. Note also the asymmetry *within* the gap: 1/2 in FP8 but 1/3 in FP4 against the B300 suggests the deficit widens at lower precision — the precision regime (FP4) where the industry's efficiency gains are concentrated. That is the least favorable possible shape for the gap.

#### "Blackwell-class": what the label means for buyers

Calling the 960DT "Blackwell-class" is a statement about *which* Nvidia generation Huawei competes with, and it has a concrete buyer implication: a Blackwell-class chip in 2026–2027 competes for the workloads Blackwell serves today — large-model training and inference at datacenter scale — but at 2024-era efficiency. For Chinese buyers with no access to Blackwell, that is a viable proposition: Blackwell-class domestic silicon beats no silicon. For the global narrative, it sets the bar Huawei must clear next: "Blackwell-class, shipping domestically" in 2027, "Rubin-class" some year after. The two-year gap is not a verdict but a starting position — and Huawei's entire system-level strategy (§3.14–§3.15) is the attempt to make the starting position matter less than the finish.

#### Huawei's counter-thesis: the system, not the chip

Huawei's counterargument — and it is the central thesis of the company's whole 2026 presentation — is that the per-chip comparison is the wrong comparison. **Morgan Stanley's Charlie Chan**, quoted by Bloomberg on September 17, 2026, gave the thesis its most quotable form: **"System-level competitiveness matters more than ever. The effective gap is narrowing through multi-die design, advanced packaging, rack-scale system architecture, optical networking, and software-hardware co-optimization."**

The five items in that list deserve item-by-item mapping, because they are a precise inventory of §3.14:

| Chan's item | Huawei's 2026 evidence |
|---|---|
| Multi-die design | 960DT packaging (claimed) |
| Advanced packaging | 3D/integration approach across the portfolio |
| Rack-scale system architecture | Atlas 960 SuperPoD: 15,488 chips, 220 cabinets |
| Optical networking | NPO Hi-ONE: 5,500 modules, >550 kW saved |
| Software-hardware co-optimization | PyTorch backend recognition; UnifiedBus fabric |

Huawei's strategy is to concede the chip and win the system: 15,488 weaker chips, bound by NPO optics and a unified fabric, presented as competitive with fewer stronger ones. It is, in essence, the same scale-up philosophy Nvidia and AMD are pursuing with NVL72 and Helios — a single coherent compute domain compensating for per-unit limits — applied from a weaker starting chip.

#### The honest verdict

Whether the arithmetic of compensation actually closes a 10x per-chip gap is unknowable without independent benchmarks, and skepticism is warranted: system-level tricks have overheads, yields, and power costs of their own. Fifteen thousand chips need fifteen thousand chips' worth of power, cooling, and failure handling — the 99.8% availability claim (§3.14) is Huawei's answer to the reliability half of that objection, and the >550 kW optical saving is its answer to the power half. But answers on paper are not answers in production.

The honest 2026 verdict, in one paragraph: Huawei is approximately two years behind per chip, claims to be narrowing the *effective* gap at system level through the five mechanisms Chan listed, and will be judged on the Q4 2027 SuperPoD delivery — the first date at which the claim becomes testable. Until then, the gap is paper-versus-paper: Huawei's claimed 4 PFLOPS FP4 against Nvidia's claimed 35–50, with the system architecture as Huawei's stated equalizer. The Register's "two years behind" is the bear case written as fact; Chan's "the effective gap is narrowing" is the bull case written as analysis. The dossier records both, and waits for silicon.

---

#### The arithmetic, shown fully

| Figure | Value | Status |
|---|---|---|
| Ascend 960DT, FP8 | 2 PFLOPS (claimed) | Huawei claim |
| Ascend 960DT, FP4 | 4 PFLOPS (claimed) | Huawei claim |
| Nvidia B300, FP8 | ~4 PFLOPS (implied by "1/2" ratio) | Implied by The Register's ratio |
| Nvidia B300, FP4 | ~12 PFLOPS (implied by "1/3" ratio) | Implied by The Register's ratio |
| Rubin GPU, FP4 | 35–50 PFLOPS (claimed) | Nvidia claim |
| 960DT vs B300 (FP8) | ~1/2 | The Register, 17/09/2026 |
| 960DT vs B300 (FP4) | ~1/3 | The Register, 17/09/2026 |
| Rubin vs 960DT (FP4) | ~9–12x (≈10x) | Derived from claimed specs |

The derived rows are arithmetic on claimed specs, not measurements — the table labels them as such. The "~10x" is a rounding of a 9–12x range, and the dossier uses it as shorthand with the range available here for precision. Note what the arithmetic cannot do: it cannot account for software efficiency, memory-system differences, or the system-level compensation of §3.14. Per-chip FLOPS ratios are the beginning of a competitiveness analysis, not the end — which is exactly Chan's point.

#### On paper gaps: a methodological note

"Two years behind" is a generation-counting statement: the 960DT's claimed per-chip throughput resembles the B300's class, and the B300 is roughly two years old relative to Rubin's 2026. Generation-counting is the industry's standard shorthand, but it has known weaknesses: it ignores the *rate* of improvement (a fast follower can close two years in one generation), it ignores system-level compensation (the whole of §3.14), and it treats the leader's roadmap as fixed (it isn't — Rubin Ultra is coming). The dossier uses "two years behind" because The Register's analysis supports it, but readers should treat it as a snapshot of claimed specs, not a forecast. Forecasts require the one thing nobody has: measured 960DT silicon.

---

