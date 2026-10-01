---
id: collect-261001-cisco/cisco/t5-switching-cisco-catalyst-3560-cx-series-pd-vlan-setup-not-working-m-p-5216847-7e84ec04-1
title: "t5-switching-cisco-catalyst-3560-cx-series-pd-vlan-setup-not-working-m-p-5216847-7e84ec04"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["license"]
source: docs/RAG/collect-261001-cisco/t5-switching-cisco-catalyst-3560-cx-series-pd-vlan-setup-not-working-m-p-5216847-7e84ec04.md
source_anchor: ""
source_lines: [1, 48]
sha256: fe3026dbd23554832c19ed7c3d9346689aef433fe735c5859c3d22a439299117
---

# t5-switching-cisco-catalyst-3560-cx-series-pd-vlan-setup-not-working-m-p-5216847-7e84ec04

- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
10-28-2024 12:50 PM - edited 10-30-2024 06:15 AM
Goal: Create 2 VLANs on a switch (for now, will expand to 4 - have 8 ports on the switch) for a test lab, not for production. Connect devices on these ports (direct connection) and #1. verify ping to VLAN IP, and other devices, and #2. verify ability to talk between devices on the same VLAN, and #3. verify ability to talk between devices on different VLANs. This switch will never connect to a router for Internet access, it will only act as a L3 Switch for a relatively small network
Device Info: Cisco Catalyst 3560-CX Series PD Cisco IOS Software Version 15.2(4)E8
show version
License level: ipbase
License Type: Default. No Valid License found
What has been done so far on this switch:
enable
config terminal
hostname no_internet_switch
vlan 10
interface Gi0/1 (and Gi0/2)
switchport mode access
switchport access vlan 10
interface vlan 10
ip address 192.168.10.1 subnet mask 255.255.255.0
<< Repeated the same seps for VLAN 20, with Gi0/3 and Gi0/4) >>
ip routing
show vlan
1   default                    active     Gi0/3, Gi0/4, Gi0/5, Gi0/6, Gi0/7, Gi0/8
10  vlan1                     active     Gi0/1, Gi0/2
20  vlan2                     active     Gi0/3, Gi0/4
show ip route
Gateway of last resort is not set
192.168.10.0/24 is variably subnetted, 2 subnets, 2 masks
C   192.168.10.0/24 is directly connected, Vlan 10
L   192.168.10.1/32 is directly connected, Vlan 10
C   192.168.20.0/24 is directly connected, Vlan 20
L   192.168.20.1/32 is directly connected, Vlan 20
show ip interface brief
Vlan1                        unassigned         Yes   unset      up                     down
Vlan10                      192.168.10.1      Yes   unset      up                     down
Vlan20                      192.168.20.1     Yes   unset      up                     down
Here is the problem:
When I got the switch, did not check everything, set this up, same VLAN ping worked. Inter VLAN communications did not work. During the course of debugging, did a factory reset and now nothing works. Can someone please point out whats going on ? I am fairly new to this area, and have seen lots of questions like this here, but none where there is a requirement to not have a Router. I have the Vlan IPs as the default gateway on the Linux devices I have connected on the switch. Any help is greatly appreciated.
Solved! Go to Solution.
- Labels:
- 
						
							
		
