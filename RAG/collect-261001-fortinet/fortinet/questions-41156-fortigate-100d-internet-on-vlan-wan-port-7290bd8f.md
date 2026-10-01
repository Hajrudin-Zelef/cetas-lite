---
id: collect-261001-fortinet/fortinet/questions-41156-fortigate-100d-internet-on-vlan-wan-port-7290bd8f
title: "questions-41156-fortigate-100d-internet-on-vlan-wan-port-7290bd8f"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/questions-41156-fortigate-100d-internet-on-vlan-wan-port-7290bd8f.md
source_anchor: ""
source_lines: [1, 44]
sha256: c484ee8e8a9a9588dc60994e8dd1100c4252f1a7fc161b26d22b4643aaed9b35
---

# questions-41156-fortigate-100d-internet-on-vlan-wan-port-7290bd8f

I am fairly new towards Fortigate firewalls and I am trying to set up one FortiGate 100D running firmware v5.0 as a router for a hotel network. I have 2 ISPs using PPPoE connection that runs on VLAN 500. I created VLANs for both wan1 and wan2 (assuming ISP_A_vlan and ISP_B_vlan) with all ISPs' credential set and I am able to connect and obtain an IP from the ISP, which is fine.
Of the 16 ports I have, I've separated it into 8 different networks in pair of 2 ports per network with DHCP enabled, where it serves different IP in different network pair. For example, Office_lan @ 10.1.1.1/24 and Access_point @ 10.1.0.1/24
Next, I created a policy for both ISP_A_vlan and ISP_B_vlan with the following configuration:
policy type: firewall
policy subtype: address
incoming interface: Access_point
source address: all
outgoing interface: ISP_A_vlan, ISP_B_vlan 
destination address: ALL
schedule: always
service: all
enable NAT: true (use destination interface address)
Similarly for Office_lan:
policy type: firewall
policy subtype: address
incoming interface: Office_lan
source address: all
outgoing interface: ISP_A_vlan, ISP_B_vlan 
destination address: ALL
schedule: always
service: all
enable NAT: true (use destination interface address)
but when I inspect the Foward Traffic log, there are no internet connection from either Office_lan or Access_point. They are able to negotiate an IP from the DHCP server as well as obtained a DNS server from the ISP itself.
With the similar setup, instead of running PPPoE on vlan wan ports, I took 2 routers (assuming wan1 for Router_A and wan2 for Router_B), configured the ISP's credential in routers then connecting both of the routers to wan1 and wan2 with almost similar configuration as follow:
policy type: firewall
policy subtype: address
incoming interface: Access_point
source address: all
outgoing interface: wan1, wan2 
destination address: ALL
schedule: always
service: all
enable NAT: true (use destination interface address)
Similarly for Office_lan:
policy type: firewall
policy subtype: address
incoming interface: Office_lan
source address: all
outgoing interface: wan1, wan2 
destination address: ALL
schedule: always
service: all
enable NAT: true (use destination interface address)
and everything works fine. Am I missing something right here?
