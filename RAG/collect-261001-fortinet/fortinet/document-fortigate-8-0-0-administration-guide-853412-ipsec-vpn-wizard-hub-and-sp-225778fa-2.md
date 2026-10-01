---
id: collect-261001-fortinet/fortinet/document-fortigate-8-0-0-administration-guide-853412-ipsec-vpn-wizard-hub-and-sp-225778fa-2
title: "IPsec VPN wizard hub-and-spoke ADVPN support"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/document-fortigate-8-0-0-administration-guide-853412-ipsec-vpn-wizard-hub-and-sp-225778fa.md
source_anchor: ""
source_lines: [142, 191]
sha256: b05067461764b98ed866968bada795f306855f6e7096f6f21c391051d3403d12
---

# IPsec VPN wizard hub-and-spoke ADVPN support

```
config vpn ipsec phase1-interface
    edit "ADVPN-S1"
        set interface "port2"
        set ike-version 2
        set peertype any
        set net-device enable
        set proposal aes128-sha256 aes256-sha256 aes128gcm-prfsha256 aes256gcm-prfsha384 chacha20poly1305-prfsha256
        set add-route disable
        set dpd on-idle
        set dhgrp 20 21
        set wizard-type spoke-fortigate-auto-discovery
        set auto-discovery-receiver enable
        set transport auto
        set remote-gw 203.0.113.249
        set psksecret ENC \<key\>
    next
end
config vpn ipsec phase2-interface
    edit "ADVPN-S1"
        set phase1name "ADVPN-S1"
        set proposal aes128-sha256 aes256-sha256 aes128gcm aes256gcm chacha20poly1305
        set dhgrp 20 21
    next
end
```
```
config firewall policy
    edit 5
        set name "vpn_ADVPN-S1_remote"
        set srcintf "ADVPN-S1"
        set dstintf "port3"
        set action accept
        set srcaddr "all"
        set dstaddr "ADVPN-S1_local"
        set schedule "always"
        set service "ALL"
    next
    edit 6
        set name "vpn_ADVPN-S1_local"
        set srcintf "port3"
        set dstintf "ADVPN-S1"
        set action accept
        set srcaddr "all"
        set dstaddr "all"
        set schedule "always"
        set service "ALL"
    next
end
```
