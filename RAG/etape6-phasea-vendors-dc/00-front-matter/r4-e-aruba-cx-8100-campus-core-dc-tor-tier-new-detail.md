---
id: etape6-phasea-vendors-dc/00-front-matter/r4-e-aruba-cx-8100-campus-core-dc-tor-tier-new-detail
title: "R4-E. Aruba CX 8100 — campus core / DC ToR tier (new detail)"
domain: front-matter
role: reference
task: reference
actors: ["AMD", "Broadcom", "Nvidia", "xAI"]
dates: ["2025-02", "2025-12", "2026-09-02", "2026-09-15", "2026-09-22"]
keywords: ["acquisition", "amd", "ethernet", "gpu", "gpus", "helios", "inference", "license", "licenses", "mi455x", "nvidia", "pricing"]
source: docs/RAG/etape6_phaseA_vendors_dc.md
source_anchor: ""
source_lines: [1722, 1790]
section: "Step 6 — Phase A: Enterprise Data-Center Switching Vendors"
sha256: ba92e557ecc46a864b4fd33a72cfcc0dfeb16753d9140d2c8a39faa832c26525
---

# R4-E. Aruba CX 8100 — campus core / DC ToR tier (new detail)

### R4-E. Aruba CX 8100 — campus core / DC ToR tier (new detail)

Not covered in any earlier pass (which focused on CX 9300/8360/8325/10000/10040):

- Positioning: **1.76 Tbps / 1,309 Mpps**, campus core/aggregation and **DC ToR/EoR**; line-rate 1/10G (SFP/SFP+) + 40/100G (QSFP+/QSFP28) in 1U; 4× 10G and 4× 25G breakout from 40/100G ports [official — HPE CX 8100 datasheet via hpe.com].
- SKUs: **R9W86A** (24× 1G/10G SFP/SFP+ + 4× 40/100G QSFP+/QSFP28, front-to-back); **R9W88A** (24× Smart Rate 100M/1/2.5/5/10G Base-T + 4× 10G SFP+ + 4× 40/100G QSFP28, Thai retail **฿183,500 excl. 7% VAT** ≈ regional snapshot) [secondary — avendor.com, itk.co.th].
- Platform: quad-core Arm Cortex-A72 @ 1.8 GHz, 16 GB RAM, 32 GB flash, **32 MB packet buffer**; max power 500W (idle 120W); dual hot-swap PSUs; VSX redundancy; limited lifetime warranty [official — datasheet].
- Software: **ArubaOS-CX** with CX Foundation license embedded (NAE, dynamic segmentation, HA, QoS, L2/L3, multicast, security, NetEdit support), Advanced tier optional; **Fabric Composer** supports the CX 8100 for software-defined leaf-spine orchestration [official — HPE datasheet].
- Price anchor: CX 8100 24XT4XF4C (Smart Rate) — **regular $26,785.44**, "call for availability" [secondary — lttpartners.com]. No Cisco-published list; treat as single-reseller anchor.

### R4-F. xAI–Dell $5B+ deal — status check, final attempt (remains unresolved)

Base §7.3 and §1.6 left this open; a targeted 2026 search found no confirmation:

- Original report: Bloomberg, **February 2025** — Dell "close to finalizing" a **$5B+ agreement** for GB200-based AI servers to xAI, "details are being finalized and still may change"; Dell and NVIDIA declined comment [secondary — Bloomberg via nation.lk, mitrade.com, timesofai.com, brandiconimage.com].
- Context: xAI's Memphis "Colossus" already ran on a mix of Dell + Supermicro servers; Dell said December 2025 it had deployed tens of thousands of GPUs at xAI's Memphis facility [secondary — Bloomberg relay].
- **No 2026 finalization or delivery confirmation located** as of 2026-09-22; status stays **[unverified]**.
- Not a deal reference: Dell's **$5.0B senior-notes offering priced September 15, 2026** (1.25B× 2029/2031, 1.5B× 2033, 1.0B× 2037) is corporate debt, unrelated to xAI [secondary — minichart.com.sg]. Do not conflate the two " $5B" items.

### R4-G. NVIDIA Mellanox heritage branding note (clarification)

- The "Mellanox" name persists in NVIDIA URLs/artifacts (network.nvidia.com, MLNX_OFED, MLNX-OS-era EOL notices) but current products carry **NVOS/NVidia branding**; Quantum-X800 systems run **NVOS**, not MLNX-OS [official — datasheet].
- Driver continuity: the mlx5 driver (DPDK 24.11) still supports the full lineage ConnectX-4 → ConnectX-7 and BlueField → BlueField-3 [secondary — pass §CC].

### R4 verification log (new open items)

