---
id: collect-261001-general-networking/general-networking/vmps-put-a-fork-in-it-1
title: "vmps-put-a-fork-in-it"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-general-networking/vmps-put-a-fork-in-it.md
source_anchor: ""
source_lines: [1, 50]
sha256: 3ebf61e6a93b8f5a6c169a1eb7c8961c946940a47644643fff3ed6594e5ad538
---

# vmps-put-a-fork-in-it

VLAN Membership Policy Server (VMPS) was a technology that, at a point in time, provided a way for organizations to control access to their networks. A few organizations embraced it and after a few years they found it to be an albatross around their neck. However, getting away from VMPS was not so easy and the longer they delay, the harder the breakup becomes.

VMPS is a Cisco proprietary solution for providing users access to a network based on their MAC address. VMPS is available on Cisco 4000/4500/5000/6000/6500 switches. As part of the configuration of VPMS a large table is constructed with MAC addresses and VLANs of valid network attached devices. This VMPS database is placed on a switch in the network environment that acts like a VMPS server for other switches to query. When computers turns on and activate their NICs the access switches in the environment use VLAN Query Protocol (VQP) that uses UDP port 1589 to determine the VLAN name for the end-user’s MAC address and then that user is assigned to that VLAN. The VMPS access system dynamically assigns an Ethernet switchport to a specific VLAN based on the user’s MAC address. This database is maintained manually and, based on the size of the organization, is often times updated anywhere from 1 to 10 times each day. The VMPS download server is a Cisco Ethernet switch and the vmps.cfg file is a file that is transmitted with TFTP to the VMPS servers. It is also possible to have a primary and a backup VMPS server for high availability.

If you are curious here is the Cisco guide for configuring VMPS on a 6500 running CatOS 8.7.

Here is a guide for Troubleshooting the Catalyst VMPS Switch.

Many others have written descriptions of VMPS. From Sean Convery’s Cisco Press Book here is what he says this about VMPS.

**Issues with VMPS:**

This system does have a couple of shortcomings. If an attacker knew what they were doing they could still assign a static IP address to their system with a locally administered MAC address to subvert this system. This system also has lots of historical MAC addresses entered into it that haven’t been removed. Therefore, an audit needs to be performed of the MAC addresses that are still valid. One way to do this would be to pull the CAM table today and then saying everything in the network today is permitted.

One issue with VMPS servers is that this protocol consumes a lot of processor power depending on the number of users and the size of the database. I have seen organizations with literally thousands of users and many thousands of entries in the database. That is because few organizations have the discipline to remove anything from the database once it is entered. If you are using VMPS and you are noticing these messages on your VMPS server then you may have a problem. The logs can show many messages related to VQP having issues. Messages like this can be abundant in the syslogs or on a CiscoWorks server.

2009 Jan 05 15:46:26 MDT -07:00 %IP-6-UDP_SOCKOVFL:UDP socket overflow from Source IP: 192.168.126.92, Destination port: 1589 The Cisco Error Message Decoder says that these messages are related to the following problems.

1. %IP-6-UDP_SOCKOVFL: UDP socket overflow from Source IP:[chars], Destination port:[dec] This message indicates that all buffers for a UDP socket on the Network Management Processor (NMP) have filled up due to excessive UDP traffic on the administrative VLAN and cannot store additional traffic. [chars] is the source IP address and [dec] is the destination port number.

Recommended Action: Remove or block the source of the UDP packets to prevent further UDP packet loss. Note Kernel messages do not indicate a problem with system performance but should be reported to your technical support representative.

Other issues with VMPS involve the vmps.txt file consuming space on the flash filesystem of the VMPS server. This experience is documented on Jonboys Blog .

The problem arises in the fact that in order to use VMPS you have to use CatOS. VMPS is not supported in Cat IOS. Therefore, any organization that has chosen VMPS has been stuck without an upgrade path. This technology obsolescence problem has prevented many organizations using VMPS from being able to use any newer feature in Cat IOS. Furthermore, VMPS has now been deprecated by Cisco and it is not recommended for customer use. In fact, organizations that use VMPS are frozen on the version of software currently used on their switches because newer versions don’t support VMPS. Therefore, organizations using VMPS are forced to choose between VMPS and security if a vulnerability is found on the current version of switch software.

**DISA has Dissed VMPS:**

The Defense Information Systems Agency (DISA) Security Technical Implementation Guides (STIGs) for Network Security doesn’t recommend its use for DOD organizations. The DISA STIG has several paragraphs on VPMS on pages 76-78. The concluding sentence ends with “For these reasons, the U.S. DOD believes that VMPS must not be used to provide port authentication or dynamic VLAN assignment.”

**What Options do VMPS Users Have?**

There are several other alternatives to VMPS and organizations are encouraged to explore other options that will have future support by Cisco and are also based on industry standards. What organizations are really looking for is a way to easily perform Network Access Control (NAC) and prevent unauthorized users from accessing the network. There are numerous styles and flavors of NAC solutions on the market. Because there are many different approaches it may be hard to differentiate them. Each system has a different way of assessing the security of the end-point and they also have different ways of containing security threats. The table below covers the handful of basic forms of policy enforcement that NAC products use.

**Enforcement Techniques:**

IEEE 802.1X The Ethernet switch only opens up the port if the end-point is properly authenticated and healthy. This is common for wireless LANs. The downside is that each end-point needs a supplicant.

VLAN Steering This technique assigns the user’s LAN switchport to a specific VLAN (guest, remediation, intranet,…). The command and control function of the NAC system must interact with the LAN switch.

DHCP Lease Management The NAC system controls the IP address that the end-point receives through DHCP

ARP Poisoning Uses ARP to control which hosts can communicate by modifying the binding of IP addresses to MAC addresses

DNS Redirection Redirects all DNS requests toward the web portal to guide a user to the authentication system

Inline Blocking A NAC system that is in-line between the computer and the core of the network can stop a specific client from communicating with the rest of the network. The closer the NAC system is to the end-point the more granular the control.

DHCP and ARP Poisoning are two of the techniques used to control which endpoints get access to the network. The NAC system first puts a host onto a private network so that it can be assessed, then changes the host’s IP as needed. DHCP control requires little change to the underlying infrastructure and is less invasive than switch-port manipulation, VLAN steering or dynamically updating router ACLs. ARP poisoning, on the other hand, uses ARP to manage the MAC-to-IP mapping used by network hosts to communicate within a single subnet. If a host sends out an ARP packet saying it’s the network router, for example, all endpoints on that segment will send it all packets bound for other segments (note that the concept is called ARP poisoning whether it’s used for good or evil).

