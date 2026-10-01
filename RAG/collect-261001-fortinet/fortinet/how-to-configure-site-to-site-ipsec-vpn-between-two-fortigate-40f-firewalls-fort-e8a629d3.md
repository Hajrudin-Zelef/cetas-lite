---
id: collect-261001-fortinet/fortinet/how-to-configure-site-to-site-ipsec-vpn-between-two-fortigate-40f-firewalls-fort-e8a629d3
title: "How to Configure Site-to-Site IPsec VPN Between Two FortiGate 40F Firewalls (FortiOS v6 and v7)"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["parameters"]
source: docs/RAG/collect-261001-fortinet/how-to-configure-site-to-site-ipsec-vpn-between-two-fortigate-40f-firewalls-fort-e8a629d3.md
source_anchor: ""
source_lines: [1, 158]
sha256: b096256f4d15684f154ed5bf45ad5fc18ccaf1987ecf951f76717096b412b631
---

# How to Configure Site-to-Site IPsec VPN Between Two FortiGate 40F Firewalls (FortiOS v6 and v7)

*Source : https://gkhan.in/how-to-configure-site-to-site-ipsec-vpn-between-two-fortigate-40f-firewalls-fortios-v6-and-v7/*

Connecting branch offices or partners through a secure VPN tunnel is one of the most common FortiGate tasks.
This guide shows how to configure a manual IPsec site-to-site VPN between two FortiGate 40F units running different firmware versions:
| Site | Model | Firmware | 
|---|---|---|
| Site B | FortiGate-40F | v7.6.3, build 3510 (GA.F) | 
| Site C | FortiGate-40F | v6.4.7, build 8726 (GA) | 
| Parameter | Site B | Site C | 
|---|---|---|
| WAN IP | 172.16.10.109 | 172.16.10.111 | 
| LAN Subnet | 192.168.20.0/24 | 192.168.30.0/24 | 
| Tunnel Name | vpn-to-SiteC | SiteB | 
| Pre-Shared Key | Forti@123 | Forti@123 | 
config vpn ipsec phase1-interface
    edit "vpn-to-siteC"
        set interface "wan"                 # WAN interface carrying IKE
        set ike-version 2                   # IKEv2 (default in v7)
        set peertype any
        set net-device disable
        set proposal aes256-sha256          # Encryption / authentication
        set dhgrp 14                        # Diffie-Hellman group
        set transport auto                  # <— New in FortiOS 7.x, auto-selects UDP 500/4500
        set remote-gw 172.16.10.111         # Peer’s WAN IP
        set psksecret "Forti@123"           # Shared key
    next
end
config vpn ipsec phase1-interface
    edit "SiteB"
        set interface "wan"
        set ike-version 2
        set peertype any
        set net-device disable
        set proposal aes256-sha256
        set dhgrp 14
        # ❌ 'set transport' not available in v6.x
        set remote-gw 172.16.10.109
        set psksecret "Forti@123"
    next
end
📝 Explanation:
ike-version 2 → modern and more stable.dhgrp must match on both sides.transport (v7 only) automatically handles NAT-Traversal; in v6 FortiGate does it by default when needed.config vpn ipsec phase2-interface
    edit "Site_C"
        set phase1name "vpn-to-siteC"
        set proposal aes256-sha256
        set src-subnet 192.168.20.0 255.255.255.0
        set dst-subnet 192.168.30.0 255.255.255.0
    next
end
config vpn ipsec phase2-interface
    edit "Site_B"
        set phase1name "SiteB"
        set proposal aes256-sha256
        set src-subnet 192.168.30.0 255.255.255.0
        set dst-subnet 192.168.20.0 255.255.255.0
    next
end
Both sides need two policies:
config firewall policy
    edit 0
        set name "LAN-to-SiteC"
        set srcintf "lan"
        set dstintf "vpn-to-siteC"
        set srcaddr "all"
        set dstaddr "all"
        set action accept
        set schedule "always"
        set service "ALL"
        set nat disable
    next
    edit 0
        set name "SiteC-to-LAN"
        set srcintf "vpn-to-siteC"
        set dstintf "lan"
        set srcaddr "all"
        set dstaddr "all"
        set action accept
        set schedule "always"
        set service "ALL"
        set nat disable
    next
end
config firewall policy
    edit 0
        set name "LAN-to-SiteB"
        set srcintf "lan"
        set dstintf "SiteB"
        set srcaddr "all"
        set dstaddr "all"
        set action accept
        set schedule "always"
        set service "ALL"
        set nat disable
    next
    edit 0
        set name "SiteB-to-LAN"
        set srcintf "SiteB"
        set dstintf "lan"
        set srcaddr "all"
        set dstaddr "all"
        set action accept
        set schedule "always"
        set service "ALL"
        set nat disable
    next
end
Each site needs a route to the opposite LAN through the VPN interface.
config router static
    edit 0
        set dst 192.168.30.0 255.255.255.0
        set device "vpn-to-siteC"
    next
end
config router static
    edit 0
        set dst 192.168.20.0 255.255.255.0
        set device "SiteB"
    next
end
get vpn ipsec tunnel summary
diagnose vpn tunnel list
diagnose vpn tunnel up name vpn-to-siteC
execute ping-options source 192.168.20.1
execute ping 192.168.30.1
diagnose debug reset
diagnose vpn ike log-filter clear
diagnose vpn ike log-filter dst-addr4 172.16.10.111
diagnose debug application ike -1
diagnose debug enable
Look for:
no proposal chosen → mismatch in Phase 1/2 encryption.AUTHENTICATION_FAILED → wrong pre-shared key.invalid id information → subnet mismatch.
Stop debug:
diagnose debug disable
config vpn ipsec phase1-interface
    edit "vpn-to-siteC"
        set dpd on-idle
        set keylife 28800
        set keepalive enable
    next
end
| Command | v7.x | v6.x | Description | 
|---|---|---|---|
| set transport auto | ✅ | ❌ | Controls NAT-T / UDP encapsulation (v7+) | 
| set ike-version 2 | ✅ | ✅ | Enables IKEv2 | 
| set net-device disable | ✅ | ✅ | Use interface-mode VPN | 
| set dpd on-idle | ✅ | ✅ | Dead Peer Detection | 
| set psksecret "..." | ✅ | ✅ | Shared key | 
| set proposal aes256-sha256 | ✅ | ✅ | Cipher / hash pair | 
| Test | Command | Expect | 
|---|---|---|
| Tunnel status | get vpn ipsec tunnel summary | up/active | 
| Routing table | get router info routing-table all | Remote LAN via VPN | 
| Ping test | execute ping 192.168.30.1 | Success | 
| Log | diagnose vpn tunnel list | Established = yes | 
With these CLI commands, any engineer—junior or senior—can configure a secure, working site-to-site IPsec VPN between different FortiGate firmware versions.
The key is to mirror Phase 1/2 parameters, ensure no NAT, and verify routing and policies.
