---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/how-to-beginners-guide-to-set-up-home-network-using-opnsense-62d9c4eb-1
title: "how-to-beginners-guide-to-set-up-home-network-using-opnsense-62d9c4eb"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/how-to-beginners-guide-to-set-up-home-network-using-opnsense-62d9c4eb.md
source_anchor: ""
source_lines: [1, 41]
sha256: a18ea7d7ddafee66c08515eb106e7d78a2c9252c6dfe07d3864e5f0b9019a5d5
---

# how-to-beginners-guide-to-set-up-home-network-using-opnsense-62d9c4eb

Beginner's Guide to Set Up a Home Network Using OPNsense
Table of Contents
My most popular guide at the time of this writing is how to set up a full network using OPNsense. In that guide, I combine many of the concepts I have written about over the years. The guide has been generally well received but as more new users started following that guide, some confusion has arisen about the concept of LAGs and when it is appropriate to use them.
A Link Aggregation Group (LAG) is useful for providing redundancy if one or more interfaces fail but also can be helpful in increasing bandwidth of multiple simultaneous streams of data. For 1 Gbps interfaces, they are much easier to saturate on a regular basis than higher speed interfaces. Normal day to day usage will likely not saturate 1 Gbps. Of course, downloading and copying data between two computers can easily saturate 1 Gbps. LAGs may be helpful for performance on slower links like 1 Gbps only if you have more than 1 device on your network saturating 1 Gbps at the same time.
Many users choose to implement a LAG without knowing for certain if they actually need it. Perhaps I did not spell that out clearly enough in my original guide. I recommend implementing a LAG on your OPNsense box only if you know for certain you have bottlenecks for network traffic going across your various internal networks and you know that your OPNsense box is capable of handling the increased capacity. LAGs can be used for redundancy/high availability purposes but to be honest, in the past I used the same cheap mini-PC appliance for OPNsense for over 5 years without any hardware failures so it is not something which occurs frequently in my experience.
As 2.5/10G interfaces become more affordable, there is less need to utilize LAGs on your network (except maybe for redundancy/high availability purposes). Therefore, I recommend taking advantage of those higher speed interfaces on your network to alleviate any potential bottlenecks. I personally no longer have any LAGs configured since I have introduced 10 Gbps interfaces on my network.
The Intention of the Original Guide
The intention of the original guide was to simply demonstrate several concepts in one comprehensive guide so that new users did not have to piece together a bunch of guides on specific topics hoping to make everything work together. The original guide was aimed more toward users who had a good understanding of networking but wanted to go beyond a basic home network to help improve security, privacy, etc. Therefore, the original guide included more configuration than what many new users would likely need to implement to simply get started.
For instance, new users may not need or want to create 5 VLANs on their networks. I included 5 VLANs as examples of the different types of VLANs that users may wish to implement on their networks so they could pick and choose which ones are appropriate to use for their home networks. In addition, LAGs (link aggregations) are completely unnecessary unless the purpose of utilizing LAGs is fully understood and there is good reason to implement them.
I have always advocated for new users to start with the basic configuration before working towards a more complex configuration. By taking one step at a time slowly, accomplishing goals becomes much easier and attainable. Implementing a more complex network all at once may be an overwhelming and frustrating experience if issues arise. Understanding the fundamentals first before moving forward is a great approach to learning and reaching goals.
Although the primary focus of this website is generally aimed toward intermediate to more advanced home network users, I try to also make the content approachable to new users the best I can. However, in order to keep guides on topic and constrained, I do have to assume that some basic level of understanding of networking has to exist when creating guides.
With that said, I have decided to create this alternate version of setting up a full network which has a greatly simplified network architecture to help beginners get started with using OPNsense to build out a full network.
I will be using a 4 port mini-PC such as the Protectli VP2420 (affiliate link) , a managed network switch such as the TP-Link T1500G-10MPS (affiliate link) , and a wireless access point capable of supporting VLANs such as the Grandstream GWN7660 (affiliate link) .
To follow along, you will need similar equipment. As I mentioned in the original guide, the way VLANs are configured for the network switch and wireless access point may vary depending on which vendor you are using so it is not possible for me to demonstrate how to do the configuration for every type of device that exists. However, once you understand the concepts of VLANs, you should be able to implement it on your specific hardware.
Note
Disclaimer: As stated in the original guide: This is yet another network architecture example you may use as a reference. I am not claiming that this architecture is the best way to implement your home network. You know your needs the best and you should be able to choose or omit portions of this guide.
Network Architecture
The example network will assume the following architecture:
- The connection from the ISP will be connected directly to the WAN interface of the OPNsense system
- The LAN interface will be connected to a smart/managed network switch and will also contain a single VLAN for untrusted devices (a router on a stick configuration)
- A wireless access point will be connected to the network switch to provide wireless for the networks
- Other devices will be connected to the network switch and will reside either in the trusted LAN or the untrusted VLAN
- The network will only use IPv4 to keep configuration simple (see the IPv6 configuration guide or the original full network guide if you wish to implement IPv6)
To keep the configuration minimal, only a single physical network interface will be used for the LAN and the VLAN which will contain all of the untrusted devices.
Even though you could utilize two different physical interfaces to accomplish the separation between trusted and untrusted devices, I want to introduce the concept of creating a single VLAN because this is a good foundational concept to learn in order to build out a more complex network. Once you understand how to create one VLAN, you will be able to create additional VLANs in the future for other purposes as you build out your network.
Logical Network Diagram
Below is a logical network diagram that will be implemented in by this guide:
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
The first step you need to do is install OPNsense on your hardware. I am using a Protectli VP2420 in my network example, but you can use any hardware compatible with OPNsense.
Go to the OPNsense download page. Choose the “vga” installer so you can install it on a USB drive. Use a program such as Etcher to flash the drive with the image you just downloaded (you do not even have to extract the image if you are using Etcher).
