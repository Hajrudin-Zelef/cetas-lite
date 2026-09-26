---
id: etape6-trackb-arista-sonic/00-arista-sonic/6-open-verification-items
title: "6. OPEN VERIFICATION ITEMS"
domain: step-6-track-b-arista-sonic-cumulus-data-center-fabric-open-
role: deep-dive
task: reference
actors: ["Broadcom", "Meta", "Microsoft", "Nvidia"]
dates: ["2025-12", "2026-02", "2026-03-12", "2026-05", "2026-05-27", "2026-06", "2026-07", "2026-07-16", "2026-09-22", "2026-11", "2027-04", "2027-07", "2027-08", "2027-11"]
keywords: ["asic", "datacenter", "distribution", "ethernet", "gpu", "gpus", "hyperscaler", "latency", "nvidia", "research", "revenue"]
source: docs/RAG/etape6_trackB_arista_sonic.md
source_anchor: ""
source_lines: [244, 319]
section: "Step 6 — Track B: Arista + SONiC + Cumulus (Data-Center Fabric & Open Networking)"
sha256: f61aecf7b2f7e1af267ccd2718f93eeec2c0c91e2186f374e87e93e94fac26e8
---

# 6. OPEN VERIFICATION ITEMS

## 6. OPEN VERIFICATION ITEMS

1. **NVIDIA Q2 2026 $3.86B figure**: reported by The Next Platform from IDC data (Sep 20, 2026) — likely total Ethernet switching revenue, not directly comparable to the Q1 $2.1B data-center-only figure. Denominator (total vs data-center segment) needs confirmation.
2. **Arista AI revenue targets**: $3.6B (Sep 2026) vs $3.25B (Feb 2026) management guidance — latest supersedes; both vendor-reported.
3. **NVIDIA #1 DC Ethernet share (21.5% Q1 2026)**: sourced to IDC via secondary investment press (ainvest, remio.ai). No primary IDC release was directly fetched; verify against IDC's published tracker before quoting.
4. **ODM $3.56B / 2.43× Q2 2026**: The Next Platform analyst's own estimate (IDC no longer publishes ODM figures). Directional, not official.
5. **Cumulus VX with support contract (Mar 29, 2026)**: single-reader report via ipSpace — unverified.
6. **Tomahawk 6 volume-shipment date (March 12, 2026)**: from marketminute/financialcontent secondary relay; Broadcom's official announcement not directly fetched.
7. **UEC 1.0.3 (July 16, 2026)**: per community wiki citing UEC release notes — not verified against ultraethernet.org directly.
8. **Broadcom "250ns latency" UE Tomahawk claim**: secondary (Medium/Teradata Labs) — vendor-adjacent, treat as marketing until benchmarked.
9. **7060XE7 availability windows** (Q4 2026 air-cooled; Q1 2027 liquid-cooled RV3-L): from official press release and Network World — GA dates can slip; verify before procurement use.
10. **Meta 1.6T RDMA-fabric quote**: from Arista's own press release (official as a quote, but a vendor-authored endorsement).
11. **Ethernet "65% of new AI back-end deployments"**: Tamar Securities feature article (secondary) — no primary source chain.
12. **Cisco SONiC on Nexus 9000 timing**: "soon" per The Register (May 27, 2026) — GA date not confirmed as of Sep 22, 2026.
13. **Arista campus $1.25B 2026 revenue target**: single secondary source (FFJ) — unverified.
14. **Arista customer-concentration figures (55% top-three)**: secondary analyst brief — approximately right directionally, but exact mix is not disclosed by Arista.

## 7. COLLECTION METADATA

- **Research date:** September 22, 2026.
- **Track:** Step 6 Track B — Arista + SONiC + Cumulus (data-center fabric & open networking).
- **Source priority used:** vendor press releases/earnings docs (Arista Q2 2026, Arista 1.6T press release, NVIDIA Cumulus docs); The Register; Network World; The Next Platform (IDC-derived figures); Stordis/Netberg/Aviz distribution release notes; secondary investment press (Zacks, ainvest, CoinCentral, Tickeron) for earnings color and analyst framing — marked [secondary] throughout.
- **Not fetched directly:** primary IDC Ethernet Switch Tracker release; ultraethernet.org spec history page; Broadcom's official Tomahawk 6 announcement; Arista EOS release notes for 2026 builds; community SONiC GitHub release tags for 2026. These are marked [unverified] or [secondary] where cited via relays.
- **Related prior coverage (same project):** Step 5 Track D (servers) covered 800G adoption, Arista Q2 $3.036B, Spectrum-6/Silicon One G300/Tomahawk 6, Cisco $9.3B hyperscaler AI orders, UEC spec status at high level. This track is the authoritative one for Arista/SONiC/Cumulus detail.
- **Nothing sent externally; read-only web research + local file write.**

---

