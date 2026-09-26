---
id: etape6-phaseb-smb-networking/06-supplementary-complementary-research-pass-2-september-22-202/supplementary-complementary-research-pass-2026-09-22-indepen
title: "Supplementary / Complementary Research Pass — 2026-09-22 (independent research passes A/B/C)"
domain: supplementary-complementary-research-pass-2-september-22-202
role: deep-dive
task: reference
actors: []
dates: ["2026-09-22"]
keywords: ["research", "ethernet", "memory", "optics", "pricing", "throughput"]
source: docs/RAG/etape6_phaseB_smb_networking.md
source_anchor: ""
source_lines: [1127, 1155]
section: "Supplementary / Complementary Research Pass #2 — September 22, 2026"
sha256: da14badff01f69100ec8df04e3a114ee77320c7ca530c689fb639a373ac8fc38
---

# Supplementary / Complementary Research Pass — 2026-09-22 (independent research passes A/B/C)

## Supplementary / Complementary Research Pass — 2026-09-22 (independent research passes A/B/C)

**Scope note:** This section contains material from three independent research passes run on 2026-09-22 against gaps in the original Parts 1–2. It is distinct from the sibling S1–S18 and R1–R6 supplementary waves added separately to this file; some topics overlap across waves (flagged where material), but this pass's research is independent and adds new facts, prices, and firmware details. Parts 1–2 above are untouched. Every fact carries a provenance tag: [official] / [vendor-reported] / [independent] / [secondary] / [unverified]. Identifiers, dates, specs, prices and verbatim source URLs are given; uncertainties are flagged explicitly.

---

### Pass A — Ubiquiti deep gaps: 400G, fiber/PON, EF-Core launch, software window Sep 10–22, homelab consensus, supply/surcharges, Pro AV

Research conducted September 22, 2026; all facts dated on/before that cutoff. **[official]** = ui.com/Ubiquiti-owned; **[vendor-reported]** = reseller/distributor specs; **[independent]** = press/reviewer/institutional; **[secondary]** = aggregator/analysis; **[unverified]** = could not be confirmed.

#### A.1. UniFi 400G (QSFP-DD) status

- No UniFi switch with 400G (QSFP-DD) ports was announced, launched, or road-mapped on/before Sep 22, 2026, in any source searched. **[unverified — absence of evidence; not proof of absence]**
- The fastest UniFi ports found in-market remain 100G QSFP28: Enterprise Campus Switch Core (32× QSFP28, $5,393 incl. $394 memory surcharge), EAV-Fiber (2× 100G QSFP28), EF-Core gateway (4× 100G QSFP28). **[independent]** https://www.notebookcheck.net/Ubiquiti-introduces-Enterprise-Campus-Switch-Core-with-3-2-Tbit-s-throughput.1389733.0.html · **[vendor-reported]** https://arcip.com.au/shop/eav-fiber-ubiquiti-unifi-eav-fiber-enterprise-audio-video-1u-aggregation-switch-with-20-x-10g-sfp-and-2-x-100g-qsfp28-800-gbps-switching-ptp-smpte-st-2110-sdvoe-aes67-timing-engine-ocxo-clock-gps-grandmaster-input-9724
- Searched: "Ubiquiti UniFi 400G QSFP-DD switch announced 2026", "store.ui.com 400G QSFP-DD switch UniFi", community/Reddit threads — no hits on a UniFi 400G product or roadmap statement. **[unverified]**

#### A.2. Fiber / PON 2026: UISP line (no UniFi-branded OLT/ONU found)

