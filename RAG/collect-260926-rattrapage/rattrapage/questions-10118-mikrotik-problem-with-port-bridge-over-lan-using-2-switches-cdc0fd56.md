---
id: collect-260926-rattrapage/rattrapage/questions-10118-mikrotik-problem-with-port-bridge-over-lan-using-2-switches-cdc0fd56
title: "questions-10118-mikrotik-problem-with-port-bridge-over-lan-using-2-switches-cdc0fd56"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-rattrapage/ai-llm/questions-10118-mikrotik-problem-with-port-bridge-over-lan-using-2-switches-cdc0fd56.md
source_anchor: ""
source_lines: [1, 27]
sha256: fcca98fc5d1ae0e097fb743584aae28349ed438d5fcb3ef63010064c8e07e860
---

# questions-10118-mikrotik-problem-with-port-bridge-over-lan-using-2-switches-cdc0fd56

I'd suggest to remove ports from bridge if it is possible to add them to switch group. When your ports are bridged - all packets going through these interface are processed using CPU while, if they would have been in the same switch group, special switch chip would process those packets, decreasing load on CPU.
So, in your case, you can make ether1 to be a master port for ether2-5. Don't forget that before assigning master port you have to remove them from bridge (bridge-local). Then, do the same for the second switch group (ports 6-10), so ether6 will be a master port for ether7-10.
/interface ethernet
  set ether1 master-port=none
  set ether2 master-port=ether1
  set ether3 master-port=ether1
  set ether4 master-port=ether1
  set ether5 master-port=ether1
  set ether6 master-port=none
  set ether7 master-port=ether6
  set ether8 master-port=ether6
  set ether9 master-port=ether6
  set ether10 master-port=ether6
After doing that you'll get two separate and independent switch groups, however, as you mentioned it in your post, you want them to be connected. In such case, there are two options:
- to use software bridge;
- connect them using cable.
Software (CPU) bridging
In this case, switch groups 1 and 2 will be switched using CPU, so any packet going from any interface in group 1 to group 2, or vice versa, will be processed using CPU.
/interface bridge
  add name=bridge-local
  port add interface=eth1 bridge=bridge-local
  port add interface=eth6 bridge=bridge-local
NOTE: don't forget to assign an IP address to the bridge and, if you're using DHCP server on Mikrotik, change its interface to bridge-local as well.
/ip address add interface=bridge-local address=192.168.1.1/24
/ip dhcp-server add interface=bridge-local address-poll=[name of your address pool]
Patch cord
Just connect any port from group 1 (ether1-5) to any port in group 2 (ether6-10). Yes, it takes 2 port and only applicable if you have ones. If all ports in one of the switch groups are already in use, I'd suggest using first option (bridging).
