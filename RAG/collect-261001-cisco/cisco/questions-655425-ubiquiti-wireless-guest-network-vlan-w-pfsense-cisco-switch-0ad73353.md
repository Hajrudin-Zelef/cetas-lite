---
id: collect-261001-cisco/cisco/questions-655425-ubiquiti-wireless-guest-network-vlan-w-pfsense-cisco-switch-0ad73353
title: "questions-655425-ubiquiti-wireless-guest-network-vlan-w-pfsense-cisco-switch-0ad73353"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/questions-655425-ubiquiti-wireless-guest-network-vlan-w-pfsense-cisco-switch-0ad73353.md
source_anchor: ""
source_lines: [1, 27]
sha256: 86bf6ed18c682520ce4fe7c4e8a0df376a71acdc2ea133cccc2c83e84b998216
---

# questions-655425-ubiquiti-wireless-guest-network-vlan-w-pfsense-cisco-switch-0ad73353

I'm trying to setup a guest wireless network in an environment that has been humming along nicely for quite some time. The wireless runs on Ubiquiti UniFi APs.
I hope this is clear. If its not, feel free to ask questions.... I have a feeling this could be confusing.
The Cisco switch in question is a SG 200 series, and the firewall is pfSense. Ubiquiti UniFi APs are configured with 2 SSIDs ("NNH" and "NNH Guest")
Before I describe more of the typography, here's the situation:
- Employees receive DHCP just fine from the proper DHCP server
- Computers connected physically to VLAN 200 receive DHCP just fine from pfSense (which is that VLAN's DHCP correct DHCP server).
- Wireless clients on NNH receive DHCP fine
- Wireless clients on "NNH Guest" do NOT get a DHCP address, and can't connect
More Details
pfSense is running as the firewall and the DHCP server for the Guest network only (the employee / primary network has a Synology NAS as the DHCP server).
The network has these two subnets
- 10.1.10.0/24 = Employees
- 10.1.200.0/24 = Guests (tagged VLAN 200)
pfSense is configured as follows:
- 3 Interfaces (WAN on interface bce0, LAN on bce1, OPT1 on VLAN 200)
- VLAN 200 is setup to run on top of LAN
There is a computer lab which is physically connected to VLAN 200. The port on the Cisco switch that connects to the lab (which has its own switch) is setup as an Access Port on VLAN 200 (untagged).
All PCs in the computer lab (physically connected) get their correct DHCP address from pfSense (10.1.200.0/24).
I've configured the ports on the Cisco switch that connect directly to the Ubiquiti UniFi APs to be in "General" mode, in a VLAN membership of 200 tagged:
Additionally, the switch's Trunk port that is connected to pfSense is a member of VLAN 200:
Ubiquiti UniFi APs are configured as follows:
- 2 SSIDs
- "NNH Guest" is setup to "use VLAN 200"
To Summarize...
Wireless clients connecting to "NNH Guest" aren't receiving a DHCP address (10.1.200.0/24) when "NNH Guest" is configured on the Ubiquiti UniFi APs to use VLAN 200.
When I take VLAN 200 out of the equation (remove VLAN 200 from the SSID in the Ubiquiti settings), clients are able to connect to the Guest network, but they get an IP address from the Employee subnet, which is obviously what I'm trying to avoid.
How can I fix this?
