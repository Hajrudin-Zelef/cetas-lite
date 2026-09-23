---
id: etape6-phasea-vendors-dc/00-front-matter/gg-aruba-cx-10040-datasheet-depth-cx-10000-price-anchor-new
title: "GG. Aruba CX 10040 datasheet depth + CX 10000 price anchor (new)"
domain: front-matter
role: reference
task: reference
actors: ["AMD", "CoreWeave", "Nvidia", "United States"]
dates: ["2019-03", "2020-04-27", "2025-05", "2026-09-22"]
keywords: ["acquisition", "amd", "backlog", "cpo", "ethernet", "full-duplex", "latency", "license", "nvidia", "research", "revenue", "rubin"]
source: docs/RAG/etape6_phaseA_vendors_dc.md
source_anchor: ""
source_lines: [906, 966]
section: "Step 6 — Phase A: Enterprise Data-Center Switching Vendors"
sha256: 4823b86c613c854fe37e0a75a0dfe46c817f797f89b2a9643eb58fef38768a76
---

# GG. Aruba CX 10040 datasheet depth + CX 10000 price anchor (new)

### GG. Aruba CX 10040 datasheet depth + CX 10000 price anchor (new)

- **CX 10040 (bundle S4R54A)**: 32× 100GbE QSFP + 6× 400GbE QSFP-DD + 2× 10GbE SFP+; **1.8 Tbps (3.6 Tbps full-duplex)** switching, **8 Tbps bidirectional** claimed throughput; latency <1 µs (<5 µs with DPU redirect); AMD **V3C48** CPU; 120 GB M.2 SSD, 64 MB SPI flash, 2× 32 GB DDR5 SODIMM; max power **1,580 W** (idle 650 W); 100–127/200–240 VAC; stacking; AOS-CX with CLI/REST/SNMP, Fabric Composer, Aruba Central, **AMD Pensando PSM** serial management; limited lifetime warranty [official — HPE datasheet via hpe.com].
- CX 10040 launch context: unveiled ~May 2025 at HPE Houston HQ; **doubles CX 10000 distributed-services performance** (per HPE); built-in firewalling, in-line encryption, precision telemetry; HPE Morpheus VM Essentials integration for unified virtualization/ bare-metal orchestration with distributed firewall micro-segmentation [vendor-reported — Converge Digest; secondary — IoT M2M Council].
- **CX 10000-48Y6C (R8S96A) price anchor**: Connection (US VAR) lists **$42,723.95** (temporarily out of stock) — consistent with the $44–47K street range in base §3.2, tighter at the low end [secondary — connection.com].

### HH. Round-3 verification log (new open items)

