---
id: collect-260926-mikrotik/mikrotik/questions-833139-mikrotik-forward-local-ip-to-another-local-ip-on-specific-port-bc6afb0f
title: "questions-833139-mikrotik-forward-local-ip-to-another-local-ip-on-specific-port-bc6afb0f"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/misc/questions-833139-mikrotik-forward-local-ip-to-another-local-ip-on-specific-port-bc6afb0f.md
source_anchor: ""
source_lines: [1, 13]
sha256: 79f1ebaf08a42d8b1e64f0f1a8181e262c924b186d02b6cc8c599fca6f7c2524
---

# questions-833139-mikrotik-forward-local-ip-to-another-local-ip-on-specific-port-bc6afb0f

I'm running a Docker container on a machine on port 8090. Let's say the IP address of that machine is 192.168.0.3. I want to forward 192.168.0.4:80 to 192.168.0.3:8090. The machine is and may only be reachable from within the local network.
What I've tried:
/ip firewall nat export
add action=dst-nat chain=srcnat dst-address=192.168.0.4 dst-port=80 \
    src-address=192.168.0.0/24 to-address=192.168.0.3 to-port=8090 \
    protocol=tcp
add action=masquerade chain=srcnat dst-address=192.168.0.4 dst-port=80 \
    src-address=192.168.0.0/24 protocol=tcp
This doesn't work.
Am I forgetting something? I've got the feeling I'm overlooking something simple but haven't been able to figure out what.
Edit:
The ip address 192.168.0.4 does not resolve to anything in my network. I just want to "assign" and forward it to 192.168.0.3:8090. The reason for this is that I want to setup a local domain name that resolves to the Docker container without having to specify a port.
192.168.0.4pingable? if so, does the firewall know about this ip address on one of its interfaces? if so, is the ARP entry for192.168.0.4the MAC address of your firewall or is it per chance flapping? try last question on a system in the same subnet witharping -I eth0 192.168.0.4(replace eth0 with the correct interface name if necessary)192.168.0.3looks way easier and more transparent to me.arpingcommand will surely result in a timeout as well. I will update my question with some additional information.
