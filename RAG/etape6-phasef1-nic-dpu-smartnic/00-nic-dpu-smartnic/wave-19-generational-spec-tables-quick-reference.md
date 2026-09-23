---
id: etape6-phasef1-nic-dpu-smartnic/00-nic-dpu-smartnic/wave-19-generational-spec-tables-quick-reference
title: "Wave 19 — Generational spec tables (quick-reference)"
domain: phase-f1-nic-dpu-and-smartnic-nvidia-amd-pensando-intel-ipu-
role: deep-dive
task: pricing
actors: ["Nvidia"]
dates: ["2026-09-22"]
keywords: ["gpu", "memory", "nvidia", "optics", "research", "rubin", "throughput"]
source: docs/RAG/etape6_phaseF1_nic_dpu_smartnic.md
source_anchor: ""
source_lines: [709, 756]
section: "Phase F1 — NIC, DPU and SmartNIC (NVIDIA, AMD Pensando, Intel IPU, Marvell, offloads, pricing)"
sha256: a9af30e82fae10f5073f8741cf9ea11d334ce9e819c82ab26ddb56d0eb6889cb
---

# Wave 19 — Generational spec tables (quick-reference)

## Wave 19 — Generational spec tables (quick-reference)

### 19.1 NVIDIA BlueField generational table

| Attribute | BlueField-2 | BlueField-3 | BlueField-4 |
|---|---|---|---|
| Launch | 2020 | 2022 | 2026 (early availability) |
| Max network | 200G | 400G / NDR | 800G |
| Arm cores | 8× A72 | 16× A78 | 64× Neoverse V2 (Grace) @ 1.7 GHz |
| Memory | 16 GB DDR | 32 GB DDR (P-series) / 16 GB (E-series) | 128 GB LPDDR5 (275 GB/s reported) |
| PCIe | Gen4 x16 | Gen5 x16 (+x16 extension) | Gen6 x16 |
| Network chip | ConnectX-6 Dx | ConnectX-7-class | ConnectX-9 |
| BMC | No (host BMC) | Integrated | Integrated |
| Host driver | MLNX_OFED | DOCA-Host (MLNX_OFED unsupported) | DOCA-Host |
| Crypto accel | BF-2 gen | 4× vs BF-2 | TBA |
| Example street price (2026-09) | $2,299 new / $380 used | $3,053–$4,686 new | Not found [gap] |

### 19.2 NVIDIA ConnectX generational table

| Generation | Max speed | PCIe | Notes |
|---|---|---|---|
| ConnectX-6 Dx | 200G | Gen4 | OVS/ASAP² mature; BlueField-2 base |
| ConnectX-7 | 400G / NDR | Gen5 | BlueField-3 base; Multi-Host |
| ConnectX-8 | 800G | Gen6 | SuperNIC; OSFP or dual QSFP112; $1,779–$2,519 street |
| ConnectX-9 | 1.6T per GPU (multi-chip) | Gen6/7 | Rubin generation; paired with BlueField-4 |

### 19.3 DPU power envelope (reported/observed)

- OCTEON 10: ~60 W (DPU400 36-core) down to sub-25 W (CN103) [vendor-reported].
- BlueField-3: dual-slot 400G cards need auxiliary power beyond 75 W slot (Lenovo/HPE cable kits) [official].
- Thor 2 400G: marketed on low power / air-cooling density vs competitors [vendor-reported].

---
*End of Phase F1. 19 waves, append-only. Research cutoff 2026-09-22. Single writer; no other workspace files modified.*

## Wave 20 — Pre-deployment validation checklist (field notes)

- [ ] Confirm PSID against NVIDIA firmware compatibility list before ordering optics/cables [official].
- [ ] Pin DOCA version to a validated release (e.g. DOCA 2.9.x LTS) and test BFB upgrade path [official].
- [ ] Verify PCIe bifurcation, slot power (75 W limit) and aux-power kit fit in target chassis [official].
- [ ] Burn-in test: line-rate packet gen (TRex/DPDK pktgen) through OVS-offload and NVMe-oF paths [independent].
- [ ] Validate RoCEv2 ECN/PFC end-to-end with DCQCN counters; or UEC NSCC/RCCC if Pollara/Vulcano [independent].
- [ ] Measure TLS/IPsec throughput vs CPU-offload baseline to confirm crypto engine engagement [independent].
- [ ] Register DPU BMC in OOB management VLAN; rotate default BMC credentials [official].
- [ ] Document grey-market exclusions: warranty only via authorized channel/OEM [secondary].

---
*End of Phase F1. 20 waves, append-only. Research cutoff 2026-09-22. Single writer; no other workspace files modified.*
