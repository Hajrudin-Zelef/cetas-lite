---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/how-to-beginners-guide-to-set-up-home-network-using-opnsense-62d9c4eb-5
title: "how-to-beginners-guide-to-set-up-home-network-using-opnsense-62d9c4eb"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "parameters"]
source: docs/RAG/collect-261001-opnsense-pfsense/how-to-beginners-guide-to-set-up-home-network-using-opnsense-62d9c4eb.md
source_anchor: ""
source_lines: [236, 270]
sha256: 76947e8ca53cd1a70e19e8e09fe234693bfbcc7f1f811965c253d587d9eaf13f
---

# how-to-beginners-guide-to-set-up-home-network-using-opnsense-62d9c4eb

I am going to make the assumption in this guide that you have a new switch that has not been configured or one that you factory reset to the default values.
When configuring the network switch, I recommend that you plug a computer directly into the switch so that you can get it set up before you plug it into OPNsense to avoid any potential issues such as having a static IP address on the switch which conflicts with the interfaces you have configured in OPNsense.
Connect Directly to Switch
To avoid losing connectivity when configuring VLANs on your network switch, you could plug into the same port on your switch that will eventually be plugged into OPNsense because the port will be set up to allow both untagged and tagged VLAN traffic for all of your networks. In our example, you would plug into port 1.
Once you are plugged into your switch, you will need to determine if your switch has DHCP enabled by default. The easiest way to know is to look at your network status on whatever device you are using to configure the switch to see if you have an IP address assigned. If you do not have an IP address, you will need to manually enter your IP address. However, you will need to consult the manual to see what the default IP address of your switch is set to.
Since I am using a TP-Link managed switch for this example, the default IP address is 192.168.0.1. In this case, manually setting the IP address of the device you are using to configure the switch to 192.168.0.10 with a subnet mask of 255.255.255.0 will be sufficient.
The TP-Link switches have a web interface and newer models can be configured using the Omada Software Controller. The Omada software is not required to configure the switch, so I am going to simply use the web interface. If you can access http://192.168.0.1 or https://192.168.0.1 successfully after configuring the static IP address then you are good to go.
The default username of the TP-Link switch is admin and the password is admin, but you will need to consult the user manual of the switch you are using. Of course, I recommend changing your password after you sign in.
Change the Switch’s Interface IP Address
The first thing you should do is change the IP address of the network switch to a static IP address that resides on the LAN so you can access it later to make changes. In my example, I will set it to 192.168.1.2 so the switch will be included on the LAN network since the router is using 192.168.1.1/24 for the LAN interface. This step is important so you do not lose access to your network switch.
Click on the “System” tab at the top of the page, and then “System IP” on the left side menu. You should see the default IP address of 192.168.0.1. Enter 192.168.1.2 for the “IP Address”. You may also enter 192.168.1.1 for the “Default Gateway”.
Click “Apply”. You will likely lose connectivity immediately and will have to start using the http://192.168.1.2 web address. You will need to change your device’s IP address to be 192.168.1.10 so it can be on the same subnet as the network switch.
Be sure to log back in and click on the “Save” button in the upper right hand corner to make sure the changes are persistent once you know you are able to connect to the new IP address. If you do not click “Save” on TP-Link switches, the changes will be lost when you reboot the switch.
Physical Diagram of Connected Devices
Before proceeding with the VLAN configuration, refer to the following physical diagram for devices/clients connected to the network (see also the physical network infrastructure diagram). It will be beneficial to see which Ethernet ports the devices are plugged into when configuring the switch.
VLAN Configuration
Go to the “L2 Features” page and click on the “VLAN > 802.1Q VLAN” left side menu to see the list of VLANs. By default, every port is assigned VLAN1 which is designated for untagged ports. The default behavior of a managed switch is exactly like an unmanaged switch so everything is on the same flat network.
Click the “Add” button to create new VLANs.
VLAN 10 (UNTRUSTED)
When the “VLAN Config” dialog box opens, you will see two main sections for “Untagged Ports” and “Tagged Ports”. This may be confusing if you are new to VLANs.
The “Untagged Ports” section is where you select the port(s) you wish to add to a particular VLAN for all of your wired devices. You can only have an “untagged port” assigned to a single VLAN.
The “Tagged Ports” section is only used for ports that are connected to VLAN-aware devices such as routers, switches, wireless access points, and even servers (such as virtualization servers). “Tagged ports” can be assigned multiple VLANs unlike “untagged ports” so that multiple VLANs can pass through one port. Think of tagged ports as an aggregation of multiple VLANs. You will often see the term “trunk port” used to refer to the tagged ports.
Enter the “VLAN ID” of 10 for the UNTRUSTED network, and the “VLAN Name” of “UNTRUSTED”. An IoT device such as a smart TV is connected to the UNTRUSTED network on port 5 so it is selected in the “Untagged Ports” section. Select ports 1 and 2 as the tagged ports include the ports connected to OPNsense as well as the wireless access point.
On the “Port Config” section, select port 5 and set the PVID to 10 as the VLAN ID. This step is important. Otherwise, the port will not be properly tagged with the desired VLAN ID.
VLAN 1 (LAN)
The trusted LAN will be on the default VLAN 1 so we do not need to make any changes for the trusted devices which are connected to ports 3 and 4 in the diagram above since those ports are already in the default VLAN 1.
Connect switch to OPNsense and the AP to the Switch
After the switch has been configured, it is time to plug it into OPNsense to see if the VLAN configuration was successful! Plug the LAN interface of OPNsense into port 1 on the switch.
For a quick test of the VLAN, try plugging your device into port 5 and check your device’s IP address. If you receive an IP address in the 192.168.10.x network, then your configuration is working properly!
You may also plug the Grandstream (or other) wireless access point into port 2 before proceeding with the configuration in the next section.
Configure Wireless Access Points
The last device which needs configured is the wireless access point. When using a firewall mini-PC such as the Protectli VP2420, I recommend using external wireless access points rather than using a USB or built-in wireless module on the mini-PC. WiFi performance will be much greater and more reliable when using dedicated wireless access point(s). The APs I prefer to use have a wired backhaul connection.
A wired backhaul means the wireless AP has an Ethernet cable that is plugged either directly into your router or into a switch that is plugged into your router. Many wireless APs can be powered via PoE (Power over Ethernet) in which case you will need to plug it into a PoE network switch or use a PoE injector. Some APs have a barrel jack to plug it into AC power which might be helpful for some users depending on the location of your AP.
Not all wireless access points support creating VLANs but APs such as the Grandstream I am using in this example support this feature. When APs support VLANs, you have the ability to add your wireless devices to the same VLANs as your wired devices, which is convenient since it allows you to group untrusted wired and wireless devices together, for example.
Because OPNsense and the network switch have been configured with a VLAN, all that remains is the VLAN configuration of the AP. There are other parameters you may tweak as well to improve wireless performance, but I am going to focus on the VLAN configuration to get the devices connected. I will leave WiFi performance tuning as an exercise for the reader.
