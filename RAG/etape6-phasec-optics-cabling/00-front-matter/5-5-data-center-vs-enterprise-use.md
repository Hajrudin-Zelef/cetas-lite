---
id: etape6-phasec-optics-cabling/00-front-matter/5-5-data-center-vs-enterprise-use
title: "5.5 Data center vs enterprise use"
domain: front-matter
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "distribution", "ethernet", "latency", "sol", "training"]
source: docs/RAG/etape6_phaseC_optics_cabling.md
source_anchor: ""
source_lines: [874, 925]
section: "Step 6 — Phase C: Optics, Cabling & Interconnect Infrastructure"
sha256: 4e5412fde66c91775a9c09a692c72aa95753f28973d1df691a2a647faf33bb78
---

# 5.5 Data center vs enterprise use

### 5.5 Data center vs enterprise use

#### 5.5.1 Why fiber dominates server access at 25G+
- **Power**: 10GBASE-T PHY consumes **2–5 W per port at each end**, distance-dependent — roughly **3–4×** an SFP+ solution at ~0.7 W/port regardless of distance [secondary — QSFPTEK, FS.com blog, bonelinks.com]. One source cites 3–6 W/port [secondary — l-p.com, Jan 2026] — minor conflict, treat as 2–6 W range. In dense switches this drives cooling load ("for every watt consumed, typically two additional watts needed for cooling" [secondary — blog.router-switch.com]).
- **Latency**: 10GBASE-T ≈ 2.6 µs per link (block encoding per PHY standard) vs SFP+ fiber ≈ 0.1–0.3 µs [secondary — QSFPTEK]. 10GBASE-T delay characterized as "same order of magnitude as SSD latency" [secondary — blog.router-switch.com].
- **Reach/flexibility**: passive SFP+ DAC limited to ~10 m and cannot be field-terminated (fixed factory lengths, inventory overhead); fiber reaches kilometers; copper structured cabling reaches 100 m (Cat6A) and field-terminates in under a minute [secondary — bonelinks, cn.10gtek.com].
- **Cost structure for <30 m at 25/40G**: Cat8 positioned as cheaper/more flexible than QSFP+ + twinax DAC assemblies for short data-center links, though DAC remains entrenched for ToR server attach [vendor-reported — cablewholesale.com; secondary — QSFPTEK].

