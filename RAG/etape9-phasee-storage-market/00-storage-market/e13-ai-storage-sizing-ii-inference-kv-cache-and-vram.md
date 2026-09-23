---
id: etape9-phasee-storage-market/00-storage-market/e13-ai-storage-sizing-ii-inference-kv-cache-and-vram
title: "E13 — AI storage sizing II: inference KV-cache and VRAM"
domain: step-9-phase-e-storage-memory-market-2026
role: deep-dive
task: reference
actors: ["AMD", "AWS", "Google", "Meta", "Nvidia", "Samsung", "United States"]
dates: ["2026-07"]
keywords: ["inference", "accelerator", "amd", "attention", "compute", "consumer", "cost", "datacenter", "dram", "energy", "fp8", "gpu"]
source: docs/RAG/etape9_phaseE_storage_market.md
source_anchor: ""
source_lines: [171, 237]
section: "Step 9 Phase E — Storage & Memory Market 2026"
sha256: 314b9d95b74a9c72e8804c069ed81042c1a22d91b5fe98c87be0b6f335765d8a
---

# E13 — AI storage sizing II: inference KV-cache and VRAM

## E13 — AI storage sizing II: inference KV-cache and VRAM

- **The baseline industry formula:** Memory (GB) = Parameters (B) × Bytes/param × (1 + Overhead), with **20–50% overhead** for KV cache, activations, and framework buffers [secondary]. Source: http://gpusmith.com/articles/en/pdfs/llm-inference-hardware-sizing-guide.pdf
- **Weight footprints (BF16/FP16):** 7B → 14 GB; 13B → 26 GB; 34B → 68 GB; **70B → 140 GB**; INT8 halves, 4-bit quarters (70B → 35 GB at 4-bit) [secondary]. Source: https://github.com/mosesy5688-cell/ai-nexus/blob/HEAD/src/pages/knowledge/vram.md
- **KV-cache mechanics:** per token, KV bytes = 2 (K+V) × layers × KV-heads × head_dim × bytes/value. Worked example — **Llama-3.1-8B-Instruct BF16: 128 KiB/token**, so 16K context ≈ 2 GB for one sequence; on a 24 GB card, 0.90×24 − 16.1 ≈ **5.5 GB KV pool ≈ ~45K cached tokens** [secondary]. Source: https://github.com/cachebox-project/inference-cache/blob/HEAD/docs/reference-stack/GPU-RUNBOOK.md
- **70B-class KV figures:** **20–40 GB at 32K context** (FP16, GQA); with 10 concurrent users at 32K, ~112 GB FP16 cache (56 GB FP8) [secondary]. GQA (grouped-query attention) saves up to ~4× vs old MHA designs [secondary]. Source: https://github.com/tieubao/til/blob/HEAD/notes/ai/turboquant-kv-cache-compression.md
- **Extreme-context example:** Llama-3.1 70B FP16 chatbot, 16 users × 128K context: **140 GB weights + 16 × ~38 GB KV ≈ 748 GB total** — this is why long-context serving is a rack-scale (NVL72-class) problem [secondary]. Source: https://medium.com/@donmccasland_57353/strategic-gpu-selection-for-llm-serving-on-google-cloud-a-guide-to-vram-planning-and-82d97de2870c
- **Compression relief (2026):** Google Research's **TurboQuant** (ICLR 2026) compresses KV cache 16-bit → 3–4 bits with near-zero quality loss (~5×) — directly targeting what Jensen Huang called the #1 long-context bottleneck at GTC 2026 [secondary].
- **Offload relief:** LMCache-style tiering moves KV blocks to **host DRAM** — it relieves GPU KV pressure but *adds a host-memory requirement*; it does not shrink the weights footprint [secondary].
- **GPU-per-model table (healthy KV pool, BF16):** ~8B → 24 GB (1 card); 13–34B → 48–80 GB (1 card); **70B → 160–200 GB (2–4 cards)**; 70B FP8/INT4 → 48–80 GB (1–2 cards); 100B+/MoE → 320 GB+ (8 cards) [secondary]. Card count is driven first by *weights fitting*, then by KV pool/concurrency; multi-card needs NVLink/NVSwitch (bare metal, not multi-GPU VMs) [secondary]. Source: https://github.com/cachebox-project/inference-cache/blob/HEAD/docs/reference-stack/GPU-RUNBOOK.md
- **Local-inference buying rules (Memeburn, Sept 2026):** 8–12 GB GPU memory for light chat; **24 GB VRAM (or 32–48 GB unified) for a serious coding assistant**; ~48–64 GB accelerator memory for a good 70B without heavy CPU offload; **96–128 GB for gpt-oss-120b-class / large MoE**; "a model that barely fits is already telling you the machine is too small" [secondary]. Source: https://memeburn.com/how-much-ram-and-vram-do-you-actually-need-for-local-ai-in-2026/

