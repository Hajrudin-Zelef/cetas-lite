---
id: collect-261001-huawei/huawei/hedex-api-pages-edoc1100331435-aem10132-05-resources-dc-dc-cfg-ipsec-0053-html-287aa4a2-1
title: "Assign an IP address to an interface on RouterA."
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/hedex-api-pages-edoc1100331435-aem10132-05-resources-dc-dc-cfg-ipsec-0053-html-287aa4a2.md
source_anchor: ""
source_lines: [1, 108]
sha256: cc8325bac152c7aa372d4657816e1ca70304e3ec5df9ac60743388c4c2b595d9
---

# Assign an IP address to an interface on RouterA.

As shown in Figure 1, RouterA (remote small-scale branch gateway) and RouterB (headquarters gateway) communicate through the Internet. The headquarters and branch networks are not planned uniformly. The branch subnet is 10.1.1.0/24 and the headquarters subnet is 10.1.2.0/24. The DHCP server is located on the headquarters network and allocates an IP address to the branch gateway.
The enterprise requires that traffic between headquarters and branch networks should be securely transmitted and the headquarters gateway should manage the branch gateway with simplified configuration in centralized manner. An Efficient VPN policy in client mode can be used to establish an IPSec tunnel to protect traffic. This method facilitates IPSec tunnel establishment and maintenance.
In client mode, RouterA requests an IP address from RouterB to establish an IPSec tunnel, and requests the DNS domain name, DNS server IP addresses, and WINS server IP addresses for the branch subnet.
The configuration roadmap is as follows:
Configure IP addresses and static routes for interfaces on RouterA and RouterB so that routes between RouterA and RouterB are reachable.
Configure the DHCP server address on RouterB so that IP addresses can be dynamically allocated through DHCP.
Configure RouterB as the responder to use an IPSec policy template to establish an IPSec tunnel with RouterA.
Configure an Efficient VPN policy in client mode on RouterA. RouterA as the initiator establishes an IPSec tunnel with RouterB.
# Assign an IP address to an interface on RouterA.
<Huawei> system-view
[Huawei] sysname RouterA
[RouterA] interface gigabitethernet 1/0/0
[RouterA-GigabitEthernet1/0/0] ip address 60.1.1.1 255.255.255.0
[RouterA-GigabitEthernet1/0/0] quit
[RouterA] interface gigabitethernet 2/0/0
[RouterA-GigabitEthernet2/0/0] ip address 10.1.1.1 255.255.255.0
[RouterA-GigabitEthernet2/0/0] quit
# Configure a static route to the peer on RouterA. This example assumes that the next hop address in the route to RouterB is 60.1.1.2.
[RouterA] ip route-static 60.1.2.0 255.255.255.0 60.1.1.2
[RouterA] ip route-static 10.1.2.0 255.255.255.0 60.1.1.2
# Assign an IP address to an interface on RouterB. The IP address of GigabitEthernet4/0/0 must be on the same network segment as the IP address assigned by the DHCP server.
<Huawei> system-view
[Huawei] sysname RouterB
[RouterB] interface gigabitethernet 1/0/0
[RouterB-GigabitEthernet1/0/0] ip address 60.1.2.1 255.255.255.0
[RouterB-GigabitEthernet1/0/0] quit
[RouterB] interface gigabitethernet 2/0/0
[RouterB-GigabitEthernet2/0/0] ip address 10.1.2.1 255.255.255.0
[RouterB-GigabitEthernet2/0/0] quit
[RouterB] interface gigabitethernet 3/0/0
[RouterB-GigabitEthernet3/0/0] ip address 10.1.3.1 255.255.255.0
[RouterB-GigabitEthernet3/0/0] quit
[RouterB] interface gigabitethernet 4/0/0
[RouterB-GigabitEthernet4/0/0] ip address 100.1.1.3 255.255.255.0
[RouterB-GigabitEthernet4/0/0] quit
# Configure a static route to the peer on RouterB. This example assumes that the next hop address in the route to RouterA is 60.1.2.2.
[RouterB] ip route-static 60.1.1.0 255.255.255.0 60.1.2.2
[RouterB] ip route-static 10.1.1.0 255.255.255.0 60.1.2.2
[RouterB] ip route-static 100.1.1.0 255.255.255.0 60.1.2.2
# Enable DHCP, create a DHCP server group, and add DHCP servers to the DHCP server group.
[RouterB] dhcp enable
[RouterB] dhcp server group dhcp-ser1
[RouterB-dhcp-server-group-dhcp-ser1] dhcp-server 10.1.3.2
[RouterB-dhcp-server-group-dhcp-ser1] gateway 100.1.1.3
[RouterB-dhcp-server-group-dhcp-ser1] quit
# In the service scheme view, configure the resources to be allocated, including the IP address, DNS domain name, DNS server IP addresses, and WINS server IP addresses.
[RouterB] aaa
[RouterB-aaa] service-scheme schemetest 
[RouterB-aaa-service-schemetest] dhcp-server group dhcp-ser1
[RouterB-aaa-service-schemetest] dns-name mydomain.com.cn
[RouterB-aaa-service-schemetest] dns 2.2.2.2
[RouterB-aaa-service-schemetest] dns 2.2.2.3 secondary
[RouterB-aaa-service-schemetest] wins 3.3.3.2
[RouterB-aaa-service-schemetest] wins 3.3.3.3 secondary
[RouterB-aaa-service-schemetest] quit
[RouterB-aaa] quit
# Configure an IKE proposal and an IKE peer, and bind the service scheme to the IKE peer.
[RouterB] ike proposal 5
[RouterB-ike-proposal-5] dh group14
[RouterB-ike-proposal-5] authentication-algorithm sha2-256
[RouterB-ike-proposal-5] encryption-algorithm aes-128
[RouterB-ike-proposal-5] quit
[RouterB] ike peer rut3
[RouterB-ike-peer-rut3] version 1
[RouterB-ike-peer-rut3] undo version 2
[RouterB-ike-peer-rut3] exchange-mode aggressive
[RouterB-ike-peer-rut3] pre-shared-key cipher YsHsjx_202206
[RouterB-ike-peer-rut3] ike-proposal 5
[RouterB-ike-peer-rut3] service-scheme schemetest
[RouterB-ike-peer-rut3] quit
# Configure an IPSec proposal and establish an IPSec policy using an IPSec policy template.
[RouterB] ipsec proposal prop1
[RouterB-ipsec-proposal-prop1] esp authentication-algorithm sha2-256
[RouterB-ipsec-proposal-prop1] esp encryption-algorithm aes-128
[RouterB-ipsec-proposal-prop1] quit
[RouterB] ipsec policy-template temp1 10
[RouterB-ipsec-policy-templet-temp1-10] ike-peer rut3
[RouterB-ipsec-policy-templet-temp1-10] proposal prop1
[RouterB-ipsec-policy-templet-temp1-10] quit
[RouterB] ipsec policy policy1 10 isakmp template temp1
# Enable SHA-2 to be compatible with RFC standard algorithm versions.
[RouterB] ipsec authentication sha2 compatible enable
# Apply the IPSec policy to an interface.
[RouterB] interface gigabitethernet 1/0/0
[RouterB-GigabitEthernet1/0/0] ipsec policy policy1
[RouterB-GigabitEthernet1/0/0] quit
# Configure an Efficient VPN policy in client mode and specify the remote address and pre-shared key.
[RouterA] ipsec efficient-vpn evpn mode client
[RouterA-ipsec-efficient-vpn-evpn] remote-address 60.1.2.1 v1
[RouterA-ipsec-efficient-vpn-evpn] pre-shared-key cipher YsHsjx_202206
[RouterA-ipsec-efficient-vpn-evpn] dh group14
[RouterA-ipsec-efficient-vpn-evpn] quit
# Apply the Efficient VPN policy to the interface.
[RouterA] interface gigabitethernet 1/0/0 
[RouterA-GigabitEthernet1/0/0] ipsec efficient-vpn evpn
[RouterA-GigabitEthernet1/0/0] quit
# After the configurations are complete, PC A can ping PC B successfully. You can run the display ipsec statistics command to view packet statistics.
# Run the display ike sa command on RouterA. The following information is displayed:
[RouterA] display ike sa
IKE SA information :
    Conn-ID  Peer            VPN   Flag(s)   Phase   RemoteType  RemoteID
  -----------------------------------------------------------------------------
    26       60.1.2.1:500          RD|ST     v1:2    IP          60.1.2.1
    25       60.1.2.1:500          RD|ST     v1:1    IP          60.1.2.1
                                   
  Number of IKE SA : 2 
  -----------------------------------------------------------------------------
                                                           
