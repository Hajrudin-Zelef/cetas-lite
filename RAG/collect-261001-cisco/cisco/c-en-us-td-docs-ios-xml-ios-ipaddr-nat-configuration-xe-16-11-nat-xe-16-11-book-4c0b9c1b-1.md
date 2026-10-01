---
id: collect-261001-cisco/cisco/c-en-us-td-docs-ios-xml-ios-ipaddr-nat-configuration-xe-16-11-nat-xe-16-11-book-4c0b9c1b-1
title: "c-en-us-td-docs-ios-xml-ios-ipaddr-nat-configuration-xe-16-11-nat-xe-16-11-book--4c0b9c1b"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-ios-xml-ios-ipaddr-nat-configuration-xe-16-11-nat-xe-16-11-book--4c0b9c1b.md
source_anchor: ""
source_lines: [1, 96]
sha256: f17492dd8618b9ab5071b5e1743737cb9d495eadc824827c4c61dab35f8ed383
---

# c-en-us-td-docs-ios-xml-ios-ipaddr-nat-configuration-xe-16-11-nat-xe-16-11-book--4c0b9c1b

IP Addressing: NAT Configuration Guide, Cisco IOS XE Gibraltar 16.11.x
Bias-Free Language
Bias-Free Language
The documentation set for this product strives to use bias-free language. For the purposes of this documentation set, bias-free is defined as language that does not imply discrimination based on age, disability, gender, racial identity, ethnic identity, sexual orientation, socioeconomic status, and intersectionality. Exceptions may be present in the documentation due to language that is hardcoded in the user interfaces of the product software, language used based on RFP documentation, or language that is used by a referenced third-party product. Learn more about how Cisco is using Inclusive Language.
The ability of Network Address Translation (NAT) to consistently represent a local IP address as a single global IP address
is termed paired address pooling. Paired address pooling is supported only on Port Address Translation (PAT).
Prior to the introduction of the Paired-Address-Pooling Support feature, if you have a PAT configuration, and you need a
new global address or port, the next available address in the IP address pool is allocated. There was no mechanism to ensure
that a local address is consistently mapped to a single global address. The Paired-Address-Pooling Support feature provides
the ability to consistently map a local address to a global address.
Starting from IOS XE Polaris 16.8 release, you can specify an NAT pool for which PAP support is to be activated. This feature
is helpful when you have to apply PAP support to a specific dynamic NAT traffic stream.
Restrictions for Paired-Address-Pooling Support in NAT
Paired address pooling uses more memory, and the scaling of translations is much lower than standard Network Address Translation
(NAT) configuration due to the following reasons:
Use of a new data structure that tracks each local address.
Use of the paired-address-pooling limit. When the number of users on a global address reaches the configured limit, the next
global address is used for paired address pooling. The paired-address-pooling limit uses more memory and requires more global
addresses in the address pool than standard NAT.
Two IP address pools with same IP addresses in two different mapping is not supported.
The following example shows two non-VRF mappings. The addresses used in these two pools mappings should not overlap.
ip nat pool natpool1 83.0.0.56 83.0.0.56 prefix-length 24
ip nat pool natpool2 83.0.0.56 83.0.0.56 prefix-length 24
ip nat inside source list acl2 pool natpool2 overload
ip nat inside source list acl1 pool natpool1 overload
This following example is a combination of non-VRF and VRF-to-global mappings. In this example as well, sharing IP addresses
in pools are not supported.
ip nat pool natpool1 82.0.0.15 82.0.0.15 prefix-length 24
ip nat pool natpool2 82.0.0.15 82.0.0.15 prefix-length 24
ip nat inside source list acl2 pool natpool2 overload //non-vrf mapping//
ip nat inside source list acl1 pool natpool1 vrf vrf1 overload //vrf mapping//
The only case where same pools can be used in two different mapping is for the match-in-vrf mappings.
Information About Paired-Address-Pooling Support in NAT
An IP address pool is a group of IP addresses. You create an IP address pool by assigning a range of IP addresses and a
name to it. You allocate or assign addresses in the pool to users.
The ability of Network Address Translation (NAT) to consistently represent a local IP address as a single global IP address
is termed paired address pooling. A local address is any address that appears on the inside of a network, and a global address
is any address that appears on the outside of the network. You can configure paired address pooling only for Port Address
Translation (PAT) because dynamic and static NAT configurations are paired configurations by default. PAT, also called overloading,
is a form of dynamic NAT that maps multiple, unregistered IP addresses to a single, registered IP address (many-to-one) by
using different ports. Paired address pooling is supported in both classic (default) and carrier-grade NAT (CGN) mode.
In a paired-address-pooling configuration, a local address is consistently represented as a single global address. For example,
if User A is paired with the global address G1, that pairing will last as long as there are active sessions for User A.
If there are no active sessions, the pairing is removed. When User A has active sessions again, the user may be paired with
a different global address.
If a local address initiates new sessions, and resources (ports) are insufficient for its global address, packets are dropped.
When the number of users on a global address reaches the configured limit, the next global address is used for paired address
pooling. When a user who is associated with a global address through paired address pooling is unable to get a port number,
then the packet is dropped, the NAT drop code is incremented, and Internet Control Message Protocol (ICMP) messages are not
sent.
Paired-address-pooling uses the fill-it-up method for address selection. The fill-it-up method fits (adds) the maximum
possible users into a single global address before going to the next global address.
If you change the Network Address Translation (NAT) configuration mode to paired-address-pooling configuration mode and vice
versa, all existing NAT sessions are removed.
To configure NAT paired-address-pooling mode, use the ip nat settings pap command. To remove it, use the no ip nat settings pap command.
After you configure paired-address-pooling mode, all pool-overload mappings will act in the paired-address-pooling manner.
Based on your NAT configuration, you can use NAT static or dynamic rules.
SUMMARY STEPS
enable
configure terminal
ip nat settings pap [limit {1000 | 120 | 250 | 30 | 500 | 60}]
ip nat pool name start-ip end-ip {netmask netmask | prefix-length prefix-length}
Configuring Paired-Address-Pooling Support For a NAT Pool
Note
If you change the Network Address Translation (NAT) configuration mode to paired-address-pooling configuration mode and vice
versa, all existing NAT sessions are removed.
To configure NAT paired-address-pooling mode, use the ip nat settings pap command. To remove it, use the no ip nat settings pap command.
After you configure paired-address-pooling mode, all pool-overload mappings will act in the paired-address-pooling manner.
Based on your NAT configuration, you can use NAT static or dynamic rules.
SUMMARY STEPS
enable
configure terminal
ip nat settings pap [limit {1000 | 120 | 250 | 30 | 500 | 60}]
ip nat pool name start-ip end-ip {netmask netmask | prefix-length prefix-length}
Example: Configuring Paired Address Pooling Support in NAT
The following example shows how to configure paired address pooling along with Network Address Translation (NAT) rules. This
example shows a dynamic NAT configuration with access lists and address pools. Based on your NAT configuration, you can configure
static or dynamic NAT rules.
Device# configure terminal
Device(config)# ip nat settings pap
Device(config)# ip nat pool net-208 192.168.202.129 192.168.202.158 netmask 255.255.255.240
Device(config)# access-list 1 permit 192.168.34.0 0.0.0.255
Device(config)# ip nat inside source list 1 pool net-208 overload
Device(config)# interface gigabitethernet 0/0/1
Device(config-if)# ip address 10.114.11.39 255.255.255.0
Device(config-if)# ip nat inside
Device(config-if)# exit
Device(config)# interface gigabitethernet 0/1/2
Device(config-if)# ip address 172.16.232.182 255.255.255.240
Device(config-if)# ip nat outside
Device(config-if)# end
Additional References for Paired-Address-Pooling Support in NAT
The Cisco Support website provides extensive online resources,
including documentation and tools for troubleshooting and
resolving technical issues with Cisco products and technologies.
To receive security and technical information about your
