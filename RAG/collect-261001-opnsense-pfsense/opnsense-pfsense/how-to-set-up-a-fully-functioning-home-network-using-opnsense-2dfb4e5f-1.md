---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/how-to-set-up-a-fully-functioning-home-network-using-opnsense-2dfb4e5f-1
title: "how-to-set-up-a-fully-functioning-home-network-using-opnsense-2dfb4e5f"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/how-to-set-up-a-fully-functioning-home-network-using-opnsense-2dfb4e5f.md
source_anchor: ""
source_lines: [1, 48]
sha256: b8b40bd8e8a4950b82174b3d71c9649fd07f697e8764f90c3e659b8e80a7c1b2
---

# how-to-set-up-a-fully-functioning-home-network-using-opnsense-2dfb4e5f

Set Up a Fully Functioning Home Network Using OPNsense
Table of Contents
Since the publication of this guide, I have written a new guide which is aimed more towards beginners. The beginner’s guide focuses more on the bare minimum requirements to get started with building a full network using OPNsense. In that guide, I only demonstrate using the LAN interface with a single VLAN where untrusted devices may be placed. I feel this would be a good first step in learning how to build a more advanced home network since it provides a basic foundation on which to build upon.
I often write about specific networking topics in an effort to limit the scope of each topic being discussed. However, I thought it would be helpful to create one comprehensive guide which pulls together many concepts to demonstrate how all the pieces fit together using a real world example.
While beginners may follow this guide, this guide is aimed towards individuals which already have some background knowledge on networking but wish to take their home networks to the next level.
Much of this guide will be focused on OPNsense configuration since it lies at the core of the network infrastructure, but I will also discuss managed switches and wireless access points.
When it comes to the network switches and wireless access points, I cannot include an example of all varieties of those products. Therefore, you will have to adapt this guide to the hardware you are using. I personally prefer to use managed network switches which have a web interface for configuration rather than switches which only provide a command line interface (even though I am not personally afraid of using the command line).
Because there is so much to cover, I will attempt to be more concise than usual in certain areas. In some cases I may refer to other guides I have written if you wish to go into more depth about a particular topic.
Disclaimer: Please consider that I am merely presenting one network architecture that may be used as a reference. I am not endorsing or claiming that this architecture is the best way to implement your home network – it is simply one way you could implement it. In the end, you have to decide what meets your needs the best.
Network Architecture
The example network will assume the following architecture:
- The connection from the ISP will utilize a modem or a modem/router in bridge mode (to avoid double NAT)
- The modem is connected to the first port of a four port OPNsense box for the WAN interface
- The LAN interface will be connected to a smart/managed network switch to be used as a dedicated management interface and will only contained untagged network traffic
- Two interfaces on the OPNsense box will be set up as a LAG (Link Aggregation Group) to the network switch to provide extra bandwidth/redundancy across networks and will only contain tagged VLAN traffic
- A wireless access point with VLAN support will be connected to the network switch to provide wireless for the networks
- Other devices will be connected to the network switch in various VLANs
- The network will use dual stack IPv4 & IPv6*
In the example described above, it is important to note that the management LAN interface (where all of the critical network infrastructure will be managed) will only have untagged network traffic while the LAG interface will only have tagged VLAN traffic.
Note
Some users prefer to set up a dedicated management VLAN where all network infrastructure can be managed and avoid using the default untagged network, which is often depicted as “VLAN 1” on network switches. The primary reason is the security concern of users plugging devices into ports that are not assigned to any VLANs since they would have access to the most sensitive part of the network.
When using the default “VLAN 1” as your management network, you need to be sure to assign all unused ports to a VLAN to avoid unwanted devices on your management VLAN. You could assign unused ports to the GUEST network or to a sinkhole network with no network access. You could even disable the ports on the switch but that seems a bit extreme and inconvenient.
For home networks in particular, this risk is likely much lower than on business/enterprise networks especially since most guests in your home typically use WiFi rather than plugging directly into your network switch (or a wall jack which is connected to your network switch).
Regardless of whether or not you use the untagged LAN or have a separate VLAN for managing your critical network infrastructure, the OPNsense team recommends separating untagged and tagged VLAN traffic so they are not on the same interface to prevent potential security issues due to carelessly written firewall rules on the parent interface. This actually caught several users by surprise including myself.
Do not worry because this guide will demonstrate separating untagged and untagged interfaces by using one of the interfaces on the router/firewall box as a dedicated management interface and the remaining interfaces for VLANs only.
Note
* For the IPv6 configuration demonstrated in this guide, it will be based on the assumptions that the ISP is providing dynamic IPv6 addresses and that IPv6 will only be configured to allow Internet access. Similar to IPv4, the IPv6 networks will be configured to isolate them from each other.
I will not demonstrate using IPv6 for your internal networks in this guide. There are some challenges with using dynamic IPv6 addresses on your internal network especially when IPv6 addresses change. For this reason and more, I will only use IPv4 examples for allowing access between devices. IPv6 configuration will be limited to Internet access only.
Logical Network Diagram
To have a visual representation of what this guide will implement, I have included the following logical network diagram.
Note
Please keep in mind this is a logical diagram so the LAG which has the 5 associated VLANs is only using 2 physical interfaces rather than 3 interfaces. Some users may have interpreted this diagram as “use all of your remaining interfaces in the LAG”.
Depending on your needs, you will likely not see many benefits going beyond 2 interfaces. Even 2 interfaces might be overkill for certain networks if you do not have a lot of traffic going across VLANs.
If you have 2.5 Gbps interfaces or faster on your OPNsense box, you may not need to use any LAGs since you have more bandwidth available, and therefore you can skip over the LAG section in this guide.
Physical Diagram of Network Infrastructure
Refer to the image below for the physical connections of the network infrastructure that I will be using.
To keep the physical network infrastructure image less cluttered, the image below only shows the basic network infrastructure described in my example and not the various devices connected to the network (see also the physical diagram of connected devices).
TL;DR Overview
Because this guide will be longer than my usual content, below is an outline of what will be done to complete the full network example I am discussing:
- Install OPNsense
- Configure OPNsense
- Configure network switch
- Configure wireless access point
The full setup may seem daunting, but I am including a lot of configuration information with brief descriptions in an effort to show all of the necessary or recommended values to build a fully functioning network with OPNsense.
Let us get started with installing OPNsense!
Install OPNsense
The first step you need to do is install OPNsense on your desired hardware. I am using a Protectli VP2410 in my network example, but you can use any hardware compatible with OPNsense.
Go to the OPNsense download page. Choose the “vga” installer so you can install it on a USB drive. Use a program such as Etcher to flash the drive with the image you just downloaded (you do not even have to extract the image if you are using Etcher).
