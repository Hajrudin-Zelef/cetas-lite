---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/how-to-set-up-a-fully-functioning-home-network-using-opnsense-2dfb4e5f-9
title: "how-to-set-up-a-fully-functioning-home-network-using-opnsense-2dfb4e5f"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-opnsense-pfsense/how-to-set-up-a-fully-functioning-home-network-using-opnsense-2dfb4e5f.md
source_anchor: ""
source_lines: [499, 530]
sha256: 67fb677cc60db81d1aa2d2fcc580341f013f1c8a9249bf43281ea0ef32f7d3c5
---

# how-to-set-up-a-fully-functioning-home-network-using-opnsense-2dfb4e5f

The IPCAM network is where you can put IP security cameras that is isolated from your network and the Internet. This can be useful if you have a NVR box where you access all your camera feeds. In this network, you could simply allow access to DNS and everything else will be blocked.
You do not really need to create any rules at all to block all access (access is “deny all” by default when no rules exist), but I found that if you allow DNS, the firewall logs are a lot cleaner because you do not see all of the DNS blocks. IP cameras often try to phone home for cloud services so if DNS is blocked they may try even harder to query DNS, which increases spam to your firewall logs. Since all other access is blocked, even if the IP addresses are resolved for domain names, access to external servers is still blocked.
| Action | TCP/IP Version | Protocol | Source | Dest / Invert | Destination | Dest Port | Description | 
|---|---|---|---|---|---|---|---|
| Pass | IPv4+6 | TCP/UDP | IPCAM net | unchecked | IPCAM address | 53 (DNS) | Allow access to DNS | 
Configure Switch
With OPNsense being configured, you are actually more than halfway done because a bulk of the network infrastructure was completed in OPNsense. You will need a managed or smart network switch that is capable of supporting VLANs in order to make use of VLANs in your network. If you only need basic VLAN functionality, smart switches are often cheaper than managed switches since they offer more basic features than managed switches.
I am going to make the assumption in this guide that you have a new switch that has not been configured or one that you factory reset to the default values. The reason being that an already configured switch will need the current settings changed and it is easier to start at the beginning to learn the steps you need to take.
For a new or reset switch, I recommend that you plug a computer directly into the switch so that you can get it set up before you plug it into OPNsense to avoid any potential issues such as having DHCP running on your switch (it could conflict with DHCP on OPNsense) or a static IP address on the switch which conflicts with the interfaces you have configured in OPNsense.
Connect Directly to Switch
For the configuration of the switch, I recommend plugging into the port that your management LAN/VLAN will be connected to your OPNsense box (in our example, that would be port 6) since you will want to be connected to one of the untagged interfaces so you do not lose connectivity to the switch’s web interface after configuring VLANs (see note below), which then would require you to move over to an untagged port.
As mentioned earlier, you may want a dedicated VLAN for network management, but I am going to demonstrate using the untagged LAN for management (and how you can still make it a safe approach by defaulting unused ports to a VLAN such as the GUEST network – it is a home network so you are free to build it how you feel comfortable securing it).
Note that you can typically configure managed switches to listen on different interfaces so you can still have access to the switch’s web interface from within various VLANs, but since I am configuring a dedicated management network, I do not need to allow web access on different VLANs. That is why I am recommending you to connect to one of the untagged ports for this example.
Once you are plugged into your switch, you will need to determine if your switch has DHCP enabled by default. The easiest way to know is to look at your network status on whatever device you are using to see if you have an IP address assigned. If you do not have an IP address, you will need to manually enter your IP address. However, you will need to consult the manual to see what the default IP address of your switch is set to.
Since I am using a TP-Link managed switch for this example, the default IP address is 192.168.0.1. In this case, manually setting the IP address of the device you are using to configure the switch to 192.168.0.10 with a subnet mask of 255.255.255.0 will be sufficient.
The TP-Link switches have a web interface and newer models can be configured using the Omada Software Controller. The Omada software is not required to configure the switch, so I am going to simply use the web interface. If you can access http://192.168.0.1 or https://192.168.0.1 successfully after configuring the static IP address then you are good to go.
The default username of the TP-Link switch is admin and the password is admin, but you will need to consult the user manual of the switch you are using. Of course, I recommend changing your password after you sign in.
Change the Switch’s Interface IP Address
The first thing you should do is change the IP address of the network switch to a static IP address that resides in your management network. In my example, I will set it to 192.168.1.2 so the switch will be included on the LAN network since the router is using 192.168.1.1/24 for the LAN interface. This step is important because you want to make sure all your network infrastructure is accessible within your management network.
Click on the “L3 Features” tab at the top of the page, and then “Interface” on the left side menu. You should see the default IP address. Click on “Edit IPv4” to edit the IPv4 address.
Enter 192.168.1.2 and click “Apply”. You will likely lose connectivity immediately and will have to start using the http://192.168.1.2 web address. There may be the possibility you will need to change your IP address on your PC/laptop to be in the same subnet if you cannot connect to the new IP address on the switch (I cannot easily test this because I do not have an extra switch laying around that I am able to reset without taking down parts of my network).
Be sure to log back in and click on the “Save” button in the upper right hand corner to make sure the changes are persistent once you know you are able to connect to the new IP address. If you do not click “Save” on TP-Link switches, the changes will be lost when you reboot the switch.
Disable DHCP
You should check to make sure DHCP is enabled or not. Most likely, you already know based on whether or not your PC obtained an IP address automatically or not.
On the “L3 Features” page, click on the “DHCP Service > DHCP Server” side menu. Then make sure the “Enable” checkbox for the “DHCP Server” is unchecked. Click “Apply” and “Save” to persist the changes.
LAG Configuration
On the “L2 Features” page for the “LAG” configuration on the “LAG Table” section, you may select a different hashing algorithm. I like to choose an algorithm that includes both the source and destination as part of the hash so that traffic can be distributed more evenly.
The “LACP Config” section is where the LAG configuration resides. For my example, port 2 and 4 will be a LAG that is connected to OPNsense. Check both ports, choose “Enable” for the “Status”, “Group ID” of 1, leave the port priority at the default value, and choose “Active” for the “Mode”. Click the “Apply” button when you are finished.
Physical Diagram of Connected Devices
Before proceeding with the VLAN configuration, refer to the following physical diagram for devices/clients connected to the network (see also the physical network infrastructure diagram). It will be beneficial to see which Ethernet ports the devices are plugged into when configuring the switch.
VLAN Configuration
Go to the “L2 Features” page and click on the “VLAN > 802.1Q VLAN” left side menu to see the list of VLANs. By default, every port is assigned VLAN1 which is designated for untagged ports. The managed switch will behave like a basic unmanaged switch by default.
