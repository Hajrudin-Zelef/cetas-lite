---
id: collect-261001-mikrotik/mikrotik/etape6-trackc-huawei-mikrotik-2
title: "Step 6 — Track C: Huawei + MikroTik (Networking Hardware)"
domain: mikrotik
role: reference
task: reference
actors: ["AWS", "China", "EU", "Huawei", "Microsoft", "United States"]
dates: ["2026-03", "2026-04", "2026-05-26", "2026-09", "2026-09-22"]
keywords: ["aws", "compute", "ethernet", "graviton", "hyperscaler", "incident", "license", "nand", "pricing", "revenue", "scout"]
source: docs/RAG/collect-261001-mikrotik/etape6_trackC_huawei_mikrotik.md
source_anchor: ""
source_lines: [58, 132]
sha256: 85964a753d67bfb1ca92db7fe6e65e385a0aa8484c4f49d9035b6b4368e9c639
---

# Step 6 — Track C: Huawei + MikroTik (Networking Hardware)

**Dell'Oro — High End Routing and Aggregation, Q2 2026** (reported ~Sep 10, 2026) **[independent — https://www.prnewswire.com/news-releases/high-end-routing-and-aggregation-market-grew-25-percent-in-2q-2026-according-to-delloro-group-302872936.html]**:
- Segment revenue **+25% YoY**; vendor **direct sales to cloud providers +94% YoY** (hyperscaler AI infrastructure push).
- Trailing-4-quarter vendor rank: **High End Routing & Aggregation: 1 Cisco, 2 Huawei, 3 Nokia**; **Cloud Provider segment: 1 Cisco, 2 HPE Juniper, 3 Nokia**; **Communication Service Provider segment: 1 Huawei, 2 Cisco, 3 Nokia**.

