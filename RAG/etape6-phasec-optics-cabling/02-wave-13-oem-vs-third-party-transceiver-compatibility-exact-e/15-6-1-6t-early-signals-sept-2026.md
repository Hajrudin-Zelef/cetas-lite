---
id: etape6-phasec-optics-cabling/02-wave-13-oem-vs-third-party-transceiver-compatibility-exact-e/15-6-1-6t-early-signals-sept-2026
title: "15.6 1.6T early signals (Sept 2026)"
domain: wave-13-oem-vs-third-party-transceiver-compatibility-exact-e
role: deep-dive
task: regulation
actors: ["Cohere", "CoreWeave", "EU", "Nvidia", "United States"]
dates: ["2026-03", "2026-04", "2026-09", "2026-09-22"]
keywords: ["alignment", "consumer", "ethernet", "gpu", "gpus", "nvidia", "optics", "pricing", "rubin", "serdes", "training", "vera rubin"]
source: docs/RAG/etape6_phaseC_optics_cabling.md
source_anchor: ""
source_lines: [2700, 2748]
section: "Wave 13 — OEM vs third-party transceiver compatibility: exact error messages, TAC/warranty policy, FEC, failure behavior"
sha256: afe9dbb9c0d699e227470038fd466cdc0ab14277f48208157e8cc2ac75d08c6e
---

# 15.6 1.6T early signals (Sept 2026)

### 15.6 1.6T early signals (Sept 2026)

