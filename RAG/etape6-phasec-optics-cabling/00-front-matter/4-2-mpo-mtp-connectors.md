---
id: etape6-phasec-optics-cabling/00-front-matter/4-2-mpo-mtp-connectors
title: "4.2 MPO/MTP connectors"
domain: front-matter
role: reference
task: reference
actors: ["Nvidia", "United States"]
dates: []
keywords: ["alignment", "nvidia", "optics"]
source: docs/RAG/etape6_phaseC_optics_cabling.md
source_anchor: ""
source_lines: [564, 622]
section: "Step 6 — Phase C: Optics, Cabling & Interconnect Infrastructure"
sha256: 24fcfe356b2618cdec1fbea9b7c7f9ee190489436a24a655393843c6fef3d23c
---

# 4.2 MPO/MTP connectors

### 4.2 MPO/MTP connectors

#### 4.2.1 Fiber counts & applications

| Connector | Fibers/rows | Applications |
|---|---|---|
| **MPO-8** | 8 fibers (1×8; or 8 used positions of 12-f ferrule) | 40G-SR4, 100G-PSM4, 100G-SR4/400G-SR4/400G-DR4 (4 Tx + 4 Rx, 100% utilization) [vendor-reported — FS.com https://www.fs.com/blog/what-are-mtpmpo8-cables-15692.html] |
| **MPO-12** | 12 fibers (1×12) | 40G-SR4/100G-SR4 (8 of 12 active), 400G-SR4/400G-DR4 (8 of 12 active); legacy 100G-SR10; 2×400G via dual MPO-12 (800G 2xSR4/2xDR4) [vendor-reported — holightoptic.com https://www.holightoptic.com/mpo-fiber-patch-cords-how-to-choose-the-right-one-for-your-high-density-network/; medium/@aicplight888] |
| **MPO-16** | 16 fibers (1×16, MT-16 ferrule, same external footprint as MPO-12) | 400G-SR8 (8×50G), 400G-DR8, 800G-SR8/800G-DR8 (8 Tx + 8 Rx, no waste); Terabit BiDi MSA 800G-SR8.2 (16 fibers, 2 λ) [vendor-reported — FS.com; CommScope ordering guide https://webresources.commscope.com/download/assets/CO-120443-EN_NVIDIA+NVL72+DC+Ordering+Guide.pdf/e6cd640a10b611f1944a3a47d01923dd] |
| **MPO-24** | 24 fibers (2×12) | Trunk backbones, breakouts (MPO-24 → 12×LC, MPO-24 → 3×MPO-8); 100GBASE-SR10 (CFP); hyperscale spine/leaf aggregation [vendor-reported — FS.com; cablinginstall.com https://www.cablinginstall.com/connectivity/product/55362193/800g-breakout-module-with-front-facing-mpo-port] |

- MTP = US Conec's enhanced MPO (1996, 100% interoperable), adds metal pin clamp, elliptical pins, removable housing; MTP-16 = 16-fiber MT ferrule in same footprint [vendor-reported — US Conec catalog https://www.usconec-cn.com/media/k2pbdmyl/us-conec-product-catalog.pdf; FS.com timeline https://www.fs.com/blog/the-development-of-mtp-cables-12629.html]
- **Corrections/flags:** (a) IEEE lists 400GBASE-SR4.2 (400G-BiDi) as a **typical 12-fiber MPO** (8 fibers used), NOT MPO-24 [official — IEEE 802.3 NGMMF https://ieee802.org/3/NGMMF/public/Jan18/lingle_ngmmf_03_jan18.pdf; Arista FAQ http://sit.www.arista.com/assets/data/pdf/Datasheets/Arista-400G_Optics_FAQ.pdf]. Task item "MPO-24 for 200G/400G-SR4.2" is not supported — MPO-24 is used for trunk/backbone and breakout-to-3×MPO-8, not the SR4.2 optic itself. (b) MPO-8 is "not strictly a standards-recognized interface" — it uses 8 positions of a 12-f ferrule [vendor-reported — FS.com]. (c) 400G-SR8 per IEEE 802.3cm = 400GBASE-SR8 with 16 fibers [official — IEEE].

#### 4.2.2 Male vs female (pinned/unpinned)
- Male (pinned): 2 alignment pins protrude from MT ferrule. Female (unpinned): alignment holes, no pins. Mated pairs must be one pinned + one unpinned [vendor-reported — FS.com https://www.fs.com/uk/blog/a-comprehensive-guide-to-mtp-connector-13324.html]
- Convention: **equipment/transceiver ports are typically unpinned (female)**; trunk cables are often pinned (male); cassettes/modules present unpinned rear ports [secondary — network-switch.com https://network-switch.com/blogs/networking/mpo-connectors-explained-2025]
- Mating male↔male or female↔female = mis-mate; male↔male can damage pins/ferrules [vendor-reported — holightoptic.com]

#### 4.2.3 APC vs UPC polish
- **UPC** (Ultra Physical Contact, slightly domed 0° polish): standard for **multimode** MPO systems; return loss ≥20 dB (MPO MM) [vendor-reported — FS.com]
- **APC** (Angled Physical Contact, 8° angle polish): standard for **single-mode** MPO/MTP parallel optics (e.g., 400G-DR4, 800G-DR8, NVIDIA InfiniBand); return loss ≥60 dB (SM APC) [vendor-reported — FS.com; network-switch.com]
- Never mate APC↔UPC (8° angle vs flat causes high loss/damage) [secondary — industry rule]
- Example: FS OS2 harnesses default to "Female APC 0.35dB" on the MPO side and "LC UPC Duplex" on breakout legs [vendor-reported — FS.com product pages]

#### 4.2.4 Insertion loss (IL) specs — standard vs low-loss/elite

| Grade | SM (APC/UPC) | MM |
|---|---|---|
| **Standard MTP** | typ 0.25 dB, max 0.75 dB (per fiber) | typ 0.20 dB, max 0.60 dB |
| **MTP Elite (low-loss)** | typ 0.10 dB, max 0.35 dB | typ 0.10 dB, max 0.35 dB |

[vendor-reported — FS.com (citing US Conec MTP/Elite figures) https://www.fs.com/uk/blog/a-comprehensive-guide-to-mtp-connector-13324.html; consistent with ecer.com factory spec table ≤0.35/≤0.75 SM, ≤0.35/≤0.60 MM https://www.ecer.com/corp/detail500-uuu46gt-pu8hd19-mpo-mtp-breakout-fiber-optic-cable-armored-fiber-optic-patch-cable-om3-om4.html]
- Optical return loss: ≥60 dB SM (8° APC), ≥20 dB MM UPC; durability ≥500 mating cycles; operating −40 to +80 °C [vendor-reported — ecer.com; holightoptic.com lists SM ≥60 dB, MM ≥30 dB — **minor conflict flagged**: MM RL cited as ≥20 dB (FS) vs ≥30 dB (Holight); both vendor-claimed]
- FS.com sells all trunk/harness/cassette lines at "0.35 dB max" low-loss as standard [vendor-reported — FS.com product pages]

### 4.3 Polarity methods

#### 4.3.1 The three methods (TIA-568 / TIA-568.3-D; "Type A/B/C")

| Method | Key orientation | Fiber mapping (MPO-12) | Patch cords in a full link |
|---|---|---|---|
| **Method A (straight-through)** | Key up → Key down | 1→1, 12→12 (straight) | A-to-B at one end + A-to-A at the other end; polarity fixed in cassettes |
| **Method B (reversed)** | Key up → Key up | 1→12, 2→11 … (full reversal) | A-to-B patch cords both ends; preferred for parallel optics (40G/100G-SR4, 400G/800G) |
| **Method C (pair-flipped)** | Key up → Key down | 1↔2, 3↔4 … (adjacent pairs crossed inside trunk) | A-to-B duplex cords; **duplex breakout only** |

[secondary — TIA-568-C.0/-3-D described via cablinginstall.com https://www.cablinginstall.com/connectivity/article/16465981/maintaining-fiber-optic-polarity-with-array-cabling; FS.com https://www.fs.com/blog/polarity-and-mtp-technology-in-40100g-transmission-5211.html; holightoptic.com https://www.holightoptic.com/understand-mtp-mpo-polarity-fiber-cable/]
- **Recommendation:** Method B for MPO-12 parallel-optics links (keeps patching consistent at both ends, aligns Tx/Rx of QSFP/OSFP transceivers) [secondary — holightoptic.com]
- Method C is not recommended if migration to parallel optics is planned [secondary]
- TIA-568 formally defines polarity for **single-row 12-fiber** MPO only; MPO-24 (2-row) mappings are vendor-specific — always request the manufacturer's polarity diagram (common procurement pitfall) [secondary — medium.com @Yanissss, Sept 2026 https://medium.com/@Yanissss/mpo-mtp-polarity-types-explained-type-a-b-c-52ff209ac7ed]
- Adapters: Type A adapter = opposed-key (key-up/key-down); Type B adapter = aligned-key (key-up/key-up) [secondary — cablinginstall.com]

#### 4.3.2 Common mistakes
1. Mixing A/B/C within one channel ("golden rule: pick one method site-wide") [secondary]
2. Wrong gender (male↔male / female↔female) — breaks alignment, risks ferrule damage [secondary]
3. Using Method C trunks for parallel-optics migration [secondary]
4. Assuming 24F/16F polarity works like 12F without the vendor map (lane-order errors) [secondary]
5. Skipping inspection/cleaning — field loss is usually a contamination problem, not a polarity one [vendor-reported — holightoptic.com]

