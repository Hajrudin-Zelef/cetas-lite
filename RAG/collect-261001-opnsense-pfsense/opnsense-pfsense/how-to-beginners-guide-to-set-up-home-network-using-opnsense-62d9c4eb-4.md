---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/how-to-beginners-guide-to-set-up-home-network-using-opnsense-62d9c4eb-4
title: "how-to-beginners-guide-to-set-up-home-network-using-opnsense-62d9c4eb"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/how-to-beginners-guide-to-set-up-home-network-using-opnsense-62d9c4eb.md
source_anchor: ""
source_lines: [186, 235]
sha256: 1406ece4a51719a71a92c7dda75612a7a2319d5a60f2474b194e705ea413db43
---

# how-to-beginners-guide-to-set-up-home-network-using-opnsense-62d9c4eb

| Listen Port | Leave as default 53 | 
| Network Interfaces | Choose “All (recommended)” (should be default) | 
| DNSSEC | Check “Enable DNSSEC Support” (only if the upstream DNS servers support this option – if unsure, leave unchecked) | 
| DHCP Registration | Check “Register DHCP leases” (to use hostnames of DHCP clients) | 
| DHCP Static Mappings | Check “Register DHCP static mappings” (to use hostnames of static DHCP clients) | 
| DNS Cache | Check “Flush DNS cache during reload” (to clear the cache after making changes to Unbound) | 
| Local Zone Type | transparent (the default value) | 
Click the “Save” button at the bottom of the page and then click the “Apply changes” button at the top of the page to reload the Unbound service to apply configuration changes.
Firewall Configuration
Firewall rules are critical for providing increased security among the devices in your network. Having a solid understanding in this area will be crucial in helping you lock down your network tighter.
As you likely know, no software or hardware is fully impenetrable, which is why it is important to have several layers of defense when protecting your network.
Firewall rules work in conjunction with VLANs to isolate and limit access to various devices on your network.
Firewall: Aliases
Firewall aliases are useful when you want to use more than one IP/network address, port numbers, etc. in a firewall rule, you want to reuse values across multiple rules, or you simply want your rules to be easier to read and maintain. Instead of seeing 192.168.10.10 as the source for the firewall rule, you could create an alias called MyPC, which is much easier to understand what is being allowed or blocked.
The first alias you should create is one which contains all of the RFC 1918 private IPv4 address ranges so that any future networks you create will also be isolated and protected from each other. If you create an alias which only has the network addresses of the LAN, USER, IOT, and GUEST networks, you may forget to add new network addresses if you are adding a new VLAN, which means you may accidentally leave access open to your new network since it is not in the alias used to block access.
Note
I will be making use of firewall aliases in the firewall rule examples, so if you see a name instead of an IP address for the “Source” or “Destination” it means I am using either a built-in firewall alias or the custom firewall aliases I describe in the this section.
The names you see in the firewall rules are not hostnames of devices on the network because you can only use hostnames in firewall aliases. Firewall rules only allow you to enter a single IP/network address or a single firewall alias (aliases may contain more than one value). If you wish to use multiple IP addresses or network addresses in a single firewall rule, you have to create an alias containing those addresses and use that alias in the firewall rule.
Visit the “Firewall > Aliases” page and click the “+” button at the bottom of the page to create the following PrivateNetworks alias:
| Option | Value | 
|---|---|
| Enabled | Checked | 
| Name | PrivateNetworks | 
| Type | Network(s) | 
| Content | 10.0.0.0/8,172.16.0.0/12,192.168.0.0/16 | 
| Description | All local IPv4 networks | 
When you start writing firewall rules, you will notice that aliases such as “LAN net” and “UNTRUSTED net” are created automatically based on the interface names so you do not need to create aliases for each network or interface IP address, which is convenient.
Click the “Apply” button to ensure the alias changes are applied. If you do not click “Apply”, the new aliases will not be available to select when creating firewall rules.
Firewall: Rules: [LAN]
The LAN network will already have the “allow all IPv4” and “allow all IPv6” rules created by default from the OPNsense installation. In order to isolate the two networks in the example used in this guide, those rules will no longer be used.
To avoid confusion with updating the existing rules in the LAN interface, you may remove the two allow all IPv4/IPv6 rules on the “Firewall > Rules > LAN” page. Do not click “Apply” until you have added the rules below!. Enter the following rules in the same order shown in the following table. Make sure you have the destination invert option checked on the 3rd rule!
| Action | TCP/IP Version | Protocol | Source | Dest / Invert | Destination | Dest Port | Description | 
|---|---|---|---|---|---|---|---|
| Pass | IPv4 | TCP/UDP | LAN net | unchecked | LAN address | 53 (DNS) | Allow access to DNS on the LAN interface | 
| Pass | IPv4 | ICMP | LAN net | unchecked | any | any | Allow ICMPv4 from LAN to all networks | 
| Pass | IPv4 | any | LAN net | checked | PrivateNetworks | any | Block access to other internal networks but allow access to the Internet | 
These rules will isolate the LAN from any other local network (including the UNTRUSTED network) and allow access to the Internet. If you want to allow the LAN to reach anything specific in your UNTRUSTED network, you simply just need to add a firewall rule above the bottom rule. Notice that I am using “LAN net” as the source instead of “any” to help ensure we do not allow any potential security holes since there are also a VLAN residing on the LAN interface.
Since the LAN is the trusted network, I included the 2nd rule to allow all ICMPv4 from the LAN to all other networks so that it is possible to use ping and other network utilities to help make it easier to troubleshoot network issues.
Many users prefer to block ICMPv4 on their local networks either entirely or only on some of their networks in an attempt to make it more difficult to discover devices on the network. However, in reality it likely does not provide a significant amount of security. For the reasons described in the article linked in this paragraph, you may want to consider enabling ICMPv4 on all networks since it can help improve the responsiveness of your network, waste less network resources, and provide easier troubleshooting on your internal networks.
Firewall: Rules: [UNTRUSTED]
For the UNTRUSTED network, only access to the Internet will be allowed. This will fully separate your untrusted devices from your trusted devices. If you wish to follow this strict security model, you should never create any rules that allow access to devices on your LAN network if you want the maximum protection.
However, if you want your untrusted devices to access your NAS on the LAN network, you could create a rule allowing very specific access, which is still better than allowing full access but is still less than ideal if you want to keep untrusted devices from communicating with your trusted devices.
| Action | TCP/IP Version | Protocol | Source | Dest / Invert | Destination | Dest Port | Description | 
|---|---|---|---|---|---|---|---|
| Pass | IPv4 | TCP/UDP | UNTRUSTED net | unchecked | UNTRUSTED address | 53 (DNS) | Allow access to DNS on the UNTRUSTED interface | 
| Pass | IPv4 | any | UNTRUSTED net | checked | PrivateNetworks | any | Block access to other internal networks but allow access to the Internet | 
Note
If you need access to your NAS from two different networks, for instance, you may make use of “multi-homing” if your NAS has more than one network interface. Essentially you can connect the NAS to both networks so the traffic to/from the NAS does not have to pass through the firewall. Multi-homing can minimize wasteful network usage since less bandwidth intensive traffic needs to route through your firewall.
Configure Switch
With OPNsense being configured, you are actually more than halfway done because most of the network infrastructure configuration was completed in OPNsense. As mentioned at the beginning of this guide, you will need a network switch that is capable of supporting VLANs in order to follow along with the rest of this guide.
