---
id: collect-261001-general-networking/general-networking/infrastructure-security-and-segmentation-3
title: "infrastructure-security-and-segmentation"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-general-networking/infrastructure-security-and-segmentation.md
source_anchor: ""
source_lines: [152, 216]
sha256: c86ae185b6317f1a92082be612aaec244f6f90339aff4709b5a6f486b26aaace
---

# infrastructure-security-and-segmentation

While forcing an association change in the CAM table by spoofing the source MAC address is effective, it is not very practical. The device that actually owns the MAC address may send a frame any time and update the table. ARP poisoning, on the other hand, is easier to execute and lasts a longer time.

With ARP poisoning, the ARP table of a router, switch, or target host is changed to map an IP address to the MAC address of a different host. This causes IP packets to be switched to a rogue device instead of the intended destination.

Two methods can be used to poison the ARP table. The first involves using a forged ARP response to a broadcasted query. The second involves using a Gratuitous ARP (GARP) packet to announce a false IP-to-MAC address mapping. The new mapping announced with the GARP packet overwrites the existing one and causes the packets to be delivered to the newly mapped MAC address.

ARP poisoning attacks can be prevented with the DAI security feature of a Cisco switch. When this feature is enabled, the switch checks all ARP packets against the DHCP snooping database to validate mappings being announced. If a mapping does not match the information in the database, it is dropped. Just as with DHCP snooping, interfaces can be configured as trusted or untrusted for DAI. Only ARP packets from untrusted interfaces are inspected. If the network segment contains hosts that do not use DHCP, an ARP access list must be configured for DAI to use.

DAI can be enabled for a VLAN with the **ip arp inspection vlan** *vlan-ID* command. *VLAN-ID* can be a single VLAN, a comma-separated list, or a range.

An ARP access list is created using the **arp access-list** *acl-name* command. Within the access list, individual entries are defined using the **permit ip host** *ip-address* **mac host** *mac-address* command. The access list can be applied to a VLAN with the **ip arp inspection filter** *filter-name* **vlan** *vlan-ID* command.

Example 2-45 shows DAI configuration on a switch. In this example, DAI is enabled on VLAN 1, an ARP inspection access list is applied for a static IP address, and Gi0/10 uplink is configured as a trusted interface.

#### **Example 2-45** *Configuring Dynamic ARP Inspection*

`SW1(config)#**ip arp inspection vlan 1**
SW1(config)#**arp access-list static-map**
SW1(config-arp-nacl)#**permit ip host 192.168.1.24 mac host 1001.1001.1001**
SW1(config-arp-nacl)#**exit**
SW1(config)#**ip arp inspection filter static-map vlan 1**
SW1(config)#**interface Gi0/10**
SW1(config-if)#**ip arp inspection trust**`

DAI configuration and operation can be verified with the **show ip arp inspection** command, as shown in Example 2-46.

#### **Example 2-46** *Verifying Dynamic ARP Inspection*

`SW1#**show ip arp inspection**
—removed for brevity—
 Vlan     Configuration    Operation   ACL Match         Static ACL
 ----     -------------    ---------   ---------         -----------
    1     Enabled          Active      static-map         No
 Vlan     ACL Logging      DHCP Logging      Probe Logging
 ----     -----------      ------------      -------------
    1     Deny             Deny              Off
 Vlan     Forwarded        Dropped     DHCP Drops      ACL Drops
 ----     ---------        -------     ----------      ---------
    1            19              6              6              0
—removed for brevity—`

#### Segmentation

Flat networks provide a high degree of flexibility because all resources and users are connected to the same network and can access anything. With the increasing popularity of wireless networks, accessing networks is even easier. The operational and administrative ease make flat networks very popular. You are probably wondering why we are talking about segmentation. Flat networks are not secure! With a flat network, anybody who has access can try to move laterally and may compromise the whole network. In fact, most worms do this.

The most common form of network architecture places critical assets in a central location such as a data center and segments that off while the rest of the network is kept flat. The problem with this approach is that it does not take into consideration the potential to escalate privilege with a lateral attack.

With the increasing numbers and types of devices, such as the IoT and mobile, it is more important than ever before to segment traffic even at the access layer. Such segmentation provides a few key benefits:

- **Network performance:** Segmentation divides broadcast domains and improves network performance. A very large and flat network has degraded performance due to the number of broadcast packets.
- **Reduced area of attack:** When an endpoint is compromised, an attacker or an agent such as a worm attempts to move laterally to compromise other endpoints. Segmentation limits the number of devices that can be targeted.
- **Compliance:** Certain standards, such as the Payment Card Industry Data Security Standard (PCI DSS), require relevant traffic to be kept separate. Segmentation helps achieve that while using the same physical infrastructure.
- **Secure mobility and guest access:** Users are increasingly mobile within and outside offices. On the other hand, visitors and employees require Internet access for their devices. Segmentation allows traffic from guests, non-corporate devices and remote users to be kept separate from normal business traffic.
- **Monitoring and analytics:** The difficulty involved in monitoring and analyzing traffic increases with the size of the network. With segmentation of traffic, monitoring can also be divided into smaller pieces while also adding context to analytics.

Segmentation at the Layer 2 data plane can be achieved with either VLANs or Cisco TrustSec. VLANs work by placing endpoints into logical broadcast domains, while TrustSec works by tagging traffic with scalable (formerly security) group tags (SGTs). TrustSec and SGTs are covered in detail in *Integrated Security Technologies and Solutions, Volume II*. While configuring and managing VLANs is beyond the scope of this book, this section looks at private VLANs (PVLANs).

Typically, devices within a VLAN belong to the same broadcast domain and can communicate directly without going through a Layer 3 device. Each VLAN is assigned a separate IP subnet, and inter-VLAN traffic is routed through a Layer 3 device. In a multitenant environment or segments with few devices, using multiple VLANs and subnets is not always feasible. In such situations, PVLANs can be used to provide segmentation within a VLAN.

PVLANs treat the existing VLAN domain as the *primary VLAN* and then create subdomains called *secondary VLANs*. Traffic within the secondary VLANs is kept separate and has to be routed by a Layer 3 device to reach another secondary VLAN. The traffic between secondary VLANs can be filtered at Layer 3 to provide additional security. PVLANs work by assigning switch ports into three types:

- **Promiscuous ports:** These ports are part of the primary VLAN and can communicate with all other ports in that VLAN, including all ports in secondary VLANs. These ports generally connect to the VLAN gateway to allow traffic to be routed between different secondary VLANs and other primary VLANs.
- **Isolated ports:** These ports are part of a secondary isolated VLAN. PVLANs block all traffic to isolated ports except traffic from promiscuous ports. Similarly, traffic received from an isolated port is forwarded only to promiscuous ports. Isolated ports provide an absolute separation for the hosts connected to them.
- **Community ports:** These ports are part of a secondary community VLAN. Community ports communicate with other ports in the same community VLAN and with promiscuous ports. These interfaces are isolated at Layer 2 from all other interfaces in other communities and from isolated ports within their private VLAN.

