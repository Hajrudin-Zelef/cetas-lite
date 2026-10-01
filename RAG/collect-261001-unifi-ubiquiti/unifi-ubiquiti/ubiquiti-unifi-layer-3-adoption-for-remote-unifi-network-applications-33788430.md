---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/ubiquiti-unifi-layer-3-adoption-for-remote-unifi-network-applications-33788430
title: "..."
domain: unifi-ubiquiti
role: reference
task: reference
actors: ["AWS"]
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/ubiquiti-unifi-layer-3-adoption-for-remote-unifi-network-applications-33788430.md
source_anchor: ""
source_lines: [1, 82]
sha256: 3e51a78b61454cd9085024ea61d0e1a35628576e11bf5ec03439eb421cf1eae6
---

# ...

Layer 3 adoption is the process of adopting a UniFi device to a remote UniFi Network application.
You might use Layer 3 adoption for applications located in the cloud (e.g. on Amazon EC2) or NOC.
For regular device adoption, see UniFi – Device adoption.
Overview
In many deployments where it’s not possible to have the UniFi Network host running on-premise, you can run the UniFi Network application in the Cloud or your NOC. For example, for a large-scale project with many devices there are a few possible methods for the adoption of devices:
- Take a laptop to the device’s site to perform adoption via Chrome browser (easiest method).
- When you’re at the site, open a browser and navigate to Cloud: either the UniFi Remote Access Portal or the UniFi Network application (when launched using Cloud).
- Create a virtual application instance on Amazon EC2.
- Either configure the DHCP server or DNS server.
Initial setup
Please make sure you’re familiar with how a regular L2 adoption on UniFi works (where the devices and UniFi Network application are on the same network) before attempting L3 (remote) adoption. Also, remember that in order to adopt, the following conditions must be true in order to have internet access and also have access to the router from within the network (locally):
1. WAN port connected to the Internet.
2. LAN port connected locally to access management features on the router (USG or third party).
UniFi APs have a default inform URL http://unifi:8080/inform. Thus, the purpose of using DHCP option 43 or DNS is to allow the AP to know the IP of the UniFi Network application host.
If you encounter discovery issues please use the UniFi – Troubleshooting Device Adoption article to help you troubleshoot the issue.
After installing the Discovery tool plugin (freely available in Chrome Web Store) on a computer running Chrome browser, any locally-available, unmanaged UniFi Devices (i.e., same L2 network as your computer) will appear as “Pending Adoption” in the UniFi Cloud Access Portal as well as your UniFi Network application itself (in the Devices section in both cases). To access the application remotely Remote Access will have to be enabled.
Via UniFi OS
1. Go to https://unifi.ui.com and login with your Ubiquiti SSO credentials.
2. Navigate to the Devices section.
3. The device to be adopted will appear as ready to be adopted. Click Adopt.
Via the UniFi Remote Access Portal
1. Go to https://network.unifi.ui.com/ and log in with your Ubiquiti SSO credentials.
2. Go to the Devices section and locate the model with the Pending Adoption status. Click ADOPT.
3. In the Adopt window that will appear, select the UniFi Network host and the site that will be adopting the device (for multi-site hosts) and click Adopt.
Via the UniFi Network application
1. Launch UniFi Network, go to the Devices section, find the device that is to be adopted with the status “Pending Adoption” and click Adopt under Actions.
DNS
You’ll need to configure your DNS server to resolve ‘unifi’ to your UniFi Network host’s IP address. Make sure that the device can resolve the UniFi Network domain name. For example, if you are setting http://XYZ:8080/inform, then ping from the device to determine if XYZ is resolvable/reachable. Or you may also use FQDN for the application inform URL: http://FQDN:8080/inform
Troubleshooting: Device (with static IP) fails to connect to the L3 UniFi Network application
- When configuring a device from DHCP to static in the UniFi Network application, make sure you have put the IP of DNS. If not, then the device cannot contact DNS to resolve UniFi Network’s domain name.
- If the device has been reset, make sure that you have “informed” the device twice (using the Discovery Utility) about the UniFi Network application’s location. See steps in the section above.
DHCP Option 43
If using Ubiquiti’s EdgeMAX routers, then DHCP option 43 can be done by just entering the IP address of the UniFi Network host in the “unifi” field on the DHCP-server.NOTE: The UniFi Security Gateway (USG) will not use DHCP option 43 to add the UniFi Network application location when obtaining a DHCP lease on the WAN interface.
To use DHCP option 43 you’ll need to configure your DHCP Server. We provide some third party examples below, but please refer to the manufacturer’s support documentation for up to date instructions. For example:
Linux’s ISC DHCP server: dhcpd.conf
# ...
option space ubnt;
option ubnt.unifi-address code 1 = ip-address;
class "ubnt" {
        match if substring (option vendor-class-identifier, 0, 4) = "ubnt";
        option vendor-class-identifier "ubnt";
        vendor-option-space ubnt;
}
subnet 10.10.10.0 netmask 255.255.255.0 {
        range 10.10.10.100 10.10.10.160;
        option ubnt.unifi-address 201.10.7.31;  ### UniFi Network host IP ###
        option routers 10.10.10.2;
        option broadcast-address 10.10.10.255;
        option domain-name-servers 168.95.1.1, 8.8.8.8;
        # ...
}
Cisco CLI
# assuming your UniFi is at 192.168.3.10
ip dhcp pool <pool name>
network <ip network> <netmask>
default-router <default-router IP address>
dns-server <dns server IP address>
option 43 hex 0104C0A8030A # 192.168.3.10 -> CO A8 03 0A
# Why 0104C0A8030A ?
#
# 01: suboption
# 04: length of the payload (must be 4)
# C0A8030A: 192.168.3.10
Mikrotik CLI
/ip dhcp-server option add code=43 name=unifi value=0x0104C0A8030A
/ip dhcp-server network set 0 dhcp-option=unifi
# Why 0104C0A8030A ?
#
# 01: suboption
# 04: length of the payload (must be 4)
# C0A8030A: 192.168.3.10
User Tip: Find more DHCP Option 43 instructions in the User Notes & Tips section.
SSH
If you can SSH into the device, it’s possible to do L3 adoption via CLI command:
1. Make sure the device is running updated firmware. See this guide: UniFi – Changing the Firmware of a UniFi Device.
2. Make sure the device is in the factory default state. If it’s not, run the following command:
sudo syswrapper.sh restore-default
3. SSH into the device and type the following and hit enter, substituting “ip-of-host” with the IP address of the host of the UniFi Network application:
set-inform http://ip-of-host:8080/inform
4. After issuing the set-inform, the UniFi device will show up for adoption in the Devices section of UniFi Network. Once you click Adopt, the device will appear to go offline or have the status of “Adopting” then proceed to “Provision” and “Connected”.
Source :
https://help.ui.com/hc/en-us/articles/204909754-UniFi-Layer-3-Adoption-for-Remote-UniFi-Network-Applications
