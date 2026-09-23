---
id: etape9-phasee-storage-market/00-storage-market/e12b-checkpoint-math-for-moe-and-400b-models
title: "E12b — Checkpoint math for MoE and 400B+ models"
domain: step-9-phase-e-storage-memory-market-2026
role: deep-dive
task: architecture
actors: ["DeepSeek", "Google", "Nvidia", "Samsung", "United States"]
dates: []
keywords: ["moe", "agents", "consumer", "cost", "datacenter", "deepseek", "dram", "fine-tuning", "fp8", "gpu", "gpus", "hbm"]
source: docs/RAG/etape9_phaseE_storage_market.md
source_anchor: ""
source_lines: [354, 419]
section: "Step 9 Phase E — Storage & Memory Market 2026"
sha256: f21103bd66cdcfa90bf7a5151818b1ed58c78487b5bd540c7f17106c0ed321b2
---

# E12b — Checkpoint math for MoE and 400B+ models

## E12b — Checkpoint math for MoE and 400B+ models

- **MoE checkpoints follow total parameters, not active parameters:** a 400B-total/17B-active MoE (Llama 4 Maverick class) still checkpoints ~400B × 8–12 B ≈ **3.2–4.8 TB**; DeepSeek-V3 class (671B total) ≈ **5.4–8 TB** [secondary].
- **Node-count reference (H200 141 GB nodes, from community training guide):** 70B → 280 nodes; 120B → 480; 300B → 1,200; 700B → 2,800; 1,000B → 4,000; DeepSeek-V3 671B → 108 nodes listed (MoE efficiency) [secondary]. Source: https://github.com/yy29/ai-dev-tips-platform/blob/HEAD/manual_llm_pretraining.md
- **Storage-per-training-node rule of thumb:** 2–4× the single checkpoint size in fast NVMe per node (one live checkpoint + one staging + OS/scratch) — for a 70B run (~700 GB checkpoint), **~1.5–3 TB usable NVMe per node**; for 405B (~6 TB checkpoint), **~12–24 TB per node**, which is why 8× 3.84 TB (24 TB raw) node-local configs are the 2026 training-node standard [secondary].
- **Shared tier:** N × checkpoint-size × retention-count on the parallel file system or object store; with hourly checkpoints and 24-hour retention on a 70B run: 24 × 700 GB ≈ **17 TB/day of checkpoint write traffic** before compression/dedup [secondary].
- **Compression note:** safetensors + zstd typically yields 1.3–1.8× on BF16 weights; optimizer states compress poorly (high entropy) — budget shared-tier capacity on the uncompressed figure [secondary].

## E13b — KV-cache worked examples by model class