1. Meraki MS450/MS355 hardware prices — reseller data only; Cisco publishes no list; gray-market MS355 listings vary 9× ($1,033–$9,350), indicating non-comparable channels [secondary].
2. Spectrum-XGS 2026 deployments — only CoreWeave confirmed; no independent ship confirmation of the firmware-delivered rollout beyond the vendor announcement [gap].
3. BlueField-4 2026 ship date — announced for early availability with Vera Rubin in 2026; no GA confirmation located [gap].
4. Quantum-X (IB) CPO H2 2025 vs Spectrum-X (Ethernet) CPO H2 2026 schedule — original vendor schedule; independent confirmations not located [gap].
5. Mellanox acquisition timeline details (NVIDIA's $6.9B acquisition, announced March 2019, closed April 27, 2020) — well-established but not re-verified against primary sources in this pass; treat as context [secondary].
6. NVIDIA $11B networking quarter (Q2 FY2027, Aug 2026) cited in round-2 §O — reported at earnings; not independently audited [vendor-reported].
7. DOCA Platform Framework versioning vs DOCA SDK versioning — this pass tracks the platform/orchestration train (v26.x), not the DOCA SDK data-plane release line; do not conflate them [gap].
8. X1600 network processor specs — roadmap reporting only; no spec sheet located [unverified].

### II. Round-3 sources (verbatim URLs)

- https://www.pcnation.com/cisco-ms450-12-hw-5rr548
- https://www.lttpartners.com/products/meraki-ms450-aggregation-switch
- https://idmproducts.com/cisco-ms450-12-hw-meraki-ms450-aggregation-switch-manageable-100-gigabit-ethernet-100gbase-x-3-layer-supported-modular-250-w-power-consumption-optical-fiber-1u-high-rack-mountable-lifetime-limited-warranty/?setCurrencyId=1
- https://signalgroup.com.au/meraki-ms450-12-hw-ms450-aggregation-switch/
- https://www.backtotheoffice.co.uk/products/cisco-meraki-ms450-12-network-switch-managed-l3-1u-grey
- https://ciscomeraki.systems/shop/cisco/meraki-ms/ms450-12-hw/
- https://networkequipment.net/products/cisco-meraki-ms355-48x2-hw-ref?_pos=20&_fid=2515010fc&_ss=c
- https://www.networktigers.com/products/ms355-48x2-hw-cisco-switch
- https://www.networkgenetics.net/cisco-meraki-ms355-48x2-hw-24x-10gb-rj-45-24x-rj-45-2x-qsfp-unclaimed-switch/
- https://www.hummingbirdnetworks.com/cisco-meraki-lic-ms355-48x2-3yr/
- https://www.superwarehouse.com/cisco-meraki-ms355-enterprise-license-and-lic-ms355-48x2-3yr.html
- https://4tekgear.com/meraki-ms355-48x2-1-year-hardware-licensing.html
- https://www.techpowerup.com/340218/nvidia-links-data-centers-into-a-unified-supercomputer-with-spectrum-xgs-ethernet
- https://www.quiverquant.com/news/NVIDIA+Introduces+Spectrum-XGS+Ethernet+to+Enable+Giga-Scale+AI+Super-Factories+Across+Distributed+Data+Centers
- https://www.crnasia.com/news-network/2025/nvidia-new-spectrum-xgs-tech-forms-ai-super-factories-by-linking-multiple-data-centers
- https://www.pipelinepub.com/news/22379
- https://network.nvidia.com/pdf/eol/LCR-000674.pdf
- https://network.nvidia.com/pdf/eol/LCR-000509.pdf
- https://network.nvidia.com/pdf/eol/LCR-000266.pdf
- https://network.nvidia.com/pdf/eol/LCR-000443.pdf
- https://github.com/asterfusion/helium_dpu/blob/HEAD/ET3600/dpdk-24.11/doc/guides/nics/mlx5.rst
- https://github.com/nvidia/doca-platform/blob/HEAD/docs/public/release-notes/v26.10.0.md
- https://github.com/nvidia/doca-platform/blob/HEAD/docs/public/release-notes/v26.4.0.md
- https://github.com/nvidia/doca-platform/blob/HEAD/docs/public/release-notes/v25.7.0.md
- https://github.com/nvidia/doca-platform/blob/HEAD/docs/public/release-notes/v25.10.1.md
- https://docs.nvidia.com/networking/display/nvidia-bluefield-3-dpu-nic-firmware-release-notes-v32-43-2408.2408.pdf
- https://www.servethehome.com/nvidia-bluefield-4-with-64-arm-cores-and-800g-networking-announced-for-2026/
- https://www.sdxcentral.com/analysis/nvidias-bluefield-4-a-first-look-at-the-dpu-built-to-run-ai-factories/
- https://convergedigest.com/nvidia-charts-ai-data-center-roadmap-through-2028/
- https://www.techpowerup.com/334337/nvidia-commercializes-silicon-photonics-with-infiniband-and-ethernet-switches
- https://www.pulse2.com/dell-technologies-q2-revenue-hits-record-47-billion-as-ai-server-backlog-reaches-95-billion-and-fy27-outlook-rises-to-192-billion/
- https://sharesify.com/dell-q2-fy2027-ai-demand-surges-but-expectations-are-now-sky-high/
- https://www.hpe.com/psnow/generateDDS/HPE Aruba Networking CX 10040 32p QSFP28 100G 6p QSFP-DD 400G Front-to-Back 4xFan 2xPSU AC Bundle Digital Data Sheet-PSN1014899894CAEN.pdf?oid=1014899894&cc=CA&lc=EN&softroll=0&prelaunch=false&print=§ion=&prelaunchSection=&softrollSection=&deepLink=&utm_source=&utm_campaign=&utm_content=&utm_term=&isLinearized=false&contentDisposition=attachment
- https://www.connection.com/product/hpe-aruba-cx-10000-48y6c-switch/r8s96a/41376956
- https://convergedigest.com/hpe-doubles-distributed-switch-performance-with-aruba-cx-10040-amd-dpus/
- https://www.iotm2mcouncil.org/iot-library/news/iot-newsdesk/hpe-expands-aruba-wired-and-wireless-portfolio/

**Round-3 collection metadata:** read-only web research (browser_search, 2026-09-22); no live-browser visits; nothing sent externally. No identifiers guessed. All new facts carry provenance tags; new open items are listed in §HH above. Existing sections §§1–8, §§A–M, §§N–X, and the "round 2" §§N–S were not modified.


---

