---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/how-to-set-up-a-fully-functioning-home-network-using-opnsense-2dfb4e5f-11
title: "how-to-set-up-a-fully-functioning-home-network-using-opnsense-2dfb4e5f"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "parameters"]
source: docs/RAG/collect-261001-opnsense-pfsense/how-to-set-up-a-fully-functioning-home-network-using-opnsense-2dfb4e5f.md
source_anchor: ""
source_lines: [570, 603]
sha256: a738242e4ce9866d18bd72598b03b8d5ad62228641b98957c7029145de75817e
---

# how-to-set-up-a-fully-functioning-home-network-using-opnsense-2dfb4e5f

After the switch has been configured, it is time to plug it into OPNsense to see if the LAG and VLAN configuration was successful! Plug the LAN interface of OPNsense into port 6 on the switch and the two interfaces of the LAG on OPNsense into ports 2 and 4 on the switch.
You can tell if the LAG is configured properly by viewing the “LAG Table” on the “LAG” page under “L2 Features”. If you see both port 2 and port 4 as “Members” of the LAG, you know it is functioning properly since it has detected both ports in the LAG configuration. Both ends of the LAG must be configured in order for the LAG to show all of the connected members.
Also plug the UniFi AP into port 1 before proceeding with the configuration in the next section.
Note
As everything is currently configured, you will need to make sure you have either a dedicated device such as something simple like a Raspberry Pi plugged into your LAN so you can access your OPNsense or network switch’s interface. To do that you need to configure one of your switch’s ports to be untagged (VLAN 1) so the device will be on the same network as your network infrastructure.
Alternatively, you could create a firewall rule on the USER interface to allow your PC/laptop access to the web interfaces. However, that opens one hole into the management network. If you do not have any spare devices or you do not want to have to temporarily plug your PC/laptop into the LAN, a firewall rule would be the most convenient option.
Configure Wireless Access Points
Last but not least, the wireless access point needs to be configured. When using a firewall appliance such as the Protectli VP2410, I recommend using external wireless access points rather than using a USB or built-in wireless module. WiFi performance will be much greater and more reliable when using dedicated wireless access point(s) – in particular, APs which have a wired backhaul connection.
A wired backhaul means the wireless AP has an Ethernet cable that is plugged either directly into your router or into a switch that is plugged into your router. Many wireless APs can be powered via PoE (Power over Ethernet) in which case you will need to plug it into a PoE network switch or a PoE injector.
Not all wireless access points support creating VLANs but APs such as UniFi support this feature which means you can add your wireless devices to the same VLANs as your wired devices. I often see others put their WiFi on a separate subnet/VLAN, but that is not necessary (unless you really want your wireless devices completely isolated from your wired devices). For me personally, having my wireless devices on some of the same networks as my wired devices is very convenient.
Because OPNsense and the network switch have been configured for VLANs, all that remains is the VLAN configuration of the AP. There are other parameters you may tweak as well to improve wireless performance, but I am going to focus on the VLAN configuration to get the devices connected. I will leave WiFi performance tuning as an exercise for the reader (perhaps a future topic idea).
I am going to assume the UniFi Controller is already installed and is running on the LAN. By default, the UniFi Controller should be able to find the AP which is plugged into port 1 since port 1 has been configured to allow untagged VLAN 1 traffic as well as VLANs 20, 30, 40, and 50. It not easy for me to demonstrate these steps without having an unused or new AP that I can provision using the UniFi Controller software.
Since the UniFi APs only allow for 4 VLANs on each AP (1 VLAN per SSID), I will demonstrate setting up 4 VLANs.
Tip
Keep in mind that if your wireless spectrum is crowded by having several close neighbors with WiFi, you may wish to limit the number of VLANs because the UniFi APs use separate radios for each VLAN. I have read that performance on UniFi APs can start to diminish when using more than 3 SSIDs but that could depend on how crowded the spectrum is and how many devices are connected to your WiFi network(s).
Create New VLAN Networks
The first thing you need to do is create the VLAN networks. Go to the “Settings” page by clicking the gear button at the bottom of the left side menu and then click on the “Networks” section. By default you will see a “Default” network with the 192.168.1.1 address range. You can leave this default network. Click on “Create New Network” to create a VLAN.
Enter the “Network Name” of “USER (20)” and the “VLAN ID” of 20, which is the same VLAN ID used for the OPNsense/switch VLAN configuration.
Repeat this step with the other VLANs using the following values:
| Network Name | VLAN ID | 
|---|---|
| IOT (30) | 30 | 
| GUEST (40) | 40 | 
| IPCAM (50) | 50 | 
Create New WiFi Networks
Once the networks are added, you can create new WiFi SSIDs – 1 SSID for each VLAN. Click on the “WiFi” section above the “Networks” section. Then you can click on “Create New WiFi”.
Enter the “Name” of your WiFi network. This network will be visible to anyone with range of your network so I recommend using a more descriptive name than the name of your VLANs. What I like to do is use a common base name and then append the VLAN to SSID so I can tell the networks apart from each other.
Next add the “Password” and select the “Network” that you just created (such as “USER (20)”, etc).
Repeat this process for the other 3 WiFi networks. All of your AP(s) should be provisioned with the new SSIDs and you will be able to connect to them with your wireless devices.
Next Steps
At this point you should have a fully functioning home network with several VLANs to isolate groups of devices on your home network for improved security.
You may want to explore other areas where you can add more features or add more layers of security. I have written another guide for several security related features that you may reference.
I hope this guide helps demonstrate the steps involved with setting up the network infrastructure for a more advanced home network. My goal was to keep the topic constrained as much as possible while also providing a comprehensive example to help you get started with your home network.
If you found this guide to be too complex to begin your journey, please check out my beginner’s guide on creating a full network, which has a greatly simplified network architecture from which you can build upon later as your knowledge and experience grows.
