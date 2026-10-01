---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/questions-1067399-opnsense-nat-port-forward-forward-multiple-protocols-and-ports-8b847400
title: "OPNsense NAT/Port Forward: Forward multiple protocols and ports"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/questions-1067399-opnsense-nat-port-forward-forward-multiple-protocols-and-ports-8b847400.md
source_anchor: ""
source_lines: [1, 19]
sha256: f8ce5e0e640550b3f6a5b7c0c34fa69cf63d42ba3abe2e92b738dcd81df340c1
---

# OPNsense NAT/Port Forward: Forward multiple protocols and ports

*Score : 1 | Source : https://serverfault.com/questions/1067399/opnsense-nat-port-forward-forward-multiple-protocols-and-ports*

I want to forward ICMP and specific TCP and UDP ports on OPNsense but I'm unable to find a concise solution. Specifically I want to forward ICMP, http, https and UDP 32768-65535.
I'm adding a new port forward in the port forwarding section ("Firewall>NAT>Port Forward"). Here if I select "any" protocol, then I can not specify TCP/UDP ports. If I select TCP/UDP in protocol then specified ports will be open for both TCP and UDP and I can not specify ICMP with this. I can create separate rules for separate protocols but it seems unintuitive.
Will be glad to provide any further clarification if required.
Thanks in advance.
Update: I understand only TCP/UDP has concept of ports. I want to forward ICMP port for testing/reachability check and I'm forwarding an IP from my BGP network not one assigned to any interface like WAN. UDP 32768-65535 for a videoconferencing app (BigBlueButton). The UDP ports are not required in my case as my bigbluebutton and coturn instance are both inside the firewall.
I think it would be more intuitive if I could be able to list all ports/forwards for a NAT mapping at a single place. I felt there may be some way to enter a list like the following to a NAT. TCP/80, TCP/443, TCP/22, UDP/100:200, ICMP
From the current answer I think this is not available. I Will submit a feature request.

---

### Reponse (acceptee) — score 1

Only TCP and UDP has the concept of ports. If you specify any as protocol, you can't specify ports, as it's not relevant for most protocols.
You'll have to make multiple forwarding rules:
This is not a problem; you can have as many forwarding rules as you want, all forwarding to the same destination.
