---
id: collect-260926-rattrapage/rattrapage/r-mikrotik-comments-15bufpi-for-those-running-vlans-on-their-switches-db29dd30-2
title: "2023-07-28 14:28:27 by RouterOS 7.10"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2023-07-28"]
keywords: ["compute", "ethernet", "voice"]
source: docs/RAG/lot-rattrapage/ai-llm/r-mikrotik-comments-15bufpi-for-those-running-vlans-on-their-switches-db29dd30.md
source_anchor: ""
source_lines: [3, 92]
sha256: dd750e631529ac15ff589ef550b4ba2ad56f0c39b5fc3a7cbffdc645ddcc18ce
---

# 2023-07-28 14:28:27 by RouterOS 7.10

    Please would someone running successful VLANS on their switches be kind enough to export their config here (Without sensitive data of course) in order for me to wrap my head around VLAN filtering in bridges.
No matter what config I try, I just cannot grasp the concept.
In the past I've used bridging interfaces to VLAN interfaces but I know it's not the right way and CPU intensive.
Specifically looking for those with VLAN filtering in the the bridge.
Thanks in advance.
Section des commentaires
The thing is that not all VLAN configuration methods are suitable for all devices. This is not clear in tutorials, including the recent one by Mikrotik
What is also not clear, is that if you abstract your bridges differently than the physical chips, you can get weird routing and bottlenecks. Also, if you have multiple chips, each one should be treated like a separate switch/bridge. Mixing ports will result in CPU load, unless you bridge them with an ethernet cable and configure them accordingly as separate switches. RSTP is also not supported by some chips, forcing offloading to switch off.
You need to check the VLAN switching support page and follow the one suitable for your device and chipsets.
https://help.mikrotik.com/docs/display/ROS/Basic+VLAN+switching#BasicVLANswitching-Otherdeviceswithabuilt-inswitchchip
https://help.mikrotik.com/docs/display/ROS/Bridging+and+Switching#BridgingandSwitching-BridgeHardwareOffloading
https://help.mikrotik.com/docs/display/ROS/Switch+Chip+Features
Edit: That might be more relevant for routers, but it's worth noting for any device that uses hardware acceleration
Thanks for this, something to also consider.
are you sure about that? i think that was the case years ago, but in recent years i don't think i've ever seen a device that supports switching but can't offload with spanning-tree running.
Here's a stripped-down configuration from my CRS328 that includes only a couple of relevant interfaces:
/interface bridgeadd name=interfaceSwitch vlan-filtering=yes/interface bridge portadd bridge=interfaceSwitch interface=ether1add bridge=interfaceSwitch interface=ether2 pvid=172/interface bridge vlanadd bridge=interfaceSwitch comment=vlanData tagged=ether1 untagged=ether2 vlan-ids=172/ip addressadd address=192.0.2.1/25 interface=interfaceSwitch
I've got ether1 as a trunk carrying the default VLAN (1, native) and the data VLAN (172, tagged); ether2 is an access interface with the PVID and native VLAN set to the data VLAN (172). The switch has L3 reachability on the default VLAN via 192.0.2.1 but is only running VLAN 172 with L2 switching.
If we want L3 switching and want to apply an IP address to the data VLAN, interfaceSwitch needs to be added to the data VLAN's tagged interface list. Then we can add a VLAN interface and apply an address to it.
/interface vlanadd interface=interfaceSwitch name=interfaceVlan172 vlan-id=172/ip addressadd address=192.0.2.129/25 interface=interfaceVlan172
Clear as mud, but I hope that helps.
Edit: Couldn't get the formatting to take.
It's a start, much appreciated.
192.0.2.129 is just a placeholder here? For documentation reasons?
All bridge vlan filtering does is apply the following logic: "Is the vlan tag on this packet in my /interface/bridge/vlan list? If yes, forward. If no, drop"
The port PVID (port vlan id) setting kinda does the inverse and says "if I see an untagged packet coming in to this port, add this vlan tag"
The /interface/bridge/vlan list then allows mutliple vlan tags to ingress or egress through specified ports. Untagged ports get the tag for packets from that vlan stripped, so they pair well with a port with a matching PVID.
Once you get the hang of it, it's not too terrible. I'll be the first to admit it can be a steep learning curve though.
Fucking thank you. I work with Cisco and HP devices and the way Mikrotik does VLANs did not compute. Out of all the documentation and explanations I’ve found, yours is the simplest and clearest.
Yeah def a weird learning curve. I struggled a lot at first, but once you get the hang of it, it's as cantanko describes here.
It really took a day with my switch, a console cable (port not on bridge) and continiously nuking the config with MikroTik's wiki open (CRS328 switching/bridging/l3 section) to fully wrap my head around it. Now that I did that - it's so damn simple. You just gotta trust in the bridge lol
What makes it extra complicated is the hardware chips and limitations that have to be respected in your configuration and bridge abstractions. If you violate chip boundaries, you will get weird routing and bottlenecks
thank you!
this is so simple,it should be stickied on mikrotik wiki...you forget !one! port in config and the day is gone,searching forums and wikis...
MikroTik actually just released some pro informative videos on the subject on their own YouTube channel!
It is hard to wrap around, but let's do a walk through on how I do mine. Let's make some assumptions:
- You'll be using these as hybrid ports (PVIDs for the ports because PCs, but VLANs tagged for things like Voice.)
- We're doing a Router on a Stick configuration.
- One port is a dedicated management port for the OH SHIT reasons.
- Three VLANs: LAN, Voice, Management (we'll use IDs 10, 50, 99)
- Trunk port is ether1, and we'll pretend it is a 48port switch. The OH SHIT port is #48. Port 2 is going to have an AP, like hAP ac2 attached.
Name things as you wish, I tend to use a structured name scheme such as: <Itemtype><Desc/VLAN><Attachment> For example, the first line below, Int for Interfaces, the VLANID, and which physical interface it is attached to.
From here, you just add ports to the Interface List for the PVID, and they will add to the bridge and be assigned that PVID as untagged, then add whichever tags you need in the Bridge/VLAN section.
I like the organized thinking here. I'll apply to the rest of my life as well!
Config below.
Connected network devices: Apeiron - RB5009, Permeant - CSS610.
4 VLAN's exist on the network but currently only one gets tagged on this switch.
# 2023-07-28 14:28:27 by RouterOS 7.10
# model = CRS309-1G-8S+
/interface bridge
add admin-mac=xxxxxxxxxx auto-mac=no dhcp-snooping=yes \
ingress-filtering=no name=bridge vlan-filtering=yes
/interface ethernet
set [ find default-name=sfp-sfpplus1 ] name=sfp1-Trunk-Permeant
set [ find default-name=sfp-sfpplus2 ] name=sfp2
set [ find default-name=sfp-sfpplus3 ] name=sfp3-NUC
set [ find default-name=sfp-sfpplus4 ] name=sfp4
set [ find default-name=sfp-sfpplus5 ] name=sfp5
set [ find default-name=sfp-sfpplus6 ] name=sfp6-NAS-TS-664
set [ find default-name=sfp-sfpplus7 ] name=sfp7
set [ find default-name=sfp-sfpplus8 ] name=sfp8-Trunk-Apeiron
/interface vlan
add interface=bridge name=Guest-VL-140 vlan-id=140
add interface=bridge name=Misc-VL-130 vlan-id=130
add interface=bridge name=Things-VL-120 vlan-id=120
add interface=bridge name=Users-VL-110 vlan-id=110
/interface lte apn
set [ find default=yes ] ip-type=ipv4 use-network-apn=no
/interface wireless security-profiles
set [ find default=yes ] supplicant-identity=MikroTik
/port
set 0 name=serial0
/interface bridge port
add bridge=bridge comment=defconf ingress-filtering=no interface=ether1
add bridge=bridge comment=defconf ingress-filtering=no interface=\
sfp1-Trunk-Permeant trusted=yes
add bridge=bridge comment=defconf ingress-filtering=no interface=sfp2
add bridge=bridge comment=defconf ingress-filtering=no interface=sfp3-NUC \
pvid=110
add bridge=bridge comment=defconf ingress-filtering=no interface=sfp4 pvid=\
110
add bridge=bridge comment=defconf ingress-filtering=no interface=sfp5 pvid=\
110
add bridge=bridge comment=defconf ingress-filtering=no interface=\
sfp6-NAS-TS-664 pvid=110
add bridge=bridge comment=defconf ingress-filtering=no interface=sfp7 pvid=\
110
add bridge=bridge comment=defconf ingress-filtering=no interface=\
sfp8-Trunk-Apeiron trusted=yes
add bridge=bridge interface=Users-VL-110 pvid=110
