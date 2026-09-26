---
id: collect-260926-rattrapage/rattrapage/r-mikrotik-comments-15bufpi-for-those-running-vlans-on-their-switches-db29dd30-3
title: "2023-07-28 14:28:27 by RouterOS 7.10"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-rattrapage/ai-llm/r-mikrotik-comments-15bufpi-for-those-running-vlans-on-their-switches-db29dd30.md
source_anchor: ""
source_lines: [93, 170]
sha256: a978dd6e8a11fbf469ef4cb0ee2bb7d2400ea185946066fcfaa1e17bf1ed80d5
---

# 2023-07-28 14:28:27 by RouterOS 7.10

add bridge=bridge interface=Things-VL-120 pvid=120
add bridge=bridge interface=Misc-VL-130 pvid=130
add bridge=bridge interface=Guest-VL-140 pvid=140
/ip settings
set max-neighbor-entries=8192
/ipv6 settings
set disable-ipv6=yes
/interface bridge vlan
add bridge=bridge tagged=bridge,sfp1-Trunk-Permeant,sfp8-Trunk-Apeiron \
untagged=Users-VL-110,sfp3-NUC,sfp4,sfp5,sfp6-NAS-TS-664,sfp7 vlan-ids=\
110
add bridge=bridge tagged=bridge,sfp8-Trunk-Apeiron,sfp1-Trunk-Permeant \
untagged=Things-VL-120 vlan-ids=120
add bridge=bridge tagged=bridge,sfp1-Trunk-Permeant,sfp8-Trunk-Apeiron \
untagged=Misc-VL-130 vlan-ids=130
add bridge=bridge tagged=bridge,sfp8-Trunk-Apeiron,sfp1-Trunk-Permeant \
untagged=Guest-VL-140 vlan-ids=140
/interface ovpn-server server
set auth=sha1,md5
/ip dhcp-client
add comment="Akasha IP" interface=bridge
/ip dns
set servers=192.168.x.x
/ip service
set telnet disabled=yes
set ftp disabled=yes
set www address=192.168.x.0/24,192.168.x.0/24,10.x.x.0/24
set ssh address=192.168.x.0/24,192.168.x.0/24,10.x.x.0/24
set api disabled=yes
set winbox address=192.168.x.0/24,192.168.x.0/24,10.x.x.0/24
set api-ssl disabled=yes
/routing bfd configuration
add disabled=no
/system clock
set time-zone-name=Europe/London
/system identity
set name=Akasha
/system leds
set 0 disabled=yes
set 1 disabled=yes
set 4 disabled=yes
set 5 disabled=yes
set 10 disabled=yes
set 11 disabled=yes
set 14 disabled=yes
set 15 disabled=yes
/system leds settings
set all-leds-off=immediate
/system note
set show-at-login=no
/system ntp client
set enabled=yes
/system ntp client servers
add address=0.pool.ntp.org
add address=2.pool.ntp.org
/system routerboard settings
set boot-os=router-os
/system swos
set static-ip-address=192.168.x.x
Thank you!
This article Using RouterOS to VLAN your network - MikroTik helped me with this. It has config examples to look through.
Great thanks
It's still a work in progress as I continue changing things but this has been working great.
Edit: /system identitiy and below is not required just part of my export. Currently IP settings are missing as I'm directly wired to the switch, but will be added when I'm finished.
Depends on the switch model.
CRS1x/2x you have to do it via the Switch menu if you want HW offloaded vlans.
CRS3x you can do it bridge>vlans.
CRS1x/2x : https://help.mikrotik.com/docs/pages/viewpage.action?pageId=103841836#CRS1xx/2xxseriesswitchesexamples-VLAN
CRS3x and most RB: https://www.youtube.com/watch?v=YI0gPoCQDmM
Also: BUY a WOOBM. This is your anti lockout if you have a USB port on your switch.
My quick summary of VLANs configuration on Mikrotik:
/interface/bridge/port defines ingress behavior. The frame-types parameter determines if frames are tagged on ingress. The PVID parameter specifies the VLAN id added.
/interface/bridge/vlan defines egress behavior. Each bridge has a single vlan table. The table contains entries (rows) for each VID. Each VID can only appear once in a table. The entry (row) for a VID will contain a list of member ports that tagged and member ports that are untagged.
The /interface/bridge/* stuff configures L2 stuff, switch-level processing.
/interface/vlan exists to add layer 3 functionality on a vlan. For example IP addressing, DHCP, and IP routing on top of VLANs.
Oh god THANK YOU! You finally made sense.
VLANs in general are not that hard. Why does Mikrotik have to make to insanely unnecesarely artificially complicated?
This of a vlan another ethernet cable going inside the ethernet cable.
