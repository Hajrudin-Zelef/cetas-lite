---
id: collect-261001-huawei/huawei/hedex-api-pages-edoc1100331435-aem10132-05-resources-dc-dc-cfg-ipsec-0063-html-e03cd1d8-1
title: "Assign an IP address to an interface on RouterA."
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/hedex-api-pages-edoc1100331435-aem10132-05-resources-dc-dc-cfg-ipsec-0063-html-e03cd1d8.md
source_anchor: ""
source_lines: [1, 117]
sha256: b98cae084dcbf0c2fa5fb98938747125dca90a438dab6a04a7196c95a7def08b
---

# Assign an IP address to an interface on RouterA.

As shown in Figure 1, RouterA and RouterB communicate through the NAT gateway. RouterA is located on the subnet at 10.1.0.2/24, and RouterB is located on the subnet at 10.2.0.2/24.
The enterprise wants to protect traffic exchanged between RouterA and RouterB.
RouterA and RouterB communicate through the NAT gateway, so NAT traversal must be enabled for establishing an IPSec tunnel. The configuration roadmap is as follows:
Configure IP addresses and static routes for interfaces on RouterA and RouterB so that routes between RouterA and RouterB are reachable.
Configure an ACL on RouterA to define data flows to be protected.
Configure IPSec proposals to define the method used to protect IPSec traffic.
Configure IKE peers to define IKE negotiation attributes.
Configure IPSec policies on RouterA and RouterB. RouterB uses an IPSec policy template to create an IPSec policy.
Apply IPSec policy groups to interfaces.
# Assign an IP address to an interface on RouterA.
<Huawei> system-view
[Huawei] sysname RouterA
[RouterA] interface gigabitethernet 1/0/0
[RouterA-GigabitEthernet1/0/0] ip address 192.168.0.2 255.255.255.0
[RouterA-GigabitEthernet1/0/0] quit
[RouterA] interface gigabitethernet 2/0/0
[RouterA-GigabitEthernet2/0/0] ip address 10.1.0.1 255.255.255.0
[RouterA-GigabitEthernet2/0/0] quit
# Configure a static route to the peer on RouterA. This example assumes that the next-hop address in the route to RouterB is 192.168.0.1.
[RouterA] ip route-static 0.0.0.0 0.0.0.0 192.168.0.1
# Assign an IP address to an interface on RouterB.
<Huawei> system-view
[Huawei] sysname RouterB
[RouterB] interface gigabitethernet 1/0/0 
[RouterB-GigabitEthernet1/0/0] ip address 1.2.0.1 255.255.255.0
[RouterB-GigabitEthernet1/0/0] quit
[RouterB] interface gigabitethernet 2/0/0
[RouterB-GigabitEthernet2/0/0] ip address 10.2.0.1 255.255.255.0
[RouterB-GigabitEthernet2/0/0] quit
# Configure a static route to the peer on RouterB. This example assumes that the next hop address in the route to RouterA is 1.2.0.2.
[RouterB] ip route-static 10.1.0.0 255.255.255.0 1.2.0.2
[RouterB] ip route-static 192.168.0.0 255.255.255.0 1.2.0.2
[RouterA] acl number 3101
[RouterA-acl-adv-3101] rule permit ip source 10.1.0.0 0.0.0.255 destination 10.2.0.0 0.0.0.255
[RouterA-acl-adv-3101] quit
# Create an IPSec proposal on RouterA.
[RouterA] ipsec proposal tran1
[RouterA-ipsec-proposal-tran1] esp authentication-algorithm sha2-256
[RouterA-ipsec-proposal-tran1] esp encryption-algorithm aes-128 
[RouterA-ipsec-proposal-tran1] quit
# Create an IPSec proposal on RouterB.
[RouterB] ipsec proposal tran1
[RouterB-ipsec-proposal-tran1] esp authentication-algorithm sha2-256
[RouterB-ipsec-proposal-tran1] esp encryption-algorithm aes-128 
[RouterB-ipsec-proposal-tran1] quit
# Set the local ID type to name on RouterA.
[RouterA] ike local-name rta
# Set the local ID type to name on RouterB.
[RouterB] ike local-name rtb
# Create an IKE proposal on RouterA.
[RouterA] ike proposal 5
[RouterA-ike-proposal-5] encryption-algorithm aes-128
[RouterA-ike-proposal-5] authentication-algorithm sha2-256
[RouterA-ike-proposal-5] dh group14
[RouterA-ike-proposal-5] quit
# Configure an IKE peer on RouterA.
[RouterA] ike peer rta
[RouterA-ike-peer-rta] version 1
[RouterA-ike-peer-rta] undo version 2
[RouterA-ike-peer-rta] exchange-mode aggressive 
[RouterA-ike-peer-rta] ike-proposal 5
[RouterA-ike-peer-rta] pre-shared-key cipher YsHsjx_202206
[RouterA-ike-peer-rta] local-id-type fqdn
[RouterA-ike-peer-rta] remote-address 1.2.0.1
[RouterA-ike-peer-rta] remote-id rtb
[RouterA-ike-peer-rta] nat traversal
[RouterA-ike-peer-rta] quit
# Create an IKE proposal on RouterB.
[RouterB] ike proposal 5
[RouterB-ike-proposal-5] encryption-algorithm aes-128
[RouterB-ike-proposal-5] authentication-algorithm sha2-256
[RouterB-ike-proposal-5] dh group14
[RouterB-ike-proposal-5] quit
# Configure an IKE peer on RouterB.
[RouterB] ike peer rtb
[RouterB-ike-peer-rta] version 1
[RouterB-ike-peer-rta] undo version 2
[RouterB-ike-peer-rtb] exchange-mode aggressive 
[RouterB-ike-peer-rtb] ike-proposal 5
[RouterB-ike-peer-rtb] pre-shared-key cipher YsHsjx_202206
[RouterB-ike-peer-rtb] local-id-type fqdn
[RouterB-ike-peer-rtb] remote-id rta
[RouterB-ike-peer-rtb] nat traversal
[RouterB-ike-peer-rtb] quit
# Create an IPSec policy for IKE negotiation on RouterA.
[RouterA] ipsec policy policy1 10 isakmp
[RouterA-ipsec-policy-isakmp-policy1-10] security acl 3101
[RouterA-ipsec-policy-isakmp-policy1-10] ike-peer rta
[RouterA-ipsec-policy-isakmp-policy1-10] proposal tran1
[RouterA-ipsec-policy-isakmp-policy1-10] quit
# Create an IPSec policy for IKE negotiation using a policy template on RouterB.
[RouterB] ipsec policy-template temp1 10
[RouterB-ipsec-policy-templet-temp1-10] ike-peer rtb
[RouterB-ipsec-policy-templet-temp1-10] proposal tran1
[RouterB-ipsec-policy-templet-temp1-10] quit
[RouterB] ipsec policy policy1 10 isakmp template temp1
Run the display ipsec policy command on RouterA and RouterB to view the configurations of the IPSec policies.
# Apply the IPSec policy group to the interface of RouterA
[RouterA] interface gigabitethernet 1/0/0
[RouterA-GigabitEthernet1/0/0] ipsec policy policy1
[RouterA-GigabitEthernet1/0/0] quit
# Apply the IPSec policy group to the interface of RouterB.
[RouterB] interface gigabitethernet 1/0/0
[RouterB-GigabitEthernet1/0/0] ipsec policy policy1
[RouterB-GigabitEthernet1/0/0] quit
# After the configurations are complete, PC A can ping PC B successfully. Data exchanged between PC A and PC B is encrypted. You can run the display ipsec statistics command to view packet statistics.
# Run the display ike sa command on RouterA. The following information is displayed:
[RouterA] display ike sa
IKE SA information :
  Conn-ID  Peer            VPN   Flag(s)   Phase   RemoteType  RemoteID
  ---------------------------------------------------------------------------
     15    1.2.0.1:4500          RD|ST     v1:2    FQDN        rtb
     14    1.2.0.1:4500          RD|ST     v1:1    FQDN        rtb
                                   
  Number of IKE SA : 2 
  ---------------------------------------------------------------------------
                                                           
