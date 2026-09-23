---
id: etape6-phaseb-smb-networking/03-supplementary-complementary-research-pass-september-22-2026/s10-uisp-fiber-xgs-unifi-uisp-pon-line-new-detail-complement
title: "S10. UISP Fiber XGS — UniFi/UISP PON line (new detail; complements §1 fiber)"
domain: supplementary-complementary-research-pass-september-22-2026
role: deep-dive
task: reference
actors: ["Apple", "Google", "United States"]
dates: ["2026-04-02", "2026-05", "2026-09", "2026-09-22"]
keywords: ["ethernet", "license", "memory", "pricing", "research"]
source: docs/RAG/etape6_phaseB_smb_networking.md
source_anchor: ""
source_lines: [483, 533]
section: "Supplementary / Complementary Research Pass — September 22, 2026"
sha256: 62c1ae9dfd4b038452d97e0bb6fa5fe8ec8ffa49c0957451b14d27cacac12c4b
---

# S10. UISP Fiber XGS — UniFi/UISP PON line (new detail; complements §1 fiber)

## S10. UISP Fiber XGS — UniFi/UISP PON line (new detail; complements §1 fiber)

- **UISP-FIBER-OLT-XGS** — 1U XGS-PON OLT: 8× XGS/XG/G-PON ports + 4× 1/10/25G SFP28 uplinks, 10 Gbps symmetric per PON, up to 2,048 clients (256/port), dual hot-swap AC/DC PSU, 70W max, in-band/out-of-band/Bluetooth/serial mgmt, NDAA compliant [vendor-reported via retailer specs]
- **UISP-FIBER-XGS** ONU — 10G symmetric over 20 km SMF, 1× SC/APC XGS-PON WAN + 1× 1/2.5/5/10G RJ45 LAN, USB-C/PoE-powered, 8W max [vendor-reported]
- **UACC-UF-OM-XGS** — 10G XGS/XG-PON SFP+ transceiver for the OLT [vendor-reported]
- Pricing observed (high regional variance — **flag**): ZA **R70,838.85** incl. VAT, in stock (miro.co.za, Sep 2026) [independent]; Kenya KSh 441,719.17 gross, 50+ availability (senetic.co.ke) [independent]; US list not captured — **unknown** [unverified]
- Sources: https://itcart.ae/products/ubiquiti-unifi-uisp-uisp-fiber-olt-xgs ; https://ithub.ae/product/ubiquiti-unifi-uisp-fiber-xgs/ ; https://miro.co.za/search?s=UB-UF-OLT-XGS

## S11. UniFi ecosystem 2026 — G6 Entry intercom + G6 180 camera (new detail)

- **UVC-G6-Entry** (UniFi G6 Entry smart door access/intercom): 5MP main + 8MP package cameras, face recognition, NFC + Apple/Google Wallet, two-way intercom, IP55/IK07, PoE+, integrated with Protect + Access; **released 20 Jan 2026** (Galaxus listing) [secondary]; retailer specs [secondary]
- **UVC-G6-180**: dual-4K 180° camera (16MP combined 7680×2160), face + license-plate recognition, 20 m IR, IP66, microSD slot, NDAA [secondary]
- These expand UniFi beyond networking into physical security/access — relevant to SMB full-stack positioning [independent assessment]
- Sources: https://www.galaxus.at/en/s4/product/ubiquiti-unifi-g6-entry-cable-ethernet-doorbells-door-intercom-system-66467156 ; https://failsafe.mk/wp-content/uploads/2026/02/UniFi-G6-180-Tech-Specs.pdf

## S12. UniFi 400G — status re-check (Sep 22, 2026)

- No confirmed 400G UniFi product as of Sep 22, 2026 — unchanged from main §1.3. The 2026 ceiling is the **ECS-Core (32× 100G)** for core, **ECS-Aggregation (6× 100G uplinks)** for aggregation [official]. EFG Core with 4× 100G remains preview-only per main §9.4 [unverified].

