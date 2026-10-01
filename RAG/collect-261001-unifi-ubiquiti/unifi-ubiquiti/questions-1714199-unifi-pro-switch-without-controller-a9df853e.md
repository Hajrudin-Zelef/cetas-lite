---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/questions-1714199-unifi-pro-switch-without-controller-a9df853e
title: "Unifi pro switch without controller"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-unifi-ubiquiti/questions-1714199-unifi-pro-switch-without-controller-a9df853e.md
source_anchor: ""
source_lines: [1, 10]
sha256: a04ea65b32903826c1b67b202c8d8b09b16c76dffd10295bc58971c51ae10c35
---

# Unifi pro switch without controller

*Score : 3 | Source : https://superuser.com/questions/1714199/unifi-pro-switch-without-controller*

We have a UniFi Switch Pro 48 PoE and we want to use it without a unifi controller but when we connect to the switch and change the configuration it is not saved, after a reboot the configuration is restored to default
Enter command for save configuration in cli
write memory:
copy flash:running-config nvram:startup-config
in shell mode:
save
