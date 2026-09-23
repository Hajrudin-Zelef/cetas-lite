---
id: etape6-phaseb-smb-networking/06-supplementary-complementary-research-pass-2-september-22-202/overview
title: "Supplementary / Complementary Research Pass #2 — September 22, 2026"
domain: supplementary-complementary-research-pass-2-september-22-202
role: deep-dive
task: reference
actors: ["China", "United States"]
dates: ["2026-09-22"]
keywords: ["research", "cost", "license", "memory", "nand", "throughput"]
source: docs/RAG/etape6_phaseB_smb_networking.md
source_anchor: ""
source_lines: [839, 888]
section: "Supplementary / Complementary Research Pass #2 — September 22, 2026"
sha256: 2f56d0c8f4c342b1635e37eec2db15f5529d24cefc3a730a7d5608376b182008
---

# Supplementary / Complementary Research Pass #2 — September 22, 2026

**Scope:** second complementary pass. New material only — nothing in the base report or Pass #1 was altered. New topics: **MikroTik** (homelab/SMB angle, new vendor to this file), **Grandstream Wi-Fi 7 launch evidence** (corrects Pass #1/S13 gap claim that "no confirmed Grandstream Wi-Fi 7 AP launch found"), cross-vendor **Wi-Fi 7 AP comparison table**, Zyxel XGS/XMG street prices, Aruba Instant On 2026 status, and the **2026 cost shock** (memory tariffs/surcharges, Dell'Oro ASP +19%, Cisco price hikes, FCC China-router ban effects). Plus remaining open items.

Provenance tags: [official] = vendor price list/store/datasheet; [vendor-reported] = vendor claim via trade press; [independent] = third-party test/retailer listing; [secondary] = press/blog/forum; [unverified] = single weak source or inferred.

---

## T1. MikroTik — homelab/SMB angle (new vendor to this file)

MikroTik was covered in the enterprise track but is absent from this SMB file; it is the price leader in the homelab segment and materially changes the SMB price floor [independent assessment]. Current 2026 models/prices:

- **CRS310-8G+2S+IN** (Cloud Router Switch): 8× 2.5GbE + 2× 10G SFP+ cages (SFP+ cages also support 1G/2.5G), Marvell 98DX226S switch chip, dual-core ARM, 256 MB RAM / 32 MB storage, RouterOS v7 license 5, VLANs/Jumbo/LAG/ACL, some hardware-offloaded L3, rackmount ears included [vendor-reported via retailers].
  Source: https://www.comx-computers.co.za/MT-RBCRS310-8G-2S-IN-MikroTik-Cloud-Router-Switch-8x-Buy-p-290412.php
- **CRS310-1G-5S-4S+IN**: compact fiber switch — 5× 1G SFP + 4× 10G SFP+, ARM v7, 256 MB RAM, RouterOS L5, 1U rackmount; street **£161.99 inc VAT** (linitx, Sep 2026) [independent].
  Source: https://linitx.com/category/mikrotik-home-office-switches/1156
- **CRS312-4C+8XG-RM**: 8× 10G RJ45 + 4× 10G combo (RJ45/SFP+), 240 Gbps switching, 178 Mpps, dual PSU, RouterOS or SwOS boot [secondary/YouTube].
  Source: https://www.youtube.com/watch?v=dyrzxPnrG7I
- **CRS309-1G-8S+IN**: 8× 10G SFP+ + 1× GbE, dual-core 800 MHz, 512 MB RAM, dual-boot RouterOS/SwOS, 162 Gbps, 119 Mpps, **fanless** [independent/retailer specs].
  Source: https://www.otech.com.bd/mikrotik-crs309-1g-8sin-–-high-performance-10g-fiber-switch
- **CRS326-24G-2S+RM**: 24× GbE + 2× SFP+, street **£171.59 inc VAT** (Sep 2026) [independent]; desktop variant CRS326-24G-2S+IN £173.99 [independent].
  Source: https://linitx.com/category/mikrotik-home-office-switches/1156
- **CRS418-8P-8G-2S+RM**: 8× 1G PoE+ (802.3af/at + 24V passive) + 8× 1G non-PoE + 2× 10G SFP+, 1 GB RAM, L3 hardware offloading, 150W PoE budget, dual-redundant PSU, mgmt/console ports; street **£185.99 inc VAT** (Sep 2026) [independent].
  Source: https://linitx.com/category/mikrotik-home-office-switches/1156
