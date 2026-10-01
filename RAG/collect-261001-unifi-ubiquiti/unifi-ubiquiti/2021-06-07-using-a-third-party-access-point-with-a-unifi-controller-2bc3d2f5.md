---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/2021-06-07-using-a-third-party-access-point-with-a-unifi-controller-2bc3d2f5
title: "2021-06-07-using-a-third-party-access-point-with-a-unifi-controller-2bc3d2f5"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-unifi-ubiquiti/2021-06-07-using-a-third-party-access-point-with-a-unifi-controller-2bc3d2f5.md
source_anchor: ""
source_lines: [1, 20]
sha256: 95dead47084c2c9b057b4d17732c008f4b93c20c03a703add13dd34919bede73
---

# 2021-06-07-using-a-third-party-access-point-with-a-unifi-controller-2bc3d2f5

In this guide I will be going over how to use a third-party access point with a router running the Unifi controller. By using VLAN tagging to assign wireless devices their own network.
My Hardware
- Third-party Access Point
  - For my third-party access point, I will be using the Asus RT-AX88U
- Unifi Controller
  - For a Unifi controller, I will be using Ubiquiti’s Dream Machine Pro (UDM Pro)
Start off by setting up both devices according to their manuals. After setup is complete connect the AP to the UDM Pro through an ethernet cable. Make sure you set up the access point in access point mode, like my access point most can also be set up as routers. Do not do this or your devices will not appear on the controller, only the router will appear. In router mode, clients connected directly to the UDM Pro will not be able to connect to devices behind the router’s firewall. In access point mode the device will let the UDM Pro do all the routing for its wireless clients allowing the controller to see all the devices.
Network
A network is needed in order to start assigning IP addresses to your new wireless devices. To create a network first, go to your controller’s portal. You can use the gateway IP or the Unifi Portal and click on your controller’s network settings. Next click settings > networks > add new network. Give the network a name, and select content filtering if needed. Under advance set a VLAN ID, gateway IP/subnet, DHCP range, and DNS server. Here are my settings to give you some ideas.
| Name | WiFi | 
| VLAN ID | 20 | 
| Gateway IP/Subnet | 10.10.0.1/24 | 
| DHCP Range | 10.10.0.100 – 10.10.0.200 | 
| DNS Server 1 | 8.8.8.8 | 
VLAN
A virtual local area network is needed when connecting the UDM Pro to a third-party access point otherwise issues will arise where the AP loses internet connection. Here are the steps to create the VLAN.
First, go to your controller’s portal. You can use the gateway IP or the Unifi Portal and click on your controller’s network settings.
Under Unifi devices click on your UDM Pro or other controller and select the ports tab. Select the port that you connected your AP to. If you used an unmanaged switch somewhere in the middle of your ethernet connection this may cause issues and mislabel other devices as being on your AP network. For me, I connect the AP directly to port 8. I’ll select port 8 and change the port profile to the name of the network I created specifically for this AP. This will resolve the issue of the AP losing network connectivity. By allowing the controller to properly assign the wireless devices a network.
Firewall
If you would like to put a firewall between your new wireless devices and the rest of your networks this section is for you. While in the Network Unifi Portal as before, go to settings > security > firewall. Select the LAN tab and click create new rule. Make sure the type is set to LAN in and the action to drop. For source > source type, select network and then choose your AP network for the IPv4 subnet. For destination choose network and select the network you want to prevent your wireless devices from accessing. Click apply changes when you are done. If you also want to prevent the destination network from talking to your wireless devices you will need to create another rule reversing the source network and destination network. You will also need to repeat theses steps for however many networks you want to block your AP’s devices from connecting too.
