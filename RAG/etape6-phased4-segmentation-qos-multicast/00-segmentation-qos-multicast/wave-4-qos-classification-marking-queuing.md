---
id: etape6-phased4-segmentation-qos-multicast/00-segmentation-qos-multicast/wave-4-qos-classification-marking-queuing
title: "Wave 4 — QoS: classification, marking, queuing"
domain: step-6-phase-d4-segmentation-qos-multicast-network-services
role: deep-dive
task: reference
actors: ["Nvidia"]
dates: []
keywords: ["asic", "ethernet", "nvidia", "parameters", "training", "voice"]
source: docs/RAG/etape6_phaseD4_segmentation_qos_multicast.md
source_anchor: ""
source_lines: [102, 148]
section: "Step 6 — Phase D4: Segmentation, QoS, Multicast & Network Services"
sha256: a75fb6d648d21e1874e76543b79f6b40d965a7a3d46667dc479ee0388e082cfd
---

# Wave 4 — QoS: classification, marking, queuing

## Wave 4 — QoS: classification, marking, queuing

### 4.1 Classification & marking fields
- L2: 802.1p/CoS (3-bit PCP in 802.1Q tag, 8 values). L3: DSCP (6-bit in IP ToS byte, 64 values) and legacy IP Precedence (3-bit) `[official]` — Cisco QoS config guides.
- Default Cisco CoS→DSCP map: CoS 0→0, 1→8, 2→16, 3→24, 4→32, 5→40, 6→48, 7→56; default DSCP→CoS map: DSCP 0–7→CoS 0, 8–15→1, 16–23→2, 24–31→3, 32–39→4, 40–47→5, 48–55→6, 56–63→7; default IP-Precedence→DSCP: IPP n → DSCP 8n `[official]` — Cisco IOS/IOS-XE QoS guides (2960-XR, 9500, 9400). https://www.Cisco.com/c/en/us/td/docs/switches/lan/catalyst9500/software/release/16-5/configuration_guide/qos/b_165_qos_9500_cg/b_165_qos_9500_cg_chapter_01.pdf · https://www.cisco.com/c/en/us/td/docs/switches/lan/catalyst9400/software/release/16-6/configuration_guide/qos/b_166_qos_9400_cg/b_166_qos_9400_cg_chapter_01.pdf
- Trust models: trust CoS, trust DSCP, trust IP-precedence, or untrusted (port default); on trust-DSCP boundaries between QoS domains, DSCP-to-DSCP mutation maps remap values `[official]` — Cisco 3560/2960-XR guides.
- Policing can remark via policed-DSCP map (mark-down of out-of-profile traffic); default policed-DSCP map is a null map (no markdown) `[official]` — Cisco.
- Well-known DSCP codepoints operators actually use: EF 46 (voice), AFxy family, CS6 48/CS7 56 (network control); RoCEv2 deployments commonly use DSCP 26 or 48 for RDMA traffic and mark CNP with DSCP 46 — values are operator/vendor conventions, see Wave 6 `[secondary]`.

### 4.2 Queuing & scheduling
- Scheduler families: Strict Priority (SP — one queue always served first, starvation risk for lower queues), Weighted Round Robin (WRR), Deficit Round Robin / Deficit Weighted RR (DWRR, byte-fair), and hybrids like 1P3Q (one strict-priority + three WRR) `[official]` — Cisco QoS guides describe 1p3q8t egress structures (1 strict-priority, 3 WRR, 8 WRED thresholds each); Sophos docs confirm SP vs WRR as the two canonical choices `[vendor-reported]`.
- Drop mechanisms: tail drop vs WRED (Weighted Random Early Detection) with per-color thresholds; ECN-marking variants covered in Wave 6 `[official]` — Cisco; Dell OS10 WRED/ECN config `[official]` (Wave 6).
- DC default posture: two-queue defaults on some platforms (queue 0 = control traffic, queue 1 = everything else, per Cisco Catalyst default wired QoS) `[official]` — Cisco 2960-XR/9500 guides. Note: Catalyst campus defaults are NOT Nexus DC defaults — do not cross-apply.
- Vendor QoS models are NOT comparable 1:1: Cisco Nexus (MQC class-map/policy-map, 8 queues, queuing + network-qos + qos system classes), Arista EOS (traffic-class → traffic-class mapping, ECN/WRED per TC, PFC per priority), Dell OS10 (class-map type queuing/network-qos, qos-map traffic-class), NVIDIA Spectrum (shared-buffer, lossless pools, RoCE profile via `mlnx_qos`) — each has its own queue counts, buffer architectures, and defaults `[official]` — per-vendor docs; commands in Wave 6.
- Gap: a consolidated queue-count/buffer-size table per 2026 ASIC was not built — open item `[unverified]`.

