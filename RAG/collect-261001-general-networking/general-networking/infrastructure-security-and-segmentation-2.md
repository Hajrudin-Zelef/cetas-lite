---
id: collect-261001-general-networking/general-networking/infrastructure-security-and-segmentation-2
title: "infrastructure-security-and-segmentation"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/infrastructure-security-and-segmentation.md
source_anchor: ""
source_lines: [61, 151]
sha256: df8b1e5ee305d0807a52411a0c6c6f5606744bce5e7f2bd2bda5c77c839bc8ac
---

# infrastructure-security-and-segmentation

- **Protect:** When the number of secure MAC addresses exceeds the maximum allowed limit on an interface, packets with unknown source addresses are dropped silently. The switch continues to drop frames until a sufficient number of addresses have been removed or the maximum number of allowed addresses has been increased in the configuration. The switch does not generate any alerts or logs in this mode. This mode can be enabled on an interface with the**switchport port-security violation protect** command.
- **Restrict:** Like the Protect mode, this mode also drops frames, but it generates alerts in the form of SNMP traps and syslog, and it also increases the violation counter for that interface. This mode can be enabled on an interface with the**switchport port-security violation restrict** command.
- **Shutdown:** An interface in this mode enters an error-disabled status and shuts down as soon as a port security violation occurs. An interface can be brought out of this state with the**errdisable recovery cause psecure-violation** command in the global configuration mode or with the**shutdown** command followed by the**no shutdown command** on the interface. While this mode is the default, it can be reenabled on an interface with the**switchport port-security violation shutdown** command.

With dynamic secure addressing, it is important to allow the learned addresses to age out of the CAM table. Without proper aging, new devices will not be able to connect on the interface. The aging timer can be configured with the **switchport port-security aging time** *minutes* command.

The aging timer can be configured to be absolute so that it starts as soon as the address is learned, or it can be configured to start when a period of inactivity begins. The type of timer can be configured with the **switchport port-security aging type** {**absolute**| **inactivity**} command.

Example 2-41 shows an interface configured to learn four addresses, out of which one is configured statically and three can be learned dynamically. The interface is also configured to remove the dynamically learned addresses after 10 minutes of inactivity.

#### **Example 2-41** *Port Security Aging*

`SW1(config)#**interface Gi0/9**
SW1(config-if)#**switchport port-security**
SW1(config-if)#**switchport port-security mac-address 1001.1001.1001**
SW1(config-if)#**switchport port-security aging time 10**
SW1(config-if)#**switchport port-security aging type inactivity**`

Port security configuration of an interface can be verified with the **show port-security interface** *interface* command, as shown in Example 2-42.

#### **Example 2-42** *Verifying Port Security Configuration*

`SW1#**show port-security int Gi0/9**
Port Security              : Enabled
Port Status                : Secure-down
Violation Mode             : Shutdown
Aging Time                 : 10 mins
Aging Type                 : Inactivity
SecureStatic Address Aging : Disabled
Maximum MAC Addresses      : 3
Total MAC Addresses        : 1
Configured MAC Addresses   : 1
Sticky MAC Addresses       : 0
Last Source Address:Vlan   : 1001.1001.1001:1
Security Violation Count   : 0`

#### DHCP Snooping

One of the key protocols in modern networks is Dynamic Host Configuration Protocol (DHCP). It provides IP addressing, default gateway, and other information to endpoints as they connect to the network and enables them to communicate. When an endpoint connects, it broadcasts a DHCP request, and any DHCP server in the network can respond to it. That is where the problem with DHCP lies. As you can imagine, the DHCP communication has no security built into it. Any host can claim to be a DHCP server and respond to requests, while any endpoint can forge DHCP requests to get an address assigned. This situation can be exploited to carry two attacks:

- **DHCP spoofing MITM:** An attacker can set up a rogue DHCP server to respond to DHCP requests. This DHCP server responds to requests with its own IP address as the default gateway. This causes the victims to send their traffic through the rogue gateway, where it can be read, stored, and modified before being sent to the real gateway, resulting in a MITM attack.
- **DHCP starvation DoS:** Because a DHCP server responds to any client request, it is trivial to forge thousands of DHCP requests with different MAC addresses and cause the server to exhaust its address pool. When the pool is exhausted, legitimate clients do not receive IP addresses, resulting in a DoS situation. This attack is often a precursor to a DHCP spoofing MITM attack, just described. If the legitimate DHCP pool is exhausted, a larger number of endpoints receive IP addresses from the rogue DHCP server.

While a DHCP starvation attack can be prevented with port security, as described earlier, mitigating DHCP spoofing attacks requires the use of DHCP snooping on Cisco switches.

*DHCP snooping* is a security feature that acts as a filtering mechanism between DHCP servers and clients. It works by defining trusted and untrusted interfaces. Trusted interfaces are those from which DHCP server messages can be expected, while all other interfaces are untrusted. This prevents rogue DHCP servers from responding to requests.

In addition to filtering rogue DHCP server packets, DHCP snooping also builds a database of clients connected to untrusted interfaces, along with the IP address, MAC address, and VLAN ID of each one. When a DHCP packet is received from an untrusted interface, its content is validated against the database. This prevents spoofed DHCP requests from being sent out.

Finally, DHCP snooping has a rate-limiting function that limits the number of DHCP packets allowed on an untrusted interface. This helps prevent starvation attacks against the DHCP server.

DHCP snooping configuration can be broken down into four steps:

- **Step 1.** Enable DHCP snooping on VLANs that require it with the**ip dhcp snooping vlan***vlan-id* command. Multiple VLANs can be specified, separated by commas, or a range of VLANs can be specified by using a dash.
- **Step 2.** Configure trusted interfaces from which DHCP server messages are expected with the**ip dhcp snooping trust** command. All interfaces are untrusted by default.
- **Step 3.** Configure DHCP rate limiting on untrusted interfaces with the**ip dhcp limit rate***limit* command.
- **Step 4.** Enable DHCP snooping globally with the**ip dhcp snooping** command.

Example 2-43 shows DHCP snooping configuration on a switch. In this example, interface GigabitEthernet0/10 is the trunk interface toward the DHCP server, and interfaces GigabitEthernet0/1 to 9 are access ports for endpoints. Interface Gi0/10 is configured as trusted, while others have rate limits applied to them.

#### **Example 2-43** *Configuring DHCP Snooping*

`SW1(config)#**ip dhcp snooping vlan 1,4**
SW1(config)#**interface Gi0/10**
SW1(config-if)#**ip dhcp snooping trust**
SW1(config-if)#**exit**
SW1(config)#**interface range Gi0/1-9**
SW1(config-if-range)#**ip dhcp snooping limit rate 5**
SW1(config-if)#**exit**
SW1(config)#**ip dhcp snooping**`

DHCP snooping can be verified with the **show ip dhcp snooping** command, as shown in Example 2-44.

#### **Example 2-44** *Verifying DHCP Snooping*

`SW1#**show ip dhcp snooping**
Switch DHCP snooping is enabled
DHCP snooping is configured on following VLANs:
1,4
DHCP snooping is operational on following VLANs:
1
— removed for brevity -
DHCP snooping trust/rate is configured on the following Interfaces:
Interface                  Trusted    Allow option    Rate limit (pps)
----------------------     --------   ------------    ----------------
GigabitEthernet0/1         no         no              5
— remove for brevity —
GigabitEthernet0/10        yes        yes             unlimited`

#### The ARP Table and Dynamic ARP Inspection (DAI)

