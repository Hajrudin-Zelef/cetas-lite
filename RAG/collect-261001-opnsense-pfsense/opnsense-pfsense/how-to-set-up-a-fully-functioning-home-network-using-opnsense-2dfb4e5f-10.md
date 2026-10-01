---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/how-to-set-up-a-fully-functioning-home-network-using-opnsense-2dfb4e5f-10
title: "how-to-set-up-a-fully-functioning-home-network-using-opnsense-2dfb4e5f"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/how-to-set-up-a-fully-functioning-home-network-using-opnsense-2dfb4e5f.md
source_anchor: ""
source_lines: [531, 569]
sha256: 5d2f38f968ed36eecbb9a866d549f15a4bfeef126f8100334819df2a9cb1205c
---

# how-to-set-up-a-fully-functioning-home-network-using-opnsense-2dfb4e5f

Click the “Add” button to create new VLANs. When the “VLAN Config” dialog box opens, you will see two main sections for “Untagged Ports” and “Tagged Ports”. This may be confusing if you are new to VLANs.
The “Untagged Ports” section is where you select the ports you wish to add to a particular VLAN for all of your wired devices. You can only have an “untagged port” assigned to one VLAN.
The “Tagged Ports” section is only used for ports that are connected to VLAN-aware devices such as routers, switches, wireless access points, and even servers (such as virtualization servers). “Tagged ports” can be assigned multiple VLANs unlike “untagged ports” so that multiple VLANs can pass through one port. Think of tagged ports as an aggregation of multiple VLANs. You will often see the term “trunk port” used to refer to the tagged ports.
VLAN 10 (DMZ)
Enter the “VLAN ID” of 10 for the DMZ network, which is the same value used for OPNsense for the DMZ. They must match. The “VLAN Name” should be “DMZ” so the description matches with the OPNsense VLAN configuration so there is less confusion when maintaining your network configuration.
The DMZ network contains the web server so select port 9 in the “Untagged Ports” section since that is where it is connected in the example network.
You will need to set port 2 and 4 as the “Tagged Ports”. It is more important to have LAG1 set as a tagged port as shown in the next step than setting the individual ports as tagged. I like to tag the individual ports “just in case” (you decide to remove the LAG in the future, for instance), but it is likely not necessary as long as you have the LAG configured.
Note that for the DMZ, you do not need to tag the wireless access point port on port 1 unless you plan to have wireless devices in your DMZ (and you configure the DMZ VLAN on your access points).
Click on the “LAGS” section on the “Tagged Ports” so that you can select LAG1. This step is important to ensure your VLAN traffic is allowed on the LAG interface.
On the “Port Config” section, select port 9 and set the PVID to 10 as the VLAN ID.
VLAN 20 (USER)
Enter the “VLAN ID” of 20 for the USER network, and the “VLAN Name” of “USER”. A PC is connected to the USER network on port 13 so it is selected in the “Untagged Ports” section. Select ports 1, 2, and 4 as the tagged ports include the ports connected to OPNsense as well as the wireless access point.
Include tagging the LAG1 interface to allow the USER network on the LAG.
On the “Port Config” section, select port 13 and set the PVID to 20 as the VLAN ID.
VLAN 30 (IOT)
Enter the “VLAN ID” of 30 for the IOT network, and the “VLAN Name” of “IOT”. A printer is connected to the IOT network on port 11 so it is selected in the “Untagged Ports” section. Select ports 1, 2, and 4 as the tagged ports include the ports connected to OPNsense as well as the wireless access point.
Include tagging the LAG1 interface to allow the IOT network on the LAG.
On the “Port Config” section, select port 11 and set the PVID to 30 as the VLAN ID.
VLAN 40 (GUEST)
For the GUEST network, you should select any unused ports to be on the GUEST network so that devices that are later plugged into the ports do not default to the management LAN network. This is very important if you are worried about users or anything potentially malicious from being connected to your most critical network.
Since I am used the untagged LAN network for management, it means every other port needs assigned to some VLAN to keep devices on the LAN. For this reason, that is why many users (especially for businesses/enterprises) prefer to have a dedicated management VLAN for their network infrastructure.
Even if you have your own dedicated VLAN network for your critical network infrastructure, I still think it is a good idea to default all of the unused ports either to a sinkhole network that has no network connectivity or to a more restrictive network such as the GUEST network. That way if anyone connects to your network, they are already isolated from your network. If you decide you need to change it later for a new device on your network, it is a simple configuration change on your switch.
Enter the “VLAN ID” of 40 for the GUEST network, and the “VLAN Name” of “GUEST”. I do not have any wired devices connected to the GUEST network in my example but all unused ports are selected in the “Untagged Ports” section. Select ports 1, 2, and 4 as the tagged ports include the ports connected to OPNsense as well as the wireless access point.
If you plan to have a dedicated wired device to access all of the management network’s web interfaces, be sure to leave one port available to connect another LAN device. Otherwise, you will not have access to the management interfaces unless you created the appropriate firewall rules (as mentioned in the firewall rules section above).
Include tagging the LAG1 interface to allow the GUEST network on the LAG.
On the “Port Config” section, select the same unused ports as earlier for the GUEST network and set the PVID to 40 as the VLAN ID.
VLAN 50 (IPCAM)
Enter the “VLAN ID” of 50 for the IPCAM network, and the “VLAN Name” of “IPCAM”. IP security cameras are connected to the IPCAM network on ports 3, 5, and 7 so it is selected in the “Untagged Ports” section. Select ports 1, 2, and 4 as the tagged ports include the ports connected to OPNsense as well as the wireless access point.
Include tagging the LAG1 interface to allow the IPCAM network on the LAG.
On the “Port Config” section, select ports 3, 5, and 7 to set the PVID to 50 as the VLAN ID.
VLAN 1 (LAN)
VLAN 1 is often used by managed switches as the default for all ports. Any port that is set to VLAN 1 is considered to be untagged and not belonging to any other VLAN (some switches may allow you to change this default behavior).
I saved this configuration until last because some switches may not let you remove ports from VLAN 1 until they are assigned to other VLANs since all ports must be assigned to a VLAN even if it is the default VLAN 1. The ports that are assigned to other VLANs may already be removed from VLAN 1, but I noticed with the newer TP-Link switches, the ports may still show as being assigned to untagged VLAN 1 even when assigned to another VLAN. Most likely the switch is ignoring VLAN 1 when it is assigned to other ports, but I like to fix that so that it easy to see which ports are assigned to other VLANs.
Select port 6, which is connected to the untagged LAN interface of OPNsense, and port 1, which is connected to the UniFi wireless access point. The UniFi AP needs to have the untagged VLAN 1 assigned so that the wireless AP can be managed from the LAN using the UniFi Controller software, which will reside on the LAN as well.
Please be sure to remove VLAN 1 from ports 2 and 4 because only tagged traffic will be on those ports used for the LAG (remember, to keep the untagged and tagged VLAN traffic more separated in OPNsense to reduce the likelihood of accidentally creating rules that could allow access between untagged and tagged networks, which could be a real security concern).
One thing that you should change on the “LAGS” area of the “Untagged Ports” section is to deselect the LAG1 as shown in the screenshot below. The reason is that the LAG should only carry tagged traffic so the tagged and untagged traffic can be separated on the OPNsense interfaces.
To further ensure only tagged traffic is allowed on the LAG, go to the “Port Config” section, select ports 2 and 4, and set the “Acceptable Frame Types” to “Tagged Only”.
You can do the same thing for the LAG itself by clicking on the “LAGS” tab.
Connect switch to OPNsense and the AP to the Switch
