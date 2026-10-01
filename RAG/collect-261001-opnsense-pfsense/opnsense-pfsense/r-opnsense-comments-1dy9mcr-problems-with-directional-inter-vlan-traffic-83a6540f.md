---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/r-opnsense-comments-1dy9mcr-problems-with-directional-inter-vlan-traffic-83a6540f
title: "r-opnsense-comments-1dy9mcr-problems-with-directional-inter-vlan-traffic-83a6540f"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/r-opnsense-comments-1dy9mcr-problems-with-directional-inter-vlan-traffic-83a6540f.md
source_anchor: ""
source_lines: [1, 66]
sha256: 382d5e4dc5a555a1c2fdf840edb1ef5e620665fe4ccbac06b6ff66f25f90e53b
---

# 
       Problems with directional inter VLAN traffic 

      
    Status Quo:

    I have a Fujitsu Futro S920 as OPNSense router behind a cable modem, and there I configured 5 VLANs, VLAN20, VLAN30 etc. to VLAN60. all in the `192.168.` range with a subnet mask of `/24` and the third octet is the same as the VLAN ID which is also part of the name of the VLAN. So VLAN20 has VLAN ID 20 and the network is `192.168.20.0/24`. The VLAN Interfaces are all on the same physical Interface called LAN01 which is the `192.168.10.0/24` network. On every Interface is DHCP configured and active. Important for my question are LAN01, VLAN20 and VLAN50. Here are the firewall rules for these:
  


I also have the TL-SG108PE switch behind it. It is in LAN01 with a static IP adress. On port 3 is an AP, on port 4 is a server and on port 1 is my PC. Port 2 & 7 are empty. The other are not important. On the switch I configured the following:


    The server on port 4 gets an IP address via DHCP in the correct VLAN, VLAN50 therefore `192.168.50.X`, and the same is true for my PC, VLAN20 therefore `192.168.20.X`.
  


The problem:

Even though I have a firewall rule which allow traffic from VLAN20 in VLAN50, I can't ping or ssh into the server, from my PC. But I can ping the Interface IP in VLAN50, from my PC. Also as you can see above there is a firewall rule to allow traffic from VLAN20 in LAN01, and I can ping the switch and AP (which are on LAN01) and also visit there web interface, this works. Why this but not the other?


Any help appreciated

BR

Maybe the server has additional settings blocking access from another subnet?

It is a Raspberry Pi 4B, with Ubuntu 22.04. I doubt it, because I configured nothing except auto updates

Hi,

    you could either do serious debugging by packet tracing:

when trying to SSH from VLAN20 to VLAN50 please do the following:

* Enable logging for the VLAN20 -> VLAN50 pass rule and check in firewall's live log if there's any SSH traffic going through

* If there's no traffic: Run a packet capture on VLAN20 to see if the packets even arrive at OPNsense

* If there's SSH traffic: Run a packet capture on VLAN50 to see if the packets returned from the server are correctly adressed or missing
  

    **Or** you could check why port 1 is untagged on several VLANs at the same time and what would change if it would be untagged on VLAN 20 only (the PVID). Honestly, inconsistent tagging/pvid assignment ist usually the root cause for such a mess.
  

Hi,

    Or you could check why port 1 is untagged on several VLANs at the same time and what would change if it would be untagged on VLAN 20 only (the PVID). Honestly, inconsistent tagging/pvid assignment ist usually the root cause for such a mess.


The thing is I know it is kind of odd, but the real untagged setting is the PVID setting, at least this is what I gathered as reading the Manual and different posts to this switch. And if I remove the ports from the default VLAN (1) or switch them to untagged the switch webinterface is unreachable and unpingable. I also thought before posting it is something with the switch but every post I read has the same configuration but not my problems

Edit: Sometimes it is direct unreachable, then ok I know was my bad I configured something for sure wrong, but other times only after a reboot

There is no real or unreal untagged setting. There is an untagged setting and no port shall have more than one untagged VLAN. The connected device would not be able to distinguish traffic from the different networks.

Unfortunately, TPLink choses the worst way of configuration where PVID can be set independently. Well, the PVID can be thought of the VLAN ID used when untagged traffic comes in. If there's only one untagged VLAN it is clear that PVID and the VLAN ID are the same.

In order to avoid connection loos I would recommend to use VLAN 1 as management VLAN: Keep one port untagged + pvid 1, feel free to also put it as vlan to OPNsense in order to access it whenever you want. The untagged port is kind of a fallback if everything goes down. Don't forget to test. :-)

Regarding your problem: Ports 1,4,5,6 are untagged in VLAN 1 but have a different PVID. Cross check this with OPNsense: The physical connection port shall know all VLAN IDs and those should be available as tagged on the switch"s port.

Set PVID to match the desired untagged VLAN on the port and remove them from untagged VLAN 1.

Not sure what you mean, since the PVID is already set to match the untagged VLAN ID 50. And why does the changing of the default matter? I don#t have these problems with my PC. But I will try it
