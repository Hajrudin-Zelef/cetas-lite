---
id: collect-261001-fortinet/fortinet/ghostkellz-fortigate-d2e2ec09-2
title: "Get policy summary"
domain: fortinet
role: reference
task: reference
actors: ["AWS", "Lambda"]
dates: []
keywords: ["aws"]
source: docs/RAG/collect-261001-fortinet/ghostkellz-fortigate-d2e2ec09.md
source_anchor: ""
source_lines: [292, 569]
sha256: d4336fdf3976071b717c9f15202675afcd0e26a8ed21a5e11790b58333c7cee7
---

# Get policy summary

        set update-static-route enable
    next
    edit "WAN2-Failover"
        set srcintf "wan2"
        set server "8.8.8.8" "1.1.1.1"
        set protocol ping
        set gateway-ip 10.0.0.1
        set interval 500
        set failtime 3
        set recoverytime 3
        set update-static-route enable
    next
end
# Static routes with distance (lower = preferred)
config router static
    edit 1
        set dst 0.0.0.0/0
        set gateway 192.168.1.1
        set device "wan1"
        set distance 10
    next
    edit 2
        set dst 0.0.0.0/0
        set gateway 10.0.0.1
        set device "wan2"
        set distance 20
    next
endconfig system sdwan
    set status enable
    config members
        edit 1
            set interface "wan1"
            set gateway 192.168.1.1
            set priority 0
        next
        edit 2
            set interface "wan2"
            set gateway 10.0.0.1
            set priority 10
        next
    end
    config health-check
        edit "Failover-Check"
            set server "8.8.8.8" "1.1.1.1"
            set protocol ping
            set interval 500
            set failtime 3
            set recoverytime 3
            set members 1 2
        next
    end
    config service
        edit 1
            set name "Primary-WAN1-Failover-WAN2"
            set mode priority
            set dst "all"
            set src "all"
            set health-check "Failover-Check"
            set priority-members 1 2
        next
    end
end# Check link monitor status
diagnose sys link-monitor interface
# Check which WAN is active
get router info routing-table all
# SD-WAN member status
diagnose sys sdwan member
# Real-time failover events
diagnose sys sdwan health-check
# Check failover history in logs
execute log filter category event
execute log filter field msg "link-monitor"
execute log display# Enable session pickup on failover (keeps connections alive)
config system ha
    set session-pickup enable
end
# Or for SD-WAN, sessions re-establish automatically
# Check active sessions
diagnose sys session list
External threat intelligence feeds for blocking malicious IPs, domains, and URLs.
# Create external threat feed connector
config system external-resource
    edit "Malicious-IPs"
        set type address
        set resource "https://example.com/threat-feed/malicious-ips.txt"
        set refresh-rate 60
    next
    edit "TOR-Exit-Nodes"
        set type address
        set resource "https://check.torproject.org/torbulkexitlist"
        set refresh-rate 1440
    next
end
# Use in firewall policy
config firewall policy
    edit 0
        set name "Block-Threat-Feed"
        set srcintf "wan1"
        set dstintf "lan"
        set srcaddr "Malicious-IPs" "TOR-Exit-Nodes"
        set dstaddr "all"
        set action deny
        set schedule "always"
        set service "ALL"
        set logtraffic all
    next
endconfig system external-resource
    edit "Malicious-Domains"
        set type domain
        set resource "https://example.com/threat-feed/malicious-domains.txt"
        set refresh-rate 60
    next
end
# Use in DNS filter profile
config dnsfilter profile
    edit "Block-Malicious"
        config ftgd-dns
            # Enable categories as needed
        end
        set external-ip-blocklist "Malicious-Domains"
    next
end# Enable FortiGuard outbreak prevention
config ips global
    set fail-open enable
end
config ips sensor
    edit "Default"
        set scan-botnet-connections block
    next
