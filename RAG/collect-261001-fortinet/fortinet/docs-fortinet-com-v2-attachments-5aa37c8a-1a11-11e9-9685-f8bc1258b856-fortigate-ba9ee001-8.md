---
id: collect-261001-fortinet/fortinet/docs-fortinet-com-v2-attachments-5aa37c8a-1a11-11e9-9685-f8bc1258b856-fortigate-ba9ee001-8
title: "docs-fortinet-com-v2-attachments-5aa37c8a-1a11-11e9-9685-f8bc1258b856-fortigate--ba9ee001"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/docs-fortinet-com-v2-attachments-5aa37c8a-1a11-11e9-9685-f8bc1258b856-fortigate--ba9ee001.md
source_anchor: ""
source_lines: [1192, 1372]
sha256: 18a08e4b54f0c7a0599b315007ee0307e2572cf3e19630ca1575a4e00723e4d0
---

# docs-fortinet-com-v2-attachments-5aa37c8a-1a11-11e9-9685-f8bc1258b856-fortigate--ba9ee001

Example 1: Remote Sites with Different Subnets IPsec VPN in Transparent Mode
ah=sha1 key=20 e316113eb6ea03b014b3a5f9c1a3bd386637801a
The above tunnel is up!
Verify that destination local hosts are seen in the ARP table (necessary for IPsec despite being in TP
mode)
FGT2 # get system arp
Address Age(min) Hardware Addr Interface
10.2.2.10 2 00:50:56:00:76:04 root.b
10.2.2.254 0 00:09:0f:30:29:e4 root.b
Using the debug flow command on the initiator side (example on FortiGate1)
FGT1 # diagnose debug flow filter addr 10.1.1.10
FGT1 # diagnose debug flow show console enable
FGT1 # diagnose debug enable
FGT1 # diagnose debug flow trace start 50
FGT1 # id=36870 trace_id=615 msg="vd-root received a packet(proto=1, 10.1.1.10:512-
>10.2.2.10:8) from port1."
id=36870 trace_id=615 msg="allocate a new session-00000636"
id=36870 trace_id=615 msg="Allowed by Policy-1: encrypt"
id=36870 trace_id=615 msg="enter IPsec tunnel-to_FGT2"
id=36870 trace_id=615 msg="SA is not ready yet, drop"
id=36870 trace_id=616 msg="vd-root received a packet(proto=1, 10.1.1.10:512->10.2.2.10:8)
from port1."
id=36870 trace_id=616 msg="Find an existing session, id-00000636, original direction"
id=36870 trace_id=616 msg="enter IPsec tunnel-to_FGT2"
id=36870 trace_id=616 msg="encrypted, and send to 10.2.2.100 with source 10.1.1.100"
id=36870 trace_id=616 msg="send out via dev-port2, dst-mac-00:09:0f:30:29:e0"
id=36870 trace_id=617 msg="vd-root received a packet(proto=1, 10.1.1.10:512->10.2.2.10:8)
from port1."
id=36870 trace_id=617 msg="Find an existing session, id-00000636, original direction"
id=36870 trace_id=617 msg="enter IPsec tunnel-to_FGT2"
id=36870 trace_id=617 msg="encrypted, and send to 10.2.2.100 with source 10.1.1.100"
id=36870 trace_id=617 msg="send out via dev-port2, dst-mac-00:09:0f:30:29:e0"
id=36870 trace_id=618 msg="vd-root received a packet(proto=1, 10.2.2.10:512->10.1.1.10:0)
from port2."
id=36870 trace_id=618 msg="Find an existing session, id-00000636, reply direction"
id=36870 trace_id=618 msg="send out via dev-port1, dst-mac-00:50:56:00:76:03"
id=36870 trace_id=619 msg="vd-root received a packet(proto=1, 10.1.1.10:512->10.2.2.10:8)
from port1."
id=36870 trace_id=619 msg="Find an existing session, id-00000636, original direction"
id=36870 trace_id=619 msg="enter IPsec tunnel-to_FGT2"
id=36870 trace_id=619 msg="encrypted, and send to 10.2.2.100 with source 10.1.1.100"
id=36870 trace_id=619 msg="send out via dev-port2, dst-mac-00:09:0f:30:29:e0"
id=36870 trace_id=620 msg="vd-root received a packet(proto=1, 10.2.2.10:512->10.1.1.10:0)
from port2."
id=36870 trace_id=620 msg="Find an existing session, id-00000636, reply direction"
id=36870 trace_id=620 msg="send out via dev-port1, dst-mac-00:50:56:00:76:03"
36 Transparent Mode for FortiOS 5.6.3
Fortinet Technologies Inc.

