---
id: etape6-phasea-vendors-dc/01-round-5-complementary-research-pass-2026-09-22-deep-datashee/r4-b-cumulus-linux-eol-table-refresh-september-2026-5-19-0-e
title: "R4-B. Cumulus Linux — EOL table refresh (September 2026) + 5.19.0 evidence"
domain: round-5-complementary-research-pass-2026-09-22-deep-datashee
role: deep-dive
task: reference
actors: ["Broadcom", "Nvidia", "United States"]
dates: ["2025-12", "2026-02", "2026-03-29", "2026-05", "2026-07", "2026-09", "2026-11", "2027-04", "2027-07", "2027-08", "2027-11"]
keywords: ["ethernet", "nvidia", "pricing"]
source: docs/RAG/etape6_phaseA_vendors_dc.md
source_anchor: ""
source_lines: [1946, 1992]
section: "Round-5 complementary research pass (2026-09-22) — deep datasheet and radix detail"
sha256: 35a8dd64a8972f9e50709180db1dfdcc5b203bdef3b1fb0871262e326f31858d
---

# R4-B. Cumulus Linux — EOL table refresh (September 2026) + 5.19.0 evidence

### R4-B. Cumulus Linux — EOL table refresh (September 2026) + 5.19.0 evidence

- NVIDIA's official "Cumulus Linux Release Versioning and Support Policy" KB (crawl ~Sep 2026) confirms and extends the EOL picture from this file's §F [official]:
  - 5.15.z → EOL **November 2026** (imminent); 5.16.z → April 2027; 5.17.z → July 2027; 5.18.z → August 2027; 5.12.z → February 2026; 5.13.z → May 2026; 5.14.z → July 2026 (all now EOL); LTS: 5.9.z → April 2027, 5.11.z → November 2027; **4.3.1-and-later on Broadcom → EOL December 2025 (already passed)**; 5.y.z supports Spectrum-based switches only.
  Source: https://docs.nvidia.com/networking-ethernet-software/knowledge-base/Support/Support-Offerings/Cumulus-Linux-Release-Versioning-and-Support-Policy/
- **5.19.0 exists**: the Cumulus 5.13 release-notes "fixed in" column references **"Fixed in 5.19.0"** (issue 4508836 — Cumulus allows mismatched-speed ports in one LACP bond without error) [official — docs.nvidia.com]. No dedicated 5.19.0 "What's New" page was located via search; 5.19.0 is therefore confirmed by fix-table reference only — **flag as thin confirmation**.
  Source: https://docs.nvidia.com/networking-ethernet-software/cumulus-linux-513/Whats-New/rn/
- **NVIDIA security bulletin 5817** (GitHub, ~Sep 2026): CVE-2026-24184 (Cumulus Linux GA, affected 0.0–5.16 → updated version 5.16), CVE-2026-24183 (5.15 → 5.16); LTS 5.11 → 5.11.5, LTS 5.9 → 5.9.5. Same bulletin: CVE-2026-24185 affects NVOS on GB300 and IBSwitch XDR (patched 25.0.2.4438 / 25.0.2.6077) [official].
  Source: https://github.com/nvidia/product-security/blob/HEAD/2026/5817/5817.md
- Community note: ipSpace.net (2026-03-29 revision) reports newer **Cumulus VX images are accessible to NVIDIA customers with support contracts** (previously unobtainable), while criticizing the shift of demos into the NVIDIA Air cloud platform [secondary — ipSpace opinion].
  Source: https://blog.ipspace.net/2025/06/cumulus-linux-gone/

### R4-C. Meraki MS210 — full model and pricing detail (new)

The base §4.1 named the MS210 family (5 models, 24/48× 1G, PoE variants 370W) without prices. September-2026 street/MSRP data [independent — retailer listings]:

- **MS210-48FP-HW** (48× 1GbE PoE+, 740W budget, 4× SFP): SHI MSRP **$9,515.87**, SHI street **$4,490.00**; publicsector SHI MSRP **$10,724.87**, street **$7,888.00** — **flag: ~$1.2K MSRP variance between the two SHI storefronts**; gray-market refurbished as low as **$368.99** (cablesandkits, CK Certified) [independent].
  Sources: https://www.shi.com/product/34444595/CISCO-MERAKI-MS210- ; https://www.publicsector.shidirect.com/Product/34444595/CISCO-MERAKI-MS210-48FP-1G-L2CLD-MNGD-48X-GIGE-740W-POE-SWITCH ; https://www.cablesandkits.com/networking/switches/meraki-switches/ms210-48fp-hw/pro-27668/
