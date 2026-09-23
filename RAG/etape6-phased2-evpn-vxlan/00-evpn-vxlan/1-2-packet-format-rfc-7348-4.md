---
id: etape6-phased2-evpn-vxlan/00-evpn-vxlan/1-2-packet-format-rfc-7348-4
title: "1.2 Packet format (RFC 7348 §4)"
domain: step-6-phase-d-wave-2-evpn-vxlan-overlay
role: deep-dive
task: quantization
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/etape6_phaseD2_evpn_vxlan.md
source_anchor: ""
source_lines: [54, 100]
section: "Step 6 — Phase D wave 2: EVPN-VXLAN overlay"
sha256: eab63718eebacf37e7b01e6aa3f00125bdd23e117ec3efc717b0a70729680a0e
---

# 1.2 Packet format (RFC 7348 §4)

### 1.2 Packet format (RFC 7348 §4)

The wire format stacks as [official][secondary]:

```
Outer Eth | Outer IP (src=local VTEP, dst=remote VTEP) |
UDP (dport 4789, sport=flow entropy) | VXLAN (flags, VNI:24) |
Inner Eth | Inner payload
```

Details worth recording for capacity planning and troubleshooting:

- Encap overhead is **50 bytes** over an IPv4 underlay (14 outer Ethernet +
  20 outer IPv4 + 8 UDP + 8 VXLAN) [secondary]. Over IPv6 the outer IP is
  40 bytes, so 70 bytes total.
- UDP destination port is **4789**, the IANA-assigned VXLAN port [official].
  Legacy implementations (pre-RFC) used 8472; Linux's vxlan driver historically
  defaulted to 8472, which still causes interop gotchas with modern gear
  expecting 4789 [independent].
- The UDP **source port** carries flow entropy derived from the inner frame
  (RFC 7348 §5 references the RFC 6335 dynamic range). This is how the underlay
  gets ECMP entropy for VXLAN tunnels — without it, all tunneled traffic would
  hash to one path [official][secondary].
- UDP checksum is 0 over IPv4 (legal per §4.3; same choice as GTP-U) [official].
- VXLAN header: 8 bytes; first byte is flags (only the I-flag, bit 3, is
  defined — set to 1 for a valid VNI); next 3 bytes reserved; last 4 bytes carry
  the 24-bit VNI in the low 3 bytes [official: RFC 7348 §5].
- No fragmentation or PMTU handling: a grown frame exceeding the underlay MTU is
  dropped. Underlay MTU must therefore be raised (commonly 9216 or at least
  1550+) or overlay MTU clamped — a classic silent-failure cause [independent].

### 1.3 VNI (VXLAN Network Identifier)

- 24-bit field → 16,777,216 values, 0 to 16777215 [official].
- VNI 0 and reserved ranges: VNI 0 is reserved; operational practice reserves
  low values; exact reserved ranges vary by vendor — do not assume portability
  of a given "well-known" VNI [secondary][unverified].
- **L2VNI** vs **L3VNI**: an L2VNI identifies a bridged L2 segment (mapped to a
  local VLAN on each VTEP); an L3VNI identifies a VRF / routing domain used for
  inter-subnet routing (see IRB, Wave 2). This distinction is EVPN-era
  terminology, not RFC 7348's [secondary].
- VNI scope is global across the fabric: both ends of a tunnel must use the
  same VNI or traffic is dropped silently [secondary].
- Mapping convention: most vendors map VLAN X ↔ VNI 10000+X or similar local
  conventions; there is no standard mapping — it is per-VTEP configuration
  [secondary][vendor-reported].

