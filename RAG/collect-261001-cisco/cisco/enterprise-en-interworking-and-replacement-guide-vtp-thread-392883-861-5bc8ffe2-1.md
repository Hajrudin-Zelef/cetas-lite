---
id: collect-261001-cisco/cisco/enterprise-en-interworking-and-replacement-guide-vtp-thread-392883-861-5bc8ffe2-1
title: "Run the show running-config command to check the interface configuration."
domain: cisco
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/enterprise-en-interworking-and-replacement-guide-vtp-thread-392883-861-5bc8ffe2.md
source_anchor: ""
source_lines: [1, 47]
sha256: c966e2549b64ab857b82f4eed6a1366f46f703a903a12b16f8454742556ad9e9
---

# Run the show running-config command to check the interface configuration.

Hi there, Community!
This post is about an Interworking and Replacement Guide VTP. Please see more details about the topic below.
The VLAN Central Management Protocol (VCMP), a Layer 2 protocol in the Open System Interconnection (OSI) model, transmits VLAN information and ensures consistent VLAN information on the Layer 2 network. VCMP is a Huawei proprietary protocol. Huawei S series switches support VCMP starting from V200R005.
Purpose
In most cases, switches on an enterprise network need to synchronize VLAN information with each other to ensure that they can correctly forward data. On a small-scale enterprise network, the network administrator can log in to each switch to configure and maintain VLANs. On a large-scale enterprise network, a lot of switches are deployed, so a large amount of VLAN information needs to be configured and maintained. If the network administrator manually configures and maintains all VLANs, the workload is heavy and VLAN information may be inconsistent.
To address the preceding issue, VCMP is used to implement centralized VLAN management. The network administrator needs to create and delete VLAN information only on one switch. The changes on the switch are automatically synchronized to other switches in a specified scope so that no manual operation is required on these switches. In this way, the configuration workload is reduced and VLAN information consistency is ensured.
VCMP can only help the network administrator synchronize VLAN information but not dynamically assign VLANs. VCMP is often used with Link-type Negotiation Protocol (LNP) to simplify user configurations.
VCMP that is configured on a switch of a Layer 2 network brings in the following benefits:
l Implements centralized VLAN management and maintenance, and reduces the network maintenance workload.
l Implements the plug-and-play function of access switches.
Basic Concepts
VCMP involves two important entities:
l VCMP domain
A VCMP domain is composed of switches that have the same VCMP domain name and are connected through trunk or hybrid interfaces. All switches in the VCMP domain must use the same domain name, and each switch can join only one VCMP domain. Switches in different VCMP domains cannot synchronize VLAN information.
A VCMP domain specifies the scope for the administrative switch and managed switches. Switches in a VCMP domain are managed by the administrative switch. There is only one administrative switch and multiple managed switches in a VCMP domain.
l VCMP roles
VCMP determines attributes of switches based on VCMP roles. Table 1-1 describes VCMP roles.
Table 1-1 VCMP roles
| VCMP Role | Description | 
| Server | The VCMP server synchronizes VLAN information to other switches in the local VCMP domain. The VLAN information that is created and deleted on the VCMP server is broadcast in a VCMP domain. | 
| Client | A VCMP client belongs to a specified VCMP domain and synchronizes VLAN information with the VCMP server. The VLAN information that is created and deleted on a VCMP client is not broadcast in a VCMP domain, but is overwritten by VLAN information sent by the VCMP server. Huawei S series switches are VCMP clients by default. | 
| Transparent | A VCMP transparent switch is not affected by VCMP management behaviors, and does not affect other switches in the local VCMP domain. The VCMP transparent switch transparently forwards VCMP packets to only trunk or hybrid links. The VLAN information that is created and deleted on a VCMP transparent switch is not affected by the VCMP server and is not broadcast in a VCMP domain. In this way, some switches that do not need to be managed by VCMP can forward VCMP packets. | 
| Silent | Deployed at the edge of a VCMP domain, a VCMP silent switch does not affect other switches in the local VCMP domain and is not affected by VCMP management behaviors. The VCMP silent switch prevents VCMP packets in a VCMP domain from being transmitted to other VCMP domains. A VCMP silent switch directly discards received VCMP packets but not forward them. The VLAN information that is created and deleted on a VCMP silent switch is not affected by the VCMP server and is not broadcast in a VCMP domain. | 
VCMP Packets
VCMP enables switches of different roles to exchange VCMP packets to implement centralized VLAN management. VCMP packets can be only transmitted in VLAN 1 on trunk or hybrid interfaces. To retain the same VLAN information on the VCMP server and clients, VCMP defines two types of multicast packets: Summary-Advert and Advert-Request. Table 1-2 describes the functions and applicable scenarios of the two types of packets.
Table 1-2 VCMP packets
| Packet Type | Function | Applicable Scenario | 
| Summary-Advert | The VCMP server sends Summary-Advert packets to other devices in the local VCMP domain to notify them of the domain name, device ID, configuration revision number, and VLAN information. | l The VCMP server sends a Summary-Advert packet every 5 minutes to ensure real-time synchronization of VLAN information on the VCMP server and clients and to prevent VLAN information loss due to packet loss. l The VCMP server configuration is changed. For example, VLANs are created or deleted, the VCMP domain name or device ID is changed, and the VCMP server restarts. l The VCMP server receives Advert-Request packets from VCMP clients in the same VCMP domain. | 
| Advert-Request | A VCMP client sends Advert-Request packets to the VCMP server to request VLAN information. | l A VCMP client is added. l A VCMP client restarts or a client interface becomes Up. l A VCMP client changed from a VCMP server, silent, or transparent sends Advert-Request packets to the VCMP server. If the VCMP server does not respond, the VCMP client sends 10 Advert-Request packets every 5 minutes. | 
l Figure 1-1 shows the format of a Summary-Advert packet.
− Code: indicates a Summary-Advert packet when the value is 0x01.
− Followers: is not used currently, and is set to 0.
− Updater Identity: indicates the VCMP server ID.
− Configuration Revision Number: determines whether VLAN information sent from the VCMP server is newer than the local VLAN information. The VCMP client determines whether to synchronize VLAN information with the VCMP server based on this field. The value is an 8-digit hexadecimal number. The four left-most bits indicate the change of the VCMP domain or device ID and the four right-most bits indicate the VLAN change. Upon a VLAN change on the VCMP server, the configuration revision number is automatically increased. When the VCMP domain name or device ID changes, the four left-most bits of the configuration revision number are recalculated and the four right-most bits are reset.
− Digest: indicates the packet digest. When a VCMP domain is configured with a password, the switch uses the password and bits 25 to 62 of packets to calculate the digest. When no password is configured, the switch uses bits 25 to 68 of packets to calculate the digest.
− VLAN bitmap: indicates the VLAN information on the VCMP server.
Figure 1-1 Format of a Summary-Advert packet
l Figure 1-2 shows the format of an Advert-Request packet.
− Code: indicates an Advert-Request packet when the value is 0x02.
− Reserved: indicates the reserved field that has a fixed value of 0.
Figure 1-2 Format of an Advert-Request packet
The VLAN Trunking protocol (VTP), a layer 2 protocol in the Open System Interconnection (OSI) model, manages VLAN creation, deletion, and renaming in a domain. VTP is a Cisco proprietary protocol.
VTP involves two important entities:
l VTP domain
A VTP domain is composed of switches that have the same VTP domain name and are connected through trunk links.
Switches in a VTP domain share VLAN information, and each switch can join only one VTP domain. Switches in different VTP domains cannot share VLAN information.
l Working mode
