---
id: collect-261001-cisco/cisco/enterprise-en-interworking-and-replacement-guide-vtp-thread-392883-861-5bc8ffe2-3
title: "Run the show running-config command to check the interface configuration."
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/enterprise-en-interworking-and-replacement-guide-vtp-thread-392883-861-5bc8ffe2.md
source_anchor: ""
source_lines: [100, 248]
sha256: c62d75b0aaad15ab124b4b8813b8a16ecdd597f44e50fc9a000f84740538bc2d
---

# Run the show running-config command to check the interface configuration.

In Figure 1-6, a Huawei S series switch directly connects to a Cisco VTP server. The Cisco switch and user hosts connected to the Huawei S series switch need to communicate in VLAN 10.
Figure 1-6 Hybrid networking of the C-H model
Configuration Roadmap
The configuration roadmap is as follows:
1. Check the configuration of the Cisco switch.
2. Create a VLAN manually on the Huawei S series switch and add interfaces to the VLAN.
Procedure
Step 1 Check the Cisco VTP server configuration. The display depends on the device configuration.
# Run the show running-config command to check the interface configuration.
! 
hostname VTP_Sever 
! 
interface GigabitEthernet5/1 
 switchport trunk encapsulation dot1q 
 switchport mode trunk 
! 
interface GigabitEthernet5/2 
 switchport trunk encapsulation dot1q 
 switchport mode trunk 
! 
interface GigabitEthernet5/3 
 switchport trunk encapsulation dot1q 
 switchport mode trunk 
!
If the interface configuration is incorrect, perform the following operations to configure the interface.
VTP_Sever# configure terminal 
VTP_Sever(config)# interface gigabitethernet 5/1 
VTP_Sever(config-if)# switchport trunk encapsulation dot1q   
VTP_Sever(config-if)# switchport mode trunk   
VTP_Sever(config-if)# exit 
VTP_Sever(config)# interface gigabitethernet 5/2 
VTP_Sever(config-if)# switchport trunk encapsulation dot1q 
VTP_Sever(config-if)# switchport mode trunk 
VTP_Sever(config-if)# exit 
VTP_Sever(config)# interface gigabitethernet 5/3 
VTP_Sever(config-if)# switchport trunk encapsulation dot1q 
VTP_Sever(config-if)# switchport mode trunk 
VTP_Sever(config-if)# exit 
# Run the show vlan brief command to check whether VLAN 10 has been created. If VLAN 10 is created, perform the following operation to create VLAN 10.
VTP_Sever(config)# vlan 10
# Run the show vtp status command to check whether the VTP working mode is server and whether the domain name is the same as that on the client. Run the show vtp password command to check whether the password is the same as that on the client.
If the VTP configuration is different from that on the client, perform the following operations to configure the VTP server.
VTP_Sever(config)# vtp domain Cisco   
VTP_Sever(config)# vtp mode server   
VTP_Sever(config)# vtp password Cisco   
Step 2 Check the Cisco VTP client configuration. The configurations of two VTP clients are the same. The following information is used for reference only.
! 
hostname VTP_Client 
! 
interface GigabitEthernet0/1 
 switchport access vlan 10 
 switchport trunk encapsulation dot1q 
 switchport mode access 
! 
interface GigabitEthernet0/2 
 switchport access vlan 10 
 switchport trunk encapsulation dot1q 
 switchport mode access 
! 
interface GigabitEthernet0/48 
 switchport trunk encapsulation dot1q 
 switchport mode trunk 
