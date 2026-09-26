---
id: etape9-phaseb-form-factors-interfaces/00-form-factors-interfaces/32-document-revision-note
title: "32. Document revision note"
domain: step-9-phase-b-storage-form-factors-interfaces-hardware-angl
role: deep-dive
task: hardware
actors: ["Broadcom"]
dates: ["2026-09-22"]
keywords: ["research"]
source: docs/RAG/etape9_phaseB_form_factors_interfaces.md
source_anchor: ""
source_lines: [740, 751]
section: "Step 9 — Phase B: Storage Form Factors & Interfaces (Hardware Angle)"
sha256: 81c8aea4f50e50101463d5d04baf6f3800e14737a169ee7855cd404279e3612e
---

# 32. Document revision note

- Cable example: Broadcom 05-60002-00 — 1 m SlimSAS x8 SFF-8654 to 2× x4 SFF-8643 Mini-SAS HD (bridges 95xx adapters to older backplanes) [secondary](https://www.scan.co.uk/products/1m-broadcom-05-60002-00-slimsas-x8-sff-8654-to-two-x4-sff-8643-mini-sas-hd-nvme-connection).
- PCIe switch example: HighPoint Rocket 1624A — Broadcom PEX89048, Gen5 x16 host, 2× MCIO 8i, up to 16 NVMe SSDs [vendor-reported](https://www.scan.co.uk/products/highpoint-rocket-1624a-hba-adapter-2x-mcio-gen5-x8-pcie-50-x16-32-gb-s-hot-swap-broadcom-pex89048-sw).
- UBM controller example: Microchip EEC1005-UB2 — SFF-TA-1005 v1.4, Cortex-M4, up to 16 drives, $7.33 @10k (Mar 2024) [vendor-reported](https://www.storagenewsletter.com/2024/03/14/microchip-eec1005-ub2-universal-backplane-management-controller/).

## 32. Document revision note

- **v1.0 — 2026-09-22**: Initial single-writer completion. 31 sections, all research current to 2026-09-22.
- Waves: header + EDSFF (§1–4) → M.2/U.2/U.3 (§5 header region) → other form factors, comparison matrix, NVMe spec family (§5–7) → command sets, NVMe-oF, advanced features (§8–10) → PCIe generations, SAS/SATA (§11–12) → backplanes, cabling (§13–14) → server integration, decision guide, gaps, glossary, sources (§15–19) → timeline, interop, lane budget, retimers, CXL, software (§20–25) → capacity, security, firmware, thermal (§26–29) → cheat sheet (§30) → SKU tables (§31).
- No other workspace files were modified during this task.
- Suggested maintenance: re-verify NVMe 2.4 drive availability, PCIe 6.0 SSD launches, and SAS-48G ratification when revisiting after Q1 2027.

- RAG ingestion: chunk on `##` section boundaries; each section is self-contained with its own citations.
