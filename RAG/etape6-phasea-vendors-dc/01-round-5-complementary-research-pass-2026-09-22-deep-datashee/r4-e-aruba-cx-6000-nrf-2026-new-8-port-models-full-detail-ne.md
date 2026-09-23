---
id: etape6-phasea-vendors-dc/01-round-5-complementary-research-pass-2026-09-22-deep-datashee/r4-e-aruba-cx-6000-nrf-2026-new-8-port-models-full-detail-ne
title: "R4-E. Aruba CX 6000 — NRF 2026 new 8-port models, full detail (new)"
domain: round-5-complementary-research-pass-2026-09-22-deep-datashee
role: deep-dive
task: reference
actors: ["Huawei", "Intel"]
dates: ["2026-01-12", "2026-05-20", "2026-05-28", "2026-09-22"]
keywords: ["asic", "compute", "ethernet", "intel", "latency", "nvidia"]
source: docs/RAG/etape6_phaseA_vendors_dc.md
source_anchor: ""
source_lines: [1993, 2070]
section: "Round-5 complementary research pass (2026-09-22) — deep datasheet and radix detail"
sha256: 5b069a483aa025ce20d928bdc86ad37ac27722ead7c13387e7c4448ceddc1aaf
---

# R4-E. Aruba CX 6000 — NRF 2026 new 8-port models, full detail (new)

### R4-E. Aruba CX 6000 — NRF 2026 new 8-port models, full detail (new)

Base timeline mentioned the NRF 2026 addition without detail. Expanded from the January 12, 2026 HPE announcement and Network World coverage [official / secondary]:

- HPE announced **new 8-port models in the HPE Aruba Networking CX 6000 Switch Series** at NRF 2026 (announcement dated **January 12, 2026**): compact, **completely silent (fanless) design** for checkout lanes and overhead crawl spaces [official].
  Source: https://www.networkworld.com/article/4115610/nrf-2026-hpe-expands-network-server-products-for-retailers.html
- New 8-port model: Layer 2, up to **104 Gbps non-blocking bandwidth**, **77.3 Mpps forwarding**, HPE Aruba Networking ASIC architecture, programmable Aruba CX OS [secondary].
- CX 6000 family composition: **five fixed 1U models** with 24 and 48 access ports of 1GbE + four built-in 1G SFP uplinks; 24-port PoE models up to **370W**, 48-port PoE models up to **740W** (802.3at Class 4, 30W/port) — for PoS terminals, IoT, cameras, staff systems [secondary].
- Street anchor: **R8N85A** (CX 6000 48G Class4 PoE 4SFP 370W — 48× 10/100/1000BASE-T + 4× 1G SFP, 104 Gbps, 77.3 Mpps, 8,192 MAC, 1.9 µs latency) — **$1,469.99** new (NetworkTigers) [independent].
  Source: https://www.networktigers.com/products/r8n85a-hpe-switch-new
- Announcement context: integration of AI-based **Marvis with HPE Juniper Networking analytics**; a new network sensor supporting Wi-Fi 7 with end-user activity view; HPE Nonstop Compute NS9 X5 and NS5 X5 upgrades [official].

### R4-F. Gartner Magic Quadrant 2026 — additional detail (complements base §U)

Base §U named Arista as a 2026 Gartner MQ Leader for enterprise wired/wireless LAN. Additional 2026 MQ developments [vendor-reported / secondary]:

- **HPE claims** "furthest in Completeness of Vision and highest for Ability to Execute" in the **2026 Gartner Magic Quadrant for Enterprise Wired and Wireless LAN Infrastructure — fifth consecutive time** [vendor-reported].
  Source: https://www.hpe.com/emea_europe/en/networking/magic-quadrant-wired-wireless.html
