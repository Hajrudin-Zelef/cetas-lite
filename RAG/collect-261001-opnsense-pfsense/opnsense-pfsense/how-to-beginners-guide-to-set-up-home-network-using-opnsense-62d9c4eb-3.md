---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/how-to-beginners-guide-to-set-up-home-network-using-opnsense-62d9c4eb-3
title: "how-to-beginners-guide-to-set-up-home-network-using-opnsense-62d9c4eb"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["cost"]
source: docs/RAG/collect-261001-opnsense-pfsense/how-to-beginners-guide-to-set-up-home-network-using-opnsense-62d9c4eb.md
source_anchor: ""
source_lines: [116, 185]
sha256: 88a6bec14fac3a43291085a94f398aef3a9c480c23db9e782b32922c2d068658
---

# how-to-beginners-guide-to-set-up-home-network-using-opnsense-62d9c4eb

Not only is utilizing VLANs cost effective, but it allows you to make better use of your network resources. Limiting broadcast domains can improve performance on your network if you have a large number of devices especially if they utilize a significant amount of bandwidth. When combining VLANs with firewalls, you can also improve the overall security of your network by limiting access between groups of devices.
In this guide, I am going to create one VLAN for untrusted devices on your network. The original guide provided 5 different VLAN examples but to keep this guide simplified, only a single VLAN will be demonstrated to help you get started with segregating your devices to improve network security.
To create a VLAN in OPNsense, go to the “Interfaces > Other Types > VLAN” page. When creating the VLAN, you will use the LAN interface as the parent interface. As mentioned above, VLANs require a physical interface in which to create logical networks.
| Option | Value | 
|---|---|
| Device | Leave empty to automatically generate a name | 
| Parent | igc1 (use the LAN interface as the parent) | 
| VLAN tag | 10 | 
| VLAN priority | You may use the default “Best Effort” or select priorities (not sure how much it impacts actual performance) | 
| Description | UNTRUSTED | 
Interfaces: Assignments
After the VLAN is created, you will be able to assign it to an interface. You can think of an “interface” as not only the address of the physical port itself but also the gateway to an entire network. That concept may seem confusing to new users, but creating a new interface assignment is how you create separate physical or logical networks in OPNsense (and other router platforms).
When creating an interface you can specify the size of the network, which limits the total number of devices that can be connected to each network. The interface acts as the gateway for each network where traffic may enter or exit.
On the “Interfaces > Assignments” page, you can create a new interface by clicking on the “+” button in the “New interface” section of the page. The dropdown box only shows unassigned physical/logical interfaces. Once you assign the interface, it will no longer be included in the dropdown.
The WAN and LAN interfaces should already be assigned from the OPNsense installation so I will only mention setting up the VLAN interface assignment.
Select the UNTRUSTED VLAN listed in the “Network port” dropdown box (it should be the only value available to select in the dropdown box) and enter the appropriate “Description” of UNTRUSTED. The “Description” is what is displayed on the “Interfaces” section in the left side menu so it is important to use a short name to indicate the purpose of each interface you assign. Otherwise, the interfaces will show up as “OPT1”, “OPT2”, etc., which will be very confusing if you have multiple networks to manage.
Click the “Save” button when you are finished.
Interface Pages
Each interface has its own page under the “Interfaces” menu on the left side of the OPNsense user interface. They will appear as [WAN], [LAN], and [UNTRUSTED]. Please go to the appropriate interface pages to modify the configuration as described below.
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
| IPv6 Configuration Type | None | 
Interfaces > [LAN], [UNTRUSTED]
You will need to edit the settings for the LAN and UNTRUSTED interface. For the LAN interface, all of the default values might be fine but I am including the settings below as a reference.
Use the following common values for the options of both interfaces:
| Option | Value | 
|---|---|
| Enable | “Enable Interface” should be checked by default by the OPNsense installation | 
| Lock | Check “Prevent interface removal” so you cannot easily remove the interface from the “Interfaces > Assignments” page | 
| Block private networks | Unchecked (all internal networks should have this unchecked) | 
| Block bogon networks | Unchecked (all internal networks should have this unchecked) | 
| IPv4 Configuration Type | Static IPv4 | 
| IPv6 Configuration Type | None | 
| IPv4 Upstream Gateway | Auto-detect | 
And use the following IPv4 values in the table below for each corresponding interface:
| Interface | Description | IPv4 address |  | 
|---|---|---|---|
| [LAN] | LAN | 192.168.1.1/24 (This should be the default value) |  | 
| [UNTRUSTED] | UNTRUSTED | 192.168.10.1/24 |  | 
DHCP Configuration
Once the interfaces are assigned and enabled, you will want to enable DHCP on the interfaces so that all of your devices will automatically be assigned IP addresses when they are plugged into your network switch or join your WiFi network.
For the DHCP settings, you may want to enable a wider range of IP addresses if you have more than 100 devices on any of your networks, but for most users the ranges I specify below should be sufficient.
If you plan to have some devices use static IP addresses (which is recommended when hosting various apps/services on your network), I recommend that you do not set the DHCP IP address range to include the full subnet (such as 192.168.1.2 - 192.168.1.254) so that you have some IP addresses available for static IPs. The static IP addresses need to be outside of the DHCP range you specify.
Do not forget to click the “Save” button after configuring each interface.
DHCPv4
To reduce the length of this guide, refer to the table below to enter the IP address ranges for each interface’s DHCPv4 page by going to the “Services > DHCPv4” menu and clicking on each interface’s page such as “Services > DHCPv4 > [LAN]”.
For each interface below, be sure to click the “Enable” checkbox.
| Interface | Range from | Range to | 
|---|---|---|
| [LAN] | 192.168.1.100 | 192.168.1.200 | 
| [UNTRUSTED] | 192.168.10.100 | 192.168.10.200 | 
DNS Configuration
I think configuring the DNS options in OPNsense can be a bit confusing for new users (I struggled at first too) primarily because there are a couple of places where you may specify DNS information. There are various approaches to how you may configure DNS so depending on the approach taken is where you need to enter the DNS configuration.
For simplicity when you start learning how to configure OPNsense, you may simply use the ISP’s DNS servers which happens to be the default DNS configuration in OPNsense. In this guide, I am going to simply leave the Unbound DNS options mostly at the default settings. You may explore the various DNS topics I have written about on this site for further configuration options.
System: Settings: General
Leave the “DNS servers” boxes blank and check the option “Allow DNS server list to be overridden by DHCP/PPP on WAN”. This should be the default configuration, but I wanted to mention it to ensure they are set properly.
Click the “Save” button to apply the changes.
Unbound DNS: General
On the “Services > Unbound DNS > General” page, set the following configuration values:
| Option | Value | 
|---|---|
| Enable | Check “Enable Unbound” (if not enabled already) | 