1. Helios ship dates and Juniper scale-up switch SKU/pricing — announced for 2026; no SKU or pricing located [gap].
2. DOCA SDK 3.5.0 full changelog — release-notes page updated 2026-09-02 but API-change table listed "TBD" at crawl time [official, incomplete].
3. Q3200-RA / Q3400-RA street pricing — quote-only, not located [gap].
4. CX 8100 street prices — single-reseller anchors only (lttpartners, Thai retailer); no HPE list [secondary].
5. xAI–Dell $5B+ finalization — no 2026 confirmation located after targeted search [gap — final attempt, closing as unresolved].

### R4 sources (verbatim URLs)

- https://www.tomshardware.com/tech-industry/semiconductors/hpe-adopts-amd-helios-rack-architecture-for-2026-ai-systems
- https://www.theregister.com/special-features/2025/12/02/hpe-to-ship-rack-scale-ai-system-using-amds-helios-in-2026/2254787
- https://www.techradar.com/pro/the-ai-race-explodes-as-hpe-deploys-amds-helios-racks-crushing-limits-with-venice-cpus-and-insane-gpu-density
- https://ai-daily.news/articles/hpe-to-offer-amd-helios-ai-racks-globally-by-2026
- https://mlq.ai/news/amd-helios-mi455x-rack-scale-platform-surfaces-with-72-gpu-design-ualink-over-ethernet-interconnect/
- https://vmvirtualmachine.com/hpe-and-amd-announce-helios-ai-rack-with-ethernet-network/
- https://docs.nvidia.com/doca/sdk/changes+and+new+features/index.html
- https://docs.nvidia.com/doca/sdk/doca-legal-notices-and-3rd-party-licenses/index.html
- https://catalog.ngc.nvidia.com/orgs/nvidia/doca/resources/doca_hbn/3.4.0/version-history
- https://docs.nvidia.com/doca/sdk/snap-virtio-fs-service-release-notes/index.html
- https://www.delltechnologies.com/asset/de-at/products/networking/technical-support/nvidia-quantum-x800-q3200-ra-and-q3400-ra-datasheet.pdf
- http://gpusmith.com/datasheets/nvidia-quantum-x800-datasheet.pdf
- https://www.aicplight.com/products/switches-nics/nvidia-infiniband-switches/quantum-x800-switches/q3200-ra/
- https://www.aicplight.com/products/switches-nics/nvidia-infiniband-switches/quantum-x800-switches/q3400-ra/
- https://www.hpe.com/psnow/generateDDS/HPE Aruba Networking CX 8100 Switch Series data sheet-PSN1014733547PHEN.pdf?oid=1014733547&cc=PH&lc=EN&softroll=0&prelaunch=false&print=§ion=&prelaunchSection=&deepLink=&utm_source=&utm_campaign=&utm_content=&utm_term=&isLinearized=false&contentDisposition=attachment
- https://shop.levata.com/media/pdf/Aruba/cx8100.pdf
- https://www.lttpartners.com/products/aruba-cx-8100-24xt4xf4c-ethernet-switch
- https://www.avendor.com/products/aruba-cx-8100-24xf4c-ethernet-switch
- https://www.itk.co.th/network-switch/hp-switch/aruba-cx-8100-switch-series/r9w88a
- https://nvidianews.nvidia.com/_gallery/download_pdf/5e989f8aed6ae55d9f69d451/
- https://www.ainvest.com/news/mellanox-nvidia-important-acquisition-noticed-2606/
- https://datacentrereview.com/2020/04/nvidia-edges-closer-to-completing-6-9-billion-mellanox-acquisition/
- https://www.minichart.com.sg/2026/09/15/dell-raises-us5-billion-in-senior-notes-offering/
- https://nation.lk/online/dell-secures-deal-worth-more-than-5b-in-elon-musks-xai-startup-company-296563.html

**Round-4 collection metadata:** read-only web research (browser_search, 2026-09-22); no live-browser visits; nothing sent externally. No identifiers guessed. All new facts carry provenance tags; new open items are listed in the R4 verification log above. Earlier sections were not modified.

---

## Supplementary / Complementary Research Pass — round 4 addendum (2026-09-22, "QFX5140 + HPE Helios fabric")

**Scope note:** Six concurrent round-4 sibling passes already exist in this file and were read in full; overlap was checked item by item (Cumulus 5.18.1, 802.1X telemetry, Enterprise SONiC 4.6.0 Bookworm/FRR 8.2.2/Broadcom SDK 6.5.35, CVE-2026-73749, NVIDIA Mellanox acquisition timeline, AMD Helios rack-scale AI architecture are all covered). This addendum adds ONLY the two data-center-networking items that remain genuinely absent: (a) the new Juniper QFX5140 switch for AI inference edge, and (b) HPE's "Helios" rack-scale Ethernet fabric enabled by the HPE–Juniper combination — distinct from AMD's "Helios" MI455X rack architecture already covered above. Nothing above was modified. Provenance legend from the base report applies. All facts as of 2026-09-22. Research method: browser_search only; no live-browser visits; nothing sent externally.

### ADD-1. Juniper QFX5140 — new AI-inference-edge switch (2026) [secondary]

