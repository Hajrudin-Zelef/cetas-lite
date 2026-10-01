---
id: collect-261001-fortinet/fortinet/shalabyx-fortigate-admin-blob-head-skill-md-48a02cf1-3
title: "Use VIP in policy"
domain: fortinet
role: reference
task: reference
actors: ["China"]
dates: []
keywords: ["agent", "license"]
source: docs/RAG/collect-261001-fortinet/shalabyx-fortigate-admin-blob-head-skill-md-48a02cf1.md
source_anchor: ""
source_lines: [540, 822]
sha256: 4b3c1cf1874d1ffb84c5ce9f5e32150ba5177895a464d264ad9058960a09e1a0
---

# Use VIP in policy

    next
    edit "Block-CN"
        set type geography
        set country "CN"     # China
    next
end
# Group them
config firewall addrgrp
    edit "Blocked-GeoIP"
        set member "Block-Countries" "Block-CN"
    next
endconfig firewall policy
    edit 99
        set name "Block-GeoIP"
        set srcintf "wan1"
        set dstintf "internal"
        set srcaddr "Blocked-GeoIP"
        set dstaddr "all"
        set action deny
        set schedule "always"
        set service "ALL"
        set logtraffic all
    next
end
Place this policy above any ACCEPT policies.
Policy & Objects → Addresses → Create New → Geography
config dlp sensor
    edit "DLP-Sensor"
        config filter
            edit 1
                set name "Credit-Cards"
                set type creditcard
                set action block
                set log enable
            next
            edit 2
                set name "SSN"
                set type ssn
                set action block
                set log enable
            next
        end
    next
end
# Apply to firewall policy
config firewall policy
    edit 1
        set dlp-sensor "DLP-Sensor"
        set utm-status enable
    next
end
Security Profiles → Data Loss Prevention
config user ldap
    edit "AD-LDAP"
        set server "192.168.1.10"
        set cnid "sAMAccountName"
        set dn "DC=company,DC=com"
        set type regular
        set username "CN=svc-fortigate,OU=Service Accounts,DC=company,DC=com"
        set password "LDAPpassword123!"
    next
endconfig user radius
    edit "RADIUS-Server"
        set server "192.168.1.20"
        set secret "RADIUSsecret123!"
        set auth-type auto
    next
endconfig user group
    edit "VPN-Users"
        set member "AD-LDAP"
    next
enddiagnose test authserver ldap AD-LDAP <username> <password>
diagnose test authserver radius RADIUS-Server pap <username> <password>
User & Authentication → LDAP Servers / RADIUS Servers
config user fsso
    edit "FSSO-Agent"
        set server "192.168.1.10"
        set password "FSSOpwd123!"
        set port 8000
    next
end
config user group
    edit "Domain-Users"
        set member "FSSO-Agent"
        config match
            edit 1
                set server-name "FSSO-Agent"
                set group-name "CN=Domain Users,CN=Users,DC=company,DC=com"
            next
        end
    next
endconfig firewall policy
    edit 1
        set groups "Domain-Users"
    next
enddiagnose debug authd fsso list
diagnose debug authd fsso server-statusconfig system admin
    edit "admin"
        set two-factor fortitoken
        set fortitoken "FTKMOBILE-TOKENID"
        set email-to "admin@company.com"
    next
endconfig user local
    edit "vpnuser1"
        set type password
        set passwd "UserPass123!"
        set two-factor fortitoken
        set fortitoken "FTKMOBILE-TOKENID"
        set email-to "user@company.com"
    next
endexecute fortitoken activate <token-id>
User & Authentication → FortiTokens
config firewall access-proxy
    edit "ZTNA-Proxy"
        set vip "ZTNA-VIP"
        set client-cert enable
        config api-gateway
            edit 1
                set url-map "/app1"
                set service tcp-forwarding
                config realservers
                    edit 1
                        set ip 192.168.1.100
                        set port 443
                    next
                end
            next
        end
    next
endconfig firewall policy
    edit 1
        set name "ZTNA-Policy"
        set srcintf "wan1"
        set dstintf "internal"
        set srcaddr "all"
        set dstaddr "ZTNA-VIP"
        set action accept
        set ztna-status enable
        set ztna-tags-match-logic and
    next
