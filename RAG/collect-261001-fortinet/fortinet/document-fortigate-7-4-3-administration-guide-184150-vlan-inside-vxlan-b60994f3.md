---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-3-administration-guide-184150-vlan-inside-vxlan-b60994f3
title: "document-fortigate-7-4-3-administration-guide-184150-vlan-inside-vxlan-b60994f3"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-3-administration-guide-184150-vlan-inside-vxlan-b60994f3.md
source_anchor: ""
source_lines: [1, 34]
sha256: 4b09c5419cb3d7f1ce7c1db97e4b066f17308182b7b7351a876a05c35ab68abe
---

# document-fortigate-7-4-3-administration-guide-184150-vlan-inside-vxlan-b60994f3

VLAN inside VXLAN
VLAN inside VXLAN
VLANs can be assigned to VXLAN interfaces. In a data center network where VXLAN is used to create an L2 overlay network and for multitenant environments, a customer VLAN tag can be assigned to VXLAN interface. This allows the VLAN tag from VLAN traffic to be encapsulated within the VXLAN packet.
To configure VLAN inside VXLAN on HQ1:
- Configure VXLAN:config system vxlan
   edit "vxlan1"
      set interface port1
      set vni 1000
      set remote-ip 173.1.1.1next end
- Configure system interface:config system interface
   edit vlan100
     set vdom root
     set vlanid 100
     set interface dmz
   next
   edit vxlan100
     set type vlan
     set vlanid 100
     set vdom root
     set interface vxlan1
   next
end
- Configure software-switch:config system switch-interface edit sw1 set vdom root set member vlan100 vxlan100 set intra-switch-policy implicit next end
|  | The default intra-switch-policy implicit behavior allows traffic between member interfaces within the switch. Therefore, it is not necessary to create firewall policies to allow this traffic. | 
|  | Instead of creating a software-switch, it is possible to use a virtual-wire-pair as well. See Virtual wire pair with VXLAN. | 
To configure VLAN inside VXLAN on HQ2:
- Configure VXLAN:
			config system vxlan edit "vxlan2" set interface port25 set vni 1000 set remote-ip 173.1.1.2 next end
- Configure system interface:config system interface edit vlan100 set vdom root set vlanid 100 set interface port20 next edit vxlan100 set type vlan set vlanid 100 set vdom root set interface vxlan2 next end
- Configure software-switch:			config system switch-interface edit sw1 set vdom root set member vlan100 vxlan100 next end
To verify the configuration:
Ping PC1 from PC2.
The following is captured on HQ2:
This captures the VXLAN traffic between 172.1.1.1 and 172.1.1.2 with the VLAN 100 tag inside.
