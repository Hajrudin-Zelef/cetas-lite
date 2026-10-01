---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/questions-1104425-opnsense-move-interface-to-other-hardware-port-4e270ffa
title: "OPNsense move interface to other hardware port"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/questions-1104425-opnsense-move-interface-to-other-hardware-port-4e270ffa.md
source_anchor: ""
source_lines: [1, 16]
sha256: 3bdd5a0a6c1bbf7204fd906781a6660d61f0f8b681a728313ed0c77442d9235a
---

# OPNsense move interface to other hardware port

*Score : 1 | Source : https://serverfault.com/questions/1104425/opnsense-move-interface-to-other-hardware-port*

I have an OPNsense with interfaces directy configured to the hardware ports. The corresponding switch port is also an access port.
We plan to change the switch port to a trunk port to transport multiple VLANs via this port. Is there a way to move the already configured OPNsense interface (IPs, rules, etc.) to the new corresponding vlan port?

---

### Reponse (acceptee) — score 2

I have figured it out.
- Add the VLAN interface to the hardware port via "Other Types"
- Change the FW interface assignment to the VLAN interface at "Assignments"
- Reboot to flush and recreate the state table
- In case of a cluster, start with the passive host and do a CARP Switch Over.
