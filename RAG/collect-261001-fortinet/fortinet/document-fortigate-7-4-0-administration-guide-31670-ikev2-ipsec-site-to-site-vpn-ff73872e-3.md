---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-0-administration-guide-31670-ikev2-ipsec-site-to-site-vpn-ff73872e-3
title: "document-fortigate-7-4-0-administration-guide-31670-ikev2-ipsec-site-to-site-vpn-ff73872e"
domain: fortinet
role: reference
task: reference
actors: ["AWS", "Apple"]
dates: []
keywords: ["aws"]
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-0-administration-guide-31670-ikev2-ipsec-site-to-site-vpn-ff73872e.md
source_anchor: ""
source_lines: [157, 236]
sha256: 0aa4aaefe3b0a3a5b86c873f946daecb38a1776fb11c5eb0163c4ab6a9ef96fc
---

# document-fortigate-7-4-0-administration-guide-31670-ikev2-ipsec-site-to-site-vpn-ff73872e

You must configure the interface's address with your FortiGate's address. If the address changes, you must recreate the FortiGate and VPN connection with Amazon VPC.
The tcp-mss option causes the router to reduce the TCP packets' maximum segment size to prevent packet fragmentation.
config system interface
edit "vpn-07e988ccc1d46f749-1"
set vdom "root"
set ip 169.254.44.162 255.255.255.255
set allowaccess ping
set type tunnel
set tcp-mss 1379
set remote-ip 169.254.44.161
set mtu 1427
set interface "wan1"
next
end
To configure BGP for the second VPN tunnel:
BGP is used within the tunnel to exchange prefixes between the virtual private gateway and your FortiGate. The virtual private gateway announces the prefix according to your VPC.
The local BGP ASN (65000) is configured as part of your FortiGate. If you must change the ASN, you must recreate the FortiGate and VPN connection with AWS.
Your FortiGate may announce a default route (0.0.0.0/0) to AWS. This is done using a prefix list and route map in FortiOS.
config router bgp
set as 65000
config neighbor
edit 169.254.44.161
set remote-as 64512
end
config router bgp
config neighbor
edit 169.254.44.161
set capability-default-originate enable
end
end
config router prefix-list
edit "default_route"
config rule
edit 1
set prefix 0.0.0.0 0.0.0.0
next
end
end
end
config router route-map
edit "routemap1"
config rule
edit 1
set match-ip-address "default_route"
next
end
next
end
To advertise additional prefixes to the Amazon VPC, add these prefixes to the network statement and identify the prefix you want to advertise. Ensure that the prefix is present in the routing table of the device with a valid next-hop. If you want to advertise 192.168.0.0/16 to Amazon, you would do the following:
config router bgp
config network
edit 1
set prefix 192.168.0.0 255.255.0.0
next
end
To configure firewall policies for the second VPN tunnel:
Create a firewall policy permitting traffic from your local subnet to the VPC subnet, and vice-versa.
This example policy permits all traffic from the local subnet to the VPC. First, view all existing policies using the show firewall policy command. Then, create a new firewall policy starting with the next available policy ID. In this example, running show firewall policy displayed policies 1, 2, 3, 4, and 5, so you would proceed to create policy 6.
config firewall policy
edit 6
set srcintf "vpn-07e988ccc1d46f749-1"
set dstintf internal
set srcaddr all
set dstaddr all
set action accept
set schedule always
set service ANY
next
end
config firewall policy
edit 6
set srcintf internal
set dstintf "vpn-07e988ccc1d46f749-1"
set srcaddr all
set dstaddr all
set action accept
set schedule always
set service ANY
next
end
