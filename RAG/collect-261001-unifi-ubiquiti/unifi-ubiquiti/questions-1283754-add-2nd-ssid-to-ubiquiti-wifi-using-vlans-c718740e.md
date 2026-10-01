---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-1283754-add-2nd-ssid-to-ubiquiti-wifi-using-vlans-c718740e
title: "questions-1283754-add-2nd-ssid-to-ubiquiti-wifi-using-vlans-c718740e"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-1283754-add-2nd-ssid-to-ubiquiti-wifi-using-vlans-c718740e.md
source_anchor: ""
source_lines: [1, 17]
sha256: 50a90ec0832ee37d057c82471609792fa7a22037ca316ec9032f8d8bae8a9fa6
---

# questions-1283754-add-2nd-ssid-to-ubiquiti-wifi-using-vlans-c718740e

I have a SonicWALL firewall with multiple VLANs.
ID  Network         Purpose
2   192.168.2.1/24  All Internal Traffic
3   192.168.3.1/24  All IoT Traffic
4   192.168.4.1/24  All WiFi Traffic
I am trying to create another VLAN for Guest WiFi. I've created VLAN 5 with 192.168.5.1/24.
I am using Cisco SG200 smart switches. Right now, my relevant ports are set up as:
Port  Mode    Membership
1     Trunk   1U, 2T, 3T, 4T
4     Access  4U
Port 4 then goes to my Ubiquiti UniFi access point. My UniFi has one SSID set up (MrPeanut). Its device IP is 192.168.4.2. I don't have any VLAN options set in the wireless network on the UniFi. (I also have the Network settings on the UniFi of 192.168.1.1/24 but I don't think that's being used?)
My Goal
I'm trying to add my VLAN 5 to the Ubiquiti UniFi using SSID "MrPeanutGuest". However, I can't quite get it to work. I've added MrPeanutGuest under Wireless Networks in my UniFi settings.
Questions
- Do I need to set the VLAN on MrPeanut as 4 and MrPeanutGuest as 5?
- Do I need to change my Cisco port VLANs so that port 4 includes VLAN 5? If so, I have to change it from Access to Trunk, correct?
- Is the AP Device IP of 192.168.4.2 going to be an issue?