- **Wi-Fi 6 variant:** CRS418-8P-8G-2S+5axQ2axQ-RM adds Wi-Fi 6 to the CRS418 PoE switch — **£383.99 inc VAT** [independent, linitx].
- **CSS610-8P-2S+IN**: 8× GbE PoE-out + 2× SFP+, 140W PoE budget; street **£179.99 inc VAT** [independent, linitx].
- **RB5009UG+S+IN** (router): 7× GbE + 1× 2.5G + 1× 10G SFP+, Marvell Amethyst switch chip, Armada quad-core ARMv8 1.4 GHz, 1 GB DDR4 / 1 GB NAND, USB 3.0; street **€199.90** (Galaxus, Sep 2026) [independent] — consistently cited as the "homelab router" [independent assessment].
  Source: https://www.galaxus.nl/en/s1/producttype/toplist/brand/routers-64/mikrotik-11434
- **hEX S (2025)** refresh: 5× GbE + 2.5G SFP, PoE out, USB, dual-core; street **€75.06** (Galaxus) [independent].
  Source: https://www.galaxus.nl/en/s1/producttype/toplist/brand/routers-64/mikrotik-11434
- Positioning notes [independent assessment]: RouterOS v7 depth (VLANs, LACP, L3 HW offload) plus fanless/sub-€200 price points make MikroTik the budget default in homelab/server-room cores; trade-offs consistently cited are fan noise on larger units (CRS317) and RouterOS learning curve [secondary/YouTube].
- **Flag:** MikroTik is Latvia-based but its China supply exposure to the 2026 FCC China-router ban discussion (see T6) is unclear — not found, treat as unknown [unverified].

## T2. Grandstream Wi-Fi 7 — NEW launch evidence (corrects S13 gap #7)

Pass #1 concluded "no confirmed Grandstream Wi-Fi 7 AP launch found in 2026." Retailer evidence now shows a **GWN767x Wi-Fi 7 family** shipping — correction recorded here [independent]:

- **GWN7672**: enterprise tri-band Wi-Fi 7 indoor AP **with 10G connectivity**, BLE; listed as "New — in stock" at Australian distributor InfrontTech (Sep 2026) [independent].
- **GWN7672L**: Wi-Fi 7 AP (long-range/ceiling variant); datasheet dated **Aug 6, 2026**, price **£125.10** (VoIPon, UK) [independent].
  Source: https://www.voipon.co.uk/pdf_datasheet.php?products_id=10774
- **GWN7670**: enterprise dual-band Wi-Fi 7 (802.11be) indoor AP — 3.6 Gbps aggregate wireless, 5 Gbps aggregate wired throughput, 2×2:2 MIMO, MLO/4K-QAM/MRU/preamble puncturing; listed at **R2,685** (bobshop.co.za, South Africa; listing dated ~Oct 2025 per crawl) [independent].
  Source: https://www.bobshop.co.za/grandstream-enterprise-wifi-7-indoor-access-point-gwn7670/p/657743542
- **GWN7670WM**: Wi-Fi 7 wall-mount AP (BE3600, 2×2, 2.5GbE, BLE) — listed "New — in stock" (InfrontTech AU) [independent].
- **GWN7670LR**: Wi-Fi 7 outdoor AP, long range (350 m); on sale at Wavetech (Kenya) [independent].
  Source: https://wavetech.co.ke/product-category/grandstream/
- **GWN7710R**: 6-port (5× GbE + 1× SFP) outdoor L2-lite managed PoE switch, IP66 — listed "New — in stock" (InfrontTech AU) [independent]. Shows the GWN switch line expanding into outdoor/SMB edge.
- Current GWN766x Wi-Fi 6 street context (NTS Direct, Sep 2026) [independent]: GWN7660 $105, GWN7660E $79, GWN7660EM $65, GWN7662 $129, GWN7664 $199, GWN7664E $165, GWN7664ELR $199, GWN7664LR $229, GWN7602 $65, GWN7604 $69, GWN7605 $69, GWN7605LR $99, GWN7615 $89, GWN7624 $79, GWN7625 $85, GWN7630 $109, GWN7630LR $125; GWN7661 marked EOL ($95); **GWN7303 listed "coming soon"** ($1,000 — likely placeholder price, treat as unverified).
  Source: https://shop.ntsdirect.com/category/325/Grandstream-Access-Points--and--Routers.html
- **Flag:** no vendor press release or MSRP found for the GWN767x Wi-Fi 7 line in sources searched — retailer evidence only; US list prices unconfirmed [unverified].

