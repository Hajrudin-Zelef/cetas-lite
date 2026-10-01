---
id: collect-261001-general-networking/general-networking/t5-security-sd-wan-best-practices-for-native-vlan-configuration-m-p-93478-538e25b2-3
title: "t5-security-sd-wan-best-practices-for-native-vlan-configuration-m-p-93478-538e25b2"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-security-sd-wan-best-practices-for-native-vlan-configuration-m-p-93478-538e25b2.md
source_anchor: ""
source_lines: [197, 197]
sha256: 1edad3c0d25bdacf2e6fb322c28a277dd597bd4ccaad03e471b0d5a527cebfd2
---

# t5-security-sd-wan-best-practices-for-native-vlan-configuration-m-p-93478-538e25b2

I believe it is more secure to avoid using a native VLAN on trunks and to explicitly list the VLAN traffic to be passed. In some instances, the software will require that a native VLAN be entered; if this occurs, then enter the VLAN ID of a non-existent VLAN. For fairly self-evident reasons I choose to use VLAN 101 (aka Room 101).
