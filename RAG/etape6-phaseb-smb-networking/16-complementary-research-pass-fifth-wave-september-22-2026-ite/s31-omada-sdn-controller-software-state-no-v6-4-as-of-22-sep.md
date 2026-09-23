---
id: etape6-phaseb-smb-networking/16-complementary-research-pass-fifth-wave-september-22-2026-ite/s31-omada-sdn-controller-software-state-no-v6-4-as-of-22-sep
title: "S31. Omada SDN Controller software state — no v6.4 as of 22 Sept 2026"
domain: complementary-research-pass-fifth-wave-september-22-2026-ite
role: deep-dive
task: reference
actors: ["Microsoft"]
dates: ["2026-05", "2026-07", "2026-07-01", "2026-08"]
keywords: ["optics", "packaging", "pricing", "scout"]
source: docs/RAG/etape6_phaseB_smb_networking.md
source_anchor: ""
source_lines: [2564, 2612]
section: "Complementary Research Pass — Fifth Wave (September 22, 2026), Items S25–S33"
sha256: a84571c8bcec234271230e37feddee103b2252d658afa5cb9f48e19ebf8a3806
---

# S31. Omada SDN Controller software state — no v6.4 as of 22 Sept 2026

## S31. Omada SDN Controller software state — no v6.4 as of 22 Sept 2026

### Findings `[official]`/`[vendor-reported]`/`[independent]`
- **No Omada SDN Controller v6.4 exists yet.** The latest GA is **6.2.10.17** (FreeBSD FreshPorts port updated 1 July 2026; TP-Link blog guides confirm "current version as of writing: 6.2.10.17" on 1 July 2026; requires MongoDB 7 and Java/OpenJDK 25) — https://www.freshports.org/net-mgmt/omada6/ ; https://github.com/kdpuvvadi/blog/blob/HEAD/_posts/2026-07-01-omada-sdn-controller-ubuntu-26-04.md
- Pre-release track: **Omada 6.2.14-adapted firmware** for ER7412-M2 V1 (1.2.0, released 6 July 2026), DR3220v/DR3650v series (6 July 2026).
- TP-Link Business Community firmware sticky (crawled Sept 2026): latest pre-release listed is **Fusion 2.5G V1 1.0.30 (26 Aug 2026)**; nothing newer in the public pre-release channel — https://community.tp-link.com/en/business/threads/topic/255644
- **Omada Cloud Standard / Essentials** remain the free cloud-controller tiers (see earlier S1/S19).
- Hardware note: **ER7212PC V2 ships with built-in Omada SDN Controller 6.2** (pre-release firmware updated 2 Apr 2026); ER7212PC V1 remains on controller 6.0.
- Router/gateway pre-release firmware adapted to SDN 6.2.x continues through August 2026 (ER605 V2 2.4.0 Apr 2026, ER7206 V2 Mar 2026, ER707-M2 Mar 2026, ER8411 1.4.0 Jun 2026).
- Community install guides: Ubuntu 26.04 LTS / Debian 13 support documented July 2026; homelab Kubernetes deployments (Omada + oauth2-proxy + Pocket-ID) documented Sept 2026 — https://github.com/alpar-t/homelab/blob/HEAD/config/omada-controller/README.md

### Gap status
- **RESOLVED:** v6.4 confirmed nonexistent; 6.2.10.17 is the July 2026 baseline; pre-release 6.2.14 tracked.
- S19's 6.3.0.45 (4 Sept 2026) reference from the earlier pass stands but was not independently re-confirmed in this wave — mark **[unverified]** until re-checked against TP-Link release notes.

---

## S32. UniFi Network Application — 10.7 status and release-train notes

### Findings `[secondary]`
- **No public UniFi Network Application 10.7 release notes were found** as of 22 Sept 2026. The latest documented version remains **10.6.106 (10 Sept 2026)** from the earlier pass (T2).
- Dependency note: third-party packaging documentation records that the **10.2.x line moved to Java 25** (UniFi add-on maintainer notes, GitHub zglate/addon-unifi, maintained Sept 2026) — https://github.com/zglate/addon-unifi/blob/HEAD/MAINTENANCE.md — relevant for self-hosted (Docker/Home Assistant) operators captured in homelab sections.
- Community-mirrored release notes (McCann Tech / evanmccann.net) index only older 5.x–9.x era notes; the canonical source remains community.ui.com/releases.

### Gap status
- **Retained:** 10.7 release date and changelog; UniFi OS 5.x state for Dream Machine gateways; whether 10.7 changes MongoDB/Java requirements.

---

## S33. This wave's gap closure summary

| Gap (from earlier passes) | Status after this wave |
|---|---|
| UniFi U7 Pro XGS launch status | **RESOLVED** — launched May 2026, $199 XG / $299 XGS; reseller stock confirmed; E7/E7 Campus confirmed shipping |
| Grandstream GWN7672 street price | **RESOLVED** — UK £145 ex VAT / CA $228 / IN ₹14.5–20k / KSA SAR 1,017 |
| Omada Pro line 2026 | Extended — EAP775-Wall/Outdoor, EAP783, EAP787 documented with prices; EAP776 not found (do not use) |
| UniFi UCG-Max pricing | **RESOLVED** — 9 regional price snapshots Sept 2026 |
| NETGEAR Insight subscription 2026 | Extended — pricing documented; **Basic free-tier availability now [unverified]** (possible 2026 discontinuation) |
| MikroTik 2026 hardware | **RESOLVED** — full MWC 2026 slate (hEX Pro, mAP ax, nRAY gen2, Scout, A42GO-HbeP, CRS816, hAP be³) |
| Omada controller v6.4 | **RESOLVED as nonexistent** — latest GA 6.2.10.17; pre-release 6.2.14; 6.3.0.45 [unverified] |
| UniFi Network 10.7 | **Retained** — not found; 10.6.106 remains latest documented |
| Zyxel Nebula version after 20.10 | **Retained** — no 20.11/20.12 found this wave |
| EnGenius ECW516L price | **Retained** — no price data found this wave |
| Zyxel XS1935-12HP price | **Retained** — no price data found this wave |
| Zyxel USG FLEX H successors | **Retained** — no 2026 refresh found this wave |

---

*End of Fifth Wave (S25–S33). File remains append-only; earlier sections untouched. Ready for Phase C (FS.com / optics / cabling) on Anicet's go.*