- **Arista** named a Leader (announced **May 20, 2026**): alongside the MQ news, Arista announced Cognitive Campus innovations — **ruggedized switching platforms 710HXP-28TXH and 710HXP-20TNH** (industrial/outdoor, EOS-based) and a new family of internal directional-antenna wireless APs [vendor-reported].
  Source: https://www.businesswire.com/news/home/20260520437739/en/Arista-Networks-Positioned-as-a-Leader-in-the-2026-Gartner-Magic-Quadrant-for-Enterprise-Wired-and-Wireless-LAN
- **Huawei** named a Leader for the **fourth year in a row** (May 28, 2026); **the only non-North American vendor in the Leaders quadrant**; cites Xinghe AI Campus (AI-powered O&M, full-scope security, network autonomy) and an AirEngine Wi-Fi 7 Advanced AP incorporating anticipated Wi-Fi 8 innovations (iCSSR multi-AP coordination, SmartBF beamforming, ASFN zero-roaming) [vendor-reported].
  Source: https://www.webdisclosure.com/press-release/huawei-etr-huawei-named-a-leader-in-the-2026-gartner-magic-quadrant-for-enterprise-wired-and-wireless-lan-infrastructure-for-the-fourth-year-in-a-row-p7AObeR7kp0
- Note: this is the LAN-infrastructure MQ; no dedicated data-center-networking MQ entry for 2026 was located in this pass — **flag**.

### R4-G. Dell Z9864F-ON — spec depth (street price still a gap)

- Verified hardware detail from reseller datasheets (Grabnpay): **Intel Xeon D-1714** (4 cores @ 2.3 GHz, up to 3.4 GHz Turbo), **32 GB DDR4 ECC**, 2U fixed chassis, dual AC or DC hot-swappable PSUs, four dual-rotor hot-swap fans, RJ45 serial console + USB Type-C console, 10/100/1000Base-T Ethernet management, USB 3.0 Type-A; airflow options I/O-to-PSU or PSU-to-I/O; ONIE; **TAA-compliant**, RoHS, IPv6 Ready [secondary].
  Source: https://www.grabnpay.in/products/dell-powerswitch-z9864f-on-64-800gbe-osfp112-high-density-open-networking-switch-with-dual-ac-dc-power-supplies-and-hot-swappable-fans
- Breakout: **128× 400GbE / 256× 200GbE / 320× 100GbE** via breakout (Dell spec sheet) [official].
  Source: https://delltechnologies.com/asset/ko-kr/products/networking/technical-support/dell-powerswitch-z9864f-on-spec-sheet.pdf
- AI-fabric features per spec sheet: PFC, DCBX, RoCEv2, ETS for DCB environments; multi-path hashing, adaptive routing, enhanced AI observability, Dell SmartFabric Manager support [official].
- Street price: reseller BigFastServers lists the Z9864F-ON ("ships within 1-2 business days") but **shows no price** — dollar street price for Z9864F-ON remains unconfirmed; this file's §1.7 gap stands [gap].
  Source: https://bigfastservers.com/products/dell-powerswitch-z9864f-on-with-102-4-tbps-switch-capacity-64x-800-gbe-osfp112-2x-sfp-ports
- **Reseller data-quality flag**: juaraitsolutions.com describes the Z9864F-ON as "64 ports of 400GbE" — incorrect (it is 64× 800GbE OSFP112); do not reuse that claim [unverified/incorrect].
  Source: https://juaraitsolutions.com/dell-powerswitch-z-series-switches/

### R4 verification log (new open items)

1. **Cumulus 5.19.0**: confirmed only via fix-table reference; no dedicated What's-New page located — GA scope/contents unconfirmed.
2. **Dell Enterprise SONiC 4.7.x**: no evidence of a 4.7 train through 2026-09-22; confirm whether Dell is skipping to a 2027 train.
3. **MS210 MSRP variance**: SHI vs publicsector SHI differ ~$1.2K on MS210-48FP-HW MSRP — one storefront may be stale; treat either MSRP with caution.
4. **No dedicated 2026 Gartner Magic Quadrant for data-center networking located** in this pass.
5. **Z9864F-ON street price** still not found from any retailer.
6. **Z9664F-ON / Z9432F-ON street prices**: not searched in this pass — still a gap.
7. **Accio/naddod prices** are gray-market/AI-generated content risk (accio.com explicitly flags AI-generated FAQ content); treat as street anchors, not list prices.

