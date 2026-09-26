---
id: collect-260926-mikrotik/mikrotik/mikrotik-ipsec-tunnel-with-ddns-and-nat-occursus-arca-1
title: "mikrotik-ipsec-tunnel-with-ddns-and-nat-occursus-arca"
domain: mikrotik
role: reference
task: reference
actors: []
dates: ["2021-02-09"]
keywords: []
source: docs/RAG/lot-mikrotik/forum/ipsec/mikrotik-ipsec-tunnel-with-ddns-and-nat-occursus-arca.md
source_anchor: ""
source_lines: [1, 112]
sha256: 62c6356b544771ea9dd29bd3c10beef12f3ed40b2089dd159032f8263f0d7373
---

# mikrotik-ipsec-tunnel-with-ddns-and-nat-occursus-arca

This guide describes the following situation:

- VPN site-to-site tunnel using IPSec setup is created in MikroTik routers between two private networks: 10.10.10.0/24 and 10.10.20.0/24
- Both private networks use MikroTik router as a gateway
- Each MikroTik router is behind a NAT and have private network range on WAN ports as well: 192.168.10.0/24 and 192.168.20.0/24
- Each MikroTik router has IPSec NAT-Traversal (4500/UDP) forwarded from its gateway (ISP Router)
- Both public network connections change public IP occasionally

Some more remarks:

- I didn’t find any guide which would describe this setup, so I created one.
- Before the start, make sure that you have a separate access to each router, in case you will break your connection.
- Examples contain some additional security settings which can provide better security. Before you use or change these settings, make sure you know what you are doing.
- IPSec tunnel setup in examples uses pre-shared-key authentication method, which has been chosen only for demonstrative purpose and more secure method should be considered. In any case, make sure that if you are going to use PSK method then you need to use a different secret than the one in the examples – also don’t forget that the secret needs to be the same on both sides.
- Names of interfaces on MikroTik routers in this example are:
  - WAN: ether1-gateway
  - LAN: bridge-local
- I successfully tested the setup on 2x Mikrotik hAP lite classic devices, each running behind different routers ( in one case Draytek Vigor 2700 and Ubee EVW3226, in another case TP-Link TD-W8951NB and Compal CH7465LG )
- Also there is a lot of useful documentation about IPSec VPN on MikroTik Wiki – check it out.

Update 1:

- I have experienced tunnel instability when upload link ( provided by ISP ) has been overloaded and when IPSec tunnel was configured with AES GCM. This issue occurred for me at least on Router OS versions 6.38 – 6.39.1. Therefore I have updated the example to use AES CBC, which proved to be stable. But GCM is more secure than CBC, so I recommend to upgrade the RouterOS to the latest version and try with GCM at first.

Update 2:

- RouterOS 6.38 (2016-Dec-30) added IKEv2 support as key exchange mode for IPSec. This mode can be used to improve the security of the tunnel establishment, so I’ve updated the examples in this article accordingly. IKEv2 is also more recent and updated version of the key exchange mode than previously available modes. Except the security improvements, it has embedded the “dead peer detection” and “NAT traversal”, which makes the configuration easier. Additionally, IKEv2 NAT traversal ensures that if connection cannot be created directly between two peers, port 4500/UDP is used. Therefore in RouterOS firewall you need to allow only 4500/UDP.
- RouterOS 6.41 (2017-Dec-22) introduced possibility to use DNS name as IPSec peer address instead IP address. If you will use DNS names in /ip ipsec peer configurations, you can skip the respective part in ipsec-peer-update script mentioned later in this guide, but it is still required to update the IP address of sa-dst-address in IPSec policy, in case the remote router’s public IP has changed. Examples here do not use DNS names in IPSec peer addresses.

Update 3:

- RouterOS 6.43.12 (2019-Feb-11) moved the IPSec profile outside peer configuration and RouterOS 6.44 (2019-Feb-26) added the IPSec identity menu for peers. Therefore I’ve updated the IPSec configurations with versions before RouterOS 6.43.12 and after RouterOS 6.44.

Update 4:

- RouterOS since some version around 6.46 (2019-Dec-02) requires script policy permission “test” for DNS requests using :resolve command. Script and scheduler creation commands have been updated accordingly.

Update 5:

- Refresh on compatibility and features of long-term support version 6.47.9 (2021-02-09). Change to elliptic-curve crypto in some examples, if your non-Mikrotik peer does not connect, try to switch to RSA/DHmodp/smaller curves.

