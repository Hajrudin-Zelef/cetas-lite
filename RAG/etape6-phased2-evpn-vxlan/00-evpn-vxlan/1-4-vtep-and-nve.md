---
id: etape6-phased2-evpn-vxlan/00-evpn-vxlan/1-4-vtep-and-nve
title: "1.4 VTEP and NVE"
domain: step-6-phase-d-wave-2-evpn-vxlan-overlay
role: deep-dive
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/etape6_phaseD2_evpn_vxlan.md
source_anchor: ""
source_lines: [101, 147]
section: "Step 6 — Phase D wave 2: EVPN-VXLAN overlay"
sha256: de4a885d0f39941125972def41c334ff7238d968cc4fb299399b68bc30b342ec
---

# 1.4 VTEP and NVE

### 1.4 VTEP and NVE

- **VTEP** (VXLAN Tunnel Endpoint): the device/function that encapsulates and
  decapsulates. In a DC fabric, every leaf switch is a VTEP; spine switches
  normally are not (they only see outer IP/UDP) [secondary].
- **VTEP IP**: the source/destination IP in the outer header — operational best
  practice is a loopback (commonly loopback0), stable across link failures
  [secondary][vendor-reported].
- **NVE** (Network Virtualization Edge): the logical interface representing the
  VTEP on some platforms — notably `interface nve1` on Cisco NX-OS. Other
  vendors expose the same concept under different names (Arista: the VXLAN
  interface `Vxlan1`; Cumulus/SONiC: Linux `vxlan` devices + FRR/EVPN control
  plane) [vendor-reported][secondary].
- Software VTEPs also exist: hypervisor vSwitches, Linux `ip link add ...
  type vxlan` devices, CloudStack's VXLAN plugin (which creates VTEP devices
  with `nolearning`, UDP 4789, IPv4 underlay, /32 loopback as VTEP source)
  [secondary].

### 1.5 MAC learning modes: flood-and-learn vs control-plane

Two fundamentally different ways a VTEP learns "destination MAC → remote VTEP
IP" [official][secondary]:

**Data-plane (flood-and-learn)** — RFC 7348's original model:

- Unknown unicast and broadcast/multicast (BUM) traffic is flooded to all VTEPs
  in the VNI via underlay **multicast** (one multicast group per VNI, or shared
  groups) or via **ingress replication** (head-end replication: the ingress VTEP
  sends a unicast copy to each known remote VTEP).
- MACs are learned from source addresses of decapsulated frames (classic
  learning), like a switch learning on ports.
- Pros: no control plane needed; simple. Cons: multicast in the underlay is
  operationally painful; flood-and-learn scales poorly; ARP storms replicate
  across the fabric; no host-mobility signaling [independent][secondary].

**Control-plane learning** — BGP EVPN (RFC 7432, Waves 2–3):

- MAC/IP bindings are advertised as BGP EVPN routes; BUM can use ingress
  replication driven by EVPN Type-3 (inclusive multicast) routes, or assisted
  replication.
- Pros: no underlay multicast required, ARP/ND suppression, host-mobility
  handling, multihoming. This is the production model for modern DC fabrics
  [vendor-reported][independent].
- Gartner/analyst-era consensus and every major vendor's current design guides
  position flood-and-learn as legacy/lab-only for greenfield fabrics
  [secondary][unverified — analyst wording not verified].

