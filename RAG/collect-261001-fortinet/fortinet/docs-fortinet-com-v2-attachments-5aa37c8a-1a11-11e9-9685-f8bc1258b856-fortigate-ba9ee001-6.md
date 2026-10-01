---
id: collect-261001-fortinet/fortinet/docs-fortinet-com-v2-attachments-5aa37c8a-1a11-11e9-9685-f8bc1258b856-fortigate-ba9ee001-6
title: "docs-fortinet-com-v2-attachments-5aa37c8a-1a11-11e9-9685-f8bc1258b856-fortigate--ba9ee001"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-fortinet/docs-fortinet-com-v2-attachments-5aa37c8a-1a11-11e9-9685-f8bc1258b856-fortigate--ba9ee001.md
source_anchor: ""
source_lines: [818, 1004]
sha256: f9d04e4a563cf6bde1755dc06caffbb8872f0c71dce7010f347d4b0aba41164c
---

# docs-fortinet-com-v2-attachments-5aa37c8a-1a11-11e9-9685-f8bc1258b856-fortigate--ba9ee001

Networking in Transparent Mode Configuration Example
edit "LACP_VD2_OUT"
set vdom "VD2"
set stpforward enable
set type aggregate
set member "port19" "port20"
next
edit "LACP_VD1"
set vdom "VD1"
set stpforward enable
set type aggregate
set member "port1" "port2"
next
end
When using aggregation, the stpforward setting needs to be applied
only on the port aggregation level, not on the physical port
This will also forward regular Spanning Tree BPDUs
Verification with a sniffer trace:
FGT# diagnose sniffer packet any "" 4
41.365434 port3 in llc unnumbered, ui, flags [command], length 72
41.365437 LACP_VD1 out llc unnumbered, ui, flags [command], length 72
41.365439 port2 out llc unnumbered, ui, flags [command], length 72
41.365479 LACP_VD2_IN in llc unnumbered, ui, flags [command], length 72
41.365482 LACP_VD2_OUT out llc unnumbered, ui, flags [command], length 72
41.365484 port19 out llc unnumbered, ui, flags [command], length 72
See above the CDP packet flow from port3, LACP_VD1 (port2), LACP_VD2_IN, LACP_VD2_OUT (port19).
The following sniffer trace command will filter only CDP or VTP packets :
FGT# diagnose sniffer packet port_name "ether host 01-00-
0C-CC-CC-CC"
Configuration Example
Step 1: Create VLANs and forwarding domains
config system interface
edit "vlan102_intern"
Transparent Mode for FortiOS 5.6.3
Fortinet Technologies Inc.
27

Configuration Example Networking in Transparent Mode
set forward-domain 102
set interface "port2"
set vlanid 102
next
edit "vlan102_extern"
set forward-domain 102
set interface "port3"
set vlanid 102
next
edit "vlan103_intern"
set forward-domain 103
set interface "port2"
set vlanid 103
next
edit "vlan103_extern"
set forward-domain 103
set interface "port3"
set vlanid 103
next
end
Step 2: Create the appropriate Firewall Policies
config firewall policy
edit 1
set srcintf "vlan102_extern"
set dstintf "vlan102_intern"
set srcaddr "all"
set dstaddr "all"
set action accept
set schedule "always"
set service "ANY" next
edit 2
set srcintf "vlan102_intern"
set dstintf "vlan102_extern"
set srcaddr "all"
set dstaddr "all"
set action accept
set schedule "always"
set service "ANY" next
edit 3
set srcintf "vlan103_intern"
set dstintf "vlan103_extern"
set srcaddr "all"
set dstaddr "all"
set action accept
set schedule "always"
set service "ANY" next
edit 4
set srcintf "vlan103_extern"
set dstintf "vlan103_intern"
set srcaddr "all"
set dstaddr "all"
set action accept
set schedule "always"
set service "ANY" next
end
28 Transparent Mode for FortiOS 5.6.3
Fortinet Technologies Inc.

