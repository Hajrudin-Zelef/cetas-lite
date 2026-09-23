---
id: etape6-phased4-segmentation-qos-multicast/00-segmentation-qos-multicast/wave-20-config-walkthroughs-rocev2-lossless-qos-dell-os10
title: "Wave 20 — Config walkthroughs: RoCEv2 lossless QoS (Dell OS10)"
domain: step-6-phase-d4-segmentation-qos-multicast-network-services
role: deep-dive
task: reference
actors: ["Nvidia"]
dates: []
keywords: ["asic", "nvidia"]
source: docs/RAG/etape6_phaseD4_segmentation_qos_multicast.md
source_anchor: ""
source_lines: [612, 658]
section: "Step 6 — Phase D4: Segmentation, QoS, Multicast & Network Services"
sha256: 8c78fffc12ce2ab7577d0eb525147cd52725ad0276a80a175f602ea60241797d
---

# Wave 20 — Config walkthroughs: RoCEv2 lossless QoS (Dell OS10)

## Wave 20 — Config walkthroughs: RoCEv2 lossless QoS (Dell OS10)

Verbatim-structure walkthrough from the Dell SmartFabric OS10 10.5.0 User Guide (configure-roce-on-the-switch) `[official]` — https://www.dell.com/support/manuals/en-in/dell-emc-smartfabric-os10/smartfabric-os-user-guide-10-5-0/configure-roce-on-the-switch?guid=guid-ddfa4455-8f64-4014-8543-ceb6719c904d&lang=en-us

```
! Step 1 — queuing class-maps: Q0 = lossy default, Q3 = lossless RoCE
class-map type queuing Q0
  match queue 0
class-map type queuing Q3
  match queue 3

! Step 2 — WRED+ECN profile for the lossless queue
wred wred_ecn
  random-detect ecn
  random-detect color green  minimum-threshold 1000 maximum-threshold 2000 drop-probability 100
  random-detect color yellow minimum-threshold 500  maximum-threshold 1000 drop-probability 100
  random-detect color red    minimum-threshold 100  maximum-threshold 500  drop-probability 100
  exit

! Step 3 — queuing policy: bandwidth split + ECN on Q3
policy-map type queuing policy_2Q
  class Q0
    bandwidth percent 30
    exit
  class Q3
    bandwidth percent 70
    random-detect wred_ecn
    end

! Step 4 — ETS traffic-class mapping: qos-group 3 -> queue 3 (lossless)
qos-map traffic-class 2Q
  queue 0 qos-group 0-2, 4-7
  queue 3 qos-group 3

! Step 5 — (per guide) apply policy-map to interfaces + enable PFC on the
! RoCE priority via network-qos policy; trust DSCP on fabric links
```

- Reading the numbers: green/yellow/red are WRED colors (drop precedence); thresholds are in the platform's buffer units — do not port these numbers to another ASIC `[official]` — Dell guide.
- NVIDIA `mlnx_qos` host-side complement (from Wave 6, FS.com whitepaper): DCBX OS-controlled, priority trust `dscp`, per-priority PFC bitmap, per-TC buffer sizes — the switch config above must MATCH the NIC's dscp2prio mapping or PFC pauses the wrong traffic `[vendor-reported]` — FS.com.
- Cross-check ritual: `show qos maps`, `show policy-map interface`, PFC pause counters, ECN-mark counters — all four must agree before declaring the fabric lossless `[secondary]` — operator practice.

### Wave 20 verification
- Sources: 3. Dell commands transcribed from the official 10.5.0 guide; NIC-side from FS.com whitepaper. Portability warning explicit.

---

