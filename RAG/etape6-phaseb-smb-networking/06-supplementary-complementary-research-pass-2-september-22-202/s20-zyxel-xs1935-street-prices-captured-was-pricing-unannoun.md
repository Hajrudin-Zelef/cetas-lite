---
id: etape6-phaseb-smb-networking/06-supplementary-complementary-research-pass-2-september-22-202/s20-zyxel-xs1935-street-prices-captured-was-pricing-unannoun
title: "S20. Zyxel XS1935 street prices captured (was: pricing unannounced)"
domain: supplementary-complementary-research-pass-2-september-22-202
role: deep-dive
task: pricing
actors: ["Falcon", "Intel", "Malaysia", "United States"]
dates: ["2026-03", "2026-04-10", "2026-08-07", "2026-08-26"]
keywords: ["pricing", "intel", "memory"]
source: docs/RAG/etape6_phaseB_smb_networking.md
source_anchor: ""
source_lines: [987, 1033]
section: "Supplementary / Complementary Research Pass #2 — September 22, 2026"
sha256: 40d7f42aa013f52b66c16f4638a500a44cd3d9a98942345aeab159a721d0e3bf
---

# S20. Zyxel XS1935 street prices captured (was: pricing unannounced)

## S20. Zyxel XS1935 street prices captured (was: pricing unannounced)

- **XS1935-12F** (12-port managed multi-gigabit/10G, EAN 4718937637805, L2, 240 Gbps fabric, 178 Mpps, 32K MAC table, no PoE) [independent retailer specs]:
  - Lithuania (shopro.lt, listing generated **2026-08-07**): **€448.99** ; earlier crawl 2026-04-10 showed €432.99 — mild upward drift across 2026
  - South Africa (itworkup.co.za, 2026-08-26): **R9,714**
- This resolves the S3 gap for the -12F variant; **XS1935-12HP (PoE++) and XS1935-10 prices still not located** [unverified].
- Adjacent datapoint: **XGS1935-52HP** (48× GbE PoE+ + 4× 10G SFP+) reviewed by IT Pro at **£528 ex VAT** via Broadbandbuyer (~£100 less than its predecessor); Zyxel UK comms-express shows in-stock next-day across XGS1935 SKUs (£161.91–£508.79 ex VAT range) [secondary].
- Sources: https://www.shopro.lt/index.php?route=product/pdf&product_id=540230 ; https://itworkup.co.za/index.php?route=product/pdf&product_id=610220 ; https://www.itpro.com/hardware/routers/zyxel-xgs1935-52hp-review-a-port-dense-gigabit-poe-switch-thats-priced-right-for-smbs ; https://www.comms-express.com/categories/zyxel-xgs1935-series-smart-managed-switches/

## S21. NETGEAR — two NEW M4350 models launched at ISE 2026 (was: no confirmed 2026 launches)

Per VARINDIA and Mix coverage of **ISE 2026** [secondary/vendor-reported]:
- **M4350-16M4V** — 16× 2.5G PoE++ (up to **1,130W** total PoE) with Neutrik locking connectors throughout (etherCON for network, opticalCON for fiber, powerCON for power), 4× 25G SFP28 uplinks, **modular uplink interface card slot** (cards with 4× 10G RJ45, or 2× Neutrik opticalCON Quad for 4× MMF or 4× SMF transceivers); dual powerCON on back panel; live-event hardened
- **M4350-16C** — 16× **100G** ports, designed for aggregation/core layers of large AV-over-IP deployments
- Both bring the M4350 portfolio to **18 models** total (1G→100G); TAA-compliant versions available; general availability from **March 2026** [vendor-reported]
- Also at ISE 2026: new **NETGEAR Engage Controller** with offline provisioning (design/configure networks before arriving on-site) [vendor-reported]
- New detail on existing SKU from TAA listing [secondary, creationnetworks.net]: **XSM4328CV-TAANES (M4350-24X4V)** — 24× 10G/Multi-Gig PoE+ + 4× 25G SFP28, 880W internal PSU giving 576W PoE (up to 720W with redundant budget remaining 576W), Virtual Chassis stacking with NSF hitless failover, SMPTE ST 2110 Grand Master/Boundary Clock, L3 feature set (static, policy-based, dynamic routing)
- Sources: https://www.varindia.com/news/netgear-introduces-next-gen-m4350-switches-with-rugged-design-and-offline-av-ready-network-configuration-at-ise-2026 ; https://www.mixonline.com/technology/news-products/neutrik-netgear-team-for-m4350-switch ; https://creationnetworks.net/products/netgear-xsm4328cv-taanes-28pt-m4350-24x4v-managed-switch

