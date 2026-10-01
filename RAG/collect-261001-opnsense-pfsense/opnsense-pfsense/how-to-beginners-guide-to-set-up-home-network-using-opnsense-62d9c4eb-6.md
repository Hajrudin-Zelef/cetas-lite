---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/how-to-beginners-guide-to-set-up-home-network-using-opnsense-62d9c4eb-6
title: "how-to-beginners-guide-to-set-up-home-network-using-opnsense-62d9c4eb"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/how-to-beginners-guide-to-set-up-home-network-using-opnsense-62d9c4eb.md
source_anchor: ""
source_lines: [271, 306]
sha256: 18df16c51f7ad9bfe6ad92f68cb59ce668d30ad251c974b4816834d1c6266cf3
---

# how-to-beginners-guide-to-set-up-home-network-using-opnsense-62d9c4eb

You can connect to Grandstream wireless APs via a built in web interface if you are not using their locally installed or cloud managed controller, which is very convenient when you only need to configure one device.
If you go to the “Services > DHCPv4 > Leases” page in OPNsense, you will be able to find the IP address of the Grandstream AP. The IP address should reside in your trusted LAN such as 192.168.1.104. Simply enter https://192.168.1.104 depending on the IP address that is assigned to the AP.
For Grandstream APs, the default username is admin and the default password is located on the bottom of the AP so you will need to use that for the password.
Create the Trusted Network (LAN) SSID
The first time you log into the Grandstream AP, you will have a basic configuration wizard to set up the device with the first SSID. For the first screen since there is only one AP being configured, you can essentially click “Next”. If you have more than one AP, the nice thing is that you can manage your other APs from the first AP so there is no need to install a local controller.
On the second screen, you can configure your first SSID for your wireless network. In the scenario for this guide, I am going to create 2 WiFi networks – one for your trusted network and one for the untrusted network. When you complete the configuration of the first SSID, it will by default be on your trusted network since no VLAN ID is set for the SSID.
Enter the desired “SSID” and “WPA Pre-Shared Key”, which is the password for your WiFi connection. You will want to name your trusted and untrusted WiFI network as something meaningful to you.
For the “Security Mode”, you may use WPA2 for the greatest compatibility among your clients since not everything may support WPA3. If you have all newer devices, you can likely use WPA3. Some APs offer the option to support both WPA2 and WPA3 simultaneously including the Grandstream AP. That option might be the best of both worlds for compatibility and for improved wireless security since newer devices should hopefully default to WPA3.
Make sure your current device on the “Member Devices” to include your access point. Otherwise the SSID settings will not be applied to your access point (you can fix that later if you forget to check that option).
Click “Complete” to set up the first SSID. You should now have your trusted network SSID set up! If you like, you can try connecting a mobile device to that network to see if you get an IP address in the 192.168.1.x network.
Create the Untrusted Network SSID
Go to the “SSIDs” page by clicking on the “SSIDs” menu on the left side of the page. You will see the trusted network SSID that was just created. Click on the “Add” button to create a new SSID.
Enter the “SSID” for the untrusted network. You may not want to have “UNTRUSTED” in the name, but I am using that as an illustration to make it clear which network the SSID is on. Check the “Enable SSID” box.
You will need to check the “VLAN” option and enter 10 as the “VLAN ID”. You can select your desired “Security Mode” as described earlier as well as the “WPA Pre-Shared Key” for your password.
Since the network is untrusted, you may want to consider enabling “Client Isolation” to prevent wireless clients from communicating directly with each other.
I have seen it recommended to set the “DTIM Period” to 3 to help conserve battery usage of mobile devices that are using WiFi.
Before clicking “Save”, go to the “Device Membership” tab at the top of the dialog box. Select the AP in the “Available Devices” and click the right arrow button to move it to the “Member Devices” so that the SSID gets applied to the access point.
Make sure the device is moved to the “Member Devices” box as shown below before clicking “Save”.
Click “Save” to the configuration of the second SSID. You will also need to click the “Apply” button for changes to take effect on the AP.
You may try connecting to that SSID to see if you get an IP address in the 192.168.10.x network. If you do, the VLAN configuration is working properly!
Next Steps
If all goes well, you should have a fully functioning home network with a trusted LAN and an untrusted VLAN to separate devices that may be more likely to be compromised, which helps to improve the security of your most trusted devices! Congratulations!
My hope is that you found this simplified version of the full network build to be beneficial if you are a novice user. The great thing about your home network is that you are free to build it to meet your wants and needs!
Below are a few ideas of areas to explore next should you find yourself wanting to go further on this journey.
Implementing Additional Security Features
Since you have the basics configured if you followed this guide, you may wish to go deeper into other areas to add more features or to implement more layers of security. As a reference, I have written a guide to discuss several security related features that you may wish to implement on your OPNsense system or your home network.
You may also wish to check out the original full network build guide which covers more advanced networking topics using a more complex network architecture.
Multi-Homing Device(s)
Multi-homing is the concept of placing a single device into two or more separate networks. In order to accomplish this, you need a system with two or more network interfaces. Some mini-PCs and NAS devices include more than one network interface, which is very useful if you wish to multi-home the device.
A NAS is one good example where you may want to multi-home. Routing lots of network traffic through the firewall can slow down performance significantly especially if you are running any intrusion detection/prevention services on the firewall. IDS/IPS requires a great deal of computing power in order to process all of the data packets on the network in a timely fashion.
By putting your NAS on multiple networks where access is needed, you can prevent high bandwidth traffic from traversing across networks and through the firewall. You should consider multi-homing your NAS if firewall performance is suffering. This is a topic I may explore in greater detail in future guides.
In the diagram below, you would connect both interfaces to your switch but configure each port to be on different networks. If following the example in the guide, you could include the NAS on both the LAN and the UNTRUSTED VLAN. Each network interface would have an IP address in the respective network. Then you would configure client devices on each network to access network shares, for instance, using the corresponding IP address for the NAS for each network. The clients will be able to communicate freely with the NAS without traversing the firewall.
Secure, Remote Access to Your Network
If you are interested in remotely accessing your network for various reasons, you may run IPSec, OpenVPN, WireGuard, or Zerotier using built-in functionality or plugins.
Once you have your VPN set up in OPNsense, you can create firewall rules to allow the desired access to your internal networks. For instance, you may want to create a rule to access all of the devices on your trusted network (the LAN in the example provided in this guide).
I have written guides on OpenVPN and WireGuard if you are interested in setting up those VPN services in OPNsense.
