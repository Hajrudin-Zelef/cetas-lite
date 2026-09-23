---
id: etape6-phaseb-smb-networking/01-part-1-ubiquiti-unifi/overview
title: "Part 1 — Ubiquiti UniFi"
domain: part-1-ubiquiti-unifi
role: deep-dive
task: reference
actors: []
dates: ["2026-05", "2026-09", "2026-09-22"]
keywords: ["license", "memory", "pricing", "research", "throughput"]
source: docs/RAG/etape6_phaseB_smb_networking.md
source_anchor: ""
source_lines: [9, 69]
section: "Part 1 — Ubiquiti UniFi"
sha256: 21237563a216cff53315c4737dd4bad016c01202596c7d5506bdcd095aa68b6f
---

# Part 1 — Ubiquiti UniFi

# Ubiquiti UniFi — 2026 Research (SMB/Prosumer Networking)

**Research date:** 2026-09-22. **Scope:** Ubiquiti UniFi product launches, software updates, pricing, financials, security incidents, litigation — calendar 2026. All Markdown deliverables in English per project rule.
**Provenance tags:** [official] = ui.com / Ubiquiti press release / SEC filing; [vendor-reported] = vendor claim via trade press; [independent] = third-party review/lab testing; [secondary] = news/analyst/retailer; [unverified] = single weak source or inferred.

---

## 1. Switch launches 2026 (USW Pro / Aggregation / Enterprise)

### 1.1 Enterprise Campus Switch Core — NEW, September 2026 [secondary]
- Ubiquiti "just launched" an Enterprise Campus Switch Core line (reported ~Sep 5, 2026, Notebookcheck; previously previewed at UWC London 2026, mid-2026 [unverified]).
- Two SKUs:
  - **32× QSFP28 (100G) model**: 6.4 Tbps switching capacity, 3.2 Tbps non-blocking throughput, 1,000 VLANs, 128,000 MAC addresses, dual hot-swap PSUs, 5 hot-swap rear fans, 1.3" status touchscreen, 1U full-depth. Ports configurable 100/40GbE. MC-LAG pairing supported.
  - **SFP28 model (25G)**: 1.8 Tbps non-blocking throughput.
- Management quirk: adoption happens via a QSFP28 port; the rear OOB management port cannot be used for adoption; RJ45 console port present [secondary].
- L3 features per UWC show-floor reporting: full L3 incl. MC-LAG, OSPF, BGP, VRRP [unverified — show-floor claim, InsideWire/YouTube].
- Pricing: QSFP28 model **$5,393** (includes a $394 "memory surcharge"); SFP28 model **$4,314** [secondary, Notebookcheck]. Exact retail SKUs not confirmed — do not invent SKU codes [uncertainty flagged].
- DAC cables ~$35; transceivers $169–$299 [secondary].
- Source: https://www.notebookcheck.net/Ubiquiti-introduces-Enterprise-Campus-Switch-Core-with-3-2-Tbit-s-throughput.1389733.0.html ; UWC London 2026 show coverage: https://www.youtube.com/watch?v=wBk8fXy78Y0

### 1.2 Pro XG family (2026 current lineup — 10G/25G Etherlighting L3)
Not all launched in 2026, but all are the current 2026 Pro XG generation [vendor-reported via retailer specs]:
- **USW-Pro-XG-8-PoE**: 8× 10G RJ45 PoE++ (155 W budget) + 2× SFP+; desktop/wall; **$499** [official list per STH review context].
- **USW-Pro-XG-10-PoE**: 10× 10G RJ45 PoE+++ (400 W budget) + 2× SFP+; **$699** at time of STH review (Sep 2026) [independent].
- **USW-Pro-XG-24** (non-PoE): 16× 10G + 8× 2.5G RJ45 + 2× SFP28; ZAR 26,995 retail in South Africa (Sep 2026) [secondary].
- **USW-Pro-XG-24-PoE**: 16× 10G + 8× 2.5G PoE+++ (720 W budget) + 2× SFP28 [vendor-reported].
- **USW-Pro-XG-48-PoE**: 32× 10G + 16× 2.5G RJ45 (1,080 W PoE budget) + 4× SFP28; L3, 460 Gbps throughput, 256k IPv4 routes, stacking, USP-RPS [vendor-reported].
- **USW-Pro-XG-Aggregation**: 32× SFP28 (25G), 1.6 Tbps switching / 800 Gbps non-blocking, L3, Etherlighting, touchscreen, USP-RPS; ZAR 59,975 retail (Sep 2026) [secondary/vendor-reported].
- ServeTheHome review series (Sep 2026): STH reviewed the Pro XG 8-PoE ($499) and Pro XG 10-PoE ($699), calling the latter "a sweet 10GbE switch" [independent]. Source: https://www.servethehome.com/ubiquiti-unifi-usw-pro-xg-10-poe-review-a-sweet-10gbe-switch/
- STH forum note: STH has not reviewed the USW-EnterpriseXG-24 ($1,300, 24-port 10G), partly because Ubiquiti reportedly sought review/veto rights over review units [secondary/unverified — forum claim].
- Retailer (Sep 2026): https://scoop.co.za (prices in ZAR; availability "in stock" Sep 2026)

