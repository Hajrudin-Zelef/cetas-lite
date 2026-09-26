---
id: collect-260926-mikrotik/mikrotik/questions-23960-transparent-bridging-using-mikrotik-routerboard-630407e0
title: "questions-23960-transparent-bridging-using-mikrotik-routerboard-630407e0"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "research"]
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/questions-23960-transparent-bridging-using-mikrotik-routerboard-630407e0.md
source_anchor: ""
source_lines: [1, 59]
sha256: ef839e9641ae93759067670a19aba530d9f78b7e30fdbf3582ce87ab46b79592
---

# questions-23960-transparent-bridging-using-mikrotik-routerboard-630407e0

Network Engineering is part of Stack Overflow’s open communities: specialist spaces where curiosity is welcome, knowledge is shared freely, and the best answers rise to the top.
This question shows research effort; it is useful and clear
1
This question does not show any research effort; it is unclear or not useful
Save this question.
Show activity on this post.
I have a router and a MikroTik connected by a (long) ethernet cable that I'm trying to setup with switch-like functionality.
My understanding is that the correct way to do this is to create a bridge on the MikroTik that incorporates the port that it's plugged into (a normal LAN port) and the WLAN.
When I attempt this however, there are couple of problems I seem to face:
WLAN devices cannot communicate with devices which are connected to the router (or in fact, the router itself). I assume this is some kind of routing issue in that it doesn't know where to try and send the requests, but my understanding of bridging would imply that would be a non-issue. Attempting to ping the router sends back the following response:
92 bytes from router (192.168.88.1): Redirect Host(New addr: 192.168.0.1)
Vr HL TOS Len ID Flg off TTL Pro cks Src Dst
4 5 00 0054 db8d 0 0000 3f 01 c5cb 192.168.88.254 192.168.0.1
the MikroTik itself acts as a DHCP server rather than the main router. This is still the case after setting the DHCP client to attach to the bridge-local interface. Do I want to setup a DHCP relay? Setup a separate DHCP server?
If I connect the ethernet to the ether1-gateway port, I can successfully route through it (access the router etc) but then I'm not actually connecting the network together (which kind of defeats the point).
Any pointers as to where I'm going wrong would be gratefully received, this isn't something I normally do so stupid errors are not unlikely.
I'll assume you are using a RB951 series router, which has ether1-5 and wlan1.
If that's not the case just edit the script to fit the specifications of your router.
First of all, do a System > Reset configuration (check No default configuration).
To do it, upload it to the router and type "import filename" in a Terminal.
This will make a bridge between the wireless interface and all the MikroTik ethernet ports, carrying all traffic transparently. Your main router will still act as the gateway and will take care of NAT, firewall and routing.
If the router is used for all the VLANs, you have a few options:
The router needs to use 802.1Q trunking on its LAN port. It will need to be the gateway for each
subnet (have a gateway address for each subnet in the address range
of the subnet for which it is serving as the gateway). The link
between the MikroTik and the router needs to trunk all the VLANs.
You could configure the MikroTik as a layer-3 switch. This
essentially makes the MikroTik the router for the VLANs, and it will
need a gateway address for each subnet in the address for the subnet.
The link between the MikroTik and the router should be a routed
point-to-point link. The router will need to have routes to each
subnet which point to the MikroTik as the next hop.
You could configure the MikroTik as a typical WAN router running NAT
with the WAN port connecting to the router's LAN port. This
introduces double-NAT which may cause some problems.
I have a MikroTik RB941-2nD-TC "hAP lite" running version 6.37.3 (6.34.2 also worked), that I wanted to add to an existing home router to extend wifi into a blackspot. I labeled the MikroTik as ecwa_01 (ethernet-connected-wifi-ap), to distinguish it from the existing router. The house already had an ethernet cable from the routers location to the blackspot.
Because the default config does most of what I needed I just, (from linux command line)
sudo ip a a 192.168.88.254 dev eth0; ssh admin@192.168.88.1
/system identity set name=ecwa_01
/system package enable ipv6
/ip pool remove numbers=0
/ip dhcp-server remove [find where interface=bridge]
/interface bridge port add bridge=bridge interface=ether1
/ip address add address=192.168.88.3/24 interface=bridge
Then I connect port 1 (ether1 labeled internet) of ecwa_01 to the ethernet cable that leads back to the router.
There are a few other things that I set, such as a passphrase for the device and:
/ip service disable telnet
/ip service disable api
/ip service disable winbox
/ip service disable api-ssl
/ip service disable www
/ip service disable ftp
If the customer also wants to use the MikroTik via wifi then I use a slightly modified version from the answer by mcdnl90
/interface wireless
set ssid="ecwa_01" numbers=0
set wlan1 mode=ap-bridge wireless-protocol=802.11 band=2ghz-b/g/n frequency=auto ssid="ecwa_01" disabled=no
/interface wireless security-profiles
set [find default=yes] authentication-types=wpa-psk,wpa2-psk group-ciphers=tkip,aes-ccm unicast-ciphers=tkip,aes-ccm mode=dynamic-keys wpa-pre-shared-key="**Your WiFi password**" wpa2-pre-shared-key="**Your WiFi password**"
N.B. you can cut-n-paste most of this, but you have to change **Your WiFi password** and there may be a delay when you add ether1 to the bridge. (Which is why I usually do that last.)
