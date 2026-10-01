---
id: collect-261001-cisco/cisco/c-en-us-td-docs-ios-xml-ios-sec-conn-vpnav-configuration-xe-16-12-sec-vpn-availa-86c4f878-1
title: "c-en-us-td-docs-ios-xml-ios-sec-conn-vpnav-configuration-xe-16-12-sec-vpn-availa-86c4f878"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-ios-xml-ios-sec-conn-vpnav-configuration-xe-16-12-sec-vpn-availa-86c4f878.md
source_anchor: ""
source_lines: [1, 111]
sha256: a67a7ba5e724c9c42ed2924b40179e94e843670b8f70c9a78c6d43f309b9ddf8
---

# c-en-us-td-docs-ios-xml-ios-sec-conn-vpnav-configuration-xe-16-12-sec-vpn-availa-86c4f878

The documentation set for this product strives to use bias-free language. For the purposes of this documentation set, bias-free is defined as language that does not imply discrimination based on age, disability, gender, racial identity, ethnic identity, sexual orientation, socioeconomic status, and intersectionality. Exceptions may be present in the documentation due to language that is hardcoded in the user interfaces of the product software, language used based on RFP documentation, or language that is used by a referenced third-party product. Learn more about how Cisco is using Inclusive Language.
The IPsec VPN High Availability Enhancements feature: Reverse Route Injection (RRI) and Hot Standby Router Protocol (HSRP)
with IPsec. When used together, these two features provide you with a simplified network design for VPNs and reduced configuration
complexity on remote peers when defining gateway lists.
Note
Security threats, as well as the cryptographic technologies to help protect against them, are constantly changing. For more
information about the latest Cisco cryptographic recommendations, see the
Next Generation Encryption (NGE) white paper.
Finding Feature Information
Your software release may not support all the features documented in this module. For the latest caveats and feature information,
see
Bug Search Tool and the release notes for your platform and software release. To find information about the features documented in this module,
and to see a list of the releases in which each feature is supported, see the feature information table at the end of this
module.
Use Cisco Feature Navigator to find information about platform support and Cisco software image support. To access Cisco Feature
Navigator, go to
www.cisco.com/go/cfn. An account on Cisco.com is not required.
Information About IPsec VPN High Availability Enhancements
Reverse Route Injection
Reverse Route Injection (RRI) simplifies network design for Virtual Private Networks (VPNs) in which there is a requirement
for redundancy or load balancing. RRI works with both dynamic and static crypto maps.
RRI provides the following benefits:
Enables routing of IPsec traffic to a specific VPN headend device in environments that have multiple (redundant) VPN headend
devices.
Ensures predictable failover time of remote sessions between headend devices when using IKE keepalives, especially in environments
in which remote device route flapping is common (not taking into consideration the effects of route convergence, which may
vary depending on the routing protocol used and the size of the network).
Eliminates the need for the administration of static routes on upstream devices, as routes are dynamically learned by these
devices.
In the dynamic case, as remote peers establish IPsec security associations (SAs) with an RRI-enabled router, a static route
is created for each subnet or host protected by that remote peer. For static crypto maps, a static route is created for each
destination of an extended access list rule. When RRI is used on a static crypto map with an access control list (ACL), routes
will always exist, even without the negotiation of IPsec SAs.
Note
The use of any keyword in ACLs with RRI is not supported.
When routes are created, they are injected into any dynamic routing protocol and distributed to surrounding devices. This
traffic flows, requiring IPsec to be directed to the appropriate RRI router for transport across the correct SAs to avoid
IPsec policy mismatches and possible packet loss.
The figure below shows an RRI configuration functionality topology. Remote A is being serviced by Router A and Remote B connected
to Router B, providing load balancing across VPN gateways at the central site. RRI on the central site devices ensures that
the other router on the inside of the network can automatically make the correct forwarding decision. RRI also eliminates
the need to administer static routes on the inside router.
Hot Standby Router Protocol and IPsec
Hot Standby Router Protocol (HSRP) provides high network availability by routing IP traffic from hosts on Ethernet networks
without relying on the availability of any single router. HSRP is particularly useful for hosts that do not support a router
discovery protocol, such as ICMP Router Discovery Protocol (IRDP) and do not have the functionality to switch to a new router
when their selected router reloads or loses power. Without this functionality, a router that loses its default gateway because
of a router failure cannot communicate with the network.
HSRP is configurable on LAN interfaces using standby command-line interface (CLI) commands. You can to use the standby IP
address from an interface as the local IPsec identity or local tunnel endpoint.
By using the standby IP address as the tunnel endpoint, failover can be applied to VPN routers by using HSRP. Remote VPN
gateways connect to the local VPN router via the standby address that belongs to the active device in the HSRP group. In the
event of failover, the standby device takes over ownership of the standby IP address and begins to service remote VPN gateways.
Failover can be applied to VPN routers through the use of HSRP. Remote VPN gateways connect to the local VPN router through
the standby address that belongs to the active device in the HSRP group. This functionality reduces configuration complexity
on remote peers with respect to defining gateway lists, because only the HSRP standby address needs to be defined.
The figure below shows the enhanced HSRP functionality topology. Traffic is serviced by the active Router P, which is the
active device in the standby group. In the event of failover, traffic is diverted to Router S, which is the original standby
device. Router S assumes the role of the new active router and takes ownership of the standby IP address.
Note
In case of a failover, HSRP does not facilitate IPsec state information transference between VPN routers. This means that
without this state transference, SAs to remotes will be deleted, requiring Internet Key Exchange (IKE) and IPsec SAs to be
reestablished. To make IPsec failover more efficient, it is recommended that IKE keepalives be enabled on all routers.
How to Configure IPsec VPN High Availability Enhancements
Configuring Reverse Route Injection on a Dynamic Crypto Map
Dynamic crypto map entries, like regular static crypto map entries, are grouped into sets. A set is a group of dynamic crypto
map entries all with the same dynamic map name, but each with a different dynamic sequence number. Each member of the set
may be configured for RRI.
To create a dynamic crypto map entry and enable RRI, perform the steps in this section.
SUMMARY STEPS
enable
configureterminal
cryptodynamic-mapmap-nameseq-num
settransform-set
reverse-route
DETAILED STEPS
Command or Action
Purpose
Step 1
enable
Example:
Router> enable
Enables privileged EXEC mode.
Enter your password if prompted.
Step 2
configureterminal
Example:
Router# configure terminal
Enters global configuration mode.
Step 3
cryptodynamic-mapmap-nameseq-num
Example:
Router(config)# crypto dynamic-map mymap
Creates a dynamic crypto map entry and enters crypto map configuration mode.
Step 4
settransform-set
Example:
Router(config-crypto-m)#set transform-set
Specifies which transform sets are allowed for the crypto map entry. Lists multiple transform sets in order of priority (highest
priority first).
This entry is the only configuration statement required in dynamic crypto map entries.
Step 5
reverse-route
Example:
Router(config-crypto-m)#reverse-route
Creates source proxy information.
Configuring Reverse Route Injection on a Static Crypto Map
Before configuring RRI on a static crypto map, note that:
Routes are not created based on access list 102, as reverse-route is not enabled on mymap 2. RRI is not enabled by default
and is not displayed in the router configuration.
Enable a routing protocol to distribute the VPN routes to upstream devices.
