---
id: collect-261001-fortinet/fortinet/ghostkellz-fortigate-d2e2ec09-1
title: "Get policy summary"
domain: fortinet
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: ["cost", "latency"]
source: docs/RAG/collect-261001-fortinet/ghostkellz-fortigate-d2e2ec09.md
source_anchor: ""
source_lines: [1, 291]
sha256: 597b9563cb2ec25cbbcfb48eff77322a98959419f920106983d000e43bb01654
---

# Get policy summary

Scripts, configs, and reference docs for FortiGate firewalls, FortiClient deployment, SD-WAN, and more.
Deploys FortiClient IPsec VPN configuration via GPO startup script.
How it works:
- Copies VPN XML config from network share to local machine
- Imports config using fcconfig.exe
- Creates marker file to prevent re-import on subsequent boots
GPO Setup:
- Create scheduled task: At startup + at login
- Program: powershell
- Arguments: -ExecutionPolicy Bypass -File "\\SERVER\share\deploy.ps1"
Export VPN Config (from configured machine):
& "C:\Program Files\Fortinet\FortiClient\fcconfig.exe" -m vpn -o export -f "C:\Temp\vpn.xml" -p "YourPassword"
Import VPN Config:
& "C:\Program Files\Fortinet\FortiClient\fcconfig.exe" -m vpn -o import -f "C:\Temp\vpn.xml" -p "YourPassword"# Show all policies
show firewall policy
# Get policy summary
get firewall policy
# Show specific policy by ID
show firewall policy 5
# Check policy hit count
diagnose firewall iprope list 100004config firewall policy
    edit 0
        set name "LAN-to-Internet"
        set srcintf "lan"
        set dstintf "wan1"
        set srcaddr "all"
        set dstaddr "all"
        set action accept
        set schedule "always"
        set service "ALL"
        set nat enable
        set logtraffic all
    next
endconfig firewall policy
    edit 0
        set name "Secure-Web-Access"
        set srcintf "lan"
        set dstintf "wan1"
        set srcaddr "LAN_Subnet"
        set dstaddr "all"
        set action accept
        set schedule "always"
        set service "HTTP" "HTTPS"
        set utm-status enable
        set av-profile "default"
        set webfilter-profile "default"
        set ips-sensor "default"
        set ssl-ssh-profile "certificate-inspection"
        set nat enable
        set logtraffic all
    next
end# Create address object
config firewall address
    edit "Server-10.0.0.50"
        set subnet 10.0.0.50 255.255.255.255
    next
    edit "LAN_Subnet"
        set subnet 192.168.1.0 255.255.255.0
    next
    edit "Remote-Office"
        set subnet 10.10.0.0 255.255.0.0
    next
end
# Create address group
config firewall addrgrp
    edit "Internal-Servers"
        set member "Server-10.0.0.50" "Server-10.0.0.51"
    next
end
# FQDN address
config firewall address
    edit "Google-DNS"
        set type fqdn
        set fqdn "dns.google"
    next
end# Custom service
config firewall service custom
    edit "Custom-App-8080"
        set tcp-portrange 8080
    next
    edit "Custom-App-Range"
        set tcp-portrange 8000-8100
    next
    edit "Custom-UDP"
        set udp-portrange 5000-5100
    next
end
# Service group
config firewall service group
    edit "Web-Services"
        set member "HTTP" "HTTPS" "Custom-App-8080"
    next
end# Move policy (order matters - first match wins)
config firewall policy
    move 10 before 5
end
# Delete policy
config firewall policy
    delete 15
end
# Clone/copy a policy (edit 0 creates new)
# First show the policy you want to copy, then recreate it# Explicit deny with logging
config firewall policy
    edit 0
        set name "Block-BadStuff"
        set srcintf "wan1"
        set dstintf "lan"
        set srcaddr "Blocked-IPs"
        set dstaddr "all"
        set action deny
        set schedule "always"
        set service "ALL"
        set logtraffic all
    next