Firewalls and Security in Transparent Mode
This section contains information about using firewalls and security scanning in Transparent mode. It contains
the following topics:
l Firewall Policy Look Up
l Firewall Session List
l Security Scanning
Firewall Policy Look Up
In Transparent mode, like in NAT/Route mode, a firewall policy look up is based on the source and destination
interfaces. The matching firewall policy will tell which actions to apply to the traffic, including logging and security
scanning.
The FortiGate proceeds as follows to look for a matching firewall policy in Transparent mode:
l Step 1: an Ethernet IP frame ingresses a port (or a VLAN on a port), corresponding to a specific bridge instance
(from the port VDOM and Forwarding domain). This frame contains a destination MAC address that we will call
MAC_D.
l Step 2: The FortiGate is making a MAC_D address lookup in the bridge instance to determine the port where
MAC_D has been learned. This will be the destination interface.
l Step 3: The FortiGate is then looking for a firewall policy corresponding to the couple < source interface +
destination interface >. If multiple policies with the same couple < source interface + destination interface > exist,
the FortiGate screens all of them from TOP to BOTTOM (as displayed in the configuration), until a match is found.
It is important to make sure that the most specific firewall policies are located at the top of the policy list, to make
sure that traffic is matched to the appropriate policy.
Firewall Session List
The flag br in the state line will indicate that this is a “bridged” session. See example below :
FGT# diagnose sys session list
session info: proto=17 proto_state=00 duration=59 expire=128 timeout=0 flags=000
00000 sockflag=00000000 sockport=0 av_idx=0 use=4
origin-shaper=
reply-shaper=
per_ip_shaper=
ha_id=0 hakey=0
policy_dir=0 tunnel=/
state=may_dirty br rem
statistic(bytes/packets/allow_err): org=385/5/1 reply=0/0/0 tuples=2
orgin->sink: org pre->post, reply pre->post dev=3->4/4->3 gwy=0.0.0.0/0.0.0.0
hook=pre dir=org act=noop 192.168.182.93:1025->4.2.2.1:53(0.0.0.0:0)
hook=post dir=reply act=noop 4.2.2.1:53->192.168.182.93:1025(0.0.0.0:0)
misc=0 policy_id=1 auth_info=0 chk_client_info=0 vd=1 serial=000006d3 tos=ff/ff
Transparent Mode for FortiOS 5.6.3
Fortinet Technologies Inc.
29

Security Scanning Firewalls and Security in Transparent Mode
imp2p=0 app=0
dd_type=0 dd_rule_id=0
Security Scanning
Security scanning occurs in the same manner in NAT/Route mode and Transparent mode. When a protection
profile is enabled on a firewall policy for content inspection, the FortiGate acts like a transparent proxy for the
protocols that need to be inspected.
The FortiGate will therefore intercept the TCP sessions and create its own session from client to server and
server to client. The source and destination MAC addresses of the original L2 frames are however not altered in
this communication, as described in the section Network operation : source MAC addresses in frames sent by or
through the FortiGate.
Devices in the network communicating through the FortiGate do not know the
presence of the FortiGate.
For more information about security scanning, see the Security Profiles handbook.
30 Transparent Mode for FortiOS 5.6.3
Fortinet Technologies Inc.

IPsec VPN in Transparent Mode
This section contains information about configuring IPsec virtual private networks (VPNs) in Transparent mode. It
contains the following topics:
l Using IPsec VPNs in Transparent Mode
l Example 1: Remote Sites with Different Subnets
l Example 2: Remote Sites on the Same Subnet
Using IPsec VPNs in Transparent Mode
In Transparent mode, IPsec VPN is supported in Policy-based configuration mode only.
IPsec VPN in Transparent mode can be used in those scenarios:
l Encrypt data over routed networks without changing anything on the routers. See example 1.
l Encrypt data over a non-routed transport network (extension of a LAN for example). See example 2.
The following rules apply to IPsec in Transparent mode:
l If both remote FortiGate IPsec gateways are not in the same broadcast domain (separated by routers):
l The hosts on each side must be on different subnets.
l The FortiGate management IP addresses must be in the same subnet as the local hosts. This is the preferred
option.
l If both remote FortiGate IPsec gateways are in the same broadcast domain (separated by optical switches for
examples), the hosts on each side can be :
l On the same subnet
l On different subnet if the appropriate static route is configured on the remote Fortigate
l The FortiGate management IP addresses can be in any different subnet than the local hosts
l A firewall Policy with the action IPsec is used to send traffic to the remote device into the tunnel.
Therefore, it is important to place all remote devices on the appropriate ports of the Fortigate to allow a proper
match < source interface + destination interface > . See section Transparent mode Firewall processing for more
details.
This scenario requires that the remote hosts located on the remote FortiGate’s
protected subnets have their MAC addresses hardcoded in FortiGate’s static MAC
entry list. If this is not configured then it is expected to see outage in network
communications.
Transparent Mode for FortiOS 5.6.3
Fortinet Technologies Inc.
31