- **MS210-48-HW** (48× 1GbE, non-PoE): MSRP **$6,419.03**, street **$4,721.00** [independent].
  Source: https://www.shi.com/product/34444596/CISCO-MERAKI-MS210-
- **MS210-48LP-HW** (48× 1GbE PoE+, 370W): MSRP **$8,126.44**, street **$5,977.00** [independent].
  Source: https://www.shi.com/product/34444597/CISCO-MERAKI-MS210-48LP-1G-L2CLD-MNGD-48X-GIGE-370W-POE-SWITCH
- **MS210-24-HW** (24× 1GbE + 4× SFP + 2× stacking, 128 Gbps switching, 80 Gbps stacking, fanless, 15/24W power): new **$1,938.99** (networkequipment.net) [independent].
  Source: https://networkequipment.net/products/cisco-meraki-ms210-24-hw-new
- MS210-48FP spec anchors from reseller specs: 176 Gbps switching capacity, 77.38 mpps, 32K MAC, 4094 VLANs, 9578-byte jumbo, virtual + physical stacking, PoE+ 30W/port [secondary].

### R4-D. ConnectX-8 SuperNIC — street pricing and SKU detail (new)

Base §2.4 covered ConnectX-8 SuperNIC features without street prices. September-2026 retailer data [independent]:

| SKU | Form factor | US price | UK price (ex VAT) |
|---|---|---|---|
| ConnectX-8 800G OSFP 1-port (VPI) | PCIe 6.0 x16 | **$2,519.00** (FS.com) | **£1,991.00** (FS.com UK) |
| ConnectX-8 400G QSFP112 2-port (VPI) | PCIe 6.0 x16 | **$2,519.00** (FS.com) | **£2,042.00** (FS.com UK) |
| C8240 (900-9X81Q-00CN-ST0), dual QSFP112 SocketDirect NDR | PCIe 6.0 x16 | **$1,779.00** (naddod, 110 in stock) | — |
| C8180 (900-9X81E-00EX-DT0), single OSFP XDR 800G | PCIe 6.0 x16 | **$1,779.00** (naddod, 98 in stock) | — |
| C8180 (900-9X81E-00EX-ST0), SocketDirect single OSFP | PCIe 6.0 x16 | **$1,779.00** (naddod, 50 in stock) | — |
| C8180X PCIe auxiliary kit (930-9XAX6-0025-000) | — | **$239.00** (naddod) | — |
| C8240 list (accio) | — | **$2,276.30** | — |
| ConnectX-7 400G OSFP 1-port | PCIe 5.0 x16 | $2,188.00 (FS.com) | £1,856.00 (FS.com UK) |
| ConnectX-7 200G QSFP112 2-port | PCIe 5.0 x16 | $2,129.00 (FS.com) | £1,856.00 (FS.com UK) |

Sources: https://www.FS.com/c/nvidia-ethernet-nics-4014 ; https://www.fs.com/uk/c/nvidia-ethernet-nics-4014 ; https://www.naddod.com/collections/nvidia-networking/infiniband-adapters ; https://www.accio.com/plp/connectx-8-c8240

- Datasheet detail (NVIDIA, via nvdam.widen.net): ConnectX-8 SuperNIC — max 800 Gb/s IB and Ethernet; IB rates 800/400/200/100; Ethernet 400/200/100 (plus 100/50/25); PCIe Gen6 up to 48 lanes; **Multi-Host for up to 4 hosts**; advanced dynamic routing + telemetry-based congestion control; extended SHARP; 16M I/O channels; IBTA v1.7; HHHL single-port OSFP, HHHL dual-port QSFP112, dual ConnectX-8 mezzanine [official].
  Source: https://nvdam.widen.net/content/8h0owe2dhm/original/connectx-datasheet-connectx-8-supernic-update-a4-web-zhCN-3523588-R3.pdf?u=rubmrs&use=c8xan&download=true