## S22. Grandstream Wi-Fi 7 — 2026 GA status resolved (was: GA unconfirmed)

- **GWN7670LR** is now confirmed shipping/retail [secondary/independent]:
  - B&H Photo US listing (GWN7670LR, EAN 6947273705598) — specs: Wi-Fi 7 (802.11be) dual-band 2×2:2, 3.6 Gbps aggregate (2,882 Mbps 5 GHz + 688 Mbps 2.4 GHz), 1× 2.5G RJ45 PoE + 1× 2.5G SFP, 256 clients, IP66, Bluetooth 5.3 LE
  - Nigeria (jiji.ng, Sep 2026): **₦248,500** ; microless RU: **₽12,500** incl. VAT, in stock
  - Voipon UK listing confirms MLO, 4K-QAM, MRU, preamble puncturing; "first Grandstream AP to include a PtMP bridge"
- **Full GWN7670 Wi-Fi 7 family** (from Grandstream's GWN series material) [secondary via info.teledynamics.com]:
  - **GWN7670** — 3.6 Gbps, 2× 5G RJ45, ceiling/wall, up to 175 m
  - **GWN7670WM** — in-wall variant
  - **GWN7672** — **tri-band** (2.4/5/6 GHz), 2× 5G RJ45, ceiling/wall — high-density flagship
  - **GWN7670LR** — long-range outdoor, up to 350 m, wall/pole IP66, 15.5W
  - Management: embedded controller (50 APs), GDMS Cloud, GWN Manager; 32 SSIDs (16/radio); self-power adaptation on PoE+
- **Flag:** formal 2026 launch announcement date not pinned down — retail presence confirms availability; treat as [secondary] not [official].
- Sources: https://www.bhphotovideo.com/c/product/1932825-REG/grandstream_gwn7670lr_long_range_dual_band_wi_fi.html ; https://www.voipon.co.uk/grandstream-gwn7670lr-longrange-dualband-wifi-7-access-pointbridge-gwn7670lr-p-10450.html ; https://info.teledynamics.com/blog/grandstream-expands-its-wi-fi-7-ap-lineup-with-the-gwn7670wm-and-7670lr

## S23. D-Link DQS-5000 deep-dive + Instant On availability snapshot

- **DQS-5000-56ZS — official D-Link product page detail** [official, dlink.com/usweb datasheets]:
  - Switch silicon: **Marvell Falcon** ; CPU: Intel x86 4-core ; Memory: 8 GB DDR4 SO-DIMM w/ECC ; Storage: 16 GB eMMC ; Packet memory: 24 MB
  - Forwarding rate: **2,380 Mpps** ; FDB: 128K ; router host routes: 288K IPv4 / 144K IPv6 ; unicast prefixes: 512K IPv4 / 256K IPv6 ; ARP: 192K
  - Variant **DQS-F5000-56ZS** listed as "25G Data Center Switch" (dlink.com India) with same 48× 25G SFP28 + 8× 100G QSFP28 port count — likely fixed-OS sibling of the ONIE-based 56ZS [official — SKU distinction not explained by D-Link, flag]
  - Sibling **DQS-5000-32Q28** (32× 100G QSFP28, front-to-back airflow) listed on Zoro.com (G9650079, UPC 790069435706) [secondary — price not captured]
- **Pricing:** no public street price found for DQS-5000-56ZS — quote/channel model [unverified gap retained]
- **Aruba Instant On availability snapshot, Sep 2026** [secondary, retailer listings]: Miro South Africa showed **Instant On 1930 JL686B (48G 4SFP+ 370W) sold out** with "Estimate arrival of new stock: 27-08-2026" at **R15,176.55** incl. VAT — mid-range Instant On supply is tight; lelong.com.my (Malaysia) listings for 1830/1930/1960 JL-SKUs remain active with Sep-2026 auction windows
- Sources: https://usweb.dlink.com/en/-/media/product-pages/dqs/dqs-5000-56zs/dqs500056zsa1datasheetv102ww.pdf ; https://www.dlink.com/in/en/-/media/product-pages/dqs/f5000-56zs/datasheet/dqsf500056zs_datasheet.pdf ; https://miro.co.za/07-networking-switches---poe-layer-2/6432-aruba-instant-on-1930-48-port-gb-4sfp-370w-switch-190017593753.html

---

