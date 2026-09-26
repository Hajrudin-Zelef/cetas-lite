---
id: etape6-phaseb-smb-networking/09-supplementary-research-pass-addendum-september-22-2026/part-6
title: "Supplementary Research Pass — Addendum (September 22, 2026) (part 6)"
domain: supplementary-research-pass-addendum-september-22-2026
role: deep-dive
task: reference
actors: ["United States"]
dates: ["2026-05", "2026-09-22"]
keywords: ["research", "ethernet", "pricing"]
source: docs/RAG/etape6_phaseB_smb_networking.md
source_anchor: ""
source_lines: [1784, 1817]
section: "Supplementary Research Pass — Addendum (September 22, 2026)"
sha256: 0b08247db1e3b99ad56cf51d7ea57fc9a4355ee3bff19b80e9676d9ea3314675
---

# Supplementary Research Pass — Addendum (September 22, 2026) (part 6)

| Product | Ports | Observed price | Retailer / date evidence |
|---|---|---|---|
| MikroTik CRS326-24S+2Q+RM | 24×10G SFP+ + 2×40G QSFP+ | NZD 1,210.00 incl GST (1,052.17 ex), 0 stock [secondary] | PB Tech NZ, product created 18-09-2026 (https://www.pbtech.co.nz/product/NETMKT1431/MikroTik-CRS326-24S2QRM-Cloud-Router-26-Port-10Gbp) |
| MikroTik CRS326-24S+2Q+RM | same | R11,495.00 incl VAT, in stock [secondary] | Scoop SA, page crawled ~Sep 20, 2026 (https://scoop.co.za/mikrotik-cloud-router-switch-24-port-sfp-2qsfp-crs326-24s-2q-rm.html) |
| MikroTik CRS326-24S+2Q+RM | same | €488.90 incl VAT, in stock >5 pcs [secondary] | Alza SK, page crawled ~Sep 13, 2026 (https://www.alza.sk/EN/mikrotik-crs326-24s-2q-rm-d7774774.htm) |
| MikroTik CRS326-24S+2Q+RM | same | from £488.89 [secondary] | idealo UK, page crawled ~Sep 4, 2026 (https://www.idealo.co.uk/compare/200374594/mikrotik-24-port-sfp-switch-crs326-24s-2q-rm.html) |

- Practical note: most affordable "40G" homelab paths in 2026 are QSFP28 100G ports run at 40G or broken out to 4×25G (e.g., CRS504/CRS518), because new native-40G QSFP+ SMB switches are scarce; the used market (below) is where native 40G lives [secondary/unverified synthesis of the listings above].

**C.5.3. Notable used-market references (40G and adjacent).**

| Item | Observed price | Marketplace / date evidence |
|---|---|---|
| Dell S4048-ON (48×10G SFP+ + 6×40G QSFP+), refurbished | US $352.80, in stock [secondary] | directmacro.com, page crawled ~Sep 19, 2026 (https://directmacro.com/dell-s4048-on.html) |
| Dell S4048-ON, used | US $349.99 [secondary] | Alta Technologies, page crawled ~Sep 15, 2026 (https://altatechnologies.com/collections/dell-networking-hardware) |
| Dell S6010-ON (32×40G QSFP+), used | US $157.99 [secondary] | Alta Technologies, page crawled ~Sep 15, 2026 (https://altatechnologies.com/collections/dell-networking-hardware) |
| Dell S6010-ON (32×40G QSFP+), open box, manufacturer refurbished, no PSUs/fans | US $350.00 [secondary] | eBay (ebay.ca listing), page crawled ~Jun 2026 |
| Mellanox MSX1024B-2BFS (48×10G SFP+ + 12×40G QSFP), pre-owned | C$1,818.31 [secondary] | eBay category page, crawled ~Sep 20, 2026 (https://www.ebay.ca/b/Mellanox-Technologies-Ethernet-Switch-Enterprise-Network-Switches/51268/bn_72746428) |
| Mellanox MSN2700 (32×100G), pre-owned | C$2,028.12 [secondary] | eBay, crawled ~Sep 20, 2026 (same category URL) |

- Excluded from the table: sub-€100 InfiniBand switches (e.g., Mellanox IS5030 at €99, SB7790 36×100G IB at US $690 on eBay) because they are InfiniBand, not Ethernet — flagged here so nobody mistakes them for cheap 40/100GbE [secondary].

#### C.6. Open questions (Pass C)

1. Whether the Zyxel "four new models" announcement of Sep 20, 2026 reflects genuinely new hardware or a regional launch of SKUs already listed in May 2026 accessory documentation [unverified].
2. The EAP783 launch-date discrepancy (Jan 2024 per Dong Knows vs broad channel availability mid-2026) — worth one targeted check if a clean timeline is needed [unverified].
3. The GWN7672L distributor-listing date (Feb 2026) vs official initial-firmware date (Jul 14, 2026) [unverified].
4. Exact launch dates for Grandstream GWN7670E, GWN7674, and Netgear WBE700 were not verified; none is claimed as a 2026 launch [unverified].
5. "WAX655" appears to be a non-existent SKU — recommend dropping it from the brief [unverified].

---

**Collection metadata (supplementary pass):** research conducted 2026-09-22 via three independent passes (Ubiquiti deep gaps; TP-Link/D-Link/Zyxel gaps; Netgear/HPE Instant On/Grandstream/Wi-Fi 7 table/25G-40G pricing); read-only web research, no files written by researchers, nothing sent externally. All facts dated on/before 2026-09-22. Provenance tags applied per fact; identifiers not guessed; uncertainties listed per pass (§A.8, §B.6, §C.6).

