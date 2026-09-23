---
id: etape6-phasef2-fpga/00-fpga/26-conflicts-gaps-uncertainty-consolidated
title: "26. Conflicts, gaps, uncertainty (consolidated)"
domain: phase-f2-fpga-field-programmable-gate-arrays
role: deep-dive
task: reference
actors: ["AWS", "Alibaba", "China", "Intel"]
dates: ["2025-12-20", "2026-09-22"]
keywords: ["acquisition", "amd", "aws", "benchmarks", "cost", "dsp", "hyperscaler", "intel", "memory", "pricing"]
source: docs/RAG/etape6_phaseF2_fpga.md
source_anchor: ""
source_lines: [386, 433]
section: "Phase F2 — FPGA (Field-Programmable Gate Arrays)"
sha256: 4634e9a24f567d332e05e00eeea11461a9f3f36376deaf172fc000b4d0453ece
---

# 26. Conflicts, gaps, uncertainty (consolidated)

## 26. Conflicts, gaps, uncertainty (consolidated)

- "Napatech (now Intel)": **no acquisition evidence** — do not repeat `[unverified]`
- AWS F1 EOL 2025-12-20: single third-party source — **verify against aws/aws-fpga** `[unverified]`
- Versal Premium Gen 2 production H2 2026; MoP sampling end-2026 / production H2-2027: **planned** `[vendor-reported]`
- Agilex 9 Direct RF production Q3 2026: **planned** `[vendor-reported]`
- Accolade Technology detail: **gap** `[unverified]`
- Lattice Avant 30/50 pricing and eval-board availability: **not public** as of Sep 2026 `[secondary]`
- Nexus 2 FIPS 140-3 L2: vendor claim — **verify in NIST CMVP** before citing as certified `[unverified]`
- RT PolarFire / Chinese hyperscaler FPGA status / Alibaba-T-Head: **gaps** `[unverified]`
- Independent 2026 benchmarks and TCO studies: **gap** `[unverified]`

## 27. Additional key URLs (sections 16–24)

- https://www.mouser.com/catalog/specsheets/Intel Corporation_04-17-2026_ag-overview-683458-666707.pdf
- https://www.cnx-software.com/2019/07/22/intel-agilex-soc-fpga-features-four-arm-cortex-a53-cores/
- https://www.businesswire.com/news/home/20250401117009/en/Altera-Starts-Production-Shipments-of-Industrys-Highest-Memory-Bandwidth-FPGA
- https://www.storagenewsletter.com/2025/04/04/altera-starts-production-shipments-of-agilex-7-fpga-m-series-highest-memory-bandwidth-fpga/
- https://amd.nt-rt.ru/images/manuals/k3-1.pdf
- https://www.eenewseurope.com/en/ai-and-dsp-boost-for-fpga-signal-processing/
- https://siliconangle.com/2022/11/15/amd-announces-reprogrammable-versal-ai-chips-ready-space-based-deployments/
- https://www.Theregister.Com/2022/11/15/amd_versal_space_chip/?td=keepreading
- https://www.hpcwire.com/2023/06/28/fpga-development-brings-amd-and-intel-competition-to-the-forefront/
- https://www.napatech.com/media/press-releases/napatech-leverages-latest-intel-agilex-fpga-to-launch-industrys-first-400gbps-smartnic-solutions/
- https://www.napatech.com/media/press-releases/napatech-launches-5g-upf-offload-solution-on-its-intel-agilex-fpga-based-400gbps-smartnic/
- https://cdn-adepci1.actonsoftware.com/acton/cdna/14951/f-d0dea8a4-3331-49df-910e-2faaebb4575f/1/10
- https://www.lightwaveonline.com/data-center/article/14298298/napatech-launches-400-gbps-smartnic-solutions
- https://www.edn.com/lattice-launches-small-size-fpga-platform/
- https://www.cnx-software.com/2024/12/12/lattice-unveils-nexus-2-small-fpga-platform-lattice-avant-30-and-avant-50-mid-range-devices-updated-lattice-design-software-tools/
- https://www.hackster.io/news/lattice-semiconductor-launches-the-certus-n2-small-fpga-family-built-around-the-nexus-2-platform-f0e437012f1b.amp
- https://www.futurumgroup.com/insights/lattice-mach-n2-turns-the-post-quantum-procurement-gate-into-an-fpga-moat/
- https://www.eeworldonline.com/pcie-gen-4-integration-added-to-small-form-factor-fpga/
- https://latticesemi-prod-latticesemi-live.vercel.app/view_document?document_id=54437
- https://github.com/trabucayre/openfpgaloader/blob/HEAD/doc/vendors/intel.rst
- https://github.com/machdyne/zeitlos/blob/HEAD/docs/toolchain.md
- https://www.dataweek.co.za/25105r
- https://it-online.co.za/2023/09/18/intel-expands-fpga-portfolio/

*End of Phase F2 file — FPGA. Observation date 2026-09-22.*

## 28. Agilex 5 — series detail

- Agilex 5 is Altera's power- and cost-optimized midrange tier below Agilex 7, targeting edge AI, industrial, video/broadcast, and communications `[official]`
- Agilex 5 ships in **D-Series** (higher density/performance) and **E-Series** (enhanced value) variants `[official]`
- **MTI** selected Agilex 5 E-Series for a 4T4R macro radio unit with 25GbE fronthaul (Feb 2026 announcement) `[vendor-reported]`
- Transceiver rates on Agilex 5 reach up to 28 Gbps class (midrange positioning vs 58G/112G/116G on Agilex 7) `[unverified]`
- Agilex 5 SoC variants integrate Arm application processors for single-chip edge systems `[official]`

