---
id: etape6-phasef4-virtualization-io/00-virtualization-io/8-2-nvidia-mig-profiles-the-slicing-table
title: "8.2 NVIDIA MIG profiles (the slicing table)"
domain: phase-f4-i-o-virtualization-cpu-acceleration-extensions
role: deep-dive
task: actor-profile
actors: ["Nvidia"]
dates: []
keywords: ["nvidia", "blackwell", "compute", "datacenter", "gpu", "memory"]
source: docs/RAG/etape6_phaseF4_virtualization_io.md
source_anchor: ""
source_lines: [476, 491]
section: "Phase F4 — I/O Virtualization & CPU Acceleration Extensions"
sha256: 9e6eb029d3bf467ded6ef72188fa9954a2b3751bc59d69a12c3a763c72ba09c3
---

# 8.2 NVIDIA MIG profiles (the slicing table)

### 8.2 NVIDIA MIG profiles (the slicing table)
- A100 40GB: 1g.5gb / 1g.5gb+me / 2g.10gb / 3g.20gb / 4g.20gb / 7g.40gb — max 7 instances; `+me` adds one NVDEC/JPEG+OFA media engine slice (R470+) [official].
  Source: https://docs.nvidia.com/datacenter/tesla/pdf/MIG_User_Guide.pdf
- A100 80GB: 1g.10gb / 1g.10gb+me / 1g.20gb / 2g.20gb / 3g.40gb / 4g.40gb / 7g.80gb [official].
  Source: https://docs.nvidia.com/datacenter/tesla/pdf/MIG_User_Guide.pdf
- H100 80GB (PCIe/SXM): 1g.10gb / 1g.20gb / 2g.20gb / 3g.40gb / 4g.40gb / 7g.80gb — 7 max; H100 NVL 94GB scales to 1g.12gb/1g.24gb/2g.24gb/3g.47gb/4g.47gb/7g.94gb [secondary].
  Source: https://github.com/interloperok/ai.infracalculator/blob/HEAD/docs/mig-feasibility.md
- H200 141GB: 1g.18gb / 1g.35gb / 2g.35gb / 3g.71gb / 4g.71gb / 7g.141gb; GH200 inherits the H200 profile set [secondary].
  Source: https://github.com/interloperok/ai.infracalculator/blob/HEAD/docs/mig-feasibility.md
- B200 SXM 192GB (Blackwell datasheet): 1g.23gb / 1g.45gb / 2g.45gb / 3g.90gb / 4g.90gb / 7g.180gb — max 7; B300 SXM 288GB scales the same geometry [secondary].
  Source: https://github.com/interloperok/ai.infracalculator/blob/HEAD/docs/mig-feasibility.md
- K8s exposure: `mig-strategy=mixed` exposes `nvidia.com/mig-<compute>g.<memory>gb` resources (e.g. `nvidia.com/mig-1g.10gb`); `single` keeps `nvidia.com/gpu` semantics per MIG device with a node selector; without MIG, partial-GPU sharing is impossible [secondary].
  Source: https://github.com/nvidia/deepops/blob/HEAD/workloads/examples/k8s/gpu-usage/README.md
- [unverified] B200 MIG support details come from community calculators citing the Blackwell datasheet, not the public MIG User Guide r580 captured here — re-check before fleet planning.

---