### Wave 4 verification
- Sources: 7 (Cisco ×5 official, Sophos vendor-reported). Mapping tables transcribed verbatim. Explicit non-comparable warning recorded for cross-vendor QoS defaults.

---

## Wave 5 — Data Center Bridging: PFC, ETS, DCBX, QCN

### 5.1 The DCB stack (IEEE)
- IEEE DCB adds to classic Ethernet: 802.1Qbb Priority-based Flow Control (PFC), 802.1Qaz Enhanced Transmission Selection (ETS), 802.1Qau Quantized Congestion Notification (QCN), and the Data Center Bridging Exchange (DCBx) protocol `[official]` — Dell FN IOM 9.10 / S6100-ON 9.14.2.9 / C9010 9.14.2.6 config guides; Network World DCB primer. https://www.dell.com/support/manuals/en-us/poweredge-fx2/fn-iom-9.10.0.0-config-pub/Ethernet-Enhancements-in-Data-Center-Bridging?guid=guid-d5d16c53-54f4-498c-a35a-4dbed9c68bff&lang=en-us · https://www.networkworld.com/article/749844/cisco-subnet-ethernet-adapts-for-data-center-applications-part-1.html
- Dell caveat: Dell Networking OS supports only PFC, ETS, and DCBx — NOT 802.1Qau/QCN `[official]` — same Dell guides. Other vendors vary; check per-platform.
- Purpose: I/O convergence — LAN, storage (FCoE/SCSI, ~2K payloads, no frame-loss tolerance), and IPC/HPC traffic on one wire with lossless + lossy priorities simultaneously `[official]` — Dell.

### 5.2 PFC (802.1Qbb)
- Pause granularity: per 802.1p priority (8 delineations); a receiver sends a PFC pause frame for a specific CoS when its ingress buffer crosses a threshold, and the sender stops only that priority's traffic — lossless and lossy traffic coexist on the same wire `[official]` — Network World; Dell guides.
- PFC is hop-by-hop, not end-to-end: it keeps one link lossless; it does not fix congestion, it contains the symptom of a failed congestion response `[secondary]` — Sabyasachi Kar lab notebook (2026-05), connectx-7 + Tomahawk 5 testbed. https://medium.com/@sabyasachi.kar85/how-rocev2-stays-lossless-41db8734d6d9
- Failure modes: PFC storms (pause propagation), head-of-line blocking, and PFC deadlock — circular buffer-dependency chains (common in synchronized all-to-all LLM training) freezing traffic; mitigations include PFC watchdog (drop after a bounded pause, e.g., 200 ms in the lab test), deadlock detection in switch OSs, and reducing the buffer spikes that trigger pauses `[secondary]` — Kar notebook; AICPLight RoCEv2 guide. https://medium.com/@aicplight888/rocev2-network-solutions-transceiver-cable-deployment-guide-for-ai-clusters-c55d827d7582

### 5.3 ETS (802.1Qaz)
- Bandwidth guarantees via priority groups and traffic classes: ETS assigns minimum bandwidth shares to traffic classes so no-drop and best-effort classes get deterministic service `[official]` — Network World; Dell OS10 ETS/bandwidth config `[official]` (Wave 6).
- DCBx: LLDP-based negotiation/exchange of PFC/ETS/application-priority (e.g., iSCSI/FCoE/RoCE) parameters so the link partners agree; prevents a non-lossless device from silently receiving lossless traffic `[official]` — Network World; FS.com lossless RDMA whitepaper shows DCBX mode "OS controlled" and priority trust state `dscp` in a real `mlnx_qos` dump `[vendor-reported]`. https://resource.fs.com/mall/resource/cn_lossless-network-for-rdma-white-paper-20250626115109.pdf

### 5.4 QCN (802.1Qau) — mostly historical
- QCN lets congested switches sample traffic and signal senders to slow down rather than drop; some argued it was a prerequisite for multi-hop FCoE `[secondary]` — Network World (2010-era).
- In 2026 practice, L3 ECN + DC-QCN (data-center quantized congestion notification, RFC 3168 ECN-based, sender-side rate control driven by CNPs) has displaced QCN for RoCEv2 fabrics `[secondary]` — Kar notebook; intelligentvisibility lossless Ethernet guide. https://intelligentvisibility.com/ai-networking-solutions/lossless-networking-ai
- Non-comparable warning: "DC-QCN" tuning knobs (min/max thresholds, max-mark probability) are vendor-implementation-specific; the lab values below are one testbed's, not universal.

### Wave 5 verification
- Sources: 7. Dell's QCN non-support captured as a per-vendor caveat. PFC-scope limitation (hop-by-hop) emphasized to counter the common "PFC = lossless fabric" oversimplification.

---