- **No UniFi-branded OLT or ONU (no "UOLT-4" or "UONU-XG" SKUs) was found** as of Sep 22, 2026 — all Ubiquiti PON products found are UISP/UFiber-branded, managed by the UISP app (not the UniFi Network app). The example SKUs from the brief could not be verified; treat them as unconfirmed. **[unverified]**
- **UISP-FIBER-OLT-XGS** — 8× XGS-PON ports, 4× 25G SFP28 uplinks, 2048-client capacity (256:1 split ratio), 1U rackmount, 2× hot-swappable redundant AC/DC PSUs, max 70 W, managed via UISP app 1.5.7+. MSRP **$2,499**. **[vendor-reported]** https://www.publicsector.shidirect.com/product/47460833/Ubiquiti-UISP-Fiber-OLT-XGS · https://hw.ddfu.org/en/printout/printproduct?StiId=293801 · https://www.jw.com.au/product/ubiquiti-uisp-fiber-olt-xgs-2048-cc-25g-sfp28
- **UISP-FIBER-XGS** (XGS-PON ONU/ONT) — 1× SC/APC XGS-PON (10 Gbps symmetric) + 1× 10GbE RJ45 (100/1G/2.5G/5G/10G), bridge and router modes, USB-C power, managed via UISP app 1.4.9+; works **only** with Ubiquiti OLTs. Street pricing: €120.65 ex-VAT (PL) **[vendor-reported]** https://en.cdr.pl/p9664,ubiquiti-uisp-fiber-xgs-uisp-fiber-xgs-pon-ont-terminal-1x-10ge.html · £100.79 inc VAT, 18 in stock (UK) **[vendor-reported]** https://linitx.com/category/ubiquiti-gpon/1267 · €150.90 (PL) / R3,347.65 inc VAT (ZA) **[vendor-reported]** https://interprojekt.pl/en/p/ubiquiti-uisp-fiber-xgs-eu.html · https://miro.co.za/12-fibre-optics-ubiquiti-ufiber-gpon-cpe/7103-ubiquiti-uisp-fiber-xgs-gpon-onu-ont-810084691236.html
- **UISP-FIBER-XG** (XG-PON ONU/ONT) — 10 Gbps down / 2.5 Gbps up, 1× 2.5GbE RJ45 + 24V PoE-in. £93.59 inc VAT, 24 in stock (UK). **[vendor-reported]** https://linitx.com/category/ubiquiti-gpon/1267
- **Wave-Fiber-ONU** (new "WaveFiber" GPON CPE line) — 1× SC/APC GPON WAN (ITU-G.984) + 1× 1/2.5 GbE RJ45, 24V passive PoE-in (5W max), 76.5×76.5×26.4 mm / 78 g, desktop/wall mount, OMCI-managed via UISP (remote provisioning, firmware upgrades, GEM port encryption, NAT/firewall/VLAN). £47.21 (UK, in stock, "Despatched today"); 5-pack **Wave-Fiber-ONU-5** £230.09, 38 in stock. **[vendor-reported]** https://linitx.com/product/ubiquiti-wavefiber-gpon-cpe-optical-network-unit-wave-fiber-onu/18297 · https://www.firstshop.co.za/collections/wireless-adapters/products/ubiquiti-uisp-wavefiber-onu-sc-apc-2-5gbps-ethernet-gpon-optical-network-media-converter-wave-fiber-onu-372469 · https://www.getic.com/product/ubiquiti-wavefiber-onu
  - Flag: the WaveFiber ONU is widely stocked at resellers in Sep 2026 and looks like a 2026 addition, but no official launch announcement or date was found — exact launch date **[unverified]**.
- Accessories spotted at UK retail: **UACC-UF-OM-XGS** (XGS/XG optical transceiver, 10 Gbps, 20 km) £131.99 — awaiting restock; **UACC-UF-WDM-XGS** (XGS/XG-PON + GPON coexistence WDM filter) £238.22, 9 in stock. **[vendor-reported]** https://linitx.com/category/ubiquiti-gpon/1267
- Legacy GPON hardware still sold in 2026: **UF-OLT** (8× GPON, 2× SFP+, 1024 ONUs, MSRP $1,499 at 2017 launch) and **UF-OLT-4** (~$1,190–$1,291), UF-Nano / UF-LOCO ONUs, managed via UISP/UNMS. **[vendor-reported]** https://globenewswire.com/news-release/2017/07/31/1065075/0/en/Ubiquiti-Networks-Launches-Plug-and-Play-UFiber-GPON-Platform.html · https://www.howardcomputers.com/accessories/detail.cfm?id=S20654584

#### A.3. Enterprise Firewall Core (EF-Core, "EFG Core" 4×100G) — launched, in market

