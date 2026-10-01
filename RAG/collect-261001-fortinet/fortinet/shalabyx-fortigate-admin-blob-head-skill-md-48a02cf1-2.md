---
id: collect-261001-fortinet/fortinet/shalabyx-fortigate-admin-blob-head-skill-md-48a02cf1-2
title: "Use VIP in policy"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution"]
source: docs/RAG/collect-261001-fortinet/shalabyx-fortigate-admin-blob-head-skill-md-48a02cf1.md
source_anchor: ""
source_lines: [249, 539]
sha256: d5a445393a06fd60b80444291f5eda3d55f8509ad3b4ad845387a23b713ea32e
---

# Use VIP in policy

execute log filter field dstport 443
execute log display
# Filter by action (blocked traffic)
execute log filter field action deny
execute log display
When asked to summarize logs, follow this approach:
- Identify top talkers: Find source IPs generating most traffic/blocks
- Identify blocked traffic: Focus on action=deny entries
- Identify threat events: Look for IPS, AV, or web filter alerts
- VPN events: Summarize successful/failed authentication attempts
- Policy hits: Which policies are hit most frequently
Present summaries in this format:
## Log Summary — [Date Range]
### Traffic Overview
- Total sessions analyzed: X
- Allowed: X | Blocked: X | Threat detections: X
### Top Source IPs
1. 192.168.1.10 — 1,203 sessions
2. 192.168.1.25 — 876 sessions
### Top Blocked Destinations
1. 203.0.113.99 — 45 blocks (policy: LAN-to-WAN)
### Security Events
- IPS triggers: X (top signature: [name])
- AV detections: X
- Web filter blocks: X
### VPN Activity
- IPsec tunnels up/down events: X
- SSL VPN logins: X successful, X failed
### Recommendations
- [Actionable item based on log findings]
# Check FortiAnalyzer connection
diagnose test application oftpd 2
get log fortianalyzer setting
get log fortianalyzer-cloud setting
Log & Report → Forward Traffic / System Events / Security Events
- Use filters: Source IP, Destination, Action, Policy ID
- Use Log View → Add Filter to narrow results
- Export to CSV for external analysis
Warning: Enabling VDOMs requires a reboot and will reset interface assignments. Plan a maintenance window.
config system global
    set vdom-mode multi-vdom
end
# Confirm reboot when prompted
System → Settings → Virtual Domains → Enable
config vdom
    edit "VDOM-A"
    next
    edit "VDOM-B"
    next
end
System → VDOM → Create New
config system interface
    edit "wan1"
        set vdom "VDOM-A"
    next
    edit "internal1"
        set vdom "VDOM-A"
    next
    edit "wan2"
        set vdom "VDOM-B"
    next
end
Network → Interfaces → [Select Interface] → Edit → VDOM
# Enter a specific VDOM context
config vdom
    edit "VDOM-A"
end
# Return to global context
end
# Run commands inside a VDOM
config vdom
    edit "VDOM-A"
    config system interface
        show
    end
end# Create VDOM link pair
config system vdom-link
    edit "vlink0"
        set type ppp
    next
end
# Assign each end to a VDOM
config system interface
    edit "vlink0_0"
        set vdom "VDOM-A"
        set ip 10.255.0.1 255.255.255.252
        set allowaccess ping
    next
    edit "vlink0_1"
        set vdom "VDOM-B"
        set ip 10.255.0.2 255.255.255.252
        set allowaccess ping
    next
end
# Add static routes via VDOM link
config vdom
    edit "VDOM-A"
    config router static
        edit 1
            set dst 10.20.0.0 255.255.255.0
            set device "vlink0_0"
            set gateway 10.255.0.2
        next
    end
    end
endconfig global
config system admin
    edit "vdom-a-admin"
        set password "SecurePass123!"
        set vdom "VDOM-A"
        set accprofile "prof_admin"
    next
