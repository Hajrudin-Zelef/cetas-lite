---
id: collect-261001-cisco/cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-10-configuration-g-69fa402c-1
title: "c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-10-configuration-g-69fa402c"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["throughput"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-10-configuration-g-69fa402c.md
source_anchor: ""
source_lines: [1, 79]
sha256: b926e1e5e34e1f02b45b201ab0c0a9d69a3eb4b163ca97eefc32eca3e943ec7c
---

# c-en-us-td-docs-switches-lan-catalyst9300-software-release-16-10-configuration-g-69fa402c

The documentation set for this product strives to use bias-free language. For the purposes of this documentation set, bias-free is defined as language that does not imply discrimination based on age, disability, gender, racial identity, ethnic identity, sexual orientation, socioeconomic status, and intersectionality. Exceptions may be present in the documentation due to language that is hardcoded in the user interfaces of the product software, language used based on RFP documentation, or language that is used by a referenced third-party product. Learn more about how Cisco is using Inclusive Language.
Your software release may not support all the features documented in this module. For the latest caveats and feature information,
see Bug Search Tool and the release notes for your platform and software release. To find information about the features documented
in this module, and to see a list of the releases in which each feature is supported, see the feature information table at
the end of this module.
Use Cisco Feature Navigator to find information about platform support and Cisco software image support. To access Cisco Feature
Navigator, go to https://cfnng.cisco.com/. An account on Cisco.com is not required.
Network Address Translation (NAT)
Network Address Translation (NAT) is designed for IP address conservation. It enables private IP networks that use unregistered
IP addresses to connect to the Internet. NAT operates on a device, usually connecting two networks together, and translates
the private (not globally unique) addresses in the internal network into global routable addresses, before packets are forwarded
onto another network.
NAT can be configured to advertise only one address for the entire network to the outside world. This ability provides additional
security by effectively hiding the entire internal network behind that one address. NAT offers the dual functions of security
and address conservation and is typically implemented in remote-access environments.
NAT is also used at the enterprise edge to allow internal users access to the Internet and to allow Internet access to internal
devices such as mail servers.
Cisco Catalyst 9300 Series Switch supports Stacking and NAT is supported on a stack set-up.
Benefits of Configuring NAT
Resolves the problem of IP depletion.
NAT allows organizations to resolve the problem of IP address depletion when they have existing networks and need to access
the Internet. Sites that do not yet possess Network Information Center (NIC)-registered IP addresses must acquire IP addresses,
and if more than 254 clients are present or are planned, the scarcity of Class B addresses becomes a serious issue. NAT addresses
these issues by mapping thousands of hidden internal addresses to a range of easy-to-get Class C addresses.
Provides a layer of security by preventing the client IP address from being exposed to the outside network.
Sites that already have registered IP addresses for clients on an internal network may want to hide those addresses from the
Internet so that hackers cannot directly attack clients. With client addresses hidden, a degree of security is established.
NAT gives LAN administrators complete freedom to expand Class A addressing, which is drawn from the reserve pool of the Internet
Assigned Numbers Authority. The expansion of Class A addresses occurs within the organization without a concern for addressing
changes at the LAN or the Internet interface.
Cisco software can selectively or dynamically perform NAT. This flexibility allows network administrator to use RFC 1918 addresses
or registered addresses.
NAT is designed for use on a variety of devices for IP address simplification and conservation. In addition, NAT allows the
selection of internal hosts that are available for translation.
A significant advantage of NAT is that it can be configured without requiring any changes to devices other than to those few
devices on which NAT will be configured.
How NAT Works
A device that is configured with NAT will have at least one interface to the inside network and one to the outside network.
In a typical environment, NAT is configured at the exit device between a stub domain and the backbone. When a packet leaves
the domain, NAT translates the locally significant source address into a globally unique address. When a packet enters the
domain, NAT translates the globally unique destination address into a local address. Multiple inside networks could be connected
to the device and similarly there might exist multiple exit points from the device towards outside networks. If NAT cannot
allocate an address because it has run out of addresses, it drops the packet and sends an Internet Control Message Protocol
(ICMP) host unreachable packet to the destination.
Translation and forwarding are performed in the hardware switching plane, thereby improving the overall throughput performance.
For more details on performance, refer the section on Performance and Scale numbers.
Uses of NAT
NAT can be used for the following scenarios:
To connect to the Internet when only a few of your hosts have globally unique IP address.
NAT is configured on a device at the border of a stub domain (referred to as the inside network) and a public network such
as the Internet (referred to as the outside network). NAT translates internal local addresses to globally unique IP addresses
before sending packets to the outside network. As a solution to the connectivity problem, NAT is practical only when relatively
few hosts in a stub domain communicate outside of the domain at the same time. When this is the case, only a small subset
of the IP addresses in the domain must be translated into globally unique IP addresses when outside communication is necessary,
and these addresses can be reused
Renumbering:
Instead of changing the internal addresses, which can be a considerable amount of work, you can translate them by using NAT.
NAT Inside and Outside Addresses
The term inside in a NAT context refers to networks owned by an organization that must be translated. When NAT is configured, hosts within
this network will have addresses in one space (known as the local address space) that will appear to those outside the network
as being in another space (known as the global address space).
Similarly, the term outside refers to those networks to which the stub network connects, and which are generally not under
the control of an organization. Hosts in outside networks can also be subject to translation, and can thus have local and
global addresses.
NAT uses the following definitions:
Inside local address—an IP address that is assigned to a host on the inside network. The address is probably not a routable
IP address assigned by NIC or service provider.
Inside global address—a global routable IP address (assigned by the NIC or service provider) that represents one or more inside
local IP addresses to the outside world.
Outside local address—the IP address of an outside host as it appears to the inside network. Not necessarily a routable IP
address, it is allocated from the address space that is routable on the inside.
Outside global address—the IP address assigned to a host on the outside network by the owner of the host. The address is allocated
from a globally routable address or network space.
Inside Source Address Translation—translates an inside local address to inside global address.
Outside Source Address Translation—translates the outside global address to outside local address.
Static Port Translation—translates the IP address and port number of an inside/outside local address to the IP address and
port number of the corresponding inside/outside global address.
Static Translation of a given subnet—translates a specified range of subnets of an inside/outside local address to the corresponding
inside/outside global address.
