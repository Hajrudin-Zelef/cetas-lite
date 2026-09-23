---
id: etape6-phasea-vendors-dc/00-front-matter/i-aruba-cx-street-pricing-partially-closes-base-7-items-8-10
title: "I. Aruba CX street pricing (partially closes base §7 items 8/10)"
domain: front-matter
role: reference
task: pricing
actors: ["Cohere", "Nvidia"]
dates: ["2026-05-04", "2026-05-05", "2026-06-09", "2026-09-22"]
keywords: ["pricing", "acquisition", "agent", "attribution", "cpo", "datacenter", "distribution", "ethernet", "license", "licenses", "nvidia", "research"]
source: docs/RAG/etape6_phaseA_vendors_dc.md
source_anchor: ""
source_lines: [430, 493]
section: "Step 6 — Phase A: Enterprise Data-Center Switching Vendors"
sha256: 203235464ca4f71189e88dabc67153ca6f741b58536b84fcd47bcaa192ef5766
---

# I. Aruba CX street pricing (partially closes base §7 items 8/10)

### I. Aruba CX street pricing (partially closes base §7 items 8/10)

- **CX 9300-32D (SKU R9A29A#ABA)**: 25.6 Tbps, 5 Bpps, 32× 100/200/400GbE; reseller street **~$120,600–$124,314** ("call for availability") [secondary — lttpartners.com]. Positioning: 100G leaf or 100/400G spine; pairs with CX 8325/8360/10000 leaves [secondary].
- **CX 8360-48XT4C v2 (JL707C#ABA)**: MSRP **$44,064**, street **$37,367** (SHI, updated Sep 2026); 4.8 Tbps, 1786 Mpps, 48× 1/10/25G + 4× 40/100G, MACsec on low-density + uplink ports, VSX, BGP/OSPF/VRF/EVPN/VXLAN/IPv6 [secondary — shi.com; official specs via HPE datasheet].
- No 2026 hardware refresh of CX 9300/8360 families was located; both continue as current [unverified — continuation assumed].

### J. Cisco M&A status updates (networking-adjacent)

- **Astrix Security: acquisition COMPLETED** — Cisco announced intent May 4, 2026; completion reported May 5, 2026 at ~**$400M** (Calcalist; Silicon Angle cited ~$300M — conflicting press figures, flagged). Non-human-identity / AI-agent security → Cisco Identity Intelligence, Duo, Secure Access [secondary — startuprise.org, scworld.com, dataconomy.com].
- **Galileo Technologies** (AI observability): Cisco announced intent to acquire (~Apr 2026, Burlingame CA); expected to close in **Q4 of Cisco FY2026**; strengthens Splunk Observability AI-agent monitoring [secondary — obstracts.com briefing, Apr 2026].
- Neither deal is switching-hardware; both feed Cisco's AI-security/observability narrative around the network.

### K. Quantum-X800 availability evidence + CPO context (partially closes base §7 item 5)

- HPE lists the NVIDIA InfiniBand XDR switch as **orderable: SKU P79109-B21** (HPE Store Ireland) — "2×36-port OSFP managed" Quantum-X800 for HPE; HPE datasheet confirms **115.2 Tb/s non-blocking on Q3400-RA** and lists Quantum-X Photonics CPO in features [official — buy.hpe.com; HPE datasheet PSN1014891575].
- CPO ramp context: NVIDIA CTO Gilad Shainer publicly **disputed the June 9, 2026 SemiAnalysis CPO yield-risk thesis**, stating CPO is **already shipping and ramping H2 2026** [secondary — investment research memo, Jul 2026]; Yole Group forecasts **large-scale CPO deployment in 2028–2030** [secondary]. Treat the 2026 ramp claim as vendor-disputed, not independently verified.
- NVIDIA's photonics supplier commitments: **>$6.5B deployed in 2026 YTD, incl. $2B each into Coherent and Lumentum** with multiyear purchase commitments [secondary — single investment memo; unverified].

### L. Supplementary verification log (new open items)

1. SN6800-LD / SN6810-LD / SN6600-LD GA dates and pricing — spec sheet published; commercial availability unconfirmed [gap].
2. Dell Enterprise SONiC subscription dollar prices — model structure confirmed; amounts unpublished [gap].
3. Next Platform ($3.86B) vs IDC press release ($2.5B) NVIDIA Q2 2026 figure — unresolved; likely ODM-attribution difference [unverified].
4. ODM $3.56B Q2 2026 — Next Platform estimate, not IDC data [secondary-estimate].
5. CX 9300/8360 2026 refresh — none found; continuation assumed [unverified].
6. Meraki license street prices — reseller/secondary data; Cisco does not publish list consistently [secondary].
7. Astrix final deal value — $300M vs $400M conflicting press reports [unverified].
8. Quantum-X Photonics ship confirmation — vendor claims ramping H2 2026; no independent ship confirmation located [gap].
9. NVIDIA $6.5B photonics commitments — single investment memo, unconfirmed [unverified].
10. Cisco acquisition completions beyond Astrix (Galileo close) — intent announced; close unconfirmed [gap].

### M. Supplementary sources (verbatim URLs)

- https://www.datamation.com/networks/dell-technologies-enterprise-sonic-distribution-review/
- https://www.trustradius.com/products/dell-enterprise-sonic-distribution/pricing
- https://test.direktronik.se/globalassets/_product-images--pdf/1.-natverk/1.1-lan/02.-switchar-66/sonic_by_broadcom_datasheet_4.5.0.pdf
- https://www.dell.com/support/kbdoc/en-ph/000228560/minimum-recommended-and-latest-code-versions-for-networking-products
- https://www.delltechnologies.com/asset/de-at/products/networking/technical-support/dell-powerswitch-sn6000-series-spec-sheet.pdf
- https://cdn.blueally.com/netsolutionworks/datasheets/powerswitch-sn6800-ld-spec-sheet.pdf
- https://cdn.blueally.com/netsolutionworks/datasheets/powerswitch-sn5000d-spec-sheet.pdf
- https://delltechnologies.com/asset/nl-nl/products/networking/technical-support/dell-networking-spec-sheet-sonic.pdf
- https://un5gmtkzgk7ja497aqmdywr2adtg.goodfellasdiner.com/asset/sv-se/products/networking/technical-support/nvidia-spectrum-sn5600-datasheet.pdf
- https://www.highendcomputing.co.uk/datasheets/powerswitch-sn5610.pdf
- https://stordis.com/enterprise-sonic-4-6-released/
- https://docs.nvidia.com/networking-ethernet-software/knowledge-base/Support/Support-Offerings/Cumulus-Linux-Release-Versioning-and-Support-Policy/
- https://github.com/cumulusnetworks/docs/blob/HEAD/content/cumulus-linux-517/_index.md
- https://www.idc.com/resource-center/blog/ethernet-switch-market-surges-43-4-to-18-9b-in-2q26-as-ai-infrastructure-demand-drives-record-datacenter-spending/
- https://www.nextplatform.com/connect/2026/09/20/the-genai-boom-accelerates-and-transforms-ethernet-switching/5297603
- https://www.thefastmode.com/technology-and-solution-trends/50478-delloro-group-ai-back-end-network-switch-sales-surpass-front-end-networks
- https://documentation.meraki.com/Platform_Management/Product_Information/Meraki_Licensing/General_Licensing_Information/Meraki_Per-Device_Licensing_-_Configuration
- https://www.cloudwifiworks.com/ms350-24x-licenses-renewals.asp
- https://www.linktly.com/security-software/cisco-meraki-mx-security-appliances-review/
- https://medium.com/@cacheguard/your-cisco-meraki-appliance-stops-working-when-the-license-expires-thats-not-a-bug-0882eb7ae23c
- https://www.shi.com/product/44631762/HPE-Aruba-CX-8360-48XT4C-v2
- https://www.lttpartners.com/products/hpe-cx-9300-32d-ethernet-switch-1
- https://startuprise.org/cisco-acquires-ai-security-startup-astrix-in-400m-deal/
- https://www.scworld.com/brief/cisco-acquires-astrix-security-to-bolster-ai-agent-defenses
- https://buy.hpe.com/ie/en/networking/switches/enterprise-switches/nvidia-infiniband-xdr-2x36%e2%80%91port-osfp-managed-connector-to-power-airflow-switch/p/p79109-b21
- https://github.com/47tzp4ydc9-cmyk/life-os/blob/HEAD/investment-os/narrative/research/silicon-photonics.md

**Supplementary collection metadata:** read-only web research (browser_search, 2026-09-22); no live-browser visits; nothing sent externally. No identifiers guessed. All new facts carry provenance tags; new open items are listed in §L above. Existing sections §§1–8 were not modified.

---

