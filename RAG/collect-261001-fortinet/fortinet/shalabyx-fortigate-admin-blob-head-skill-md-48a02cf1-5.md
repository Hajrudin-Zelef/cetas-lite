---
id: collect-261001-fortinet/fortinet/shalabyx-fortigate-admin-blob-head-skill-md-48a02cf1-5
title: "Use VIP in policy"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["memory", "sandbox"]
source: docs/RAG/collect-261001-fortinet/shalabyx-fortigate-admin-blob-head-skill-md-48a02cf1.md
source_anchor: ""
source_lines: [1089, 1122]
sha256: a4910501c67c1882a45bb7c2f3c7857fa257973712e169a952685c1bf18f1192
---

# Use VIP in policy

            set av-scan enable
            set sandbox-inspection enable
        end
    next
enddiagnose test application scanunitd 6
get system fortisandbox
| Task | Command | 
|---|---|
| Show interfaces | get system interface physical | 
| Show routing table | get router info routing-table all | 
| Show firewall policies | show firewall policy | 
| Show active sessions | diagnose sys session list | 
| Show HA status | get system ha status | 
| Show SD-WAN status | diagnose sys sdwan member | 
| Show VPN tunnels | diagnose vpn tunnel list | 
| Packet capture | diagnose sniffer packet <intf> <filter> 4 | 
| Flow debug | diagnose debug flow trace start 100 | 
| Show logs | execute log display | 
| CPU/Memory status | get system performance status | 
| List VDOMs | diagnose sys vd list | 
| VDOM resource usage | diagnose sys vd info | 
| Switch VDOM context | config vdom → edit "VDOM-NAME" | 
| Test LDAP auth | diagnose test authserver ldap <server> <user> <pass> | 
| Test RADIUS auth | diagnose test authserver radius <server> pap <user> <pass> | 
| Check FSSO logins | diagnose debug authd fsso list | 
| BGP summary | get router info bgp summary | 
| OSPF neighbors | get router info ospf neighbor | 
| DHCP leases | diagnose ip dhcp lease list | 
| Backup config | execute backup config tftp <file> <server> | 
| Restore config | execute restore config tftp <file> <server> | 
| API test | curl -k -H "Authorization: Bearer <key>" https://<fg>/api/v2/cmdb/system/status/ | 
| FortiSandbox status | get system fortisandbox | 
| FortiAnalyzer status | get log fortianalyzer status | 
| Managed switches | config switch-controller managed-switch |
