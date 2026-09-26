---
id: etape9-phased-data-protection-raid/00-data-protection-raid/d2-end-to-end-data-path-protection-t10-pi-dif-dix-and-nvme
title: "D2 — End-to-end data path protection: T10 PI / DIF / DIX and NVMe"
domain: step-9-phase-d-data-protection-raid-hardware
role: deep-dive
task: hardware
actors: []
dates: []
keywords: ["latency"]
source: docs/RAG/etape9_phaseD_data_protection_raid.md
source_anchor: ""
source_lines: [39, 50]
section: "Step 9 — Phase D: Data Protection & RAID Hardware"
sha256: 2740066cab97ec4bbb1aa7bd3c10a960b50fc47f7922a1ba6182c71adc974f28
---

# D2 — End-to-end data path protection: T10 PI / DIF / DIX and NVMe

- **E1.S / E3.S and OCP:** hyperscale form factors (E1.S, E3.S) carry their own power-loss and thermal specs under OCP — the PM9A3's E1.S variant exists precisely for dense flash shelves where M.2 thermals don't work [secondary].
- **PLP in the NVMe spec:** the spec defines volatile write cache behavior and flush commands, but PLP capacitors are a vendor implementation choice — "NVMe compliant" does not imply power-loss protection. Check the datasheet line item, not the protocol [secondary].
- **Host-controlled thermal management (HCTM):** NVMe's HCTM lets the host set thermal policies per drive — data-center orchestration can cap drive thermals before the controller's own throttling engages [secondary].
- **Predictable latency mode:** some enterprise NVMe drives offer deterministic-latency modes that trade peak IOPS for consistent latency — relevant to RAID rebuild QoS and tail-latency SLAs [secondary].
- **PLP and sudden vs graceful:** the drive cannot distinguish "host crashed" from "power cut" — both are unsafe shutdowns. Design as if every shutdown is sudden; graceful paths are an optimization, not the plan [secondary].
- **Multi-actuator HDDs:** dual-actuator drives (two independent head stacks) halve seek latency but don't change any reliability math in this file — treat each actuator's surfaces under the same UBER/AFR model [secondary].
- **SMR HDDs in RAID:** shingled magnetic recording drives have pathological RAID-rebuild behavior (persistent cache exhaustion); avoid SMR for RAID members — CMR only [secondary].
- **Drive-managed vs host-managed SMR:** host-managed SMR in RAID is an operational hazard; if SMR appears in procurement at all, restrict to explicitly-supported archival tiers [secondary].
- **Power-loss during rebuild:** a power cut mid-rebuild restarts or resumes the rebuild depending on controller journaling — CacheVault-protected cards resume cleanly; unprotected write-back can leave the array inconsistent [secondary].

## D2 — End-to-end data path protection: T10 PI / DIF / DIX and NVMe

