---
id: collect-260926-mikrotik/mikrotik/should-i-upgrade-routerboot-on-each-routeros-upgrade
title: "should-i-upgrade-routerboot-on-each-routeros-upgrade"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/RouterOS/should-i-upgrade-routerboot-on-each-routeros-upgrade.md
source_anchor: ""
source_lines: [1, 18]
sha256: 3ec08c9a9626188babb8583d773e36c20b7fb7725f3b4d73ff9812a6a746c1e4
---

# should-i-upgrade-routerboot-on-each-routeros-upgrade

I’ve avoided RouterBOOT firmware problems using this:

```
/system routerboard settings
set auto-upgrade=yes
```

Problem solved with **double reboots** after each ROS upgrade:

```
/system/routerboard> print
routerboard: yes
factory-firmware: 6.45.9
current-firmware: 7.11.2
upgrade-firmware: 7.11.2
```

MikroTik should make this single reboot process if auto-upgrade is enabled.
