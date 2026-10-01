---
id: collect-261001-cisco/cisco/questions-995105-comparable-ubiquiti-to-cisco-native-vlan-configuration-1a590f92
title: "questions-995105-comparable-ubiquiti-to-cisco-native-vlan-configuration-1a590f92"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-cisco/questions-995105-comparable-ubiquiti-to-cisco-native-vlan-configuration-1a590f92.md
source_anchor: ""
source_lines: [1, 38]
sha256: 403a09dfd5f4e1850d58e47a522af8faf433cb3a3f6af444ad0580ac4de8896c
---

# questions-995105-comparable-ubiquiti-to-cisco-native-vlan-configuration-1a590f92

I am trying to replace a cisco router with a ubiquiti router. On the Cisco, you don't address the base interface. You only address the vlans. Then, you choose the native vlan and that will match the untagged vlan port that connects the switch. On Ubiquiti, from what I understand, you address the base interface with the subnet matching the vlan and address the rest of your vlans but never specify a vlan for what would be "native" in cisco. Is this the correct config?
    ethernet eth2 {
    address 10.1.2.2/24
    duplex auto
    speed auto
    vif 10 {
        address 10.1.10.2/24
    }
}
In this scenario, in the unms portal, I see tons of dropped packets. I don't know why they are dropping.
Here is my switch port config, the ubnt router is connected to C23
vlan 2
   name "Management"
   untagged A1-A2,A19,B20,B22,B24,C23-C24,D20,Trk4
   tagged A4,A8,E1,E4-E8,F1-F6,F8,Trk1-Trk2
   ip address 10.1.2.10 255.255.255.0
   exit
vlan 10
   name "DATA"
   untagged A4,A8-A18,A20,B1-B19,B21,B23,C1-C22,D1-D19,E2-E3,F7,Trk1-Trk2
   tagged A19,C23-C24,E1,E4-E8,F1-F6,F8,Trk4
   no ip address
   exit
Here is the config I am trying to mimic on the cisco router:
interface Port-channel1
 description Major Interface for VLAN Routing
 no ip address
 hold-queue 150 in
!
interface Port-channel1.2
 description MANAGEMENT VLAN INTERFACE
 encapsulation dot1Q 2 native
 ip address 10.1.2.1 255.255.255.0
!
interface Port-channel1.10
 description DATA VLAN INTERFACE
 encapsulation dot1Q 10
 ip address 10.1.10.1 255.255.255.0
