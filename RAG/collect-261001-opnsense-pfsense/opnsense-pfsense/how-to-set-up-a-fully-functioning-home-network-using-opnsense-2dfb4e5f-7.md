---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/how-to-set-up-a-fully-functioning-home-network-using-opnsense-2dfb4e5f-7
title: "how-to-set-up-a-fully-functioning-home-network-using-opnsense-2dfb4e5f"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-opnsense-pfsense/how-to-set-up-a-fully-functioning-home-network-using-opnsense-2dfb4e5f.md
source_anchor: ""
source_lines: [387, 455]
sha256: 1ce12b08ffd919c159c613e0561eaec8d5ff475c9332547715b12c8d75179512
---

# how-to-set-up-a-fully-functioning-home-network-using-opnsense-2dfb4e5f

The first alias you should create is one which contains all of your networks so you can use the alias to isolate your networks. If you are only using private IPv4 addresses and/or you have static IPv6 addresses, you could simply use all of the RFC 1918 private IPv4 address ranges and your static IPv6 address range so that any future networks you create will also be isolated and protected. When adding new networks it can be easy to forget to include the new network in the alias unless perhaps you create a checklist as a reminder of what you need to do each time you add a new network.
Create a new firewall alias by visiting the “Firewall > Aliases” page and clicking the “+” button at the bottom of the page. For IPv4 addresses and static IPv6 addresses, you could create the following PrivateNetworks alias:
| Option | Value | 
|---|---|
| Enabled | Checked | 
| Name | PrivateNetworks | 
| Type | Network(s) | 
| Content | 10.0.0.0/8,172.16.0.0/12,192.168.0.0/16,(include your static IPv6 range if you have one) | 
| Description | All local networks | 
However, if you are using IPv6 networks which are (unfortunately) dynamic as is the case with my ISP, you will need to add all of your networks using the automatically generated network aliases to the PrivateNetworks alias so that the dynamic IPv6 network addresses will be kept up to date without requiring you to continually update this alias.
When you start writing firewall rules, you will quickly notice that aliases such as “LAN net” and “DMZ net” are created automatically based on the interface names. If you try to create your alias and include “LAN net”, you will realize that you cannot select them from the dropdown menu.
In my previous version of this guide, I recommended creating an firewall group so that you can include the group’s network alias in your PrivateNetworks alias as a workaround for this issue (there is nothing wrong with what I did, but it added extra unnecessary steps). If you plan to have common firewall rules that apply to all your networks, you may still want to create the firewall group so you do not have to repeat as many rules across multiple interfaces.
Someone mentioned in the comments below that you can add the automatically generated network aliases if you type __ in the “Contents” box. I thought I tried that before, but I may have overlooked it when preparing all of the details for this large guide. When you start typing the underscores, you will not see “LAN net” but rather __lan_network for the LAN and __opt1_network for your VLAN network(s).
By using these built-in network aliases, the dynamic IPv6 values will be refreshed automatically when your IPv6 addresses change. These aliases also contain the IPv4 private networks so you do not need to include the RFC 1918 private addresses unless you are worried you will forget to update this alias when you add a new network. Most likely, you may end up noticing it as you are configuring your networks or you discover access is allowed or denied when it should not be.
You should use either the PrivateNetworks alias above or the one below depending if your ISP assigns dynamic IPv6 addresses or not. Remember, if you choose the second option below, you will need to keep it updated when you add new networks unlike the option above (because the option above, all possible IPv4 and IPv6 addresses are known and therefore do not need to change).
| Option | Value | 
|---|---|
| Enabled | Checked | 
| Name | PrivateNetworks | 
| Type | Network(s) | 
| Content | __lan_network, __opt1_network, __opt2_network, __opt3_network, __opt4_network, __opt5_network | 
| Description | All local networks | 
Now for the sake of illustration purposes, I will refer to the following aliases in subsequent firewall rules to demonstrate the types of rules you may want to create to allow access to certain devices/apps across your networks.
A webserver that will be placed in the DMZ network:
| Option | Value | 
|---|---|
| Enabled | Checked | 
| Name | WebServer | 
| Type | Host(s) | 
| Content | 192.168.10.10 | 
| Description | A local web server | 
A PC that will be placed in the USER network:
| Option | Value | 
|---|---|
| Enabled | Checked | 
| Name | PC | 
| Type | Host(s) | 
| Content | 192.168.20.10 | 
| Description | A PC | 
A printer which is connected to the IOT network via Ethernet or WiFi (I recommend enabling the MDNS plugin so the printer can be automatically discoverable across your networks):
| Option | Value | 
|---|---|
| Enabled | Checked | 
| Name | Printer | 
| Type | Host(s) | 
| Content | 192.168.30.10 | 
| Description | A printer | 
Several IP cameras on the IPCAM network (notice you can use IP address ranges if you have your devices sequentially addressed):
| Option | Value | 
|---|---|
| Enabled | Checked | 
| Name | IPCameras | 
| Type | Host(s) | 
| Content | 192.168.50.10-192.168.50.12 | 
| Description | IP security cameras | 
Click the “Apply” button to ensure the alias changes are applied. If you do not click “Apply”, the new aliases will not be available to select when creating firewall rules.
Firewall: Rules: Floating
Floating rules are very helpful when you want a rule to apply to multiple interfaces. They are similar to firewall groups in the sense that the rules can apply to multiple interfaces. I prefer to use floating rules for situations where I have multiple clients on different networks that need access to the same service on one or more servers on my network. You can minimize the need to create the same basic rule on multiple interfaces. Floating rules are also helpful for IP blocklists such as Spamhaus.
One example I personally use is multiple SSH clients needing to access multiple SSH servers. Both the clients and servers are on several different networks. I have firewall aliases for both the SSH clients and the SSH servers so that I only need one floating rule to allow that access. Without this floating rule, the alternative is to create one rule on each interface to allow the clients to access the servers, which would require 2-3 more rules in my case.
For the purposes of this guide, I am only going to create one floating rule for allowing ICMPv6 on all interfaces including the WAN since IPv6 relies more heavily on ICMP than IPv4. The general recommendation is to leave it enabled. You will also receive better scores when testing your IPv6 readiness.
On the “Firewall > Rules > Floating” page, click the “+” button to create a new floating rule. Enter the following information:
| Action | Interface | TCP/IP Version | Protocol | Source | Dest / Invert | Destination | Description | 
|---|---|---|---|---|---|---|---|
| Pass | Leave blank for all interfaces | IPv6 | ICMP | any | unchecked | any | Allow ICMPv6 on all networks | 
Firewall: Rules: [LAN]
The LAN network will already have the “allow all IPv4” and “allow all IPv6” rules created by default from the OPNsense installation. However, I will tweak them so that access to your other networks is limited. If you recall from earlier, in this example I am using the untagged LAN as the management network where all of the critical network infrastructure will be managed.
Some users like to allow their LANs have full access to every other network and then limit the other networks from accessing the LAN since it can make network administration or access to other systems easier. However, I think it is good practice to also limit the reach of the LAN network even if you are only using as a management network for all of your network infrastructure and services.
Just like any other network, if a compromise happens there, isolating it could still be beneficial to protect other parts of your network (even though it would be very bad if your most critical network is compromised).
Enter the following rules on the “Firewall > Rules > LAN” page in the same order as shown below:
