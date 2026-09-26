---
id: collect-260926-rattrapage/rattrapage/questions-665833-mikrotik-rb2011u-vlanning-29582c2a
title: "questions-665833-mikrotik-rb2011u-vlanning-29582c2a"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-rattrapage/servers-reviews/questions-665833-mikrotik-rb2011u-vlanning-29582c2a.md
source_anchor: ""
source_lines: [1, 38]
sha256: 4f0768ec7720349d13626e059bd6f2a250d93bcbdae0523babb445d4d20f3cbc
---

# questions-665833-mikrotik-rb2011u-vlanning-29582c2a

I am having trouble getting port based vlanning/trunking working on a Mikrotik RB2011U. I have been following the guide located here: http://wiki.mikrotik.com/wiki/Manual:Switch_Chip_Features#Example_-_802.1Q_Trunking_with_Atheros_switch_chip_in_RouterOS_v6 The features table at the top of the wiki indicates that this device should be capable of 802.1q vlanning.
I have followed the wiki verbatim, except that I have supplied my own interface names and vlan id numbers. When I plug a host into eth3 and assign it a static IP of 10.10.10.4 I am unable to ping 10.10.10.5 (the RB2011U); however plugging a second host into eth4 and assigning it a static IP of 10.10.10.6, I am able to ping 10.10.10.4 from 10.10.10.6 and vice versa. On the 10.10.10.4 and 10.10.10.6 side I can see that when I try to ping 10.10.10.5 I am getting the an ARP reply from the RB2011U.
The second problem is that my trunk port does not seem to be working properly either. It is wired back to a Mikrotik CRS125 which I have confirmed to be working properly. In this case, pinging the RB2011U(10.10.10.5) from the CRS125(10.10.10.2) and vice virsa also does not work, however ARP does not seem to be being relayed.
In searching for answers, I see a lot of other people using bridges instead of using the switch configuration. And I did manage to get that configuration working minus the trunk port; however I configured my CRS125 to do vlan tagging in the switch config and for consistency's sake I'd prefer to have the RB2011U do the same.
EDIT: (requested configuration information)
/interface ethernet
set [ find default-name=ether1 ] name=eth1
set [ find default-name=ether2 ] name=eth2
set [ find default-name=ether3 ] master-port=eth2 name=eth3
set [ find default-name=ether4 ] master-port=eth2 name=eth4
set [ find default-name=ether5 ] master-port=eth2 name=eth5
set [ find default-name=ether6 ] name=eth6
set [ find default-name=ether7 ] master-port=eth6 name=eth7
set [ find default-name=ether8 ] master-port=eth6 name=eth8
set [ find default-name=ether9 ] master-port=eth6 name=eth9
set [ find default-name=ether10 ] master-port=eth6 name=eth10
/interface ethernet switch port
set 2 vlan-header=add-if-missing vlan-mode=secure
set 3 default-vlan-id=1 vlan-header=always-strip vlan-mode=secure
set 4 default-vlan-id=1 vlan-header=always-strip vlan-mode=secure
set 5 default-vlan-id=0 vlan-header=always-strip vlan-mode=secure
set 11 vlan-mode=secure
/interface ethernet switch vlan
add independent-learning=yes ports=eth2,eth3,eth4,switch1-cpu switch=switch1 vlan-id=1
add independent-learning=yes ports=eth2 switch=switch1 vlan-id=3
add independent-learning=yes ports=eth2 switch=switch1 vlan-id=4
add independent-learning=yes ports=eth2 switch=switch1 vlan-id=5
/interface vlan
add interface=eth2 l2mtu=1594 name=vlan1 vlan-id=1
add interface=eth2 l2mtu=1594 name=vlan3 vlan-id=3
add interface=eth2 l2mtu=1594 name=vlan4 vlan-id=4
add interface=eth2 l2mtu=1594 name=vlan5 vlan-id=5
/ip address
add address=10.10.10.5/24 interface=vlan1 network=10.10.10.0
add address=10.30.10.5/24 interface=vlan3 network=10.30.10.0
add address=10.40.10.5/24 interface=vlan4 network=10.40.10.0
add address=10.50.10.5/24 interface=vlan5 network=10.50.10.0
/interface ethernet exportalso/interface vlan exportand/ip address exporteth2is your trunk port, correct? What do you mean when you mention When I plug directly into one of the access ports of a particular vlan and try to ping the IP address I assigned to the corresponding vlan interface I get no reply; ? Can you please edit your post to reflect your configuration? For example 'When I connect a device toeth2I ping the IPx.x.x.x...'. It hard to understand what you have tried and what you connected where now.
