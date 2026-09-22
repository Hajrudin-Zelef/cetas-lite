---
id: briefing-general-tech-2026/02-gpus-accelerators/13-ascend-960dt-960pr
title: "Ascend 960DT vs 960PR: training and inference SKUs"
domain: gpus-accelerators
role: deep-dive
task: hardware
actors: ["AMD", "China", "Huawei", "Intel", "Nvidia"]
dates: ["2026-09-22"]
keywords: ["ascend", "inference", "training", "accelerator", "crescent island", "fp4", "fp8", "hbm", "helios", "kv cache", "superpod", "tokens per dollar"]
source: docs/RAG/briefing-general-tech-2026-en.md
source_anchor: "#g03-13"
source_lines: [2263, 2320]
canonical_for: ["huawei-ascend"]
sha256: 3f14da6ec183f92784e9eed2bcd3c7976df6226b17b4547bd32e5902855a0acd
---

# Ascend 960DT vs 960PR: training and inference SKUs

<a id="g03-13"></a>
### 3.13 Ascend 960DT vs 960PR: training and inference SKUs

Huawei's 960 generation splits, like its 950 predecessor, into two SKUs: the **960DT for training** and the **960PR for inference**. The schedules differ by two quarters — **Q1 2027 for the 960DT, Q3 2027 for the 960PR** — with the training chip pulled in by three quarters and the inference chip following the original plan.

#### What is verified — and what is not

What the verified record actually contains about the two SKUs is asymmetric, and the dossier should not pretend otherwise:

| SKU | Position | Schedule | Specs disclosed? |
|---|---|---|---|
| 960DT | Training | Q1 2027 (pulled in 3 quarters) | Yes — 2 PFLOPS FP8 / 4 PFLOPS FP4, 288 GB HBM (claimed) |
| 960PR | Inference | Q3 2027 | **No** — positioning and schedule only |

The honest summary is that Huawei has announced the *structure* of the 960 generation — two SKUs, two schedules, one claimed spec sheet for the training part — but not the inference part's technical details. The 960PR's power envelope, memory configuration, and precision support are all undisclosed as of September 22, 2026.

#### 288 GB HBM: reading the 960DT's memory figure

The 960DT's claimed 288 GB of HBM per chip is the most informative number in Huawei's disclosed specs, because memory capacity per chip determines what a single accelerator can hold without distribution overhead. At 288 GB, a single 960DT can hold frontier-class models at reduced precision with KV-cache headroom to spare — the per-chip capacity is competitive even if the per-chip FLOPS lag. The figure also illuminates the system strategy: 4,096 NPUs × 288 GB ≈ 1.18 PB, consistent with the 960E SuperPoD's stated 1 PB of HBM (allowing for system overhead) — the chip and system numbers corroborate each other, which is a modest but real credibility signal for Huawei's disclosure. When the 960PR's memory figure eventually appears, compare it here: inference SKUs live or die on memory capacity per watt.

#### Why the missing 960PR spec sheet matters

Given that inference is where Huawei claims its system-level advantages (see §3.14's 2.5x inference claim for the SuperPoD), the missing 960PR spec sheet is the most conspicuous gap in the Connect announcements — and the first thing to check when Huawei next discloses. An inference SKU's competitiveness turns on exactly the numbers Huawei hasn't shared: memory capacity per chip (KV cache headroom), memory bandwidth (tokens per second), and power (tokens per watt). The 960DT's 288 GB HBM figure gives analysts something to model; the 960PR gives them nothing.

One plausible reading of the asymmetry: the training chip's specs are the headline because training is where the per-chip gap (§3.15) is most visible and therefore most in need of a strong number, while the inference chip's story is a system story (the SuperPoD's 2.5x claim) that doesn't need per-chip specs to be told. That reading is interpretation, not fact — but it fits the shape of what was disclosed and what wasn't.

#### The industry-wide convergence on the split

The DT/PR split itself is worth a final note: it is the same segmentation every vendor in this chapter has converged on. Nvidia has distinct training and inference configurations of its platforms; AMD's MI450 family spans both; Intel built Crescent Island as inference-only (§3.11). The industry has decided that training and inference are different products, with different optimization targets and different buyers. Huawei's roadmap is, in this respect, fully conventional — which is itself a statement about how far China's accelerator industry has come: it now segments its product line the way the incumbents do, because it faces the same workload economics.

---

#### The 950 generation: the pattern being extended

| SKU | Position | Status (Sept 2026) |
|---|---|---|
| Ascend 950PR | Inference | In production (per Huawei) |
| Ascend 950DT | Training | Q4 2026 (announced) |
| Ascend 960DT | Training | Q1 2027 (announced, pulled in 3 quarters) |
| Ascend 960PR | Inference | Q3 2027 (announced) |

The 960 split is not an innovation but an inheritance: the 950 generation already established the DT/PR structure, with the inference SKU (950PR) shipping before the training SKU (950DT) — the reverse of the 960 order, where the training chip was pulled in ahead of the inference one. The inversion is worth noting: it suggests the pull-in was specific to the training part (perhaps driven by domestic training demand, perhaps by competitive pressure from the Rubin/Helios training narratives) rather than a uniform acceleration of the whole generation. The inference SKU kept its original schedule — which makes the missing 960PR spec sheet (§3.13) a matter of disclosure timing as much as substance.

#### Training vs. inference SKUs: the industry logic

The DT/PR split rests on a workload distinction the whole industry had converged on by 2026:

| Dimension | Training SKU (DT) | Inference SKU (PR) |
|---|---|---|
| Optimized for | Throughput on massive parallel workloads | Efficiency per token; capacity for KV cache |
| Key metric | Time-to-train; FLOPS utilization | Tokens per dollar; tokens per watt |
| Precision emphasis | FP8 (and FP4 emerging) | FP4, low-bit inference |
| Buyer | Labs training frontier models | Clouds serving models at scale |

The table is general industry logic, not Huawei disclosure — but it explains why every vendor segments this way, and why the inference SKU's undisclosed specs matter more than the training SKU's disclosed ones for evaluating Huawei's system claims. Training specs win headlines; inference specs win deployments. Huawei disclosed the former and withheld the latter.

---

