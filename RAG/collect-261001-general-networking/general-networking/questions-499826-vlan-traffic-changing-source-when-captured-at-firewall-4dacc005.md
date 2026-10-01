---
id: collect-261001-general-networking/general-networking/questions-499826-vlan-traffic-changing-source-when-captured-at-firewall-4dacc005
title: "questions-499826-vlan-traffic-changing-source-when-captured-at-firewall-4dacc005"
domain: general-networking
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/questions-499826-vlan-traffic-changing-source-when-captured-at-firewall-4dacc005.md
source_anchor: ""
source_lines: [1, 75]
sha256: ea4ace4e84141a7aca6fae7416a1c9741d478ef9b74b2215d1a45b36dd324ac1
---

# questions-499826-vlan-traffic-changing-source-when-captured-at-firewall-4dacc005

1) HP Switch config below 2) Fortinet Policy in attached image
Right now, we are a flat network of roughly 320 wireless devices, and about 100 wired devices. We have a FortiGate 300C firewall with a single internet connection, and a single internal (10.1.100.106) connection with an HP zl 5406 hanging off the FG.
Part of my summer plans is to segment off traffic into VLANS, so I am messing around with configuring that now. But, I have an extremely weird issue that I cannot figure out, and Fortinet is blaming on the HP switch.
Flat/Default Vlan 1 Network: 10.1.0.0/16 Vlan 20: 10.20.0.0/16
I have my workstation on the default Vlan. I have a laptop plugged into an untagged port for Vlan 20, and it is obtaining an IP from DHCP just as it should, from the proper 10.20.0.0 scope. On the laptop, I have 4 infinite pings going. 8.8.8.8 (Google DNS), 10.1.100.106 (FortiGate firewall), 10.1.10.15 (DHCP Server), and 10.1.10.250 (My workstation). On my workstation, I have one going for the IP of the laptop (10.20.1.50).
Google/Firewall/DHCP will be going great on the laptop, but my workstation cannot ping the laptop (Times out), but randomly, for 5-10 minutes, the Firewall will time out (But DHCP/Google keeps going) and on my workstation, the ping to the laptop will kick off like it is supposed to. Then out of nowhere, it goes back to the original performance. The way I have the policy on the firewall configured, neither of them are functioning properly. They should both be able to communicate both ways.
This has been driving me nuts for weeks. Captured packets on the FG from the laptop. When Google/Gateway/DHCP are pinging properly, the source shows as the MAC address of the laptop. When it is failing, the source is the MAC address of the switch. This has baffled me and Fortinet support.
One thing I figured out today is when I go on the HP switch and "clear arp" while Google/Gateway/DHCP are working, it will cause the firewall ping to time out and the ping from my workstation to the laptop will start working temporarily.
Now for some curveballs/other info:
1) This is my second attempt at this Vlan. I originally tried configuring Vlan 200 and it function similarily, but when the firewall would time out, so would the Google ping. Now, It is only the firewall.
2) I have a Vlan 30 that has my office printer on it. It functions just fine and acts nothing like this.
3) Enabled STP today and made no change.
4) Enabled IGMP today and made no change.
5) When my workstation can ping the laptop. in the Vlan, the laptops cannot ping the firewall. When my workstation cannot ping the laptops, the laptops CAN ping the firewall.
HP Config:
Laptop is port F1. Firewall is A1. Servers are B1-B16. Printer on Vlan 30 is C1.
; J8697A Configuration Editor; Created on release #K.15.08.0013
; Ver #02:1b.ef:f6
hostname "SGS-MDF-SW01"
module 1 type j9534a
module 2 type j9536a
module 3 type j9534a
module 4 type j9536a
module 6 type j9536a
cdp mode pass-through
timesync sntp
sntp unicast
sntp 30
sntp server priority 1 10.1.100.100
time daylight-time-rule continental-us-and-canada
time timezone -360
ip route 0.0.0.0 0.0.0.0 10.1.100.106
ip routing
snmp-server contact "Brandon" location "Middle School - 1st Floor - MDF"
vlan 1
   name "DEFAULT_VLAN"
   no untagged C1,F1
   untagged A1-A24,B1-B22,C2-C24,D1-D22,F2-F22
   ip address 10.1.100.151 255.255.0.0
   ip igmp
   exit
vlan 20
   name "Phones"
   untagged F1
   tagged A1,B1-B16
   ip address 10.20.1.1 255.255.0.0
   ip helper-address 10.1.10.15
   exit
vlan 30
   name "Printers"
   untagged C1
   tagged A1,B1-16
   ip address 10.3.1.1 255.255.0.0
   ip helper-address 10.1.10.15
   exit
vlan 40
   name "LS_Lan"
   ip address 10.40.1.1 255.255.0.0
   ip helper-address 10.1.10.15
   exit
vlan 50
   name "MS_LAN"
   ip address 10.50.1.1 255.255.0.0
   ip helper-address 10.1.10.15
   exit
vlan 60
   name "Wireless"
   ip address 10.60.1.1 255.255.0.0
   ip helper-address 10.1.10.15
   exit
vlan 80
   name "Imaging"
   ip address 10.80.1.1 255.255.0.0
   ip helper-address 10.1.10.15
   exit
