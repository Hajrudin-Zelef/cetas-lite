---
id: etape6-phased4-segmentation-qos-multicast/00-segmentation-qos-multicast/wave-22-reference-tables
title: "Wave 22 — Reference tables"
domain: step-6-phase-d4-segmentation-qos-multicast-network-services
role: deep-dive
task: reference
actors: []
dates: []
keywords: ["ethernet", "voice"]
source: docs/RAG/etape6_phaseD4_segmentation_qos_multicast.md
source_anchor: ""
source_lines: [709, 770]
section: "Step 6 — Phase D4: Segmentation, QoS, Multicast & Network Services"
sha256: a304da5cad7b85a5d083e5c6cf4cc39b9599dbe8b8b9c2b171bc5ccb7c5477fd
---

# Wave 22 — Reference tables

## Wave 22 — Reference tables

### 22.1 DSCP codepoints quick reference
| DSCP | Decimal | Name | Typical DC use |
|---|---|---|---|
| CS0 | 0 | Best effort | Default / scavenger-adjacent |
| CS1/AF11-13 | 8/10,12,14 | Priority/scavenger/AF1x | Bulk, background replication |
| AF21-23 | 18,20,22 | AF2x | Transactional data (template) |
| AF31-33 | 26,28,30 | AF3x | Signaling / storage (operator choice) |
| CS4/AF41 | 32/34 | — | Real-time video (template) |
| EF | 46 | Expedited Forwarding | Voice; also RoCE CNP marking per ZTR |
| CS6 | 48 | Network control | Routing protocols, PTP |
| CS7 | 56 | Network control | Reserved control |
`[official]` — codepoint values per RFC 2474/2597/3246/4594; "typical DC use" column is operator convention, not standard `[secondary]`.

### 22.2 IPv4 multicast address ranges
| Range | Scope |
|---|---|
| 224.0.0.0/24 | Link-local (bridged, never routed — per Cisco TRM note) |
| 224.0.1.0–238.255.255.255 | Global/organization-local scope |
| 239.0.0.0/8 | Administratively scoped (site/local) |
| 232.0.0.0/8 | SSM default range |
`[official]` — IANA registry; Cisco TRM whitepaper for the 224.0.0.0/24 bridging rule (Wave 8).

### 22.3 Key protocols & ports
| Protocol | Port / EtherType | Purpose |
|---|---|---|
| VXLAN | UDP 4789 | Overlay encapsulation (RFC 7348) |
| PIM | IP proto 103 | Multicast routing |
| IGMP | IP proto 2 | Host membership |
| MSDP | TCP 639 | Inter-RP source discovery |
| PTP | UDP 319/320 | Precision timing (IEEE 1588) |
| NTP | UDP 123 | Time sync |
| DHCP | UDP 67/68 | Address assignment |
| DCBx | LLDP TLVs | DCB capability exchange |
`[official]` — IANA/RFC/IEEE assignments.

### 22.4 Glossary (one-line)
- **BUM**: broadcast/unknown-unicast/multicast — multi-destination traffic needing special overlay handling.
- **VTEP**: VXLAN tunnel endpoint — encapsulates/decapsulates.
- **IMET/SMET**: EVPN inclusive/selective multicast Ethernet tag routes (Type 3/6).
- **DF**: designated forwarder — the one PE forwarding BUM to a multi-homed CE.
- **ESI**: Ethernet segment identifier — names a multi-homed attachment.
- **RD/RT**: route distinguisher / route target — uniqueness and import policy for VPN/EVPN routes.
- **CNP**: congestion notification packet — ECN feedback in DCQCN.
- **WRED**: weighted random early detection — probabilistic drop/mark before queues fill.
- **QCN/DC-QCN**: (data-center) quantized congestion notification — sender rate control.
- **TRM**: tenant routed multicast — EVPN-based inter-subnet multicast routing.
- **DAG**: distributed anycast gateway — same gateway IP/MAC on all leaves.

### Wave 22 verification
- Sources: RFC 2474/2597/3246/4594 (DSCP), IANA registries, RFC 7348, IEEE 1588, Cisco TRM whitepaper. Registry facts `[official]`; usage columns labeled as convention.

---

## Final verification (all 22 waves)

- Script checks after final append: line count ≥ 750 · tail intact · backtick parity even · headers well-formed.
- Waves: 1 VLAN scale · 2 VRF · 3 microsegmentation/policy · 4 QoS classification/queuing · 5 DCB · 6 ECN/RoCEv2 · 7 PIM/RP/MSDP · 8 VXLAN multicast/IGMP · 9 services · 10 EVPN control plane · 11 VRF deep-dive · 12 QoS deep-dive · 13 congestion control + UEC · 14 PIM/mVPN deep-dive · 15 services deep-dive · 16 segmentation practice · 17 multicast practice · 18 QoS practice · 19 EVPN configs · 20 RoCE QoS configs · 21 anycast-RP/TRM configs · 22 reference tables.
- Open items and conflicts consolidated in the Wave 21 final section; nothing new in Wave 22.

*End of Phase D4 — 22 waves, append-only, complete.*