**Fortune Business Insights (enterprise networking)**: Huawei listed among top vendors with ~18% enterprise networking market share (vs Cisco 22%) **[secondary — https://www.fortunebusinessinsights.com/enterprise-networking-market-105887]**.

### 1.5 International market status

- Huawei's networking growth remains anchored in **China and select emerging markets** (Middle East, Africa, Latin America, Southeast Asia) — per IDC's characterization of its router business **[independent]**.
- Export-control environment: Huawei continues to operate its networking business under US entity-list restrictions (imposed 2019); no lifting or change in 2026 found in sources fetched — treat current status as **[unverified]** for any specific 2026 regulatory change.
- No 2026 Huawei telecom/datacom wins in US, EU-5, Japan, Australia, or UK markets were found in fetched sources — consistent with the ongoing restrictions **[secondary]**.

---

## 2. MIKROTIK — 2026 product launches

### 2.1 MWC Barcelona 2026 launches (March 2026)

MikroTik announced a batch of new products at Mobile World Congress, Barcelona, March 2026 **[secondary — https://www.youtube.com/watch?v=KATjJn3RYUk; official newsletter — https://box.mikrotik.com/d/b5b340fd548c40fe97a3/files/?p=%2Fnews_132.pdf&dl=1]**:

**hEX Pro / hEX Pro PoE** (product codes **RB660UG+S+Me** / **RB660UPG+S+Me**) — next-generation SMB routers:
- CPU: **quad-core ARM Cortex-A73 @ 2.2 GHz**, ARM 64-bit; **2 GB DDR4** RAM; 128 MB NAND storage **[official — https://box.mikrotik.com/seafhttp/files/a87a30d3-5b64-46ce-b3fb-bbab549f71bd/hEX%20PRO%3AhEX%20PRO%20POE.pdf]**.
- Ports: **4× 2.5G Ethernet + 1× 10G Ethernet + 1× 10G SFP+**; **1× M.2 slot** for container storage (Docker-compatible containers); 1× USB 3.0 type A; switch chip **IPQ-9570**; RouterOS v7, license level 4.
- hEX Pro PoE adds **PoE-Out 802.3af/at** on Ether2–Ether5 (max 147 W total) and **MACSec 802.1AE** support on Ether1–Ether5.
- Powering: 24–57 V (DC jack + PoE-In 802.3bt); operating temperature **-40°C to +70°C**; max power 16 W (non-PoE) / 147 W (PoE).
- Positioning: 10G-uplink SMB router with on-device containers — a step up from the classic hEX (E50UG refresh, $59.95 MSRP, 880 MHz dual-core / 256 MB RAM) **[official/secondary — https://multilink.us/shop-by-brand/mikrotik/mikrotik-routers/]**.

**nRAY gen2** — 60 GHz wireless link reaching **up to 10 km** with 5 GHz backup link **[secondary]**.

**mAP ax** — pocket-size **travel router with Wi-Fi 6** **[secondary]**.

**MikroTik Scout** — **mesh communication device that works without internet coverage**; positioned for agriculture and rural areas **[secondary]**.

**A42GO-HbeP** — **Wi-Fi 7 access point with integrated GPON ONT in a single device**; described as the flagship product for ISPs (combines fiber termination + Wi-Fi 7 in one box) **[secondary]**.

### 2.2 Ampere Altra CCR family (Newsletter #132, April 2026)

MikroTik's April 2026 newsletter (#132) announced its **most powerful Cloud Core Router family to date, powered by Ampere Altra ARM server CPUs** — up to **64 CPU cores** and **64 GB DDR4**, aimed at core/aggregation roles in regional data centers, enterprise/campus networks, government/financial infrastructure, and telecom **[official — https://box.mikrotik.com/d/b5b340fd548c40fe97a3/files/?p=%2Fnews_132.pdf&dl=1]**.

**CCR3232-16XG-4DS-2DQ** (first model):
- **2× 200G QSFP56 + 4× 50G SFP56 + 16× 10G Ethernet** ports; **Marvell DX7335 switch chip with L3 offload**; **MACSec 802.1AE hardware encryption on all ports**; full internet routing tables, BGP, MPLS, advanced filtering; container support; RouterOS v7 **[secondary — https://www.store.mikrotikcanada.ca/ethernet-routers/750-ccr3232-16xg-4ds-2dq-4752224000026.html]**.
- Retail listing: **~$2,795** (Sil Micro, "COMING SOON", in stock per listing page) **[secondary — https://www.silmicro.com/cloud-core-routers/]** — price not confirmed by MikroTik; treat as [unverified].
- Context: community had spotted RouterOS builds for "AMPERE" on MikroTik's download page as early as Nov 2023 (forum thread), indicating a long testing cycle **[secondary — https://forum.mikrotik.com/t/mikrotik-ampere-cpu-coming-soon/171188]**.

### 2.3 ROSE Data Server RDS2216

"High-performance, all-in-one storage, networking, and container platform for enterprise environments" running **ROSE (RouterOS Edition for storage/compute)** **[secondary — https://www.techradar.com/pro/this-is-amazons-first-foray-in-servers-and-certainly-not-the-last-microtik-franken-router-is-powered-by-the-aws-graviton-1-arm-cpu]**:
- **16-core AWS Graviton 1 ARM CPU @ 2 GHz**, **32 GB DDR4**; **16 ports**: 2× 100G QSFP28, 4× 25G SFP28, 4× 10G SFP+, 2× 10G Ethernet; **20× U.2 NVMe storage slots**; NVMe-TCP block device export, encryption layers; USB ports.
- Container-ready: runs MinIO, Nextcloud, Shinobi, Frigate and other OCI-compliant containers; **no subscriptions or paywalls** (MikroTik's standard model).
- Distinctive green 1U chassis.

### 2.4 CRS switch line: 2026 pricing snapshots

- **CRS310-8G+2S+IN**: 8× 2.5G Ethernet + 2× SFP+; ARM 32-bit 98DX226S @ 800 MHz; RouterOS v7 license L5; ~**R4,090** (South Africa, Sep 22, 2026) **[secondary — https://www.comx-computers.co.za/MT-RBCRS310-8G-2S-IN-MikroTik-Cloud-Router-Switch-8x-Buy-p-290412.php]**.
- **CRS304-4XG-IN**: 4× 10G Ethernet (no modules needed) + 1× 1G management; fanless; ~**NZD 368.70** **[secondary — https://www.pbtech.com/pacific/product/NETMKT1502/MikroTik-Cloud-Router-Switch-CRS304-4XG-IN-with-4]**.
- **CRS317-1G-16S+RM**: 16× SFP+ + 1G management, dual PSU; ~NZD 871.30 **[secondary — https://www.pbtech.com/pacific/product/NETMKT1248/MikroTik-CRS317-1G-16SRM-Cloud-Router-Switch-CRS31]**.
- **CRS354-48G-4S+2Q+RM**: 48× 1G + 4× 10G SFP+ + 2× 40G QSFP+; 336 Gbps switching capacity, 235 Mpps; dual redundant PSU **[secondary — https://www.pbtech.com/product/NETMKT12401/MikroTik-CRS354-48G-4S2QRM-Cloud-Router-48-Port-Gi]**.
- MikroTik's pricing model remains: hardware price includes RouterOS license + **free software updates for the life of the product (minimum 5 years)**; no subscriptions **[official]**.

### 2.5 RouterOS v7 — 2026 release timeline

As of September 22, 2026 **[secondary — https://endoflife.ai/article-mikrotik-routeros-cve-2026-86060-only-6-49-21-fixed]**:

| Line | Channel | First release | Status |
|---|---|---|---|
| 7.21 | Long-term (previous) | Jan 12, 2026 | Last build 7.21.5 (Jul 6, 2026); in vulnerable range for Sep CVEs |
| 7.22 | Stable (superseded) | Mar 10, 2026 | Superseded by 7.23; in vulnerable range |
| 7.23 | **Long-term** | May 26, 2026 | Current LTS; 7.23.4, 7.23.5 (Sep 4), 7.23.6, 7.23.7 (Sep 16) |
| 7.24 | **Stable** | Aug 17, 2026 | Current stable; 7.24.2, 7.24.3, 7.24.4 |
| 7.25 | Beta | — | 7.25beta3 (Sep 3, 2026) |
| 6.49 | 6.x stable/long-term | Oct 6, 2021 | 6.49.21 (Sep 3, 2026) — last maintained 6.x line |
| 6.48 and earlier | None | — | End-of-life; no fixes |

### 2.6 September 2026 security incident ("MikroTrick")

