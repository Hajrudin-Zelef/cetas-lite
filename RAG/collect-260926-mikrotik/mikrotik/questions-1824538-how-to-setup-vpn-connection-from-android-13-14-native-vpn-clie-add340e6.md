---
id: collect-260926-mikrotik/mikrotik/questions-1824538-how-to-setup-vpn-connection-from-android-13-14-native-vpn-clie-add340e6
title: "2024-06-16 21:14:19 by RouterOS 7.13.2"
domain: mikrotik
role: reference
task: reference
actors: []
dates: ["2024-06-16"]
keywords: []
source: docs/RAG/lot-mikrotik/forum/vpn/questions-1824538-how-to-setup-vpn-connection-from-android-13-14-native-vpn-clie-add340e6.md
source_anchor: ""
source_lines: [1, 35]
sha256: f21aab1184ca6851fd474ae3a989675f2f1dc49405022826b37834fe225d657d
---

# 2024-06-16 21:14:19 by RouterOS 7.13.2

Here is a working configuration of ipsec ikev2 / psk vpn which works WITHOUT CERTIFICATES etc.:
notes:
1.this configuration is NOT touching the "default" profile, "default" identity etc. So it should work in parallel with other VPN types, for instance in paralell with L2TP/ipsec VPN which is creating dynamic identity/peer and cannot use anything else than default. So this configuration is glued together by a group named "ike2-group"
- Android still claims this VPN as "insecure" however I did not dig deeper, I wanted to just "make it work" because L2TP was removed. And I could not really play with certificates etc. and it is supposed to work paralelly with existing VPN configurations.
- You need to alter below scripts a bit, by filling in the [TEXT IN BRACKETS] with your names/passwords etc.
- you need to create address pool for the VPN connections first, and give the pool's name as  [ADDRESS_POOL] below
- [FULL_DOMAIN_NAME_OF_ROUTER] is DNS name under which router will be available (like www.google.com)
- [SECRET] is your pre-shared key.
- IMPORTANT!!!! In Android you have to give such VPN settings:
"name" whatever you like.
"type" is "IKEv2/IPSec PSK"
"Server address" the same as in  [FULL_DOMAIN_NAME_OF_ROUTER] 
"IPsec identifier" the same as in  [FULL_DOMAIN_NAME_OF_ROUTER] 
"Pre shared key" the same as in [SECRET]
Especially please note the "IPSec identifier".
- Maybe proposal could be simplified. I was adding everything till it started to work.
Here is the configuration code:
# 2024-06-16 21:14:19 by RouterOS 7.13.2
# model = RB3011UiAS
/ip ipsec policy group
add name=ike2-group
/ip ipsec mode-config
add address-pool=[ADDRESS_POOL] name=ike2-config
/ip ipsec profile
add dh-group=ecp256,ecp384,ecp521,modp8192,modp6144,modp4096,modp3072,modp2048 enc-algorithm=aes-256,aes-192,aes-128 hash-algorithm=sha512 name=ike2-profile proposal-check=claim
/ip ipsec peer
add exchange-mode=ike2 name=ike2-peer passive=yes profile=ike2-profile secret=[SECRET]
/ip ipsec proposal
add auth-algorithms=sha512,sha256 enc-algorithms=aes-256-cbc,aes-256-ctr,aes-256-gcm,aes-192-cbc,aes-192-ctr,aes-192-gcm,aes-128-cbc,aes-128-ctr,aes-128-gcm name=ike2-proposal pfs-group=\
    modp4096
/ip ipsec identity
add comment="identity to be used in ikev2" generate-policy=port-strict mode-config=ike2-config my-id=fqdn:[FULL_DOMAIN_NAME_OF_ROUTER]\
  peer=ike2-peer policy-template-group=ike2-group
/ip ipsec policy
add comment="policy to be used in ike2-identity and ike2-policy" dst-address=0.0.0.0/0 group=ike2-group proposal=ike2-proposal src-address=0.0.0.0/0 template=yes
