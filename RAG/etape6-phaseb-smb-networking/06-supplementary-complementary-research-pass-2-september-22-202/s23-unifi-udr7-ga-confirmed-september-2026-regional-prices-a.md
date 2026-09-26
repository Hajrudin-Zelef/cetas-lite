---
id: etape6-phaseb-smb-networking/06-supplementary-complementary-research-pass-2-september-22-202/s23-unifi-udr7-ga-confirmed-september-2026-regional-prices-a
title: "S23. UniFi UDR7 — GA confirmed, September 2026 regional prices/availability (was: name only)"
domain: supplementary-complementary-research-pass-2-september-22-202
role: deep-dive
task: reference
actors: ["EU", "Qualcomm", "United States"]
dates: ["2026-05", "2026-09"]
keywords: ["dram", "memory", "pricing", "throughput"]
source: docs/RAG/etape6_phaseB_smb_networking.md
source_anchor: ""
source_lines: [1090, 1126]
section: "Supplementary / Complementary Research Pass #2 — September 22, 2026"
sha256: 02edadde7e95ced7665206387eabecbb21f6c388977475b663f2afdffdd55da0
---

# S23. UniFi UDR7 — GA confirmed, September 2026 regional prices/availability (was: name only)

- **ECW526** (Cloud7 2x2x2 tri-band Wi-Fi 7, 10G PoE+ uplink, 9.4 Gbps aggregate, Qualcomm platform, 320 MHz, MLO, 4096-QAM): official store **$369.00** (store.engeniustech.com, Sep 2026); street **under $300** per Dong Knows review (~May 2026, updated ~125 days before Sep 22, 2026) [official/independent]
- Dong Knows assessment [independent]: "virtually the Wi-Fi 7 version" of EnGenius's Wi-Fi 6E AP, cloud-managed only (no local operation), performance "subdued" vs the 10G uplink — but recommended as a quick Wi-Fi 7 upgrade path
- **ECW520** (Cloud7 2x3x3 Lite tri-band Wi-Fi 7, 2.5 GbE uplink, 802.3at): listed on the 2026 EnGenius switch/AP category staging page — budget Wi-Fi 7 tier below ECW526 [vendor-reported]
- **ECW516L** (2x2:2/3x3:3 asymmetric Wi-Fi 7): also on the 2026 category page [vendor-reported]
- **ECS5512F** (Cloud Managed 12-Port 10-Gigabit Half-Rack Aggregate Fiber Switch): 12× 10G, half-rack form factor — new SKU in the 2026 switch lineup page, complements ECS2530FP/ECS2552FP aggregation options [vendor-reported]
- Current 2026 EnGenius switch segmentation per official staging page [vendor-reported]: Layer 3 models / 48-port L2 / 24-port L2 / half-rack aggregate (ECS5512F)
- Sources: https://store.engeniustech.com/products/ecw526-cloud-managed-2x2x2-indoor-tri-band-wifi-7-access-point ; https://dongknows.com/engenius-ecw526-wi-fi-7-access-point-review/ ; https://www.engeniustech.com/cloud-managed-network-switch.html ; https://static.engeniuscdn.com/wp-content/uploads/2026/02/12114759/DS_EnGenius-Cloud-APs_v6.0.pdf

## S23. UniFi UDR7 — GA confirmed, September 2026 regional prices/availability (was: name only)