#### 5.5.2 Where copper still wins [industry-convention synthesis — tag individual claims]
- **Out-of-band (OOB) management networks**: BMCs expose dedicated RJ45 management ports — Dell iDRAC supports dedicated/shared OOB Ethernet, IPMI, Serial-over-LAN, RAC serial [official — Dell iDRAC9 User's Guide, dl.dell.com]; HPE iLO/Redfish over TCP 443, IPMI 2.0 over UDP 623 [secondary — github.com/nexus-substrate infrastructure doc]. Enterprise practice is a physically/logically separate copper management network (typically 1G) aggregating iDRAC/iLO/IPMI ports [unverified — widely described industry convention; no shipment/stat source found].
- **Serial console servers**: traditional copper/serial console aggregation for switches/routers/PDUs persists alongside IPMI SOL [secondary — Dell docs describe DB9 serial console paths; "copper still wins for console" is convention, tag unverified].
- **PoE endpoints**: Wi-Fi APs, cameras, access control, VoIP, lighting — copper is the only medium delivering power + data on one cable [official-class — IEEE 802.3bt; secondary — multiple].
- **ToR/management and KVM networks**: 1G copper remains standard for switch management ports and KVM [unverified — industry convention].

#### 5.5.3 NBASE-T (2.5G/5G) and Wi-Fi 6E/7 backhaul
- IEEE 802.3bz (ratified 2016) defines 2.5GBASE-T and 5GBASE-T operating on Cat5e or better to 100 m (PHY derived from 10GBASE-T at 1/4 and 1/2 signaling rate) [secondary — Wikipedia]. Enables multi-gigabit AP uplinks without recabling (70 billion meters of Cat5e/6 installed base cited by Cisco at ratification) [secondary — slashgear/afterdawn, 2016].
- Purpose-built for 802.11ac Wave 2 / 802.11ax (Wi-Fi 6) and 802.11be (Wi-Fi 7) APs; eliminates dual-port link aggregation previously needed for >1G AP uplinks [secondary — Wikipedia].
- Vendor guidance: Siemon recommends Class EA/Cat6A or higher for new AP uplink cabling; CommScope/TIA TSB-5021 draft similarly recommended Cat6A for new 2.5G/5GBase-T projects [secondary — cablinginstall, 2015-2016]. This positions Cat6A as the enterprise default for Wi-Fi 7 backhaul [secondary].

### 5.6 PoE (power over Ethernet)

#### 5.6.1 Standards and power levels [secondary — spottersecurity.com; Skyworks white paper; Ethernet Alliance]
| Type | Standard | Common name | Pairs | PSE max out | PD max in |
|---|---|---|---|---|---|
| Type 1 | 802.3af | PoE | 2 | 15.4 W | 12.95 W |
| Type 2 | 802.3at | PoE+ | 2 | 30 W | 25.5 W |
| Type 3 | 802.3bt | PoE++ / 4PPoE / UPOE | 2 or 4 (4-pair mandatory Class 5–6) | 60 W | 51 W |
| Type 4 | 802.3bt | PoE++ | 4 (mandatory) | 90 W | 71.3 W |
Classes: 1–3 (Type 1), 4 (Type 2), 5–6 (Type 3), 7–8 (Type 4). Note: "90 W switch port" ≠ 90 W at device — cable I²R loss accounts for the PSE→PD delta [secondary — maisvch.com, updated Aug 2026]. Extended-power/Autoclass features can raise delivered power if channel length is known [secondary — Microsemi/802.3bt white paper via scribd].

#### 5.6.2 Heat rise in bundles
- IEEE 802.3bt task-force design limit: **<10 °C temperature rise**, modeled at 300 mA in all conductors (~51 W delivered per 100 m) [secondary — eeworldonline/5gtechnologyworld].
- TIA TSB-184-A: recommends **≤15 °C rise at bundle center**; rise depends on bundle size, energized pairs, current (PoE++ = up to 960 mA over 4 pairs), gauge, construction [vendor-reported — Panduit].
- ISO/IEC TS 29125:2017 recommends max 10 °C between bundle center and outer edge [secondary — cablinginstall].
- NEC 2020 Table 725.144 bundle limits at 45 °C ambient (per Cisco doc): 24 AWG/60 °C cable ≤37 cables; 23 AWG/60 °C ≤61; 24 AWG/75 °C ≤91; 23 AWG/75 °C ≤192 [vendor-reported — Cisco UPOE/802.3bt doc via scribd; tag secondary-vendor].
- Bundling is "the most significant thermal factor": inner cables in tight bundles cannot shed heat; open perforated trays beat closed conduit [vendor-reported — VersaTek, crawl Dec 2025].

#### 5.6.3 Cable recommendations for PoE++
- **Cat6A with 23 AWG solid copper** for all 802.3bt applications [vendor-reported — VersaTek]; Cisco: Cat6A 23 AWG or larger for Type 4 at higher data speeds; 22 AWG minimum Cat5e acceptable only for low-data-rate (≤1 Gbps) lighting-style Type 4 [vendor-reported — Cisco].
- **Shielded cable** improves heat distribution (and EMI) [vendor-reported — VersaTek; secondary — cablinginstall].
- Keep bundles ≤24 where possible; use 70–90 °C rated cable in warm/enclosed spaces [vendor-reported — VersaTek — treat as vendor best-practice, not standards text].
- FS.com rates its Cat6A/Cat8 patch cords and panels for af/at/bt [vendor-reported].

### 5.7 Testing

#### 5.7.1 Permanent link vs channel
- **Permanent link**: fixed portion only — patch panel to patch panel (data center) or telecom-room patch panel to work-area outlet (LAN); **excludes** patch/equipment cords; **max 90 m** [secondary — cablinginstall, "Channel, Permanent Link… Oh My!"].
- **Channel**: end-to-end including patch and equipment cords (device to device); **max 100 m total**, of which **≤10 m combined patch cords** (typically ≤5 m per cord; standards recommend 5 m equipment cords) [secondary — cablinginstall; Quabbin; scribd training deck].
- **Cat8 channel**: 30 m total, 2-connector configuration; **24 m permanent link** [secondary — TTI Fiber].
- Certification is end-to-end by category: "a Cat6A permanent link patched with Cat5e cords is, as a system, Cat5e" [secondary — github.com/ronutz structured-cabling guide, updated Sept 2026].
- ANSI/TIA-568.2-D added the **MPTL (modular plug terminated link)** configuration [secondary — cablinginstall].