## S13. Remaining gaps (not researched in this pass)

1. UniFi U7 Pro XGS / U7 Pro XG Wall / E7 family: still unconfirmed — no new 2026 evidence surfaced.
2. Netgear 2026 financials / Orbi / Insight managed Wi-Fi business: excluded from scope; not researched.
3. Aruba Instant On 1930R successor (if any): not confirmed.
4. TP-Link Omada Pro controller SKU + pricing: still not found.
5. Regional MSRP vs street comparability: prices above are raw snapshots; no currency normalization applied.

---

## Supplementary / Complementary Research Pass — September 22, 2026 (third wave; S14–S18)

**Scope:** adds only items absent from the existing sections and the S1–S13 supplementary wave above. Omada Pro pricing, USG FLEX H (incl. May 2026 comparison doc), NETGEAR M4350 detail, Zyxel NWA/BE6500 Wi-Fi 7 lineup, UISP Fiber XGS OLT/ONU, UniFi ECS confirmations, memory surcharge, G6 ecosystem, and the 400G re-check were already covered in S2–S12 and are **not** repeated. All claims carry source-class tags per the file convention. No existing section was altered.

## S14. UISP Fiber — XG (non-symmetric) ONU variant + UK price snapshots (was: XGS ONU only, ZA/KE prices only)

- **UISP-FIBER-XG** ONU (distinct from the XGS ONU in S10) [vendor-reported via retailer specs]: XG-PON with 9.95 Gbps downstream / 2.44 Gbps upstream, 1× 2.5GbE RJ45 LAN port, 20 km reach, 8 W max [secondary]
- **UK price snapshots, September 2026** [secondary/independent]:
  - UISP-FIBER-OLT-XGS: **£2,699.99** inc VAT (linitx.com)
  - UISP XGS/XG-PON + GPON coexistence WDM filter: **£238.22** inc VAT (linitx.com)
  - Flag: high regional variance vs S10 (ZA R70,838.85 incl. VAT; Kenya KSh 441,719.17) — no currency normalization applied
- **Note:** UISP Fiber remains the service-provider PON ecosystem, distinct from UniFi Network switching — do not conflate
- Sources: https://itcart.ae/products/ubiquiti-unifi-uisp-fiber-xg ; https://linitx.com/category/ubiquiti-gpon/1267 ; https://ithub.ae/product/ubiquiti-unifi-uisp-fiber-xgs/

## S15. D-Link Nuclias Unity — cloud platform launched April 2, 2026 (was: Oct 2025 hardware controllers only)

- **[vendor-reported]** D-Link launched **Nuclias Unity** (announced Dec 16, 2025; general availability **April 2, 2026**): unified cloud management covering business switches and enterprise APs, with router and IP-camera support planned
- **[vendor-reported]** Features: topology visualization, remote provisioning, automated alerts, role-based access control (RBAC), device lifecycle management, per-port utilization and device-health monitoring, port profiles for bulk VLAN/QoS/security changes, CPU/memory/PoE monitoring, Wi-Fi capacity/channel analytics
- **[vendor-reported]** Positioning vs existing lines: Nuclias Cloud remains supported for DBA APs and DBS switches; Nuclias Unity targets single-site and multi-site organizations; on-premises alternatives are the DNH-1000 hardware controller (license-included management) and DNC-5000 software controller (license-free Hyper management)
- **Assessment:** this is the major 2026 D-Link networking announcement the earlier pass missed; it complements (not replaces) the Oct 2025 DNH-1000/DNH-3000/DNC-5000 hardware line in S8 [independent]
- Sources: https://aapnews.aap.com.au/aapreleases/cision20260401AE25731 ; https://globalwarming.einnews.com/pr_news/880562668/d-link-announces-upcoming-launch-of-nuclias-unity-cloud-network-management-for-single-site-and-multi-site-organizations

