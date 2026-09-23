---
id: etape6-phaseb-smb-networking/02-part-2-tp-link-omada-d-link-zyxel-engenius-grandstream/2-d-link
title: "2. D-Link"
domain: part-2-tp-link-omada-d-link-zyxel-engenius-grandstream
role: deep-dive
task: reference
actors: []
dates: ["2026-05", "2026-09", "2026-09-01", "2026-09-07", "2026-09-17"]
keywords: ["lean", "license", "pricing"]
source: docs/RAG/etape6_phaseB_smb_networking.md
source_anchor: ""
source_lines: [256, 305]
section: "Part 2 — TP-Link / Omada, D-Link, Zyxel, Engenius, Grandstream"
sha256: 9240cc557815915c52486716ac4526f7853c9969f7b95744e91a76d91c506121
---

# 2. D-Link

## 2. D-Link

### 2.1 2026 switch launches

- **DXS-3130 series** (announced ~Nov 2025, current 2026 flagship) [vendor-reported]: 24-port 10GbE Multi-Gig (1/2.5/5/10G) L3 stackable, 4× 25GbE SFP28 uplinks, PoE++ up to 60W/port, 790W standard budget expandable to 1,440W with redundant PSU, up to 9 units stacked / 200 Gbps stack bandwidth, OSPFv2/v3, RIP, VRRP, ECMP, PBR, PIM/IGMP/MLD, 6 kV surge protection, NDAA/TAA compliant, Lifetime Warranty (NA). Management: web/CLI/SNMP/sFlow + optional **Nuclias Connect Hyper (license included, no subscription fee)** or Nuclias Controller DNH-1000 [vendor-reported].
  Source: https://www.einpresswire.com/article/865984960/d-link-north-america-launches-dxs-3130-series-10gbe-multi-gig-stackable-layer-3-managed-switches
- **DMS-1250 series** (announced ~Oct 2025 per PR, eight 2.5G smart-managed models + unmanaged DMS-108P/1016/1024) [vendor-reported]: 10→28-port 2.5G downlinks, 10G SFP+ uplinks, up to 90W/port PoE (bt), 6 kV surge protection on all 2.5G ports, LACP, static routing, RADIUS/TACACS+, Safeguard Engine, MIT (Made-in-Taiwan) [vendor-reported].
  Source: https://www.sdxcentral.com/news/d-link-unveils-25g-switches-to-bridge-networking-speed-gaps/ ; https://www.prnewswire.co.uk/news-releases/d-link-unveils-2-5g-multi-gigabit-switches-for-versatile-edge-connectivity-302539550.html
- DMS-1250 specs (retail confirmations): DMS-1250-28P = 24× 2.5G PoE++ (90W max on some SKUs; PoE+ 475W budget) + 4× 10G SFP+, L3 Lite, Nuclias Controller/D-View 8/Web/CLI/SNMP, lifetime warranty, NDAA/TAA [secondary].
- **Pricing** [independent]: DMS-1250-28P — SHI $898.00 (MSRP $1,104.99); AU RRP A$2,699.95 (street ~A$2,252.60 tristaronline, A$2,380.95 eBay AU, A$1,419.84 nsoffice — **wide variance, flag**); DMS-1250-28 (non-PoE) — AU RRP A$1,499.95, street A$1,239.13, UK £247.79 ex VAT (comms-express) — note UK price seems low vs AU; treat as unverified single-retailer listing; DMS-1250-18P AU RRP A$1,899.95 [independent].
  Sources: https://www.shi.com/product/50481144/D-Link-DMS-1250-28P ; https://www.digitalreviews.net/news/pressers/d-link-dms-1250-multi-gigabit-smart-switch-anz/ ; https://tristaronline.com.au/products/d-link-dms-1250-28p-switch

### 2.2 Nuclias cloud management — 2026 status