IPsec VPN in Transparent Mode Example 1: Remote Sites with Different Subnets
The message "id=36870 trace_id=615 msg="SA is not ready yet,
drop " simply means that the tunnel was not up yet.
Using the debug flow command on the receiver side (example on FortiGate2)
FGT2 # diagnose debug flow filter addr 10.1.1.10
FGT2 # diagnose debug flow show console enable
FGT2 # diagnose debug enable
FGT2 # diagnose debug flow trace start 50
FGT2 # id=36870 trace_id=51 msg="vd-root received a packet(proto=1, 10.1.1.10:512-
>10.2.2.10:8) from port2."
id=36870 trace_id=51 msg="allocate a new session-00000435"
id=36870 trace_id=51 msg="Allowed by Policy-1:"
id=36870 trace_id=51 msg="send out via dev-port1, dst-mac-00:50:56:00:76:04"
id=36870 trace_id=52 msg="vd-root received a packet(proto=1, 10.2.2.10:512->10.1.1.10:0)
from port1."
id=36870 trace_id=52 msg="Find an existing session, id-00000435, reply direction"
id=36870 trace_id=52 msg="enter IPsec tunnel-to_FGT1"
id=36870 trace_id=52 msg="encrypted, and send to 10.1.1.100 with source 10.2.2.100"
id=36870 trace_id=52 msg="send out via dev-port2, dst-mac-00:09:0f:30:29:e4"
Using the sniffer trace (example on FortiGate2)
FGT2 # diagnose sniffer packet any "host 10.2.2.10" 4
9.460021 root.b out arp who-has 10.2.2.10 tell 10.2.2.100
9.460028 port2 out arp who-has 10.2.2.10 tell 10.2.2.100
9.460034 port1 out arp who-has 10.2.2.10 tell 10.2.2.100
9.460462 port1 in arp reply 10.2.2.10 is-at 0:50:56:0:76:4
9.460462 root.b in arp reply 10.2.2.10 is-at 0:50:56:0:76:4
[...]
49.477368 port2 in 10.1.1.10 -> 10.2.2.10: icmp: echo request
49.477444 port1 out 10.1.1.10 -> 10.2.2.10: icmp: echo request
49.477898 port1 in 10.2.2.10 -> 10.1.1.10: icmp: echo reply
50.510023 port2 in 10.1.1.10 -> 10.2.2.10: icmp: echo request
50.510079 port1 out 10.1.1.10 -> 10.2.2.10: icmp: echo request
50.510524 port1 in 10.2.2.10 -> 10.1.1.10: icmp: echo reply
The above ARP process in Transparent mode with IPsec is allowing the Fortigate to:
l Identify the MAC address of the destination device 10.2.2.10
l Populate the MAC table (see below), which in turn will give a destination interface
and allow a Firewall policy look-up
Check the FDB entries for the destination
FGT2 # diagnose netlink brctl name host root.b
port no device devname mac addr ttl
[..]
1 2 port1 00:50:56:00:76:04 0
Transparent Mode for FortiOS 5.6.3
Fortinet Technologies Inc.
37

Example 2: Remote Sites on the Same Subnet IPsec VPN in Transparent Mode
Example 2: Remote Sites on the Same Subnet
This example provides a configuration example for IPsec VPN tunnels between two FortiGate in Transparent
Mode in the same subnet separated by a L2 transparent network and one remote subnet on the second site.
This scenario requires that PC1’s MAC address is added to the FortiGate’s static MAC
table. The preferred scenario would be to have a router installed between the 2
FortiGate’s.
The expectation for this example is that PC1 will be able to communicate via the IPsec tunnel with Server1 in the
same subnet, and Server2 in a different subnet.
The requirements for this example are:
l The default gateway (FGT3) for PC1 and all remote device must be behind port2 of FGT1, in order for this
FortiGate to match the appropriate Encrypt firewall policy (port1 --> port2)
l Despite being in Transparent mode, FGT2 must have a valid route to Server2
l FGT3 is used as a router between subnet 10.1.1.0/24 and 10.3.3.0/24.
PC1 MAC address added to FGT2 static MAC entries.
Server1 MAC address added to FGT1 static MAC entries.
Configuration of FortiGate 1 (FGT1):
Only relevant parts of configuration are provided.
config system settings
set opmode transparent
set manageip 10.1.1.100/255.255.255.0
end
config router static
edit 1
set gateway 10.1.1.252
next
38 Transparent Mode for FortiOS 5.6.3
Fortinet Technologies Inc.

IPsec VPN in Transparent Mode Example 2: Remote Sites on the Same Subnet
end
config system mac-address-table
edit 00:50:56:00:76:04 ==>Server1
set interface port2
next
end
config firewall address
edit "all"
next
edit "Server1"
set subnet 10.1.1.20 255.255.255.255
next
edit "Server2"
set subnet 10.3.3.30 255.255.255.255
next
edit "10.1.1.0/24"
set subnet 10.1.1.0 255.255.255.0
next
edit "gateway"
set subnet 10.1.1.254 255.255.255.255
next
end
config vpn ipsec phase1
edit "to_FGT2"
set proposal 3des-sha1 aes128-sha1 des-md5
set remote-gw 10.1.1.200
set psksecret fortinet
next
end
config vpn ipsec phase2
edit "to_FGT2"
set keepalive enable
set phase1name "to_FGT2"
set proposal 3des-sha1 aes128-sha1
set src-subnet 10.1.1.0 255.255.255.0
next
end
config firewall policy
edit 1
set srcintf "port1"
set dstintf "port2"
set srcaddr "10.1.1.0/24"
set dstaddr "Server1"
set action ipsec
set schedule "always"
set service "ANY"
set inbound enable
set outbound enable
set vpntunnel "to_FGT2"
next
edit 2
set srcintf "port1"
Transparent Mode for FortiOS 5.6.3
Fortinet Technologies Inc.
39

