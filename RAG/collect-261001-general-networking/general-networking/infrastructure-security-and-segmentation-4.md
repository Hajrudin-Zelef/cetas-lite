---
id: collect-261001-general-networking/general-networking/infrastructure-security-and-segmentation-4
title: "infrastructure-security-and-segmentation"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["exploit", "memory"]
source: docs/RAG/collect-261001-general-networking/infrastructure-security-and-segmentation.md
source_anchor: ""
source_lines: [217, 291]
sha256: f098f4e038b25899cdd728654bd92b2f8b3edfc627587651f6e03ebf8e5c0f32
---

# infrastructure-security-and-segmentation

A PVLAN can work across multiple switches, and VTPv3 can be used to carry PVLAN information across the domain. Configuring PVLAN can be broken down into four steps:

- **Step 1. Defining the secondary VLANs:** Each secondary VLAN should be configured as required. The**private-vlan** {**community** |**isolated** } command is used in VLAN configuration mode for this.
- **Step 2. Defining the primary VLAN:** The primary VLAN is configured for the PVLAN, and the secondary VLANs are associated with it. A VLAN can be declared primary with the**private-vlan primary** command in VLAN configuration mode. Secondary VLAN associations are also declared in that mode, using the**private-vlan association***vlan-ID* command. Multiple VLANs can be specified as a comma-separated list.
- **Step 3. Configuring a promiscuous port:** A switch interface can be configured as a promiscuous port with the**switchport mode private-vlan promiscuous** command. The PVLANs should then be mapped to the interface with the**switchport private-vlan mapping***primary-vlan-ID* ,*secondary-vlan-list* command, where*secondary-vlan-list* is a comma-separated list of all secondary VLANs of that PVLAN.
- **Step 4. Configuring member ports:** The PVLAN can be enabled on each participating interface with the**switchport mode private-vlan host** command. The primary and secondary VLANs can be mapped to the interface with the**switchport private-vlan host-association***primary-vlan-ID secondary- vlan-ID* command.

Example 2-47 shows the configuration of private VLANs with interface Gi0/10 configured as the promiscuous port in primary VLAN 10, interface Gi0/11 configured as an isolated port in secondary VLAN 20, and interface Gi0/12 configured as a community port in secondary VLAN 30.

#### **Example 2-47** *Configuring Private VLANs*

`SW1(config)#**vlan 20**
SW1(config-vlan)#**private-vlan isolated**
SW1(config-vlan)#**exit**
SW1(config)#**vlan 30**
SW1(config-vlan)#**private-vlan community**
SW1(config-vlan)#**exit**
SW1(config)#**vlan 10**
SW1(config-vlan)#**private-vlan primary**
SW1(config-vlan)#**private-vlan association 20,30**
SW1(config-vlan)#**exit**
SW1(config)#**interface Gi0/10**
SW1(config-if)#**switchport mode private-vlan promiscuous**
SW1(config-if)#**switchport private-vlan mapping 10 20,30**
SW1(config-if)#**exit**
SW1(config)#**interface Gi0/11**
SW1(config-if)#**switchport mode private-vlan host**
SW1(config-if)#**switchport private-vlan host-association 10 20**
SW1(config-if)#**exit**
SW1(config)#**interface Gi0/12**
SW1(config-if)#**switchport mode private-vlan host**
SW1(config-if)#**switchport private-vlan host-association 10 30**`

#### Attacks Against Segmentation

VLANs and PVLANs are both subject to some attacks. The primary motivation behind the attacks is to send traffic outside the segment that the attacker belongs to without going through a Layer 3 device and any filtering configured there. Whereas VLANs are subject to *VLAN hopping attacks*, PVLANs are suspect to the unimaginatively named *private VLAN attacks*.

In a VLAN hopping attack, two methods can be used to send traffic outside the VLAN without going through a router:

- **Basic VLAN hopping:** To execute a basic VLAN hopping attack, the attacker establishes a trunk link with the switch and is then able to tag frames with any VLAN. As mentioned earlier, Cisco switch interfaces have DTP enabled by default, which allows trunk negotiation on any interface if not disabled.
- **Double tagging:** In this type of attack, the attacker sends a frame with two 802.1q tags. The first tag specifies a VLAN that the attacker’s host actually belongs to, and the second frame specifies a VLAN of the destination host. This attack attempts to exploit the fact that most trunks have their native VLANs set to the same one as the hosts, and they allow frames in native VLANs to be sent without any VLAN tags. This results in the first tag being stripped at the source switch and being delivered across a trunk to the destination switch, where the second tag is read and the frame is delivered to the destination.

Basic VLAN hopping attacks can easily be mitigated by disabling DTP negotiation on interfaces and by configuring non-trunk interfaces in access mode.

Double tagging attacks can be mitigated by either configuring the native VLAN on trunk links to be an unused VLAN or by forcing trunks to tag frames in native VLANs also. The native VLAN of a trunk interface can be changed with the **switchport trunk native vlan** *vlan-ID* command, and native VLAN tagging can be enabled with the **vlan dot1q tag native** global configuration command.

In a PVLAN attack, the attacker attempts to send a packet to a host in another isolated or community VLAN. This is done by sending a crafted IP packet with the following:

- Real source MAC and IP addresses
- A real destination IP address
- The destination MAC address of the gateway router instead of the destination host

Because the destination MAC address belongs to the gateway router connected to a promiscuous port, the switch delivers it. Because the router only looks at the destination IP address, it routes the packet to the destination. This results in the packet being delivered outside the PVLAN’s secondary VLAN.

PVLAN attacks can be mitigated by applying an ACL on the router interface, connected to the promiscuous port, to drop packets that originate from and are destined to the same IP subnet.

#### Traffic Filtering at Layer 2

While filtering and access lists are generally associated with routers and firewalls, they can also be applied at Layer 2 interfaces and to VLANs to provide granular security. The following are some of the benefits of using access lists at Layer 2:

- **Contextual filtering:** Filtering at Layer 3 is generally based on subnets of the source traffic. A subnet can have multiple types of devices, such as IP phones, workstations, printers, and such. This context of the device is lost when filtering is done at that level. On the other hand, the context of the device is known at the switch interface it connects to, and filtering can be designed based on that. For example, an IP phone only needs to communicate with a certain set of services, so filtering can be applied to drop traffic destined to any other service.
- **Containing lateral attacks:** In most cases, devices in a subnet do not need to communicate with each other. Most services are centrally located, and there is little need for endpoints in a subnet to communicate with each other. For example, in many cases workstation-to-workstation or printer-to-printer communication is not required. Restricting intra-subnet communication with Layer 2 ACLs reduces the chance for lateral movement from a compromised host.
- **Reduced load on Layer 3:** When filtering is applied right at source, the rest of the network has to do less. This improves network performance and simplifies configuration

Filtering on a Cisco switch can be configured with port access control lists (PACLs) or VLAN access control lists (VACLs).

PACLs are standard, extended, or named IP ACLs, and named MAC address ACLs applied to a switch interface. The syntax for creating PACLs is the same as the syntax for creating ACLs on any Cisco IOS router.

When the PACL is applied, it filters incoming traffic on an interface. A few restrictions apply to PACL:

- Log, reflect, and evaluate keywords cannot be used.
- Physical and logical link protocols such as CDP, STP, DTP, and VTP cannot be filtered with a PACL.
- Ingress traffic is evaluated against PACLs before any other ACLs, such as a VACL.
- PACLs take up Ternary Content-Addressable Memory (TCAM) space and should be kept as small as possible. Generally, 20 to 30 Access Control Entries (ACEs) per PACL is an acceptable value.

