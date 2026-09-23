---
id: etape6-phasea-vendors-dc/00-front-matter/e-white-box-odm-landscape-fills-base-gap-on-dell-vs-odm
title: "E. White-box / ODM landscape (fills base gap on Dell vs ODM)"
domain: front-matter
role: reference
task: reference
actors: ["Alibaba", "Broadcom", "Huawei", "Nvidia"]
dates: ["2026-02", "2026-05", "2026-07", "2026-09-20", "2026-11", "2027-04", "2027-07", "2027-08", "2027-11"]
keywords: ["datacenter", "ethernet", "hyperscaler", "license", "licenses", "nvidia", "revenue"]
source: docs/RAG/etape6_phaseA_vendors_dc.md
source_anchor: ""
source_lines: [379, 429]
section: "Step 6 — Phase A: Enterprise Data-Center Switching Vendors"
sha256: 676b77fc40a9cf3617e3c5feb047d09e451c3c9fbf6ee491d9b5a0006a1f1f86
---

# E. White-box / ODM landscape (fills base gap on Dell vs ODM)

### E. White-box / ODM landscape (fills base gap on Dell vs ODM)

- **Dell'Oro Q2 2026** (AI Back-end Networks report): **Celestica retained the leading position in Ethernet AI back-end switch sales**, followed closely by NVIDIA; Arista third; **Cisco gained the most share** (ranked fourth). 800G = vast majority of AI back-end shipments/revenue; 1.6T began sampling, expected to ramp H2 2026 [independent — Dell'Oro via thefastmode.com].
- Next Platform estimates the **collective ODM (white-box) Ethernet switch business at ~$3.56B in Q2 2026 (+2.43× YoY)**; IDC no longer publishes ODM-collective figures [secondary — Next Platform, 2026-09-20]. This is an estimate, not IDC data.
- **Enterprise SONiC 4.6 added 13 new ODM/white-box platform entries** (vs 4.5.1): Celestica **DS4100** (Tomahawk 4, 16× 800G), UfiSpace **S9321-64EO** (Tomahawk 5, 64× 800G), Edgecore **DCS511** (Tomahawk 4, 32× 400G), Supermicro **SSE-T8164S Rev2** (Tomahawk 5, 64× 800G), Micas 6940-64OC-R (Tomahawk 5), Aria ARIA-SW-800G-64-TH5, Alpha Networks SNK6010-320F (Tomahawk 5), Edgecore AIS800-32D, UfiSpace S9311-64D/S9301-32DB/S7801-54XS/S6301-56STP, Micas M2-W6510-48GT4-RA [secondary — STORDIS engineering breakdown, ~Aug 2026].
- Dell's counter-positioning: open-NOS (Enterprise SONiC) + OEM silicon (Broadcom + NVIDIA) + validated AI Factory blueprints, rather than a proprietary-NOS closed stack [vendor-reported].

### F. Cumulus Linux release train through 5.18 (extends base §3.3)

From NVIDIA's official support-policy KB [official — docs.nvidia.com]:

| Cumulus version | EOL |
|---|---|
| 5.9.z (LTS) | April 2027 |
| 5.11.z (LTS) | November 2027 |
| 5.12.z | February 2026 |
| 5.13.z | May 2026 |
| 5.14.z | July 2026 |
| 5.15.z | November 2026 |
| 5.16.z | April 2027 |
| **5.17.z** | July 2027 |
| **5.18.z** | August 2027 |

- Cumulus Linux **5.y.z supports Spectrum-based switches only**; latest Broadcom-based release remains **4.3.z** (in maintenance mode; no new features) [official].
- A **Cumulus Linux 5.17 User Guide** exists in the CumulusNetworks docs repo (draft state) [official — GitHub cumulusnetworks/docs].

### G. Market data update: IDC Q2 2026 + Dell'Oro AI back-end

**IDC Q2 2026 Ethernet tracker** (published ~Sep 21, 2026) [independent — IDC press release]:

- Total Ethernet switch market: **$18.9B (+43.4% YoY)**; datacenter segment **$12.3B (+64.5% YoY)** (~65% of revenue).
- **Cisco: $5.4B (+36.2%), 28.7% share — #1 overall**; datacenter $2.2B (18.2% DC share). Router revenue $1.6B (+31.6%, 34.7% router share).
- **NVIDIA: $2.5B (+181.1%), 20.4% DC share — #1 branded vendor in datacenter Ethernet**; Spectrum-X "continues to win hyperscaler and AI-native cloud provider deployments at scale."
- **Arista: $2.5B (+37.6%)**; 13.4% total share, 18.7% DC share (~90.9% of revenue from DC).
- **Huawei: $1.5B (+28.6%)**, 8.2% share; router revenue $1.3B (+21.6%, 29.9% router share).
- Note the **unresolved discrepancy** with Next Platform's analysis ($3.86B / "2.8×" for NVIDIA, HPE post-Juniper $1.15B +6.1%, Huawei $1.55B): Next Platform's figures include its own ODM-inclusive estimates and differ from IDC's published release [secondary — flagged, both sources cited].

**Dell'Oro Q2 2026**: AI back-end switch sales **surpassed front-end networks for the first time**; Ethernet further extended its AI back-end lead; market to stay **supply-constrained (not demand-constrained) for 1–2 years** [independent — Dell'Oro via PRNewswire/ADVFN].

### H. Meraki licensing model — detail (partially closes base §7 item 12)

- Meraki operates **three licensing models**: **co-term** (single org-wide expiration date), **per-device licensing (PDL)**, and **subscription** [secondary — meraki-dashboard-exporter API conformance notes, corroborated by Meraki docs].
- Cisco **changed the co-term → per-device conversion process**: conversions are now handled via a revised support-guided process; organizations on subscription fall through to per-device behavior [official — Meraki documentation, "Meraki Per-Device Licensing – Configuration"].
- **1-day licenses exist only under PDL**; one org can hold max **50K 1-day licenses** [official — Meraki docs].
- Price points found (2026):
  - **LIC-MS350-24X-1YR**: list $557.95 (street ~$262.84); 3YR list $1,255.38 (street ~$591.40); 5YR list $2,092.30 (street ~$985.66); 10YR list $4,184.60 (street ~$1,971.31) [secondary — cloudwifiworks.com].
  - **MX licenses (annual, July 2026)**: MX67 Enterprise $275 / Advanced Security $595; MX95 $1,399 / $2,199; MX250 $5,999 / $7,999 [secondary — linktly.com].
  - MX67 3-year Advanced Security typically $450–550; annual per-device costs range **$300–$1,500** by tier [secondary — CacheGuard, May 2026].
  - MR57 (Wi-Fi 7 AP): $1,299 hardware + $199/yr Enterprise license (Q2 2026 list) [secondary — Alibaba buying guide].
- License expiry is hard: **Meraki devices stop working when the license expires** — cloud-dependency and license lock-in remain the key TCO caveat for Meraki [secondary — CacheGuard, May 2026].

