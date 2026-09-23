---
id: etape6-phasea-vendors-dc/00-front-matter/supplementary-complementary-research-pass-round-2-2026-09-22
title: "Supplementary / Complementary Research Pass — round 2 (2026-09-22)"
domain: front-matter
role: reference
task: reference
actors: ["Huawei", "Nvidia"]
dates: ["2026-06", "2026-09-22", "2027-12", "2028-11"]
keywords: ["research", "compute", "datacenter", "disaggregated", "distribution", "ethernet", "license", "nvidia", "rack-scale", "revenue", "rubin", "vera rubin"]
source: docs/RAG/etape6_phaseA_vendors_dc.md
source_anchor: ""
source_lines: [629, 690]
section: "Step 6 — Phase A: Enterprise Data-Center Switching Vendors"
sha256: b329287ffbb7d1e4126dcd008d3c68d1f28e39a5960f47f910e406c59e69ba93
---

# Supplementary / Complementary Research Pass — round 2 (2026-09-22)

## Supplementary / Complementary Research Pass — round 2 (2026-09-22)

**Scope note:** A first supplementary pass (§§A–M above) already exists, written by a parallel research process and independently spot-checked (its IDC Q2 2026 market figures were verified against IDC's published Q2 2026 Ethernet Switch Tracker blog — all numbers match: Cisco $5.4B/28.7% overall; NVIDIA $2.5B +181.1%/20.4% DC; Arista $2.5B +37.6%; Huawei $1.5B/8.2%; HPE 6.1% incl. Juniper; DC segment $12.3B +64.5%; 800GbE 41.2% of DC revenue). This round-2 pass adds only material not present in the base §§1–8 or in §§A–M. Nothing above was modified.

### N. Dell SmartFabric OS10 lifecycle and licensing status (new detail)

- Dell's published **SmartFabric OS10 Product Lifecycle Policy**: Standard Support lasts **36 months from the RTW of each Major release train** (e.g., 10.5.2.x, 10.5.3.x); minor/maintenance/patch releases do not reset the support window [official — Dell OS10 lifecycle policy PDF, dl.dell.com].
- Current train status (Dell SmartFabric OS10 Hardware Compatibility List, 2026): **10.6.1.x — End of Maintenance November 2028**; 10.6.0.x — December 2027 [official — Dell KB 000192674].
- OS10 is **not end-of-life**: it remains actively documented and supported across Z/S/N/E/MX platforms, while Enterprise SONiC is Dell's declared strategic/disaggregated NOS. Dell's official guidance frames OS10 → Enterprise SONiC as a **migration path** (tech brief H19795.1, June 2026), not a forced sunset [official — Dell docs].
- Licensing: an OS10 image download ships with a **120-day trial license**; a **perpetual license** is required beyond the trial (installed per switch, no license change needed for upgrades) [official — Dell OS10 install/upgrade guide].
- Interpretation: Dell is running a **dual-NOS strategy in 2026** — OS10 maintained for installed base, Enterprise SONiC for new AI-fabric/disaggregated wins (closing the "OS10 strategy" ambiguity left in base §1.4).

### O. Aruba CX 8325 / 8350-class campus-core and DC-ToR detail (new)

Base §3.3 listed CX 9300/8360 gaps and the supplementary pass priced the 9300/8360. The 8325/8350 tier was not detailed:

- **CX 8325 series** (AOS-CX, 1U): **6.4 Tbps, 2,000 Mpps**; 48Y8C (48× 25G + 8× 100G), 32C (32× 100G), 48Y8C/32C positioned for **campus core/aggregation and DC ToR/EoR**; VSX redundancy, BGP/OSPF/VRF-lite, EVPN-VXLAN (dynamic VXLAN for campus + DC segmentation), NAE, REST/Python [official — HPE datasheets via HPE Store].
- **CX 8325P-32C**: telco variant with **PTP (1588v2) Class C, SyncE (G.8262.1), GNSS** — supports 5G RAN boundary-clock use cases [official — HPE Store Taiwan].
- **CX 8325H**: half-width fixed switches (2.16 Tbps: 18× 1/10/25G + 4× 40/100G; 4 Tbps: 16× 40/100G) for edge/colocation DC [official — HPE Store].
- **CX 8360 v2 series**: up to **4.8 Tbps, 1,786 Mpps**, compact 1U; HPE Smart Rate (1/2.5/5G) + 10/25/40/50/100G; low-density **MACsec** on 10/25G ports and 40/100G uplinks (8360 48-port); VSX, BGP/OSPF/VRF/EVPN/VXLAN/IPv6 [official — HPE 8360-16Y2C v2 datasheet].
- Positioning note: 8325/8360 fill the **25/100G leaf and campus-core tier below the 9300 spine** (25.6 T) — the Aruba stack below the Juniper-side AI-fabric flagships, consistent with HPE's "Aruba owns the campus" split [secondary].

### P. SmartFabric Manager + Enterprise SONiC on Spectrum-X convergence (new 2026 development)

- Dell's 2026 AI Factory refresh: **Enterprise SONiC Distribution by Dell now supports NVIDIA Spectrum-X platforms alongside Cumulus OS**, and **SmartFabric Manager extends to Dell's Enterprise SONiC on Spectrum-X** — aiming at hyperscale-like networking setup with fewer manual steps [secondary — ciol.com, 2026].
- Rack-scale context: PowerEdge XE9812 (Vera Rubin) with Spectrum-X/Quantum networking under the Dell AI Factory banner; Dell AI Factory with NVIDIA customer base exceeded **5,000 in Q1 FY27** [secondary — investment memo/GitHub wiki; unverified].
- Significance: Dell is converging its two management stories (SmartFabric Manager ↔ Enterprise SONiC ↔ Spectrum-X) — the software counterpart to the hardware dual-sourcing covered in base §1.5/C [secondary].

### Q. Leadership in the covered networking units (2026) — finding: no changes located

- **NVIDIA networking**: Kevin Deierling continues as **Senior Vice President of Networking at NVIDIA** (appearing as the networking lead in 2026 Dell/NVIDIA AI events) [secondary — SiliconANGLE, 2025-10 event coverage]. No 2026 change of NVIDIA networking leadership was located.
- **Dell networking (ISG)**: Arun Narayanan, **SVP of Compute and Networking Portfolio Management** (Infrastructure Solutions Group), remains Dell's public face on networking/AI-fabric portfolio (GTC 2025 interviews, 2026 coverage) [secondary — CRN].
- **HPE Networking**: Rami Rahim leads the combined unit (covered in step-6 track A). No 2026 leadership change located.
- **Meraki/Cisco**: no 2026 leadership change in the Meraki business unit located.
- No M&A in 2026 affecting these four vendors' switching businesses beyond Astrix/Galileo (already covered in §J — both are security/observability, not switching) [gap — deal-completion tracking continues].

### R. Round-2 verification log (new open items)

1. OS10 post-10.6.x roadmap (a 10.7 train or OS10 EOL announcement) — nothing published as of 2026-09-22 [gap].
2. Dell Enterprise SONiC subscription dollar amounts — still unpublished; only structure confirmed (§A) [gap].
3. SmartFabric Manager on Enterprise SONiC/Spectrum-X — GA status vs roadmap phrasing not pinned down [unverified].
4. CX 8325/8350 2026 refresh — none located; continuation assumed [unverified].
5. 2026 leadership/executive changes in Dell Networking, NVIDIA networking, Meraki — none found, but executive moves are easy to miss in trade press [gap].

### S. Round-2 sources (verbatim URLs)

- https://www.dell.com/support/kbdoc/en-ca/000192674/smartfabric-os10-hardware-compatibility-list
- https://dl.dell.com/content/manual2232496-dell-smartfabric-os10-product-lifecycle-policy.pdf?language=en-us
- https://www.dell.com/support/kbdoc/en-us/000221543/dell-networking-smartfabric-os10-getting-started-basics-guide
- https://www.dell.com/support/manuals/en-us/smartfabric-os10-emp-partner/ee-upgrade-downgrade/installing-smartfabric-os10?guid=guid-9be9-4abb-99cf-b2671091f3e0&lang=en-us
- https://www.hpe.com/psnow/generateDDS/HPE Aruba Networking CX 8360-16Y2C v2 16 port 25G SFPSFP+SFP28 2-port 100G QSFP+QSFP28 Switch Digital Data Sheet-PSN1014652420BYRU.pdf?oid=1014652420&cc=BY&lc=RU&softroll=0&prelaunch=false&print=§ion=&prelaunchSection=&softrollSection=&deepLink=&utm_source=&utm_medium=&utm_campaign=&utm_content=&utm_term=&isLinearized=false&contentDisposition=attachment
- https://buy.hpe.com/tw/zh_TW/networking/switches/fixed-port-l3-managed-ethernet-switches/aruba-8325-switch-products/aruba-cx-8325-%e4%ba%a4%e6%8f%9b%e5%99%a8%e7%b3%bb%e5%88%97/p/r9f64a
- https://buy.bluum.com/aruba-jl627a-b2e-8325-32c-ethernet-switch/
- https://www.idc.com/resource-center/blog/ethernet-switch-market-surges-43-4-to-18-9b-in-2q26-as-ai-infrastructure-demand-drives-record-datacenter-spending/
- https://www.ciol.com/enterprise/dell-advances-enterprise-ai-with-nvidia-integration-10785615
- https://siliconangle.com/2025/10/09/dell-ai-data-platform-event-join-thecube-dellaidataplatform/
- https://www.crn.com/news/ai/2025/how-dell-lenovo-and-supermicro-are-adapting-to-nvidia-s-fast-ai-chip-transitions

**Round-2 collection metadata:** read-only web research (browser_search + one page fetch, 2026-09-22); no live-browser visits; nothing sent externally. No identifiers guessed. All new facts carry provenance tags; new open items are listed in §R. Existing sections §§1–8 and §§A–M were not modified.

---

