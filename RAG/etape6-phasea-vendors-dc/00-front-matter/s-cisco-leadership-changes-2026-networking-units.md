---
id: etape6-phasea-vendors-dc/00-front-matter/s-cisco-leadership-changes-2026-networking-units
title: "S. Cisco leadership changes, 2026 (networking units)"
domain: front-matter
role: reference
task: reference
actors: ["Broadcom", "China", "CoreWeave", "Microsoft", "Nebius", "Nvidia", "SpaceX", "United States"]
dates: ["2026-05", "2026-09", "2026-09-08"]
keywords: ["acquisition", "cpo", "ethernet", "full-duplex", "gpu", "latency", "memory", "nvidia", "optics", "pricing", "revenue", "ualink"]
source: docs/RAG/etape6_phaseA_vendors_dc.md
source_anchor: ""
source_lines: [542, 591]
section: "Step 6 — Phase A: Enterprise Data-Center Switching Vendors"
sha256: 203834f64837a31648e3c0574c9a617a9234a83901dc838230aad6d4333fcf93
---

# S. Cisco leadership changes, 2026 (networking units)

### S. Cisco leadership changes, 2026 (networking units)

- **Martin Lund exits**: Common Hardware Group (CHG) head Martin Lund left Cisco (~spring 2026) to lead an unnamed "early-stage AI startup" (later described as becoming CEO of an AI startup); Cisco consolidated the **CHG (silicon, systems, optics) under chief product officer Jeetu Patel's Product Organization** [secondary — SDxCentral; Fierce Network, ~Apr–May 2026].
- Lund had joined Cisco in early 2024 from Microsoft (Azure for Operators, ex-Metaswitch CEO, 12 years leading Broadcom's networking business); his team was credited with establishing Cisco's AI work across **silicon (Silicon One), systems, and optics** [secondary — SDxCentral].
- **Jeetu Patel** promoted to President & CPO at the start of Cisco FY2026, consolidating **Networking + Security + Collaboration** under one product organization [secondary — bizkonnect.com, 2026-08].
- Other FY2026-start leadership moves: **Mark Patterson** EVP→CFO (succeeding Scott Herren); **Oliver Tuszik** EVP Global Sales (effective Apr 27, 2026); **Steve Clayton** CCO (from Microsoft, effective Jan 5, 2026); **Eyal Dagan** to EVP Strategic Projects; Chuck Robbins remains CEO [secondary — bizkonnect.com].
- Patel's stated posture (interview, Sep 2026): "The core networking business is actually central to the success… I anticipate it growing at a very aggressive pace" — networking as the framework for security, AI GPU interconnect, and observability [secondary — SDxCentral].

### T. Dell'Oro campus-switch Q2 2026 — new detail (complements DC data)

Pass #1 (§E/G) covered Dell'Oro AI back-end. The **campus** report adds (2026-09-08) [independent — Dell'Oro Group press release]:

- Component shortages hit campus-switch availability and prices in 2Q 2026; **campus-switch ASP grew 19%**, driving revenue growth [independent].
- "Every vendor outside China grew campus switch ASP on a year-over-year basis… a direct consequence of AI-fueled component shortages driving up component costs and resulting in list price increases of nearly all campus switches" — Siân Morgan, Senior Director [independent].
- **Cisco, Arista, Ubiquiti and Extreme all gained share YoY** in 2Q 2026 (Cisco + Arista for the second consecutive quarter) [independent].
- Public **cloud-managed switch revenue** bounced back and outgrew the total market — "flywheel effects of a recurring revenue model" (relevant to Meraki/Aruba/Central positioning) [independent].
- Memory shortages expected to shape the market into 2027, with rising ASPs and availability issues [independent].

### U. HPE (Aruba + Juniper) competitive framing — enterprise AI (new angle)

A September 2026 trade analysis compares enterprise AI-network strategies (summarized; secondary framing) [secondary — thecodew.com, 2026-09]:

- **Core AI silicon**: Cisco = Silicon One G300 (102.4T) + support for NVIDIA Spectrum-X silicon in the same chassis; **HPE = Broadcom Tomahawk 6 (first OEM deal, Dec 2025) + QFX support for open UALink**; Arista = TH6 Etherlink up to 1.6T (7060XE7).
- **Fabric management**: Cisco = Nexus One (single plane across ACI, Hyperfabric, SONiC); HPE = merging Aruba Central + Juniper Mist under unified AIOps ("build once, deploy twice"); Arista = single EOS + CloudVision.
- **UEC roles**: Cisco = member; **HPE = founding member (Juniper pre-acquisition was not a member)**; Arista = founding member, most vocal Ethernet-over-IB advocate.
- **Estimated Q1 2026 DC Ethernet shares**: Cisco ~14%, HPE (Aruba+Juniper) ~7% (Juniper legacy figure; combined share not yet separately disclosed), Arista ~19% — note: differs from IDC Q2 figures; treat as analyst estimates [secondary].
- **Arista named Gartner Magic Quadrant Leader (2026)** for enterprise wired/wireless LAN; Q2 2026 revenue $3.036B (+37.7% YoY), FY2026 guided above $10B (analyst summary; directionally consistent with Arista earnings) [secondary].

### V. Aruba CX 9300 / CX 10000 — updated spec and pricing detail

- **CX 9300-32D (S1D07A#ARM, SHI listing)**: 32× 100/200/400G QSFP-DD + 2× 1/10G SFP+; **MSRP $119,995.00, street $102,382.00** (updated Sep 2026) [secondary — shi.com]. HPE Store SKUs: **S1D07A** (front-to-back) / **S1D08A** (back-to-front) 2AC-PSU bundles [official — buy.hpe.com].
- **CX 9300S** variant: 32p QSFP28 100G + 8p QSFP-DD 400G bundles, new transceivers and PSU options (HPE "What's New" block) [official].
- Pod-scale claim: CX 9300-32D supports DC PODs of up to **16,834× 25GbE servers or 8,192× 100GbE servers** (with 4× 100G SN transceivers) [official].
- **CX 9300X-32D (R9A29A/R9A30A)** reseller spec sheet: 25.6 Tbps, **9,520 Mpps**, 2.4 GHz x86 CPU, 16 GB DDR4 + 120 GB SSD, **64 MB dynamic deep-buffer packet buffer**, 256K MAC entries; EVPN-VXLAN scalable multi-site overlays; 32× 400G QSFP-DD (breakout 4× 100G / 8× 50G) [secondary — grabnpay.in].
- **CX 10000-48Y6C datasheet detail** (R8P13A = front-to-back, R8P14A = back-to-front bundles): **1.8 Tbps routing / 3.6 Tbps full-duplex switching**, **2× Pensando Elba 3 GHz DPUs**, 32 MB packet buffer, **<1.3 µs latency**, 1,940 Mpps, 550W, 9.75 kg; routing: static IPv4/IPv6, RIPv2/RIPng, OSPF/OSPFv3, BGP-4, MP-BGP w/ IPv6, PBR, ECMP, GRE; management: CLI/REST/SNMP, Fabric Composer, NetEdit [official-via-reseller — HPE datasheet via eetgroup.com].
- **CX 10000-48Y6C (R8P14A#ABA)**: regular/reseller price **$62,618.16** ("call for availability") — higher than the $44–47K street cited in base §3.2, reflecting different bundles/resellers; flagged [secondary — lttpartners.com].

### W. New open items (Supplementary Pass #2 verification log)

1. Martin Lund departure exact date (spring 2026 window; no official Cisco announcement text located) [gap].
2. Dell Enterprise SONiC dollar prices — still unpublished; structure known, amounts not [gap].
3. Meraki MS150 UK reseller prices vs US MSRP — 24T/48T price inversion is a listing artifact, not Cisco pricing [secondary].
4. IDC Q2 2026 ($2.5B / 20.4% DC) vs Next Platform ($3.86B / 2.8×) NVIDIA figures — Pass #1 §L.3 unresolved; both cited [unverified].
5. thecodew.com Q1 2026 share table (Cisco ~14%, HPE ~7%, Arista ~19%) vs IDC — analyst estimates, unreconciled [secondary].
6. CX 10000 street prices — $44–47K (base) vs $62.6K (this pass) — different bundles/resellers [secondary].
7. Dell OS10 10.6.1.x release notes for 1.6T platforms (Z9964): OS10 support on Z9964 not verified in the 10.6.0.9 supported-hardware list [gap].
8. Spectrum-6 end-user ship dates at CoreWeave/Microsoft/Nebius/SpaceXAI/Tesla — named as first adopters; ship dates not confirmed [unverified].

### X. Supplementary Pass #2 sources (verbatim URLs)

