---
id: etape5-trackb-amd/00-amd/10-open-verification-items
title: "10. OPEN VERIFICATION ITEMS"
domain: step-5-track-b-amd-hardware-instinct-gpus-epyc-ryzen-cpus
role: deep-dive
task: hardware
actors: ["AMD", "Google", "Meta", "Microsoft", "Nvidia", "OpenAI", "Oracle", "TSMC"]
dates: ["2026-09-22"]
keywords: ["2nm", "amd", "benchmarks", "compute", "copilot", "fp4", "gpu", "hbm3", "hbm4", "helios", "lpddr5x", "memory"]
source: docs/RAG/etape5_trackB_amd.md
source_anchor: ""
source_lines: [220, 260]
section: "Step 5 — Track B: AMD Hardware (Instinct GPUs + EPYC/Ryzen CPUs)"
sha256: ca61ed8ec387d9a5fff59869e4dc6eeae5492fc3461e275b5c32ed3742bef2b8
---

# 10. OPEN VERIFICATION ITEMS

| Product | Architecture / Process | Memory | Peak compute (dense) | Power | Status Sep 2026 |
|---|---|---|---|---|---|
| MI300X | CDNA 3 | 192 GB HBM3, 5.3 TB/s | ~1.3 PFLOPS FP16 | 750 W | Mature; ~$6/hr cloud; full-capacity pricing signals |
| MI325X | CDNA 3 | 256 GB HBM3E, 6 TB/s | 2.61 PFLOPS FP16 | 1,000 W | Shipping |
| MI350X | CDNA 4 / TSMC N3P | 288 GB HBM3E, 8 TB/s | 18.4 PFLOPS FP4 | 1,000 W | GA H2 2025; $8.60/hr OCI |
| MI355X | CDNA 4 / TSMC N3P | 288 GB HBM3E, 8 TB/s | 20.1 PFLOPS FP4 | 1,400 W | Flagship shipping; Oracle 30k order |
| MI455X | CDNA 5 / TSMC 2nm | 432 GB HBM4, 19.6 TB/s | 40 PFLOPS FP4 | TBD (liquid) | Volume H2 2026 |
| EPYC 9965 (Turin) | Zen 5c | DDR5-6400 12ch | 192c/384t | 500 W | Adopted (Google, OVHcloud) |
| EPYC Venice (9006) | Zen 6 / TSMC 2nm | DDR5-8000 16ch | 256c/512t | TBD | Volume production from Jul 2026 |
| Ryzen AI 9 HX 475 | Zen 5/5c + RDNA 3.5 + XDNA 2 (60 TOPS) | LPDDR5X-8533 | 12c/24t | 28 W | Announced CES 2026 (Gorgon Point) |
| Ryzen AI Max+ 395 | Zen 5 + RDNA 3.5 40CU + XDNA 2 (50 TOPS) | 128 GB unified | 16c/32t | 45–120 W | Shipping; ~30 mini-PC/laptop designs |

---

## 10. OPEN VERIFICATION ITEMS

1. **MI350 vs B200 head-to-heads** — AMD's 10–30% claims are vendor-reported; independent third-party benchmarks at scale remain thin as of Sep 22, 2026.
2. **MI300X/MI350 list prices** — AMD does not publish; all figures are cloud $/hr or community estimates. Older "~$15,000" claims have no verified 2026 source.
3. **MI400/MI455X final specs and exact volume timing** — H2 2026 volume is AMD's stated schedule; actual ramp and TDP figures unconfirmed.
4. **MI430X** — sovereign-AI/HPC variant; specs not published by AMD as of Sep 2026.
5. **"MI450" vs "MI455X" naming** — sources mix the two; MI455X is the flagship rack part; Meta's "custom MI450" variant details unconfirmed.
6. **MI500 (2027)** — announced at Analyst Day; no specs.
7. **Helios bandwidth figures** — 260 TBps (Register) vs 43 TB/s scale-out (Cirrascale/BusinessWire) vs per-GPU 1.8 TB/s/dir (STH): marketing maxima; apples-to-oranges.
8. **Oracle 30,000 MI355X order** — from a 2025 earnings call via secondary reporting; fulfillment status unconfirmed.
9. **Venice pricing** — no MSRP published as of Sep 2026.
10. **Turin 2026 pricing trends** — no 2026 SKU/pricing updates located in this research.
11. **"First Copilot+ desktop CPU"** — thinly sourced; AMD's Gorgon Point desktop APUs' certification details unconfirmed.
12. **Gorgon Halo / Medusa Halo** — rumor-grade (community research repo).
13. **Venice CCD core type** — resolved post-Advancing AI 2026 (Zen 6 / Zen 6c split), but per-SKU core maps incomplete.
14. **Wolfe Research $90–120B Meta-deal estimate** — analyst projection, not committed revenue.
15. **Spheron $3.59/hr MI300X vs $2.65/hr H100** — live marketplace snapshot, Sep 22, 2026; supply-driven, moves constantly.

---

## 11. COLLECTION METADATA

- **Research date:** September 22, 2026.
- **Search passes:** 8 web-search batches covering MI300/MI350 specs & benchmarks, MI400/Helios architecture, Oracle/Meta/OpenAI partnerships, EPYC Venice announcement, Ryzen Gorgon Point & Strix Halo, MI300X cloud pricing.
- **Source mix:** AMD announcements via press coverage (official-by-report), The Register, ServeTheHome, Tom's Hardware, TechPowerUp, CRN, TweakTown, Notebookcheck, HotHardware, BusinessWire (Cirrascale), plus secondary aggregators and community research repos (flagged as such).
- **Cross-references:** ROCm stack → `etape4_trackC_cuda_rocm_pytorch.md` (Step 4 Track C); NVIDIA comparison data → Step 3 Track D.
- **Nothing was sent externally; no live-browser visits; read-only web research + local file write.**
