---
id: etape6-phasea-vendors-dc/00-front-matter/s-cisco-leadership-changes-2026-networking-units
title: "S. Cisco leadership changes, 2026 (networking units)"
domain: front-matter
role: reference
task: reference
actors: ["Broadcom", "China", "CoreWeave", "Microsoft", "Nebius", "Nvidia", "SpaceX", "United States"]
dates: ["2026-05", "2026-09", "2026-09-08", "2026-09-22"]
keywords: ["acquisition", "compute", "cpo", "datacenter", "distribution", "ethernet", "full-duplex", "gpu", "gpus", "hyperscaler", "latency", "memory"]
source: docs/RAG/etape6_phaseA_vendors_dc.md
source_anchor: ""
source_lines: [542, 628]
section: "Step 6 — Phase A: Enterprise Data-Center Switching Vendors"
sha256: 42a7f01f27786cd5d1138e62bff183fb0f8dfaae12811e880fecc07950a8c49b
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

- https://monitor.spacescience.ro/dellfw/archives/firmware/10.6.0.9/SmartFabric_OS10_10.6.0.9_Release_Notes.pdf
- https://www.dell.com/support/manuals/en-in/smartfabric-os10-emp-partner/ee-upgrade-downgrade/revision-history?guid=guid-b7d10060-f564-41ea-88bc-db1b40b84fd1&lang=en-us
- https://dl.dell.com/content/manual44947528-dell-smartfabric-os10-and-smartfabric-services-security-configuration-guide-release-10-6-1-june-2026.pdf?language=en-us
- https://www.dell.com/support/manuals/en-lv/smartfabric-os10-emp-partner/smartfabric-os10-rn-10.5.6.7-pub/Dell-SmartFabric-OS10-Release-Notes?guid=guid-8951a213-856f-406a-abb1-9b6e72f539c0&lang=en-us
- https://www.4rfv.co.uk/industrynews/332702/nvidia_spectrum_6_debuts_for_gigascale_ai_factories_built_on_vera_rubin
- https://cxotoday.com/hardware/nvidia-debuts-spectrum-6-switches-to-power-ai-networks-of-the-future/
- https://www.networkworld.com/article/4200086/nvidia-unveils-spectrum-x-networking-platform-designed-to-connect-millions-of-gpus.html
- https://www.techrepublic.com/article/news-nvidia-gtc-2026-live-updates/
- https://docs.nvidia.com/networking/display/nvidia-connectx-9-supernic-firmware-release-notes-v82-48-1000-february-2026-ga-release.pdf
- https://www.sdxcentral.com/news/nvidia-reveals-next-gen-dpu-to-help-offload-gigascale-ai-infrastructure/
- https://www.servethehome.com/nvidia-launches-next-generation-rubin-ai-compute-platform-at-ces-2026/
- https://www.neoteo.com/en/coreweaves-reported-nvidia-vera-rubin-deployment-reaches-504-gpus
- https://www.trustradius.com/products/dell-enterprise-sonic-distribution/pricing
- https://delltechnologies.com/asset/sv-se/products/networking/technical-support/dell-networking-spec-sheet-sonic.pdf
- https://www.shi.com/product/49737007/Cisco-Meraki-MS150-48MP-4X
- https://networkwarehouse.co.uk/products/meraki-ms150-48mp-4x
- https://networkwarehouse.co.uk/products/meraki-ms150-24p-4g
- https://www.pc-canada.com/item/meraki-ms150-48mp-4x-ethernet-switch/ms150-48mp-4x
- https://www.sdxcentral.com/news/cisco-hardware-group-head-heads-for-the-door/
- https://www.fierce-network.com/cloud/cisco-unites-hardware-and-software-under-jeetu-patel-martin-lund-exits
- https://www.bizkonnect.com/blogs/cisco-s-org-chart-keeps-changing-here-s-how-sales-teams-keep-up
- https://www.sdxcentral.com/analysis/cisco-touts-networking-as-central-to-success-key-for-platform-growth/
- https://www.wcia.com/business/press-releases/cision/20260908SF41495/market-prices-skyrocket-driving-campus-switch-market-to-growth-in-2q-2026-according-to-delloro-group/
- https://www.thefastmode.com/technology-and-solution-trends/50478-delloro-group-ai-back-end-network-switch-sales-surpass-front-end-networks
- https://www.thecodew.com/2026/08/beyond-the-hyperscaler-enterprise-ai-ethernet-fabrics.html
- https://www.idc.com/resource-center/blog/ethernet-switch-market-surges-43-4-to-18-9b-in-2q26-as-ai-infrastructure-demand-drives-record-datacenter-spending/
- https://www.shi.com/Product/47703495/HPE-Aruba-CX-9300-32D-32xQSFP-DD-400G-2xSFP-10G-Switch
- https://buy.hpe.com/us/en/networking/switches/fixed-port-l3-managed-ethernet-switches/aruba-9300-switch-products/hpe-aruba-networking-cx-9300-switch-series/p/s1d08a
- https://www.grabnpay.in/products/aruba-cx-9300x-32d-r9a29a-r9a30a-32-400-gbe-qsfp-dd-high-density-core-aggregation-switch
- https://product-images.eetgroup.com/documents/Doc_34786093.pdf
- https://www.lttpartners.com/products/aruba-cx-10000-48y6c-distributed-services-back-to-front-6-fans-2-psu-switch-bdl

**Supplementary Pass #2 collection metadata:** read-only web research (browser_search, 2026-09-22); no live-browser visits; nothing sent externally. No identifiers guessed. All new facts carry provenance tags; new open items are listed in §W above. Existing sections were not modified.


---

