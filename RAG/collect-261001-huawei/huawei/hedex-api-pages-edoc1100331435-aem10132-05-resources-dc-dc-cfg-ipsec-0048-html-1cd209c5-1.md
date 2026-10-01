---
id: collect-261001-huawei/huawei/hedex-api-pages-edoc1100331435-aem10132-05-resources-dc-dc-cfg-ipsec-0048-html-1cd209c5-1
title: "Assign an IP address to an interface on RouterA."
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/hedex-api-pages-edoc1100331435-aem10132-05-resources-dc-dc-cfg-ipsec-0048-html-1cd209c5.md
source_anchor: ""
source_lines: [1, 121]
sha256: 0f9550ebd688aa0c88bf78857f45fcd5e87927f85881c070cd28e083877af1a2
---

# Assign an IP address to an interface on RouterA.

As shown in Figure 1, RouterA (branch gateway) and RouterB (headquarters gateway) communicate through the Internet. RouterA uses two egress links in backup or load balancing mode. The branch subnet is 10.1.1.0/24 and the headquarters subnet is 10.1.2.0/24.
The Enterprise wants to protect traffic between the branch subnet and headquarters subnet. If an active/standby switchover occurs or the egress link becomes faulty, IPSec services need to be smoothly switched. IPSec tunnels can be set up between the branch gateway and headquarters gateway because they communicate over the Internet. The two outbound interfaces negotiate with their peers to establish IPSec SAs respectively. When one interface alternates between Up and Down states and an active/standby switchover occurs, the two peers need to perform IKE negotiate again to generate IPSec SAs. The IKE re-negotiation causes IPSec service interruption in a short time. To ensure that IPSec services are smoothly switched, the two egress links on the branch gateway and the headquarters gateway only negotiate a shared IPSec SA.
The branch gateway uses a loopback interface to establish an IPSec tunnel with the headquarters gateway, and the two egress links and the headquarters gateway only negotiate a shared IPSec SA. The configuration roadmap is as follows:
Configure IP addresses and static routes for interfaces on RouterA and RouterB so that routes between RouterA and RouterB are reachable.
Configure ACLs to define data flows to be protected.
Configure IPSec proposals to define the method used to protect IPSec traffic.
Configure IKE peers to define IKE negotiation attributes.
Configure IPSec policies and reference ACLs and IPSec proposals in the IPSec policies to determine the methods used to protect data flows.
Apply IPSec policy groups to interfaces. Configure a multi-link shared IPSec policy group on RouterA so that the IPSec policy group can be shared by multiple interfaces.
# Assign an IP address to an interface on RouterA.
<Huawei> system-view
[Huawei] sysname RouterA
[RouterA] interface gigabitethernet 1/0/0
[RouterA-GigabitEthernet1/0/0] ip address 70.1.1.1 255.255.255.0
[RouterA-GigabitEthernet1/0/0] quit
[RouterA] interface gigabitethernet 2/0/0
[RouterA-GigabitEthernet2/0/0] ip address 80.1.1.1 255.255.255.0
[RouterA-GigabitEthernet2/0/0] quit
[RouterA] interface gigabitethernet 3/0/0 [RouterA-GigabitEthernet3/0/0] ip address 10.1.1.1 255.255.255.0 [RouterA-GigabitEthernet3/0/0] quit
[RouterA] interface loopback 0
[RouterA-LoopBack0] ip address 1.1.1.1 255.255.255.255
[RouterA-LoopBack0] quit
# Configure a static route to the peer on RouterA. This example assumes that the next hop addresses corresponding to the two outbound interfaces in the route to RouterB are 70.1.1.2 and 80.1.1.2.
[RouterA] ip route-static 10.1.2.0 255.255.255.0 70.1.1.2 preference 10
[RouterA] ip route-static 10.1.2.0 255.255.255.0 80.1.1.2 preference 20
[RouterA] ip route-static 60.1.1.0 255.255.255.0 70.1.1.2 preference 10
[RouterA] ip route-static 60.1.1.0 255.255.255.0 80.1.1.2 preference 20
# Assign an IP address to an interface on RouterB.
<Huawei> system-view
[Huawei] sysname RouterB
[RouterB] interface gigabitethernet 1/0/0
[RouterB-GigabitEthernet1/0/0] ip address 60.1.1.1 255.255.255.0
[RouterB-GigabitEthernet1/0/0] quit
[RouterB] interface gigabitethernet 3/0/0
[RouterB-GigabitEthernet3/0/0] ip address 10.1.2.1 255.255.255.0
[RouterB-GigabitEthernet3/0/0] quit
# Configure a static route to the peer on RouterB. This example assumes that the next hop address in the route to RouterA is 60.1.1.2.
[RouterB] ip route-static 1.1.1.1 255.255.255.255 60.1.1.2
[RouterB] ip route-static 10.1.1.0 255.255.255.0 60.1.1.2
[RouterB] ip route-static 70.1.1.0 255.255.255.0 60.1.1.2
[RouterB] ip route-static 80.1.1.0 255.255.255.0 60.1.1.2
# Configure an ACL on RouterA to define data flows sent from 10.1.1.0/24 to 10.1.2.0/24.
[RouterA] acl number 3101
[RouterA-acl-adv-3101] rule permit ip source 10.1.1.0 0.0.0.255 destination 10.1.2.0 0.0.0.255
[RouterA-acl-adv-3101] quit
# Configure an ACL on RouterB to define data flows sent from 10.1.2.0/24 to 10.1.1.0/24.
[RouterB] acl number 3101
[RouterB-acl-adv-3101] rule permit ip source 10.1.2.0 0.0.0.255 destination 10.1.1.0 0.0.0.255
[RouterB-acl-adv-3101] quit
# Create an IPSec proposal on RouterA.
[RouterA] ipsec proposal prop
[RouterA-ipsec-proposal-prop] esp authentication-algorithm sha2-256
[RouterA-ipsec-proposal-prop] esp encryption-algorithm aes-128
[RouterA-ipsec-proposal-prop] quit
# Create an IPSec proposal on RouterB.
[RouterB] ipsec proposal prop
[RouterB-ipsec-proposal-prop] esp authentication-algorithm sha2-256
[RouterB-ipsec-proposal-prop] esp encryption-algorithm aes-128
[RouterB-ipsec-proposal-prop] quit
[RouterA] ike proposal 5
[RouterA-ike-proposal-5] encryption-algorithm aes-128
[RouterA-ike-proposal-5] authentication-algorithm sha2-256
[RouterA-ike-proposal-5] dh group14
[RouterA-ike-proposal-5] quit
# Create an IKE proposal on RouterB.
[RouterB] ike proposal 5
[RouterB-ike-proposal-5] encryption-algorithm aes-128
[RouterB-ike-proposal-5] authentication-algorithm sha2-256
[RouterB-ike-proposal-5] dh group14
[RouterB-ike-proposal-5] quit
# Configure an IKE peer on RouterA, reference the IKE proposal, and set the pre-shared key and remote ID.
[RouterA] ike peer rut
[RouterA-ike-peer-rut] version 1
[RouterA-ike-peer-rut] undo version 2
[RouterA-ike-peer-rut] ike-proposal 5
[RouterA-ike-peer-rut] pre-shared-key cipher YsHsjx_202206
[RouterA-ike-peer-rut] remote-address 60.1.1.1
[RouterA-ike-peer-rut] quit
# Configure an IKE peer on RouterB, reference the IKE proposal, and set the pre-shared key and remote ID.
[RouterB] ike peer rut
[RouterB-ike-peer-rut] version 1
[RouterB-ike-peer-rut] undo version 2
[RouterB-ike-peer-rut] ike-proposal 5
[RouterB-ike-peer-rut] pre-shared-key cipher YsHsjx_202206
[RouterB-ike-peer-rut] remote-address 1.1.1.1
[RouterB-ike-peer-rut] quit
# Create an IPSec policy on RouterA.
[RouterA] ipsec policy policy1 10 isakmp
[RouterA-ipsec-policy-isakmp-policy1-10] ike-peer rut
[RouterA-ipsec-policy-isakmp-policy1-10] proposal prop
[RouterA-ipsec-policy-isakmp-policy1-10] security acl 3101
[RouterA-ipsec-policy-isakmp-policy1-10] quit
# Create an IPSec policy on RouterB.
[RouterB] ipsec policy policy1 10 isakmp
[RouterB-ipsec-policy-isakmp-policy1-10] ike-peer rut
[RouterB-ipsec-policy-isakmp-policy1-10] proposal prop
[RouterB-ipsec-policy-isakmp-policy1-10] security acl 3101
[RouterB-ipsec-policy-isakmp-policy1-10] quit
[RouterA] ipsec policy policy1 shared local-interface loopback 0
[RouterA] interface gigabitethernet 1/0/0
[RouterA-GigabitEthernet1/0/0] ipsec policy policy1
[RouterA-GigabitEthernet1/0/0] quit
[RouterA] interface gigabitethernet 2/0/0
[RouterA-GigabitEthernet2/0/0] ipsec policy policy1
[RouterA-GigabitEthernet2/0/0] quit
# Apply the IPSec policy group to the interface of RouterB.
[RouterB] interface gigabitethernet 1/0/0
[RouterB-GigabitEthernet1/0/0] ipsec policy policy1
[RouterB-GigabitEthernet1/0/0] quit
# After the configurations are complete, PC A can ping PC B successfully. Data exchanged between PC A and PC B is encrypted. You can run the display ipsec statistics command to view packet statistics.
# Run the display ike sa command on RouterA. The following information is displayed:
[RouterA] display ike sa
IKE SA information :
  Conn-ID  Peer          VPN   Flag(s)   Phase   RemoteType  RemoteID
  --------------------------------------------------------------------------
   937    60.1.1.1:500         RD|ST     v1:2    IP          60.1.1.1
   936    60.1.1.1:500         RD|ST     v1:1    IP          60.1.1.1
                                   
  Number of IKE SA : 2
  --------------------------------------------------------------------------
                                                           