end
Policy & Objects → ZTNA → Access Proxy
config router ospf
    set router-id 1.1.1.1
    config area
        edit 0.0.0.0
        next
    end
    config network
        edit 1
            set prefix 192.168.1.0 255.255.255.0
            set area 0.0.0.0
        next
        edit 2
            set prefix 10.0.0.0 255.255.255.0
            set area 0.0.0.0
        next
    end
    config redistribute "connected"
        set status enable
    end
    config redistribute "static"
        set status enable
    end
endget router info ospf neighbor
get router info ospf status
get router info routing-table ospf
diagnose ip router ospf all enable
diagnose debug enableconfig router bgp
    set as 65001
    set router-id 1.1.1.1
    config neighbor
        edit "203.0.113.254"
            set remote-as 65002
            set activate enable
            set soft-reconfiguration enable
        next
    end
    config network
        edit 1
            set prefix 192.168.0.0 255.255.0.0
        next
    end
    config redistribute "connected"
        set status enable
    end
endget router info bgp summary
get router info bgp neighbors
get router info routing-table bgp
diagnose ip router bgp all enable
diagnose debug enableconfig router policy
    edit 1
        set input-device "internal"
        set src 192.168.10.0 255.255.255.0    # Finance VLAN
        set dst 0.0.0.0 0.0.0.0
        set gateway 203.0.113.254
        set output-device "wan1"
    next
    edit 2
        set input-device "internal"
        set src 192.168.20.0 255.255.255.0    # Guest VLAN
        set dst 0.0.0.0 0.0.0.0
        set gateway 198.51.100.254
        set output-device "wan2"
    next
endget router info routing-table all
diagnose ip route listconfig firewall shaper traffic-shaper
    edit "VoIP-Priority"
        set guaranteed-bandwidth 2048      # 2 Mbps guaranteed
        set maximum-bandwidth 10240        # 10 Mbps max
        set priority high
    next
    edit "BulkTraffic"
        set guaranteed-bandwidth 512
        set maximum-bandwidth 5120
        set priority low
    next
endconfig firewall shaping-policy
    edit 1
        set name "VoIP-Shaping"
        set service "SIP" "H323"
        set srcintf "internal"
        set dstintf "wan1"
        set srcaddr "all"
        set dstaddr "all"
        set traffic-shaper "VoIP-Priority"
        set traffic-shaper-reverse "VoIP-Priority"
    next
    edit 2
        set name "Bulk-Shaping"
        set service "BitTorrent"
        set srcintf "internal"
        set dstintf "wan1"
        set srcaddr "all"
        set dstaddr "all"
        set traffic-shaper "BulkTraffic"
    next
end
Policy & Objects → Traffic Shapers / Shaping Policy
config system api-user
    edit "api-admin"
        set comments "API user for automation"
        set api-key <generated-key>
        config trusthost
            edit 1
                set ipv4-trusthost 192.168.1.0 255.255.255.0
            next
        end
        set accprofile "super_admin"
    next
end
System → Administrators → Create New → REST API Admin → Generate API Key
# Get firewall policies
curl -k -H "Authorization: Bearer <API_KEY>" \
  https://192.168.1.1/api/v2/cmdb/firewall/policy/
# Get interfaces
curl -k -H "Authorization: Bearer <API_KEY>" \
  https://192.168.1.1/api/v2/cmdb/system/interface/
# Get routing table
curl -k -H "Authorization: Bearer <API_KEY>" \
  https://192.168.1.1/api/v2/monitor/router/ipv4/
# Create address object
curl -k -X POST -H "Authorization: Bearer <API_KEY>" \
  -H "Content-Type: application/json" \
  -d '{"name":"Test-Host","type":"ipmask","subnet":"10.0.0.1/32"}' \
  https://192.168.1.1/api/v2/cmdb/firewall/address/
# Backup config via API
curl -k -H "Authorization: Bearer <API_KEY>" \
  https://192.168.1.1/api/v2/monitor/system/config/backup?scope=global \
  -o fortigate-backup.confconfig system automation-trigger
    edit "IPS-Event"
        set event-type ips-signature
        set license-type forticare
    next
end
config system automation-action
    edit "Email-Alert"
        set action-type email
        set email-to "admin@company.com"
        set email-from "fortigate@company.com"
        set message "IPS Event detected: %%log%%"
    next
end
