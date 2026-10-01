---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/how-to-create-lag-with-vlans-on-proxmox-server-using-l2-switch-and-opnsense-2a883bff-2
title: "how-to-create-lag-with-vlans-on-proxmox-server-using-l2-switch-and-opnsense-2a883bff"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-opnsense-pfsense/how-to-create-lag-with-vlans-on-proxmox-server-using-l2-switch-and-opnsense-2a883bff.md
source_anchor: ""
source_lines: [29, 61]
sha256: 20eed13682991fc9be62448cfcc93881b3e0b292df94d6bc74c004438574e5a9
---

# how-to-create-lag-with-vlans-on-proxmox-server-using-l2-switch-and-opnsense-2a883bff

On my switch, I have my OPNsense router connected to port 16 as shown below. The PVID should be set to 1 for the trunk port rather than setting it to a specific VLAN ID like other ports when you are adding a device to a specific VLAN. On the newer TP-Link models, you can set the “Acceptable Frame Types” to Admit All (to allow both tagged and untagged traffic) or “Tagged Only” (if you want to allow the untagged VLAN 1 traffic to pass through). On older models, you had to set the “Link Type” to TRUNK. You do not have to do that extra step on the new models if you simply leave it on the default “Admit All” option.
Create your desired VLANs using port 16 as the tagged interface since that will be the VLAN trunk to OPNsense. Every VLAN you create on the switch will need to have port 16 selected as tagged if you want the VLAN traffic to pass through the OPNsense interface.
In my example below, I am creating a DMZ VLAN. The VLAN ID needs to be the same ID you use in the OPNsense VLAN configuration. You may select any ports of connected devices that you want to be in the DMZ (ports 9 and 10, for example) in the “Untagged Ports” section. Repeat this process for any other VLANs you wish to create.
Create the LAG to be used by Proxmox
Once you have the VLANs set up in OPNsense and your switch configured properly, next you will need to configure the LAG to be used by Proxmox. It is ok to set up the LAG on the switch before you configure it on Proxmox because when you are using LACP, the LAG will not become active until the other end (Proxmox) has been configured with LACP.
Go to the “L2 Features” page and click on “LAG” then “LACP Config”. Select all of the interfaces that are connected to your Proxmox server and set the status to Enabled to enable the LAG for those those interfaces.
The “Group ID” is the LAG number. Since I have a LAG configured already to another switch, I am setting my Proxmox LAG as LAG2, but you may use a different number.
The “Port Priority” defaults to 32768 but you can set it to a different value. Lower numbers indicate a higher priority. If you set them all to the same value, they will have the same priority. You could set them to different priorities so if one fails it will use the next highest priority. To be honest, I do not know why in this scenario the port priority would be important within a LAG because if it switches to another port when it dies or the cable goes bad, it does not matter to me which port is used as long as the rest of the LAG still functions properly. There may be valid reasons that I am not aware of.
For the “Mode”, set it to Active so that it negotiates the LAG via LACP. At least one side of your LAG needs to be set to “Active”.
Be sure to click “Apply” to create the LAG.
Note
If you click on the “LAG Table” tab after applying your changes, you will notice that the members (ports) of the LAG will not be listed yet since Proxmox has not been configured to use the LAG. Once Proxmox is configured, you should see the 4 ports show up as members of the LAG.
If you need more information on how to create a LAG on your switch, you can view my guide on how to create a LAG.
Assign the Appropriate VLANs to the LAG for Proxmox
Once the LAG is configured, you need to assign VLANs to the LAG similar to how you assign VLANS to other ports/interfaces on your switch. I am assuming you have existing VLANs set up on your switch (as described earlier) so you will need to go to the “VLAN Config” page to edit your existing VLANs to now include the LAG interface that was just created. Continuing the example from earlier, I will be editing the DMZ VLAN.
The LAG interface selection for VLANs is on a separate tab within each “Untagged/Tagged Ports” section. Click the “LAGS” tab in the “Tagged Ports” section to select LAG2 to assign the DMZ VLAN to the LAG. Then click the “Save” button.
Repeat this process for each VLAN you wish to use on the Proxmox LAG.
If you only want to allow tagged network traffic on the Proxmox LAG, select Tagged Only for the “Acceptable Frame Types”. Otherwise, you can allow untagged traffic (VLAN 1 in particular) by leaving it as Admit All.
Warning
Setting the “Acceptable Frame Types” to the proper value may be important for you if you use the default untagged VLAN 1 as your management VLAN. If you forget the set the VLAN ID when creating a new LXC or VM, it will be placed in your management VLAN which may be a potential security risk depending what you are running. That is why it is often recommended to use a dedicated tagged VLAN for your management VLAN instead of using the default untagged VLAN 1.
Also, there is the possibility of someone plugging into an unused Ethernet port and get access your management VLAN if you leave your used ports as untagged on VLAN 1. Granted, this is probably less of a serious issue on a home network unless you have untrusted tech savvy users connecting to your network who wish to do you harm.
Persist Changes on your Switch
Your switch may require you to click a “Save” changes button that will make the configuration persistent across reboots. Otherwise, you will lose all your changes the next time you power up your switch.
The reason switch configuration is often designed this way is to allow you to test changes on your network with less risk. If you messed something up really bad, you could simply reboot your switch to revert back to your previous configuration, which was (hopefully) functioning properly. It is important that you only persist changes only when you know everything is working correctly.
On newer TP-Link switches, you only need to click the “Save” button in the upper right hand corner of the screen to persist all of the configuration changes you have made.
Configure the Proxmox Network Interfaces
With the switch configured, it is time to configure Proxmox. To minimize downtime if you are modifying an existing Proxmox instance, you may leave your LXCs/VMs running while you make the network configuration changes because it will not take effect until you click the “Apply Configuration” in Proxmox. However, after you apply your changes, you will need to reboot each service after making changes to the network interfaces.
Warning
Please consider the consequences when proceeding with the following changes if you are modifying an existing Proxmox instance rather than starting from scratch. I was able to take such a risk because I am making this change on my home network, but if you are doing something like this in a production environment, you should exercise greater caution.
Many guides I have seen show how to manually edit the /etc/network/interfaces file directly in order to make network changes in Proxmox. If you are more comfortable with making such changes, feel free to do so, but the web interface is also very easy to use especially if you are not familiar with the syntax of Linux network interfaces. You could copy/paste the network configuration that you find in such guides, but I am going to show the web interface method of how to update the network configuration, which I think will benefit newer users to Linux/Proxmox.
If you are interested in learning more about the syntax of the /etc/network/interfaces file, you can see what the web interface generates after you change your network configuration.
Create a Linux Bond
Go to your Proxmox server on the web interface and select “Network” under the “System” sidebar menu. Click on “Create” and choose “Linux Bond”. You will need to enter the “Name” of the bond interface. By default it will be bond0 if you do not have any other bonds configured.
