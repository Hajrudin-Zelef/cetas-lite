---
id: etape6-phasec-optics-cabling/00-front-matter/7-4-fiber-cabling-os2-om4-patch-cords-mtp-mpo-trunks-cassett
title: "7.4 Fiber cabling (OS2/OM4 patch cords, MTP/MPO trunks, cassettes)"
domain: front-matter
role: reference
task: reference
actors: ["Broadcom", "EU", "Huawei", "Intel", "Nvidia", "United States"]
dates: ["2025-12-22", "2026-04"]
keywords: ["3nm", "dsp", "ethernet", "intel", "latency", "lpo", "nvidia", "optics", "pricing", "wavelength"]
source: docs/RAG/etape6_phaseC_optics_cabling.md
source_anchor: ""
source_lines: [1214, 1243]
section: "Step 6 — Phase C: Optics, Cabling & Interconnect Infrastructure"
sha256: 671194f15e93e448b042926c64c060140a186be1ac88f5059c44dd7acc51fdf8
---

# 7.4 Fiber cabling (OS2/OM4 patch cords, MTP/MPO trunks, cassettes)

- **25G/10G/1G DAC/AOC** [official, fs.com category page crawled ~102 days ago]: 10G SFP+ passive DAC 1m US$14.00 (125.5K sold); 10G SFP+ AOC 1m US$33.00 (12K sold); 10G SFP+ active DAC 1m US$58.00; 10G SFP+ industrial AOC 1m US$74.00; 25G SFP28 passive DAC 1m US$33.00 (23.4K sold), 0.5m US$33.00; 25G SFP28 active DAC 1m US$40.00; 25G SFP28 AOC 1m US$65.00 (3.6K sold); 25G SFP28 industrial AOC 1m US$128.00; 1G SFP passive DAC 0.5m US$14.00. AWG/bend: 10G/25G passive DAC = 30AWG twinax; 25G SFP28 passive DAC 21–21.25mm diameter; active DACs ≤0.1–0.5W; AOCs ≤0.8W [official].
- **100G DAC/AOC** [official]: 100G QSFP28 passive DAC (implied 1m, 30AWG) US$136.00 (210 sold); 100G QSFP28 active DAC 1m US$156.00 (196 sold); 100G QSFP28→4x 25G SFP28 active breakout DAC 3m US$175.00; active breakout DAC 2m US$261.00; 100G QSFP28→2x 50G QSFP28 breakout AOC 30m US$405.00; 100G QSFP28→4x LC breakout AOC 15m US$281.00; 100G SFP-DD 2x 50G PAM4 AOC 5m US$409.00 (22 sold); 100G SFP112 PAM4 passive DAC 1.5m US$74.00 (28AWG, ≤0.15W). Huawei-specific "HW AOC-Q28-S28-7M 100G QSFP28 to 4x25G SFP28 breakout AOC" (P/N QSFP-100G-4SAO07, SKU 70530) US$264.00 [official].
- **200G DAC/AOC**: 200G QSFP-DD→8x 25G SFP28 passive breakout DAC 3m (Dell compatible, P/N QDD-200G-8SPC03, SKU 153055) €392.70 / €330.00 VAT excl. (EU) and £372.00 / £310.00 VAT excl. (UK) [official]; 200G QSFP56 passive DAC 2m SGD 149.33 GST incl. (above) [official]. 200G AOC straight-through pricing not captured [unverified gap].
- **50G/56G DAC/AOC**: category exists ("56G/50G DAC/AOC" in DAC category tree); no specific SKU/price surfaced this wave [unverified gap].
- **800G AOC**: see section 7.2 (only 800G OSFP AOC 1m found; no QSFP-DD 800G AOC or longer-length pricing captured) [official, partial].

### 7.4 Fiber cabling (OS2/OM4 patch cords, MTP/MPO trunks, cassettes)

**MTP/MPO trunks** [official]: 144F MTP-12 OM4 UPC trunk 5m US$1,603.00; 144F MTP-12 OS2 APC 5m US$1,524.00; 96F MTP-8 OM4 UPC 5m US$1,500.00; 96F MTP-8 OS2 APC 5m US$1,389.00; 48F MTP-12 OM4 UPC 30m Type A US$1,222.00; 48F MTP-12 OS2 APC 30m Type A US$769.00; 48F MTP-12 OM3 UPC 30m US$906.00. All US Conec MTP®, OFNP plenum, low IL (0.20dB MMF / 0.35dB SMF) [official]. Custom trunks: OM4 MTP-12 (P/N XXMTPOM4, SKU 30962) from US$85.08 (8–288 fibers, lengths 1–30m, polarity/gender/jacket selectable); OS2 MTP-12 (P/N XXMTPSMF, SKU 30976) from US$105.08 (8–144 fibers); MTP-8 hybrid OS2→LC (P/N HD-XXMTPLCSMF, SKU 209711) from US$105.25; 12F MTP-12→6x LC OS2 hybrid 30m (P/N HD-12FMTPLCSMF, SKU 209731) €177.31 / €149.00 VAT excl. (EU) [official].
**MTP breakouts/conversions** [official]: MTP-24→3x MTP-8 OM4 UPC 1m US$207.00; MTP-24→3x MTP-8 OS2 APC 1m US$307.00; 2x MTP-12→3x MTP-8 OM4 3m US$274.00; MTP-12→2x MTP-4 OM4 (NVIDIA IB) 1m US$122.00; MTP-16→2x MTP-8 OM4 crossover 1m US$150.00 (340 sold); MTP-16→2x MTP-8 OS2 crossover 1m US$249.00.
**Patch cords** [official]: FS official examples: OS2 LC-LC 0.5m OFNR (P/N SMLCDX, SKU 88527) SGD 6.54 GST incl. [official, SG]; armored OS2 LC-LC 1m US$10.00 (P/N AM-OS2LCSCDX, SKU 182476); industrial armored OS2 LC-LC 30m US$102.00 (P/N MG-OS2LCDX, SKU 106591); OM4 LC-LC armored 3m US$16.00 (SKU 41028) / 5m US$23.00 (P/N AM-OM4LCDX, SKU 41027); FTTA OS2 LC-LC 5m US$22.00 (P/N FTTA-OS2LCDX, SKU 97938) [official]. Third-party reseller listing of FS SKUs (for reference, not official): standard duplex LC-LC OS2 2m ≈US$5.70 (SKU 40192) and OM4 2m ≈US$6.55 (SKU 40220) [secondary].
**Custom bulk fiber**: OS2 indoor 4F from US$15.00; OM4 indoor 4F from US$17.00; OM5 4F from US$4.00 [official]. **Cassettes**: no FS cassette/enclosure SKU or price surfaced this wave [unverified gap — category exists under fiber connectivity but uncaptured].

