---
id: collect-261001-fortinet/fortinet/ghostkellz-fortigate-d2e2ec09-3
title: "Get policy summary"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/ghostkellz-fortigate-d2e2ec09.md
source_anchor: ""
source_lines: [570, 702]
sha256: a05ab07ab4209007f7ce50663882bbc4f199a5fa7a7275fb17bfbcb8cff49ec5
---

# Get policy summary

end
# Auto-block attacker
config system automation-stitch
    edit "Auto-Block-Brute-Force"
        set status enable
        set trigger "Brute-Force-Detected"
        config actions
            edit 1
                set action "Block-Attacker-IP"
                set required enable
            next
            edit 2
                set action "Email-Admin"
            next
        end
    next
end
# Daily backup via webhook
config system automation-stitch
    edit "Daily-Backup-Notify"
        set status enable
        set trigger "Daily-6AM"
        config actions
            edit 1
                set action "Backup-Webhook"
            next
        end
    next
end
| Event | Log ID | 
|---|---|
| Admin login | 32001 | 
| Admin logout | 32002 | 
| Failed admin login | 32003 | 
| Config change | 32102 | 
| IPS alert | 16384 | 
| AV detection | 8192 | 
| Web filter block | 12288 | 
| VPN tunnel up | 37138 | 
| VPN tunnel down | 37139 | 
| HA failover | 41216 | 
| Interface up | 20480 | 
| Interface down | 20481 | 
Use these in action messages:
| Variable | Description | 
|---|---|
| %%log.srcip%% | Source IP | 
| %%log.dstip%% | Destination IP | 
| %%log.srcport%% | Source port | 
| %%log.dstport%% | Destination port | 
| %%log.srcmac%% | Source MAC | 
| %%log.action%% | Action taken | 
| %%log.logdesc%% | Log description | 
| %%log.msg%% | Log message | 
| %%log.date%% | Date | 
| %%log.time%% | Time | 
| %%log.user%% | Username | 
| %%log.devname%% | Device name | 
# Show triggers
show system automation-trigger
# Show actions
show system automation-action
# Show stitches
show system automation-stitch
# Check automation status
diagnose automation test <stitch-name>
# Debug automation
diagnose debug application autod -1
diagnose debug enable# Show system status
get system status
get system performance status
# Reboot
execute reboot
# Firmware upgrade
execute restore image tftp <filename> <tftp-server-ip>
# Factory reset
execute factoryreset# Show interfaces
get system interface
show system interface
# Bring interface up/down
config system interface
    edit "wan1"
        set status up
    next
end# IPsec tunnel status
get vpn ipsec tunnel summary
diagnose vpn ike gateway list
diagnose vpn tunnel list
# SSL VPN
get vpn ssl monitor
diagnose vpn ssl list# Active sessions
get system session list
diagnose sys session filter clear
diagnose sys session filter dport 443
diagnose sys session list
# Sniffer
diagnose sniffer packet any "host 192.168.1.100" 4 0 l
diagnose sniffer packet wan1 "port 443" 4 100# Via CLI (to TFTP)
execute backup config tftp <filename> <tftp-server-ip>
# Via CLI (to USB)
execute backup config usb <filename>execute restore config tftp <filename> <tftp-server-ip>
execute restore config usb <filename># SCP backup example (from external host)
scp admin@<fortigate-ip>:sys_config /backups/fortigate-$(date +%Y%m%d).conf# Enable debug
diagnose debug enable
diagnose debug console timestamp enable
# Debug flow (traffic)
diagnose debug flow filter addr 192.168.1.100
diagnose debug flow trace start 100
diagnose debug flow trace stop
# Debug IPsec
diagnose vpn ike log filter dst-addr4 <peer-ip>
diagnose debug app ike -1
diagnose debug enable
# Stop all debug
diagnose debug disable
diagnose debug reset
| Issue | Command | 
|---|---|
| Policy not matching | diagnose debug flow trace start 100 | 
| IPsec tunnel down | diagnose vpn ike gateway list | 
| High CPU | get system performance top | 
| Session table full | get system session status | 
| DNS issues | execute ping-options source <interface> thenexecute ping 8.8.8.8 | 
| Failover not working | diagnose sys link-monitor interface | 
| SD-WAN issues | diagnose sys sdwan health-check | 
# View logs
execute log filter category event
execute log filter device disk
execute log display
# Real-time log
execute log filter field action deny
execute log display
MIT - Do whatever you want with it.