- **7B class (Llama-3.1-8B, BF16):** 16.1 GB weights; KV 128 KiB/token → 32K context ≈ 4 GB single-sequence; a 24 GB card holds weights + ~45K cached tokens at 90% utilization [secondary].
- **70B class (BF16):** 140 GB weights; 2× H200 141 GB (282 GB total) or 4× A100 80 GB for a healthy pool; KV at 32K ≈ 20–40 GB; 10 concurrent users ≈ 112 GB FP16 / 56 GB FP8 cache [secondary].
- **405B class (BF16):** ~810 GB weights alone — 8× H100/H200 single node minimum, and the KV pool at long context dominates: this is the NVL72-rack use case (Google's A4X class: 130 TB/s aggregate NVLink-domain bandwidth cited for GB200/GB300 NVL72) [secondary].
- **Quantization ladder for 70B:** BF16 140 GB → FP8 ~70 GB → INT4 ~35 GB weights; KV cache quantizes independently (FP8 cache halves the FP16 figure; TurboQuant-class 3–4-bit methods cut it ~5×) [secondary].
- **Context-length cost table (7B class, FP16, per sequence):** 2K → +0.5 GB; 8K → +2 GB; 32K → +8 GB; 128K → +32 GB — linear in tokens, which is why 128K+ agents "deserve one tier more" of GPU memory [secondary]. Source: https://github.com/mosesy5688-cell/ai-nexus/blob/HEAD/src/pages/knowledge/vram.md
- **Production heuristic:** allocate ~80% of VRAM to weights, reserve ~20% for KV — then verify against peak concurrency, not the average [secondary].

## E20 — Buyer FAQ (2026 market edition)

1. **Should I buy SSDs now or wait?** Every 2026 source says buy on need — indices rose 1.7–3.9% in a single September week; "the best time to buy was yesterday." No relief signal exists before mid-2027 [secondary].
2. **Is DDR5 worth it vs DDR4 for a homelab in 2026?** On used enterprise gear, DDR4 ECC RDIMM is dramatically cheaper per GB; DDR5 only if the platform requires it. Retail DDR5 at ~$18.44/GB is the single most inflated component in 2026 builds [secondary].
3. **New consumer NVMe or used enterprise SATA for a Proxmox datastore?** Used enterprise SATA/SAS at $78–127/TB beats new consumer NVMe on $/TB and endurance; new NVMe wins on latency/IOPS per drive and simplicity (no sleds/HBAs) [secondary].
4. **How much TBW do I really need?** Homelab VM/file workloads rarely exceed a few TB/month — a 1,200 TBW 2 TB TLC drive (1.6 TB/day for 2 years) is overkill; even 0.3 DWPD consumer ratings suffice [secondary].
5. **QLC for a NAS?** Yes for media (read-heavy); no for VM images, databases, or surveillance writes — TLC for anything write-sustained [secondary].
6. **DRAM-less SSD with HMB — acceptable?** Yes for mainstream/desktop use in 2026; for sustained random-write server duty, prefer DRAM-equipped or enterprise drives [secondary].
7. **Will a Dell-branded SSD work in my HPE server?** Physically maybe, but expect monitoring/fan/firmware friction — cross-vendor OEM drives are the worst of both lock-ins; prefer vendor-neutral (Samsung/Micron/Kioxia retail-enterprise) SKUs on plain HBAs [secondary].
8. **What does "Percentage Used 100" mean on NVMe SMART?** Rated endurance consumed — the drive owes you nothing further; replace on next maintenance window [secondary].
9. **Can I trust "0% wear" on a used drive?** If SMART is genuine, yes for endurance — but check power-on hours, reallocated sectors, and error logs too; wear is not the only failure axis [secondary].
10. **Is RAID a backup?** No — RAID (and ZFS redundancy) is uptime, not backup; the 3-2-1(-1-0) rule still applies, and ransomware-proofing needs immutability (see Step 7 Phase D) [secondary].
11. **Should homelabs buy Gen5 SSDs in 2026?** Only with PCIe 5.0 platforms and a bandwidth need; Gen4 TLC remains the value king [secondary].
12. **How do I size checkpoints for fine-tuning a 70B?** ~700 GB per full checkpoint (8–12 B/param); LoRA/PEFT checkpoints are megabytes — full vs adapter checkpointing changes storage by 3+ orders of magnitude [secondary].
13. **How much VRAM for a 70B at 4-bit?** ~35 GB weights + KV/context headroom → 48 GB class card (L40S/RTX 6000 Ada) is the comfortable single-card answer [secondary].
14. **Why are GPUs so expensive in 2026?** VRAM >80% of high-end GPU BOM; GDDR7 and HBM shortages transmit straight to street prices (RTX 5090 $1,999 → $4,329) [secondary].
15. **Will CXL memory help my homelab?** Not in 2026 — CXL 3.x memory is datacenter-only at cutoff; no consumer platform or price exists [unverified].
16. **Is computational storage dead?** Samsung SmartSSD is mothballed; the ideas moved into CXL, DPUs, and ZNS/FDP host-managed placement [secondary].
17. **New vs used HDDs for bulk?** Used 8 TB at ~$19/TB and recertified 16 TB at ~$41/TB are the 2026 bulk kings; run a full badblocks/SMART extended test on arrival [secondary].
18. **What about PLC NAND?** Roadmap-only at cutoff; ignore for purchasing [unverified].
19. **Are YMTC drives safe to buy?** YMTC is the #3 NAND shipper (14%, Q2 2026) with a 10–20% price edge; US-policy risk affects supply chains, not drives already on your shelf — availability varies by region [secondary].
20. **How do I spot a fake SSD?** Full-surface write test, SMART/model-string check, wear check on "new" units, weight/inspection vs teardowns, traceable sellers [secondary].
21. **E1.S vs U.2 for a new build?** E1.S is the hyperscale 2026+ form factor; U.2 has the mature used market and adapter ecosystem — homelabs should stay U.2/M.2 in 2026 [secondary].
22. **Does HBM pricing affect me?** Indirectly — every HBM wafer tightens DDR5/GDDR supply; your DDR5 kit price is downstream of the HBM allocation war [secondary].
23. **What's the $/GB ladder in 2026?** NAND retail ~$0.16 → used enterprise flash ~$0.08–0.13 → new enterprise flash ~$0.30–1.17 → DDR5 retail ~$18.44 → GDDR (embedded in GPU BOM) → HBM spot ~$58–97 [secondary].
24. **One-sentence 2026 strategy?** Buy flash for the workload you have (TLC for writes, QLC/HDD for bulk), verify every used drive with smartctl, and budget for DRAM — it's the scarcest dollar in the build [secondary].

## E21 — Decision matrices

**Drive choice matrix (homelab/small-business, 2026):**

| Workload | Buy | NAND | Form | Why |
|---|---|---|---|---|
| Proxmox VM datastore | Used enterprise SATA/SAS 3.84 TB | TLC, 1 DWPD RI | 2.5" | $78–127/TB, endurance verified via SMART |
| ZFS SLOG / DB WAL | New enterprise NVMe (or Optane-class if found used) | TLC, high DWPD | U.2/M.2 | Write-heavy tier = buy new (E9) |
| Media library (Jellyfin/Plex) | Used HDD 8–16 TB | n/a | 3.5" | $19–41/TB; SSD premium wasted on sequential reads |
| Boot / OS | Small used enterprise SATA or new consumer NVMe | Any | 2.5"/M.2 | Don't pay $/TB premiums at 240–500 GB |
| Laptop/desktop | New Gen4 TLC NVMe 1–2 TB | TLC + HMB/DRAM | M.2 | $65–140/TB, warranty, no adapter friction |
| AI dataset staging (warm) | New QLC NVMe or used enterprise TLC | QLC acceptable | U.2/M.2 | Read-heavy, capacity-first |

**Memory choice matrix (2026):**

| Need | Buy | $/GB anchor | Caveat |
|---|---|---|---|
| Homelab server RAM | Used DDR4 ECC RDIMM | ~$1.5–3/GB (eBay-class) | Match rank/voltage to platform |
| New workstation | DDR5 retail kits | ~$18.44/GB (Sept 2026) | The most inflated component; buy only what the board needs |
| GPU inference memory | More VRAM, not faster VRAM | Embedded in GPU street price | VRAM >80% of high-end GPU BOM in 2026 |
| Future expansion | Wait for CXL (datacenter) | n/a | No consumer CXL memory market at cutoff |