UDR7 is shipping broadly as of Sep 22, 2026 [independent]:
- BG (rlan.bg): **€320.23 incl VAT**, 90+ in stock, delivery ~Sep 30, 2026
- EU distributor (aerial.net): **€268.49 ex VAT**, 1 pc in stock + 110 pcs ETA Sep 30
- UK (senetic.co.uk): **£247.68 ex VAT** / £297.22 incl VAT, 8 pcs, delivery ~Oct 1, 2026
- ZA (scoop.co.za): **R6,775.00 incl VAT** (60 units across branches) / (geewiz.co.za): R6,242 incl VAT, in stock at external supplier
- IL (senetic.co.il): ₪1,172.66 incl VAT, 2 pcs
- ZA (miro.co.za): **sold out, new stock ETA Nov 5, 2026** — R6,941.40 incl VAT retail
- Specs confirmed across retailers [vendor-reported via retailers]: quad-core, 3 GB RAM, 1× 10G SFP+ WAN + 1× 2.5G RJ45 WAN + 3× 2.5G LAN (1× 802.3af PoE-out), tri-band Wi-Fi 7 (10.7 Gbps aggregate: 5.7 + 4.3 + 0.688 Gbps), 64 GB microSD included, 0.96" LCD, manages 30+ UniFi devices / 300+ clients, ~2.3 Gbps IDS/IPS throughput, Protect/Access/Talk pre-installed
- **Flag:** US store.ui.com list price not captured in this pass — do not assume
- Sources: https://rlan.bg/en/29104-ubiquiti-udr7-unifi-dream-router-7 ; https://www.senetic.co.uk/product/UDR7 ; https://miro.co.za/01-wi-fi-ubiquiti-unifi-cloud-gateways/8256-ubiquiti-unifi-cloud-gateway-dream-router-7-udr7-810084699980.html ; https://scoop.co.za/ubiquiti-unifi-dream-router-7-tri-band-cloud-gateway-udr7.html

## S24. Grandstream GWN7700M unmanaged + 2026 supply/availability notes (was: GWN7700MP only; no supply section)

- **GWN7700M** (non-PoE sibling of the GWN7700MP in Part 2 §5): unmanaged 5× 2.5G RJ45 + 1× 10G SFP+, fanless metal desktop/wall; **€62.92 ex VAT** (€76.21 incl VAT, CZ discomp, in stock, ships within 48h, Sep 2026) [independent]
- **No new Grandstream switch model announced in 2026 found** — GWN7700M/MP, GWN7800 Pro, GWN7810/7816, GWN7821P/22P, GWN7830/31/32 remain the current line [unverified — absence of evidence]; GWN78xx firmware train documented (e.g. 1.0.13.18, Apr 2025) with GDMS cloud-management compatibility fixes [secondary]
- **Supply/availability snapshots, Sep 22, 2026** [independent]:
  - US Section 301 tariff surcharge now itemized at checkout by US retailers (e.g. +$73.13 on MikroTik CRS504 at balticnetworks.com) — tariff pass-through is becoming explicit in networking hardware pricing
  - Zyxel XGS2220-30HP: very low stock at EU reseller (1 unit, Malta)
  - Ubiquiti UDR7: mixed — broad EU/UK/IL stock, sold out in ZA (restock Nov 5); UTR Long-Range (S18) still selling out within hours
  - QNAP QSW-M7308R-4X (4× 100G + 8× 25G, half-width, MSRP $1,927): **backordered** at directdial.com (CA, $1,898 CAD) and pc-canada.com (MSRP $1,985.32); but **>10 pcs in stock** at galaxus.de (€1,615.02 incl VAT, −10%); SGD $2,215.55 (aceperipherals.com); ₹214,500 incl GST (tanotis.com, India) — availability is channel-dependent
  - Aruba Instant On 1960 S0F35A: NZD $2,876.66 (PB Tech, listed Sep 22, 2026); UK £604.66 ex VAT, awaiting stock
- **Flag:** no 2026 DRAM/memory-price-driven switch price hikes beyond Ubiquiti's surcharge (S9) were independently verified in this pass
- Sources: https://www.discomp.cz/grandstream-gwn7700m-unmanaged-network-switch-5x-2-5gb-portu-1-sfp-_d127130.html ; https://www.balticnetworks.com/products/mikrotik-crs504-4x-qsfp28-100gbps-cloud-router-switch ; https://www.directdial.com/ca/item/qnap-half-rackmount-switch-qsw-m7308r-4x-us-layer-2-12-port-managed-switch-4-p/qsw-m7308r-4x-us ; https://www.galaxus.de/en/s1/product/qnap-qsw-m7308r-4x-managed-switch-4-port-100gbe-8-port-25gbe-half-rackmount-design-12-ports-network--39444919


---

