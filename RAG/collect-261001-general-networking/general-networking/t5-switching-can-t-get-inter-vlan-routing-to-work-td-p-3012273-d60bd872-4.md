---
id: collect-261001-general-networking/general-networking/t5-switching-can-t-get-inter-vlan-routing-to-work-td-p-3012273-d60bd872-4
title: "t5-switching-can-t-get-inter-vlan-routing-to-work-td-p-3012273-d60bd872"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/t5-switching-can-t-get-inter-vlan-routing-to-work-td-p-3012273-d60bd872.md
source_anchor: ""
source_lines: [593, 800]
sha256: 79fcc89be51ec0b9ae8bf559cbbe0653733daf94b3fd001af7d5f0b40c5eb6d2
---

# t5-switching-can-t-get-inter-vlan-routing-to-work-td-p-3012273-d60bd872

			LAN Switching
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-05-2017 11:23 AM
Your configuration looks fine to me.
On sw2 can you ping 192.168.20.1 using 192.168.20.3 as the source IP (use an extended ping).
If you can suggest you look at PC settings.
Jon
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-05-2017 11:49 AM
Hello Jon, thanks for your reply.
The ping was successful!
SW2#ping
Protocol [ip]:
Target IP address: 192.168.20.1
Repeat count [5]:
Datagram size [100]:
Timeout in seconds [2]:
Extended commands [n]: y
Source address or interface: 192.168.20.3
Type of service [0]:
Set DF bit in IP header? [no]:
Validate reply data? [no]:
Data pattern [0xABCD]:
Loose, Strict, Record, Timestamp, Verbose[none]:
Sweep range of sizes [n]:
Type escape sequence to abort.
Sending 5, 100-byte ICMP Echos to 192.168.20.1, timeout is 2 seconds:
!!!!!
Success rate is 100 percent (5/5), round-trip min/avg/max = 1/201/1000 ms
SW2#
And the PC settings are as follows:
VLAN10/PC
192.168.10.10
255.255.255.0
192.168.10.1
VLAN20/PC
192.168.20.10
255.255.255.0
192.168.20.1
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-05-2017 12:01 PM
It looks like a problem with either the PCs or the ports on the switch that the PCs connect to.
Jon
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-05-2017 11:54 AM
As Jon noted there doesn't appear to be anything wrong with you config; I would however make a best practice suggestion that will simplify your life when troubleshooting and save clutter on your network:
There really is no need to create a SVI on each switch for every VLAN in the network. Create a management VLAN that is independent of your data networks. Assign only one VLAN in each switch, that of the management VLAN. This would include creating a management VLAN sub-interface on fa0/0 of the 1841 and each of the 3550s. Removing VLAN 10 & 20 from each of the 3550s.
When you get the trunk connected correctly you will be able to ping the management VLAN of each device. Once this happens everything else will fall into place. Here is a good configuration example using management VLANs on the network equipment.
http://www.cisco.com/c/en/us/support/docs/switches/catalyst-3750-series-switches/45002-intervlan3750-45002.html
Cheers,
Sam
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-05-2017 02:27 PM
Thanks for your reply.
That seems like a very good best practice suggestion, but I haven't understood it fully. For example, doesn't the port f0/1 on SW2 need to have the following config?
interface FastEthernet0/1
switchport access vlan 20
switchport mode access
So, should I remove the vlans that dont directly connect to the switches, and leave the ones that do?
And would I have to assign the management vlan to any ports?
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-05-2017 05:18 PM
You are correct in what you are saying; in that fa0/1 needs to be configured for access on VLAN 20. When you create the trunks between each device you can specify which VLANs are allowed through the trunk or if you do not specify switchport trunk allowed vlan command ALL VLANs are allowed through the trunk.
Any time you create interVLAN routing you have to have a Layer 3 device (router or L3 switch) that is assigned IP addresses for a given VLAN. Each VLAN must be assigned a Switch Virtual Interface (SVI). In your case this is the 1841, this must have a trunk interface for each VLAN you have throughout your network. The trunk between the 1841 and SW1 will carry whichever VLANs you have assigned in the 1841. Remember that trunking is a Layer 2 protocol and doesn't care what IP address is assigned to the SVI at this point. The only thing your trunk knows is that you currently have 2 VLANs (10 & 20, 3 if you add a management VLAN).
On the switch side of the trunk the only thing the switch sees are the VLANs that the router is sending, 10, 20 & management. In order to assign a client on the switch to a specific VLAN you simply issue the commands that you have on fa0/1 on SW1 & SW2. Of course you could have fa0/34 assigned to VLAN 10 and fa0/35 assigned to VLAN 20 without any problems. The trunk between your SW1 and SW2 simply mirrors what the router is sending to SW1 provided you allow the same VLANs through the trunk. One word of warning, when you assign to a switchport be sure that the VLAN is in the VLAN database. On the 3550 inter vlan data from the enable prompt and type add vlan x (x=vlan number). To review which VLANs are on a given switch issue the show vlan command.
The management VLAN is simply there for communication (telnet, SSH) to the individual switches. You do not need to assign any ports to the management VLAN. Simply make sure that the VLAN is in the vlan database. Once the SVI is created and the VLAN is in the database you should be able to ping each device in the management network. Finally assign the IP address of the management interface in the router as the gateway on both switches.
Regards,
Sam
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-05-2017 11:52 PM
Thanks for your clarification. I was under the impression, that if I gave a port access to a vlan, (e.g f0/1 on SW2) that the corresponding SVI was needed in the switch with a configured ip address. But instead, I should use the sub-interfaces f0/0.10 & f0/0.20 on the router as the vlan SVI's, am I correct in this assumption?
The design makes much more sense now that i understand it more. I will configure my network as soon as possible and let you know how everything works. Thanks again for taking the time to explain!
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-06-2017 02:58 AM
You are correct, using the example config referenced above you will see that the 2950 and 2948 have a sole VLAN (10). All other VLANs are defined on the L3 device, in this case the 3750 stack.
Regards,
Sam
- Mark as New
- Bookmark
- Subscribe
- Mute
- Subscribe to RSS Feed
- Permalink
- Report Inappropriate Content
03-06-2017 01:33 PM
Hello Sam,
I believe I have configured the network correctly, but I am still unable to ping between the hosts. I can now ping my default gateway, 192.168.10.1, from VLAN10/PC, but no other sub interfaces on the router.
SW1
SW1#show run
Building configuration...
Current configuration : 2415 bytes
!
version 12.1
no service pad
service timestamps debug uptime
service timestamps log uptime
no service password-encryption
!
hostname SW1
!
!
ip subnet-zero
!
!
spanning-tree mode pvst
spanning-tree extend system-id
!
!
!
!
interface FastEthernet0/1
switchport access vlan 10
switchport mode access
!
interface FastEthernet0/2
switchport trunk encapsulation dot1q
switchport trunk allowed vlan 1,10,20,99
switchport mode trunk
!
interface FastEthernet0/3
switchport trunk encapsulation dot1q
switchport trunk allowed vlan 1,10,20,99
switchport mode trunk
!
-----Omitted lines-----
!
interface Vlan1
no ip address
shutdown
!
interface Vlan99
ip address 192.168.99.3 255.255.255.0
!
ip default-gateway 192.168.99.1
ip classless
ip http server
!
!
line con 0
logging synchronous
line vty 0 4
logging synchronous
login
line vty 5 15
logging synchronous
login
!
!
End
SW1#show vlan br
VLAN Name Status Ports
---- -------------------------------- --------- -------------------------------
1 default active Fa0/4, Fa0/5, Fa0/6, Fa0/7
Fa0/8, Fa0/9, Fa0/10, Fa0/11
Fa0/12, Fa0/13, Fa0/14, Fa0/15
Fa0/16, Fa0/17, Fa0/18, Fa0/19
Fa0/20, Fa0/21, Fa0/22, Fa0/23
Fa0/24, Gi0/1, Gi0/2
10 VLAN10 active Fa0/1
99 MANAGEMENT active
