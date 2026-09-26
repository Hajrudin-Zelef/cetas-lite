---
id: collect-260926-mikrotik/mikrotik/arborescence-des-pages-7-1
title: "Introduction"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet", "parameters"]
source: docs/RAG/lot-mikrotik/RouterOS/arborescence-des-pages-7.md
source_anchor: ""
source_lines: [1, 92]
sha256: 6aecda9da1caf9c35f5ddf8de79dcbf8ded93f820be1e597fa71164730fd6c19
---

# Introduction

The ZeroTier network hypervisor is a self-contained network virtualization engine that implements an Ethernet virtualization layer similar to VXLAN built atop a cryptographically secure global peer-to-peer network. It provides advanced network virtualization and management capabilities on par with an enterprise SDN switch, but across both local and wide area networks and connecting almost any kind of app or device.

MikroTik has added ZeroTier to RouterOS v7.1rc2 as a separate package for the **ARM/ARM64** architecture. 

Wait, so what can I use it for?

- Hosting a game server at home (useful for LAN only games) or simply creating a LAN party with your friends;
- Accessing LAN devices behind NAT directly;
- Accessing LAN devices via SSH without opening port to the Internet;
- Using your local Pi-Hole setup from anywhere via the Internet;

Important

## Video tutorial

# Required Network Configuration

## What ports does ZeroTier use?

It listens on three 3 UDP ports:

- 9993 - The default
- A random, high numbered port derived from your ZeroTier address
- A random, high numbered port for use with UPnP/NAT-PMP mappings

That means your *peers* could be listening on any port. To talk with them directly, you need to be able to send them to any port.

## Recommended Local Network and Internet Gateway Configuration

These ZeroTier recommended guidelines are consistent with the vast majority of typical deployments using commodity gateways and access points:

- Don't restrict outbound UDP traffic.
- Supporting either UPnP or NAT-PMP on your network can greatly improve performance by allowing ZeroTier endpoints to map external ports and avoid NAT traversal entirely.
- IPv6 is recommended and can greatly improve direct connection reliability if supported on both ends of a direct link. If present it should be implemented without NAT (NAT is wholly unnecessary with IPv6 and only adds complexity) and with a stateful firewall that permits bidirectional UDP conversations.
- Don't use "symmetric" NAT. Use "full cone" or "port restricted cone" NAT. Symmetric NAT is extremely hostile to peer-to-peer traffic and will degrade VoIP, video chat, games, WebRTC, and many other protocols as well as ZeroTier.
- No more than one layer of NAT should be present between ZeroTier endpoints and the Internet. Multiple layers of NAT introduce connection instability due to chaotic interactions between states and behaviors at different levels. **No Double NAT.**
- NATs should have a port mapping or connection timeout no shorter than 60 seconds.
- Place no more than about 16,000 devices behind each NAT-managed external IP address to ensure that each device can map a sufficient number of ports.
- Switches and wireless access points should allow direct local traffic between local devices. Turn off any "local isolation" features. Some switches might allow finer-grained control, and on these, it would be sufficient to allow local UDP traffic to/from 9993 (or in general).

# Configuration example

By default, ZeroTier is designed to be zero-configuration. A user can start a new ZeroTier node without having to write configuration files or provide the IP addresses of other nodes. It’s also designed to be fast. Any two devices in the world should be able to locate each other and communicate almost instantly so the following example will enable ZeroTier on RouterOS device and connect one mobile phone using the ZeroTier application.

1. Register on my.zerotier.com and **Create A Network** , obtain the*Network ID* , in this example:*1d71939404912b40* ;
2. Download and Install ZeroTier NPK package in RouterOS, you can find under in the "Extra packages", upload package on the device and reboot the unit;
3. Enable the default (official) ZeroTier instance:
4. Add a new network, specifying the network ID you created in the ZeroTier cloud console:
5. Verify ZeroTier configuration:
6. Now you might need to allow connections from the ZeroTier interface to your router, and **optionally** , to your other LAN interfaces:
7. Install a ZeroTier client on your smartphone or computer, follow the ZeroTier manual on how to connect to the same network from there.
8. If **"****Access Control"** is set to**"Private"** , you must authorize nodes before they become members:

### Peer

ZeroTier`s peer is an informative section with a list of nodes that your node knows about. Nodes can not talk to each other unless they are joined and authorized on the same network.

# Parameters

| Property | Description | 
|---|---|
| **name**  (s*tring* ; default:**zt1** ) | Instance name. | 
| **port** (*number;*  default:**9993** ) | Port number the instance listen to. | 
| **identity** (*string* ; default)*sensitive* | Instance 40-bit unique address. | 
| **interface** (string; default:**all)** | List of interfaces that are used in order to discover ZeroTier peers, by using ARP and IP type connections. | 
| **route-distance** (number; default:**1** ) | Route distance for routes obtained from planet/moon servers. | 

| Property | Description | 
|---|---|
| **allow-default** (*string; yes \| no)* | A network can override the systems default route (force VPN mode). | 
| **allow-global**  (*string; yes \| no)* | ZeroTier IP addresses and routes can overlap public IP space. | 
| **allow-managed**  (*string; yes \| no)* | ZeroTier managed IP addresses and routes are assigned. | 
| **arp-timeout** (*number* ; default:**auto** ) | ARP timeouts value. | 
| **comment** (*string* ; Default: ) | Descriptive comment for the interfaces. | 
| **copy-from**  | Allows copying existing interfaces configuration. | 
| **disable-running-check**  (*string; yes \| no)* | Force interface in "running" state. | 
| **instance**  (*string* ; Default: **zt1** ) | ZeroTier instance name. | 
| **name**  (s*tring* ; default:**zerotier1** ) | A short name. | 
| **network**  (*string* ; Default) | 16-digit network ID. | 

# Controller

RouterOS implements ZeroTier functionality in the role of a node where most of the network configuration must be done on the ZeroTier webpage dashboard. However, in situations where you would prefer to do all the configuration on your own device, RouterOS offers to host your own controller

A common misunderstanding is to conflate network controllers with root servers (planet and moons). Root servers are connection facilitators that operate at the **VL1 level**. Network controllers are configuration managers and certificate authorities that belong to the **VL2 level.** Generally, root servers don’t join or control virtual networks and network controllers are not root servers, though it is possible to have a node do both.

Every ZeroTier instance has a self-hosting network controller that can be used to host virtual networks. __A controller is responsible for admitting members to the network, and issuing default configuration information including certificates.__ Controllers can in theory host up to 2^24 networks and serve many millions of devices (or more), but we recommend spreading large numbers of networks across many controllers for load balancing and fault tolerance reasons.

## Parameters

