---
id: etape9-phasee-storage-market/00-storage-market/part-13
title: "Step 9 Phase E — Storage & Memory Market 2026 (part 13)"
domain: step-9-phase-e-storage-memory-market-2026
role: deep-dive
task: reference
actors: []
dates: []
keywords: ["memory", "consumer", "datacenter", "dram", "gpu", "inference", "nand"]
source: docs/RAG/etape9_phaseE_storage_market.md
source_anchor: ""
source_lines: [402, 419]
section: "Step 9 Phase E — Storage & Memory Market 2026"
sha256: 1d001de543932119f3f9872ec36dab8e5476d6397de52daa1161160ed8e10cf1
---

# Step 9 Phase E — Storage & Memory Market 2026 (part 13)

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

