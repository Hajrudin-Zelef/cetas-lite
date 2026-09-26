---
id: etape6-phasef1-nic-dpu-smartnic/00-nic-dpu-smartnic/part-5
title: "Phase F1 — NIC, DPU and SmartNIC (NVIDIA, AMD Pensando, Intel IPU, Marvell, offloads, pricing) (part 5)"
domain: phase-f1-nic-dpu-and-smartnic-nvidia-amd-pensando-intel-ipu-
role: deep-dive
task: pricing
actors: ["China", "Lambda", "United States"]
dates: ["2026-09-15"]
keywords: ["nvidia", "pricing", "ethernet"]
source: docs/RAG/etape6_phaseF1_nic_dpu_smartnic.md
source_anchor: ""
source_lines: [169, 203]
section: "Phase F1 — NIC, DPU and SmartNIC (NVIDIA, AMD Pensando, Intel IPU, Marvell, offloads, pricing)"
sha256: 7fb4fa6ea67044d558bce5a982d8f39e1f7d0f43ffafa5d41b87747abd955204
---

# Phase F1 — NIC, DPU and SmartNIC (NVIDIA, AMD Pensando, Intel IPU, Marvell, offloads, pricing) (part 5)

- FS.com US [secondary] [source: https://www.FS.com/c/nvidia-ethernet-nics-4014 ]:
  - ConnectX-8 VPI 800G OSFP 1-port (PCIe 6.0, Secure Boot, Crypto, IB+Ethernet RoCE) — **$2,519.00** ("New").
  - ConnectX-8 VPI 400G QSFP112 2-port (PCIe 6.0) — **$2,519.00**.
  - ConnectX-7 VPI 400G OSFP 1-port (PCIe 5.0) — **$2,188.00**.
  - ConnectX-7 VPI 200G QSFP112 2-port (PCIe 5.0) — **$2,129.00** ("Hot").
  - ConnectX-6 Dx 100G QSFP56 2-port (PCIe 4.0, RoCE) — **$1,550.00**.
  - ConnectX-6 Dx 100G OCP 3.0 2-port — **$1,649.00** ("Hot").
  - ConnectX-5 100G QSFP28 2-port (PCIe 3.0) — **$736.00** ("Hot").
  - ConnectX-5 100G QSFP28 1-port — **$429.00**.
  - ConnectX-6 Lx 25G SFP28 2-port (PCIe 4.0) — **$469.00**.
- FS.com UK [secondary] [source: https://www.fs.com/uk/c/nvidia-ethernet-nics-4014 ]:
  - ConnectX-8 800G OSFP 1-port — **£2,389.20** (£1,991.00 VAT excl.), 55 sold.
  - ConnectX-8 400G QSFP112 2-port — **£2,450.40** (£2,042.00 VAT excl.), 61 sold.
  - ConnectX-7 400G OSFP 1-port — £2,227.20 (£1,856.00), 483 sold.
  - ConnectX-7 200G QSFP112 2-port — £2,227.20 (£1,856.00), 297 sold.
  - ConnectX-6 Dx 100G 2-port — £1,426.80 (£1,189.00), **9.7K sold**, 17 reviews — highest-volume SKU.
  - ConnectX-6 Lx 25G 2-port — £435.60 (£363.00), 1.4K sold.
- NADDOD wholesale [secondary] [source: https://www.naddod.com/collections/nvidia-networking/infiniband-adapters ]:
  - ConnectX-7 MCX75310AAS-NEAT (NDR/400GbE) — **$1,599.00** (1.2k+ in stock, 12.1k+ sold).
  - ConnectX-7 MCX755106AS-HEAT (NDR200 dual QSFP112) — **$1,599.00** (1.3k+ in stock, 3k+ sold).
  - ConnectX-8 C8240 (900-9X81Q-00CN-ST0, dual QSFP112) — **$1,779.00** (110 in stock, 375 sold).
  - ConnectX-8 C8180 (900-9X81E-00EX-DT0, single OSFP XDR 800G) — **$1,779.00** (98 in stock, 318 sold).
  - C8180X PCIe auxiliary kit — **$239.00**.
  - ConnectX-6 VPI MCX653105A-ECAT (HDR100/EDR/100GbE) — **$1,065.00** (90 sold).
- Other channels: ecer.com ConnectX-8 SuperNIC listing **$2,500/pc** MOQ 1 (China origin, third-party "Comelink" brand — likely grey-market, treat as [unverified]); Wamatek (Germany) C8180 at **$1,903.99** (was $2,646.78) [secondary].
- BlueField-3 street pricing:
  - Used B3240 (900-9D3B6-00CN-AB0, dual 400G QSFP112) on eBay — **$3,999.00** used [secondary] [source: https://www.ebay.com/itm/306620041473 ].
  - SHI: B3210E (900-9D3B6-00CC-EA0, 100G E-series) — **$3,053.00** (MSRP $4,686.00); public-sector SHI listing shows $4,425.00 (MSRP $6,794.00) — channel price variance is significant [secondary].
  - Newegg Q&A page: B3210E at **$2,628.99**, out of stock [secondary].
  - Compsource: B3210E **$4,040.60** new, out of stock; last updated 2026-09-15 [secondary].
  - LambdaTek UK: B3220 P-series (900-9D3B6-00CV-AA0, 200GbE) — **£3,009.90** (£3,611.88 inc VAT) [secondary].
  - PC-Canada: B3220 **900-9D3B6-00SV-AA0-DELL** — volume pricing $6,981.70–$7,228.99 (MSRP $11,579.78), out of stock [secondary] — Dell OEM premium visible.
- BlueField-2 (previous gen) street: $2,043–$2,299 new (SHI), $379.99 used (eBay) — see Wave 1.
- Pattern: 400G/800G DPUs and SuperNICs price in the **$1,800–$4,500** band retail; OEM/MSRP is typically 30–60% higher; used market collapses 70–85% [analysis].