- Nuclias Connect Hyper: license included, **no subscription fee** (per DXS-3130 launch materials) [vendor-reported].
- Nuclias Controller (DNC) / Nuclias Hub (DNH-1000): on-premises/hub management for DMS-1250 etc. [vendor-reported].
- No major 2026 Nuclias platform-version announcement found in sources searched; D-View 8 remains the central NMS [independent]. **Flag: Nuclias 2026 roadmap thinly sourced.**

### 2.3 Market status

- D-Link (TWSE: 2332) continues to lean on "Made-in-Taiwan", NDAA/TAA compliance and lifetime warranty as differentiators for public sector/education; positioning the DMS-1250 series at SMB and DXS-3130 at mid-enterprise [vendor-reported]. No 2026 market-share figures found — **flag**.

---

## 3. Zyxel

### 3.1 Nebula — 2026 updates

- **Nebula 20.10** (announced 2026-09-01) [vendor-reported via press]: identity federation, passwordless access, expanded topology visibility, MSP-oriented tools (centralized event logs across orgs, one-click device replacement with config/license auto-transfer). Some functions require Nebula Pro Pack or MSP Pack license [independent/secondary].
  Sources: https://technode.global/2026/09/03/zyxel-nebula-20-10-identity-federation-network-management/ ; https://securitybrief.asia/story/zyxel-adds-identity-tools-to-nebula-cloud-platform (2026-09-07)
- **Nebula 19.30** (~May 2026) [secondary]: MSP multi-tenant control — consolidated cross-customer overview, org-level change logs, role-based access (Admins & Teams), onboarding/offboarding workflows, customer grouping. 100+ Nebula-compatible products across wired/wireless/security/mobile FWA [secondary].
  Source: https://securitybrief.asia/story/zyxel-nebula-19-30-boosts-msp-multi-tenant-control
- **Licensing** [official]: Nebula Base Pack = free; Plus Pack and Pro Pack = paid per-device subscriptions. SHI lists Nebula Plus Pack 1-yr at **MSRP $19.99 / street $14.00** (part LIC-NCC-PLUS-1YR) [independent]; Pro Pack pricing not captured — **flag**.
  Sources: https://www.zyxel.com/library/assets/solutions/nebula/pdf/nebula-licensing-table.pdf ; https://www.shi.com/product/42545001/Zyxel-Nebula-Plus-Pack
- 100+ products in Nebula ecosystem; Zyxel claims 1M+ businesses across 150 markets served [vendor-reported].

### 3.2 2026 switch launches

- **XS1935 series** (announced ~September 2026) [independent/secondary]: 10GbE on every port, compact 10/12-port formats; three models — XS1935-10, XS1935-12HP (PoE++ 90W/port, 400W budget), XS1935-12F (all-fiber). Nebula cloud or local management. Target: Wi-Fi 7 APs, NAS, AV-over-IP. **Pricing/availability not announced** [unverified].
  Source: https://thetechrevolutionist.com/2026/zyxel-xs1935-10gbe-poe-switches-smb-networks.html
- **GS1915 series** (announced 2026-09-17) [secondary]: 8-port cloud-managed PoE/non-PoE smart switches under "Just Connect" SMB push; Nebula app onboarding; entry-level pricing implied, **exact price not published — flag** [unverified].
  Source: https://technologyreseller.uk/just-connect-small-businesses-with-zyxels-new-easy-to-use-switch-powered-by-nebula/
- Older references: XGS1935 series (10G SFP+ uplink Lite-L3, 2024), XGS2220 L3 access (2023, street $849.99–$1,899.99) — included only as context, not 2026 launches.

### 3.3 Market position

- Zyxel continues a cloud-first SMB/MSP strategy with the Nebula platform (free Base tier as wedge, paid Plus/Pro for MSP features); 2026 updates are MSP/identity-focused rather than hardware-focused. AI-powered security/wireless positioning ("AI- and cloud-powered") [vendor-reported/independent assessment].

---

