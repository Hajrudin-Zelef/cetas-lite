---
id: collect-261001-fortinet/fortinet/questions-36569-connecting-to-fortigate-5-4-with-vpnc-ipsec-2deec0a4
title: "questions-36569-connecting-to-fortigate-5-4-with-vpnc-ipsec-2deec0a4"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/questions-36569-connecting-to-fortigate-5-4-with-vpnc-ipsec-2deec0a4.md
source_anchor: ""
source_lines: [1, 55]
sha256: 9ae7da734184375ac9a74951e3d716f979b34ca0cf3b1282f77a08dcb4c1ab69
---

# questions-36569-connecting-to-fortigate-5-4-with-vpnc-ipsec-2deec0a4

I have a fortigate 90b(?) running 5.4. wan1 is set to a publically accessible address. Our internal network is 192.168.100.0
I can connect via ipsec and I can access the gateway, which is the fortigate, at 192.168.100.1, when connected, but I can't traverse the internal network or the internet through the gateway. I gave up on split tunneling until I can figure things out. The VPN is setup as:
   edit "IPSEC_Cisco"
        set type dynamic
        set interface "wan1"
        set mode aggressive
        set peertype one
        set mode-cfg enable
        set proposal aes256-md5 aes256-sha1
        set dhgrp 2
        set wizard-type dialup-cisco
        set xauthtype auto
        set authusrgrp "vpnUsers"
        set peerid "linux"
        set ipv4-start-ip 192.168.100.170
        set ipv4-end-ip 192.168.100.180
        set ipv4-netmask 255.255.255.0
        set dns-mode auto
        set ipv4-split-include "all"
        set psksecret ENC [key]
    next
Phase 2
config vpn ipsec phase2-interface
    edit "IPSEC_Cisco"
        set phase1name "IPSEC_Cisco"
        set proposal aes256-md5 aes256-sha1
        set pfs disable
        set keepalive enable
    next
config firewall policy
    edit 17
        set name "IPSEC_Cisco_Internet"
        set uuid e8bc5412-ab3d-51e6-db87-ba2ff7d3aea2
        set srcintf "IPSEC_Cisco"
        set dstintf "wan1"
        set srcaddr "all"
        set dstaddr "all"
        set action accept
        set schedule "always"
        set service "ALL"
        set nat enable
    next
    edit 18
        set name "IPSEC_CISCO_internal"
        set uuid bbe449ee-ac0b-51e6-de02-0ec32669842e
        set srcintf "IPSEC_Cisco"
        set dstintf "internal"
        set srcaddr "all"
        set dstaddr "all"
        set action accept
        set schedule "always"
        set service "ALL"
        set nat enable
    next
I've tried everything I know of, which isn't much. Any insight appreciated.