## 8. SUPPLEMENTARY RESEARCH (independent second pass, 2026-09-22)

Additive findings from a second research pass, all provenance-tagged. Nothing above was removed or altered.

### 8.1 Microsoft Fairwater on SONiC (OCP session, ~June 2026)

- An OCP session featuring **Mehak Mahajan (Broadcom)** and **Guohan Lu (Microsoft Azure)** described **Microsoft's Fairwater AI data center as 500,000 GPUs in one data center — "the world's largest AI infrastructure" — running on SONiC** `[vendor-reported]` (YouTube/OCP session, crawled ~June 2026).
- Fairwater fabric design points from the session:
  - **Multi-plane rail topology** connecting 500K+ GPUs, built with **51.2T ASICs** (Tomahawk 5 class).
  - **BGP scaled to 512 sessions per pizza-box switch** in SONiC (FRR upgrades, sub-second convergence).
  - **SRv6 source routing** — giving the NIC control over traffic spreading; deterministic multipathing to avoid ECMP hash collisions (a companion session, "AI Backend: Deploying SRv6 uSID and SONiC for Deterministic Load Balancing," deep-dives Microsoft's production SRv6 uSID deployment in SONiC).
  - **Packet trimming on Tomahawk 5** — recovering from drops before RDMA times out.
  - **High-frequency streaming telemetry at millisecond cadence** via IPFIX + OpenTelemetry.
  - All four technologies (BGP scaling, SRv6, packet trimming, high-frequency telemetry) are **live in SONiC 2025.11 and running in production at Fairwater** `[vendor-reported]`.
  - Next horizon named in the session: **Ethernet for Scale-Up Networking (ESN)** and the **SONiC Scale-Up Working Group** `[vendor-reported]`.
- ⚠️ "World's largest" and the 500,000-GPU figure are session claims by Microsoft/Broadcom speakers; no independent audit `[vendor-reported]`. Session URLs: https://www.youtube.com/watch?v=1SpuQtOBmjE and https://www.youtube.com/watch?v=yB-_4lkSsEs

### 8.2 Community SONiC 202605 release

- **SONiC 202605** (May 2026): **Debian 13 (Trixie)**, SONiC kernel **6.12.41**, **SAI 1.18.1**, **FRR 10.5.4**, **Redis 8.0.2**, **Docker 28.2.1**, **Python 3.13.5** `[secondary]` (btw.media analysis).
- New/evaluated items classified **Alpha** in the release: OpenConfig YANG dial-out telemetry, limited multi-ASIC warm reboot, BMC Redfish workflows, self-encrypting-drive password operations, telemetry VRF binding, event/alarm framework — release inclusion is not a promise of stable behavior across all platforms `[secondary]`.
- Analyst note from the same source: **"supports SONiC" is too broad for procurement** — a meaningful platform statement identifies hardware, ASIC, image provider, SONiC release, SAI/SDK, tested features, and support owner `[secondary]`.

### 8.3 Nokia–Microsoft SONiC collaboration (context, deal from Nov 2024)

- Nokia expanded its multi-year agreement to supply **Microsoft Azure** datacenter networks with a **5-year deal** (announced Nov 2024): **Nokia 7250 IXR-10e** platform for multi-terabit-scale interconnect plus a custom management top-of-rack switch; **SONiC-based Nokia data-center routers and switches** deployed in greenfield sites and in Microsoft's **100GE→400GE migration**; Nokia's global footprint extended to **30+ countries** `[official]` (Nokia/Nasdaq press release).
- David Maltz (Technical Fellow/CVP, Microsoft Azure Networking): "Over the past six years we have worked with Nokia's engineers to develop their routers running SONiC" `[official]`. The collaboration extended to **chassis-based SONiC support for high-capacity datacenter roles** `[secondary]`.

### 8.4 Cumulus Linux release-train update (NVIDIA official docs, crawled Sept 2026)

| Version | EOL |
|---|---|
| 5.12.z | February 2026 |
| 5.13.z | May 2026 |
| 5.14.z | July 2026 |
| 5.15.z | November 2026 |
| 5.16.z | April 2027 |
| 5.17.z | July 2027 |
| 5.18.z | August 2027 |
| 5.9.z (LTS/ESR) | April 2027 |
| 5.11.z (LTS) | November 2027 |

- **Cumulus Linux 5.y.z supports Spectrum-based switches only**; **4.3.z (Broadcom) is in maintenance mode with no new features planned, EOL December 2025** `[official]`.
- **5.13**: new platform **NVIDIA SN5600D** (800G Spectrum-4 DC version) `[official]`.
- **5.15.0**: NVUE command updates to **standardize commands across Cumulus Linux, NVIDIA OS (NVOS), and host-based networking** — NVIDIA recommends customers review NVUE command changes before upgrading `[official]`.

### 8.5 Arista Q2 2026 — additional color (remio.ai analysis, secondary)

