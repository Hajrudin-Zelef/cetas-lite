---
id: collect-261001-fortinet/fortinet/docs-fortinet-com-v2-attachments-5aa37c8a-1a11-11e9-9685-f8bc1258b856-fortigate-ba9ee001-2
title: "docs-fortinet-com-v2-attachments-5aa37c8a-1a11-11e9-9685-f8bc1258b856-fortigate--ba9ee001"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/docs-fortinet-com-v2-attachments-5aa37c8a-1a11-11e9-9685-f8bc1258b856-fortigate--ba9ee001.md
source_anchor: ""
source_lines: [204, 326]
sha256: cf92dc786d50fee4d68c6169003eb73032fa5f017426d4633cb9de4bf3d84a0e
---

# docs-fortinet-com-v2-attachments-5aa37c8a-1a11-11e9-9685-f8bc1258b856-fortigate--ba9ee001

Transparent Mode Installation
This section contains information about installing a FortiGate in Transparent mode and using a virtual wire pair,
to simplify Transparent mode. It also contains information about configuring the FortiGate's management IP .
Installing a FortiGate in Transparent mode
Changing to Transparent mode removes most configuration changes made in
NAT/Route mode. To keep your current NAT/Route mode configuration, backup the
configuration using the System Information widget, found in the Dashboard.
1. Before connecting the FortiGate unit to your network, go to the Dashboard and locate enter the following
command into the CLI Console:
config system settings
set opmode transparent
set manageip <address and netmask>
set gateway <address>
end
2. Access the web-based manager by browsing to the new management IP.
3. (Optional) The FortiGate unit’s DNS Settings are set to use FortiGuard DNS servers by default, which is sufficient
for most networks. However, if you need to change the DNS servers, go to Network > DNS and add Primary and
SecondaryDNS servers. Select Apply.
4. If your network uses IPv4 addresses, go to Policy & Objects > IPv4 Policy and select Create New to add a
security policy that allows users on the private network to access the Internet.
If your network uses IPv6 addresses, go to Policy & Objects > IPv6 Policy and select Create New to add a
security policy that allows users on the private network to access the Internet. If the IPv6 menu option is not
available, go to System > Feature Visibility, turn on IPv6, and select Apply. For more information on IPv6
networks, see the IPv6 Handbook.
Transparent Mode for FortiOS 5.6.3
Fortinet Technologies Inc.
11

Using a Virtual Wire Pair to Simplify Transparent Mode Transparent Mode Installation
Set the Incoming Interface to the internal interface and the Outgoing Interface to the Internet-facing interface
(typically WAN1). You will also need to set Source Address, Destination Address, Schedule, and Service
according to your network requirements. You can set these fields to the default all/ANY settings for now but should
create the appropriate objects later after the policies have been verified.
5. Make sure the Action is set to ACCEPT. Select OK.
It is recommended to avoid using any security profiles, such as AntiVirus or web
filtering, until after you have successfully installed the FortiGate unit. After the
installation is verified, you can apply any required security profiles.
For more information about using security profiles, see the Security Profiles handbook.
6. Go to the Dashboard and locate the System Resources widget. Select Shutdown to power off the FortiGate
unit.
Alternatively, you can also use the CLI command execute shutdown .
7. Connect the FortiGate unit between the internal network and the router.
8. Connect the Internet-facing interface to the router’s internal interface and connect the internal network to the
FortiGate using an internal port (typically port 1).
9. Power on the FortiGate unit. You will experience downtime before the FortiGate unit starts up completely.
Results
Users on the internal network are now able to browse to the Internet. They should also be able to connect to the
Internet using any other protocol or connection method that you defined in the security policy.
If a FortiGate unit operating in Transparent mode is installed between your internet
network and a server that is providing a network service to the internal network, such
as DNS or DHCP, you must add a wan1-to-internal policy to allow the server’s
response to flow through the FortiGate unit and reach the internal network.
Using a Virtual Wire Pair to Simplify Transparent Mode
A virtual wire pair consists of two interfaces that have no IP addresses and all traffic received by one interface in
the pair can only be forwarded out the other; as controlled by firewall policies. Since the interfaces do not have IP
addresses, you can insert a virtual wire pair into a network without having to make any changes to the network.
12 Transparent Mode for FortiOS 5.6.3
Fortinet Technologies Inc.

Transparent Mode Installation Management IP configuration
Interfaces used in a virtual wire pair cannot be used for admin access to the ISFW
FortiGate. Before creating a virtual wire pair, make sure you have a different port
configured to allow admin access using your preferred protocol.
1. Go to Network > Interfacesand select Create New > Virtual Wire Pair . Add two ports to the virtual wire pair.
These ports cannot be part of a switch, such as the default internal/lan interface.
2. Go to Policy & Objects > IPv4 Virtual Wire Pair Policy and create a policy will allow traffic to flow between
the two ports. Give the policy an appropriate Name. Select the direction that traffic is allowed to flow. Configure
the other firewall options as needed.
3. If necessary, create a second virtual wire pair policy allowing traffic to flow between the ports in the opposite
direction.
Traffic can now flow between the two ports. Go to FortiView > Policies to see traffic flowing through both
policies.
Management IP configuration
A FortiGate in Transparent mode can be assigned with a single IP address for remote access management and
multiple static routes can be configured. This can be used if in-band management wants to be applied.
When out-of-band management is desired (dedicated interface for remote management access), it is
recommended to use a separate VDOM in NAT/Route mode.
In-band management details and example
The management IP address is bound to all ports or VLANs belonging to the same VDOM. Remote access
services are subject to the same rules as in NAT/Route mode, and must be enabled/disabled on each port.
Example of management IP configuration in Transparent mode:
config system settings
set manageip 10.1.1.100/255.255.255.0
end
config router static
edit 1
set gateway 10.1.1.254
next
end
config system interface
edit port1
set allowaccess ping ssh https snmp
end
It is also possible to add a second IP address for management and additional default routes:
config system settings
set opmode transparent
set manageip 192.168.182.136/255.255.254.0 10.1.1.1/255.255.255.0
end
Transparent Mode for FortiOS 5.6.3
Fortinet Technologies Inc.
13

Management IP configuration Transparent Mode Installation
config router static
edit 1
set gateway 192.168.183.254
next
edit 2
set gateway 10.1.1.254
next
end
ping-server (dead gateway detection) is not supported in Transparent mode.
Out-of-band management details and example
When VDOM is enabled and the VDOMs are operating in Transparent mode, it is recommended, to avoid L2
loops and allow more routing flexibility, to keep one VDOM (generally the root VDOM) in NAT/Route mode, with
one or more VLAN or physical interface as out-of-band management.
The management VDOM must have IP connectivity to the Internet to allow
communication with the FDS and retrieve services information (AV,IPS, Fortiguard
filtering, Support…). All syslog and FortiManager communication also go through the
management VDOM.
14 Transparent Mode for FortiOS 5.6.3
Fortinet Technologies Inc.