## E14 — Cost per GB: VRAM vs HBM vs DDR5 vs NAND

- **Retail DDR5 (Sept 2026): ~$18.44/GB** (MemoryPriceChart basket median) [secondary].
- **Retail NAND (Sept 2026): ~$0.16/GB** ($159.94/TB mainstream NVMe) — a **~115× gap** between retail DRAM and retail NAND per GB [secondary].
- **Enterprise SSD (Q3 2026): ~$0.75/GB** ($753/TB for 30 TB TLC) vs **enterprise HDD ~$0.04/GB** ($40.5/TB) [secondary].
- **Used enterprise SSD: ~$0.08–0.13/GB** ($78–127/TB) — roughly **6–9× cheaper per GB than new enterprise NVMe** ($0.30–1.17/GB) [secondary].
- **HBM (spot, Aug 2026): ~$58/GB contract-class (36 GB stack at $2,100 spot) up to ~$97/GB (HBM4 16-layer at $3,500)** — **3,000–5,000× retail DDR5 per GB**, priced as strategic allocation [secondary].
- **The GPU memory tax:** GDDR7-driven BOM inflation pushed the RTX 5090 from $1,999 MSRP to $4,329 street (mid-July 2026); VRAM is >80% of high-end GPU BOM [secondary]. For inference buyers, the $/GB ladder is: NAND < used enterprise flash < new enterprise flash < DDR5 < GDDR < HBM — and each rung is 2–100× the previous [secondary].
- **Why this matters for AI infra planning:** training clusters are sized by HBM capacity per GPU and checkpoint *bandwidth*; inference fleets are sized by **weights + KV cache per card**; bulk dataset storage stays on the cheapest tier that meets the ingest rate (HDD/object for cold, QLC/new-NVMe for warm, node-local NVMe for hot) [secondary].

## E15 — Computational storage and open firmware: 2026 status

- **Samsung SmartSSD: effectively mothballed.** The concept (2018) put NAND + HBM + RDIMM next to an AMD Xilinx FPGA inside the SSD for server-less compute; Gen2 launched 2022 [secondary]. By 2025 it had "all but disappeared from Samsung's portfolio" — still buyable on Amazon under the **AMD Xilinx brand at $517.70 for 3.84 TB**, but a Gen3 device whose novelty and complexity made it a hard sell; COVID-19 then generative AI (which demanded *capacity*, not in-drive compute) killed the business case [secondary]. Source: https://www.techradar.com/pro/samsung-and-amd-made-a-revolutionary-ssd-together-then-it-was-left-to-wither-in-the-shadows-and-nobody-knows-exactly-why
- **Assessment:** computational storage devices (CSDs) were "an interesting but niche market, closer to traditional servers" — nice, but without AI-hardware growth potential; Samsung mothballed after Gen2 despite 2022 claims of "great potential" [secondary].
- **What replaced the CSD story:** (1) **CXL-attached memory/NAND** (see E6) as the industry's chosen "compute near data" vehicle; (2) **DPUs/SmartNICs** doing storage offload on the network side (see Step 6 Phase F1); (3) **in-drive AI for the drive itself** (predictive failure, ZNS/FDP placement) rather than general compute [secondary].
- **NGD Systems:** the other notable CSD vendor — no 2026 product/market signal surfaced in this research pass; treat as dormant-or-acquired, flagged as a gap [unverified].
- **Open-source SSD firmware (OpenSSD / Cosmos+ OpenSSD):** the community FPGA-based open firmware platform remains a research/education vehicle (universities, FTL research); no evidence of 2026 commercial adoption or a new platform revision in this pass — gap flagged [unverified].
- **Open-channel SSDs / ZNS as the "open" successor:** the industry's openness energy moved from open firmware to **open interfaces** — NVMe ZNS (zone namespaces) and FDP (flexible data placement) expose placement control to the host without opening the FTL; hyperscalers (Meta, Google) drive ZNS adoption for write-amplification reduction. Details are in the sibling SSD-hardware file [secondary].

## E16 — Enterprise SSD roadmap: PCIe Gen6 and beyond