Update 6:

- To simplify the setup with new features available, this guide now:
  - describes only requirements for IKEv2
  - uses no dedicated script for update of peer’s IP address as peers already support hostnames
  - uses no dedicated script for update of IPSec policy sa-dst-address as that is handled now differently by RouterOS

## IPSec

### VPN Tunnel

IPSec tunnel will provide secure site-to-site VPN.

##### MikroTik router 1

/ip ipsec
profile add name="secure-profile" hash-algorithm=sha512 enc-algorithm=aes-256,aes-128 dh-group=ecp521
peer add name="vpn01" comment="vpn01" address="router2.sn.mynetname.net" exchange-mode=ike2 profile=secure-profile
identity add comment="vpn01" auth-method=pre-shared-key secret=REPLACE_THIS_WITH_RANDOM_SECRET peer=vpn01
proposal add name="secure-proposal" auth-algorithms=sha512 enc-algorithms=aes-256-gcm pfs-group=ecp521
policy add comment="vpn01" dst-address=10.10.20.0/24 src-address=10.10.10.0/24 tunnel=yes proposal=secure-proposal peer=vpn01

##### MikroTik router 2

/ip ipsec
profile add name="secure-profile" hash-algorithm=sha512 enc-algorithm=aes-256,aes-128 dh-group=ecp521
peer add name="vpn01" comment="vpn01" address="router1.sn.mynetname.net" exchange-mode=ike2 profile=secure-profile
identity add comment="vpn01" auth-method=pre-shared-key secret=REPLACE_THIS_WITH_RANDOM_SECRET peer=vpn01
proposal add name="secure-proposal" auth-algorithms=sha512 enc-algorithms=aes-256-gcm pfs-group=ecp521 
policy add comment="vpn01" dst-address=10.10.10.0/24 src-address=10.10.20.0/24 tunnel=yes proposal=secure-proposal peer=vpn01

### Firewall

If some rules are used in NAT tables, they need to exclude IPSec traffic, so they will not translate IP addresses in them. This is managed now in default configuration for masquerade rule by ipsec-policy=out,none. In case of manual configuration add this parameter to your masquerade rule. Alternatively you can exclude IPSec traffic by using IPSec accept rule before the NAT rules ( see NAT bypass on MikroTik Wiki for more info ). The same thing applies for a destination NAT ( known sometimes as port forwarding or DMZ host ) – in such case dstnat rules need ipsec-policy=in,none parameter or NAT bypass in opposite direction before them. Here are examples of NAT bypass rules in both directions.

##### MikroTik router 1

/ip firewall
nat add comment="vpn01" action=accept chain=srcnat dst-address=10.10.20.0/24 src-address=10.10.10.0/24 place-before=0
nat add comment="vpn01" action=accept chain=dstnat dst-address=10.10.10.0/24 src-address=10.10.20.0/24 place-before=0

##### MikroTik router 2

/ip firewall
nat add comment="vpn01" action=accept chain=srcnat dst-address=10.10.10.0/24 src-address=10.10.20.0/24 place-before=0
nat add comment="vpn01" action=accept chain=dstnat dst-address=10.10.20.0/24 src-address=10.10.10.0/24 place-before=0

In addition, IPSec IKE traffic needs to be allowed by firewall. For IKEv2 this traffic is 4500/UDP.

##### MikroTik router 1

/ip firewall
filter add comment="ipsec-ike-natt" chain=input dst-port=4500 in-interface=ether1-gateway protocol=udp
filter add comment="vpn01" chain=forward dst-address=10.10.10.0/24 in-interface=ether1-gateway ipsec-policy=in,ipsec src-address=10.10.20.0/24

##### MikroTik router 2

/ip firewall
filter add comment="ipsec-ike-natt" chain=input dst-port=4500 in-interface=ether1-gateway protocol=udp
filter add comment="vpn01" chain=forward dst-address=10.10.20.0/24 in-interface=ether1-gateway ipsec-policy=in,ipsec src-address=10.10.10.0/24

## Public IP Change Adaptation

### IP Cloud

IP Cloud is used as a dynamic DNS system for lookup of remote site’s public IP. This step can be skipped if different DDNS system is used. Time update via IP Cloud is disabled for a case when NTP is used, however you can enable it if necessary.

##### MikroTik router 1 and MikroTik router 2

/ip cloud set ddns-enabled=yes update-time=no