end
end
System → Administrators → Create New → Select VDOM
config vdom
    edit "VDOM-A"
    config system vdom-property
        set session-count 100000
        set ipsec-phase1 50
        set ipsec-phase2 50
        set dialup-tunnel 100
        set firewall-policy 500
    end
    end
end
System → VDOM → [Select VDOM] → Edit → Resource Limits
Warning: All interfaces, policies, and routes in the VDOM will be removed.
# First unassign all interfaces from the VDOM, then:
config vdom
    delete "VDOM-B"
endconfig system global
    set vdom-mode no-vdom
end# List all VDOMs
diagnose sys vd list
# Show VDOM resource usage
diagnose sys vd info
# Show interfaces per VDOM
show system interface | grep vdom
# Check routing table inside a VDOM
config vdom
    edit "VDOM-A"
    get router info routing-table all
    end
endconfig antivirus profile
    edit "AV-Profile"
        set comment "Default AV profile"
        config http
            set av-scan enable
            set outbreak-prevention enable
        end
        config ftp
            set av-scan enable
        end
        config smtp
            set av-scan enable
        end
    next
endconfig ips sensor
    edit "IPS-Profile"
        config entries
            edit 1
                set rule all
                set action default
                set status enable
            next
        end
    next
endconfig webfilter profile
    edit "WebFilter-Profile"
        config ftgd-wf
            config filters
                edit 1
                    set category 26        # Malicious websites
                    set action block
                next
                edit 2
                    set category 61        # Phishing
                    set action block
                next
            end
        end
        set log-all-url enable
    next
endconfig application list
    edit "AppControl-Profile"
        config entries
            edit 1
                set category 2             # P2P
                set action block
            next
            edit 2
                set category 27            # Social media
                set action monitor
            next
        end
    next
endconfig dnsfilter profile
    edit "DNS-Filter"
        set block-botnet-domains enable
        config ftgd-dns
            config filters
                edit 1
                    set category 26
                    set action block
                next
            end
        end
    next
endconfig firewall policy
    edit 1
        set av-profile "AV-Profile"
        set ips-sensor "IPS-Profile"
        set webfilter-profile "WebFilter-Profile"
        set application-list "AppControl-Profile"
        set dnsfilter-profile "DNS-Filter"
        set ssl-ssh-profile "deep-inspection"
        set utm-status enable
        set logtraffic all
    next
end
Policy & Objects → Firewall Policy → Edit → Security Profiles
config firewall ssl-ssh-profile
    edit "deep-inspection"
        set comment "Full SSL deep inspection"
        config https
            set ports 443
            set status deep-inspection
        end
        config ftps
            set ports 990
            set status deep-inspection
        end
        config imaps
            set ports 993
            set status deep-inspection
        end
        set caname "Fortinet_CA_SSL"
        set untrusted-caname "Fortinet_CA_Untrusted"
    next
end# Export FortiGate CA cert for distribution to clients
# GUI: System → Certificates → Download Fortinet_CA_SSL
execute vpn certificate ca export Fortinet_CA_SSLconfig firewall ssl-ssh-profile
    edit "deep-inspection"
        config ssl-exempt
            edit 1
                set type fqdn
                set fqdn "banking-site.com"
            next
            edit 2
                set type wildcard-fqdn
                set wildcard-fqdn "*.microsoft.com"
            next
        end
    next
end
Security Profiles → SSL/SSH Inspection
config firewall DoS-policy
    edit 1
        set name "DoS-WAN"
        set interface "wan1"
        set srcaddr "all"
        set dstaddr "all"
        set service "ALL"
        config anomaly
            edit "tcp_syn_flood"
                set status enable
                set log enable
                set action block
                set threshold 2000
            next
            edit "udp_flood"
                set status enable
                set log enable
                set action block
                set threshold 2000
            next
            edit "icmp_flood"
                set status enable
                set log enable
                set action block
                set threshold 500
            next
        end
    next
end
Policy & Objects → DoS Policy → Create New
diagnose ips anomaly listconfig firewall address
    edit "Block-Countries"
        set type geography
        set country "RU"     # Russia