### 1.3 100G/400G status
- **100G is shipping**: Enterprise Campus Switch Core 32× QSFP28 (above); earlier USW-Leaf (48× 25G + 6× 100G, pre-2026); Enterprise Campus Aggregation (covered by LTT "100Gb/s Switches from Ubiquiti," Apr 2026) [secondary].
- **400G: no confirmed 400G UniFi product as of Sep 22, 2026** [uncertainty — none found]. Previewed (not launched): "EFG Core" with 4× 100GbE ports shown at UWC London 2026 for enterprise deployments [unverified — YouTube show coverage].

---

## 2. Gateways 2026

### 2.1 UDM-Beast — NEW flagship, launched ~May 2026, $1,499
- Previewed at ISC West 2026 (Apr 2026) as "UDM Beast"; Ubiquiti official launch video "Introducing: UniFi Dream Machine Beast"; YouTube reviews ("just launched") cluster ~early May 2026 [vendor-reported/secondary]. A Sep 21, 2026 news item re-reports the launch — treat May 2026 as GA; Sep item likely recycled [uncertainty flagged].
- Specs [vendor-reported]: 8-core Arm Neoverse N2 @ 2.1 GHz, 16 GB RAM, 128 GB SSD, 14 ports (2× 1GbE RJ45, 8× 10GbE RJ45, 2× 10G SFP+, 2× 25G SFP28), 1.3" LCM touch display, 2× 3.5" NVR drive bays.
- Throughput: **25 Gbps IPS/IDS**; 750+ UniFi devices, 7,500+ clients, 100 HD cameras (40× 4K); zone-/domain-/application-based firewall; dual-unit redundancy (Shadow Mode / VRRP); long-term software updates claimed to cut TCO.
- Positioned above UDM-Pro-Max ($599) and alongside EFG ($1,999) [secondary comparisons].
- Sources: https://en.wedoany.me/shortnews/157024.html ; https://www.youtube.com/watch?v=qutlMzMboC8 (Ubiquiti official) ; https://www.youtube.com/watch?v=ZlX4m4wT1oY

### 2.2 Enterprise Fortress Gateway (EFG) — $1,999 (launched 2024, current 2026 flagship enterprise)
- 25G cloud gateway: 500+ UniFi devices / 5,000+ clients, 12.5 Gbps routing with IDS/IPS, (2) 25G SFP28 + (2) 10G SFP+ + (2) 2.5GbE RJ45 (two WAN-remappable), dual hot-swap PSUs, Shadow Mode HA (VRRP, requires UniFi OS 4.0+ and a paired EFG), license-free NeXT AI Inspection (SSL/TLS decryption), 90 days professional phone support [official/vendor-reported].
- Companion **UXG-Enterprise** (controller-adopted, no on-box Network app) also $1,999 [vendor-reported]. An "upcoming independent gateway counterpart … Gateway Enterprise, same price point" mentioned in an EFG review [unverified].
- Sources: https://www.ispsupplies.com/Ubiquiti-Networks-EFG ; https://www.hostifi.com/blog/uxg-enterprise-vs-enterprise-fortress-gateway-which-one-to-buy

### 2.3 Other gateways (2026 current range; prices = long-standing ui.com list, not re-verified Sep 2026)
- UDM-Pro $379 · UDM-SE $499 · UDM-Pro-Max $599 · Cloud Gateway Ultra $129 · Cloud Gateway Max ~$199 · Cloud Gateway Fiber ~$279 [vendor-reported/secondary — flag: not re-verified].
- 2026 additions in gateway/edge family: **UDR7 (Dream Router 7, Wi-Fi 7)**; **UniFi Express 7 (UX7)** incl. 3-pack whole-home positioning (UWC London 2026); **UDR-5G-Max** (Wi-Fi 7 + dual-SIM 5G failover, with display); **UniFi 5G (U5G)** compact 5G add-on device; "AirWire" (Wi-Fi 7 client adapter) [secondary/unverified].
- New storage/NVR hardware referenced in 2026: **ENVR Core** (Enterprise NVR Core), **EF-Core**, **UNAS Pro 4**, UNVR-G2 (from homelab/ISC West coverage) [secondary/unverified].

---

