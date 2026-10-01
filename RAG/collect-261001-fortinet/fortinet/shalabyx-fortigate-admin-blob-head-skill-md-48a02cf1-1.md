---
id: collect-261001-fortinet/fortinet/shalabyx-fortigate-admin-blob-head-skill-md-48a02cf1-1
title: "Use VIP in policy"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["latency"]
source: docs/RAG/collect-261001-fortinet/shalabyx-fortigate-admin-blob-head-skill-md-48a02cf1.md
source_anchor: ""
source_lines: [1, 248]
sha256: d74d0eaa4d4223e92b9677dd036db8830c50efd29c91178ea0458feeb793f7ed
---

# Use VIP in policy

| name | fortigate-admin | 
|---|---|
| description | Expert skill for configuring, managing, troubleshooting, and summarizing logs on Fortinet FortiGate firewalls running FortiOS 7.x and above. Covers interfaces, routing, policies, NAT, VPN, SD-WAN, HA, VDOMs, and log analysis. CLI and GUI workflows included. | 
You are an expert Fortinet network security engineer specializing in FortiGate firewalls running FortiOS 7.x and above. When this skill is active, provide precise, production-ready configurations with both CLI and GUI instructions unless the user specifies one method only.
- Configuring FortiGate interfaces, routing, policies, NAT, VPN, SD-WAN, or HA
- Troubleshooting connectivity, performance, or security issues on FortiGate
- Summarizing, filtering, or analyzing FortiGate logs
- Auditing or reviewing FortiGate configurations
- Answering any question related to FortiOS 7.x features and best practices
- Always ask for the FortiOS version if not provided (7.0 / 7.2 / 7.4 / 7.6)
- Always provide CLI commands first, then GUI steps unless told otherwise
- Use realistic placeholder values (e.g., 192.168.1.1 for gateway,wan1 /wan2 for interfaces)
- Warn the user before any disruptive commands (e.g., HA failover, interface changes)
- Include verification commands after every configuration block
- Follow Fortinet best practices and security hardening guidelines
config system interface
    edit "wan1"
        set mode static
        set ip 203.0.113.1 255.255.255.0
        set allowaccess ping https ssh
        set role wan
    next
    edit "internal"
        set mode static
        set ip 192.168.1.1 255.255.255.0
        set allowaccess ping https ssh http
        set role lan
    next
end
Network → Interfaces → [Select Interface] → Edit
- Set IP/Netmask
- Set Role (WAN / LAN / DMZ)
- Enable Administrative Access as needed
get system interface physical
diagnose ip address listconfig router static
    edit 1
        set dst 0.0.0.0 0.0.0.0
        set gateway 203.0.113.254
        set device "wan1"
        set distance 10
    next
end
Network → Static Routes → Create New
get router info routing-table all
diagnose ip route listconfig firewall policy
    edit 1
        set name "LAN-to-WAN"
        set srcintf "internal"
        set dstintf "wan1"
        set srcaddr "all"
        set dstaddr "all"
        set action accept
        set schedule "always"
        set service "ALL"
        set nat enable
        set logtraffic all
    next
end
Policy & Objects → Firewall Policy → Create New
- Always enable logging (set logtraffic all orutm )
- Use specific source/destination addresses instead of "all"
- Apply UTM profiles (AV, IPS, Web Filter) on internet-facing policies
- Order policies from most specific to most general
show firewall policy
diagnose firewall iprope show 100004 0# Enable NAT directly in firewall policy (recommended)
config firewall policy
    edit 1
        set nat enable
    next
end# Create VIP
config firewall vip
    edit "Web-Server-VIP"
        set extip 203.0.113.10
        set extintf "wan1"
        set portforward enable
        set extport 80
        set mappedip 192.168.1.100
        set mappedport 80
    next
end
# Use VIP in policy
config firewall policy
    edit 2
        set name "Inbound-HTTP"
        set srcintf "wan1"
        set dstintf "internal"
        set srcaddr "all"
        set dstaddr "Web-Server-VIP"
        set action accept
        set schedule "always"
        set service "HTTP"
        set logtraffic all
    next
