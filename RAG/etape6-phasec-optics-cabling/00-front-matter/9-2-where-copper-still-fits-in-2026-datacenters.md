---
id: etape6-phasec-optics-cabling/00-front-matter/9-2-where-copper-still-fits-in-2026-datacenters
title: "9.2 Where copper still fits in 2026 datacenters"
domain: front-matter
role: reference
task: reference
actors: ["EU", "United States"]
dates: []
keywords: ["cost", "dsp", "energy", "ethernet", "latency", "optics", "power delivery", "pricing"]
source: docs/RAG/etape6_phaseC_optics_cabling.md
source_anchor: ""
source_lines: [1407, 1454]
section: "Step 6 — Phase C: Optics, Cabling & Interconnect Infrastructure"
sha256: 04405d6ebab694d1f69f225ad1bbdb21b304091e2f6ea806ffeafd240db7d533
---

# 9.2 Where copper still fits in 2026 datacenters

### 9.2 Where copper still fits in 2026 datacenters

- **Out-of-band management fabric**: BMC (Dell iDRAC, HPE iLO), IPMI/Redfish over LAN, serial-over-LAN console, console servers, KVM-over-IP, switched PDUs (SNMP/telemetry), and building-automation/monitoring networks remain overwhelmingly RJ45 copper, typically 1 Gbps [secondary]. Dell iDRAC9 docs describe a "Dedicated Gigabit Ethernet port" and best practice of a separate management network [official — Dell iDRAC9 User's Guide]. Copper's advantages here: every device ships with a copper NIC, 100 m reach suits row/end-of-row OOB aggregation, and the management network does not need dataplane bandwidth [unverified — analysis, no hard source].
- **10GBASE-T top-of-rack in smaller/enterprise DCs**: Cat6A structured copper still deployed for server ToR where RJ45 ecosystem, auto-negotiation down to 1G, and installer familiarity matter [vendor-reported — FS.com Cat8/40GBASE-T and Cat6A materials]. Auto-negotiation allows incremental switch/server upgrades without forklift replacement [vendor-reported — CommScope].
- **PoE in DC facilities**: cameras, access points, sensors, and some DCIM environmental sensors run on PoE/PoE++ over Cat6A in the facilities/support network (see §9.5) [vendor-reported — FS.com; secondary].
- **Not** the server dataplane in hyperscale: production server attachment at 10G+ is SFP+/SFP28 DAC, AOC, or fiber (see §9.4) [secondary].

### 9.3 Price data points (single-unit list, Sept 2026, before volume discounts)

**Cat6A patch cords (S/FTP, 26 AWG, TIA-568.2-D / ISO Class EA) — FS.com** [vendor-reported]:
- 3 m: US$7.40 / €5.95 / SGD 7.74 / AUD 8.47 (SKU/P-N C6ASFTPSGPVC family)
- 1.5 m: SGD 5.12; 1.8 m: SGD 5.56; 1.2 m: AUD 4.73
- 0.3 m 24-pack: US$69.00 (≈$2.88/ea, multi-pack); 1.2 m 10-pack (UK): £31.20 (≈£3.12/ea)
- FS claims these exceed TIA-568.2-D and are "Fluke tested," PoE/PoE+/PoE++ capable, 550 MHz sweep [vendor-reported]

**Other vendors — Cat6A patch** [vendor-reported]:
- Monoprice SlimRun 30 AWG UTP Cat6A 10-packs: 3 ft $26.89, 5 ft $31.78, 7 ft $37.33, 1 ft $26.06 (≈$2.6–$3.8/ea); 5-packs: 5 ft $18.58, 3 ft $17.15
- StarTech 1 ft (0.3 m) Cat6a S/FTP 27 AWG, LSZH: CAD $14.99 single (PC-Canada), volume tiers to $12.89 at 100+

**Patch panels — FS.com** [vendor-reported]:
- FDMP Cat6a 110-style shielded recessed 24-port 1RU (P/N FDMP-C6A1UFPS24110): AUD 149.60 / €105.91 (€89.00 VAT excl.)
- FDMP Cat6a shielded coupler recessed 24-port 1RU (FDMP-C6A1UFPS24): AUD 121.00 / €85.68 / £81.60 / SGD 112.27
- FS C6AP-S24FT1U Cat6a coupler 24-port 1RU: €117.81 (€99.00 VAT excl.) / SGD 153.69
- FS Cat6 24/48-port range (listed as 10G Base-T capable): 24-port unshielded from £49.20–£66.00; 24-port shielded £60.00–£94.80; 48-port unshielded £84.00–£100.80; 48-port shielded £109.00–£130.80

**Keystone jacks** [vendor-reported]:
- Primus Cable: Cat6A shielded US$6.99; Cat6A unshielded $3.50; Cat8 toolless 2 GHz $11.99
- CableLeader Cat6A shielded toolless $5.06; Telecom Specialties (Vertical Cable) Cat6A shielded $8.86 / unshielded $4.89
- trueCABLE Cat6A toolless shielded 2-pack $17.99 (≈$9.00/ea); L-com specialty panel-mount Cat6A kit $47.19; ACT (EU) Cat6A shielded toolless from €41.75

**Bulk cable — GAP**: no single-unit 305 m-box or per-meter Cat6A pricing surfaced (cablesupply.com lists a 1000 ft shielded Cat6A bulk reel but the price was not shown in the fetched text). No Cat8 bulk-reel pricing found. Flagged as missing.

### 9.4 10GBASE-T vs SFP+ DAC/fiber for short server links

- **Power**: Early 10GBASE-T PHYs 6–8 W; modern PHYs 3–5 W; varies with link distance 2–5 W per port per end [secondary — l-p.com table; QSFPTEK; FS.com comparison]. SFP+ fiber module ≈0.7–1.5 W per port; passive DAC <0.5 W [secondary — QSFPTEK; EE Times]. FS's own 10G SFP+ passive DAC (3 m) is spec'd at max 0.1 W [official — FS.com product spec]. Net: 10GBASE-T burns ~3–4× the power of SFP+ [vendor-reported — FS.com], generating 2–4× the heat per port, raising inlet temperatures, fan load, and limiting port density [secondary].
- **Latency**: 10GBASE-T PHY specifies 2.6 µs transmit–receive pair (block encoding, LDPC FEC, echo-cancellation DSP) with block-size budget under 2 µs; SFP+ simplified electronics ≈0.1–0.3 µs per link (0.1 µs fiber, 0.3 µs DAC per QSFPTEK's table) [secondary — QSFPTEK, FS.com].
- **Cost comparison**: FS 10G SFP+ passive DAC 3 m: US$20.00 / €19.04 (€16.00 VAT excl.) / SGD 25.07 [vendor-reported]; PBTech lists the same FS DAC at AUD ~30.22–30.78 [vendor-reported]. A Cat6A structured link of equivalent function costs ≈$7.40 (3 m FS cord) plus a share of panel/jack/switch PHY — but note the PHY cost sits in the switch/NIC, and 10GBASE-T switch ports carry a persistent power premium. No single-unit price for a 10GBASE-T SFP+ RJ45 module was found — flagged gap.
- **Why hyperscalers moved away**: power density at thousands of ports (thermal/cooling cost), higher latency, inability to reach modern port densities on 10GBASE-T PHYs, and DAC/AOC undercutting on both cost and power for ≤7 m intra-rack/adjacent-rack links [secondary — EE Times 2009-era analysis (dated, but the power/density argument has persisted and worsened at higher speeds); QSFPTEK; FS.com]. DAC/AOC recommendation: ≤3 m passive DAC; 5–30 m AOC; >100 m optics [secondary — network-switch.com]. 10GBASE-T's remaining edge: familiar RJ45, auto-negotiation to 1G/10G, 100 m reach over installed Cat6A, no transceiver inventory [vendor-reported — FS.com, CommScope].
- Mitigation note: IEEE 802.3az Energy Efficient Ethernet can cut 10GBASE-T link power >50% at low utilization [secondary — EE Times].

### 9.5 PoE standards & bundling/heating

- **IEEE 802.3af (Type 1)**: 15.4 W PSE / 12.95 W PD, 2-pair; **802.3at (Type 2, PoE+)**: 30 W / 25.5 W, 2-pair; **802.3bt Type 3 (PoE++)**: 60 W / 51 W, 4-pair; **802.3bt Type 4**: 90 W / 71.3 W PD, 4-pair [official — IEEE 802.3bt-2018 figures via Skyworks white paper and CablingInstall; classes 1–8]. CONFLICT/nuance: FS.com marketing says Type 4 "up to 90W… theoretical maximum 100W" and "at least 71W" at the PD [vendor-reported] — treat 90 W PSE / 71.3 W PD as the standards numbers; "100 W" appears in vendor chipset/PoE++ marketing (e.g., trueCABLE jack "up to 100W" [vendor-reported]) and is not the IEEE figure.
- bt adds: mandatory classification refinements, Autoclass (cable-loss-aware power budgeting), power demotion, short MPS standby (20 mW vs 200 mW, enabling IoT) [secondary — Ethernet Alliance white paper]; supports 2.5/5/10GBASE-T data rates [vendor-reported — FS.com].
- **TIA TSB-184-A** (Guidelines for Supporting Power Delivery Over Balanced Twisted-Pair Cabling): caps temperature rise at **15 °C** in the center of a cable bundle; provides modeled rise tables by category, bundle size, current, and open-air vs conduit [official — via Panduit technical bulletin and DOE Connected Lighting study, secondary]. Example modeled values (PoE Type 4, 24-cable bundle): Cat5e 7.91 °C (air) / 11.29 °C (duct); Cat6 6.00/8.67; Cat6A 5.13/7.11; Cat8 2.90/4.74 [secondary — Lightera bulletin reproducing TSB-184-A Table 2]. Practical rule: bundles ≤24 cables stay under 15 °C for Type 4 [secondary — Lightera]. **Cat6A is TIA's recommended cable for new PoE installations**; shielded (metallic elements) recommended to mitigate rise [secondary — DOE report; Panduit; FS.com].
- **Addendum TSB-184-A-1** covers 28 AWG patch cords: TIA-568.2-D permits 22–28 AWG cords, but 28 AWG cords are limited to 15 m in a channel with length derating per Annex G for their higher insertion loss [official — via Quabbin/electronics360 secondary].
- NEC 725.144 (2017) sets bundle-size/current limits by AWG [secondary — DOE report].

