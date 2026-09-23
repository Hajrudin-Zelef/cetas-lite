---
id: etape6-phasea-vendors-dc/01-round-5-complementary-research-pass-2026-09-22-deep-datashee/s5-meraki-ms210-ms225-street-pricing-closes-round-4-oo-item-
title: "S5. Meraki MS210/MS225 street pricing — closes round-4 §OO item 4 (new)"
domain: round-5-complementary-research-pass-2026-09-22-deep-datashee
role: deep-dive
task: pricing
actors: ["EU", "Nvidia", "United States"]
dates: ["2026-08", "2026-08-25", "2026-09", "2026-09-14", "2026-09-17", "2026-10"]
keywords: ["pricing", "agentic", "compute", "datacenter", "dram", "latency", "license", "neocloud", "nvidia", "rack-scale"]
source: docs/RAG/etape6_phaseA_vendors_dc.md
source_anchor: ""
source_lines: [2192, 2242]
section: "Round-5 complementary research pass (2026-09-22) — deep datasheet and radix detail"
sha256: 7c18f31fccf2f0ad52b2b2f4c7886a38a223c91574f9d59fb8ef3aad8f8580aa
---

# S5. Meraki MS210/MS225 street pricing — closes round-4 §OO item 4 (new)

### S5. Meraki MS210/MS225 street pricing — closes round-4 §OO item 4 (new)

The MS210/MS225 access tier was the last unpriced Meraki stack family in this file. September 2026 anchors for **MS225-48LP-HW** (48× 1G PoE/PoE+ 370W, 4× SFP+ uplinks):

- US: **MSRP $9,650.79**, street **$4,553.00** [secondary — SHI Direct, publicsector.shidirect.com].
- EU: **€3,527.00** (out of stock) [secondary — bechtle.com/de-en]; companion listing **€3,374.00** [secondary — bechtle.com/be]; Meraki license anchors — **LIC-MS225-48LP-3YR €918.99**, **LIC-MS225-48LP-5YR €1,517.00** [secondary — bechtle.com].
- AU/NZ: AUD **$6,864.00** (sold out, ex tax; AU power cord and license sold separately) [secondary — cloudmanagedwifi.com.au]; NZD **$11,856.52 ex GST** ($13,635.00 inc GST), zero stock across all NZ branches, listing created 2026-09-17 [secondary — pbtech.co.nz].
- Gray-market/refurbished: **$560.99** CK-certified refurbished [secondary — cablesandkits.com] — roughly 12% of street price for a new unit, consistent with the ~40–55% discount-to-list pattern at authorized resellers for the MS390/MS355/MS450 (rounds 3–4) and the steep gray-market drops noted there.
- Spec anchors: 176 Gbps switching capacity, 127.98 Mpps forwarding, **80 Gbps dedicated physical stacking** (2× QSFP), 9,578-byte jumbo frames, 32K MAC, RPS2300 redundancy support, PoE available simultaneously on all ports [secondary — cablesandkits.com; secondary — pbtech.co.nz].
- MS210 vs MS225 positioning (new): the MS225 line (5 models, up to 740W PoE+) is the stackable 10G-uplink access family; MS210 is the non-stackable/legacy sibling — both remain current in the 2026 lineup per base §4.1 [secondary — reseller taxonomy; independent assessment].

### S6. Meraki MS250 street pricing — portfolio completion (new)

The MS250 (L3 branch-access, hot-swappable PSU) had only a passing mention in earlier passes. **MS250-48LP-HW** (48× 1G PoE+ 370W, 4× SFP+, L3 with OSPFv2, warm spare/VRRP, 80 Gbps stacking, 176 Gbps switching) September 2026 anchors:

- NZ: **$15,262.61 ex GST** ($17,552.00 inc GST), zero stock across all branches, listing created 2026-09-14 [secondary — pbtech.co.nz].
- AU: **$12,778.21** (RRP $26,029.12; ~51% discount-to-list), $14,056.03 incl GST [secondary — techforgood.com.au].
- US new/unclaimed: **$2,995.00** (new, unclaimed, hardware only) [secondary — networkequipment.net]; **$1,253.00** used unclaimed [secondary — spwindustrial.com].
- US refurbished: **$2,528.85** [secondary — techatlantix.com]; **$549.99** CK-certified refurbished [secondary — cablesandkits.com]; **$220.00** refurbished (90-day warranty — flagged: unusually low, gray-market channel) [secondary — serversupply.com].
- Spec detail: 2 GB DRAM, 256 MB flash, 32K MAC, 4,094 VLANs, up to 9,216-byte jumbo, 2 hot-swappable variable-speed fans (front-to-back), OSPF and warm spare noted as mutually exclusive [secondary — techatlantix.com; secondary — pbtech.co.nz].