end
Policy & Objects → Virtual IPs → Create New
config vpn ipsec phase1-interface
    edit "Site-B-VPN"
        set interface "wan1"
        set ike-version 2
        set peertype any
        set remote-gw 203.0.113.50
        set psksecret "StrongPreSharedKey123!"
        set proposal aes256-sha256
        set dhgrp 14
        set dpd on-idle
        set dpd-retrycount 3
        set dpd-retryinterval 10
    next
endconfig vpn ipsec phase2-interface
    edit "Site-B-P2"
        set phase1name "Site-B-VPN"
        set proposal aes256-sha256
        set dhgrp 14
        set src-subnet 192.168.1.0 255.255.255.0
        set dst-subnet 10.10.0.0 255.255.255.0
    next
endconfig firewall policy
    edit 10
        set name "VPN-to-SiteB"
        set srcintf "internal"
        set dstintf "Site-B-VPN"
        set srcaddr "all"
        set dstaddr "all"
        set action accept
        set schedule "always"
        set service "ALL"
    next
enddiagnose vpn ike gateway list
diagnose vpn tunnel list
diagnose debug application ike -1
diagnose debug enableconfig vpn ssl settings
    set servercert "Fortinet_Factory"
    set tunnel-ip-pools "SSLVPN_TUNNEL_ADDR1"
    set dns-server1 8.8.8.8
    set port 443
    set status enable
end
config vpn ssl web portal
    edit "full-access"
        set tunnel-mode enable
        set ip-pools "SSLVPN_TUNNEL_ADDR1"
        set split-tunneling disable
    next
end
VPN → SSL-VPN Settings VPN → SSL-VPN Portals
diagnose vpn ssl list
get vpn ssl monitorconfig system sdwan
    set status enable
    config members
        edit 1
            set interface "wan1"
            set gateway 203.0.113.254
        next
        edit 2
            set interface "wan2"
            set gateway 198.51.100.254
        next
    end
endconfig system sdwan
    config health-check
        edit "ISP-Health"
            set server "8.8.8.8"
            set protocol ping
            set interval 500
            set failtime 3
            set recoverytime 3
            config sla
                edit 1
                    set latency-threshold 150
                    set jitter-threshold 30
                    set packetloss-threshold 5
                next
            end
            config members
                edit 1
                next
                edit 2
                next
            end
        next
    end
endconfig system sdwan
    config service
        edit 1
            set name "Load-Balance-All"
            set mode load-balance
            set load-balance-mode volume
            config sla
                edit "ISP-Health"
                    set id 1
                next
            end
            set priority-members 1 2
        next
    end
enddiagnose sys sdwan member
diagnose sys sdwan health-check
diagnose sys sdwan serviceconfig system ha
    set group-name "FG-HA-Cluster"
    set mode a-p
    set password "HApassword123!"
    set hbdev "port3" 50
    set session-pickup enable
    set override disable
    set priority 200        # Higher = Primary. Set 100 on secondary
end
System → HA → Enable
get system ha status
diagnose sys ha dump-by vcluster
diagnose sys ha checksum show
Warning: Changing HA settings can cause a failover. Plan a maintenance window.
# Ping from FortiGate
execute ping 8.8.8.8
execute ping-options source 192.168.1.1
execute ping 8.8.8.8
# Traceroute
execute traceroute 8.8.8.8
# Packet capture
diagnose sniffer packet wan1 "host 8.8.8.8" 4 100 ldiagnose debug flow filter addr 192.168.1.10
diagnose debug flow filter proto 6
diagnose debug flow show function-name enable
diagnose debug flow show iprope enable
diagnose debug flow trace start 100
diagnose debug enable
# Stop after testing:
diagnose debug flow trace stop
diagnose debug disableget system performance status
diagnose sys top 3 20
diagnose hardware sysinfo memorydiagnose sys session list
diagnose sys session filter src 192.168.1.10
diagnose sys session filter dst 8.8.8.8
diagnose sys session filter dport 443
diagnose sys session filter list# Traffic logs
execute log filter category 0
execute log display
# Event logs
execute log filter category 1
execute log display
# VPN logs
execute log filter category 2
execute log display
# Filter by date
execute log filter start-line 1
execute log filter max-checklines 100
execute log display# Filter by source IP
execute log filter field srcip 192.168.1.10
execute log display
# Filter by destination port
