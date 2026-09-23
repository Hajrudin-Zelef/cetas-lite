---
id: etape6-phasec-optics-cabling/02-wave-13-oem-vs-third-party-transceiver-compatibility-exact-e/15-2-800g-nics-dpus
title: "15.2 800G NICs / DPUs"
domain: wave-13-oem-vs-third-party-transceiver-compatibility-exact-e
role: deep-dive
task: regulation
actors: ["AMD", "AWS", "Broadcom", "Cohere", "CoreWeave", "Crusoe", "Intel", "Lambda", "Meta", "Nvidia", "Oracle", "United States", "xAI"]
dates: ["2024-07", "2025-03", "2025-10", "2025-11", "2026-08", "2026-09", "2026-09-22"]
keywords: ["2nm", "amd", "aws", "blackwell", "chiplet", "cpo", "dsp", "ethernet", "gpu", "gpus", "intel", "nvidia"]
source: docs/RAG/etape6_phaseC_optics_cabling.md
source_anchor: ""
source_lines: [2654, 2699]
section: "Wave 13 — OEM vs third-party transceiver compatibility: exact error messages, TAC/warranty policy, FEC, failure behavior"
sha256: 9f3d0fb858e9ced41d3947ff365ef3454274eae28bf776580f45043de1a27b67
---

# 15.2 800G NICs / DPUs

### 15.2 800G NICs / DPUs