!
VTP_Client# configure terminal 
VTP_Client(config)# interface gigabitethernet 0/1 
VTP_Client(config-if)# switchport mode access   
VTP_Client(config-if)# switchport access vlan 10   
VTP_Client(config-if)# exit 
VTP_Client(config)# interface gigabitethernet 0/2 
VTP_Client(config-if)# switchport mode access 
VTP_Client(config-if)# switchport access vlan 10 
VTP_Client(config-if)# exit 
VTP_Client(config)# interface gigabitethernet 0/48   
VTP_Client(config-if)# switchport trunk encapsulation dot1q   
VTP_Client(config-if)# switchport mode trunk 
VTP_Client(config-if)# exit 
# Run the show vtp status command to check whether the VTP working mode is client and whether the domain name is the same as that on the server. Run the show vtp password command to check whether the password is the same as that on the server.
If the VTP configuration is different from that on the server, perform the following operations to configure the VTP client.
VTP_Client(config)# vtp domain Cisco   
VTP_Client(config)# vtp mode client   
VTP_Client(config)# vtp password Cisco   
Step 3 Configure the Huawei S series switch.
<HUAWEI> system-view 
[HUAWEI] vlan 10   
[HUAWEI-vlan10] quit 
[HUAWEI] interface GigabitEthernet1/0/48 
[HUAWEI-GigabitEthernet1/0/48] port link-type trunk   
[HUAWEI-GigabitEthernet1/0/48] port trunk allow-pass vlan 2 to 4094   
[HUAWEI-GigabitEthernet1/0/48] quit 
[HUAWEI] interface GigabitEthernet1/0/1 
[HUAWEI-GigabitEthernet1/0/1] port link-type access   
[HUAWEI-GigabitEthernet1/0/1] port default vlan 10   
[HUAWEI-GigabitEthernet1/0/1] quit 
Step 4 Verify the configuration.
Run the display vlan 10 command to check whether interfaces on the Huawei S series switch have been added to VLAN 10.
----End
When VTP is enabled on a Cisco switch to synchronize VLAN information, a Huawei S series switch with a downstream Cisco switch connected cannot process VTP packets. Therefore, the Huawei S series switch needs to transparently transmit VTP packets.
l When a Huawei S series switch sets up a Layer 2 tunnel to transparently transmit VTP packets, the destination multicast address must map to the unused multicast address to prevent address conflicts.
l The VTP tunnel must be set up on an interface of the Huawei S series switch in VLAN 1 where VTP packets are transmitted.
In Figure 1-7, a Huawei S series switch is directly connected to the Cisco VTP server and client. The Huawei S series switch needs to transparently transmit VTP packets to the Cisco VTP client, and user hosts need to communicate in VLAN 10.
Figure 1-7 Hybrid networking of the C-H-C model
1. Check the configuration of Cisco switches.
2. Configure Layer 2 transparent transmission on the Huawei S series switch to transparently transmit VTP packets.
3. Create a VLAN manually on the Huawei S series switch and add interfaces to the VLAN.
! 
hostname VTP_Sever 
! 
interface GigabitEthernet5/1 
 switchport trunk encapsulation dot1q 
 switchport mode trunk 
! 
interface GigabitEthernet5/3 
 switchport trunk encapsulation dot1q 
 switchport mode trunk 
!
VTP_Sever# configure terminal 
VTP_Sever(config)# interface gigabitethernet 5/1 
VTP_Sever(config-if)# switchport trunk encapsulation dot1q   
VTP_Sever(config-if)# switchport mode trunk   
VTP_Sever(config-if)# exit 
VTP_Sever(config)# interface gigabitethernet 5/3 
VTP_Sever(config-if)# switchport trunk encapsulation dot1q 
VTP_Sever(config-if)# switchport mode trunk 
VTP_Sever(config-if)# exit 
If the VTP configuration is incorrect, perform the following operations to configure the VTP server.
If the VTP configuration is incorrect, perform the following operations to configure the VTP client.
# Configure Layer 2 transparent transmission on the Huawei S series switch.
<HUAWEI> system-view 
[HUAWEI] l2protocol-tunnel vtp group-mac 0100-5e00-0011   
[HUAWEI] interface GigabitEthernet1/0/48 
[HUAWEI-GigabitEthernet1/0/48] l2protocol-tunnel vtp vlan 1   
[HUAWEI-GigabitEthernet1/0/48] quit 
[HUAWEI] interface GigabitEthernet1/0/46 
[HUAWEI-GigabitEthernet1/0/46] l2protocol-tunnel vtp vlan 1   
[HUAWEI-GigabitEthernet1/0/46] quit 
# Add interfaces on the Huawei S series
[HUAWEI] vlan 10   
[HUAWEI-vlan10] quit 
[HUAWEI] interface GigabitEthernet1/0/48 
[HUAWEI-GigabitEthernet1/0/48] port link-type trunk   
[HUAWEI-GigabitEthernet1/0/48] port trunk allow-pass vlan 2 to 4094   
[HUAWEI-GigabitEthernet1/0/48] quit 
[HUAWEI] interface GigabitEthernet1/0/46 
[HUAWEI-GigabitEthernet1/0/46] port link-type trunk    
[HUAWEI-GigabitEthernet1/0/46] port trunk allow-pass vlan 2 to 4094   
[HUAWEI-GigabitEthernet1/0/46] quit 
l Run the display l2protocol-tunnel group-mac vtp command to check the Layer 2 transparent transmission configuration on the Huawei S series switch.
l Run the display vlan 10 command to check whether interfaces on the Huawei S series switch have been added to VLAN 10.
Furthermore, the Huawei S series switch uses VCMP to synchronize VLAN information, and interfaces on the switch use Link-type Negotiation Protocol (LNP), reducing the configuration and maintenance workload.
