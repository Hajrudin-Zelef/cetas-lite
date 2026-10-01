---
id: collect-261001-huawei/huawei/questions-22558-assistance-with-routing-between-layer-3-switches-with-ospf-e17f63a8
title: "questions-22558-assistance-with-routing-between-layer-3-switches-with-ospf-e17f63a8"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/questions-22558-assistance-with-routing-between-layer-3-switches-with-ospf-e17f63a8.md
source_anchor: ""
source_lines: [1, 20]
sha256: ae4532f46b438bffe805fb5b668e84377f502cd819464c35aea41a1c7b34e3e7
---

# questions-22558-assistance-with-routing-between-layer-3-switches-with-ospf-e17f63a8

I have two Cisco 3560's in a branch office. One acts as an access switch configured with two vlans:
interface Vlan10
 ip address 192.168.0.254 255.255.255.0
interface Vlan20
 ip address 172.16.0.254 255.255.255.0
This is connected to the same model switch in a Comms room that will be routing back to the Core network. The Core for this exercise is also a 3560 layer 3 switch connected as a point to point link.
I have added the VLANs to all switches, but with different SVIs (i.e x.x.x.253/24 and x.x.x.252/24)
I have also enabled IP Routing on each switch. I would like to use OSPF to route the traffic over the network and have configured this.
I have also configured the interconnecting interfaces with 'no switchport' and assigned an IP address. These networks have been added to the OSPF routing table.
Interfaces on the Access switch have been configured with VLANs and the PCs are assigned static IP Addresses
Access_Switch#
interface GigabitEthernet0/1
 switchport access vlan 10
 switchport mode access
end
interface GigabitEthernet0/2
 switchport access vlan 20
 switchport mode access
end
With this information should I be able to ping a host connected to the core (or even telnet the Core switch) from a host connected to the Access switch in the branch office? Thanks in advance :)
