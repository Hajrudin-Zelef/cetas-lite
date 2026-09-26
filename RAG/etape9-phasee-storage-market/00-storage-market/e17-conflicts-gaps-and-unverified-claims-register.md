---
id: etape9-phasee-storage-market/00-storage-market/e17-conflicts-gaps-and-unverified-claims-register
title: "E17 — Conflicts, gaps, and unverified claims register"
domain: step-9-phase-e-storage-memory-market-2026
role: deep-dive
task: reference
actors: ["Google", "Nvidia", "Samsung", "United States"]
dates: ["2026-07"]
keywords: ["consumer", "datacenter", "dram", "gpu", "hbm", "hbm4", "merger", "nand", "nvidia", "pricing", "rubin"]
source: docs/RAG/etape9_phaseE_storage_market.md
source_anchor: ""
source_lines: [204, 237]
section: "Step 9 Phase E — Storage & Memory Market 2026"
sha256: 3c7184b5eb53bd8ca44a161040ac36184cd59e04c2e86c640c0be6348e60d28b
---

# E17 — Conflicts, gaps, and unverified claims register

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

