---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/how-to-set-up-a-fully-functioning-home-network-using-opnsense-2dfb4e5f-5
title: "how-to-set-up-a-fully-functioning-home-network-using-opnsense-2dfb4e5f"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/how-to-set-up-a-fully-functioning-home-network-using-opnsense-2dfb4e5f.md
source_anchor: ""
source_lines: [242, 329]
sha256: f4a9ddd26a4b2c405b40b90e25895fc7642cfaf600cdfa5a773c7258fdd95768
---

# how-to-set-up-a-fully-functioning-home-network-using-opnsense-2dfb4e5f

I have seen mention of resolving performance issues by assigning the physical parent interface when the parent interface is not being utilized. If you are experiencing poor performance across your VLANs, you may need to assign the parent interface without fully configuring the interface.
Personally, I have not noticed this issue with the Protectli VP2410 and the network configuration described in this guide. I tried assigning/unassigning the parent interface with no noticeable impact to performance.
Interface Pages
Each interface has its own page under the “Interfaces” menu on the left side of the OPNsense user interface. They will appear as [WAN], [LAN], etc. Go to the appropriate interface pages below to modify the configuration described below.
Interfaces > [WAN]
For the WAN interface, you may not have to change much of the configuration especially if your ISP uses DHCP, but for the sake of completeness I will list out the configuration settings with brief explanations for your reference.
| Option | Value | 
|---|---|
| Enable | “Enable Interface” should be checked by default by the OPNsense installation | 
| Lock | Check “Prevent interface removal” so you cannot easily remove the interface from the “Interfaces > Assignments” page | 
| Description | WAN (the default value from the OPNsense installation) | 
| Block private networks | Checked (should be checked if connected directly to the Internet, otherwise you should uncheck it) | 
| Block bogon networks | Checked (should be checked if connected directly to the Internet, otherwise you should uncheck it) | 
| IPv4 Configuration Type | DHCP (if your ISP uses DHCP) | 
| IPv6 Configuration Type | DHCPv6 (if your ISP uses DHCP) | 
If you want to use IPv6 with your VLANs and your ISP supports prefixes greater than /64 (such as Comcast Xfinity), you may enter configuration similar to the following under the “DHCPv6 client configuration”:
| Option | Value | 
|---|---|
| Request only an IPv6 prefix | Unchecked | 
| Prefix delegation size | 60 | 
| Send IPv6 prefix hint | Checked | 
| Use IPv4 connectivity | Unchecked | 
| Use VLAN priority | Disabled | 
Interfaces > [LAN], [DMZ], etc.
For the LAN/VLAN interfaces, use the following common values for the options of every interface:
| Option | Value | 
|---|---|
| Enable | “Enable Interface” should be checked by default by the OPNsense installation | 
| Lock | Check “Prevent interface removal” so you cannot easily remove the interface from the “Interfaces > Assignments” page | 
| Block private networks | Unchecked (all internal networks should have this unchecked) | 
| Block bogon networks | Unchecked (all internal networks should have this unchecked) | 
| IPv4 Configuration Type | Static IPv4 | 
| IPv6 Configuration Type | Track Interface (if your ISP allows prefix delegations) | 
| IPv4 Upstream Gateway | Auto-detect | 
| IPv6 Interface | WAN | 
| Manual configuration | Check “Allow manual adjustment of DHCPv6 and Router Advertisements” | 
And use the following IPv4/IPv6 values in the table below for each corresponding interface:
| Interface | Description | IPv4 address | IPv6 Prefix ID | 
|---|---|---|---|
| [LAN] | LAN | 192.168.1.1/24 | 0 | 
| [DMZ] | DMZ | 192.168.10.1/24 | 1 | 
| [USER] | USER | 192.168.20.1/24 | 2 | 
| [IOT] | IOT | 192.168.30.1/24 | 3 | 
| [GUEST] | GUEST | 192.168.40.1/24 | 4 | 
| [IPCAM] | IPCAM | 192.168.50.1/24 | 5 | 
DHCP Configuration
Once the interfaces are enabled, you will most likely want to enable DHCP on the interfaces so that all of your devices will automatically be assigned IP addresses when they are plugged into your network switch or join your WiFi network.
For the DHCP settings, you may want to enable a wider range of IP addresses if you have more than 100 devices on any of your networks, but for most users the ranges I specify below should be sufficient.
If you plan to have some devices use static IP addresses (which is recommended when hosting various apps/services on your network), I recommend that you do not set the DHCP IP address range to include the full subnet (192.168.1.2 - 192.168.1.254, for example) so that you have some available addresses for static IPs.
Do not forget to click the “Save” button after configuring each interface.
DHCPv4
To reduce the length of this guide, refer to the table below to enter the IP address ranges for each interface’s DHCPv4 page by going to the “Services > DHCPv4” menu and clicking on each interface’s page.
For every interface below, be sure to click the “Enable” checkbox.
| Interface | Range from | Range to | 
|---|---|---|
| [LAN] | 192.168.1.100 | 192.168.1.200 | 
| [DMZ] | 192.168.10.100 | 192.168.10.200 | 
| [USER] | 192.168.20.100 | 192.168.20.200 | 
| [IOT] | 192.168.30.100 | 192.168.30.200 | 
| [GUEST] | 192.168.40.100 | 192.168.40.200 | 
| [IPCAM] | 192.168.50.100 | 192.168.50.200 | 
Add at the bottom of the each interface’s page, click the “+” button to add static DHCP reservations for devices that are hosting apps/services or devices where you wish to apply specific firewall rules (although you can use hostnames in firewall aliases for clients with dynamic IP addresses, I have found it to be problematic at times – a static IP is the most reliable approach).
You can also add static DHCP reservations directly from the “Services > DHCPv4 > Leases” page. It has the added benefit of prefilling the MAC address. Either way, you will need to enter the same information.
For demonstration purposes, I will use several static DHCPv4 IP reservations that will be referenced by firewall rule aliases and included in several firewall rules. I am using randomly generated MAC addresses in the table as examples, but you will need to use your actual MAC addresses.
| Interface | MAC address | IP address | Hostname | 
|---|---|---|---|
| [DMZ] | 0e:66:c8:ee:a8:ee | 192.168.10.10 | webserver | 
| [USER] | 00:49:c0:9b:a8:cc | 192.168.20.10 | pc | 
| [IOT] | 3e:03:06:01:45:d3 | 192.168.30.10 | printer | 
| [IPCAM] | ee:f2:4d:ca:92:44 | 192.168.50.10 | ipcam1 | 
| [IPCAM] | c3:c7:cb:db:66:d6 | 192.168.50.11 | ipcam2 | 
| [IPCAM] | 23:d1:6f:17:31:8e | 192.168.50.12 | ipcam3 | 
DHCPv6
As with DHCPv6 to reduce the length of this guide, refer to the table below to enter the IP address ranges for each interface’s DHCPv6 page by going to the “Services > DHCPv6” section and clicking on each interface’s page.
For IPv6 addresses, you may specify only the second half of the address such as ::1000. The first half of the address is used by default and if you only specify 1 segment, the rest are filled with zeroes (it is how IPv6 notation works to allow shorter addresses to be expressed).
For every interface below, be sure to click the “Enable” checkbox.
| Interface | Range from | Range to | 
|---|---|---|
| [LAN] | ::1000 | ::2000 | 
| [DMZ] | ::1000 | ::2000 | 
| [USER] | ::1000 | ::2000 | 
| [IOT] | ::1000 | ::2000 | 
| [GUEST] | ::1000 | ::2000 | 
| [IPCAM] | ::1000 | ::2000 | 
Warning
You will not be able to assign the above ranges for DHCPv6 when you are using DHCPv6 on the WAN interface and “track interface” for each local network interface (LAN, DMZ, etc.) as shown in this guide unless the OPNsense box is connected to your modem. Otherwise if you are not connected, there will be no available IPv6 address ranges to assign since they are assigned by the ISP. You will see an error message on the page that says there are no addresses available to assign. Once you are connected to your modem and have DHCPv6 addresses assigned by the ISP, you should be able to complete this step.
Router Advertisements
For IPv6, there is an additional step which I think is good to do in order to help support all possible IPv6 clients on your network. I like to set the “Router Advertisements” to “Assisted” so that both DHCPv6 and SLAAC is used because not all clients may support DHCPv6.
