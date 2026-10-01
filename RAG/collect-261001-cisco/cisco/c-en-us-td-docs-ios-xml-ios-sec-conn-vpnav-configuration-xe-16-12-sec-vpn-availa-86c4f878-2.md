---
id: collect-261001-cisco/cisco/c-en-us-td-docs-ios-xml-ios-sec-conn-vpnav-configuration-xe-16-12-sec-vpn-availa-86c4f878-2
title: "c-en-us-td-docs-ios-xml-ios-sec-conn-vpnav-configuration-xe-16-12-sec-vpn-availa-86c4f878"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-en-us-td-docs-ios-xml-ios-sec-conn-vpnav-configuration-xe-16-12-sec-vpn-availa-86c4f878.md
source_anchor: ""
source_lines: [112, 284]
sha256: 19665a63bc11fd7ad9aadeb1ae2af51a5197ef759a0ea45e9ffdb7bc1c2d2f81
---

# c-en-us-td-docs-ios-xml-ios-sec-conn-vpnav-configuration-xe-16-12-sec-vpn-availa-86c4f878

If Cisco Express Forwarding (CEF) is run on a VPN router configured for RRI, adjacencies need to be formed for each RRI injected
network through the next hop device. As the next hop is not explicitly defined in the routing table for these routes, proxy-ARP
should be enabled on the next hop router, which allows the CEF adjacency to be formed using the Layer 2 addresses of that
device. In cases where there are many RRI injected routes, adjacency tables may become quite large, as an entry is created
for each device from each of the subnets represented by the RRI route.
To add RRI to a static crypto map set, perform the steps in this section.
SUMMARY STEPS
enable
configureterminal
cryptomapmap-nameseq-numipsec-isakmp
setpeerip-address
reverse-route
matchaddress
set transform-set
transform-set-name
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
cryptomapmap-nameseq-numipsec-isakmp
Example:
Router(config)#crypto map mymap 3 ipsec-isakmp
Adds a dynamic crypto map set to a static crypto map set and enters interface configuration mode.
Step 4
setpeerip-address
Example:
Router(config-if)#set peer 209.165.200.248
Specifies an IPsec peer IP address in a crypto map entry.
Step 5
reverse-route
Example:
Router (config-if)#reverse-route
Creates dynamic static routes based on crypto access control lists (ACLs).
Step 6
matchaddress
Example:
Router(config-if)# match address
Specifies an extended access list for a crypto map entry.
Step 7
set transform-set
transform-set-name
Example:
Router (config-if)# set transform-set my_t_set1
Specifies which transform sets are allowed for the crypto map entry. List multiple transform sets in order of priority (highest
priority first).
Configuring HSRP with IPsec
When configuring HSRP with IPsec, the following conditions may apply:
When HSRP is applied to a crypto map on an interface, the crypto map must be reapplied if the standby IP address or the standby
name is changed on that interface.
If HSRP is applied to a crypto map on an interface, and you delete the standby IP address or the standby name from that interface,
the crypto tunnel endpoint is reinitialized to the actual IP address of that interface.
If you add the standby IP address and the standby name to an interface with the requirement IPsec failover, the crypto map
must be reapplied with the appropriate redundancy information.
Standby priorities should be equal on active and standby routers. If they are not, the higher priority router takes over
as the active router. If the old active router comes back up and immediately assumes the active role before having time to
report itself, standby and sync connections will be dropped.
The IP addresses on the HSRP-tracked interfaces on the standby and active routers should both be either lower or higher on
one router than the other. In the case of equal priorities (an HA requirement), HSRP will assign the active state-based IP
address. If an addressing scheme exists so that the public IP address of router A is lower than the public IP address of router
B, but the opposite is true for their private interfaces, an active/standby-standby/active split condition could exist, which
will break connectivity.
Note
To configure HSRP without IPsec, refer to the “Configuring IP Services“ module in the
IP Application Services Configuration Guide.
To apply a crypto map set to an interface, perform the steps in this section.
SUMMARY STEPS
enable
configureterminal
interfacetypeslot/port
standbynamegroup-name
standbyipip-address
cryptomapmap-nameredundancy[standby-name]
DETAILED STEPS
Command or Action
Purpose
Step 1
enable
Example:
Router>enable
Enables privileged EXEC mode.
Enter your password if prompted.
Step 2
configureterminal
Example:
Router#configure terminal
Enters global configuration mode.
Step 3
interfacetypeslot/port
Example:
Router(config)#interface GigabitEthernet 0/0
Specifies an interface and enters interface configuration mode.
Step 4
standbynamegroup-name
Example:
Router(config-if)#standby name mygroup
Specifies the standby group name.
Step 5
standbyipip-address
Example:
Router(config-if)#standby ip 209.165.200.249
Specifies the IP address of the standby groups
This command is required for one device in the group.
Step 6
cryptomapmap-nameredundancy[standby-name]
Example:
Router (config-if)#crypto map mymap redundancy
Specifies the IP redundancy address as the tunnel endpoint for IPsec.
Configuration Examples for IPsec VPN High Availability Enhancements
Example: Configuring Reverse Route Injection on a Dynamic Crypto Map
In the following example, using the
reverse-route command in the definition of the dynamic crypto map template ensures that routes are created for any remote proxies (subnets
or hosts), protected by the connecting remote IPsec peers.
crypto dynamic mydynmap 1
set transform-set my-transform-set
reverse-route
This template is then associated with a “parent” crypto map statement and then applied to an interface.
Example: Configuring Reverse Route Injection on a Static Crypto Map
RRI is a good solution for topologies that require encrypted traffic to be diverted to a VPN router and all other traffic
to a different router. In these scenarios, RRI eliminates the need to manually define static routes on devices.
RRI is not required if a single VPN router is used, and all traffic passes through the VPN router during its path in to and
out of the network.
If you choose to manually define static routes on the VPN router for remote proxies and have these routes permanently installed
in the routing table, RRI should not be enabled on the crypto map instance that covers the same remote proxies. In this case,
there is no possibility of user-defined static routes being removed by RRI.
Routing convergence can affect the success of a failover based on the routing protocol used to advertise routes (link state
versus periodic update). We recommend that a link state routing protocol such as OSPF be used to help speed convergence time
by ensuring that routing updates are sent as soon as a change in routing state is detected.
In the following example, RRI is enabled for mymap 1, but not for mymap 2. Upon the application of the crypto map to the
interface, a route is created based on access-list 101 analogous to the following:
IP route 172.17.11.0 255.255.255.0 FastEthernet 0/0
crypto map mymap 1 ipsec-isakmp
set peer 172.17.11.1
reverse-route
set transform-set my-transform-set
match address 101
crypto map mymap 2 ipsec-isakmp
set peer 10.1.1.1
set transform-set my-transform-set
match address 102
access-list 101 permit ip 192.168.1.0 0.0.0.255 172.17.11.0 0.0.0.255
access-list 102 permit ip 192.168.1.0 0.0.0.255 10.0.0.0 0.0.255.255
interface FastEthernet 0/0
crypto map mymap
Example: Configuring HSRP with IPsec
The following example shows how all remote VPN gateways connect to the router via 192.168.0.3. The crypto map on the interface
binds this standby address as the local tunnel endpoint for all instances of the crypto map named
mymap and at the same time ensures that HSRP failover is facilitated between an active and standby device belonging to the same
standby group named group1.
Note that RRI also provides the ability for only the active device in the HSRP group to be advertising itself to inside devices
as the next hop VPN gateway to the remote proxies. If there is a failover, routes are deleted on the formerly active device
and created on the newly active device.
crypto map mymap 1 ipsec-isakmp
set peer 10.1.1.1
reverse-route
set transform-set esp-aes-sha
match address 102
Interface FastEthernet 0/0
ip address 192.168.0.2 255.255.255.0
standby name group1
standby ip 192.168.0.3
crypto map mymap redundancy group1
access-list 102 permit ip 192.168.1.0 0.0.0.255 10.0.0.0 0.0.255.255
