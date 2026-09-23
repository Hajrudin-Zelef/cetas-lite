---
id: etape9-phasea-enterprise-ssd/00-enterprise-ssd/28-gen4-vs-gen5-purchase-decision-2026
title: "28. Gen4 vs Gen5 purchase decision (2026)"
domain: step-9-enterprise-ssd-hardware-phase-a
role: deep-dive
task: hardware
actors: ["Samsung"]
dates: []
keywords: ["cost", "datacenter", "dram", "ipo", "latency", "memory", "nand", "throughput", "training"]
source: docs/RAG/etape9_phaseA_enterprise_ssd.md
source_anchor: ""
source_lines: [582, 648]
section: "Step 9 — Enterprise SSD Hardware (Phase A)"
sha256: 23749a7881df15853f8ee492a26a865c88ed2bcbe582a8fc67780ee5d37e7222
---

# 28. Gen4 vs Gen5 purchase decision (2026)

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

- Samsung PM9A3: Gen4 x4, V6 128L TLC, 960 GB–15.36 TB, 6.8/4.0 GB/s, 1M/180K IOPS, 1 DWPD — the used-market Gen4 staple [official](https://download.semiconductor.samsung.com/resources/brochure/Samsung%20PM9A3%20NVMe%20PCIe%20SSD.pdf).
- Samsung PM1743: Gen5 x4, 1.92–15.36 TB, 14.0/7.1 GB/s, 2.5M/360K IOPS, 1 DWPD, dual-port — first Gen5 enterprise SSD [secondary](https://www.shidirect.com/product/45832454/THINKSYSTEM-2.5IN-U.3-PM1743-15.36TB-READ-INTENSIVE-NVME-PCIE-5.0).
- Samsung PM9D3a: Gen5 x4, 960 GB–30.72 TB, 12.0/6.8 GB/s, 2M/400K IOPS, 1 DWPD, OCP 2.5 — Gen5 mainstream [official](https://semiconductor.samsung.com/ssd/datacenter-ssd/).
- Samsung BM1743: QLC high-density datacenter line (named on Samsung's page; specs not captured) [official](https://semiconductor.samsung.com/ssd/datacenter-ssd/).
- Kioxia CM7-R: Gen5 x4, 112L TLC, to 30.72 TB, 14.0/6.75 GB/s, 2.7M/310K IOPS, 1 DWPD, dual-port — performance flagship [secondary](https://www.directdial.com/us/item/kioxia-7-68tb-cm7-r-series-enterprise-2-5-nvme-ssd-solid-state-drive/kcmyxrug7t68).
- Kioxia CM7-V: Gen5 x4, 112L TLC, to 12.8 TB, 3 DWPD, dual-port — mixed-use flagship [secondary](https://www.Scan.co.uk/products/kioxia-64tb-cm7-v-u3-sie-pcie-gen5-1x4-2x2-u3-15mm-mix-use-3dwpd-enterprise-ssd).
- Kioxia CM9: teased, BiCS8 CBA, 3.4M read IOPS target — not shipping [vendor-reported](https://www.blocksandfiles.com/container-storage/2025/05/16/kioxia-teases-high-speed-ssd-aimed-at-ai-workloads/1601142).
- Kioxia CD8P-R: Gen5 x4, 112L TLC, to 30.72 TB, 12.0/5.5 GB/s, 2M/200K IOPS, 1 DWPD, single-port [secondary](https://www.digitec.ch/en/s1/product/kioxia-cd8p-r-series-kcd8xpug7t68-ssd-7680-gb-ssd-53141468).
- Kioxia CD8P-V: Gen5 x4, 112L TLC, to 12.8 TB, 2M/400K IOPS, 3 DWPD, single-port [secondary](https://www.galaxus.ch/en/s1/product/kioxia-x121-cd8p-v-dssd-u2-pcie-sie-3200-gb-25-ssd-49654389).
- Kioxia LC9: 245.8 TB QLC (announced; specs not captured) [secondary](https://www.blocksandfiles.com/ai-ml/2025/10/28/sk-hynix-aims-for-ai-flash-glory-with-ain-trifecta/1605493).
- Micron 9550 PRO: Gen5 x4, 232L TLC, 3.84–30.72 TB, 14.0/7.6 GB/s, 2.8M/400K IOPS, 1 DWPD, in-house stack [official](https://www.mouser.se/pdfDocs/9550-nvme-ssd-product-brief.pdf).
- Micron 9550 MAX: Gen5 x4, 232L TLC, 3.2–25.6 TB, 14.0/7.6 GB/s, 2.8M/750K IOPS, 3 DWPD [official](https://www.mouser.se/pdfDocs/9550-nvme-ssd-product-brief.pdf).
- Micron 7450 PRO: Gen4 x4, 176L TLC, to 15.36 TB, 6.8/5.6 GB/s, 1M/400K IOPS, 1 DWPD [official](https://www.micron.com/content/dam/micron/global/public/documents/products/technical-marketing-brief/7450-nvme-ssd-tech-prod-spec.pdf).
- Micron 7450 MAX: Gen4 x4, 176L TLC, to 12.8 TB, 3 DWPD [official](https://www.micron.com/content/dam/micron/global/public/documents/products/technical-marketing-brief/7450-nvme-ssd-tech-prod-spec.pdf).
- Solidigm D7-PS1010: Gen5 x4, 176L TLC, 1.92–15.36 TB, 14.5/9.3 GB/s, 3.1M/400K IOPS, 1 DWPD, UBER 1E-18 [secondary](https://techatlantix.com/blog/post/solidigm-d7-ps1010-review).
- Solidigm D7-PS1030: Gen5 x4, 176L TLC, 1.6–12.8 TB, 3 DWPD — mixed-use twin of PS1010 [secondary](https://www.6donline.com/solidigm-d7-ps1010-and-d7-ps1030-pcie-5-0-and-176l-tlc-datacenter-ssd-performance-play/).
- Solidigm D5-P5336: Gen4, 192L QLC, 61.44/122.88 TB — read-heavy capacity king [secondary](http://www.techtimes.com/articles/326561/20260903/wall-street-files-solidigm-etf-sk-hynix-faces-september-4-ipo-deadline.htm).
- Solidigm D7-P5810: SLC write-intensive line (specs not captured) [unverified].
- SK hynix PS1010: Gen5 datacenter SSD, 3.1M read IOPS class (prior gen) [secondary](https://www.blocksandfiles.com/container-storage/2025/05/16/kioxia-teases-high-speed-ssd-aimed-at-ai-workloads/1601142).
- SK hynix PEB110: Gen5 E1.S, 2/4/8 TB, 238L 4D NAND, OCP 2.5, SPDM — qualification pending at announcement [vendor-reported](https://en.prnasia.com/releases/global/sk-hynix-develops-peb110-e1-s-for-data-centers-460322.shtml).
- SK hynix PE8111: not found in retrieved sources [unverified].
- WD SN861 (1 DWPD): Gen5 x4, 1.92–7.68 TB, 13.7/7.5 GB/s, 3.3M/430K IOPS, FDP, OCP 2.0 [secondary](https://www.singular.com.cy/wd-ultrastar-dc-sn861-wus6ba138psp9x3-ssd-data-centre-384-tb-internal-25-u2-pcie-50-x4-nvme.html?sl=el).
- WD SN861 (3 DWPD): Gen5 x4, 3.2–12.8 TB, 3.3M/800K IOPS [secondary](https://www.convergetp.co.uk/wd-ultrastar-dc-sn861-wus6ca264psp9x1-ssd-data-centre-6-4-tb-u-2-pcie-5-0-x4-nvme-1tstosto-041687/).
- Phison Pascari X200E: Gen5 x4, 176L eTLC, 1.6–25.6 TB (30.72 TB family), 14.8/8.7 GB/s, 3.2M/930K IOPS, 1/3 DWPD [independent](https://www.techpowerup.com:443/review/phison-pascari-x200e/single-page.html).
- Phison Pascari X200Z: Gen5 x4 pSLC, 0.8–3.2 TB, 14.8/9.5 GB/s, 3.1M/950K IOPS, 60 DWPD — cache tier [secondary](https://www.thessdreview.com/our-reviews/enterprise/phison-pascari-x200z-gen5-800gb-1-6tb-enterprise-ssd-review-slc-gold-commands-a-lightning-fast-60-dwpd-data-center-ssd/).
- Sandisk UltraQLC SN670: 256 TB QLC (announced; specs not captured) [secondary](https://www.blocksandfiles.com/ai-ml/2025/10/28/sk-hynix-aims-for-ai-flash-glory-with-ain-trifecta/1605493).
- Marvell Bravera SC5: 8/16ch Gen5 controller IP, 14/9 GB/s, 2M/1M IOPS targets — the merchant silicon behind third-party Gen5 drives [vendor-reported](https://aem-origin-uat.marvell.com/content/dam/marvell/en/public-collateral/storage/marvell-ssd-mv-ss1331-1333-product-brief.pdf).

---