### 7.5 Coding/programming for OEM compatibility

- FS claims transceivers "use the same software codes as the original vendors" and modules are MSA-produced with vendor ID rewritten in EEPROM [official, fs.com blog "Compatibility Guarantee for FS Transceivers", ~5 years old].
- **FS Box / fsbox.com**: in-house real-time coding device with interfaces for SFP/SFP+, XFP, QSFP+/QSFP28 modules and DAC/AOC cables; "module configurations of more than 200 brands like Cisco, Juniper, Arista and other customized reprogramming services"; batch coding, diagnostic coding via web interface [official].
- Named compatible brands: Cisco, HP/HPE, Juniper, Brocade, Dell/Dell EMC, Extreme, H3C, Arista, Huawei, Intel, IBM, Netgear, Ciena, D-Link, Avago, F5 Networks, Avaya, Alcatel-Lucent, Aruba, Allied Telesis, SMC, TRENDnet, Palo Alto, Ciena, Edge-Core, Mellanox/NVIDIA [official, fs.com SFP+/QSFP28 family pages].
- **800G-specific testing claims** [official, fs.com blog]: every 800G transceiver comprehensively tested (power, wavelength, traffic, optical performance) with test report; "tested on original devices to ensure 100% compatibility"; InfiniBand optics verified on NVIDIA Quantum-2 switches + ConnectX-7 adapters; Ethernet optics claimed compatible with NVIDIA (ETH), Arista, Juniper, Cisco. AU site lists validated device matrix: NVIDIA MQM9700/9790, Quantum-X800, Spectrum-4/3 SN5600/5400/4700/4600, ConnectX MCX75510AAS/MCX75310AAS/MCX653105A/MCX653106A, BlueField-3; Cisco N9364E-SG2-Q/O, 9364D-GX2A, 9348D-GX2A; Juniper QFX5240-64OD/QD, QFX5230-64CD, QFX5220-32CD; Arista 7060X5 series; Dell Z9664F/Z9432F/Z9332F/Z9264F-ON [official].
- 800G product P/Ns echo OEM naming (e.g., MMA4Z00-NS-FLT, MMS4X00-NS, QDD-2X400G-DR4, QDD-800G-2XDR4, OSFP-800G-2FR4) to signal drop-in compatibility with NVIDIA/Arista/Cisco/Juniper gear [official]. Product pages are headed "Xxx Compatible" (e.g., "Cisco QDD-400-AOC1M Compatible") [official]. All claims are vendor assertions, not independent verification [vendor-reported].

### 7.6 2026 new product announcements

- **No 2026-dated FS launch announcement found.** The 1.6T announcement widely reported as "FS Launches 1.6T OSFP IHS/Closed Finned Top Modules" is dated **2025-12-22** per Business Wire URL (…/20251222728263) and republication slugs "bizwire-2025-12-22" [independent]. Details [vendor-reported via Business Wire]: models OSFP-DR8-1.6T (SiPh, 1310nm, dual MTP/MPO-12 APC, 500m, 25W, NVIDIA MMS4A00/980-9IAH1-00XM00) and OSFP-2FR4-1.6T (EML, 4λ 1271/1291/1311/1331nm, dual duplex LC, 2km, 26W, NVIDIA MMS4A50-XM/980-9IAS0-00XM00); twin-port OSFP closed finned top, 8x 200G PAM4, Broadcom 3nm DSP, targeted at NVIDIA GB300/B300 InfiniBand XDR; claimed 100% verified on NVIDIA Quantum-X800 switches; IHS compatible with air- and liquid-cooled environments.
- An April 2026 third-party PDF references an FS 800G LPO launch ("FS Launches 800G LPO Module: A Power Efficiency and Latency… 800G DR8 OSFP finned top LPO module… ultralow power consumption, reduced latency" for AI/ML) [secondary, undated launch; the FS 800G LPO OSFP DR8 SKU is live at US$1,749.00 per section 7.1].
- **1.6T commercially available now**: FS site navigation lists "1.6T OSFP" as a live category [official]; 1.6T OSFP224 passive DAC 1m listed SGD 777.17 GST incl. (67 sold) [official]. No 1.6T transceiver list price captured [unverified gap].
- DAC/AOC/AEC family expanded to 1.6T (category tree shows "1.6T/800G DAC/AOC/AEC", "400G DAC/AOC/AEC", "200G DAC/AOC", "100G DAC/AOC", "56G/50G DAC/AOC", "40G DAC/AOC", "25G/10G/1G DAC/AOC") [official].

### 7.7 Shipping/stock/availability and warranty

