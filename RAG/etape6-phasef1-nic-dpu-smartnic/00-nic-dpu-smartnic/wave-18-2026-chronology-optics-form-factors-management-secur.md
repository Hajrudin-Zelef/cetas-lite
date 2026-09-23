---
id: etape6-phasef1-nic-dpu-smartnic/00-nic-dpu-smartnic/wave-18-2026-chronology-optics-form-factors-management-secur
title: "Wave 18 — 2026 chronology, optics/form factors, management & security"
domain: phase-f1-nic-dpu-and-smartnic-nvidia-amd-pensando-intel-ipu-
role: deep-dive
task: pricing
actors: ["AMD", "Broadcom", "CoreWeave", "Intel", "Nvidia", "Oracle", "Qualcomm", "United States"]
dates: ["2026-09-15", "2026-09-22"]
keywords: ["optics", "acquisition", "amd", "ethernet", "gpus", "helios", "hyperscaler", "intel", "memory", "mi400", "nvidia", "pricing"]
source: docs/RAG/etape6_phaseF1_nic_dpu_smartnic.md
source_anchor: ""
source_lines: [654, 708]
section: "Phase F1 — NIC, DPU and SmartNIC (NVIDIA, AMD Pensando, Intel IPU, Marvell, offloads, pricing)"
sha256: d38c1596e2482f27ae8cfe11296c9dde4f37cb15a02e5b76df1ad2191252e726
---

# Wave 18 — 2026 chronology, optics/form factors, management & security

## Wave 18 — 2026 chronology, optics/form factors, management & security

### 18.1 2026 news chronology (NIC/DPU relevant)

- **Jan 2026**: CES 2026 — AMD details Helios AI rack (MI400 GPUs + Venice EPYC + Vulcano NICs) [secondary].
- **Q1 2026**: Qualcomm completes Alphawave Semi acquisition ($2.4B) ahead of schedule [secondary].
- **Jan 2026**: UEC Specification 1.0.2 released (CMS congestion-control corrections) [secondary].
- **2026**: AMD Pollara 400 begins shipping; billed as first UEC 1.0-compliant NIC; Oracle first CSP deployment [vendor-reported].
- **Mar 2026 (GTC, Washington DC)**: NVIDIA reveals BlueField-4 details — 800 Gb/s, 64 Neoverse V2 Grace cores @ 1.7 GHz, 128 GB LPDDR5, ConnectX-9, PCIe Gen6; CoreWeave/OCI/Palo Alto Networks named ecosystem partners; BlueField-5 (2028/Feynman) disclosed [vendor-reported/secondary].
- **Jun 2026**: VIAVI launches first Ultra Ethernet Transport validation solution [secondary].
- **Jul 2026**: UEC Specification 1.0.3 (current) — 200G/lane PHY, LLR/CBFC fixes [secondary].
- **H2 2026**: BlueField-4 early availability in Vera Rubin systems; AI-native storage platform partner availability [vendor-reported].
- **2026-09-15**: Compsource BF-3 B3210E listing last updated ($4,040.60, out of stock) — snapshot date for pricing waves [secondary].

### 18.2 Optics and connector form factors

| Speed | Common optics | Connector | Reach (typical) | Notes |
|---|---|---|---|---|
| 100G | QSFP28 SR4/DR | QSFP28 | 100 m–500 m | BF-2 E-series |
| 200G | QSFP56 SR4/DR4 | QSFP56 | 100 m–500 m | BF-3 B3220, CX-7 |
| 400G | QSFP112 DR4/SR4 | QSFP112 | 100 m–500 m | BF-3 B3240, Pollara, Thor 2 |
| 400G | OSFP DR4 | OSFP | 500 m | CX-8 800G port configs |
| 800G | OSFP DR8/SR8 | OSFP | 100 m–500 m | CX-8, Thor Ultra, BF-4 |
| NDR IB 400G | OSFP | OSFP | 30–100 m | CX-7 IB mode |

- QSFP112: 4× 100G PAM4 lanes; OSFP: 8× 100G (800G) with better thermals; DAC passive copper reach: Thor 2 up to 5 m, BF-class DACs 1–3 m typical; AECs extend to 5–7 m [secondary].
- BlueField-3 cards use QSFP112 cages (dual-port P-series, single-port B3140L); ConnectX-8 offers OSFP or dual-QSFP112 variants [official].

### 18.3 DPU management and security architecture

- BlueField-3 carries an **integrated BMC** (separate from host BMC) enabling out-of-band DPU management; firmware update via **PLDM type-5 over MCTP over PCIe** through the platform BMC (DOCA 2.9.x; ~100 MB image; full power cycle required in current flow) [official].
- BlueField secure boot: PKA root-of-trust, flash encryption, MACsec/IPsec/TLS offload engines; crypto-enabled vs crypto-disabled SKUs for export-control domains [official].
- **DPU isolation model**: DPU runs its own OS; host cannot snoop DPU memory (hardware isolation); DPU can enforce policy on host traffic (firewall, micro-segmentation) — basis for "DPU as root of trust" and zero-trust host networking [vendor-reported].
- DOCA Management Service (DMS) GA in 2.9: lifecycle management of DPU fleet (provisioning, telemetry, updates) [official].
- Operational note: DPU fleet management is a distinct skill set (DPU OS upgrades, BFB image provisioning, OOB network for DPU management ports) — plan for DPU-specific runbooks separate from host runbooks [independent].

### 18.4 Form-factor checklist for procurement

- PCIe slot: Gen5 x16 electrical for BF-3/Pollara/Thor 2; Gen6 x16 for 800G generation (CX-8, Thor Ultra, Vulcano).
- Mechanical: FHHL single-slot (most NICs), FHHL dual-slot (B3240 P-series 400G), HHHL (some CX variants), OCP 3.0 SFF (hyperscaler/server-vendor specific).
- Power: slot 75 W limit; BF-3 B3220/B3240 need **auxiliary power cable kits** (Lenovo/HPE part numbers documented in Wave 10); 800G cards 100–150 W class — verify PSU and airflow.
- Thermal: 400G+ optics run hot; OSFP has better thermal headroom than QSFP112 at 800G; data-center ambient specs matter for DAC vs optics choice.

### 18.5 Quick reference — key model numbers (verbatim)

- NVIDIA: 900-9D3B6-00CN-AB0 (BF-3 B3240), 900-9D3B6-00CV-AAH / 00CV-AA0 (BF-3 B3220), 900-9D3B4-00CC-EA0 (BF-3 B3210L), 900-9D3B4-00EN-EA0 (BF-3 B3140L), 900-9D3D4-00EN-HA0 (B3140H SuperNIC), MBF2M516A-CECOT / -CENOT / -EECOT (BF-2), MCX823105AN-ADAT (CX-8 800G OSFP, FS.com US listing).
- AMD: Pensando Salina 400, Pollara 400, Vulcano 800 (800G, 2027).
- Intel: E2000 (Mount Evans), Oak Springs Canyon C6000X, F2000X-PL [unverified].
- Broadcom: BCM957608-P1400GDF00 (Thor 2 400G), BCM957508-P2200G/P2100G (NetXtreme), PS410T-H04 (Stingray 4×10G-T), Thor Ultra 800G.
- Marvell: OCTEON 10 CN102/CN103, CN106, DPU400 (36-core), OCTEON 12 [unverified].
- Lenovo: 4XC7A93809 / 4XC7A87752 / 4XC7B08942 / 4XC7A96568; HPE P71949-B21 (BF-3 power cable kit).

---
*End of Phase F1. 18 waves, append-only. Research cutoff 2026-09-22. Single writer; no other workspace files modified.*