### R4 sources (verbatim URLs)

- https://www.dell.com/support/kbdoc/en-ph/000228560/minimum-recommended-and-latest-code-versions-for-networking-products
- http://infohub.delltechnologies.com/static/media/client/7phukh/DAM_60dba377-fc13-4251-9dd6-c81409b4095d.pdf
- http://infohub.delltechnologies.com/static/media/client/7phukh/DAM_215fe847-2754-45d2-9f66-b54fe3bfbd6a.pdf
- https://www.dell.com/en-us/blog/open-ethernet-for-ai-nvidia-spectrum-x-with-dell-sonic/
- https://docs.nvidia.com/networking-ethernet-software/knowledge-base/Support/Support-Offerings/Cumulus-Linux-Release-Versioning-and-Support-Policy/
- https://docs.nvidia.com/networking-ethernet-software/cumulus-linux-513/Whats-New/rn/
- https://github.com/nvidia/product-security/blob/HEAD/2026/5817/5817.md
- https://blog.ipspace.net/2025/06/cumulus-linux-gone/
- https://www.shi.com/product/34444595/CISCO-MERAKI-MS210-
- https://www.publicsector.shidirect.com/Product/34444595/CISCO-MERAKI-MS210-48FP-1G-L2CLD-MNGD-48X-GIGE-740W-POE-SWITCH
- https://www.cablesandkits.com/networking/switches/meraki-switches/ms210-48fp-hw/pro-27668/
- https://www.shi.com/product/34444596/CISCO-MERAKI-MS210-
- https://www.shi.com/product/34444597/CISCO-MERAKI-MS210-48LP-1G-L2CLD-MNGD-48X-GIGE-370W-POE-SWITCH
- https://networkequipment.net/products/cisco-meraki-ms210-24-hw-new
- https://www.FS.com/c/nvidia-ethernet-nics-4014
- https://www.fs.com/uk/c/nvidia-ethernet-nics-4014
- https://www.naddod.com/collections/nvidia-networking/infiniband-adapters
- https://www.accio.com/plp/connectx-8-c8240
- https://nvdam.widen.net/content/8h0owe2dhm/original/connectx-datasheet-connectx-8-supernic-update-a4-web-zhCN-3523588-R3.pdf?u=rubmrs&use=c8xan&download=true
- https://www.networkworld.com/article/4115610/nrf-2026-hpe-expands-network-server-products-for-retailers.html
- https://www.networktigers.com/products/r8n85a-hpe-switch-new
- https://www.hpe.com/emea_europe/en/networking/magic-quadrant-wired-wireless.html
- https://www.businesswire.com/news/home/20260520437739/en/Arista-Networks-Positioned-as-a-Leader-in-the-2026-Gartner-Magic-Quadrant-for-Enterprise-Wired-and-Wireless-LAN
- https://www.webdisclosure.com/press-release/huawei-etr-huawei-named-a-leader-in-the-2026-gartner-magic-quadrant-for-enterprise-wired-and-wireless-lan-infrastructure-for-the-fourth-year-in-a-row-p7AObeR7kp0
- https://www.grabnpay.in/products/dell-powerswitch-z9864f-on-64-800gbe-osfp112-high-density-open-networking-switch-with-dual-ac-dc-power-supplies-and-hot-swappable-fans
- https://delltechnologies.com/asset/ko-kr/products/networking/technical-support/dell-powerswitch-z9864f-on-spec-sheet.pdf
- https://bigfastservers.com/products/dell-powerswitch-z9864f-on-with-102-4-tbps-switch-capacity-64x-800-gbe-osfp112-2x-sfp-ports
- https://juaraitsolutions.com/dell-powerswitch-z-series-switches/
---

