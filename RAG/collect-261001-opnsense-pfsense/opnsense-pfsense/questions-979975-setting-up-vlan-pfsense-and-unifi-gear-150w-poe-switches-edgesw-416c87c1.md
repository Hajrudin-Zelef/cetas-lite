---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/questions-979975-setting-up-vlan-pfsense-and-unifi-gear-150w-poe-switches-edgesw-416c87c1
title: "questions-979975-setting-up-vlan-pfsense-and-unifi-gear-150w-poe-switches-edgesw-416c87c1"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/questions-979975-setting-up-vlan-pfsense-and-unifi-gear-150w-poe-switches-edgesw-416c87c1.md
source_anchor: ""
source_lines: [1, 23]
sha256: b29711cf18693530997d7bbf7c8709e3b9ee471bdfba406be1e9212698789134
---

# questions-979975-setting-up-vlan-pfsense-and-unifi-gear-150w-poe-switches-edgesw-416c87c1

I'm have a tremendously difficult time setting up something that seems like it should be easy: VLANS with a pfSense gateway and a bunch of Ubiquiti UniFi gear. The frustrating thing is I've set this up before with a pfSense gateway and UniFi APs, but with a bunch of different switches.
Below I will to be as thorough as possible explaining the architecture of my organization’s Network and documenting our setup so someone could possibly see what I may have messed up.
1st, here is a JPEG of a VISO drawing I made showing the Network overview:
Important Info: I am able to SSH into the EdgeSwitch 16XG and ping the gateways of the VLANS (10.10.30.1 and 10.10.40.1). But if I try to login to the wireless networks using these VLANs no IP is given via DHCP on the pfSense appliance. Also, if I try to assign a port on any of my switches to only use a specific VLAN and then connect a laptop to that switch via RJ45 I do not get an IP (just APIPA address and ipconfig /release, ipconfig /renew does not help).
Moving on...
Here is all the configuration I have performed:
Interface assignments in pfSense. The Student Wifi will have VLAN 300, and the Guest Wifi will have VLAN 310. They are tied to the lan port (whcih you can physically see in the overview of the network above).
Screenshot of the settings within the Student Wifi interface for VLAN 300:
Screenshot of the settings within the Guest Wifi interface for VLAN 310:
Here is the screenshot for the VLANs under "Interface Assignments":
Screenshot for igb1 (lan) - VLAN 300 under "VLAN Interfaces":
Screenshot for Firewall Rules - VLAN 300 (Student Wifi):
Screenshot for Firewall Rules - VLAN 310 (Guest Wifi):
Screenshot of settings in the EdgeSwitch 16XG:
Screenshot of "Devices" in the Unifi Controller (o1/o3):
Screenshot of "Devices" in the Unifi Controller (o2/o3) - Looking at settings for one of the PoE switches:
Screenshot of "Devices" in the Unifi Controller (o3/o3) - All VLANs set to pass through the ports on this PoE switch which goes out to the APs:
Screenshot of UniFi Controller (Settings --> Networks):
Screenshot of UniFi Controller (Settings --> Networks, Editing VLAN 310):
Screenshot of UniFi Controller (Settings --> Wireless Networks):
Screenshot of UniFi Controller (Settings --> Wireless Networks, Editing VLAN 300):
Screenshot of UniFi Controller (Settings --> Wireless Networks, Editing VLAN 310):
Ok, that is all: Have I misconfigured something? What am I doing wrong? Why can't my devices get an IP address when connecting to the Wifi attributed with these two VLANs?