end
# Internet Service Database (ISDB) for app control
# Used in SD-WAN and policies automatically
get firewall internet-service-name
| Feed | URL | Type | 
|---|---|---|
| Abuse.ch Feodo Tracker | https://feodotracker.abuse.ch/downloads/ipblocklist.txt | IPs | 
| Spamhaus DROP | https://www.spamhaus.org/drop/drop.txt | CIDRs | 
| TOR Exit Nodes | https://check.torproject.org/torbulkexitlist | IPs | 
| Emerging Threats | https://rules.emergingthreats.net/fwrules/emerging-Block-IPs.txt | IPs | 
# View external resources
diagnose sys external-resource list
diagnose sys external-resource entry list <name>
# Force refresh
diagnose sys external-resource update <name>
Automation stitches let you trigger actions based on events (logs, schedules, etc.). Great for automated responses to security events, notifications, and custom workflows.
- Trigger: Event that starts the automation (log event, schedule, etc.)
- Action: What happens when triggered (email, webhook, CLI script, etc.)
- Stitch: Combines trigger + action(s)
# Event-based trigger (e.g., admin login)
config system automation-trigger
    edit "Admin-Login-Trigger"
        set event-type event-log
        set logid 32001  # Admin login event
    next
end
# IPS attack trigger
config system automation-trigger
    edit "IPS-Critical-Alert"
        set event-type event-log
        set logid 16384
        config fields
            edit 1
                set name "severity"
                set value "critical"
            next
        end
    next
end
# Schedule trigger (daily at 6am)
config system automation-trigger
    edit "Daily-6AM"
        set trigger-type scheduled
        set trigger-frequency daily
        set trigger-hour 6
        set trigger-minute 0
    next
end
# Incoming webhook trigger
config system automation-trigger
    edit "Webhook-Trigger"
        set trigger-type incoming-webhook
    next
end
# HA failover trigger
config system automation-trigger
    edit "HA-Failover"
        set event-type ha-failover
    next
end
# FortiGuard update trigger
config system automation-trigger
    edit "AV-DB-Updated"
        set event-type faz-event
        set event-name "av-db-update"
    next
end# Email notification
config system automation-action
    edit "Email-Admin"
        set action-type email
        set email-to "admin@company.com"
        set email-from "fortigate@company.com"
        set email-subject "FortiGate Alert: %%log.logdesc%%"
        set message "Event: %%log.logdesc%%\nSource: %%log.srcip%%\nTime: %%log.date%% %%log.time%%"
    next
end
# Webhook/API call
config system automation-action
    edit "Slack-Webhook"
        set action-type webhook
        set uri "https://hooks.slack.com/services/xxx/yyy/zzz"
        set http-body "{\"text\": \"FortiGate Alert: %%log.logdesc%% from %%log.srcip%%\"}"
        set port 443
        set protocol https
        set method post
    next
end
# Teams webhook
config system automation-action
    edit "Teams-Alert"
        set action-type webhook
        set uri "https://outlook.office.com/webhook/xxx"
        set http-body "{\"text\": \"**FortiGate Alert**\\n\\nEvent: %%log.logdesc%%\\nSource: %%log.srcip%%\"}"
        set port 443
        set protocol https
        set method post
    next
end
# Run CLI script
config system automation-action
    edit "Block-Attacker-IP"
        set action-type cli-script
        set script "config firewall address
    edit \"blocked-%%log.srcip%%\"
        set subnet %%log.srcip%%/32
    next
end
config firewall addrgrp
    edit \"Auto-Blocked\"
        append member \"blocked-%%log.srcip%%\"
    next
end"
    next
end
# Quarantine host
config system automation-action
    edit "Quarantine-Host"
        set action-type quarantine
        set quarantine-host-mac %%log.srcmac%%
    next
end
# Ban IP
config system automation-action
    edit "Ban-IP"
        set action-type ip-ban
    next
end
# AWS Lambda
config system automation-action
    edit "AWS-Lambda"
        set action-type aws-lambda
        set aws-api-id "your-api-id"
        set aws-region "us-east-1"
        set aws-api-key "your-key"
    next
endconfig system automation-stitch
    edit "Alert-On-IPS-Critical"
        set status enable
        set trigger "IPS-Critical-Alert"
        config actions
            edit 1
                set action "Email-Admin"
                set required enable
            next
            edit 2
                set action "Slack-Webhook"
            next
        end
    next