### S7. Meraki MS410/MS425 aggregation heritage — portfolio completion (new)

Zero prior coverage in this file; the MS400 series is the older fiber-aggregation family that the MS450 supersedes in Cisco's current positioning:

- **MS410-16-HW**: 16× Gigabit SFP + 4× 10G SFP+ uplink; **MS410-32-HW**: 32× Gigabit SFP + 4× 10G SFP+ uplink; both L3 aggregation, physical stacking up to 8 members with **160 Gbps stacking bandwidth** (MS410/425), 72 Gbps switching capacity, 1.2 µs latency, 96K forwarding table, front-to-back airflow [secondary — 1stop.com; secondary — it-market.com].
- **MS425**: the 10G SFP+ aggregation sibling (MS425-16/32 class, up to 136W full load vs 85W MS410) — referenced in the MS400 family datasheet set [secondary — 1stop.com]. Meraki's official documentation hub lists MS410/MS425 overview and specification documents as current product information [official — documentation.meraki.com "Product Information" page; official — Meraki MS switch comparison chart, telecom4good.org mirror].
- Pricing: MS410-32-HW refurbished **$6,678.26** (was $14,840.58 — ~55% discount) [secondary — priceblaze.com]; MS410-16-HW **$3,565.30** (new, in stock) [secondary — serverblink.com]; MS410-16-HW used **€83.30** "UNCLAIMED device without license" (flagged: gray-market, no license — not a market price) [secondary — it-market.com].
- L3 capabilities: DHCP failover, VRRP, OSPFv2, static IP routing [secondary — 1stop.com; secondary — it-market.com].

### S8. AOS-CX 10.17 feature delta — closes round-4 §OO item 2 (new)

Beyond the CVE-2026-23813 security bulletin (round-4 §KK), HPE's 10.17 release-update technical assets document the following feature deltas [all secondary — HPE "AOS-CX 10.17 Release Update" video technical assets, YouTube]:

- **DHCP server support on all VLANs** (previously restricted) [secondary — youtube.com, "Support DHCP on all VLANs"].
- **Multicast IVRL dynamic RP enhancements** [secondary — youtube.com, "Multicast IVRL Dynamic RP enhancements"].
- **ARP management enhancements** [secondary — youtube.com, "ARP Management Enhancement"].
- **Tunneling sub-interface underlay support** [secondary — youtube.com, "Tunneling Sub-Interface Underlay Support"].
- **Ability to mirror a specific interface of a LAG** (finer-grained port mirroring) [secondary — youtube.com, "Ability to Mirror a specific interface of a LAG"].
- Note: an older (2020-era) NetworkWorld piece on AOS-CX's edge-to-cloud strategy was surfaced but is **not** a 10.17 source — excluded from the delta list; flagged to avoid misdating [secondary — networkworld.com, 2020].

### S9. Cisco expands Secure AI Factory with NVIDIA — rack-scale era (new, August 2026)

A major 2026 development for the Cisco angle of this file, announced **2026-08-25** [vendor-reported — Cisco via newswire.telecomramblings.com; secondary — SiliconANGLE; independent — dpa-AFX via tradingview.com]:

- Cisco expanded its **Secure AI Factory with NVIDIA** architecture with a Supermicro partnership, adding **rack-scale computing** (liquid- and air-cooled high-density systems, NVL72-class support) targeting neocloud and sovereign-cloud demand [vendor-reported].
- The full-stack architecture is **NVIDIA Cloud Partner (NCP) reference compliant**, combining **Cisco Silicon One-based switches and NVIDIA Spectrum-X networking**, unified by **Cisco Nexus One**, with the entire infrastructure — compute sleds, switches, security appliances — managed via the **Cisco Cloud Control** platform (connects to round-4 §R4-B) [secondary — siliconangle.com, 2026-08-25].
- Cisco will begin offering Supermicro systems inside the Secure AI Factory in **October 2026** [independent — dpa-AFX].
- Jeetu Patel (President and Chief Product Officer, Cisco): "We are at the beginning of one of the largest datacenter buildouts in history… it starts with the right infrastructure: compute and networking, delivered as an integrated solution that is easy to deploy and secure from day one" [vendor-reported]. Justin Boitano (VP Enterprise AI, NVIDIA): full-stack infrastructure "that can get into production faster and generate more value from every watt" [vendor-reported].
- Relevance for §4 (Meraki): Cisco's enterprise networking story in 2026 is bifurcating — **Meraki hardware static, Cisco AI-fabric/agentic networking accelerating** — and Cloud Control is positioned as the unified management plane spanning both [independent assessment].