- **NVIDIA ConnectX-8** — FS.com (US) catalog lists: (a) **1-port 800G OSFP**, PCIe 6.0 ×16, **US$2,519**; (b) **2-port 400G QSFP112** (aggregate 800G), **US$2,519** [vendor-reported] (https://www.FS.com/c/nvidia-ethernet-nics-4014). FS UK: ConnectX-8 800G OSFP at **£1,991 (VAT excl.) / £2,389.20 (VAT incl.)**, 55 sold [vendor-reported] (https://www.fs.com/uk/c/nvidia-ethernet-nics-4014). **Retailer availability only — official NVIDIA ConnectX-8 announcement/GA page: not found.**
- TrendForce NVIDIA roadmap table: **ConnectX-8 (800G/port, PCIe 6.0)** under the Blackwell generation; **ConnectX-9 (1.6T)** under Rubin (2H26); **ConnectX-10 (3.2T)** under Feynman (2028) [secondary] (https://www.trendforce.com/insights/infiniband-vs-ethernet).
- **Broadcom Thor Ultra** — 800G AI Ethernet NIC adopting the open Ultra Ethernet Consortium (UEC) spec, PCIe Gen6 ×16, PCIe CEM and OCP 3.0 form factors; **sampling** (not shipping) as of October 2025 reporting [independent — Network World / TechSpot / EDN, based on Broadcom announcement] (https://www.networkworld.com/article/4072253/broadcom-drops-the-hammer-on-ai-networking-with-thor-ultra.html). **Production GA status: not found.**
- **AMD Pensando Vulcano 800 AI NIC** — August 2026 launch reporting: up to 800 Gbps Ethernet per NIC and up to 2.4 Tbps scale-out bandwidth per GPU in a 3-NIC-per-GPU configuration [secondary] (https://infotechlead.com/networking/amd-pensando-vulcano-800-challenges-nvidia-connectx-8-as-ai-networking-race-heats-up-broadcom-arista-and-marvell-intensify-competition-97488).
- Competitor 800G NIC production GA status (Broadcom, AMD, Marvell, Intel): **not found.**
- A small Medium vendor article shows 800G interconnect solutions built on NVIDIA Spectrum-4 SN5610 switches paired with ConnectX-8 C8240 series NICs [unverified] (https://medium.com/@aicplight888/1-6t-800g-switch-connectx-8-cx-8-nic-interconnection-solutions-20182fb0d4e5).

### 15.3 Named 800G AI-cluster deployments

- Asia Business Daily (dated **5 November 2025**) reports xAI's Colossus cluster uses "800G Spectrum switches" and that Meta adopted 800G infrastructure — **no official corroboration found** [secondary] (https://www.asiae.co.kr/en/article/2025110510025849323).
- **CoreWeave Vera Rubin deployment (2026)** — a 1.6T-era data point, not 800G: StorageReview reports hundreds of Rubin GPUs with ConnectX-9, BlueField-4, Spectrum-X fabric, **1.6 Tb/s per GPU** [independent] (https://www.storagereview.com/news/coreweave-brings-up-a-multi-rack-vera-rubin-nvl72-cluster-hundreds-of-rubin-gpus-1-6-tb-s-per-gpu-and-a-no-fee-archive-tier).
- Official statements from xAI, Meta, AWS, Oracle, Lambda, or Crusoe explicitly confirming **800G port/NDR/XDR deployment dates**: **not found.** Treat all named-800G-deployment claims as [secondary]/unverified unless corroborated.
- EE Times (observed 2026-09-22) reports analyst forecasts of a **40–60% 800G transceiver shortfall through 2027** and possible multi-quarter cluster delays if import restrictions bite [independent] (https://www.eetimes.com/fcc-rule-on-optical-connectivity-could-slow-ai-race/).

### 15.4 800G optic price snapshots — FS.com vs alternatives (single-retailer, observed ~2026-09-22)

- **800G OSFP DR8 500m:**
  - QSFPTEK generic SiPh (Broadcom 7nm DSP): **US$699** (3 in real-time stock, dated 13 Sept 2026) [vendor-reported] (https://www.qsfptek.com/product/103578.html).
  - QSFPTEK NVIDIA MMS4X00-NS-compatible 100m: **US$799** (19 in real-time stock, dated 9 Sept 2026) [vendor-reported] (https://www.qsfptek.com/product/103592.html).
  - FS.com Arista-compatible: **US$1,319** (~1K sold) [vendor-reported] (https://www.fs.com/c/osfp-200-400-800g-4089).
  - ATGBICS via DigiKey marketplace: **US$3,067** (qty 1, 669 in stock, ~14-day shipment) [vendor-reported] (https://www.digikey.com/en/products/detail/atgbics/OSFP-800G-DR8-FLT-MSA-AT/25582709) — marketplace listing, not DigiKey direct.
  - ATGBICS direct (UK): **£1,959** (InfiniBand variant) [vendor-reported] (https://www.atgbics.com/800g-osfp-dr8-transceiver-500m-osfp-800g-dr8-flat-top-optical-infiniband-at).
  - SHI (Axiom, Juniper OSFP-800G-DR8-P equivalent): **US$2,866** (MSRP US$3,128) [vendor-reported] (https://www.shi.com/product/50881435/Axiom-OSFP112-transceiver-module-(equivalent-to:-Juniper-OSFP-800G-DR8-P)).
- **800G QSFP-DD DR8 500m:**
  - QSFPTEK generic (EML+PIN): **US$1,021.90** (2 in real-time stock, dated 14 Sept 2026), 137 sold [vendor-reported] (https://www.qsfptek.com/product/102570.html).
  - FS.com Cisco-compatible: **US$1,429** [vendor-reported] (https://www.fs.com/c/osfp-200-400-800g-4089).
  - ATGBICS direct (UK): **£1,916** [vendor-reported] (https://www.atgbics.com/800g-qsfp-dd-dr8-fibre-transceiver-500m-qsfp-dd-800g-dr8-flat-top-optical-at).
  - ATGBICS via DigiKey marketplace (twin-port 800G DR8 InfiniBand): **US$1,641** (3,255 in stock) [vendor-reported].
- **800G OSFP SR8:** FS.com NVIDIA-compatible **US$879** [vendor-reported] (https://www.fs.com/c/osfp-200-400-800g-4089).
- **800G QSFP-DD800 (FlexOptix, official):** 2x DR4 D.134HG2.05 as low as **EUR 1,604.90** (3 pcs, +10 expected ~Dec 9, 2026); SR8 D.858HG.005.MP as low as **EUR 1,599.26** (2 pcs) [official] (https://www.flexoptix.net/en/transceiver).
- **Pricing conflict flagged:** 800G DR8 third-party pricing spans **US$699** (QSFPTEK, low stock) to **US$3,067** (DigiKey marketplace) — single-retailer snapshots with different stock levels and dates, not a market average. Westbury Photonics listed 800G OSFP DR8 at **£182** but marked "Sold out" with internally inconsistent specs ("500m across MMF") — do not treat as a comparable in-stock price [unverified] (https://www.westburyphotonics.com/products/800g-osfp-dr8-transceiver).
- **Approved Networks** list price **US$2,275** for 800G OSFP DR8 (article dated 17 July 2024); current September 2026 direct-retailer price **not found** [secondary] (https://convergedigest.com/approved-networks-adds-osfp-800g-dr8-transceiver/?amp=1).
- **10Gtek** 800G OSFP SR8 product pages exist, but **no price visible** — not found (https://store.10gtek.com/800g-sr8-osfp-optical-transceiver/p-21208).

### 15.5 Silicon photonics — shipment status (Sept 2026)

- **Intel** — reporting based on Intel's announcement: Intel shipped more than **8 million silicon-photonic PICs** in pluggable modules (100/200/400G); 200G/lane PICs for 800G/1.6T were under development; the OCI chiplet remained a prototype with select-customer co-packaging work [secondary, based on Intel announcement] (https://www.techPowerUp.com/323915/intel-demonstrates-first-fully-integrated-optical-io-chiplet).
- **Shipping SiPh-based pluggables in retail channels:** QSFPTEK 800G OSFP DR8 product page states silicon-photonic chips plus Broadcom 7nm DSP, US$699 [vendor-reported] (https://www.qsfptek.com/product/103578.html); FS 400G DR4 (SKU 128242, P/N **QDD-DR4-400G-Si**) is described by FS as silicon-photonics-based with Broadcom 7nm DSP, Cisco-compatible, US$749 [vendor-reported] (https://www.fs.com/products/128242%20.html); QSFPTEK generic 400GBASE-DR4 SiPh (Broadcom 7nm DSP) at US$674.90 [vendor-reported] (https://www.qsfptek.com/product/103832.html); PRO-OPTICS generic 400G QSFP-DD DR4 SiPh at US$699.00 [vendor-reported] (https://www.pro-optics.com/store/product/400g-qsfp-dd-dr4).
- **Ayar Labs** — promotional claims of high-volume TeraPHY Gen 3 production surfaced only on low-authority marketing-style sites; **do not use without official corroboration**. September 2026 production/sampling status: **not found**.
- **Lightmatter** — Reuters-derived report dated 31 March 2025: interposer releasing in 2025, chiplet in 2026, manufactured by GlobalFoundries. Current September 2026 shipping/customer status **not verified** [independent] (https://srnnews.com/lightmatter-releases-new-photonics-technology-for-ai-chips/).
- **Marvell ECOC 2026 demos** (20–21 Sept 2026 reporting): 2nm 800G ZR/ZR+, 1.6T ZR/coherent-lite, and a 102.4T CPO platform — demonstrations, no availability dates [secondary] (https://biztechweekly.com/at-ecoc-2026-marvells-2nm-optics-point-toward-3-2t-ai-networks-not-yet-deployments/).
- **Knowledge-base distinction (important):** Ayar Labs and Lightmatter are optical-I/O / CPO / in-package interconnect platforms — **not** ordinary shipping pluggable-transceiver vendors. Conflating them with 800G pluggable supply would be wrong.
- Direct statements from Cisco/Acacia, Marvell, Coherent, Innolight, Eoptolink, or Lumentum identifying which commercial 400G/800G modules use silicon photonics: **not found.**

