---
id: etape9-phasec-memory/00-memory/12-hbm2e-in-2026-legacy-but-still-deployed
title: "12. HBM2e in 2026: legacy but still deployed"
domain: step-9-phase-c-server-accelerator-memory-hardware
role: deep-dive
task: hardware
actors: ["AMD", "AWS", "DeepSeek", "Google", "Intel", "Meta", "Nvidia", "Samsung", "TSMC"]
dates: ["2025-04-16", "2026-03", "2026-08", "2026-08-26"]
keywords: ["hbm", "accelerator", "amd", "blackwell", "cost", "deepseek", "disaggregated", "dram", "foundry", "fp4", "fp8", "gpu"]
source: docs/RAG/etape9_phaseC_memory.md
source_anchor: ""
source_lines: [226, 294]
section: "Step 9 — Phase C: Server & Accelerator Memory Hardware"
sha256: c130ad600e4f858252476d803005f7c580a0fd8289034f2e7c13d275da10fe16
---

# 12. HBM2e in 2026: legacy but still deployed

## 12. HBM2e in 2026: legacy but still deployed

- HBM2e remains in production accelerators: Intel Gaudi 3 ships 128GB of HBM2e (see Section 17); NVIDIA
  A100 80GB uses HBM2e at ~2.039 TB/s and remains a common workhorse card [secondary].
- Gaudi 2 used 96GB HBM2e at 2.45 TB/s; Gaudi 3 moved to 128GB HBM2e at up to 3.7 TB/s [official].

## 13. HBM3 / HBM3e: the 2023–2026 workhorse

- NVIDIA H100 (Hopper): 80GB HBM3 at 3.35 TB/s (SXM5); H100 NVL PCIe variant offers 94GB [secondary].
- NVIDIA H200: 141GB HBM3e at 4.8 TB/s (six HBM3e stacks per GPU) — the memory-bound LLM inference choice
  on Hopper-class hardware [secondary].
- AMD MI300X (CDNA 3): 192GB HBM3 at 5.3 TB/s, 750W [secondary].
- AMD MI325X (CDNA 3): 256GB HBM3e at 6.0 TB/s (a 32GB reduction from the originally announced 288GB,
  cut to address the market opportunity) [secondary].
- AMD MI350X / MI355X (CDNA 4): 288GB HBM3e at 8.0 TB/s; MI350X air-cooled 1000W, MI355X liquid-cooled
  up to 1400W [secondary].
- HBM3e was the dominant AI accelerator memory standard through 2025 and into 2026; its shortage traces
  to the same wafer-allocation dynamics as HBM4 — demand (e.g., H200's six stacks per unit) outstripped
  capacity even as HBM4 ramped alongside it [secondary].
- Custom ASICs from Google, Amazon and Meta generally stayed on HBM3e through their current shipping
  generations, planning HBM4 for future iterations [secondary].

## 14. HBM4: the 2026 transition

- JEDEC finalized the HBM4 standard (JESD270-4) on April 16, 2025: 2048-bit interface, 32 channels, up to
  2 TB/s per stack [secondary].
- Samsung's HBM4: 36GB modules hitting ~3.3 TB/s via redesigned stacking architecture and advanced signal
  correction; Samsung applied 6th-gen (1c) DRAM technology ahead of competitors [secondary].
- Cost shock: Samsung's HBM4 reportedly costs Nvidia over **$500 per module**, more than double the
  ~$250 for HBM3e; SK hynix reportedly charges ~$550, partly due to TSMC base-die supply costs; Korean
  press estimates HBM4 at ~$600+ per unit [secondary].
- **Base-die outsourcing is the structural shift**: SK hynix outsources HBM4 base dies to TSMC (12nm-class);
  Samsung runs a dual track (in-house foundry + TSMC, custom HBM announced Sept 15); Micron outsources to
  TSMC starting with HBM4E (mass production targeted 2027). All three major makers now rely on TSMC for
  base dies [secondary].
- Driver: as HBM moves from standardized products to customer-specific solutions, logic in the base die
  keeps growing — NVIDIA's **NVHBM** (unveiled August 26, 2026) puts the memory controller inside the
  base die rather than on the XPU: claimed up to 30% higher bandwidth, 15% lower HBM power, 25% of XPU
  die area freed; Amazon's Annapurna Labs is first partner (Trainium4 with NVLink Fusion) [secondary].
- HBM4e previewed March 2026 (Samsung): 16 Gb/s, 2048-bit, 4.0 TB/s per stack [secondary].
- Spot-market snapshot (late August 2026): 36GB HBM3e at $2,100 (4–5× long-term contract price of
  ~500,000–700,000 won); 16-layer HBM4 in pre-mass-production at $3,500 on spot vs LTA [secondary].

## 15. NVIDIA accelerators: memory map

| Accelerator | Memory | Bandwidth | Notes |
|---|---|---|---|
| A100 80GB (Ampere) | 80GB HBM2e | ~2.039 TB/s | Still common; 400W-class PCIe/SXM |
| H100 SXM5 (Hopper) | 80GB HBM3 | 3.35 TB/s | Transformer Engine, FP8; 700W |
| H100 NVL | 2× 94GB HBM3 | ~3.9 TB/s each | PCIe inference variant |
| H200 SXM (Hopper) | 141GB HBM3e | 4.8 TB/s | 6 HBM3e stacks; memory-bound inference pick |
| B200 (Blackwell) | 192GB HBM3e (8-hi) | 8 TB/s | 208B transistors, dual-die; 1000W; FP4 |
| B300 (Blackwell Ultra) | 288GB HBM3e (12-hi) | ~8 TB/s | ~1400W; liquid strongly assumed |
| GB200 NVL72 | 13.4 TB HBM3e (72 GPUs) | 576 TB/s aggregate | Rack-scale; 120–130 kW/rack |
| GB300 NVL72 | B300-based | — | 45% higher DeepSeek-R1 inference vs GB200 (MLPerf Inference v5.1) [vendor-reported] |
| Rubin (H2 2026) | 288GB HBM4 / GPU | 13 TB/s | Vera Rubin NVL144: 75TB fast memory, 13 TB/s HBM4 BW |
| Rubin Ultra (H2 2027) | HBM4e | 4.6 PB/s | 576 GPUs, 2304 HBM4e stacks, 150TB |
| Rubin CPX | 128GB **GDDR7** | — | Purpose-built for massive-context *prefill*; disaggregated inference |
[secondary]

- Pricing signals (new purchase, 2026): H100 $25,000–40,000; H200 $30,000–40,000; B200 $40,000–45,000;
  GB200 NVL72 $2–3M per rack (~$28,000–42,000 per GPU equivalent) [secondary]. Cloud on-demand medians:
  H100 ~$2.29/hr, H200 ~$3.95/hr, B200 $3.75–9.86/hr [secondary].
- B200 detail: derived per-GPU from DGX B200 totals (1440GB memory, 64 TB/s HBM BW across 8 GPUs);
  some sources quote 180GB vs 192GB per B200 — the DGX-derived 192GB figure is used here; treat the
  180GB figure as [unverified].
- HBM is now the dominant cost: HBM accounts for ~50–60% of an AI GPU's bill of materials (2026
  estimate) [secondary].

