---
id: collect-261001-huawei/huawei/wp-content-uploads-2020-05-huawei-cico-cli-cheatsheet-new-pdf-9d083130
title: "wp-content-uploads-2020-05-huawei-cico-cli-cheatsheet-new-pdf-9d083130"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-huawei/wp-content-uploads-2020-05-huawei-cico-cli-cheatsheet-new-pdf-9d083130.md
source_anchor: ""
source_lines: [1, 73]
sha256: e78f721b8a06b104b275a0a30945cfefbdba2766586d25b02541db26d6e46d58
---

# wp-content-uploads-2020-05-huawei-cico-cli-cheatsheet-new-pdf-9d083130

Below is list of commands generally used in Cisco platforms and their equivalent Huawei CLI command set. This is quite helpful especially for Cisco networking Engineers who face challenge when introduced to a new Huawei device for configuration –
HUAWEI CISCO EQUIVALENT COMMANDS – CLI CHEATSHEET
CISCO COMMAND HUAWEI COMMAND
ping Ping
traceroute tracert
show display
show interfaces display interface
Show ip route display ip routing-table
Show ip interface display ip interface
Show version Display version
Show ip bgp Display bgp routing-table
Show clock Display clock
Show port Display port-mapping
Show flash dir flash: (on user view mode)
Show logging Display logbuffer
Show snmp Display snmp-agent statistics
Show frame-relay pvc Display fr pvc-info
Show users Display users
Show terminal length screen-length disableundo screen-length disable
enable super
disable Super 0 (number is privilege level from 0 to 3,where 3 is default and 
equivalent to “enable” on Cisco)
Conf t System-view
exit quit
end return
Show policy-map interface Display qos policy interface
send send (on user view mode)
clear counters reset counters interface
clear interface reset counters interface
terminal monitor terminal monitor
terminal length screen-length disableundo screen-length disable
terminal no monitor undo terminal monitor
Show tech display diagnostic-information
Show Controller display controller (but not relevant for non-modularchassis)
show dsl int atm 0 display dsl status interface Atm 2/0
sho crypto isakmp sa Display ike sa
sho crypto isakmp key Display ike peer
CISCO COMMAND HUAWEI COMMAND
write terminal (sh run) display current-configuration
Sh startup Display saved-configuration
[no equivalent: shows the files used for 
startup]
Display startup
Write erase Reset saved-configuration
Write mem (or wr or copy run start) save
clear counters reset (on user view mode)Reset counters interface
telnet telnet
Enable secret (conf mode) Super pass cipher (system mode)
Term mon term debu
clock clock
no undo
debug / no debug debugging / undo debugging
copy running-config Save safely
clear crypto ipsec saike sa
clear access-list counters reset acl counter all
reload reboot
shutdown shutdown
boot bootrom
Aaa hwtacacs scheme
terminal no monitor undo terminal monitor
tacacs-server hwtacacs scheme (in conf command)
snmp-server tftp-server (in conf command)
router bgp Bgp
Router rip Rip
ip tacacs hwtacacs nas-ip (this command doesn’t exist !!!)
mtu Mtu (this command doesn’t exist !!!)
clear ip cef reset ip fast-forwarding
clear ip route * reset ip routing-table statistics protocol all
Clear ip bgp Reset bgp all
Sh ip nat translation Display nat session
sho atm pvc Display atm pvc-info
debug pvc nego Debug atm all (very dangerous – might crash router)
sho crypto isakmp police Display ike proposal
