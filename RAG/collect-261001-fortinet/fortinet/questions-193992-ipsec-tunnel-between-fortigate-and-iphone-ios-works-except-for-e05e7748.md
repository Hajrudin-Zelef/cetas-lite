---
id: collect-261001-fortinet/fortinet/questions-193992-ipsec-tunnel-between-fortigate-and-iphone-ios-works-except-for-e05e7748
title: "questions-193992-ipsec-tunnel-between-fortigate-and-iphone-ios-works-except-for--e05e7748"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/questions-193992-ipsec-tunnel-between-fortigate-and-iphone-ios-works-except-for--e05e7748.md
source_anchor: ""
source_lines: [1, 23]
sha256: f1d1567b4c3693a881aa2c882e693cfb2e1cbe211f002b531349013ac3ca1ebc
---

# questions-193992-ipsec-tunnel-between-fortigate-and-iphone-ios-works-except-for--e05e7748

I have configured my Fortigate with a new VPN IPSec tunnel to allow the iOS Cisco client to connect. That works fine. I can RDP to my servers, browse to my servers via IP address, etc.
But, the iPhone does not resolve my internal IP addresses. I have added the DNS servers that serve addresses for my internal users, as well as the WINS servers, but the iPhone acts like it doesn't see them at all.
config vpn ipsec phase1-interface
    edit "iPhone_VPN"
        set type dynamic
        set interface "wan1"
        set dhgrp 2
        set proposal 3des-sha1 3des-md5
        set xauthtype auto
        set mode-cfg enable
        set authusrgrp "iPhone_VPN_Users"
        set ipv4-start-ip 10.10.99.100
        set ipv4-end-ip 10.10.99.199
        set ipv4-netmask 255.255.0.0
        set ipv4-dns-server1 10.10.2.1
        set ipv4-dns-server2 10.22.1.80
        set ipv4-wins-server1 10.10.2.1
        set ipv4-wins-server2 10.22.1.80
        set ipv4-split-include "Dialup_VPN_Networks"
        set psksecret ENC xxxxx
    next
end
For whatever reason, as far as I can tell, the iPhone does not 'see' the DNS or WINS entries. I don't know how to check it...
