---
id: collect-261001-fortinet/fortinet/document-fortigate-7-4-0-administration-guide-31670-ikev2-ipsec-site-to-site-vpn-ff73872e-2
title: "document-fortigate-7-4-0-administration-guide-31670-ikev2-ipsec-site-to-site-vpn-ff73872e"
domain: fortinet
role: reference
task: reference
actors: ["AWS", "Apple"]
dates: []
keywords: ["aws", "parameters"]
source: docs/RAG/collect-261001-fortinet/document-fortigate-7-4-0-administration-guide-31670-ikev2-ipsec-site-to-site-vpn-ff73872e.md
source_anchor: ""
source_lines: [3, 156]
sha256: 312d0cd0b4b16efcb7325b42d69e9378d1128c26593f63ab1bc910526d92e5e1
---

# document-fortigate-7-4-0-administration-guide-31670-ikev2-ipsec-site-to-site-vpn-ff73872e

                                            This is a sample configuration of an IPsec site-to-site VPN connection between an on-premise FortiGate and an AWS virtual private cloud (VPC).
AWS uses unique identifiers to manipulate a VPN connection's configuration. Each VPN connection is assigned an identifier and is associated with two other identifiers: the customer gateway ID for the FortiGate and virtual private gateway ID.
This example includes the following IDs:
- VPN connection ID: vpn-07e988ccc1d46f749
- Customer gateway ID: cgw-0440c1aebed2f418a
- Virtual private gateway ID
This example assumes that you have configured VPC-related settings in the AWS management portal as described in Create a Secure Connection using AWS VPC.
This example includes creating and configuring two tunnels. You must configure both tunnels on your FortiGate.
To configure IKEv2 IPsec site-to-site VPN to an AWS VPN gateway:
- Configure the first VPN tunnel:
- Configure the second VPN tunnel:
To configure IKE for the first VPN tunnel:
A policy is established for the supported ISAKMP encryption, authentication, Diffie-Hellman (DH), lifetime, and key parameters. These sample configurations fulfill the minimum requirements for AES128, SHA1, and DH Group 2. Category VPN connections in the GovCloud AWS region have a minimum requirement of AES128, SHA2, and DH Group 14. To take advantage of AES256, SHA256, or other DH groups such as 14-18, 22, 23, and 24, you must modify these sample configuration files. Higher parameters are only available for VPNs of category "VPN", not for "VPN-Classic".
Your FortiGate's external interface's address must be static. Your FortiGate may reside behind a device performing NAT. To ensure NAT traversal can function, you must adjust your firewall rules to unblock UDP port 4500. If not behind NAT, it is recommended to disable NAT traversal.
Begin configuration in the root VDOM. The interface name must be shorter than 15 characters. It is best if the name is shorter than 12 characters. IPsec dead peer detection (DPD) causes periodic messages to be sent to ensure a security association remains operational.
config vpn ipsec phase1-interface
edit vpn-07e988ccc1d46f749-0
set interface "wan1"
set dpd enable
set local-gw 35.170.66.108
set dhgrp 2
set proposal aes128-sha1
set keylife 28800
set remote-gw 3.214.239.164
set psksecret iCelks0UOob8z4SYMRM6zlx.rU2C3jth
set dpd-retryinterval 10
next
end
To configure IPsec for the first VPN tunnel:
The IPsec transform set defines the encryption, authentication, and IPsec mode parameters.
config vpn ipsec phase2-interface
edit "vpn-07e988ccc1d46f749-0"
set phase1name "vpn-07e988ccc1d46f749-0"
set proposal aes128-sha1
set dhgrp 2
set pfs enable
set keylifeseconds 3600
next
end
To configure the tunnel interface for the first VPN tunnel:
You must configure a tunnel interface as the logical interface associated with the tunnel. All traffic routed to the tunnel interface must be encrypted and transmitted to the VPC. Similarly, traffic from the VPC will be logically received on this interface.
You must configure the interface's address with your FortiGate's address. If the address changes, you must recreate the FortiGate and VPN connection with Amazon VPC.
The tcp-mss option causes the router to reduce the TCP packets' maximum segment size to prevent packet fragmentation.
config system interface
edit "vpn-07e988ccc1d46f749-0"
set vdom "root"
set ip 169.254.45.90 255.255.255.255
set allowaccess ping
set type tunnel
set tcp-mss 1379
set remote-ip 169.254.45.89
set mtu 1427
set interface "wan1"
next
end
To configure BGP for the first VPN tunnel:
BGP is used within the tunnel to exchange prefixes between the virtual private gateway and your FortiGate. The virtual private gateway announces the prefix according to your VPC.
The local BGP autonomous system number (ASN) (65000) is configured as part of your FortiGate. If you must change the ASN, you must recreate the FortiGate and VPN connection with AWS.
Your FortiGate may announce a default route (0.0.0.0/0) to AWS. This is done using a prefix list and route map in FortiOS.
config router bgp
set as 65000
config neighbor
edit 169.254.45.89
set remote-as 64512
end
end
end
config router bgp
config neighbor
edit 169.254.45.89
set capability-default-originate enable
end
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
To configure firewall policies for the first VPN tunnel:
Create a firewall policy permitting traffic from your local subnet to the VPC subnet, and vice-versa.
This example policy permits all traffic from the local subnet to the VPC. First, view all existing policies using the show firewall policy command. Then, create a new firewall policy starting with the next available policy ID. In this example, running show firewall policy displayed policies 1, 2, 3, and 4, so you would proceed to create policy 5.
config firewall policy
edit 5
set srcintf "vpn-07e988ccc1d46f749-0"
set dstintf internal
set srcaddr all
set dstaddr all
set action accept
set schedule always
set service ANY
next
end
config firewall policy
edit 5
set srcintf internal
set dstintf "vpn-07e988ccc1d46f749-0"
set srcaddr all
set dstaddr all
set action accept
set schedule always
set service ANY
next
end
To configure IKE for the second VPN tunnel:
A policy is established for the supported ISAKMP encryption, authentication, DH, lifetime, and key parameters. These sample configurations fulfill the minimum requirements for AES128, SHA1, and DH Group 2. Category VPN connections in the GovCloud AWS region have a minimum requirement of AES128, SHA2, and DH Group 14. To take advantage of AES256, SHA256, or other DH groups such as 14-18, 22, 23, and 24, you must modify these sample configuration files. Higher parameters are only available for VPNs of category "VPN", not for "VPN-Classic".
Your FortiGate's external interface's address must be static. Your FortiGate may reside behind a device performing NAT. To ensure NAT traversal can function, you must adjust your firewall rules to unblock UDP port 4500. If not behind NAT, it is recommended to disable NAT traversal.
Begin configuration in the root VDOM. The interface name must be shorter than 15 characters. It is best if the name is shorter than 12 characters. IPsec DPD causes periodic messages to be sent to ensure a security association remains operational.
config vpn ipsec phase1-interface
edit vpn-07e988ccc1d46f749-1
set interface "wan1"
set dpd enable
set local-gw 35.170.66.108
set dhgrp 2
set proposal aes128-sha1
set keylife 28800
set remote-gw 100.25.187.58
set psksecret IjFzyDneUtDdAT4RNmQ85apUG3y4Akre
set dpd-retryinterval 10
next
end
To configure IPsec for the second VPN tunnel:
The IPsec transform set defines the encryption, authentication, and IPsec mode parameters.
config vpn ipsec phase2-interface
edit "vpn-07e988ccc1d46f749-1"
set phase1name "vpn-07e988ccc1d46f749-1"
set proposal aes128-sha1
set dhgrp 2
set pfs enable
set keylifeseconds 3600
next
end
To configure the tunnel interface for the second VPN tunnel:
You must configure a tunnel interface as the logical interface associated with the tunnel. All traffic routed to the tunnel interface must be encrypted and transmitted to the VPC. Similarly, traffic from the VPC will be logically received on this interface.
