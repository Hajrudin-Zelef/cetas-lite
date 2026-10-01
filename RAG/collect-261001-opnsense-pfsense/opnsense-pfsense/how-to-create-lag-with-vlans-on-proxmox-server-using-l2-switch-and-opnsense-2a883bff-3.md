---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/how-to-create-lag-with-vlans-on-proxmox-server-using-l2-switch-and-opnsense-2a883bff-3
title: "how-to-create-lag-with-vlans-on-proxmox-server-using-l2-switch-and-opnsense-2a883bff"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-opnsense-pfsense/how-to-create-lag-with-vlans-on-proxmox-server-using-l2-switch-and-opnsense-2a883bff.md
source_anchor: ""
source_lines: [62, 75]
sha256: 84582f5d1c5ee40bfbb4a3b35d435b80cdd6d50860884309bf5e7896f86f4824
---

# how-to-create-lag-with-vlans-on-proxmox-server-using-l2-switch-and-opnsense-2a883bff

For the “Slaves” field, you need to enter each of your network interfaces with a space in between each interface. On my system, my interfaces are automatically named enp11s0f0 through enp11s0f3 so I entered enp11s0f0 enp11s0f1 enp11s0f2 enp11s0f3 as the slave interfaces.
The only other thing you need to select is the “Mode”. Choose LACP (802.3ad) since LACP is the LAG type being used on the switch. Click the “Create” button to create the bond.
Create a Linux Bridge Using the Linux Bond
To use any of the physical interfaces in Proxmox, a bridge is needed. On the “Network” page, click “Create” so you can choose “Linux Bridge”. By default the bridge “Name” will begin with vmbr followed by the next available number. If you have a dedicated interface for managing your Proxmox like I do (since I am using the motherboard Ethernet interface) and you have not created any other bridges yet, your bridge interface may default to vmbr1 as shown in the screenshot below.
Make sure you set the interface to be “VLAN aware” by checking the box so that VLANs can be used on the bridge. You will also need to enter bond0 in the “Bridged ports” box to use the bond that you just created. Click “Create” to create the bridge.
Apply Changes to Network Interface Configuration
Finally, you can make the Proxmox network changes persistent by clicking the “Apply Configuration” button. The Proxmox network interfaces will reload.
Go to your network switch’s LAG table to verify that the LAG is recognized on the switch similar to the screenshot below. The members (ports) will be listed when the LAG is configured properly. If you do not see 4 ports listed, check your physical connections to the switch. If you test the physical connections of the LAG, when you unplug one of the cables, you will see the number of ports in the LAG reduce by 1.
When you see that the status of the LAG is good on the switch, you know that the LAG was created successfully.
Assign LXC or VM to the Desired VLAN
You are now ready to use the new LAG and start assigning VLANs to all your LXCs/VMs. If you are creating a new LXC/VM, you will need to select vmbr1 (or whatever you named your bridge) and enter the VLAN ID when you are walking through the steps.
Otherwise, if you have existing LXCs/VMs configured in Proxmox, you will need to change the bridge interface (if you are consolidating bridged interfaces like me) of each LXC/VM and assign the appropriate VLANs.
Repeat this process for each of your LXCs/VMs. Once you have made the changes, you need to start or restart the LXCs/VMs.
If all goes well, you should be able to obtain an IP address from the range of IP addresses you designated in OPNsense in the DHCP settings!
