---
id: etape9-phasea-enterprise-ssd/00-enterprise-ssd/28-gen4-vs-gen5-purchase-decision-2026
title: "28. Gen4 vs Gen5 purchase decision (2026)"
domain: step-9-enterprise-ssd-hardware-phase-a
role: deep-dive
task: hardware
actors: ["Samsung"]
dates: []
keywords: ["cost", "datacenter", "dram", "latency", "memory", "nand", "pricing", "throughput", "training"]
source: docs/RAG/etape9_phaseA_enterprise_ssd.md
source_anchor: ""
source_lines: [576, 618]
section: "Step 9 — Enterprise SSD Hardware (Phase A)"
sha256: a852128b21aee2784c84edd7d4eb4357ebeb18d1d82a00d056b43a2878d03f2e
---

# 28. Gen4 vs Gen5 purchase decision (2026)

- $/TBW derivation: TBW = capacity × DWPD × 365 × 5 (formula matches published TBW for all rows above); prices are single-channel snapshots, currencies as listed [independent calculation from cited sources].
- Reading the table: 3-DWPD and large-capacity SKUs dominate $/TBW; small-capacity 1-DWPD SKUs are the worst value — buy the largest capacity the workload fits [independent guidance].
- OEM-branded (Lenovo/Dell/HPE) SKUs carry 2–10x markups over street pricing — the PM1743 $4,671/TB row is a channel artifact, not NAND economics [secondary](https://www.shi.com/product/45962451/Samsung-PM1743-SSD).

---

## 28. Gen4 vs Gen5 purchase decision (2026)

- Buy Gen5 when: sequential throughput >7 GB/s is needed (AI checkpointing, NVMe-oF targets, all-flash Ceph), random-read IOPS >1.5M per drive matter, or the server platform is Gen5-native and lane budget allows [independent guidance].
- Buy Gen4 when: the workload is IOPS/latency-bound at queue depths Gen4 already serves (most virtualization, databases), power/thermal budget is tight (Gen4 ~8–14 W vs Gen5 ~18–29 W), or $/TB is the deciding factor — Gen4 drives (PM9A3, Micron 7450, Kioxia CD8) are discounted as Gen5 takes over [independent guidance].
- Lane math: Gen5 x4 ≈ 2x Gen4 x4 bandwidth; a Gen4 drive in a Gen5 slot works at Gen4 speed (backward compatible) — populating Gen5 servers with discounted Gen4 drives is a legitimate cost play [independent guidance].
- Platform note: Gen5 needs platform-level signal integrity (retimers/redrivers on long traces) — backplane and cable compatibility (SFF-8654/SlimSAS Gen5-rated) must be verified; Gen4 cabling on Gen5 links causes link-training failures [independent guidance].
- EOL watch: Samsung PM9A3 and Micron 7450 remain in production at cutoff but are previous-generation; confirm firmware/EOL status with the vendor before large orders [official](https://semiconductor.samsung.com/ssd/datacenter-ssd/).

---

## 29. NVMe 2.0 enterprise feature deep dive

- ZNS (Zoned Namespaces): host-managed sequential zones eliminate device-side garbage collection for zone-aligned writes — WAF approaches 1.0; supported on Kioxia CM7-class drives and Marvell Bravera-based designs [secondary](https://www.tweaktown.com/news/87587/kioxia-cm7-series-enterprise-pcie-5-0-ssds-up-to-14gb-sec-30tb/index.html).
- FDP (Flexible Data Placement): host hints place data by lifetime/temperature without full ZNS rework — WD SN861 supports FDP; a pragmatic middle ground [secondary](https://www.singular.com.cy/wd-ultrastar-dc-sn861-wus6ba119psp9x3-ssd-data-centre-192-tb-internal-25-u2-pcie-50-x4-nvme.html).
- Multistream writes: Kioxia CM7 groups data by expected lifetime at write time, reducing GC-induced WAF on mixed workloads [secondary](https://www.tweaktown.com/news/87587/kioxia-cm7-series-enterprise-pcie-5-0-ssds-up-to-14gb-sec-30tb/index.html).
- SR-IOV (16 PF / 32 VF on Marvell SC5): hardware-isolated virtual SSDs for multi-tenant bare metal and smart storage offload [vendor-reported](https://aem-origin-uat.marvell.com/content/dam/marvell/en/public-collateral/storage/marvell-ssd-mv-ss1331-1333-product-brief.pdf).
- CMB (Controller Memory Buffer): exposes a slice of drive DRAM to the host for submission queues — cuts submission latency for NVMe-oF initiators; Kioxia CM7 supports it [secondary](https://www.tweaktown.com/news/87587/kioxia-cm7-series-enterprise-pcie-5-0-ssds-up-to-14gb-sec-30tb/index.html).
- NVMe-MI 1.2c: out-of-band management (SMBus/MCTP) — datacenter fleet management baseline on Micron 9550, WD SN861 [official](https://www.mouser.com/datasheet/2/671/Micron_Technology_04_16_2025_9550_nvme_ssd_tech_pr-3581407.pdf).
- OCP 2.5 telemetry: standardized SMART/health/event reporting for hyperscale fleets — Samsung PM9D3a, Micron 9550 [official](https://semiconductor.samsung.com/ssd/datacenter-ssd/).
- TCG Ruby SSC: Kioxia CD8P offers Ruby alongside Opal — newer, simpler security spec aimed at datacenter SEDs [secondary](https://www.crn.in/news/kioxia-launches-new-pcie-5-0-ssds-for-enterprise-and-data-center-infrastructures/).
- 512 vs 4096-byte sectors: 4Kn improves ECC efficiency and is preferred for new deployments; 512e retained for legacy OS/hypervisor compatibility — all drives in this file support both [official](https://www.mouser.com/datasheet/2/671/Micron_Technology_04_16_2025_9550_nvme_ssd_tech_pr-3581407.pdf).

---

## 30. Warranty and support notes

- Standard enterprise warranty: 5 years across Samsung, Kioxia, Micron, Solidigm, WD, Phison lines in this file — endurance (TBW) and warranty are co-terminus: whichever exhausts first ends coverage [official](https://www.mouser.com/datasheet/2/671/Micron_Technology_04_16_2025_9550_nvme_ssd_tech_pr-3581407.pdf).
- Warranty is void-on-TBW-exhaustion: `percentage_used` = 100 ends warranty even inside 5 years — monitor fleet-wide, not per-drive [independent guidance].
- OEM-branded drives (Lenovo/Dell/HPE SKUs): warranty and firmware flow through the OEM, not the NAND vendor — factor OEM support contracts into TCO [secondary](https://www.shi.com/product/45962451/Samsung-PM1743-SSD).
- RMA practicalities: enterprise RMA usually requires the original vendor/OEM channel; keep purchase invoices; SED drives must be crypto-erased (PSID revert) before return — check the return policy on data-bearing media [independent guidance].
- MTBF 2.0–2.5M hours ≈ 228–285 years per drive — a population statistic: in a 10,000-drive fleet at 2.5M h MTBF, expect ~35 drive failures/year; provision spares accordingly [independent guidance].
- AFR framing: WD SN861 projects 0.35% AFR — at 10,000 drives ≈ 35 failures/year, consistent with the MTBF math above [secondary](https://cdn.multitronic.fi/media/c/d/mmo_133086493_1752741630_7228_401777.pdf).

---

## 31. One-line cheat sheet per drive family

