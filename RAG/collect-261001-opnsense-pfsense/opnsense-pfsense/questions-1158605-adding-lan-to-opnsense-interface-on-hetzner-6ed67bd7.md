---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/questions-1158605-adding-lan-to-opnsense-interface-on-hetzner-6ed67bd7
title: "Adding LAN to opnsense interface on hetzner"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/questions-1158605-adding-lan-to-opnsense-interface-on-hetzner-6ed67bd7.md
source_anchor: ""
source_lines: [1, 12]
sha256: b97f903e048cb49f386487d065a4b547aca5259f9a7e63a7ebb9e52b5dc090cc
---

# Adding LAN to opnsense interface on hetzner

*Score : 0 | Source : https://serverfault.com/questions/1158605/adding-lan-to-opnsense-interface-on-hetzner*

I am a newbie in opnsense. I installed opnsense on hetzner. I want to add LAN to interface but I can't find LAN interface on the assignment menu. What I can add is opt1. How do I add LAN to interfaces?

---

### Reponse (acceptee) — score 0

Go to Interfaces > Assignment > Add wan1 > save and apply changes
Go to interfaces > opt1 > enable interface > IPv4 Configuration Type to DHCP > save > apply changes