- **Ethernet Alliance ECOC 2026 demo** — official press release dated **9 September 2026** announces a multi-vendor demo at ECOC 2026 (21–23 September 2026, Málaga) covering **1.6T OSFP connectivity**, 224G SerDes, link training, and the 400G/800G/1.6T ecosystem [official] (https://www.globenewswire.com/news-release/2026/09/09/3358828/0/en/1-6t-ethernet-moves-closer-to-reality-in-ethernet-alliance-s-ecoc-2026-demo.html).
- **Cisco 1.6T pluggable optics** — SDxCentral reports Cisco Live EMEA 2026 (early 2026) showcased new 1.6T pluggable optics. Exact module models, price, sampling/GA: **not found** [independent] (https://www.sdxcentral.com/news/next-gen-switches-servers-everything-unveiled-at-cisco-live-emea-2026/).
- **Coherent ECOC 2026** — secondary release: Coherent will demo a **3.2T OSFP-size technology** built from dual 1.6T optical paths (eight 425G PAM4 line-side lanes, differential EMLs) — a technology demo, **not a shipping 1.6T module** [secondary] (https://www.marketnewsdesk.com/index.php/coherent-showcases-optical-innovations-to-scale-ai-infrastructure-at-ecoc-2026/).
- **IEEE P802.3dj** — Network World reported it defines 200G/400G/800G/1.6T at 200G/lane and was on track for late-2026 completion [independent] (https://www.networkworld.com/article/4113364/ethernet-groups-keep-2026-focus-on-higher-bandwidth-ai-demands.html); a secondary report says it entered Working Group Ballot [secondary] (https://convergedigest.com/ieee-p802-3dj-moves-forward-with-200g-to-1-6t-ethernet-working-group-work/); EEWorld warned the 2026 schedule could slip [independent] (http://eeworldonline.com/how-ieee-802-3df-brings-800g-ethernet-to-life/). **Status as of 22 Sept 2026: "late-2026 target" vs. "schedule at risk" — not yet a completed standard on collected evidence.**
- A 22 September 2026 secondary article claims 1.6T engineering samples are in customer validation with volume production late 2026–2027 — use only as [secondary] (https://roboticsandautomationnews.com/2026/09/22/osfp-modules-the-complete-guide-to-400g-800g-and-1-6t-optical-transceivers-for-ai-and-hyperscale-data-centers/104982/).
- **OSFP-XD:** 1.6T OSFP-XD-specific module pricing, sampling, or GA evidence: **not found.** Strongest demo evidence is generic 1.6T OSFP — do not infer OSFP-XD from it.
- CoreWeave's 2026 Vera Rubin deployment (ConnectX-9, BlueField-4, Spectrum-X, **1.6 Tb/s per GPU**) is a 1.6T-era deployment data point [independent] (https://www.storagereview.com/news/coreweave-brings-up-a-multi-rack-vera-rubin-nvl72-cluster-hundreds-of-rubin-gpus-1-6-tb-s-per-gpu-and-a-no-fee-archive-tier).

### 15.7 Regional availability & lead times

- **FS.com fulfillment (official):**
  - US: in-stock US-warehouse orders ship within 24 hours on working days [official] (https://www.fs.com/shipping_delivery-p0001.html).
  - Europe: qualified delivery within **1–3 business days** from FS's Munich-area warehouse [official] (https://www.fs.com/eu-en/shipping_delivery-p0001.html).
  - UK: in-stock UK orders ship within 24 hours on working days [official] (https://www.fs.com/uk/shipping_delivery-p0001.html).
  - Australia/APAC: in-stock Australian warehouse orders ship within 24 business hours; standard delivery **1–5 business days** [official] (https://www.fs.com/au/shipping_delivery-p0001.html).
  - Global/local transfer: if local stock is insufficient, transfer from the global warehouse takes about **5–8 business days** [official] (https://www.fs.com/global-warehouse-h0002.html).
  - Pickup/warehouse locations include New Castle (Delaware, US), Karlsfeld (Germany), Birmingham (UK), Singapore, Dandenong South (Australia), and Tokyo [official] (https://www.fs.com/sg/pickup_warehouse-p0002.html).
- **QSFPTEK:** 800G pages show global stock but not region-specific transit times; one DAC page showed US delivery in 5 business days (March 2026) [vendor-reported] (https://www.qsfptek.com/product/102645.html).
- **FlexOptix:** prices "ex Darmstadt"; some 800G items "4–6 weeks" / "5–7 weeks after ordering"; replenishment dates displayed (e.g. +10 pcs expected ~Dec 9, 2026) [official] (https://www.flexoptix.net/en/transceiver).
- **Supply constraints:**
  - TrendForce (April 2026): EML/CW-laser and optical-alignment capacity remain tight as AI optics demand rises [independent] (https://www.trendforce.com/presscenter/news/20260420-13017.html).
  - TrendForce-derived reporting: NVIDIA capacity reservations extended EML lead times beyond 2027 [secondary] (https://www.ledinside.com/intelligence/2025/12/2025_12_09_01).
  - EE Times (observed 22 Sept 2026): analyst forecasts of a **40–60% 800G transceiver shortfall through 2027** and possible multi-quarter cluster delays if import restrictions bite [independent] (https://www.eetimes.com/fcc-rule-on-optical-connectivity-could-slow-ai-race/).
- A Medium vendor article claims typical 800G lead times of 4–8 weeks and 1.6T of 12–16 weeks — [unverified], do not present as established regional data (https://medium.com/@aicplight888/800g-vs-1-6t-transceivers-which-one-fits-your-ai-data-center-in-2026-198a2542fb3d).
- Independent, region-specific US/EU/APAC lead-time data for optics/DAC/cabling: **not found.**

### 15.8 Consolidated gaps / open-items log (all waves, updated 2026-09-22)

- **OEM list prices:** no verified OEM list prices for 100G/400G/800G optics; the only Cisco figure is a reseller-stated MSRP (QSFP-100G-SR4-S US$2,440.10). All price ratios remain indicative snapshots.
- **EEPROM check fields:** Cisco's exact checked fields (vendor name/ID/serial/security code/CRC) are secondary-only; Aruba's authentication claim is unverified; Ubiquiti and MikroTik publish no field-level detail.
- **Warranties:** Approved Networks and ProLabs exact warranty exclusions not found; 10Gtek 400G/800G warranty length not specified; FlexOptix has no lifetime claim (12-month GTC term).
- **White-box/ODM 800G systems** (Edgecore, Celestica, others): not systematically covered.
- **NVIDIA:** official third-party-optic policy statement not found; official ConnectX-8 announcement/GA page not found; Spectrum-4/Spectrum-X800 800G system model numbers/GA dates not found.
- **Dell:** official third-party-optics support policy not found; Dell SONiC lock behavior is community consensus only.
- **1.6T / OSFP-XD:** no module pricing, sampling, or GA evidence; no OSFP-XD-specific evidence — do not infer from generic OSFP demos.
- **IEEE 802.3dj:** not yet a completed standard on collected evidence (late-2026 target vs schedule-risk reports).
- **Silicon photonics:** no direct vendor statements from Cisco/Acacia, Marvell, Coherent, Innolight, Eoptolink, or Lumentum identifying SiPh-based commercial modules; Ayar Labs production status not found; Lightmatter current status unverified.
- **Regional lead times:** no independent, region-specific US/EU/APAC lead-time datasets found; only vendor-stated fulfillment terms and analyst shortfall forecasts.
- **Reliability engineering:** no independent laser-aging, FIT/MTBF, or field-failure-rate figures verified anywhere in Phase C — do not invent them.
- **Legal:** no primary Magnuson-Moss (FTC) or EU/UK competition-law sources gathered; consumer warranty law does not automatically govern enterprise/B2B support contracts.
- **Counterfeit 400G/800G:** no official 2026 vendor statements specifically targeting counterfeit 400G/800G optics; do not overgeneralize older Cisco brand-protection material.
- **Source-quality standing note:** treat search "Last Updated / Last Crawl" metadata as crawl timing, not publication dates; treat catalog presence alone as retailer availability, not official GA; never infer 1.6T OSFP-XD from generic OSFP demo coverage; treat QSFPTEK/FS.com stock counters as retailer displays, not audited inventory.

*End of Wave 15. Phase C file remains append-only; earlier waves untouched.*

---

