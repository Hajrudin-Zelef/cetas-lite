---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/how-to-set-up-a-fully-functioning-home-network-using-opnsense-2dfb4e5f-8
title: "how-to-set-up-a-fully-functioning-home-network-using-opnsense-2dfb4e5f"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/how-to-set-up-a-fully-functioning-home-network-using-opnsense-2dfb4e5f.md
source_anchor: ""
source_lines: [456, 498]
sha256: d5ec023dc2771bdfc482cf3c8f6d33435bfdf7236a705459eb85e2c7c0522546
---

# how-to-set-up-a-fully-functioning-home-network-using-opnsense-2dfb4e5f

| Action | TCP/IP Version | Protocol | Source | Dest / Invert | Destination | Dest Port | Description | 
|---|---|---|---|---|---|---|---|
| Pass | IPv4+6 | TCP/UDP | LAN net | unchecked | LAN address | 53 (DNS) | Allow access to DNS | 
| Pass | IPv4 | ICMP | LAN net | unchecked | any | any | Allow ICMPv4 from LAN to all networks | 
| Pass | IPv4+6 | any | LAN net | checked | PrivateNetworks | any | Allow access only to Internet | 
This will isolate the LAN from the other VLANs in our network and allow access to the Internet. If you want to allow the LAN to reach anything specific in your network, you simply just need to add a rule above the bottom rule.
Since the LAN will be used for network management, I included a rule to allow all ICMPv4 from the LAN to all other networks so that it is possible to use ping and other network utilities for troubleshooting purposes. Note that I am not including ICMPv6 for IPv6 because I already included a floating rule to all ICMPv6 on all networks. IPv6 relies more heavily on ICMP so it is generally not recommended to block ICMPv6.
Many users prefer to block ICMPv4 on their local networks either entirely or only on some of their networks in an attempt to make it more difficult to discover devices on the network. However, in reality it likely does not provide a significant amount of security. Rather than disabling ICMP entirely, you may consider blocking only a subset of ICMP types that may be abused especially in more sensitive networks. For the reasons described on the article linked in this paragraph, you may want to consider enabling ICMPv4 on all networks since it can help improve the responsiveness of your network, waste less network resources, and provide easier troubleshooting on your internal networks.
Firewall: Rules: [DMZ]
The DMZ network will follow a similar patter to the LAN network above. On the “Firewall > Rules > DMZ” page, enter the following values in the same order shown below:
| Action | TCP/IP Version | Protocol | Source | Dest / Invert | Destination | Dest Port | Description | 
|---|---|---|---|---|---|---|---|
| Pass | IPv4+6 | TCP/UDP | DMZ net | unchecked | DMZ address | 53 (DNS) | Allow access to DNS | 
| Pass | IPv4+6 | any | DMZ net | checked | PrivateNetworks | any | Allow access only to Internet | 
Similar to the LAN, this will isolate the DMZ from the other VLANs in our network and allow access to the Internet. In general, you probably should not allow the DMZ to access anything else in your internal network unless perhaps you have a dedicated network just for hosted apps/services and all you have in the DMZ is simply a reverse proxy (I have described this scenario on my reverse proxy guide). The purpose of the DMZ is to limit exposure to your local network if your public facing services are compromised.
Firewall: Rules: [USER]
The USER network can be used for PCs, laptops, or phones (if you do not want phones in your IOT network). The primary purpose is to separate these devices from potentially more vulnerable IOT devices. The rules will be similar to what has been shown until this point except for rules to allow access to devices/services located on other networks.
For the allow rule for the printer, I set the destination port to any to keep it simple but printers can often require several ports depending on the type of printer. For instance, HP printers have a list of ports that you may add in order to further restrict access.
| Action | TCP/IP Version | Protocol | Source | Dest / Invert | Destination | Dest Port | Description | 
|---|---|---|---|---|---|---|---|
| Pass | IPv4+6 | TCP/UDP | USER net | unchecked | USER address | 53 (DNS) | Allow access to DNS | 
| Pass | IPv4 | TCP | USER net | unchecked | WebServer | 443 (HTTPS) | Allow access to web app | 
| Pass | IPv4 | TCP/UDP | USER net | unchecked | Printer | any | Allow access to printer | 
| Pass | IPv4+6 | any | USER net | checked | PrivateNetworks | any | Allow access only to Internet | 
Ideally, you should have a dedicated device (or VM perhaps) residing on the LAN that has access to all of the web interfaces of your network infrastructure. If you do not have any dedicated devices available, you could create a rule on the USER network to allow your PC/laptop to access the management interfaces. This is less than ideal because you are poking a hole in the management network, but at least access is still restricted from a specific device to specific web interfaces and ports. Security is still better than a flat network with full access to everything.
Firewall: Rules: [IOT]
The IOT network is where you can put less trusted IOT (Internet of Things) type devices or perhaps devices that no longer receive security updates – anything that is likely more vulnerable than your other devices which are regularly updated (PCs, laptops, etc). The reason for such a network is that IOT devices have been notoriously prone to vulnerabilities. They are often developed as cheap, convenient devices which do not receive thorough security audits and are not updated as frequently (or at all if users do not perform updates).
If your phones are on the IOT network, you will likely want the second rule below to allow access to your web app(s) on your DMZ network. To make the rule even more restrictive, you could create an alias that contains the IP addresses of just your phones. To make this work easily, reserving a static IP address via DHCP is one recommendation since you do not have to set it manually on your phone (make sure you turn off private WiFi address options for your local WiFi connection so your MAC address does not change).
The third rule is useful when you want to access your camera feeds from your phone even if you put your IP cameras on their own isolated network which does not have even Internet access. As I mention in the IPCAM section below, any time you open a hole into the camera network, it introduces one more way a malicious user can get in, but I think the risk is pretty low if you limit access and have other protections in place (IDS/IPS, firewall rules, passwords, etc).
| Action | TCP/IP Version | Protocol | Source | Dest / Invert | Destination | Dest Port | Description | 
|---|---|---|---|---|---|---|---|
| Pass | IPv4+6 | TCP/UDP | IOT net | unchecked | IOT address | 53 (DNS) | Allow access to DNS | 
| Pass | IPv4 | TCP | IOT net | unchecked | WebServer | 443 (HTTPS) | Allow access to web app | 
| Pass | IPv4 | TCP | IOT net | unchecked | IPCameras | 554 | Allow access to IP camera feeds | 
| Pass | IPv4+6 | any | IOT net | checked | PrivateNetworks | any | Allow access only to Internet | 
Firewall: Rules: [GUEST]
The guest network is nice to have if you want to give out access to your WiFi network to guests. You do not know what kind of infestations those users may bring onto your secured network! This network provides you a place that is separate from your network like a DMZ but for untrusted devices that you do not personally own. You can simply allow only Internet access and block access to your local networks (similar to the DMZ – unless the DMZ has access to a dedicated app/services network as I mentioned above) or you could also allow them to access your printer as shown below.
| Action | TCP/IP Version | Protocol | Source | Dest / Invert | Destination | Dest Port | Description | 
|---|---|---|---|---|---|---|---|
| Pass | IPv4+6 | TCP/UDP | GUEST net | unchecked | GUEST address | 53 (DNS) | Allow access to DNS | 
| Pass | IPv4 | TCP/UDP | GUEST net | unchecked | Printer | any | Allow access to printer | 
| Pass | IPv4+6 | any | GUEST net | checked | PrivateNetworks | any | Allow access only to Internet | 
Firewall: Rules: [IPCAM]
