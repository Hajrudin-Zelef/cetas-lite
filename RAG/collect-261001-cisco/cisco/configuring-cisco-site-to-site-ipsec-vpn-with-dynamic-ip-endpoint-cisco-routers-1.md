---
id: collect-261001-cisco/cisco/configuring-cisco-site-to-site-ipsec-vpn-with-dynamic-ip-endpoint-cisco-routers-1
title: "configuring-cisco-site-to-site-ipsec-vpn-with-dynamic-ip-endpoint-cisco-routers"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/configuring-cisco-site-to-site-ipsec-vpn-with-dynamic-ip-endpoint-cisco-routers.md
source_anchor: ""
source_lines: [1, 156]
sha256: ca3ae359eedc5754048de6b66ea54f11998b8da8df7b2ee3ddd8a21eeb5309b4
---

# configuring-cisco-site-to-site-ipsec-vpn-with-dynamic-ip-endpoint-cisco-routers

Configuring Cisco Site to Site IPSec VPN with Dynamic IP Endpoint Cisco Routers

This article serves as an extension to our popular Cisco VPN topics covered here on Firewall.cx. While we’ve covered Site to Site IPSec VPN Tunnel Between Cisco Routers (using static public IP addresses), we will now take a look on how to configure our headquarter Cisco router to support remote Cisco routers with dynamic IP addresses. One important note to keep in mind when it comes to this implementation, is that Site-to-Site VPN networks with Dynamic remote Public IP addresses can only be brought up by the remote site routers as only they are aware of the headquarter's router Public IP address.

ISAKMP (Internet Security Association and Key Management Protocol) and IPSec are essential to building and encrypting the VPN tunnel. ISAKMP, also called IKE (Internet Key Exchange), is the negotiation protocol that allows two hosts to agree on how to build an IPsec security association. ISAKMP negotiation consists of two phases: Phase 1 and Phase 2.

Phase 1 creates the first tunnel, which protects later ISAKMP negotiation messages. Phase 2 creates the tunnel that protects data. IPSec then comes into play to encrypt the data using encryption algorithms and provides authentication, encryption and anti-replay services.

IPSec VPN Requirements

To help make this an easy-to-follow exercise, we have split it into two required steps to get the Site-to-Site IPSec Dynamic IP Endpoint VPN Tunnel to work.

Our example setup consists of the headquarter router R1 which is assigned a static public IP address, and two remote routers, R2 & R3. Both remote routers (R2 & R3) connect to the Internet and have a dynamic public IP address assigned by the ISP, as shown in the diagram below:

Our Headquarters is assigned an internal network of 10.10.10.0/24, while Remote Site 1 has been assigned network 20.20.20.0/24. and Remote Site 2 network 30.30.30.0/24. The goal is to securely connect both remote sites with our headquarters and allow full communication, without any restrictions.

Configure ISAKMP (IKE) - (ISAKMP Phase 1)

IKE exists only to establish SAs (Security Association) for IPsec. Before it can do this, IKE must negotiate an SA (an ISAKMP SA) relationship with the peer.

To begin, we’ll start working on the Headquarter router (R1).

First step is to configure an ISAKMP Phase 1 policy:

crypto isakmp policy 1

encr 3des

hash md5

authentication pre-share

group 2

lifetime 86400

The above commands define the following (in listed order):

3DES - The encryption method to be used for Phase 1.

MD5 - The hashing algorithm

Pre-share - Use Pre-shared key as the authentication method

Group 2 - Diffie-Hellman group to be used

86400 – Session key lifetime. Expressed in either kilobytes (after x-amount of traffic, change the key) or seconds. Value set is the default value.

We should note that ISAKMP Phase 1 policy is defined globally. This means that if we have five different remote sites and configured five different ISAKMP Phase 1 policies (one for each remote router), when our router tries to negotiate a VPN tunnel with each site it will send all five policies and use the first match that is accepted by both ends. Since we only have one ISAKMP policy, this will be used for all remote VPN routers.

Next we are going to define a pre-shared key for authentication with our peers (R2 & R3 routers) by using the following command:

The peers pre-shared key is set to firewallcx and note that we are defining a remote public IP address of 0.0.0.0 0.0.0.0. This tells our headquarter router that the remote routers have dynamic public IP addresses and ensures it will try to negotiate and establish a VPN tunnel with any router that requests it.

Configure IPSec

To configure IPSec we need to setup the following in order:

- Create extended ACL

- Create IPSec Transform

- Create Dynamic Crypto Maps

- Apply crypto map to the public interface

Let us examine each of the above steps.

Creating Extended ACL

Next step is to create an access-list and define the traffic we would like the router to pass through each VPN tunnel. In this example, for the first VPN tunnel it would be traffic from headquarters (10.10.10.0/24) to remote site 1 (20.20.20.0/24) and for the second VPN tunnel it will be from our headquarters (10.10.10.0/24) to remote site 2 (30.30.30.0/24). Access-lists that define VPN traffic are sometimes called crypto access-list or interesting traffic access-list.

Because we are dealing with two separate VPN tunnels, we’ll need to create one set of access-lists for each:

ip access-list extended VPN1-TRAFFIC permit ip 10.10.10.0 0.0.0.255 20.20.20.0 0.0.0.255 ! ip access-list extended VPN2-TRAFFIC

permit ip 10.10.10.0 0.0.0.255 30.30.30.0 0.0.0.255

Create IPSec Transform (ISAKMP Phase 2 policy)

Now we need to create the transform set used to protect our data. We’ve named our transform set TS:

The Crypto Map is the last step of our setup and connects the previously defined ISAKMP and IPSec configuration together. We will need one dynamic crypto map for each remote endpoint, which means a total of two crypto maps for our setup.

First we create a crypto map named VPN which will be applied to the public interface of our headquarter router, and connect it with the dynamic crypto maps we named as hq-vpn.

crypto map VPN 1 ipsec-isakmp dynamic hq-vpn

The ipsec-isakmp tag tells the router that this crypto map is an IPsec crypto map. Now we create our two dynamic crypto maps using the following configuration commands:

crypto dynamic-map hq-vpn 10

set security-association lifetime seconds 86400

set transform-set TS

match address VPN1-TRAFFIC

!

crypto dynamic-map hq-vpn 11

set security-association lifetime seconds 86400

set transform-set TS

match address VPN2-TRAFFIC

Notice how we create one dynamic map for each remote network. The configuration is similar for each dynamic crypto map, with only the instance number (10 , 11) and match address (VPN1-TRAFFIC , VPN2-TRAFFIC) changing.

Adding additional remote sites in the future is as easy as simply adding more dynamic crypto maps, incrementing the index number and specifying the match address extended access-lists for each remote network.

Apply Crypto Map To The Public Interface

The final step is to apply our crypto map to the public interface of the headquarter router, which is FastEthernet0/1. In many cases, this might be a serial or ATM (ADSL - Dialer) interface:

interface FastEthernet0/1 crypto map VPN

Note that you can assign only one crypto map to an interface.

As soon as we apply crypto map on the interface, we receive a message from the router that confirms isakmp is on: “ISAKMP is ON”.

At this point, we have completed the IPSec VPN configuration on our headquarter router and we can move to the remote endpoint routers.

Configuring Remote Endpoint Routers (Dynamic Public IP Addresses)

Our remote routers connect to the Internet and are assigned a dynamic IP address which changes periodically by the ISP. In most part, the configuration is similar to that of the headquarter router, but with a few minor changes.

In the configuration below, IP address 74.200.90.5 represents the public IP address of our headquarter router.

Remote Site 1 Router

crypto isakmp policy 1 encr 3des

hash md5

authentication pre-share

group 2

lifetime 86400

!

crypto isakmp key firewallcx address 74.200.90.5

!

ip access-list extended VPN-TRAFFIC

permit ip 20.20.20.0 0.0.0.255 10.10.10.0 0.0.0.255

Network Address Translation (NAT) is most likely to be configured to provide Internet access to internal hosts. When configuring a Site-to-Site VPN tunnel, it is imperative to instruct the router not to perform NAT (deny NAT) on packets destined to the remote VPN networks.

This is easily done by inserting a deny statement at the beginning of the NAT access lists as shown below:

