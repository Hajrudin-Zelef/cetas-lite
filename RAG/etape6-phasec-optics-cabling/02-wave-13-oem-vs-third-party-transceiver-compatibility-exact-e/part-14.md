---
id: etape6-phasec-optics-cabling/02-wave-13-oem-vs-third-party-transceiver-compatibility-exact-e/part-14
title: "Wave 13 — OEM vs third-party transceiver compatibility: exact error messages, TAC/warranty policy, FEC, failure behavior (part 14)"
domain: wave-13-oem-vs-third-party-transceiver-compatibility-exact-e
role: deep-dive
task: regulation
actors: ["Broadcom", "Nvidia"]
dates: ["2024-09", "2025-06", "2025-10", "2026-07"]
keywords: ["asic", "ethernet", "lpo", "nvidia", "optics"]
source: docs/RAG/etape6_phaseC_optics_cabling.md
source_anchor: ""
source_lines: [2643, 2653]
section: "Wave 13 — OEM vs third-party transceiver compatibility: exact error messages, TAC/warranty policy, FEC, failure behavior"
sha256: 0f511bc5a816e71d19500b6f0c6d9b33519ca5b28e811f72c2edb8310c10bff2
---

# Wave 13 — OEM vs third-party transceiver compatibility: exact error messages, TAC/warranty policy, FEC, failure behavior (part 14)

- **Arista 7800R4** — official Arista release dated **29 October 2025**: up to **576×800GbE** with 36-port 800G linecards; **7280R4** variants include **32×800GbE** and **10×800GbE + 64×100GbE**, all shipping [official] (https://investors.arista.com/Communications/Press-Releases-and-Events/Press-Release-Detail/2025/Arista-Networks-Unveils-Next-Generation-Data-and-AI-Centers/default.aspx).
- **Arista 7060X6** — secondary reporting identifies it as **64×800G on Broadcom Tomahawk 5**; the 7800R4 is described as using **Broadcom Jericho3-AI** [secondary] (https://www.demandtalk.com/news/it-infra-news/arista-launches-the-etherlink-ai-platforms-for-ai-workloads/).
- **Cisco 8122-64EH/EHF** — **64×800G**, 2RU, built on **Cisco Silicon One G200**, with QSFP-DD800 or OSFP optics options [independent] (https://www.networkworld.com/article/3564900/cisco-pumps-up-data-center-networking-with-ai-large-workloads-in-mind.html). **Cisco Silicon One G200 datasheet** [official] supports configurations from **64×800GE to 512×100GE** (https://www.cisco.com/c/en/us/solutions/collateral/silicon-one/silicon-one-g200-ds.pdf).
- **Cisco 8122X-64EF-O** — **64×800G SONiC** switch from the G200 family with Cisco 800G LPO support [independent] (https://www.networkworld.com/article/4130263/cisco-amps-up-silicon-one-line-delivers-new-systems-and-optics-for-ai-networking.html).
- **FS 800G TH5 switch** — FS says its Tomahawk-5 system offers **64×800G OSFP** (BCM78900), listed "available now" (announcement circa September 2024) [vendor-reported] (https://www.fs.com/sg/blog/fs-unveils-512t-400g-and-800g-ethernet-switches-powered-by-broadcom-tomahawk-5-10112.html).
- **Broadcom Tomahawk 5 (BCM78900)** — press-release copy: **64×800GbE**, 51.2 Tb/s aggregate, "ships" status [vendor-reported — hosted press release, not Broadcom's own page] (https://www.nasdaq.com/press-release/broadcom-ships-tomahawk-5-industrys-highest-bandwidth-switch-chip-to-accelerate-ai-ml).
- **Broadcom Tomahawk 6** — secondary report dated **4 June 2025**: Broadcom shipping the 102.4 Tb/s chip with 1.6T-port support [secondary] (https://www.storagenewsletter.com/2025/06/04/broadcom-ships-tomahawk-6-first-102-4tb-s-switch). **TH6-based shipping switch systems with model numbers and GA dates: not found.**
- **Marvell Teralynx 10 (ASIC)** — official product brief: **64×800G / 128×400G / 32×1.6T**, 51.2 Tb/s [official] (https://cn.marvell.com/content/dam/marvell/en/public-collateral/switching/marvell-teralynx-10-data-center-ethernet-switch-product-brief.pdf). A July 2026 secondary report says Teralynx 10 had begun production/customer deployment [secondary] (https://convergedigest.com/marvell-ships-51-2-tbps-ethernet-switch-for-ai-data-centers/) — **specific shipping switch-system models based on Teralynx 10: not found.**
- **NVIDIA Spectrum-4 / Spectrum-X800 800G systems** (model numbers, port counts, GA dates): **not found** in the evidence gathered — do not treat catalog presence as GA.
- White-box / ODM 800G systems (Edgecore, Celestica, etc.): not systematically covered in this wave — gap.