- **Samsung PM1763 (Gen6):** launching **early 2026**, ~2× performance of Gen5, ~25W power envelope; Samsung also rolling out **256 TB Gen5** variants and targeting **512 TB Gen6 in EDSFF 1T form factor around 2027** at 28–32 GB/s [secondary]. Source: https://wccftech.com/samsung-512-tb-pcie-gen6-ssds-2027-innogrit-preps-gen6-ai-nvme-cxl-enterprise/
- **Micron 9650:** described as the world's first Gen6 SSD, up to ~28 GB/s [secondary]. **InnoGrit:** developing Gen6 AI-enterprise SSDs targeting **25M IOPS 512B random read**, first PCIe 6.0 SSDs by 2026 [secondary]. **Silicon Motion SM8466:** Gen6 controller, up to 28 GB/s, 512 TB capacities [secondary]. **FADU:** Gen6 controller at up to 28.5 GB/s under 9W [secondary].
- **Consumer Gen6:** not expected before **2028–2029**; 2026–2027 Gen6 is enterprise-only [secondary].
- **Strategic read:** the Gen6/512 TB push is explicitly **AI-driven** — checkpoint bandwidth (see E12: ~132 GB/s aggregate measured) and node-local "Tier 0" capacity scale with GPU counts; hyperscalers want fewer, denser, faster drives per node [secondary].
- **Pricing signal:** no 2026 street prices exist yet for Gen6 enterprise drives (sampling/qualification phase at cutoff); expect the new-gen premium pattern of prior transitions (Gen5 launch pricing ~2× Gen4 $/TB, decaying over 12–18 months) — pattern observation, not a quote [secondary].

## E17 — Conflicts, gaps, and unverified claims register

**Conflicts (kept unresolved, both sides recorded):**
1. **Consumer SSD $/TB (Sept 2026):** MemoryPriceChart basket median $159.94/TB vs homelab roundup examples $65–140/TB. Resolution: different baskets (mainstream PCIe 4.0 1–2 TB vs selected TLC models), different weeks, different regions — a range, not a contradiction [secondary].
2. **HBM market shares:** Q1 2026 "latest public" 58/21/21 (SK hynix/Samsung/Micron) vs SemiAnalysis Rubin R200-class 60/30/0. Resolution: different scopes (all-HBM vs Rubin-platform allocation) and different dates [secondary].
3. **DRAM shortage depth:** SemiAnalysis "7% below demand 2026–27" vs spot-price 4–5× contract levels. Resolution: both can be true — single-digit physical shortfall with extreme price elasticity at the margin [secondary].
4. **WDC–Kioxia merger:** reported as "restarted" (TradingKey, July 2026) vs assessed "probability remains low" (Mitrade, Aug 2026; SK Hynix veto unlifted). Kept as competing assessments [secondary].
5. **SanDisk stock:** "$1,791.82 on Sept 18" (Reuters via TheStreet) vs "~470% YoY surge" — both from the same press chain; the absolute figure is unusually high and single-sourced; treat the *direction* as corroborated, the *level* as [unverified].
6. **Dell soft-block currency:** documentation "predates 16G and 17G" — current-generation behavior unconfirmed [secondary].
7. **Samsung 30% HBM4 Rubin share vs "recovery mode; yield delays":** Samsung's HBM4 mass-production claims and its qualification struggles coexist in reporting; allocation share is an estimate, not a shipment fact [secondary].

**Gaps (not found by cutoff):**
- G1. Exact 2026 $/TB for new 30 TB QLC enterprise drives (only the VDURA ratio, 22.6× HDD, was found).
- G2. SK hynix Solidigm US NAND fab: decision status.
- G3. NGD Systems and other CSD vendors: 2026 status.
- G4. OpenSSD platform: 2026 revision/adoption evidence.
- G5. Gen6 enterprise SSD street pricing (pre-launch at cutoff).
- G6. HBM4 contract (LTA) price levels — only spot figures surfaced.
- G7. YMTC/CXMT NAND output volumes in wafers/month for 2026.
- G8. Dell 16G/17G third-party drive policy documentation.
- G9. Per-vendor enterprise SSD warranty terms comparison (2026).
- G10. Counterfeit-drive prevalence statistics for 2026 (only detection guidance found).

**Unverified claims carried with the tag (do not cite as fact):**
- Google CXL datacenter deployment; NVIDIA Vera CPU CXL 3.1 support.
- SanDisk $1,791.82 share price level.
- "512 TB Gen6 in 2027" — vendor roadmap statement, not a shipped product.

