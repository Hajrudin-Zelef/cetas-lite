---
id: etape9-phasee-storage-market/00-storage-market/part-11
title: "Step 9 Phase E — Storage & Memory Market 2026 (part 11)"
domain: step-9-phase-e-storage-memory-market-2026
role: deep-dive
task: reference
actors: ["Intel", "Samsung", "United States"]
dates: []
keywords: ["cost", "datacenter", "gpu", "intel", "latency", "pricing"]
source: docs/RAG/etape9_phaseE_storage_market.md
source_anchor: ""
source_lines: [342, 353]
section: "Step 9 Phase E — Storage & Memory Market 2026"
sha256: 765a217ab758221e21315f17b95e475b752ca5e534cf304df69af2e6561b8ebf
---

# Step 9 Phase E — Storage & Memory Market 2026 (part 11)

- **The classic homelab enterprise SSDs (recurring eBay/ServerPartDeals listings, 2025–2026):** Samsung PM883/PM893 (SATA), PM9A3/PM963 (NVMe U.2), Intel/Micron 5300/5400 Pro/Max (SATA), Intel DC P4510/P4610 (NVMe U.2), Kioxia/HGST SAS lines, Micron 9300/9400 Pro/Max (NVMe U.2). These are the parts the $78–127/TB band refers to [secondary].
- **Why read-intensive ("1 DWPD / RI") pulls dominate listings:** most datacenter flash does read-heavy work, so retired RI drives typically show single-digit percentage-used — the exact drives a homelab wants for VM datastores and media [secondary].
- **Write-intensive ("3–10 DWPD / WI") pulls:** rarer, pricier per TB, and the right buy only for SLOG/ZIL, database redo, or heavy ingest — and even then, new is the 2026 guidance for write-heavy tiers (see E9) [secondary].
- **SATA vs SAS vs NVMe on the used market:** SATA enterprise SSDs are the cheapest per TB and slot into any backplane; SAS needs a SAS HBA/controller; **NVMe U.2 needs U.2 bays or M.2/U.2 adapters plus PCIe bifurcation or tri-mode controllers** — factor the adapter/cable cost (~$20–60) into the $/TB math [secondary].
- **DDR4 ECC RDIMM (used, 2026):** the homelab RAM staple — 16 GB sticks $25–40, 32 GB sticks commonly $50–90 on eBay-class listings; 64 GB LRDIMMs higher but still far below new DDR5 RDIMM pricing [secondary]. Compatibility rule: match the server generation's rank/voltage support (e.g., Dell 13G/14G DDR4 2400/2666) — check the vendor QVL or community reports before bulk-buying [secondary].
- **The mini-PC vs used-server tradeoff (2026):** an MS-01-class mini PC idles at 13–35 W in silence, saving $120–200/year vs a rack server at average US power rates; the used server wins when the job needs hundreds of GB of cheap ECC RAM, 8–24 hot-swap bays, full-height GPU slots, redundant PSUs, and a real BMC — purchase price favors the server, running cost favors the mini PC [secondary]. Source: https://pcserverandparts.com/news/best-used-server-for-home-lab-2026/
- **Where the deals cluster:** corporate refresh cycles (3–5 years) dump matched lots — same SKU, same firmware, similar wear — which is ideal for RAID/ZFS vdevs; single-drive odd lots are cheaper per unit but complicate array building [secondary].
- **ServerPartDeals pattern:** enterprise-focused retailer with rotating stock of pulled drives, publishes per-drive SMART on request; typical premium over eBay auctions but with testing and return policy — the "verified middle" between eBay luck and new retail [secondary].
- **r/homelabsales pattern:** peer-to-peer, best for full servers and RAM lots; escrow via PayPal G&S is the community norm; prices typically 10–20% under eBay after fees [secondary].
- **What "recertified" means at different sellers:** manufacturer recertified (rare for enterprise SSDs) > refurbisher-tested with SMART report and 3–5-year warranty > "pulled working" (as-is, test on arrival) > "for parts." Price should descend in that order; if it doesn't, walk away [secondary].
- **Power-cost reality check for 24/7 spinning rust:** a 4-bay NAS with 8 TB HDDs at ~$80–120 used per drive plus ~25–40 W idle is still the cheapest $/TB/year for media in 2026 — the SSD premium only pays off where IOPS or latency matter [secondary].

