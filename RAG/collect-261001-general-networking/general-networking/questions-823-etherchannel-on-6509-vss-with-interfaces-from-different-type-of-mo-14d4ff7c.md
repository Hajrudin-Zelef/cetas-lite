---
id: collect-261001-general-networking/general-networking/questions-823-etherchannel-on-6509-vss-with-interfaces-from-different-type-of-mo-14d4ff7c
title: "questions-823-etherchannel-on-6509-vss-with-interfaces-from-different-type-of-mo-14d4ff7c"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/questions-823-etherchannel-on-6509-vss-with-interfaces-from-different-type-of-mo-14d4ff7c.md
source_anchor: ""
source_lines: [1, 10]
sha256: 031283ca49c7d57deea8e09fa19fbdc9e785098aa22fc7eb90e43c75c4ed932c
---

# questions-823-etherchannel-on-6509-vss-with-interfaces-from-different-type-of-mo-14d4ff7c

We currently have 1 10GB Multi-Chassis etherchannel Between 2x 6509 (VSS) and 2x 3750-E (Stack).
On the 6509 the etherchannel interface members currently run on the supervisor module VS-S720-10G.
We want to migrate that etherchannel over to a WS-X6704-10GE module, and was considering doing it in the following procedure:
- Shut down interface Te1/5/5 (member of etherchannel)
- Remove interface from etherchannel (no channel-group)
- Add interface from WS-X6704-10GE into etherchannel (Te1/3/4)
- Move fiber cable
- Redo same procedure for Te2/5/5 and Te2/3/4
When doing this we will for a short while have a etherchannel with members from different type of modules. Is this supported by Cisco, or will this work at all?
Feel free to suggest other approaches to this migration...