end# Show all routes
get router info routing-table all
# Show route details
get router info routing-table details
# Show static routes config
show router static
# Check specific route
get router info routing-table databaseconfig router static
    # Default route via WAN1
    edit 1
        set dst 0.0.0.0/0
        set gateway 192.168.1.1
        set device "wan1"
        set distance 10
        set priority 0
    next
    # Default route via WAN2 (backup, higher distance)
    edit 2
        set dst 0.0.0.0/0
        set gateway 10.0.0.1
        set device "wan2"
        set distance 20
        set priority 0
    next
    # Route to remote network via VPN
    edit 3
        set dst 10.10.0.0/16
        set device "vpn-tunnel1"
    next
    # Blackhole route (drop traffic)
    edit 4
        set dst 192.168.99.0/24
        set blackhole enable
    next
end# Create link monitor
config system link-monitor
    edit "WAN1-Monitor"
        set srcintf "wan1"
        set server "8.8.8.8" "1.1.1.1"
        set protocol ping
        set gateway-ip 192.168.1.1
        set interval 500
        set failtime 3
        set recoverytime 3
        set update-static-route enable
    next
end
# Static route tied to link monitor
config router static
    edit 1
        set dst 0.0.0.0/0
        set gateway 192.168.1.1
        set device "wan1"
        set link-monitor-exempt enable
    next
end# Force specific traffic out a specific interface
config router policy
    edit 1
        set input-device "lan"
        set src "192.168.1.100/32"
        set dst "0.0.0.0/0"
        set output-device "wan2"
        set gateway 10.0.0.1
    next
endconfig router static
    edit 1
        set dst 0.0.0.0/0
        set gateway 192.168.1.1
        set device "wan1"
        set distance 10
        set weight 3    # 75% of traffic
    next
    edit 2
        set dst 0.0.0.0/0
        set gateway 10.0.0.1
        set device "wan2"
        set distance 10
        set weight 1    # 25% of traffic
    next
end
# Enable ECMP
config system settings
    set v4-ecmp-mode weight-based
end# Create SD-WAN zone
config system sdwan
    set status enable
    config zone
        edit "virtual-wan-link"
        next
    end
end
# Add interfaces to SD-WAN
config system sdwan
    config members
        edit 1
            set interface "wan1"
            set gateway 192.168.1.1
            set cost 0
        next
        edit 2
            set interface "wan2"
            set gateway 10.0.0.1
            set cost 10
        next
    end
endconfig system sdwan
    config health-check
        edit "Google-DNS"
            set server "8.8.8.8"
            set protocol ping
            set interval 500
            set failtime 3
            set recoverytime 3
            set members 1 2
        next
        edit "Cloudflare"
            set server "1.1.1.1"
            set protocol ping
            set members 1 2
        next
        edit "HTTP-Check"
            set server "www.google.com"
            set protocol http
            set port 80
            set members 1 2
        next
    end
endconfig system sdwan
    config service
        edit 1
            set name "Critical-Apps"
            set mode priority
            set dst "all"
            set src "all"
            set priority-members 1 2
            set health-check "Google-DNS"
        next
        edit 2
            set name "VoIP-Traffic"
            set mode sla
            set dst "all"
            set internet-service enable
            set internet-service-app-ctrl 16354 16355  # MS Teams, Zoom
            config sla
                edit "Google-DNS"
                    set latency-threshold 100
                    set jitter-threshold 20
                    set packetloss-threshold 1
                next
            end
            set priority-members 1
        next
        edit 3
            set name "Bulk-Downloads"
            set mode load-balance
            set dst "all"
            set internet-service enable
            set internet-service-app-ctrl 33182  # Downloads
            set priority-members 1 2
        next
    end
enddiagnose sys sdwan health-check
diagnose sys sdwan member
diagnose sys sdwan service
diagnose sys sdwan intf-sla-log
get router info routing-table all
For basic active/standby WAN failover without full SD-WAN:
# Create link monitors for each WAN
config system link-monitor
    edit "WAN1-Failover"
        set srcintf "wan1"
        set server "8.8.8.8" "1.1.1.1"
        set protocol ping
        set gateway-ip 192.168.1.1
        set interval 500
        set failtime 3
        set recoverytime 3
