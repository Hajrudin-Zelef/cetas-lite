---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/how-to-create-lag-with-vlans-on-proxmox-server-using-l2-switch-and-opnsense-2a883bff-1
title: "how-to-create-lag-with-vlans-on-proxmox-server-using-l2-switch-and-opnsense-2a883bff"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "research", "throughput"]
source: docs/RAG/collect-261001-opnsense-pfsense/how-to-create-lag-with-vlans-on-proxmox-server-using-l2-switch-and-opnsense-2a883bff.md
source_anchor: ""
source_lines: [1, 28]
sha256: dffa5d856d8f81e029768734d48c636f704dd7c9b4c4c3b539948ab4e616820c
---

# how-to-create-lag-with-vlans-on-proxmox-server-using-l2-switch-and-opnsense-2a883bff

How to Create a LAG with VLAN Tagging on Proxmox Server using a L2 Switch and OPNsense
Table of Contents
When I first decided to migrate from my Ubuntu server to Proxmox, I added a 4 port 1 Gbps NIC (affiliate link) in my Proxmox system to be used by my LXCs and VMs. The primary goal was to increase bandwidth for the services hosted on my server while allowing me to dedicate the network interface on the motherboard for Proxmox management. A secondary goal was to improve redundancy in case a network interface dies or the cable fails (redundancy may be a primary goal for you).
Note
With four 1 Gbps interfaces, you will not have a single 4 Gbps data stream but you may have multiple 1 Gbps streams which aggregate to a max of 4 Gbps.
My initial configuration of Proxmox had one bridge interface per physical interface with no VLANs configured in Proxmox. I had each physical interface on Proxmox assigned to a VLAN on my network switch. This configuration is pretty simple, but it only allowed me to use 1 VLAN per bridge. When migrating to Proxmox I did not want to take the time to learn all of the networking aspects of Proxmox. I like to take one step at a time when learning something new. I needed to get moved to Proxmox first and get comfortable with the new environment since it was a big departure from simply running an Ubuntu server (especially since I never had any prior hypervisor experience). While using 4 network interfaces individually is still better than using 1 interface for increasing throughput across multiple services/networks, this configuration was not very flexible.
You know the story goes… eventually your needs grow which forces you to learn something new in order to solve a particular issue. I have more than 4 VLANs on my network so there came a time when I wanted to host a service on Proxmox on a 5th VLAN. It is not practical for me to add a physical network interface for each new network I need to host a service since it does not scale. Also, we already have the technology to solve this problem: Link Aggregation Group (LAG) and VLAN tagging. I already knew this, of course, but I was not using making use of LAGs or VLAN tagging within Proxmox itself since I wanted to keep the initial Proxmox configuration as basic as possible until I had time to learn how.
After doing some initial research, I took the plunge. I have this irrational (or perhaps rational) fear that I am going to mess up my perfectly working server and have to spend hours and hours of troubleshooting to figure out how to fix it (or just revert it back to try again later). However, I was pleasantly surprised to find that I was able to migrate my existing network configuration over to using a LAG with VLAN tagging in less than 30 minutes! Part of that 30 minutes was spent verifying the connectivity of each LXC/VM (20+ containers/VMs). It was such a smooth process. I hope it will be smooth for you too if you are modifying your existing Proxmox server configuration. If you are starting fresh, it should be even easier since you can test as you go.
I thought I was going to have to create a “Linux bond”, a “Linux bridge”, and several “Linux VLAN"s in Proxmox as they are called in the user interface, but I found I only needed to create a bond for the LAG and a bridge which is configured to be VLAN aware. For each LXC/VM, I entered the VLAN ID on the assigned network interfaces to properly tag the network traffic. That is enough configuration in Proxmox for everything to work properly assuming you have your network switch and OPNsense configured properly. I did not need to create any “Linux VLAN"s (as they are called in the Proxmox web interface) that are associated to the bridge interface in Proxmox since I am not doing any virtual networking within Proxmox. If you are experienced with Proxmox, I suppose this paragraph is the TL;DR version of this guide!
Warning
Before you begin, make sure you are editing the network interfaces from an interface where you will not loose access to Proxmox should something go wrong. You may also use the console if you have direct access to the system if you are worried about losing remote access to your Proxmox server.
Physical Connections
Before continuing this guide, the picture below shows the basic physical connections which are being used. Note that I am using the motherboard Ethernet interface as the management interface to configure the 4 port NIC as a LAG. This allows me to make changes to the networking on Proxmox without losing access to the server.
Create VLANs in OPNsense to use with Network Switch
You may already have VLANs configured for your network for other devices, but if you do not, a quick rundown of the process is the following:
- After your physical interfaces are assigned, go to “Interfaces > Other Types > VLAN” to add your VLAN(s).
- Click the “+” button and select the parent interface, the VLAN tag (the ID number between 1-4094), the priority, and the description of the VLAN.
- After finishing saving your VLAN(s), click “Apply”.
- Go to the “Interfaces > Assignments” page to assign the VLANs to the desired parent interface (the physical interface on your OPNsense system). Be sure to give it a description so the interface does not display as “OPT1”, for example.
- On the “Interfaces > [VLAN]” page (where “[VLAN]” is the description you used when creating the VLAN interface), click the “Enable Interface” checkbox. Click “Static IPv4” for the “IPv4 Configuration Type” and enter the “IPv4 address” for the VLAN interface such as 192.168.10.1 for VLAN 10, for instance. Select “24” from the dropdown next to the IP address to indicate a/24 network. You may also configure IPv6 if you desire.
- Go to the “Services > DHCPv4 > [VLAN]” page to enable DHCP by clicking on the “Enable DHCP server on the VLAN interface” checkbox (substitute “VLAN” for your interface description). Enter the range such as 192.168.10.100 to192.168.10.200 for VLAN 10.
- Finally, you will need to create at least one firewall rule for your new VLAN. Otherwise all network traffic will be blocked. Navigate to the “Firewall > Rules > VLAN” page (once again substituting “VLAN” for your interface). I recommend creating one rule to allow DNS on the VLAN interface followed by a rule which blocks all private networks while allowing all other networks (which will allow Internet access) as the minimal set of rules. You can make them more restrictive once you get everything set up properly.
For more detailed information, see my guide on creating VLANs in OPNsense.
Configure the Network Switch
In my example, I am going to assume you have a smart/managed network switch attached to the physical interface where you have your VLANs assigned in OPNsense. I will be using a managed TP-Link switch (affiliate link) as an example so you may need to refer to your switch’s documentation on how to exactly configure the LAG on your particular switch.
Configure Trunk Port that is Connected to OPNsense
If you already have your network configured for VLANs with OPNsense, you may skip this step, but for completeness I will describe the basic process.
The network switch which is connected to the OPNsense interface where you assigned your VLANs will need to have a trunk port configured to allow VLAN traffic to pass through to OPNsense. This step is extremely important because if VLAN traffic cannot reach OPNsense through your network switch, you will not be able to make use of VLANs for your LXCs/VMs in Proxmox as described in this guide. You need to have VLANs working properly in your network before you can continue further.
