---
id: collect-261001-general-networking/general-networking/questions-42359-assign-wireless-devices-to-vlan-by-mac-address-e79be70a
title: "questions-42359-assign-wireless-devices-to-vlan-by-mac-address-e79be70a"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/questions-42359-assign-wireless-devices-to-vlan-by-mac-address-e79be70a.md
source_anchor: ""
source_lines: [1, 80]
sha256: 0d21076accd23f556e99236d7171b98b167020bc63d0878c9c18b61fe78baa4c
---

# questions-42359-assign-wireless-devices-to-vlan-by-mac-address-e79be70a

Answer
I wound up realizing that MAC-based VLANs isn't really what I wanted, and that Group-based VLANs are actually much more flexible.  From here on out, the information will regard to accomplishing VLAN assignment based on group membership.  
Bench Test
I followed this really great guide published by Meraki.  It is 99% generic so if you're not a Meraki shop don't worry.  There were just 2-3 settings to change in the Meraki and you can probably translate that into your system easily enough.
- Windows Server 2012 R2 (in VMware Workstation)
  - Domain Controller
  - DHCP
    - 192.168.3.1 /24 (primary scope)
    - 192.168.4.1 /24 (scope will be used for VLAN 400)
    - 192.168.5.1 /24 (scope will be used for VLAN 500)
    - Option 3 (Router) set to point to Aruba switch (192.198.3.6)
  - DNS (standard setup)
  - Certificate Services 
    - (I followed this guide.  Skip the part where they have you test revoking a cert.  It made a bit of extra work)
  - NAP
  - Connected to port 1 on the switch
- L3 Aruba 2920-24G Switch
  - Configured for routing inside test network (but not out to the Internet)
  - Don't forget IP helpers
  - sho config listed below for reference
- Cisco Meraki MR18 WAP
  - Connected to port 3 on the switch
  - Configured to use RADIUS (link in comments...I'm new to this stackexchange network and I can't post more than 2 in the body)
- 2x Lenovo ThinkPad 11e Laptops (wireless clients to test with)
  - Joined to domain
  - Named "WIRELESSLAPTOP" and "WIRELESSLAPTOP2"
- AD Structure
  - OU: "Test Machines" containing both laptops
  - OU: "VLAN Assignment Groups"
    - Group: "VLAN 400" with member "WIRELESSLAPTOP2"
    - Group: "VLAN 500" with member "WIRELESSLAPTOP"
Result
Following the guide and then expanding a little, I wound up with three network policies.  There is a VLAN 400 Policy which injects the VLANID 400 into the RADIUS-ACCEPT packet, and there's a VLAN 500 Policy which works the same way.  There is also a required default policy of sorts (explained in the guide).  Each of these policies has a "Conditions" list, and that is where you assign that policy to a group.  Just stick your computers in that group and they'll be good to go.
Now when I connect these two laptops to the network, one lands on the 400 VLAN and gets a 4.x address, and the other lands on the 500 VLAN and gets a 5.x address.  These assignments can be changed as easily as group membership.  
I believe this will be a pretty robust solution for us moving forward.  Note to the reader, this is a complex solution so while the manageability is nice, be sure it's something you'll need and use before introducing several new layers of complexity into your environment.
sho config placed here for reference
; J9727A Configuration Editor; Created on release #WB.16.03.0004
; Ver #10:08.3f.f3.b8.ee.34.79.3c.29.eb.9f.fc.f3.ff.37.ef:86
hostname "HP-2920-24G-PoEP"
module 1 type j9727a
gvrp
ip route 0.0.0.0 0.0.0.0 10.1.30.1 (config for routing to the production 
network)
ip routing
snmp-server community <removed> unrestricted
oobm
   ip address dhcp-bootp
   exit
vlan 1
   name "DEFAULT_VLAN"
   no untagged 5,15
   untagged 1-4,6-14,16-24
   ip address dhcp-bootp
   exit
vlan 300
   name "aovlan"
   untagged 5 (this is an uplink to our production environment.  Not currently 
in use)
   ip address 10.1.30.100 255.255.255.0
   exit
vlan 400
   name "TestA"
   tagged 3 (this is where the AP lives)
   ip address 192.168.4.1 255.255.255.0
   ip helper-address 192.168.3.1
   exit
vlan 500
   name "TestB"
   untagged 15 (port 15 was used in previous testing...not really important 
here)
   tagged 3 (this is where the AP lives)
   ip address 192.168.5.1 255.255.255.0
   ip helper-address 192.168.3.1
   exit
device-profile name "default-ap-profile"
   cos 0
   exit
activate software-update disable
activate provision disable
----------
