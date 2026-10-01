---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/questions-1632914-opnsense-firewall-scheduled-rule-does-not-work-59cb14ff
title: "questions-1632914-opnsense-firewall-scheduled-rule-does-not-work-59cb14ff"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/questions-1632914-opnsense-firewall-scheduled-rule-does-not-work-59cb14ff.md
source_anchor: ""
source_lines: [1, 38]
sha256: 95fe1ef882c7bef342099dbd053989ba16d5cc1d73b5bca6433c8f13f32dcf5c
---

# questions-1632914-opnsense-firewall-scheduled-rule-does-not-work-59cb14ff

I have created a schedule for internet access for a VM (10.0.64.43/27), the rule is implemented on a WAN interface but does not seem to be working. The internet access is to be allowed between 21:30 - 21:45 every Mon, Thu, and Sun yet the VM has internet access all the time.

Schedule - https://i.ibb.co/qm5FCMF/Schedules.png

WAN Rule - https://i.ibb.co/TgxLTY7/WAN-Rules.png

Rule Ineffective - https://i.ibb.co/QcBzVpD/Schedule-Failure.png

Could it be that NAT is being applied to the 10.0.64.0/27 network before packets reach the WAN and thus the rule is ineffective.

Any thoughts what might be wrong in this case.

**UPDATE**


I had a hard time understanding In and Out of the firewall in relation to Source and Destination.

Whatever I understood I implemented but only part of the scheduled rule is effective, the network 192.168.28.0 has a scheduled internet access and works fine. **The network 10.0.64.0 does not seem to be effective**.

The whole network with internet route for client VM - https://i.ibb.co/9gHG3y3/Dell-Network.png

Tracert from client (**192.168.1.21 is the 1_dell interface** and **192.168.47.2 is the NAT network in VMware Workstation**) - https://i.ibb.co/PG8YKs5/W10-Tracert-Internet.png

Schedule - https://i.ibb.co/JFqL03v/Schedule.png

Server with No Internet as in Schedule - https://i.ibb.co/HVbMPcv/Server-No-Internet.png

Alias for RFC1918 networks - https://i.ibb.co/9HXZ7t0/RFC1918.png

Internet Rule for 10.0.64.32 /27 - https://i.ibb.co/9Wn5RQv/firewallwm-RFC1918.png

Internet still accessible - https://i.ibb.co/TB7jRhd/W10-Internet.png

WAN Rule - https://i.ibb.co/YN28rzs/firewallwm-WAN-Rule.png

Not sure if my rule is incorrect or its a glitch I'm failing to understand how to implement.

4more comments
