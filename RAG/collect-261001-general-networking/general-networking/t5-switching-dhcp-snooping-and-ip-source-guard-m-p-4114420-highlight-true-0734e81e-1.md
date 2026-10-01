---
id: collect-261001-general-networking/general-networking/t5-switching-dhcp-snooping-and-ip-source-guard-m-p-4114420-highlight-true-0734e81e-1
title: "t5-switching-dhcp-snooping-and-ip-source-guard-m-p-4114420-highlight-true-0734e81e"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-switching-dhcp-snooping-and-ip-source-guard-m-p-4114420-highlight-true-0734e81e.md
source_anchor: ""
source_lines: [1, 149]
sha256: 9752f0c779c54bcc4aaf1371574a5a64673ee839df1723020b8b5280eaa93eda
---

# t5-switching-dhcp-snooping-and-ip-source-guard-m-p-4114420-highlight-true-0734e81e

DHCP Snooping and IP Source Guard
						
					
					
				
			
		
	
			
	
	
	
	
	
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
07-05-2020 06:07 AM
Hi,
I'm trying to configure DHCP snooping with IP Source Guard.
My topology (attached ) contained 3 Switches , where I created 2 VLANs 100 and 200.
the network wroks fine I can reach pc on the same VLAN from an AS-1 to AS-2, However when I enable "IP Source Guard" nothing works even PCs on the same VLAN on the same Switch can't reach each other.
As far as I know IP source Guard wroks with DHCP snooping to verify the source IP address and it should be enabled on untrusted ports, if the source IP is not on DHCP snooping binding than the packet is dropped.
- My first thought was maybe it's not working between different Switch because each switch has it's own dhcp snooping table that why packet is dropped between packet coming from other switch. That why I tried with the same VLAN (PC1 and PC6) but it's the same , ping doesn't work.
Any more clarification about how "IP Source Guard" Works and why it's not working in my case?
- CS- Config
!
ip dhcp snooping vlan 100,200
no ip dhcp snooping information option
ip dhcp snooping
!
spanning-tree mode rapid-pvst
spanning-tree extend system-id
spanning-tree vlan 1,100,200 priority 0
!
interface GigabitEthernet0/1
switchport trunk allowed vlan 1,100,200
switchport trunk encapsulation dot1q
switchport mode trunk
switchport nonegotiate
media-type rj45
negotiation auto
!
interface GigabitEthernet0/2
switchport trunk allowed vlan 1,100,200
switchport trunk encapsulation dot1q
switchport mode trunk
switchport nonegotiate
media-type rj45
negotiation auto
! 
interface Vlan100
ip address 100.1.1.1 255.255.255.0
ip helper-address 100.1.1.100
!
interface Vlan200
ip address 200.1.1.1 255.255.255.0
ip helper-address 200.1.1.100
!
AS-1-Config
!
ip dhcp snooping vlan 100,200
no ip dhcp snooping information option
ip dhcp snooping
!
interface GigabitEthernet0/0
switchport trunk allowed vlan 1,100,200
switchport trunk encapsulation dot1q
switchport mode trunk
switchport nonegotiate
media-type rj45
negotiation auto
ip dhcp snooping trust
!
interface GigabitEthernet0/2
switchport trunk allowed vlan 1,100,200
switchport trunk encapsulation dot1q
switchport mode trunk
switchport nonegotiate
media-type rj45
negotiation auto
ip dhcp snooping trust
!
!
interface GigabitEthernet1/0
switchport access vlan 100
switchport mode access
switchport nonegotiate
media-type rj45
negotiation auto
ip verify source
!
interface GigabitEthernet1/1
switchport access vlan 200
switchport mode access
switchport nonegotiate
media-type rj45
negotiation auto
ip verify source
!
AS-2 Config
!
ip dhcp snooping vlan 100,200
no ip dhcp snooping information option
ip dhcp snooping
!
interface GigabitEthernet0/0
switchport trunk allowed vlan 1,100,200
switchport trunk encapsulation dot1q
switchport mode trunk
switchport nonegotiate
media-type rj45
negotiation auto
ip dhcp snooping trust
!
interface GigabitEthernet0/1
switchport trunk allowed vlan 1,100,200
switchport trunk encapsulation dot1q
switchport mode trunk
switchport nonegotiate
media-type rj45
negotiation auto
ip dhcp snooping trust
!
interface GigabitEthernet1/0
switchport access vlan 100
switchport mode access
switchport nonegotiate
media-type rj45
negotiation auto
ip verify source
!
interface GigabitEthernet1/1
switchport access vlan 200
switchport mode access
switchport nonegotiate
media-type rj45
negotiation auto
ip verify source
!
- Labels:
- 
						
							
		
